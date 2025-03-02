package boot

import (
	"context"
	"encoding/json"
	"fmt"
	"nk/account/internal/domain"

	"github.com/segmentio/kafka-go"
)

type KafkaConfig struct {
	Broker  string
	GroupID string
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
			GroupID: s.config.GroupID,
		}),
	}
}

type EventProducer struct {
	writer *kafka.Writer
}

type EventConsumer struct {
	reader *kafka.Reader
}

func (s *EventProducer) Pub(event domain.Event) error {
	event.EventType = s.writer.Topic
	payload, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("error on write message: %v", err)
	}

	err = s.writer.WriteMessages(context.Background(), kafka.Message{
		Value: payload,
	})
	if err != nil {
		return fmt.Errorf("error on write message: %v", err)
	}

	return nil
}

func (s *EventConsumer) On(handler func(*domain.Event, error)) {
	go func() {
		for {
			msg, err := s.reader.ReadMessage(context.Background())
			if err != nil {
				handler(nil, err)
				return
			}

			var event domain.Event
			err = json.Unmarshal(msg.Value, &event)
			if err != nil {
				handler(&event, nil)
				return
			}

			event.EventType = s.reader.Config().Topic
			handler(&event, nil)
		}
	}()
}
