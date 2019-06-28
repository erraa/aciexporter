package aci

import "encoding/json"

type FaultInst struct {
	ImData     []IMDATA `json:"imdata"`
	TotalCount string   `json:"totalCount"`
}

func (faultinst *FaultInst) UnmarshalJson(data []byte) error {
	err := json.Unmarshal(data, &faultinst)
	if err != nil {
		return err
	}
	return nil
}

func (faultinst *FaultInst) Class() string {
	return "faultInst"
}

func (faultinst *FaultInst) GetTotalCount() string {
	return faultinst.TotalCount
}

func (faultinst *FaultInst) GetImData() []IMDATA {
	return faultinst.ImData
}

func (faultinst *FaultInst) GetAll() ([]map[string]string, error) {
	returnSlice := []map[string]string{}
	for _, obj := range faultinst.ImData {
		jsonAttr := map[string]string{}
		attr, _ := json.Marshal(obj.FaultInst.Attributes)
		err := json.Unmarshal(attr, &jsonAttr)
		if err != nil {
			return nil, err
		}
		returnSlice = append(returnSlice, jsonAttr)
	}
	return returnSlice, nil
}

type FAULTINST struct {
	Attributes struct {
		Ack             string `json:"ack"`
		Cause           string `json:"cause"`
		ChangeSet       string `json:"changeSet"`
		ChildAction     string `json:"childAction"`
		Code            string `json:"code"`
		Created         string `json:"created"`
		Delegated       string `json:"delegated"`
		Descr           string `json:"descr"`
		Dn              string `json:"dn"`
		Domain          string `json:"domain"`
		HighestSeverity string `json:"highestSeverity"`
		LastTransition  string `json:"lastTransition"`
		Lc              string `json:"lc"`
		Occur           string `json:"occur"`
		OrigSeverity    string `json:"origSeverity"`
		PrevSeverity    string `json:"prevSeverity"`
		Rule            string `json:"rule"`
		Severity        string `json:"severity"`
		Status          string `json:"status"`
		Subject         string `json:"subject"`
		Type            string `json:"type"`
	} `json:"attributes"`
}
