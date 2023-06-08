package usecase

import "context"

type UseCase interface {
	MakeURLShorter(ctx context.Context, url string) (string, error)
	GetOriginalURL(ctx context.Context, url string) (string, error)
}
