package collector

import (
	"github.com/erraa/aciexporter/aci"
)

var aciObjects = []aci.AciObject{&aci.FvAEPg{}, &aci.FvBD{}, &aci.FaultInst{}}

func CollectMetrics(client *aci.Client) (map[string][]map[string]string, error) {
	metrics := make(map[string][]map[string]string)
	for _, aciObject := range aciObjects {
		o, err := aci.Get(aciObject, client)
		if err != nil {
			Log.Fatalf("Error fetching object from APIC %s", err)
			return nil, err
		}
		attributes, err := o.GetAll()
		if err != nil {
			return nil, err
		}
		for _, v := range attributes {
			metrics[o.Class()] = append(metrics[o.Class()], v)
		}
	}
	return metrics, nil
}
