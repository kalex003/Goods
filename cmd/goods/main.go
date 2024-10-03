package main

import (
	"Goods/internal/app"
	"Goods/internal/config"
	storage "Goods/internal/storage/postgres"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

/*var ConnString = MustGetEnv("DATABASE_URL")

func MustGetEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Missing required environment variable: %s", key)
	}

	return value
}*/

var ConnString = "postgres://postgres:postgres@db:5432/Goods?sslmode=disable"

const createPartitionSQL = `
DO
$$
BEGIN
EXECUTE format('CREATE TABLE "goods.goodslog' || '[' || DATE(NOW()::TIMESTAMP)::TEXT || ']"' || ' PARTITION OF goods.goodslog FOR VALUES FROM (''' || DATE_TRUNC('day', NOW()::TIMESTAMP)::TEXT || ''') TO ('''|| (DATE_TRUNC('day', NOW()::TIMESTAMP) + INTERVAL '3 minutes')::TEXT  || ''');');
END
$$;
`

func worker(log *slog.Logger, db *storage.GoodsDb) {
	for {
		log.Info("Запуск создания артиции")

		_, err := db.Db.Exec(createPartitionSQL)
		if err != nil {
			log.Info("Ошибка при создании партиции: %v\n", err)
		} else {
			log.Info("Партиция успешно создана.")
		}

		// Ожидание 10 минут перед следующим запуском
		time.Sleep(10 * time.Minute)
	}
}

func main() {
	cfg := config.MustLoad()
	fmt.Println(cfg)

	log := setupLogger(cfg.Env)

	log.Info("Starting application",
		slog.String("env", cfg.Env),
		slog.Any("env", cfg),
		slog.Int("port", cfg.GRPC.Port),
	)

	application, db := app.New(log, cfg.GRPC.Port, ConnString) //добавил для крончика

	go application.GRPCServer.MustRun() //в параллельном режиме от остальной программы обрабатываем запросы

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT) //обработка заверщаюего сигнала (одного из двух сигтерм или сигинт)

	go worker(log, db)

	<-stop //пока в этот канал что-то не придет (сигнал о зщавершении программы), мы тут просто будем висеть и ждать, а сверху будет работать го рутина

	application.GRPCServer.Stop()

	log.Info("Application stopped")
}

func setupLogger(env string) *slog.Logger { //сложно чот
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}
