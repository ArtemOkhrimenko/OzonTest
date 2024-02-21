package main

import (
	"context"
	"fmt"

	"OzonTest/internal/adapters/repo"
	"OzonTest/internal/api"
	"OzonTest/internal/app"
	"OzonTest/internal/config"
	"OzonTest/internal/decoder"
	"OzonTest/serve"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {
	ctx := context.Background()

	fmt.Println("start server")

	log, err := zap.NewDevelopment()
	if err != nil {
		panic(fmt.Sprintf("zap.NewDevelopmentt: %s", err.Error()))
	}

	err = godotenv.Load("/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg, err := config.LoadConfig()
	if err != nil {
		panic(fmt.Sprintf("config.LoadConfig(): %s", err.Error()))
	}

	db, err := sqlx.ConnectContext(ctx, "pgx", cfg.DSN)
	if err != nil {
		panic(fmt.Sprintf("sqlx.Connect: %s", err.Error()))
	}

	dec, err := decoder.New(cfg.MinLen, cfg.Alphabet)
	if err != nil {
		panic("can not initialize decoder")
	}

	repository := repo.New(db)
	application := app.New(repository, dec)
	server := api.New(application)

	err = serve.Start(
		ctx,
		serve.GRPC(log.With(zap.String("module", "gRPC")), cfg.Host, cfg.GrpcPort, server),
		serve.GRPCGateWay(log.With(zap.String("module", "gRPC-Gateway")), cfg.Host, cfg.HttpPort, cfg.GrpcPort),
	)
	if err != nil {
		panic(fmt.Sprintf("serve.Start %v", err))
	}
}
