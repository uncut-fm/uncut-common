package grpc

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/cenkalti/backoff"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"strings"
	"time"
)

const (
	defaultTimeoutInSeconds = 10
	rpcPort                 = 443
)

func NewClient(ctx context.Context, address string, tracerProvider trace.TracerProvider) (*grpc.ClientConn, error) {
	var isInsecure bool
	if strings.Contains(address, "http://") {
		isInsecure = true
		address = strings.Replace(address, "http://", "", 1)
	} else {
		address = fmt.Sprintf("%s:%d", strings.Replace(address, "https://", "", 1), rpcPort)
	}

	var conn *grpc.ClientConn

	operation := func() error {
		var err error

		opts := []grpc.DialOption{grpc.WithAuthority(address),
			grpc.WithStatsHandler(otelgrpc.NewClientHandler(otelgrpc.WithTracerProvider(tracerProvider)))}

		if isInsecure {
			opts = append(opts, grpc.WithInsecure())
		} else {
			systemRoots, err := x509.SystemCertPool()
			if err != nil {
				return err
			}
			cred := credentials.NewTLS(&tls.Config{
				RootCAs: systemRoots,
			})
			opts = append(opts, grpc.WithTransportCredentials(cred))
		}

		conn, err = grpc.Dial(address, opts...)
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
