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

package metrics_test

import (
	"reflect"
	"testing"

	"github.com/prometheus/client_golang/prometheus"

	"istio.io/pkg/collateral/metrics"
)

func TestPrometheusExportedMetrics(t *testing.T) {
	want := []metrics.Exported{
		{"pilot_xds_cds_reject", "GAUGE", "Pilot rejected CDS."},
		{"pilot_xds_eds_reject", "GAUGE", "Pilot rejected EDS."},
		{"pilot_xds_lds_reject", "GAUGE", "Pilot rejected LDS."},
		{"pilot_xds_rds_reject", "GAUGE", "Pilot rejected RDS."},
	}

	if got := metrics.PrometheusExportedMetrics(); !reflect.DeepEqual(got, want) {
		t.Errorf("ExportedMetrics() = %v, want %v", got, want)
	}
}

var (
	// experiment on getting some monitoring on config errors.
	cdsReject = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "pilot_xds_cds_reject",
		Help: "Pilot rejected CDS.",
	}, []string{"node", "err"})

	edsReject = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "pilot_xds_eds_reject",
		Help: "Pilot rejected EDS.",
	}, []string{"node", "err"})

	edsInstances = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "pilot_xds_eds_instances",
		Help: "Instances for each cluster, as of last push. Zero instances is an error.",
	}, []string{"cluster"})

	ldsReject = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "pilot_xds_lds_reject",
		Help: "Pilot rejected LDS.",
	}, []string{"node", "err"})

	rdsReject = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "pilot_xds_rds_reject",
		Help: "Pilot rejected RDS.",
	}, []string{"node", "err"})
)

func init() {
	prometheus.MustRegister(cdsReject)
	prometheus.MustRegister(edsReject)
	prometheus.MustRegister(ldsReject)
	prometheus.MustRegister(rdsReject)

	cdsReject.WithLabelValues("test", "test").Add(16.45)
	edsReject.WithLabelValues("test", "test").Add(16.45)
	ldsReject.WithLabelValues("test", "test").Add(16.45)
	rdsReject.WithLabelValues("test", "test").Add(16.45)
}
