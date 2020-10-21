package eventHandler

import (
	"fmt"

	"github.com/micro/go-micro/v2/broker"
)

type eventHandlerFunc = func(broker.Event) error

func New() eventHandlerFunc {
	return func(e broker.Event) error {
		fmt.Printf("Received message body: %v\n", e.Message().Body)
		return nil
	}
}
