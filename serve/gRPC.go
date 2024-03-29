package serve

import (
	"context"
	"fmt"
	"net"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// GRPC starts gRPC server on addr, logged as service.
// It runs until failed or ctx.Done.
func GRPC(log *zap.Logger, host string, port uint16, srv *grpc.Server) func(context.Context) error {
	return func(ctx context.Context) error {
		ln, err := net.Listen("tcp", net.JoinHostPort(host, fmt.Sprintf("%d", port)))
		if err != nil {
			return fmt.Errorf("net.Listen: %w", err)
		}

		errc := make(chan error, 1)
		go func() { errc <- srv.Serve(ln) }()
		log.Info("started", zap.String("host", host), zap.Uint16("port", port))
		defer log.Info("shutdown")

		select {
		case err = <-errc:
		case <-ctx.Done():
			srv.GracefulStop() // It will not interrupt streaming.
		}
		if err != nil {
			return fmt.Errorf("srv.Serve: %w", err)
		}

		return nil
	}
}
