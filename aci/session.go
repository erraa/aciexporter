package aci

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"net/http"
	"net/http/cookiejar"
	"time"

	"fmt"

	"github.com/erraa/aciexporter/config"
	"github.com/gorilla/websocket"
)

var exporterConfig = config.GetConfig()

type Client struct {
	// hostname like https://apic.cisco.com
	hostname            string
	client              *http.Client
	loginToken          string
	loginRefreshTimeout time.Duration
	loginRefreshLast    time.Time
	insecure            bool
	socket              *websocket.Conn
}

type LoginStruct struct {
	AaaUser struct {
		Attributes struct {
			Name string `json:"name"`
			Pwd  string `json:"pwd"`
		} `json:"attributes"`
	} `json:"aaaUser"`
}

func (c *Client) Login(username, password string) error {
	apiUrl := fmt.Sprintf("%s%s", c.hostname, "/api/aaaLogin.json")
	login := LoginStruct{}
	login.AaaUser.Attributes.Name = username
	login.AaaUser.Attributes.Pwd = password

	data, err := json.Marshal(login)
	if err != nil {
		return err
	}

	resp, err := c.post(apiUrl, data)
	defer resp.Body.Close()

	if err != nil {
		return err
	}
	cookieJar, _ := cookiejar.New(nil)
	for _, cookie := range resp.Cookies() {
		if cookie.Name == "APIC-cookie" {
			cookieJar.SetCookies(resp.Request.URL, []*http.Cookie{cookie})
			c.client.Jar = cookieJar
		}
	}

	return err
}

func (c *Client) post(url string, body []byte) (*http.Response, error) {
	// Just set Insecure
	var method string = "POST"
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{
		InsecureSkipVerify: c.insecure,
	}

	data := bytes.NewReader(body)
	req, err := http.NewRequest(method, url, data)

	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(req)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s, %s",
			"statuscode not 200",
			string(resp.StatusCode),
		)
	}

	return resp, nil
}

func (c *Client) get(url string) (*http.Response, error) {
	var method string = "GET"
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{
		InsecureSkipVerify: c.insecure,
	}
	apiUrl := fmt.Sprintf("%s%s", c.hostname, url)
	req, err := http.NewRequest(method, apiUrl, nil)

	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(req)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s, %s",
			"statuscode not 200",
			string(resp.StatusCode),
		)
	}

	return resp, nil
}

func New(url string) *Client {
	insecure := true
	return &Client{
		insecure: insecure,
		hostname: url,
		client:   &http.Client{},
	}
}
