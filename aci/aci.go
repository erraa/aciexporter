package aci

import (
	"fmt"
	"io/ioutil"
)

type AciObject interface {
	GetTotalCount() string
	GetImData() []IMDATA
	GetAttributes() (map[string]string, error)
	UnmarshalJson([]byte) error
	Class() string
}

func Get(o AciObject, client *Client) (AciObject, error) {
	resp, err := client.get(fmt.Sprintf(
		"%s%s%s",
		"/api/class/",
		o.Class(),
		".json",
	))
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}

	err = o.UnmarshalJson(body)

	if err != nil {
		return nil, err
	}

	return o, nil
}
