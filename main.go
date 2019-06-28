package main

import (
	"fmt"
	"net/http"
	"os"

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
		log.Error("Unable to connect to APIC")
		os.Exit(2)
	}
	prometheus.MustRegister(collector.NewAciCollector(client, "aci"))
}

func main() {
	fmt.Println(c.APIC.BaseURI)
	fmt.Println(c.Log.Loglevel)
	http.Handle("/metrics", promhttp.Handler())
	log.Info("Beginning to serve on port :8383")
	log.Fatal(http.ListenAndServe("0.0.0.0:8383", nil))
}
