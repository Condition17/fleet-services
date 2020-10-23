package baseservice

import (
	"context"
	"encoding/json"
	"log"

	"github.com/Condition17/fleet-services/lib/auth"
	"github.com/Condition17/fleet-services/lib/topics"
	runControllerProto "github.com/Condition17/fleet-services/run-controller-service/proto/run-controller-service"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/broker"
)

type BaseHandler struct {
	Service        micro.Service
	MessagesBroker broker.Broker
}

func NewBaseHandler(service micro.Service) BaseHandler {
	return BaseHandler{Service: service, MessagesBroker: service.Server().Options().Broker}
}

func (h *BaseHandler) SendRunStateEvent(ctx context.Context, eventType string, data []byte) {
	var usrDetails []byte = auth.GetUserBytesFromContext(ctx)
	msgBody, _ := json.Marshal(&runControllerProto.Event{Type: eventType, User: usrDetails, Data: data})
	h.publishMessage(topics.RunStateTopic, &broker.Message{Body: msgBody});
}

func (h *BaseHandler) SendChunkDataToUploadQueue(ctx context.Context data []byte) {
	h.publishMessage(topics.ChunksUploadQueueTopic, &broker.Message{Body: data})
}

func (h *baseHandler) publishMessage(topic string, message *broker.Message) {
	if err := h.MessagesBroker.Publish(topics.RunStateTopic, &broker.Message{Body: msgBody}); err != nil {
		log.Printf("[Messages Broker] Failed to publish message on create. Encountered error: %v", err)
	}
}