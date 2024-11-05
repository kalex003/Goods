package partitionMaker

import (
	storage "Goods/internal/storage/postgres"
	"context"
	"log/slog"
	"time"
)

const createPartitionSQL = `
DO
$$
BEGIN
SET TIME ZONE 'Europe/Moscow';
EXECUTE format('CREATE TABLE "goods.goodslog' || '[' || DATE_TRUNC('MINUTE', NOW()::TIMESTAMP)::TEXT || ']"' || ' PARTITION OF goods.goodslog FOR VALUES FROM (''' || DATE_TRUNC('MINUTE', NOW()::TIMESTAMP)::TEXT || ''') TO ('''|| (DATE_TRUNC('MINUTE', NOW()::TIMESTAMP) + INTERVAL '5 minutes')::TEXT  || ''');');
END
$$;`

func Worker(ctx context.Context, log *slog.Logger, db *storage.GoodsDb) {
	for {
		log.Info("Запуск создания артиции")

		scriptCtx, cancel := context.WithTimeout(ctx, 10*time.Second)

		_, err := db.Db.Exec(scriptCtx, createPartitionSQL)
		if err != nil {
			log.Info("Ошибка при создании партиции: %v\n", err)
		} else {
			log.Info("Партиция успешно создана.")
		}

		cancel()

		// Ожидание 4 минут перед следующим запуском
		time.Sleep(4 * time.Minute)
	}
}
