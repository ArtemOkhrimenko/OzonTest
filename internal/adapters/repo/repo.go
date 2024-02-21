package repo

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Repo struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Repo {
	return &Repo{db: db}
}

func (r *Repo) SaveURL(ctx context.Context, url string) (int, error) {
	const query = `insert into
						urls
							(url)
						values
							($1)
						returning id`

	var id int

	err := r.db.GetContext(ctx, &id, query, url)
	if err != nil {
		return 0, fmt.Errorf("db.GetContext: %w", err)
	}

	return id, nil
}

func (r *Repo) GetURL(ctx context.Context, id int) (string, error) {
	const query = `select * from urls where id = $1;`

	var url Url
	err := r.db.GetContext(ctx, &url, query, id)
	if err != nil {
		return "", fmt.Errorf("db.SelectContext: %w", err)
	}

	return url.Url, nil
}

func (r *Repo) GetURLbyName(ctx context.Context, link string) (int, error) {
	const query = `select * from urls where url = $1;`

	var url Url
	err := r.db.GetContext(ctx, &url, query, link)
	if err != nil {
		return 0, fmt.Errorf("db.GetContext: %w", err)
	}

	return url.ID, nil
}
