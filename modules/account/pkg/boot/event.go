package boot

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

type KafkaConfig struct {
	Broker  string
	GroupId string
}

type EventBus struct {
	config *KafkaConfig
}

func NewEventBus(config *KafkaConfig) *EventBus {
	return &EventBus{
		config: config,
	}
}

func (s *EventBus) NewProducer(topic string) *EventProducer {
	return &EventProducer{
		writer: kafka.NewWriter(kafka.WriterConfig{
			Brokers:  []string{s.config.Broker},
			Topic:    topic,
			Balancer: &kafka.LeastBytes{},
		}),
	}
}

func (s *EventBus) NewConsumer(topic string) *EventConsumer {
	return &EventConsumer{
		reader: kafka.NewReader(kafka.ReaderConfig{
			Brokers: []string{s.config.Broker},
			Topic:   topic,
			GroupID: s.config.GroupId,
		}),
	}
}

type EventProducer struct {
	writer *kafka.Writer
}

type EventConsumer struct {
	reader *kafka.Reader
}

func (s *EventProducer) Pub(eventId string) error {
	err := s.writer.WriteMessages(context.Background(), kafka.Message{
		Key:   []byte(eventId),
		Value: []byte(eventId),
	})

	if err != nil {
		return fmt.Errorf("error on write message: %v", err)
	}

	return nil
}

func (s *EventConsumer) On(handler func(eventId string, err error)) {
	go func() {
		for {
			msg, err := s.reader.ReadMessage(context.Background())
			if err != nil {
				handler("", err)
			} else {
				handler(string(msg.Value), nil)
			}
		}
	}()
}
