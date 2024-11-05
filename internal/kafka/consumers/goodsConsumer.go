package consumers

import (
	"Goods/internal/kafka/models"
	storage "Goods/internal/storage/postgres"
	"context"
	"encoding/json"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/segmentio/kafka-go"
	"log/slog"
	"time"
)

var (
	kafkaBroker = "kafka:9092"
	kafkaTopic  = "new_goods_topic."
	groupID     = "consumer-group-1"
)

// Worker — основной процесс чтения сообщений из Kafka и записи их в базу данных
func Worker(ctx context.Context, log *slog.Logger, db *storage.GoodsDb) {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{kafkaBroker},
		Topic:    kafkaTopic,
		GroupID:  groupID,
		MinBytes: 1e3,  // 1KB
		MaxBytes: 10e6, // 10MB
	})

	defer func() {
		if err := reader.Close(); err != nil {
			log.Error("Ошибка при закрытии читателя Kafka", err)
		}
	}()

	for {
		select {
		case <-ctx.Done():
			log.Info("Kafka consumer остановлен по сигналу контекста")
			return
		default:
			// Создаем контекст с таймаутом 30 секунд для каждого сообщения

			// Чтение сообщения из Kafka
			msg, err := reader.ReadMessage(ctx)
			if err != nil {
				if err == context.DeadlineExceeded {
					log.Error("Таймаут при чтении сообщения из Kafka")
				} else {
					log.Error("Ошибка при чтении сообщения из Kafka: %v", err)
				}
				continue
			}

			log.Info("Получено сообщение из Kafka: %s", msg.Value)

			msgCtx, cancel := context.WithTimeout(ctx, 30*time.Second)

			// Обработка сообщения и запись в базу данных
			err = processMessage(msgCtx, log, db, msg.Value)
			if err != nil {
				log.Error("Ошибка при обработке сообщения: %v", err)
			}

			// Завершение контекста после обработки сообщения
			cancel()
		}
	}
}

// processMessage — функция для обработки сообщения и записи его в базу данных
func processMessage(ctx context.Context, log *slog.Logger, db *storage.GoodsDb, message []byte) error {
	var importedGood models.KafkaModel

	// Десериализация сообщения из JSON
	err := json.Unmarshal(message, &importedGood)
	if err != nil {
		return fmt.Errorf("ошибка при десериализации сообщения: %w", err)
	}

	// SQL-запрос для вставки данных в таблицу imported_goods
	/*	query := `
			INSERT INTO import.goods_imported
			(log_dt, goods_id, place_id, sku_id, wbsticker_id, barcode, state_id, ch_employee_id, office_id, wh_id, tare_id, tare_type, ch_dt, is_del)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)
		`

		// Выполнение SQL-запроса с использованием контекста
		_, err = db.Db.Exec(ctx, query,
			importedGood.LogDt, importedGood.GoodsId, importedGood.PlaceId, importedGood.SkuId,
			importedGood.WbstickerId, importedGood.Barcode, importedGood.StateId, importedGood.ChEmployeeId,
			importedGood.OfficeId, importedGood.WhId, importedGood.TareId, importedGood.TareType,
			importedGood.ChDt, importedGood.IsDel,
		)*/

	sql, args, err := squirrel.StatementBuilder.
		Insert("").
		Into("import.goods_imported AS g").
		Columns("log_dt", "goods_id", "place_id", "sku_id", "wbsticker_id", "barcode", "state_id", "ch_employee_id", "office_id", "wh_id", "tare_id", "tare_type", "ch_dt", "is_del").
		Values(importedGood.LogDt, importedGood.GoodsId, importedGood.PlaceId, importedGood.SkuId,
			importedGood.WbstickerId, importedGood.Barcode, importedGood.StateId, importedGood.ChEmployeeId,
			importedGood.OfficeId, importedGood.WhId, importedGood.TareId, importedGood.TareType,
			importedGood.ChDt, importedGood.IsDel).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()

	// Выполняем запрос
	_, err = db.Db.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("ошибка при вставке данных в базу: %w", err)
	}

	log.Info("Данные успешно записаны в базу данных для товара ID: %d", importedGood.GoodsId)
	return nil
}
