package app

import (
	grpcapp "Goods/internal/app/grpc"
	"Goods/internal/services/goods"
	"Goods/internal/storage/postgres"
	storage "Goods/internal/storage/postgres"
	"log/slog"
)

type App struct {
	GRPCServer *grpcapp.App
}

func New(
	log *slog.Logger,
	grpcport int,
	ConnString string, // чет сложно, тут надо разбираться до конца нормально
) (*App, *storage.GoodsDb) {

	GoodsDb, err := postgres.New(ConnString)

	if err != nil {
		panic(err)
	}

	GoodsService := goods.New(log, GoodsDb, GoodsDb, GoodsDb, GoodsDb)

	grpcApp := grpcapp.New(log, GoodsService, grpcport)

	return &App{
		GRPCServer: grpcApp,
	}, GoodsDb
}
