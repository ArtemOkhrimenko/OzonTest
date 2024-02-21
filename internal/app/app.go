package app

import "context"

//go:generate mockgen -source=app.go -destination mock.app_test.go -package app_test

type (
	Repo interface {
		SaveURL(ctx context.Context, url string) (int, error)
		GetURL(ctx context.Context, id int) (string, error)
		GetURLbyName(ctx context.Context, link string) (int, error)
	}
	Decoder interface {
		Encode(id int) (string, error)
		Decode(id string) int
	}
)

type App struct {
	repo    Repo
	decoder Decoder
}

func New(repo Repo, decoder Decoder) *App {
	return &App{
		repo:    repo,
		decoder: decoder,
	}
}
