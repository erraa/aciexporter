package main

import (
	"fmt"
	"net/http"

	"github.com/erraa/aciexporter/aci"
	"github.com/erraa/aciexporter/collector"
	"github.com/erraa/aciexporter/config"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/log"
)

var c = config.GetConfig()

func init() {
	client := aci.New(c.APIC.BaseURI)
	err := client.Login(
		c.APIC.Username,
		c.APIC.Password,
	)
	if err != nil {
		panic(err)
	}
	prometheus.MustRegister(collector.NewAciCollector(client, "aci"))
}

func main() {
	fmt.Println(c.APIC.BaseURI)
	fmt.Println(c.Log.Loglevel)
	http.Handle("/metrics", promhttp.Handler())
	log.Info("Beginning to serve on port :8080")
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}
