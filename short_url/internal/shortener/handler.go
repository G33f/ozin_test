package shortener

import (
	"ShortURL/internal/shortener/model"
	"context"
)

type Handler interface {
	MakeURLShorter(ctx context.Context, req *model.CrateRequest) (*model.CrateResponse, error)
	GetOriginalURL(ctx context.Context, in *model.GetRequest) (*model.GetResponse, error)
}
