package xhttp

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/opentracing/opentracing-go"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestMiddleware(t *testing.T) {
	t.Run("nop", func(t *testing.T) {
		req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "/", nil)
		assert.Nil(t, err)
		rr := httptest.NewRecorder()
		NopMiddleware().Next(func(writer http.ResponseWriter, request *http.Request) {
			fmt.Fprint(writer, "Hello,world")
		}).ServeHTTP(rr, req)
		assert.Equal(t, 200, rr.Result().StatusCode)
		assert.EqualValues(t, []byte("Hello,world"), rr.Body.Bytes())
	})
	t.Run("recovery", func(t *testing.T) {
		req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "/", nil)
		assert.Nil(t, err)
		rr := httptest.NewRecorder()
		RecoveryMiddleware().Next(func(writer http.ResponseWriter, request *http.Request) {
			panic("panic test")
		}).ServeHTTP(rr, req)
		assert.Equal(t, 200, rr.Result().StatusCode)
		assert.Contains(t, rr.Header(), "X-Recovery-Info")
	})
	t.Run("log", func(t *testing.T) {
		req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "/", nil)
		assert.Nil(t, err)
		rr := httptest.NewRecorder()
		LogMiddleware(zap.L()).Next(func(writer http.ResponseWriter, request *http.Request) {
			fmt.Fprint(writer, "Hello,world")
		}).ServeHTTP(rr, req)
		assert.Equal(t, 200, rr.Result().StatusCode)
	})
	t.Run("metric", func(t *testing.T) {
		req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "/", nil)
		assert.Nil(t, err)
		rr := httptest.NewRecorder()
		MetricMiddleware(promauto.NewHistogramVec(prometheus.HistogramOpts{
			Name: "http_request_seconds",
			Help: "how long time spent on process http request",
		}, []string{"url", "method"}), func(r *http.Request) []string {
			return []string{r.URL.Path, r.Method}
		}).Next(func(writer http.ResponseWriter, request *http.Request) {
			fmt.Fprint(writer, "Hello,world")
		}).ServeHTTP(rr, req)
		assert.Equal(t, 200, rr.Result().StatusCode)
		assert.Contains(t, rr.Header(), "X-Response-Time")
	})
	t.Run("trace", func(t *testing.T) {
		req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "/", nil)
		assert.Nil(t, err)
		rr := httptest.NewRecorder()
		TraceMiddleware(opentracing.GlobalTracer()).Next(func(writer http.ResponseWriter, request *http.Request) {
			fmt.Fprint(writer, "Hello,world")
		}).ServeHTTP(rr, req)
		assert.Equal(t, 200, rr.Result().StatusCode)
		assert.Contains(t, rr.Header(), "X-Device-Id")
	})
}
