package collector

import (
	"strconv"

	"github.com/erraa/aciexporter/aci"
)

var aciObjects = []aci.AciObject{&aci.FvBD{}}

func CollectMetrics(client *aci.Client) (map[string]int, error) {
	metrics := make(map[string]int)
	for _, aciObject := range aciObjects {
		o, err := aci.Get(aciObject, client)
		if err != nil {
			Log.Fatalf("Error fetching object from APIC %s", err)
			return nil, err
		}
		metrics[o.Class()], err = strconv.Atoi(o.GetTotalCount())
	}

	return metrics, nil
}
