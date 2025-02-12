package monitoring

import (
    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promauto"
)

type Metrics struct {
    RequestCounter   *prometheus.CounterVec
    RequestDuration  *prometheus.HistogramVec
    ErrorCounter     *prometheus.CounterVec
    ActiveUsers      prometheus.Gauge
    CacheHitRate     prometheus.Gauge
}

func NewMetrics(namespace string) *Metrics {
    return &Metrics{
        RequestCounter: promauto.NewCounterVec(
            prometheus.CounterOpts{
                Namespace: namespace,
                Name:     "http_requests_total",
                Help:     "Total number of HTTP requests",
            },
            []string{"method", "endpoint", "status"},
        ),
        RequestDuration: promauto.NewHistogramVec(
            prometheus.HistogramOpts{
                Namespace: namespace,
                Name:     "http_request_duration_seconds",
                Help:     "HTTP request duration in seconds",
            },
            []string{"method", "endpoint"},
        ),
        ErrorCounter: promauto.NewCounterVec(
            prometheus.CounterOpts{
                Namespace: namespace,
                Name:     "error_count_total",
                Help:     "Total number of errors",
            },
            []string{"type"},
        ),
        ActiveUsers: promauto.NewGauge(prometheus.GaugeOpts{
            Namespace: namespace,
            Name:     "active_users",
            Help:     "Number of currently active users",
        }),
        CacheHitRate: promauto.NewGauge(prometheus.GaugeOpts{
            Namespace: namespace,
            Name:     "cache_hit_rate",
            Help:     "Cache hit rate percentage",
        }),
    }
}
