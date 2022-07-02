package lazerpay

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	BASE_URL = "https://api.lazerpay.engineering/api/v1"
)

type Config struct {
	ApiPubKey string
	ApiSecKey string
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
	"payment_link": {
		"create": "/payment-links",
		"update": "/payment-links",
		"fetch":  "/payment-links",
	},
	"transfer": {
		"transfer-crypto": "/transfer",
	},
}

func (ctx *Context) GetEndpoint(endpoint []string) string {
	return Endpoints[endpoint[0]][endpoint[1]]
}

type Request struct {
	Method     string
	Route      []string
	Queries    map[string]string
	Identifier string
	Auth       string
}

func (ctx *Context) GenerateURL(route []string, queries map[string]string, identifier string) string {
	var queryStr string = "?"
	for k, v := range queries {
		queryStr += fmt.Sprintf("%s=%s&", k, v)
	}
	endpoint := ctx.GetEndpoint(route)
	if identifier != "" {
		return fmt.Sprintf(`%s%s/%s%s`, BASE_URL, endpoint, identifier, queryStr)
	}
	return fmt.Sprintf(`%s%s%s`, BASE_URL, endpoint, queryStr)
}

func (ctx Context) SendRequest(reqData Request, data interface{}) (b []byte, err error) {
	url := ctx.GenerateURL(reqData.Route, reqData.Queries, reqData.Identifier)
	var (
		req  *http.Request
		body *bytes.Buffer
	)
	if data != nil {
		jsonBytes, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		body = bytes.NewBuffer(jsonBytes)
		req, err = http.NewRequest(reqData.Method, url, body)
		if err != nil {
			return nil, err
		}
	} else {
		req, err = http.NewRequest(reqData.Method, url, nil)
		if err != nil {
			return nil, err
		}
	}

	switch reqData.Auth {
	case "PUB_KEY":
		req.Header.Add("x-api-key", ctx.Config.ApiPubKey)
	case "SEC_KEY":
		req.Header.Add("Authorization", fmt.Sprintf(`Bearer %s`, ctx.Config.ApiSecKey))
	case "PUB_SEC_KEY":
		req.Header.Add("x-api-key", ctx.Config.ApiPubKey)
		req.Header.Add("Authorization", fmt.Sprintf(`Bearer %s`, ctx.Config.ApiSecKey))
	default:
		{
			req.Header.Add("x-api-key", ctx.Config.ApiPubKey)
			req.Header.Add("Authorization", fmt.Sprintf(`Bearer %s`, ctx.Config.ApiSecKey))
		}
	}

	req.Header.Add("Content-Type", "application/json")
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
	if resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusCreated {
		return b, nil
	}
	return b, APIError{
		Content: string(b),
	}
}

type APIError struct {
	Content string
}

func (e APIError) Error() string {
	return fmt.Sprintf(`unexpected error occured; error: %v;`, e.Content)
}

type CustomResp struct {
	Status     string      `json:"status"`
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
}
