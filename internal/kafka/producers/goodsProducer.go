package producers

import (
	"Goods/internal/kafka/models"
	storage "Goods/internal/storage/postgres"
	"context"
	"encoding/json"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/segmentio/kafka-go"
	"log"
	"log/slog"
	"time"
)

var (
	kafkaBroker      = "kafka:9092"
	kafkaTopic       = "new_goods_topic."
	kafkaPartition   = 1
	kafkaReplication = 1
)

func createTopic() error {
	// Подключение к брокеру
	conn, err := kafka.Dial("tcp", kafkaBroker)
	if err != nil {
		return fmt.Errorf("не удалось подключиться к Kafka: %w", err)
	}
	defer conn.Close()

	// Проверка, существует ли топик
	topics, err := conn.ReadPartitions(kafkaTopic)
	if err == nil && len(topics) > 0 {
		log.Printf("Топик %s уже существует\n", kafkaTopic)
		return nil
	}

	// Создание нового топика
	topicConfig := kafka.TopicConfig{
		Topic:             kafkaTopic,
		NumPartitions:     kafkaPartition,
		ReplicationFactor: kafkaReplication,
	}

	err = conn.CreateTopics(topicConfig)
	if err != nil {
		return fmt.Errorf("не удалось создать топик: %w", err)
	}

	return nil
}

// Worker — основной процесс отправки сообщений в Kafka
func Worker(ctx context.Context, log *slog.Logger, db *storage.GoodsDb) {
	for {
		select {
		case <-ctx.Done():
			log.Info("Kafka producer остановлен по сигналу контекста")
			return
		default:
			// Создаем контекст с таймаутом для каждой отправки сообщения
			msgCtx, cancel := context.WithTimeout(ctx, 10*time.Second)

			// Создание топика (если необходимо)
			err := createTopic()
			if err != nil {
				log.Error("Ошибка при создании топика: %v", err)
			} else {
				log.Info("Топик успешно создан.")
			}

			// Отправка данных в Kafka
			err = processGoodsLog(msgCtx, log, db)
			if err != nil {
				log.Error("Ошибка при отправке данных в Kafka: %v", err)
			} else {
				log.Info("Данные успешно отправлены в Kafka.")
			}

			// Завершение контекста после выполнения или по таймауту
			cancel()

			// Ожидание 3 минуты до следующей обработки
			time.Sleep(3 * time.Minute)
		}
	}
}

// processGoodsLog — функция для извлечения данных из таблицы и отправки их в Kafka
func processGoodsLog(ctx context.Context, log *slog.Logger, db *storage.GoodsDb) error {
	// Извлечение данных из таблицы goodslog
	/*	rows, err := db.Db.Query(ctx, `
	    SELECT gl.log_dt, gl.goods_id, gl.place_id, gl.sku_id, gl.wbsticker_id, gl.barcode,
	           gl.state_id, gl.ch_employee_id, gl.office_id, gl.wh_id, gl.tare_id, gl.tare_type,
	           gl.ch_dt, gl.is_del
	    FROM goods.goodslog gl
	    WHERE gl.log_dt >= NOW() AT TIME ZONE 'Europe/Moscow' - INTERVAL '3 minutes'
	`)*/

	sql, _, err := squirrel.StatementBuilder.
		Select("gl.log_dt", "gl.goods_id", "gl.place_id", "gl.sku_id", "gl.wbsticker_id", "gl.barcode", "gl.state_id", "gl.ch_employee_id", "gl.office_id", "gl.wh_id", "gl.tare_id", "gl.tare_type", "gl.ch_dt", "gl.is_del").
		From("goods.goodslog AS gl").
		Where("gl.log_dt >= NOW() AT TIME ZONE 'Europe/Moscow' - INTERVAL '3 minutes'").
		PlaceholderFormat(squirrel.Dollar).
		ToSql()

	if err != nil {
		return fmt.Errorf("ProcessGoodsLog: unable to build query: %w", err)
	}

	rows, err := db.Db.Query(ctx, sql)
	if err != nil {
		return fmt.Errorf("SelectGoodsByTare: %w", err)
	}

	if err != nil {
		return fmt.Errorf("ошибка при выполнении запроса к базе данных: %w", err)
	}

	defer rows.Close()

	// Инициализация Kafka writer
	writer := &kafka.Writer{
		Addr:     kafka.TCP(kafkaBroker),
		Topic:    kafkaTopic,
		Balancer: &kafka.LeastBytes{},
	}
	defer writer.Close()

	rowsProcessed := 0
	for rows.Next() {
		rowsProcessed++
	}

	if err := rows.Err(); err != nil {
		return err
	}

	if rowsProcessed == 0 {
		return fmt.Errorf("Запрос не вернул данных, отправлять нечего")
	}

	for rows.Next() {
		var logEntry models.KafkaModel

		if err := rows.Scan(&logEntry.LogDt, &logEntry.GoodsId, &logEntry.PlaceId, &logEntry.SkuId, &logEntry.WbstickerId, &logEntry.Barcode, &logEntry.StateId, &logEntry.ChEmployeeId, &logEntry.OfficeId, &logEntry.WhId, &logEntry.TareId, &logEntry.TareType, &logEntry.ChDt, &logEntry.IsDel); err != nil {
			log.Error("Ошибка при чтении данных из таблицы: %v", err)
			continue
		}

		// Сериализация данных в JSON
		message, err := json.Marshal(logEntry)
		if err != nil {
			log.Error("Ошибка при сериализации данных: %v", err)
			continue
		}

		// Отправка сообщения в Kafka с контекстом
		err = writer.WriteMessages(ctx, kafka.Message{
			Key:   []byte(fmt.Sprintf("%s", logEntry.LogDt)),
			Value: message,
		})

		if err != nil {
			if err == context.DeadlineExceeded {
				log.Error("Контекст завершен по таймауту при отправке сообщения")
			} else {
				log.Error("Ошибка при отправке сообщения в Kafka: %v", err)
			}
		} else {
			log.Info("Сообщение успешно отправлено в Kafka: %s", message)
		}

		// Проверяем, не завершен ли контекст досрочно
		select {
		case <-ctx.Done():
			log.Info("Операция отправки данных в Kafka была отменена")
			return ctx.Err()
		default:
			// Продолжаем работу
		}
	}

	return nil
}
