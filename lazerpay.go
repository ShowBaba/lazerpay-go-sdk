package lazerpay

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	API_URL_LIVE = "https://api.lazerpay.engineering/api/v1"
)

var (
	BASE_URL string
)

type Config struct {
	ApiPubKey string
	ApiSecKey string
	Live      bool
}

type Context struct {
	Config Config
}

func NewContext(c Config) Context {
	return Context{c}
}

var Endpoints = map[string]map[string]string{
	"payment": {
		"create": "/transaction/initialize",
		"verify": "/transaction/verify",
	},
}

func (ctx *Context) GetEndpoint(endpointType string, action string) string {
	return Endpoints[endpointType][action]
}

func (ctx *Context) GetBaseURL(c Config) string {
	if c.Live {
		return API_URL_LIVE
	}
	return "" // return a sandbox URL instead
}

func (ctx Context) SendRequest(mtd string, url string, data interface{}, params map[string]string, authType string) (b []byte, err error) {
	var (
		req  *http.Request
		body *bytes.Buffer
	)
	var paramStr string = "?"
	for k, v := range params {
		paramStr += fmt.Sprintf("%s=%s&", k, v)
	}
	url += paramStr
	switch mtd {
	case "POST":
		jsonBytes, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		body = bytes.NewBuffer(jsonBytes)
		req, err = http.NewRequest("POST", url, body)
		if err != nil {
			return nil, err
		}
	case "GET":
		req, err = http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, err
		}
	}

	switch authType {
	case "PUB_KEY":
		req.Header.Add("x-id-key", ctx.Config.ApiPubKey)
	case "SEC_KEY":
		req.Header.Add("Authorization", fmt.Sprintf(`Bearer %s`, ctx.Config.ApiSecKey))
	case "PUB_SEC_KEY":
		req.Header.Add("x-id-key", ctx.Config.ApiPubKey)
		req.Header.Add("Authorization", fmt.Sprintf(`Bearer %s`, ctx.Config.ApiSecKey))
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	b, err = ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	return b, APIError{
		StatusCode: resp.StatusCode,
		Content:    string(b),
	}
}

// error from lazerpay API
type APIError struct {
	StatusCode int
	Content    string
}

func (e APIError) Error() string {
	return fmt.Sprintf(`unexpected error occured; error: %v; code: %v\n`, e.Content, e.StatusCode)
}
