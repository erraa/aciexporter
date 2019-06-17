package aci

import (
	"encoding/json"
	"time"
)

//func (e Epg) Class() string {
//	return "fvAEPg"
//}

type FvAEPg struct {
	ImData     []IMDATA `json:"imdata"`
	TotalCount string   `json:"totalCount"`
}

//type IMDATA struct {
//	fvAEPg FVAEPG `json:"fvAEPg"`
//}

func (epg *FvAEPg) UnmarshalJson(data []byte) error {
	err := json.Unmarshal(data, &epg)
	if err != nil {
		return err
	}
	return nil
}

func (epg *FvAEPg) GetTotalCount() string {
	return epg.TotalCount
}

func (epg *FvAEPg) GetImData() []IMDATA {
	return epg.ImData
}

func (epg *FvAEPg) GetAll() ([]map[string]string, error) {
	returnSlice := []map[string]string{}
	for _, obj := range epg.ImData {
		jsonAttr := map[string]string{}
		attr, _ := json.Marshal(obj.FvAEPg.Attributes)
		err := json.Unmarshal(attr, &jsonAttr)
		if err != nil {
			return nil, err
		}
		returnSlice = append(returnSlice, jsonAttr)
	}
	return returnSlice, nil
}

func (epg *FvAEPg) GetAttributes() (map[string]string, error) {
	returnMap := map[string]string{}
	attr, err := json.Marshal(epg.ImData[0].FvAEPg.Attributes)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(attr, &returnMap)
	return returnMap, nil
}

func (epg *FvAEPg) Class() string {
	return "fvAEPg"
}

