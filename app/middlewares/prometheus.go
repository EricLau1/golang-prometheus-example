package middlewares

import (
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"net/http"
)

var (
	totalRequests = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Número de requisições GET",
		},
		[]string{"path"},
	)

	responseStatus = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_response_status",
			Help: "Status da resposta HTTP",
		},
		[]string{"status"},
	)

	httpDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_response_time_seconds",
			Help:    "Duração das respostas HTTP",
			Buckets: []float64{0.05, 0.1, 0.2, 0.3, 0.4, 0.5}, // buckets in milliseconds
		},
		[]string{"path"},
	)
)

func init() {
	prometheus.MustRegister(totalRequests)
	prometheus.MustRegister(responseStatus)
	prometheus.MustRegister(httpDuration)
}

func Metrics(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		var (
			route   = mux.CurrentRoute(request)
			path, _ = route.GetPathTemplate()

			timer = prometheus.NewTimer(httpDuration.WithLabelValues(path))
			rw    = newResponseWriter(writer)
		)

		next.ServeHTTP(rw, request)

		responseStatus.WithLabelValues(rw.StringStatus()).Inc()
		totalRequests.WithLabelValues(path).Inc()

		_ = timer.ObserveDuration()
	})
}
