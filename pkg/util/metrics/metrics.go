package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var podName string
var clusterName string

// RegisterPodName will set the name of the current pod.
func RegisterPodName(name string) {
	podName = name
}

// RegisterClusterName will set the name of the current cluster.
func RegisterClusterName(name string) {
	clusterName = name
}

// RegisterOperatorMetric will register a single operator metric.
func RegisterOperatorMetric(metric prometheus.Collector) {
	assertPodName()
	prometheus.MustRegister(metric)
}

// RegisterAgentMetric will register a single agent metric.
func RegisterAgentMetric(metric prometheus.Collector) {
	assertPodName()
	assertClusterName()
	prometheus.MustRegister(metric)
}

func newCounter(namespace string, subsystem string, name string, help string, labels []string) *prometheus.CounterVec {
	return prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      name,
			Help:      help,
		},
		labels,
	)
}

func newGauge(namespace string, subsystem string, name string, help string, labels []string) *prometheus.GaugeVec {
	return prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      name,
			Help:      help,
		},
		labels,
	)
}

func assertPodName() {
	if podName == "" {
		panic("Metrics package requires podName. Unable to register metrics")
	}
}

func assertClusterName() {
	if clusterName == "" {
		panic("Metrics package requires clusterName. Unable to register metrics")
	}
}

func eventLabels() prometheus.Labels {
	labels := prometheus.Labels{
		"podName": podName,
	}
	if clusterName != "" {
		labels["clusterName"] = clusterName
	}
	return labels
}

func statusLabels(status innodb.InstanceStatus) prometheus.Labels {
	return prometheus.Labels{
		"podName":        podName,
		"clusterName":    clusterName,
		"instanceStatus": string(status),
	}
}