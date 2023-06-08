package repo

import "context"

type Repo interface {
	AddShortURL(ctx context.Context, url string, shortURL string) error
	GetURL(ctx context.Context, shortURL string) (string, error)
}
