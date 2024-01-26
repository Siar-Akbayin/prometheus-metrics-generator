package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type GaugeWithLabels struct {
	gauge       *prometheus.GaugeVec
	labelNames  []string
	updateCount int
}

func CreateAndRegisterGauge(cardinality int) GaugeWithLabels {
	baseName := fmt.Sprintf("sample_gauge_%d", cardinality)
	labelNames := make([]string, cardinality)
	for i := 0; i < cardinality; i++ {
		labelNames[i] = fmt.Sprintf("dim%d", i+1)
	}
	gaugeVec := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: baseName,
			Help: fmt.Sprintf("A sample gauge metric with %d cardinality", cardinality),
		},
		labelNames,
	)
	prometheus.MustRegister(gaugeVec)
	return GaugeWithLabels{gauge: gaugeVec, labelNames: labelNames, updateCount: 0}
}

func UpdateGauge(gaugeWithLabels *GaugeWithLabels) {
	labelValues := make([]string, len(gaugeWithLabels.labelNames))
	for i := range labelValues {
		labelValues[i] = fmt.Sprintf("%d", i+1) // Assign sequential values to labels
	}
	gaugeWithLabels.gauge.WithLabelValues(labelValues...).Set(rand.Float64() * 3000)
	gaugeWithLabels.updateCount++
}

// Generates one gauge metric for each cardinality from 1 to 3000
func main() {
	const maxCardinality = 3000

	gauges := make([]GaugeWithLabels, maxCardinality)
	for i := 1; i <= maxCardinality; i++ {
		gauges[i-1] = CreateAndRegisterGauge(i)
	}

	go func() {
		for {
			for i := 0; i < maxCardinality; i++ {
				UpdateGauge(&gauges[i])
			}
			if gauges[maxCardinality-1].updateCount >= 1 {
				break
			}
			time.Sleep(100 * time.Millisecond)
		}
	}()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8082", nil)
}
