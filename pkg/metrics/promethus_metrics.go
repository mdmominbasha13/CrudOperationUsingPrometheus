package metrics

import "github.com/prometheus/client_golang/prometheus"

var (
	RequestCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "endpoint", "status"},
	)
	RequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Duration of HTTP requests in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "endpoint", "status"},
	)
	InProgressRequests = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "http_requests_in_progress",
			Help: "Number of HTTP requests in progress",
		},
		[]string{"method", "endpoint"},
	)
)

func InitMetric() {
	prometheus.MustRegister(RequestCounter, RequestDuration, InProgressRequests)
}
