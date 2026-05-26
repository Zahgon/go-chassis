package metrics

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var registries = make(map[string]NewRegistry)
var prometheusRegistry = prometheus.NewRegistry()

// NewRegistry create a registry
type NewRegistry func(opts Options) Registry

// Registry holds all of metrics collectors
// name is a unique ID for different type of metrics
type Registry interface {
	CreateGauge(opts GaugeOpts) error
	CreateCounter(opts CounterOpts) error
	CreateSummary(opts SummaryOpts) error
	CreateHistogram(opts HistogramOpts) error

	GaugeSet(name string, val float64, labels map[string]string) error
	GaugeAdd(name string, val float64, labels map[string]string) error
	CounterAdd(name string, val float64, labels map[string]string) error
	SummaryObserve(name string, val float64, Labels map[string]string) error
	HistogramObserve(name string, val float64, labels map[string]string) error

	GaugeValue(name string, labels map[string]string) float64
	CounterValue(name string, labels map[string]string) float64
	SummaryValue(name string, labels map[string]string) (uint64, float64)

	Reset(name string) error
}

var defaultRegistry Registry

// CreateGauge init a new gauge type
func CreateGauge(opts GaugeOpts) error { _ = "STUB: not implemented"; return nil }

// CreateCounter init a new counter type
func CreateCounter(opts CounterOpts) error { _ = "STUB: not implemented"; return nil }

// CreateSummary init a new summary type
func CreateSummary(opts SummaryOpts) error { _ = "STUB: not implemented"; return nil }

// CreateHistogram init a new summary type
func CreateHistogram(opts HistogramOpts) error { _ = "STUB: not implemented"; return nil }

// GaugeSet set a new value to a collector
func GaugeSet(name string, val float64, labels map[string]string) error {
	_ = "STUB: not implemented"
	return nil
}

// GaugeAdd set a new value to a collector
func GaugeAdd(name string, val float64, labels map[string]string) error {
	_ = "STUB: not implemented"
	return nil
}

// CounterAdd increase value of a collector
func CounterAdd(name string, val float64, labels map[string]string) error {
	_ = "STUB: not implemented"
	return nil
}

// SummaryObserve gives a value to summary collector
func SummaryObserve(name string, val float64, labels map[string]string) error {
	_ = "STUB: not implemented"
	return nil
}

// HistogramObserve gives a value to histogram collector
func HistogramObserve(name string, val float64, labels map[string]string) error {
	_ = "STUB: not implemented"
	return nil
}

// Reset clear collector metrics
func Reset(name string) error { _ = "STUB: not implemented"; return nil }

func GaugeValue(name string, labels map[string]string) float64 { _ = "STUB: not implemented"; return 0 }

func CounterValue(name string, labels map[string]string) float64 {
	_ = "STUB: not implemented"
	return 0
}

func SummaryValue(name string, labels map[string]string) (uint64, float64) {
	_ = "STUB: not implemented"
	return 0, 0
}

// CounterOpts is options to create a counter options
type CounterOpts struct {
	// Key is the key set joining with '_', Name will be ignored when Key is not empty
	Key    string
	Name   string
	Help   string
	Labels []string
}

// GaugeOpts is options to create a gauge collector
type GaugeOpts struct {
	Key    string
	Name   string
	Help   string
	Labels []string
}

// SummaryOpts is options to create summary collector
type SummaryOpts struct {
	Key        string
	Name       string
	Help       string
	Labels     []string
	Objectives map[float64]float64
}

// HistogramOpts is options to create histogram collector
type HistogramOpts struct {
	Key     string
	Name    string
	Help    string
	Labels  []string
	Buckets []float64
}

// Options control config
type Options struct {
	FlushInterval          time.Duration
	EnableGoRuntimeMetrics bool
}

// InstallPlugin install metrics registry
func InstallPlugin(name string, f NewRegistry) { _ = "STUB: not implemented"; return }

// Init load the metrics plugin and initialize it
func Init() error {
	_ = "STUB: not implemented"
	// TODO name should be configurable
	return nil
}

// GetSystemPrometheusRegistry return prometheus registry which go chassis use
func GetSystemPrometheusRegistry() *prometheus.Registry { _ = "STUB: not implemented"; return nil }
