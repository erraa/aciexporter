package collector

import (
	"sync"

	"github.com/erraa/aciexporter/aci"
	"github.com/erraa/aciexporter/acilogger"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
)

var Log = acilogger.Log.WithFields(logrus.Fields{
	"package": "collector",
	"file":    "collector.go",
})

type AciCollector struct {
	client  *aci.Client
	metrics map[string]*prometheus.Desc
	mutex   sync.Mutex
}

func NewAciCollector(aciClient *aci.Client, namespace string) *AciCollector {
	return &AciCollector{
		client: aciClient,
		metrics: map[string]*prometheus.Desc{
			"fvBD": newGlobalMetric(namespace, "total_bds", "Total amount of bds"),
		},
	}
}

func (c *AciCollector) Describe(ch chan<- *prometheus.Desc) {
	for _, m := range c.metrics {
		ch <- m
	}
}

func (c *AciCollector) Collect(ch chan<- prometheus.Metric) {
	c.mutex.Lock() // To protect metrics from concurrent collects
	defer c.mutex.Unlock()

	metrics, err := CollectMetrics(c.client)
	if err != nil {
		Log.Fatalf("Error getting stats %s", err)
	}

	ch <- prometheus.MustNewConstMetric(c.metrics["fvBD"],
		prometheus.GaugeValue, float64(metrics["fvBD"]))
}

func newGlobalMetric(namespace string, metricName string, docString string) *prometheus.Desc {
	return prometheus.NewDesc(namespace+"_"+metricName, docString, nil, nil)
}
