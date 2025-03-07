package tracing

import (
	"context"
	"errors"
	"fmt"
	"github.com/jaegertracing/jaeger/pkg/otelsemconv"
	"github.com/uncut-fm/uncut-common/pkg/logger"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
	"strings"
	"sync"
	"time"
)

const (
	PLATFORM_POSTGRESQL_TRACER_NAME           = "db_platform"
	AUTH_POSTGRESQL_TRACER_NAME               = "db_auth"
	MAIN_REDIS_TRACER_NAME                    = "redis_main"
	SECONDARY_REDIS_TRACER_NAME               = "redis_secondary"
	EVENTS_PUBLISHER_TRACER_NAME              = "events_publisher"
	BLOCKCHAIN_REQUESTS_PUBLISHER_TRACER_NAME = "blockchain_requests_publisher"
	TRANSCODER_REQUESTS_PUBLISHER_TRACER_NAME = "transcoder_requests_publisher"
)

var once sync.Once

func InitOtel(serviceName, otlpURL, env string, logger logger.Logger) trace.TracerProvider {
	once.Do(func() {
		otel.SetTextMapPropagator(
			propagation.NewCompositeTextMapPropagator(
				propagation.TraceContext{},
				propagation.Baggage{},
			))
	})

	exp, err := createOtelExporter("otlp", otlpURL)
	if err != nil {
		logger.Error("cannot create exporter", err)
	}

	res, err := resource.New(
		context.Background(),
		resource.WithSchemaURL(otelsemconv.SchemaURL),
		resource.WithAttributes(otelsemconv.ServiceNameKey.String(serviceName)),
		resource.WithAttributes(attribute.Key("deployment.environment").String(env)),
		resource.WithTelemetrySDK(),
		resource.WithHost(),
		resource.WithOSType(),
	)
	if err != nil {
		logger.Error("resource creation failed", err)
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exp, sdktrace.WithBatchTimeout(1000*time.Millisecond)),
		sdktrace.WithResource(res),
	)

	return tp
}

func createOtelExporter(exporterType, otlpURL string) (sdktrace.SpanExporter, error) {
	var exporter sdktrace.SpanExporter
	var err error

	switch exporterType {
	case "jaeger":
		return nil, errors.New("jaeger exporter is no longer supported, please use otlp")
	case "otlp":
		var opts []otlptracehttp.Option
		if !withSecure(otlpURL) {
			opts = []otlptracehttp.Option{otlptracehttp.WithInsecure(), otlptracehttp.WithEndpointURL(otlpURL)}
		}
		exporter, err = otlptrace.New(
			context.Background(),
			otlptracehttp.NewClient(opts...),
		)
	default:
		return nil, fmt.Errorf("unrecognized exporter type %s", exporterType)
	}

	return exporter, err
}

// withSecure instructs the client to use HTTPS scheme, instead of hotrod's desired default HTTP
func withSecure(url string) bool {
	return strings.HasPrefix(url, "https://")
}
