package collector

import (
	"github.com/erraa/aciexporter/aci"
)

var aciObjects = []aci.AciObject{&aci.FvBD{}}

func CollectMetrics(client *aci.Client) (map[string]map[string]string, error) {
	metrics := make(map[string]map[string]string)
	for _, aciObject := range aciObjects {
		o, err := aci.Get(aciObject, client)
		if err != nil {
			Log.Fatalf("Error fetching object from APIC %s", err)
			return nil, err
		}
		metrics[o.Class()] = map[string]string{"count": o.GetTotalCount()}
		attributes, err := o.GetAttributes()
		if err != nil {
			return nil, err
		}
		for k, v := range attributes {
			metrics[o.Class()][k] = v
		}
	}
	return metrics, nil
}
