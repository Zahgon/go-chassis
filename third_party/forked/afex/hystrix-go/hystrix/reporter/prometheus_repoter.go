// Copyright 2016 Csergő Bálint github.com/deathowl
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//     http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package reporter

// Forked from github.com/deathowl
// Some parts of this file have been modified to make it functional in this package
import (
	"sync"
	"time"

	"github.com/go-chassis/go-chassis/v2/third_party/forked/afex/hystrix-go/hystrix"
	"github.com/prometheus/client_golang/prometheus"
)

var onceInit sync.Once

var desc = map[string]string{
	"Consumer.attempts":          "all requests count to target provider service",
	"Consumer.errors":            "if a request is failed because of error or timeout or circuit open, it will increase",
	"Consumer.successes":         "if request is not timeout or failed, it will increase",
	"Consumer.failures":          "if request return error, it will increase",
	"Consumer.rejects":           "if circuit open, then all request will be reject immediately, it will increas",
	"Consumer.shortCircuits":     "after circuit open, it will increase",
	"Consumer.timeouts":          "after timeout, it will increase ",
	"Consumer.fallbackSuccesses": "if fallback is executed and no error returns, it will increase",
	"Consumer.fallbackFailures":  "if fallback is executed and error returns, it will increase",
	"Consumer.totalDuration":     "how long all requests consumed totally",
	"Consumer.runDuration":       "how long a request consumed",

	"Provider.attempts":          "all requests count to target provider service",
	"Provider.errors":            "if a request is failed because of error or timeout or circuit open, it will increase",
	"Provider.successes":         "if request is not timeout or failed, it will increase",
	"Provider.failures":          "if request return error, it will increase",
	"Provider.rejects":           "if circuit open, then all request will be reject immediately, it will increas",
	"Provider.shortCircuits":     "after circuit open, it will increase",
	"Provider.timeouts":          "after timeout, it will increase ",
	"Provider.fallbackSuccesses": "if fallback is executed and no error returns, it will increase",
	"Provider.fallbackFailures":  "if fallback is executed and error returns, it will increase",
	"Provider.totalDuration":     "how long all requests consumed totally",
	"Provider.runDuration":       "how long a request consumed",
}

// GetDesc retrieve metric doc
func GetDesc(name string) string { _ = "STUB: not implemented"; return "" }

var FlushInterval time.Duration //interval to update prom metrics
var gauges map[string]prometheus.Gauge
var gaugeVecs map[string]*prometheus.GaugeVec

// GetPrometheusSinker get prometheus configurations
func GetPrometheusSinker() { _ = "STUB: not implemented"; return }

func flattenKey(key string) string { _ = "STUB: not implemented"; return "" }

func gaugeVecFromNameAndValue(name string, val float64, labels prometheus.Labels) {
	_ = "STUB: not implemented"
	return
}

// ReportMetricsToPrometheus report metrics to prometheus registry, you can use GetSystemPrometheusRegistry to get prometheus registry. by default chassis will report system metrics to prometheus
func ReportMetricsToPrometheus(cb *hystrix.CircuitBreaker) error {
	_ = "STUB: not implemented"
	return nil
}

func init() {
	hystrix.InstallReporter("Prometheus", ReportMetricsToPrometheus)
}
