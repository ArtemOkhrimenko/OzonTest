package serve

import (
	"context"
	"fmt"
	"net"

	zip_proto "OzonTest/api/zip"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// GRPCGateWay starts HTTP-proxy server for gRPC serer, for using gRPC endpoints from WEB.
func GRPCGateWay(log *zap.Logger, host string, port uint16, grpcPort uint16) func(context.Context) error {
	return func(ctx context.Context) error {

		conn, err := grpc.DialContext(ctx, net.JoinHostPort(host, fmt.Sprintf("%d", grpcPort)),
			grpc.WithBlock(),
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)
		if err != nil {
			return fmt.Errorf("grpc_helper.Dial: %w", err)
		}

		gwmux := runtime.NewServeMux()
		// Register Greeter
		err = zip_proto.RegisterZipHandler(ctx, gwmux, conn)
		if err != nil {
			return err
		}

		return HTTP(log, host, port, gwmux)(ctx)
	}
}
