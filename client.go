package gofb

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Response struct {
	Schema struct {
		Fields []Field `json:"fields"`
	} `json:"schema"`
	Data          [][]interface{} `json:"data", omitempty`
	ExecutionTime int             `json:"execution-time"`
	Error         interface{}     `json:"error", omitempty`
}

type Field struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	BaseType string `json:"base-type"`
	TypeInfo string `json:"-"` // only decimal is using this field, and it is seen as float
}

type Client struct {
	opt *Options
	QueryURL string
	APIKey   string
}

func NewClient(opt *Options) *Client {
	opt.init()
	return &Client{opt: opt, QueryURL: opt.QueryURL, APIKey: opt.APIKey}
}

func (c *Client) Query(query string) (*Response, error) {

	u, err := url.Parse(c.QueryURL)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, c.QueryURL, bytes.NewBuffer([]byte(query)))
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	if u.APIKey != "" {
		req.Header.Add("Authorization", c.APIKey)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response Response

	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
