package reporter

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// PrometheusReporter is a reporter that reports to prometheus
type PrometheusReporter struct {
	missCounter prometheus.Counter
	hitCounter  prometheus.Counter
}

// NewPrometheusReporter creates a new prometheus reporter
// using NewPrometheusCounter to create the counters
func NewPrometheusReporter(
	serviceName string,
	cacheName string,
) *PrometheusReporter {
	return &PrometheusReporter{
		missCounter: NewPrometheusCounter(
			serviceName,
			cacheName,
			"miss",
		),
		hitCounter: NewPrometheusCounter(
			serviceName,
			cacheName,
			"hit",
		),
	}
}

// NewPrometheusCounter creates a new prometheus counter
// serviceName: the name of the service
// cacheName: the name of the cache
// counterType: the type of the counter
// build the counter name as: serviceName_counterType_total
func NewPrometheusCounter(
	serviceName string,
	cacheName string,
	counterType string,
) prometheus.Counter {
	return promauto.NewCounter(prometheus.CounterOpts{
		Name: fmt.Sprintf("%s_%s_total", serviceName, counterType),
		Help: "The total number of " + counterType,
		ConstLabels: map[string]string{
			"service": serviceName,
			"cache":   cacheName,
		},
	})
}

// ReportMiss reports a cache miss
func (pr *PrometheusReporter) ReportMiss() {
	pr.missCounter.Inc()
}

// ReportHit reports a cache hit
func (pr *PrometheusReporter) ReportHit() {
	pr.hitCounter.Inc()
}
