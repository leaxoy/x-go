package sentry

import (
	"github.com/getsentry/raven-go"
)

// sentry init default client use env
// client.SetDSN(os.Getenv("SENTRY_DSN"))
// client.SetRelease(os.Getenv("SENTRY_RELEASE"))
// client.SetEnvironment(os.Getenv("SENTRY_ENVIRONMENT"))

func Init(serviceName string) {
	raven.SetDefaultLoggerName(serviceName)
}
