package repo

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
)

type postgres struct {
	client *pgxpool.Pool
}

func (r *postgres) AddShortURL(ctx context.Context, url string, shortURL string) error {
	q := `insert into urls (url, short_url)
			  values ($1, $2)
			  returning urls.id;`
	var id int64
	err := r.client.QueryRow(ctx, q, url, shortURL).Scan(&id)
	if err != nil {
		return err
	}
	return nil
}

func (r *postgres) GetURL(ctx context.Context, shortURL string) (string, error) {
	var url string
	q := `select urls.url from urls
			  where urls.short_url = $1;`
	err := r.client.QueryRow(ctx, q, shortURL).Scan(&url)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return "", nil
		}
		return "", err
	}
	return url, nil
}

func NewPostgresRepo(client *pgxpool.Pool) Repo {
	return &postgres{
		client: client,
	}
}
