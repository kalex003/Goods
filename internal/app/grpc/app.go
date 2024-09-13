package grpcapp

import (
	grpcgoods "Goods/internal/grpc/goods"
	"fmt"
	"google.golang.org/grpc"
	"log/slog"
	"net"
)

type App struct {
	log        *slog.Logger
	grpcServer *grpc.Server
	port       int
}

func New(log *slog.Logger, goodsService grpcgoods.Goods, port int) *App {
	grpcServer := grpc.NewServer()

	grpcgoods.Register(grpcServer, goodsService)
	return &App{
		log:        log,
		grpcServer: grpcServer,
		port:       port,
	}
}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *App) Run() error {
	const op = "grpcapp.Run"

	log := a.log.With(
		slog.String("op", "op"),
		slog.String("port", string(rune(a.port))),
	)

	addr, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	log.Info("grpc server is running", slog.String("addr", addr.Addr().String()))

	if err := a.grpcServer.Serve(addr); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
func (a *App) Stop() {
	const op = "grpcapp.Stop"
	a.log.With(slog.String("op", "op")).Info("grpc server is stopping", slog.Int("addr", a.port))
	a.grpcServer.GracefulStop()
}
