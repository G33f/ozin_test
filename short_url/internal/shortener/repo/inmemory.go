package repo

import (
	"ShortURL/internal/logging"
	"ShortURL/internal/shortener"
	"context"
	"sync"
)

type inMemory struct {
	uRLHash map[string]string
	mut     *sync.RWMutex

	logger *logging.Logger
}

func (i *inMemory) AddShortURL(ctx context.Context, url string, shortURL string) error {
	i.mut.Lock()
	i.uRLHash[shortURL] = url
	i.mut.Unlock()
	return nil
}

func (i *inMemory) GetURL(ctx context.Context, shortURL string) (string, error) {
	i.mut.RLock()
	val, ok := i.uRLHash[shortURL]
	i.mut.RUnlock()
	if !ok {
		return "", nil
	}
	return val, nil
}

func NewInMemoryRepo(log *logging.Logger, mut *sync.RWMutex) shortener.Repo {
	return &inMemory{
		uRLHash: map[string]string{},
		logger:  log,
		mut:     mut,
	}
}
