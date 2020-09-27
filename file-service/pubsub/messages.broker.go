package pubsub

import (
	"log"

	"github.com/micro/go-micro/v2/broker"
)

type MessagesBroker struct {
	Broker broker.Broker
}

func (m *MessagesBroker) PublishEvent(topic string, data []byte) error {
	if err := m.Broker.Publish(topic, &broker.Message{Body: data}); err != nil {
		log.Printf("[Messages Broker] Failed to publish message to topic '%s'. Encountered error: %v", topic, err)
		return err
	}

	return nil
}
