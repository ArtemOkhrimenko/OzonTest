package api

import (
	"context"

	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"google.golang.org/grpc"

	zip_proto "OzonTest/api/zip"
)

type application interface {
	GetShortUrl(ctx context.Context, longURL string) (string, error)
	GetLongURL(ctx context.Context, shortURL string) (string, error)
}

type api struct {
	app application
}

func New(app application) *grpc.Server {
	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			grpc_recovery.UnaryServerInterceptor(),
		),
	)

	zip_proto.RegisterZipServer(grpcServer, &api{
		app: app,
	})

	return grpcServer
}
