package metrics

import (
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"

	dto "github.com/prometheus/client_model/go"
)

var onceEnable sync.Once

// PrometheusExporter is a prom exporter for go chassis
type PrometheusExporter struct {
	FlushInterval time.Duration
	lc            sync.RWMutex
	lg            sync.RWMutex
	ls            sync.RWMutex
	lh            sync.RWMutex
	counters      map[string]*prometheus.CounterVec
	gauges        map[string]*prometheus.GaugeVec
	summaries     map[string]*prometheus.SummaryVec
	histograms    map[string]*prometheus.HistogramVec
}

// NewPrometheusExporter create a prometheus exporter
func NewPrometheusExporter(options Options) Registry {
	_ = "STUB: not implemented"
	return *new(Registry)
}

// EnableRunTimeMetrics enable runtime metrics
func EnableRunTimeMetrics() { _ = "STUB: not implemented"; return }

// CreateGauge create collector
func (c *PrometheusExporter) CreateGauge(opts GaugeOpts) error {
	_ = "STUB: not implemented"
	return nil
}

// GaugeSet set value
func (c *PrometheusExporter) GaugeSet(name string, val float64, labels map[string]string) error {
	_ = "STUB: not implemented"
	return nil
}

// GaugeAdd add value, can be negative
func (c *PrometheusExporter) GaugeAdd(name string, val float64, labels map[string]string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *PrometheusExporter) GaugeValue(name string, labels map[string]string) float64 {
	_ = "STUB: not implemented"
	return 0
}

func (c *PrometheusExporter) CounterValue(name string, labels map[string]string) float64 {
	_ = "STUB: not implemented"
	return 0
}

func (c *PrometheusExporter) SummaryValue(name string, labels map[string]string) (uint64, float64) {
	_ = "STUB: not implemented"
	return 0, 0
}

// CreateCounter create collector
func (c *PrometheusExporter) CreateCounter(opts CounterOpts) error {
	_ = "STUB: not implemented"
	return nil
}

// CounterAdd increase value
func (c *PrometheusExporter) CounterAdd(name string, val float64, labels map[string]string) error {
	_ = "STUB: not implemented"
	return nil
}

// CreateSummary create collector
func (c *PrometheusExporter) CreateSummary(opts SummaryOpts) error {
	_ = "STUB: not implemented"
	return nil
}

// SummaryObserve set value
func (c *PrometheusExporter) SummaryObserve(name string, val float64, labels map[string]string) error {
	_ = "STUB: not implemented"
	return nil
}

// CreateHistogram create collector
func (c *PrometheusExporter) CreateHistogram(opts HistogramOpts) error {
	_ = "STUB: not implemented"
	return nil
}

// HistogramObserve set value
func (c *PrometheusExporter) HistogramObserve(name string, val float64, labels map[string]string) error {
	_ = "STUB: not implemented"
	return nil
}

// Reset reset a collector metrics
func (c *PrometheusExporter) Reset(name string) error { _ = "STUB: not implemented"; return nil }

func Split(key string) (string, string, string) { _ = "STUB: not implemented"; return "", "", "" }

func init() {
	registries["prometheus"] = NewPrometheusExporter
}

func getValue(name string, labels map[string]string, getV func(m *dto.Metric) float64) float64 {
	_ = "STUB: not implemented"
	return 0
}

func getSummaryValue(name string, labels map[string]string, getV func(m *dto.Metric) *dto.Summary) (uint64, float64) {
	_ = "STUB: not implemented"
	return 0, 0
}

func family(name string) *dto.MetricFamily { _ = "STUB: not implemented"; return nil }

func matchLabels(m *dto.Metric, labels map[string]string) bool {
	_ = "STUB: not implemented"
	return false
}
