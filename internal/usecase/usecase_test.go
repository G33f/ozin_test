package usecase

import (
	mock_shortener "ShortURL/mocks/mock_repo"
	"context"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
	"testing"
)

var (
	ctx = context.Background()
)

func Test_useCase_GetOriginalURL(t *testing.T) {
	tt := []struct {
		name    string
		uc      useCase
		testUrl string
		expStr  string
		expErr  error
		wantErr bool
	}{
		{
			name:    "Normal case",
			uc:      useCase{},
			testUrl: "Gj5SeHClQi",
			expStr:  "youtube.com12",
			expErr:  nil,
			wantErr: false,
		},
		{
			name:    "no error 2",
			uc:      useCase{},
			testUrl: "Gj5SeHClQi",
			expStr:  "",
			expErr:  fmt.Errorf("URL does not exist"),
			wantErr: true,
		},
	}

	ctrl := gomock.NewController(t)
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			mcRepo := mock_shortener.NewMockRepo(ctrl)
			uc := &useCase{
				repo:   mcRepo,
				logger: nil,
			}
			mcRepo.EXPECT().GetURL(ctx, tc.testUrl).Return(tc.expStr, nil)
			got, err := uc.GetOriginalURL(ctx, tc.testUrl)
			assert.Equal(t, tc.wantErr, err != nil)
			assert.Equal(t, tc.expStr, got)
			assert.Equal(t, tc.expErr, err)
		})
	}
}

func Test_useCase_MakeURLShorter(t *testing.T) {
	tt := []struct {
		name      string
		uc        useCase
		testUrl   string
		expGetUrl string
		shortURL  string
		expStr    string
		expErr    error
		wantErr   bool
	}{
		{
			name:      "Short URL dose not exist",
			uc:        useCase{},
			testUrl:   "youtube.com12",
			expGetUrl: "",
			shortURL:  "Gj5SeHClQi",
			expStr:    "Gj5SeHClQi",
			expErr:    nil,
			wantErr:   false,
		},
		{
			name:      "short URL is exist",
			uc:        useCase{},
			testUrl:   "youtube.com12",
			expGetUrl: "youtube.com12",
			shortURL:  "Gj5SeHClQi",
			expStr:    "Gj5SeHClQi",
			expErr:    nil,
			wantErr:   false,
		},
		{
			name:      "exist URL white same short URL",
			uc:        useCase{},
			testUrl:   "youtube.com12",
			expGetUrl: "youtube.com121",
			shortURL:  "Gj5SeHClQi",
			expStr:    "0Y5Qr2ByRl",
			expErr:    nil,
			wantErr:   false,
		},
		{
			name:      "exist URL white same short URL",
			uc:        useCase{},
			testUrl:   "youtube.com12",
			expGetUrl: "youtube.com12",
			shortURL:  "Gj5SeHClQi",
			expStr:    "",
			expErr:    fmt.Errorf("sql error: some"),
			wantErr:   true,
		},
	}
	ctrl := gomock.NewController(t)
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			mcRepo := mock_shortener.NewMockRepo(ctrl)
			uc := &useCase{
				repo:   mcRepo,
				logger: nil,
			}
			mcRepo.EXPECT().GetURL(ctx, tc.shortURL).Return(tc.expGetUrl, tc.expErr)
			mcRepo.EXPECT().GetURL(ctx, tc.expStr).Return("", tc.expErr)
			mcRepo.EXPECT().AddShortURL(ctx, tc.testUrl, tc.expStr).Return(nil)
			got, err := uc.MakeURLShorter(ctx, tc.testUrl)
			assert.Equal(t, tc.wantErr, err != nil)
			assert.Equal(t, tc.expStr, got)
			assert.Equal(t, tc.expErr, err)
		})
	}
}
