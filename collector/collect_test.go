package collector

import (
	"testing"

	"github.com/erraa/aciexporter/aci"
	"github.com/erraa/aciexporter/config"
)

func TestCollectMetrics(t *testing.T) {
	c := config.GetConfig()
	client := aci.New(c.APIC.BaseURI)
	err := client.Login(c.APIC.Username, c.APIC.Password)
	if err != nil {
		t.Errorf("%s", err)
	}
	metrics, err := CollectMetrics(client)
	if err != nil {
		t.Errorf("%s", err)
	}
	t.Log(metrics)
}