type FVAEPG struct {
	Attributes struct {
		ChildAction         string    `json:"childAction"`
		ConfigIssues        string    `json:"configIssues"`
		ConfigSt            string    `json:"configSt"`
		Descr               string    `json:"descr"`
		Dn                  string    `json:"dn"`
		ExtMngdBy           string    `json:"extMngdBy"`
		FwdCtrl             string    `json:"fwdCtrl"`
		IsAttrBasedEPg      string    `json:"isAttrBasedEPg"`
		IsSharedSrvMsiteEPg string    `json:"isSharedSrvMsiteEPg"`
		LcOwn               string    `json:"lcOwn"`
		MatchT              string    `json:"matchT"`
		ModTs               time.Time `json:"modTs"`
		MonPolDn            string    `json:"monPolDn"`
		Name                string    `json:"name"`
		NameAlias           string    `json:"nameAlias"`
		PcEnfPref           string    `json:"pcEnfPref"`
		PcTag               string    `json:"pcTag"`
		PrefGrMemb          string    `json:"prefGrMemb"`
		Prio                string    `json:"prio"`
		Scope               string    `json:"scope"`
		Status              string    `json:"status"`
		TriggerSt           string    `json:"triggerSt"`
		TxID                string    `json:"txId"`
		UID                 string    `json:"uid"`
	} `json:"attributes"`
	Children []struct {
		FvCEp struct {
			Attributes struct {
				ChildAction string    `json:"childAction"`
				ContName    string    `json:"contName"`
				Encap       string    `json:"encap"`
				ID          string    `json:"id"`
				Idepdn      string    `json:"idepdn"`
				IP          string    `json:"ip"`
				LcC         string    `json:"lcC"`
				LcOwn       string    `json:"lcOwn"`
				Mac         string    `json:"mac"`
				McastAddr   string    `json:"mcastAddr"`
				ModTs       time.Time `json:"modTs"`
				MonPolDn    string    `json:"monPolDn"`
				Name        string    `json:"name"`
				NameAlias   string    `json:"nameAlias"`
				Rn          string    `json:"rn"`
				Status      string    `json:"status"`
				UID         string    `json:"uid"`
				UUID        string    `json:"uuid"`
				VmmSrc      string    `json:"vmmSrc"`
			} `json:"attributes"`
			Children []struct {
				FvIP struct {
					Attributes struct {
						Addr        string    `json:"addr"`
						ChildAction string    `json:"childAction"`
						LcOwn       string    `json:"lcOwn"`
						ModTs       time.Time `json:"modTs"`
						MonPolDn    string    `json:"monPolDn"`
						Rn          string    `json:"rn"`
						Status      string    `json:"status"`
						UID         string    `json:"uid"`
					} `json:"attributes"`
					Children []struct {
						FvReportingNode struct {
							Attributes struct {
								ChildAction string    `json:"childAction"`
								ID          string    `json:"id"`
								LcC         string    `json:"lcC"`
								LcOwn       string    `json:"lcOwn"`
								ModTs       time.Time `json:"modTs"`
								Rn          string    `json:"rn"`
								Status      string    `json:"status"`
							} `json:"attributes"`
						} `json:"fvReportingNode"`
					} `json:"children"`
				} `json:"fvIp,omitempty"`
				FvRsCEpToPathEp struct {
					Attributes struct {
						ChildAction  string    `json:"childAction"`
						ForceResolve string    `json:"forceResolve"`
						LcC          string    `json:"lcC"`
						LcOwn        string    `json:"lcOwn"`
						ModTs        time.Time `json:"modTs"`
						RType        string    `json:"rType"`
						Rn           string    `json:"rn"`
						State        string    `json:"state"`
						StateQual    string    `json:"stateQual"`
						Status       string    `json:"status"`
						TCl          string    `json:"tCl"`
						TDn          string    `json:"tDn"`
						TType        string    `json:"tType"`
					} `json:"attributes"`
				} `json:"fvRsCEpToPathEp,omitempty"`
			} `json:"children"`
		} `json:"fvCEp,omitempty"`
		FvRsGraphDef struct {
			Attributes struct {
				ChildAction  string    `json:"childAction"`
				ForceResolve string    `json:"forceResolve"`
				LcOwn        string    `json:"lcOwn"`
				ModTs        time.Time `json:"modTs"`
				RType        string    `json:"rType"`
				Rn           string    `json:"rn"`
				State        string    `json:"state"`
				StateQual    string    `json:"stateQual"`
				Status       string    `json:"status"`
				TCl          string    `json:"tCl"`
				TDn          string    `json:"tDn"`
				TType        string    `json:"tType"`
			} `json:"attributes"`
		} `json:"fvRsGraphDef,omitempty"`
		FvRsProv struct {
			Attributes struct {
				ChildAction      string    `json:"childAction"`
				CtrctUpd         string    `json:"ctrctUpd"`
				ExtMngdBy        string    `json:"extMngdBy"`
				ForceResolve     string    `json:"forceResolve"`
				LcOwn            string    `json:"lcOwn"`
				MatchT           string    `json:"matchT"`
				ModTs            time.Time `json:"modTs"`
				MonPolDn         string    `json:"monPolDn"`
				Prio             string    `json:"prio"`
				RType            string    `json:"rType"`
				Rn               string    `json:"rn"`
				State            string    `json:"state"`
				StateQual        string    `json:"stateQual"`
				Status           string    `json:"status"`
				TCl              string    `json:"tCl"`
				TContextDn       string    `json:"tContextDn"`
				TDn              string    `json:"tDn"`
				TRn              string    `json:"tRn"`
				TType            string    `json:"tType"`
				TnVzBrCPName     string    `json:"tnVzBrCPName"`
				TriggerSt        string    `json:"triggerSt"`
				UID              string    `json:"uid"`
				UpdateCollection string    `json:"updateCollection"`
			} `json:"attributes"`
			Children []struct {
				FvCollectionCont struct {
					Attributes struct {
						ChildAction  string    `json:"childAction"`
						CollectionDn string    `json:"collectionDn"`
						LcOwn        string    `json:"lcOwn"`
						ModTs        time.Time `json:"modTs"`
						MonPolDn     string    `json:"monPolDn"`
						Name         string    `json:"name"`
						NameAlias    string    `json:"nameAlias"`
						Rn           string    `json:"rn"`
						Status       string    `json:"status"`
					} `json:"attributes"`
				} `json:"fvCollectionCont"`
			} `json:"children"`
		} `json:"fvRsProv,omitempty"`
		FvRsPathAtt struct {
			Attributes struct {
				ChildAction  string    `json:"childAction"`
				Descr        string    `json:"descr"`
				Encap        string    `json:"encap"`
				ExtMngdBy    string    `json:"extMngdBy"`
				ForceResolve string    `json:"forceResolve"`
				InstrImedcy  string    `json:"instrImedcy"`
				LcC          string    `json:"lcC"`
				LcOwn        string    `json:"lcOwn"`
				ModTs        time.Time `json:"modTs"`
				Mode         string    `json:"mode"`
				MonPolDn     string    `json:"monPolDn"`
				PrimaryEncap string    `json:"primaryEncap"`
				RType        string    `json:"rType"`
				Rn           string    `json:"rn"`
				State        string    `json:"state"`
				StateQual    string    `json:"stateQual"`
				Status       string    `json:"status"`
				TCl          string    `json:"tCl"`
				TDn          string    `json:"tDn"`
				TType        string    `json:"tType"`
				UID          string    `json:"uid"`
			} `json:"attributes"`
		} `json:"fvRsPathAtt,omitempty"`
		FvRsDomAtt struct {
			Attributes struct {
				ChildAction  string    `json:"childAction"`
				ClassPref    string    `json:"classPref"`
				ConfigIssues string    `json:"configIssues"`
				Delimiter    string    `json:"delimiter"`
				Encap        string    `json:"encap"`
				EncapMode    string    `json:"encapMode"`
				EpgCos       string    `json:"epgCos"`
				EpgCosPref   string    `json:"epgCosPref"`
				ExtMngdBy    string    `json:"extMngdBy"`
				ForceResolve string    `json:"forceResolve"`
				InstrImedcy  string    `json:"instrImedcy"`
				LcOwn        string    `json:"lcOwn"`
				ModTs        time.Time `json:"modTs"`
				Mode         string    `json:"mode"`
				MonPolDn     string    `json:"monPolDn"`
				NetflowDir   string    `json:"netflowDir"`
				NetflowPref  string    `json:"netflowPref"`
				PrimaryEncap string    `json:"primaryEncap"`
				RType        string    `json:"rType"`
				ResImedcy    string    `json:"resImedcy"`
				Rn           string    `json:"rn"`
				State        string    `json:"state"`
				StateQual    string    `json:"stateQual"`
				Status       string    `json:"status"`
				TCl          string    `json:"tCl"`
				TDn          string    `json:"tDn"`
				TType        string    `json:"tType"`
				TriggerSt    string    `json:"triggerSt"`
				TxID         string    `json:"txId"`
				UID          string    `json:"uid"`
			} `json:"attributes"`
		} `json:"fvRsDomAtt,omitempty"`
		FvRsCons struct {
			Attributes struct {
				ChildAction      string    `json:"childAction"`
				CtrctUpd         string    `json:"ctrctUpd"`
				DeplInfo         string    `json:"deplInfo"`
				ExtMngdBy        string    `json:"extMngdBy"`
				ForceResolve     string    `json:"forceResolve"`
				LcOwn            string    `json:"lcOwn"`
				ModTs            time.Time `json:"modTs"`
				MonPolDn         string    `json:"monPolDn"`
				Prio             string    `json:"prio"`
				RType            string    `json:"rType"`
				Rn               string    `json:"rn"`
				State            string    `json:"state"`
				StateQual        string    `json:"stateQual"`
				Status           string    `json:"status"`
				TCl              string    `json:"tCl"`
				TContextDn       string    `json:"tContextDn"`
				TDn              string    `json:"tDn"`
				TRn              string    `json:"tRn"`
				TType            string    `json:"tType"`
				TnVzBrCPName     string    `json:"tnVzBrCPName"`
				TriggerSt        string    `json:"triggerSt"`
				UID              string    `json:"uid"`
				UpdateCollection string    `json:"updateCollection"`
			} `json:"attributes"`
			Children []struct {
				FvCollectionCont struct {
					Attributes struct {
						ChildAction  string    `json:"childAction"`
						CollectionDn string    `json:"collectionDn"`
						LcOwn        string    `json:"lcOwn"`
						ModTs        time.Time `json:"modTs"`
						MonPolDn     string    `json:"monPolDn"`
						Name         string    `json:"name"`
						NameAlias    string    `json:"nameAlias"`
						Rn           string    `json:"rn"`
						Status       string    `json:"status"`
					} `json:"attributes"`
				} `json:"fvCollectionCont"`
			} `json:"children"`
		} `json:"fvRsCons,omitempty"`
		FvRsCustQosPol struct {
			Attributes struct {
				ChildAction        string    `json:"childAction"`
				ForceResolve       string    `json:"forceResolve"`
				LcOwn              string    `json:"lcOwn"`
				ModTs              time.Time `json:"modTs"`
				MonPolDn           string    `json:"monPolDn"`
				RType              string    `json:"rType"`
				Rn                 string    `json:"rn"`
				State              string    `json:"state"`
				StateQual          string    `json:"stateQual"`
				Status             string    `json:"status"`
				TCl                string    `json:"tCl"`
				TContextDn         string    `json:"tContextDn"`
				TDn                string    `json:"tDn"`
				TRn                string    `json:"tRn"`
				TType              string    `json:"tType"`
				TnQosCustomPolName string    `json:"tnQosCustomPolName"`
				UID                string    `json:"uid"`
			} `json:"attributes"`
		} `json:"fvRsCustQosPol,omitempty"`
		FvRsBd struct {
			Attributes struct {
				ChildAction  string    `json:"childAction"`
				ExtMngdBy    string    `json:"extMngdBy"`
				ForceResolve string    `json:"forceResolve"`
				LcOwn        string    `json:"lcOwn"`
				ModTs        time.Time `json:"modTs"`
				MonPolDn     string    `json:"monPolDn"`
				RType        string    `json:"rType"`
				Rn           string    `json:"rn"`
				State        string    `json:"state"`
				StateQual    string    `json:"stateQual"`
				Status       string    `json:"status"`
				TCl          string    `json:"tCl"`
				TContextDn   string    `json:"tContextDn"`
				TDn          string    `json:"tDn"`
				TRn          string    `json:"tRn"`
				TType        string    `json:"tType"`
				TnFvBDName   string    `json:"tnFvBDName"`
				UID          string    `json:"uid"`
			} `json:"attributes"`
			Children []struct {
				FvSubnetBDDefCont struct {
					Attributes struct {
						BddefDn     string    `json:"bddefDn"`
						ChildAction string    `json:"childAction"`
						LcOwn       string    `json:"lcOwn"`
						ModTs       time.Time `json:"modTs"`
						MonPolDn    string    `json:"monPolDn"`
						Name        string    `json:"name"`
						NameAlias   string    `json:"nameAlias"`
						Rn          string    `json:"rn"`
						Status      string    `json:"status"`
					} `json:"attributes"`
				} `json:"fvSubnetBDDefCont"`
			} `json:"children"`
		} `json:"fvRsBd,omitempty"`
	} `json:"children"`
}
