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
	err      error
	BASE_URL string
)

type Config struct {
	apiPubKey string
	apiSecKey string
	Live      bool
}

type Context struct {
	Config Config
}

func NewContext(c Config) Context {
	BASE_URL = GetBaseURL(c)
	return Context{c}
}

var Endpoints = map[string]map[string]string{
	"payment": {
		"create": fmt.Sprintf(`%s/transaction/initialize`, BASE_URL),
		"verify": fmt.Sprintf(`%s/transaction/verify`, BASE_URL),
	},
}

func (ctx *Context) GetEndpoint(endpointType string, action string) string {
	return Endpoints[endpointType][action]
}

func GetBaseURL(c Config) string {
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
	switch mtd {
	case "POST":
		jsonBytes, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		body = bytes.NewBuffer(jsonBytes)
	case "GET":
		body = nil
	}
	var paramStr string = "?"
	if params != nil {
		for k, v := range params {
			paramStr += fmt.Sprintf("%s=%s&", k, v)
		}
	}
	url += paramStr
	req, err = http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}

	switch authType {
	case "PUB_KEY":
		req.Header.Add("x-id-key", ctx.Config.apiPubKey)
	case "SEC_KEY":
		req.Header.Add("Authorization", fmt.Sprintf(`Bearer %s`, ctx.Config.apiSecKey))
	case "PUB_SEC_KEY":
		req.Header.Add("x-id-key", ctx.Config.apiPubKey)
		req.Header.Add("Authorization", fmt.Sprintf(`Bearer %s`, ctx.Config.apiSecKey))
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
	return fmt.Sprintf(`unexpected error occured; error: %v; code: %v`, e.Content, e.StatusCode)
}
