package aci

import (
	"io/ioutil"
	"testing"

	"github.com/erraa/aciexporter/config"
)

var client Client
var c config.Config

func getconfig() config.Config {
	return config.GetConfig()
}

func TestLogin(t *testing.T) {
	c := getconfig()
	client := New(c.APIC.BaseURI)
	err := client.Login(c.APIC.Username, c.APIC.Password)
	if err != nil {
		t.Errorf("%s", err)
	}
}

func TestGet(t *testing.T) {
	c := getconfig()
	client := New(c.APIC.BaseURI)
	err := client.Login(c.APIC.Username, c.APIC.Password)
	if err != nil {
		t.Errorf("%s", err)
	}
	resp, err := client.get("/api/class/fvAEPg.json")
	_, err = ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		t.Errorf("%s", err)
	}
}
