package baseservice

import (
	"cloud.google.com/go/pubsub"
	"context"
	"encoding/json"
	"github.com/Condition17/fleet-services/lib/auth"
	topics "github.com/Condition17/fleet-services/lib/communication"
	runStateEvents "github.com/Condition17/fleet-services/run-controller-service/events"
	runControllerProto "github.com/Condition17/fleet-services/run-controller-service/proto/run-controller-service"
	"github.com/micro/go-micro/v2"
	"log"
)

const GoogleProjectId = "fleet-295921"
type BaseHandler struct {
	Service        micro.Service
	PubSubClient  *pubsub.Client
}

func NewBaseHandler(service micro.Service) BaseHandler {
	pubSubClient, err := pubsub.NewClient(context.Background(), GoogleProjectId)

	if err != nil {
		log.Fatalf("Error initializing pubsub client for google project id %v\n", GoogleProjectId)
		return BaseHandler{}
	}

	return BaseHandler{Service: service, PubSubClient: pubSubClient}
}

func (h *BaseHandler) SendServiceError(ctx context.Context, testRunId uint32, err error) {
	eventData, _ := json.Marshal(&runControllerProto.ServiceErrorEventData{Source: h.Service.Name(), TestRunId: testRunId, Error: []byte(err.Error())})
	h.SendRunStateEvent(ctx, runStateEvents.ServiceError, eventData)
}

func (h *BaseHandler) SendRunStateEvent(ctx context.Context, eventType string, data []byte) {
	msgData, _ := json.Marshal(
		&runControllerProto.Event{
			Type: eventType,
			Meta: &runControllerProto.EventMetadata{
				Authorization: auth.GetAuthorizationBytesFromContext(ctx),
			},
			Data: data,
		})
	h.publishMessage(topics.RunStateTopic, msgData)
}

func (h *BaseHandler) SendChunkDataToUploadQueue(ctx context.Context, data []byte) {
	h.publishMessage(topics.ChunksUploadQueueTopic, data)
}

func (h *BaseHandler) SendStorageUploadedChunkData(ctx context.Context, data []byte) {
	h.publishMessage(topics.StorageUploadedChunksTopic, data)
}

func (h *BaseHandler) SendEventToWssQueue(ctx context.Context, eventType string, data []byte) {
	var userBytes []byte

	userBytes = auth.GetUserBytesFromContext(ctx)
	if len(userBytes) == 0 {
		userBytes = auth.GetUserBytesFromDecodedToken(ctx)
	}

	msgData, _ := json.Marshal(
		&runControllerProto.WssEvent{
			Type:   eventType,
			Target: userBytes,
			Data:   data,
		})
	h.publishMessage(topics.WssTopic, msgData)
}

func (h *BaseHandler) publishMessage(topic string, msgData []byte) {
	var msgPreview string
	if len(msgData) < 80 {
		msgPreview = string(msgData)
	} else {
		msgPreview = string(msgData[:80])
	}
	log.Printf("Writing to topic %s: %s...\n\n\n", topic, msgPreview)

	t := h.PubSubClient.Topic(topic)
	res := t.Publish(context.Background(), &pubsub.Message{
		Data: msgData,
		//OrderingKey: fmt.Sprintf("%v", time.Now().Unix()),
	})

	go func(res *pubsub.PublishResult) {
		if _, err := res.Get(context.Background()); err != nil {
			log.Printf("[PubSub] Failed to publish message . Encountered error: %v", err)
			return
		}
	}(res)
}
