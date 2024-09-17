package smartcontract

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/uncut-fm/uncut-common/model"
	"github.com/uncut-fm/uncut-common/pkg/blockchain/polygon/smartcontract/abi"
	"github.com/uncut-fm/uncut-common/pkg/logger"
	"sync"
)

type LogsFilterer struct {
	log                  logger.Logger
	ethClient            ethclient.Client
	collectionsFilterers sync.Map
	storeFilterer        *abi.StoreFilterer
	eventsProcessor      EventsProcessor
	eventsPublisher      EventsPublisher
}

type EventsPublisher interface {
	SendCollectionCreatedEvent(c context.Context, newCollection model.CollectionCreatedEventInfo, isLive bool, blockchainEvent *abi.StoreCollectionCreated) error
}

type EventsProcessor interface {
	ProcessTransferEvent(c context.Context, blockchainEvent *abi.CollectionTransferSingle, isLive bool, collectionAddress string) error
	ProcessNFTMintedEvent(c context.Context, blockchainEvent *abi.StoreNFTMinted, isLive bool) error
	ProcessNFTPriceChangedEvent(c context.Context, blockchainEvent *abi.StoreNFTPriceChanged, isLive bool) error
}

func NewLogsSyncer(log logger.Logger, storeAddress string, ethClient *ethclient.Client, publisher EventsPublisher, processor EventsProcessor) (*LogsFilterer, error) {
	storeFilterer, err := abi.NewStoreFilterer(common.HexToAddress(storeAddress), ethClient)
	if err != nil {
		return nil, err
	}
	return &LogsFilterer{
		log:             log,
		ethClient:       *ethClient,
		storeFilterer:   storeFilterer,
		eventsProcessor: processor,
		eventsPublisher: publisher,
	}, nil
}

func (l *LogsFilterer) SyncEvent(ctx context.Context, log types.Log) (bool, error) {
	collectionFilterer, err := l.getCollectionFilterer(log.Address)
	if err == nil {
		blockchainEvent, err := collectionFilterer.ParseTransferSingle(log)
		if err == nil {
			err = l.eventsProcessor.ProcessTransferEvent(ctx, blockchainEvent, false, log.Address.Hex())
			if err != nil {
				return true, err
			}
		}
	}

	collectionEvent, err := l.storeFilterer.ParseCollectionCreated(log)
	if err == nil {
		err = l.eventsPublisher.SendCollectionCreatedEvent(ctx, model.CollectionCreatedEventInfo{
			Name:           collectionEvent.Name,
			Address:        collectionEvent.CollectionAddress.Hex(),
			CreatorAddress: collectionEvent.Creator.Hex(),
			BlockNumber:    int(collectionEvent.Raw.BlockNumber),
			ShowID:         int(collectionEvent.ShowId.Int64()),
		}, false, collectionEvent)
		return true, l.log.CheckError(err, l.SyncEvent)
	}

	nftMintedEvent, err := l.storeFilterer.ParseNFTMinted(log)
	if err == nil {
		err = l.eventsProcessor.ProcessNFTMintedEvent(ctx, nftMintedEvent, false)
		return true, l.log.CheckError(err, l.SyncEvent)
	}

	nftPriceChangedEvent, err := l.storeFilterer.ParseNFTPriceChanged(log)
	if err == nil {
		err = l.eventsProcessor.ProcessNFTPriceChangedEvent(ctx, nftPriceChangedEvent, false)
		return true, l.log.CheckError(err, l.SyncEvent)
	}

	return false, nil
}

func (l *LogsFilterer) getCollectionFilterer(address common.Address) (*abi.CollectionFilterer, error) {
	filtererAny, ok := l.collectionsFilterers.Load(address)
	if ok {
		return filtererAny.(*abi.CollectionFilterer), nil
	}

	filterer, err := abi.NewCollectionFilterer(address, &l.ethClient)
	if err != nil {
		return nil, err
	}

	l.collectionsFilterers.Store(address, filterer)

	return filterer, nil
}
