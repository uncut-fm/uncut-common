package pubsub

import (
	"context"
	common_model "github.com/uncut-fm/uncut-common/model"
	"github.com/uncut-fm/uncut-common/pkg/blockchain/polygon/smartcontract/abi"
	"github.com/uncut-fm/uncut-common/pkg/logger"
	"github.com/uncut-fm/uncut-common/pkg/pubsub"
)

type EventsProcessor struct {
	log       logger.Logger
	publisher pubsub.PublisherInterface
}

func NewProcessor(log logger.Logger, publisher pubsub.PublisherInterface) *EventsProcessor {
	return &EventsProcessor{
		log:       log,
		publisher: publisher,
	}
}

type EventToProcess struct {
	Event             common_model.BlockchainEvent
	CollectionAddress string
}

func (e EventsProcessor) ProcessTransferEvent(c context.Context, blockchainEvent *abi.CollectionTransferSingle, isLive bool, collectionAddress string) error {
	err := e.publisher.Publish(c, &EventToProcess{
		Event: common_model.BlockchainEvent{
			EventType:       common_model.TransferBlockchainEventType,
			PickedLive:      isLive,
			BlockchainEvent: blockchainEvent,
		},
		CollectionAddress: collectionAddress,
	})

	return e.log.CheckError(err, e.ProcessTransferEvent)
}

func (e EventsProcessor) ProcessNFTMintedEvent(c context.Context, blockchainEvent *abi.StoreNFTMinted, isLive bool) error {
	err := e.publisher.Publish(c, &EventToProcess{
		Event: common_model.BlockchainEvent{
			EventType:       common_model.NftMintedBlockchainEventType,
			PickedLive:      isLive,
			BlockchainEvent: blockchainEvent,
		},
	})

	return e.log.CheckError(err, e.ProcessNFTMintedEvent)
}

func (e EventsProcessor) ProcessNFTPriceChangedEvent(c context.Context, blockchainEvent *abi.StoreNFTPriceChanged, isLive bool) error {
	err := e.publisher.Publish(c, &EventToProcess{
		Event: common_model.BlockchainEvent{
			EventType:       common_model.NftPriceChangedBlockchainEvent,
			PickedLive:      isLive,
			BlockchainEvent: blockchainEvent,
		},
	})

	return e.log.CheckError(err, e.ProcessNFTPriceChangedEvent)
}
