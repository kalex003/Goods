package producers

import (
	kafkamodels "Goods/internal/kafka/models"
	"context"
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"log/slog"
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

func Initkafka(ctx context.Context, Log *slog.Logger) {

	err := createTopic()
	if err != nil {
		Log.Error("Ошибка при создании топика: %v", err)
	} else {
		Log.Info("Топик успешно создан.")
	}

	// Инициализация Kafka writer
	Writer := &kafka.Writer{
		Addr:     kafka.TCP(kafkaBroker),
		Topic:    kafkaTopic,
		Balancer: &kafka.LeastBytes{},
	}

	<-ctx.Done()

	err = Writer.Close()
	if err != nil {
		Log.Error("Ошибка при закрытии writer: %v", err)
	} else {
		Log.Info("Kafka writer закрыт по сигналу контекста")
	}

}

// processGoodsLog — функция для извлечения данных из таблицы и отправки их в Kafka
func SendMessage(message kafkamodels.KafkaModel, msgctx context.Context, log *slog.Logger, writer kafka.Writer) error {

	// Сериализация данных в JSON
	jsonmessage, err := json.Marshal(message)
	if err != nil {
		log.Error("Ошибка при сериализации данных: %v", err)
		return err
	}

	// Отправка сообщения в Kafka с контекстом
	err = writer.WriteMessages(msgctx, kafka.Message{
		Key:   []byte(fmt.Sprintf("%s", message.LogDt)),
		Value: jsonmessage,
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

	return nil
}
