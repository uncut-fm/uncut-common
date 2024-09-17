package pubsub

import (
	"context"
	common_model "github.com/uncut-fm/uncut-common/model"
	"github.com/uncut-fm/uncut-common/pkg/blockchain/polygon/smartcontract/abi"
	"github.com/uncut-fm/uncut-common/pkg/logger"
	"github.com/uncut-fm/uncut-common/pkg/pubsub"
)

type EventsPublisher struct {
	log       logger.Logger
	publisher pubsub.PublisherInterface
}

func NewPublisher(log logger.Logger, publisher pubsub.PublisherInterface) *EventsPublisher {
	return &EventsPublisher{
		log:       log,
		publisher: publisher,
	}
}

func (e EventsPublisher) SendCollectionCreatedEvent(c context.Context, newCollection common_model.CollectionCreatedEventInfo, isLive bool, blockchainEvent *abi.StoreCollectionCreated) error {
	err := e.publisher.Publish(c, &common_model.BlockchainEvent{
		EventType:       common_model.NftCollectionCreatedBlockchainEvent,
		ParsedEventInfo: newCollection,
		PickedLive:      isLive,
		BlockchainEvent: blockchainEvent,
	})

	return e.log.CheckError(err, e.SendCollectionCreatedEvent)
}
