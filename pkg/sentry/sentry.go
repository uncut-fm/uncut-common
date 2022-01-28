package sentry

import "github.com/getsentry/sentry-go"

func Init(dsn, environment string, sampleRate float32) error {
	return sentry.Init(sentry.ClientOptions{
		Dsn:              dsn,
		Debug:            true,
		AttachStacktrace: true,
		Environment:      environment,
		TracesSampler: sentry.TracesSamplerFunc(func(ctx sentry.SamplingContext) sentry.Sampled {
			// do not send events on local environment
			if environment == "local" {
				return sentry.SampledFalse
			}
			return sentry.UniformTracesSampler(sampleRate).Sample(ctx)
		}),
	})
}
