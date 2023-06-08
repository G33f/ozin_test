package shortener

import (
	"ShortURL/internal/shortener/model"
	"context"
)

type Handler interface {
	MakeURLShorter(ctx context.Context, req *model.URL) (*model.URL, error)
	GetOriginalURL(ctx context.Context, in *model.URL) (*model.URL, error)
}
