package collector

import (
	"fmt"
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

func NewAciCollector(aciClient *aci.Client, namespace string) *AciCollector {
	metrics := map[string]*prometheus.Desc{}
	for _, obj := range aciObjects {
		aci.Get(obj, aciClient)
		attributes, _ := obj.GetAll()
		name := fmt.Sprintf("total_%s", obj.Class())
		obj.GetAll()
		labels := []string{}
		for k := range attributes[0] {
			labels = append(labels, k)
		}
		metrics[obj.Class()] = newGlobalMetric(namespace, name, fmt.Sprintf("Total amount of %s", name), labels)
	}
	acicollector := &AciCollector{
		client:  aciClient,
		metrics: metrics,
	}
	return acicollector
}

type AciCollector struct {
	client  *aci.Client
	metrics map[string]*prometheus.Desc
	mutex   sync.Mutex
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

	for k, v := range metrics {
		var count int = 1
		if err != nil {
			Log.Fatalf("Error %s", err)
		}
		for _, obj := range v {
			labelValues := []string{}
			for _, labelValue := range obj {
				labelValues = append(labelValues, labelValue)
			}
			ch <- prometheus.MustNewConstMetric(c.metrics[k],
				prometheus.GaugeValue, float64(count), labelValues...)
		}
	}

}

func newGlobalMetric(namespace string, metricName string, docString string, labels []string) *prometheus.Desc {
	return prometheus.NewDesc(namespace+"_"+metricName, docString, labels, nil)
}
