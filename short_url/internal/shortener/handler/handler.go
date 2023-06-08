package handler

import (
	"ShortURL/internal/logging"
	"ShortURL/internal/shortener"
	"ShortURL/internal/shortener/model"
	"context"
)

type handler struct {
	useCase shortener.UseCase
	log     *logging.Logger
}

func (h *handler) MakeURLShorter(ctx context.Context, req *model.CrateRequest) (*model.CrateResponse, error) {
	shorter, err := h.useCase.MakeURLShorter(ctx, req.Url)
	if err != nil {
		return nil, err
	}
	return &model.CrateResponse{Url: shorter}, nil
}

func (h *handler) GetOriginalURL(ctx context.Context, in *model.GetRequest) (*model.GetResponse, error) {
	shorter, err := h.useCase.GetOriginalURL(ctx, in.Url)
	if err != nil {
		return nil, err
	}
	return &model.GetResponse{Url: shorter}, nil
}

func NewHandler(useCase shortener.UseCase, log *logging.Logger) shortener.Handler {
	return &handler{
		useCase: useCase,
		log:     log,
	}
}
