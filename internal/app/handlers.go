package app

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
)

func (a *App) GetShortUrl(ctx context.Context, longURL string) (string, error) {
	validURL := strings.SplitAfterN(longURL, "/", 2)

	idURL, err := a.repo.GetURLbyName(ctx, longURL)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		id, err := a.repo.SaveURL(ctx, longURL)
		if err != nil {
			return "", fmt.Errorf("a.repo.SaveURL: %w", err)
		}

		encodeID, err := a.decoder.Encode(id)
		if err != nil {
			return "", fmt.Errorf("a.decoder.Encode: %w", err)
		}

		return validURL[0] + encodeID, nil
	case err != nil:
		return "", fmt.Errorf("a.repo.GetURLbyName: %w", err)
	}

	Id, err := a.decoder.Encode(idURL)
	if err != nil {
		return validURL[0] + Id, fmt.Errorf("a.decoder.Encode: %w", err)
	}

	return validURL[0] + Id, nil
}

func (a *App) GetLongURL(ctx context.Context, shortURL string) (string, error) {
	validURL := strings.SplitAfterN(shortURL, "/", 2)

	url := a.decoder.Decode(validURL[1])

	res, err := a.repo.GetURL(ctx, url)
	if err != nil {
		return "", fmt.Errorf("a.repo.GetURL: %w", err)
	}

	return validURL[0] + res, nil
}
