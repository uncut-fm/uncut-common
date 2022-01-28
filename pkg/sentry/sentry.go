package sentry

import (
	"github.com/getsentry/sentry-go"
	"github.com/uncut-fm/uncut-common/pkg/config"
)

func Init(sentryConfigs config.SentryConfigs, environment string) error {
	if sentryConfigs.SampleRate == 0 {
		sentryConfigs.SampleRate = 1
	}
	return sentry.Init(sentry.ClientOptions{
		Dsn:              sentryConfigs.DSN,
		Debug:            true,
		AttachStacktrace: true,
		Environment:      environment,
		TracesSampler: sentry.TracesSamplerFunc(func(ctx sentry.SamplingContext) sentry.Sampled {
			// do not send events on local environment
			if environment == "local" {
				return sentry.SampledFalse
			}
			return sentry.UniformTracesSampler(sentryConfigs.SampleRate).Sample(ctx)
		}),
	})
}
