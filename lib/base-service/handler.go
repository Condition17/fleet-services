package baseservice

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/Condition17/fleet-services/lib/auth"
	topics "github.com/Condition17/fleet-services/lib/communication"
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
	msgBody, _ := json.Marshal(
		&runControllerProto.Event{
			Type: eventType,
			Meta: &runControllerProto.EventMetadata{
				User:          auth.GetUserBytesFromContext(ctx),
				Authorization: auth.GetAuthorizationBytesFromContext(ctx),
			},
			Data: data,
		})
	h.publishMessage(topics.RunStateTopic, &broker.Message{Body: msgBody})
}

func (h *BaseHandler) SendChunkDataToUploadQueue(ctx context.Context, data []byte) {
	h.publishMessage(topics.ChunksUploadQueueTopic, &broker.Message{Body: data})
}

func (h *BaseHandler) SendEventToWssQueue(ctx context.Context, eventType string, data []byte) {
	fmt.Printf("User bytes from context: %v\n", auth.GetUserBytesFromContext(ctx))
	msgBody, _ := json.Marshal(
		&runControllerProto.WssEvent{
			Type:   eventType,
			Target: auth.GetUserBytesFromContext(ctx),
			Data:   data,
		})
	h.publishMessage(topics.WssTopic, &broker.Message{Body: msgBody})
}

func (h *BaseHandler) publishMessage(topic string, message *broker.Message) {
	log.Printf("Writing to topic %s: %s", topic, string(message.Body))
	if err := h.MessagesBroker.Publish(topic, message); err != nil {
		log.Printf("[Messages Broker] Failed to publish message on create. Encountered error: %v", err)
	}
}
