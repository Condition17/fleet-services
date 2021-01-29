package baseservice

import (
	"context"
	"encoding/json"
	"log"

	"github.com/Condition17/fleet-services/lib/auth"
	topics "github.com/Condition17/fleet-services/lib/communication"
	runStateEvents "github.com/Condition17/fleet-services/run-controller-service/events"
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

func (h *BaseHandler) SendServiceError(ctx context.Context, testRunId uint32, err error) {
	eventData, _ := json.Marshal(&runControllerProto.ServiceErrorEventData{Source: h.Service.Name(), TestRunId: testRunId, Error: []byte(err.Error())})
	h.SendRunStateEvent(ctx, runStateEvents.ServiceError, eventData)
}

func (h *BaseHandler) SendRunStateEvent(ctx context.Context, eventType string, data []byte) {
	msgBody, _ := json.Marshal(
		&runControllerProto.Event{
			Type: eventType,
			Meta: &runControllerProto.EventMetadata{
				Authorization: auth.GetAuthorizationBytesFromContext(ctx),
			},
			Data: data,
		})
	h.publishMessage(topics.RunStateTopic, &broker.Message{Body: msgBody})
}

func (h *BaseHandler) SendChunkDataToUploadQueue(ctx context.Context, data []byte) {
	h.publishMessage(topics.ChunksUploadQueueTopic, &broker.Message{Body: data})
}

func (h *BaseHandler) SendStorageUploadedChunkData(ctx context.Context, data []byte) {
	h.publishMessage(topics.StorageUploadedChunksTopic, &broker.Message{Body: data})
}

func (h *BaseHandler) SendEventToWssQueue(ctx context.Context, eventType string, data []byte) {
	var userBytes []byte

	userBytes = auth.GetUserBytesFromContext(ctx)
	if len(userBytes) == 0 {
		userBytes = auth.GetUserBytesFromDecodedToken(ctx)
	}

	msgBody, _ := json.Marshal(
		&runControllerProto.WssEvent{
			Type:   eventType,
			Target: userBytes,
			Data:   data,
		})
	h.publishMessage(topics.WssTopic, &broker.Message{Body: msgBody, Header: map[string]string{"orderingKey": "wssEventKey"}})
}

func (h *BaseHandler) publishMessage(topic string, message *broker.Message) {
	log.Printf("Writing to topic %s: %s...", topic, string(message.Body)[:80])
	if err := h.MessagesBroker.Publish(topic, message); err != nil {
		log.Printf("[Messages Broker] Failed to publish message on create. Encountered error: %v", err)
	}
}
