package drpc

import (
	"context"
	"fmt"
	"github.com/cenkalti/backoff"
	"log"
	"storj.io/drpc/drpcconn"
	"storj.io/drpc/drpcmigrate"
	"strings"
	"time"
)

const (
	defaultTimeoutInSeconds = 10
	rpcPort                 = 443
)

func NewClient(ctx context.Context, address string) (*drpcconn.Conn, error) {
	if strings.Contains(address, "http://") {
		address = strings.Replace(address, "http://", "", 1)
	} else {
		address = fmt.Sprintf("%s:%s", strings.Replace(address, "https://", "", 1), rpcPort)
	}

	var conn *drpcconn.Conn

	operation := func() error {
		rawconn, err := drpcmigrate.DialWithHeader(ctx, "tcp", address, drpcmigrate.DRPCHeader)
		if err != nil {
			return err
		}

		conn = drpcconn.New(rawconn)
		return nil
	}

	timeout := time.Second * defaultTimeoutInSeconds

	b := backoff.NewExponentialBackOff()
	b.MaxElapsedTime = timeout
	err := backoff.Retry(operation, b)

	if err != nil {
		return nil, err
	}

	log.Println("Successfully connected to drpc server, address: ", address)
	go monitorContext(ctx, conn)

	return conn, nil
}

func monitorContext(ctx context.Context, conn *drpcconn.Conn) {
	<-ctx.Done()
	log.Println("closing drpc conn")
	conn.Close()
}
