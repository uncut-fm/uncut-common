package tracing

import (
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel/trace"
	"net/http"
)

func NewHTTPClient(tp trace.TracerProvider) *http.Client {
	return &http.Client{
		Transport: NewTransport(tp),
	}
}

func NewTransport(tp trace.TracerProvider) *otelhttp.Transport {
	return otelhttp.NewTransport(
		http.DefaultTransport,
		otelhttp.WithTracerProvider(tp),
	)
}
