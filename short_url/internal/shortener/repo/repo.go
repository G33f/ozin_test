package repo

import (
	"ShortURL/internal/logging"
	"ShortURL/internal/shortener"
	"ShortURL/internal/storage"
	"ShortURL/internal/utils"
	"context"
	"fmt"
)

const repoError = "repo error: "

const notFind = "no rows in result set"

type repo struct {
	client storage.Client
	logger *logging.Logger
}

func (r *repo) AddShortURL(ctx context.Context, url string, shortURL string) error {
	q := `insert into urls (url, short_url)
			  values ($1, $2)
			  returning urls.id;`
	q = utils.FormatQuery(q)
	var id int64
	err := r.client.QueryRow(ctx, q, url, shortURL).Scan(&id)
	if err != nil {
		return err
	}
	return nil
}

func (r *repo) GetURL(ctx context.Context, shortURL string) (string, error) {
	var url string
	q := `select urls.url from urls
			  where urls.short_url = $1;`
	q = utils.FormatQuery(q)
	err := r.client.QueryRow(ctx, q, shortURL).Scan(&url)
	if err == fmt.Errorf(notFind) {
		return "", nil
	}
	return url, err
}

func NewRepo(client storage.Client, log *logging.Logger) shortener.Repo {
	return &repo{
		client: client,
		logger: log,
	}
}
