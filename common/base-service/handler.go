package baseservice

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/Condition17/fleet-services/common/auth"
	runControllerProto "github.com/Condition17/fleet-services/run-controller-service/proto/run-controller-service"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/broker"
)

const runStateTopic = "test-run-state"

type Handler struct {
	Service        micro.Service
	MessagesBroker broker.Broker
}

func NewHandler(service micro.Service) Handler {
	return Handler{Service: service, MessagesBroker: service.Server().Options().Broker}
}

func (h *Handler) sendRunStateEvent(ctx context.Context, eventType string, data []byte) {
	var usrDetails []byte = auth.GetUserBytesFromContext(ctx)
	msgBody, _ := json.Marshal(&runControllerProto.Event{Type: eventType, User: usrDetails, Data: data})
	fmt.Printf("Sending run state event on topic %s - Body: %v\n", runStateTopic, msgBody)

	if err := h.MessagesBroker.Publish(runStateTopic, &broker.Message{Body: msgBody}); err != nil {
		log.Printf("[Messages Broker] Failed to publish message on create. Encountered error: %v", err)
	}
}
