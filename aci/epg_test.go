package aci

import (
	"fmt"
	"testing"
)

func TestEpg(t *testing.T) {
	c := getconfig()
	client := New(c.APIC.BaseURI)
	err := client.Login(c.APIC.Username, c.APIC.Password)
	if err != nil {
		t.Errorf("%s", err)
	}
	e := &FvAEPg{}
	epgData, err := Get(e, client)
	attr, err := epgData.GetAll()
	if err != nil {
		t.Errorf("%s", err)
	}
	fmt.Println(attr)
}
