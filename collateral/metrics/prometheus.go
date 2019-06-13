// Copyright 2019 Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

// PrometheusExportedMetrics should only be used to collected exported metrics for
// the offline generation of collateral. It is not suitable for production usage.
func PrometheusExportedMetrics() []Exported {
	// For collateral, we are NOT interested in the process/go metrics that Prometheus
	// exports by default. Rather, we want to focus on the process-specific metrics
	prometheus.Unregister(prometheus.NewProcessCollector(prometheus.ProcessCollectorOpts{}))
	prometheus.Unregister(prometheus.NewGoCollector())

	f, err := prometheus.DefaultGatherer.Gather()
	if err != nil {
		return []Exported{}
	}

	out := make([]Exported, 0, len(f))
	for _, metricFam := range f {
		out = append(out, Exported{metricFam.GetName(), metricFam.GetType().String(), metricFam.GetHelp()})
	}

	return out
}
