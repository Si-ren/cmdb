package global

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	RequestTotal *prometheus.CounterVec
	ResponseTime *prometheus.HistogramVec
)

func init() {
	// 定义指标对象

	RequestTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests.",
		},
		[]string{"method", "function", "status"},
	)
	ResponseTime = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "http_response_time_seconds",
			Help: "HTTP response time distribution.",
			Buckets: []float64{
				0.01, 0.05, 0.1, 0.5, 1, 5, 10,
			},
		},
		[]string{"method", "function", "status"},
	)
}
