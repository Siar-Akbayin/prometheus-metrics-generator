// metrics/metric_generator.go
package metrics

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"math/rand"
)

// Example of creating a gauge with variable cardinality
func createGaugeWithCardinality(baseName string, maxCardinality int) prometheus.GaugeVec {
	labelNames := make([]string, maxCardinality)
	for i := 0; i < maxCardinality; i++ {
		labelNames[i] = fmt.Sprintf("dim%d", i+1)
	}
	return *prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: baseName,
			Help: "A sample gauge metric with variable cardinality",
		},
		labelNames,
	)
}

// Updating metrics with varying cardinality
func updateGauge(gaugeVec prometheus.GaugeVec, maxCardinality int) {
	labelValues := make([]string, maxCardinality)
	for i := 0; i < maxCardinality; i++ {
		labelValues[i] = fmt.Sprintf("%d", i+1)
	}
	gaugeVec.WithLabelValues(labelValues...).Set(rand.Float64() * 100)
}
