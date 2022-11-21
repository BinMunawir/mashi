package services

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/BinMunawir/mashi/src/delivery/infra/configs"
	"github.com/segmentio/kafka-go"
)

type MessageBus interface {
	Produce(m []byte, topic string) error
	Consume(topic string) ([]byte, error)
}

type Kafka struct{}

func NewKafka() MessageBus {
	return Kafka{}
}

func (k Kafka) Consume(topic string) ([]byte, error) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        []string{configs.KAFKA_HOST},
		Topic:          topic,
		Partition:      0,
		GroupID:        "main",
		MinBytes:       10e3, // 10KB
		MaxBytes:       1e6,  // 1MB
		SessionTimeout: time.Second * 10,
	})

	m, err := r.ReadMessage(context.Background())
	if err != nil {
		log.Println(err)
		return nil, nil
	}
	fmt.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))

	r.Close()
	return m.Value, nil
}

func (k Kafka) Produce(m []byte, topic string) error {
	conn, err := kafka.DialLeader(context.Background(), "tcp", configs.KAFKA_HOST, topic, 0)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	_, err = conn.WriteMessages(kafka.Message{Value: m})
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	if err := conn.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}

	return nil
}
