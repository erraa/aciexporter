package aci

import (
	"encoding/json"

	"github.com/erraa/aciexporter/acilogger"
	"github.com/sirupsen/logrus"
)

var Log = acilogger.Log.WithFields(logrus.Fields{
	"package": "aci",
	"file":    "fvbd.go",
})

type FvBD struct {
	ImData     []IMDATA `json:"imdata"`
	TotalCount string   `json:"totalCount"`
	FvBD       FVBD     `json:"fvBD"`
}

type IMDATA struct {
	FvBD   FVBD   `json:"fvBD"`
	FvAEPg FVAEPG `json:"fvAEPg"`
}

type FVBD struct {
	Attributes struct {
		OptimizeWanBandwidth     string `json:"OptimizeWanBandwidth"`
		ArpFlood                 string `json:"arpFlood"`
		BcastP                   string `json:"bcastP"`
		ChildAction              string `json:"childAction"`
		ConfigIssues             string `json:"configIssues"`
		Descr                    string `json:"descr"`
		EpClear                  string `json:"epClear"`
		EpMoveDetectMode         string `json:"epMoveDetectMode"`
		ExtMngdBy                string `json:"extMngdBy"`
		IntersiteBumTrafficAllow string `json:"intersiteBumTrafficAllow"`
		IntersiteL2Stretch       string `json:"intersiteL2Stretch"`
		IPLearning               string `json:"ipLearning"`
		LcOwn                    string `json:"lcOwn"`
		LimitIPLearnToSubnets    string `json:"limitIpLearnToSubnets"`
		LlAddr                   string `json:"llAddr"`
		Mac                      string `json:"mac"`
		McastAllow               string `json:"mcastAllow"`
		ModTs                    string `json:"modTs"`
		MonPolDn                 string `json:"monPolDn"`
		Mtu                      string `json:"mtu"`
		MultiDstPktAct           string `json:"multiDstPktAct"`
		Name                     string `json:"name"`
		NameAlias                string `json:"nameAlias"`
		OwnerKey                 string `json:"ownerKey"`
		OwnerTag                 string `json:"ownerTag"`
		PcTag                    string `json:"pcTag"`
		Rn                       string `json:"rn"`
		Scope                    string `json:"scope"`
		Seg                      string `json:"seg"`
		Status                   string `json:"status"`
		Type                     string `json:"type"`
		UID                      string `json:"uid"`
		UnicastRoute             string `json:"unicastRoute"`
		UnkMacUcastAct           string `json:"unkMacUcastAct"`
		UnkMcastAct              string `json:"unkMcastAct"`
		Vmac                     string `json:"vmac"`
	} `json:"attributes"`
	Children []struct {
		FvRsBDToNdP struct {
			Attributes struct {
				ChildAction   string `json:"childAction"`
				ForceResolve  string `json:"forceResolve"`
				LcOwn         string `json:"lcOwn"`
				ModTs         string `json:"modTs"`
				MonPolDn      string `json:"monPolDn"`
				RType         string `json:"rType"`
				Rn            string `json:"rn"`
				State         string `json:"state"`
				StateQual     string `json:"stateQual"`
				Status        string `json:"status"`
				TCl           string `json:"tCl"`
				TContextDn    string `json:"tContextDn"`
				TDn           string `json:"tDn"`
				TRn           string `json:"tRn"`
				TType         string `json:"tType"`
				TnNdIfPolName string `json:"tnNdIfPolName"`
				UID           string `json:"uid"`
			} `json:"attributes"`
		} `json:"fvRsBDToNdP"`
		FvRsBdToEpRet struct {
			Attributes struct {
				ChildAction      string `json:"childAction"`
				ForceResolve     string `json:"forceResolve"`
				LcOwn            string `json:"lcOwn"`
				ModTs            string `json:"modTs"`
				MonPolDn         string `json:"monPolDn"`
				RType            string `json:"rType"`
				ResolveAct       string `json:"resolveAct"`
				Rn               string `json:"rn"`
				State            string `json:"state"`
				StateQual        string `json:"stateQual"`
				Status           string `json:"status"`
				TCl              string `json:"tCl"`
				TContextDn       string `json:"tContextDn"`
				TDn              string `json:"tDn"`
				TRn              string `json:"tRn"`
				TType            string `json:"tType"`
				TnFvEpRetPolName string `json:"tnFvEpRetPolName"`
				UID              string `json:"uid"`
			} `json:"attributes"`
		} `json:"fvRsBdToEpRet"`
		FvRsCtx struct {
			Attributes struct {
				ChildAction  string `json:"childAction"`
				ExtMngdBy    string `json:"extMngdBy"`
				ForceResolve string `json:"forceResolve"`
				LcOwn        string `json:"lcOwn"`
				ModTs        string `json:"modTs"`
				MonPolDn     string `json:"monPolDn"`
				RType        string `json:"rType"`
				Rn           string `json:"rn"`
				State        string `json:"state"`
				StateQual    string `json:"stateQual"`
				Status       string `json:"status"`
				TCl          string `json:"tCl"`
				TContextDn   string `json:"tContextDn"`
				TDn          string `json:"tDn"`
				TRn          string `json:"tRn"`
				TType        string `json:"tType"`
				TnFvCtxName  string `json:"tnFvCtxName"`
				UID          string `json:"uid"`
			} `json:"attributes"`
		} `json:"fvRsCtx"`
		FvRsIgmpsn struct {
			Attributes struct {
				ChildAction        string `json:"childAction"`
				ForceResolve       string `json:"forceResolve"`
				LcOwn              string `json:"lcOwn"`
				ModTs              string `json:"modTs"`
				MonPolDn           string `json:"monPolDn"`
				RType              string `json:"rType"`
				Rn                 string `json:"rn"`
				State              string `json:"state"`
				StateQual          string `json:"stateQual"`
				Status             string `json:"status"`
				TCl                string `json:"tCl"`
				TContextDn         string `json:"tContextDn"`
				TDn                string `json:"tDn"`
				TRn                string `json:"tRn"`
				TType              string `json:"tType"`
				TnIgmpSnoopPolName string `json:"tnIgmpSnoopPolName"`
				UID                string `json:"uid"`
			} `json:"attributes"`
		} `json:"fvRsIgmpsn"`
		FvRtBd struct {
			Attributes struct {
				ChildAction string `json:"childAction"`
				LcOwn       string `json:"lcOwn"`
				ModTs       string `json:"modTs"`
				Rn          string `json:"rn"`
				Status      string `json:"status"`
				TCl         string `json:"tCl"`
				TDn         string `json:"tDn"`
			} `json:"attributes"`
		} `json:"fvRtBd"`
		FvSubnet struct {
			Attributes struct {
				ChildAction string `json:"childAction"`
				Ctrl        string `json:"ctrl"`
				Descr       string `json:"descr"`
				ExtMngdBy   string `json:"extMngdBy"`
				IP          string `json:"ip"`
				LcOwn       string `json:"lcOwn"`
				ModTs       string `json:"modTs"`
				MonPolDn    string `json:"monPolDn"`
				Name        string `json:"name"`
				NameAlias   string `json:"nameAlias"`
				Preferred   string `json:"preferred"`
				Rn          string `json:"rn"`
				Scope       string `json:"scope"`
				Status      string `json:"status"`
				UID         string `json:"uid"`
				Virtual     string `json:"virtual"`
			} `json:"attributes"`
		} `json:"fvSubnet"`
	} `json:"children"`
}

func (bd *FvBD) UnmarshalJson(data []byte) error {
	err := json.Unmarshal(data, &bd)
	if err != nil {
		return err
	}
	return nil
}

func (bd *FvBD) GetTotalCount() string {
	return bd.TotalCount
}

func (bd *FvBD) GetImData() []IMDATA {
	return bd.ImData
}

func (bd *FvBD) GetAttributes() (map[string]string, error) {
	returnMap := map[string]string{}
	attr, err := json.Marshal(bd.ImData[0].FvBD.Attributes)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(attr, &returnMap)
	return returnMap, nil
}

func (bd *FvBD) Class() string {
	return "fvBD"
}
