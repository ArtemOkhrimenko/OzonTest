package api

import (
	"context"
	"fmt"

	zip_proto "OzonTest/api/zip"
)

func (a *api) GenerateShortLink(ctx context.Context, req *zip_proto.GenerateShortLinkRequest) (*zip_proto.GenerateShortLinkResponse, error) {

	shortLink, err := a.app.GetShortUrl(ctx, req.LongLink)
	if err != nil {
		return nil, fmt.Errorf("a.app.GetShortUrl: %w", err)
	}

	return &zip_proto.GenerateShortLinkResponse{ShortLink: shortLink}, nil
}

func (a *api) GetLongLink(ctx context.Context, req *zip_proto.GetLongLinkRequest) (*zip_proto.GetLongLinkResponse, error) {

	longLink, err := a.app.GetLongURL(ctx, req.ShortLink)
	if err != nil {
		return nil, fmt.Errorf("a.app.GetLongURL: %w", err)
	}

	return &zip_proto.GetLongLinkResponse{LongLink: longLink}, nil
}
