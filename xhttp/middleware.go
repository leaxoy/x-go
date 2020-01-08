package xhttp

import (
	"fmt"
	"net/http"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/zap"
)

type Middleware interface {
	Next(next http.HandlerFunc) http.HandlerFunc
}

type MiddlewareFn func(next http.HandlerFunc) http.HandlerFunc

func (fn MiddlewareFn) Next(next http.HandlerFunc) http.HandlerFunc { return fn(next) }

func NopMiddleware() Middleware {
	return MiddlewareFn(func(n http.HandlerFunc) http.HandlerFunc {
		return n
	})
}

func RecoveryMiddleware() Middleware {
	return MiddlewareFn(func(next http.HandlerFunc) http.HandlerFunc {
		return func(writer http.ResponseWriter, request *http.Request) {
			defer func() {
				if x := recover(); x != nil {
					zap.L().Error("failed to handle http request", zap.Any("panic", x), zap.Stack("stacktrace"))
					writer.Header().Set("X-Recovery-Info", fmt.Sprintf("%+v\n", x))
				}
			}()
			next(writer, request)
		}
	})
}

func LogMiddleware(l *zap.Logger) Middleware {
	return MiddlewareFn(func(next http.HandlerFunc) http.HandlerFunc {
		return func(writer http.ResponseWriter, request *http.Request) {
			now := time.Now()
			next(writer, request)
			if l != nil {
				l.Info("handle http request",
					zap.Duration("latency", time.Since(now)),
					zap.String("host", request.Host),
					zap.String("path", request.URL.Path),
					zap.String("method", request.Method),
				)
			}
		}
	})
}

func MetricMiddleware(histogram *prometheus.HistogramVec, f func(r *http.Request) []string) Middleware {
	return MiddlewareFn(func(next http.HandlerFunc) http.HandlerFunc {
		return func(writer http.ResponseWriter, request *http.Request) {
			timer := prometheus.NewTimer(histogram.WithLabelValues(f(request)...))
			next(writer, request)
			dura := timer.ObserveDuration()
			writer.Header().Set("X-Response-Time", dura.String())
		}
	})
}

func TraceMiddleware(tracer opentracing.Tracer) Middleware {
	if tracer == nil {
		tracer = opentracing.GlobalTracer()
	}
	return MiddlewareFn(func(next http.HandlerFunc) http.HandlerFunc {
		return func(writer http.ResponseWriter, request *http.Request) {
			span, ctx := opentracing.StartSpanFromContextWithTracer(request.Context(), tracer, request.URL.Path)
			span.SetBaggageItem("device_id", "device_id")
			request = request.WithContext(ctx)
			next(writer, request)
			span.Finish()
			writer.Header().Set("X-Device-Id", span.BaggageItem("device_id"))
		}
	})
}
