package grpc

import (
	"context"
	"fmt"
	"github.com/cenkalti/backoff"
	"google.golang.org/grpc"
	"log"
	"strings"
	"time"
)

const (
	defaultTimeoutInSeconds = 10
	rpcPort                 = 443
)

func NewClient(ctx context.Context, address string) (*grpc.ClientConn, error) {
	if strings.Contains(address, "http://") {
		address = strings.Replace(address, "http://", "", 1)
	} else {
		address = fmt.Sprintf("%s:%d", strings.Replace(address, "https://", "", 1), rpcPort)
	}

	var conn *grpc.ClientConn

	operation := func() error {
		var err error
		conn, err = grpc.Dial(address, grpc.WithAuthority(address), grpc.WithInsecure())
		return err
	}

	timeout := time.Second * defaultTimeoutInSeconds

	b := backoff.NewExponentialBackOff()
	b.MaxElapsedTime = timeout
	err := backoff.Retry(operation, b)

	if err != nil {
		return nil, err
	}

	log.Println("Successfully connected to grpc server, address: ", address)
	go monitorContext(ctx, conn)

	return conn, nil
}

func monitorContext(ctx context.Context, conn *grpc.ClientConn) {
	<-ctx.Done()
	log.Println("closing grpc conn")
	conn.Close()
}
