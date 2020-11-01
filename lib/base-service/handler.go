package baseservice

import (
	"context"
	"encoding/json"
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
				User:  auth.GetUserBytesFromContext(ctx),
				Token: auth.GetTokenBytesFromContext(ctx),
			},
			Data: data,
		})
	h.publishMessage(topics.RunStateTopic, &broker.Message{Body: msgBody})
}

func (h *BaseHandler) SendChunkDataToUploadQueue(ctx context.Context, data []byte) {
	h.publishMessage(topics.ChunksUploadQueueTopic, &broker.Message{Body: data})
}

func (h BaseHandler) SendNotificationToWssQueue(ctx context.Context, data []byte) {
	h.sendEventToWssQueue(context.Background(), "notification", data)
}

func (h *BaseHandler) sendEventToWssQueue(ctx context.Context, eventType string, data []byte) {
	msgBody, _ := json.Marshal(
		&runControllerProto.WssEvent{
			Type: eventType,
			Data: data,
		})
	h.publishMessage(topics.WssTopic, &broker.Message{Body: msgBody})
}

func (h *BaseHandler) publishMessage(topic string, message *broker.Message) {
	log.Printf("Writing to topic %s: %s", topic, string(message.Body))
	if err := h.MessagesBroker.Publish(topic, message); err != nil {
		log.Printf("[Messages Broker] Failed to publish message on create. Encountered error: %v", err)
	}
}
