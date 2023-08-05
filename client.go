package gofb

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
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

type client struct {
	opt      *Options
	queryURL string
	apiKey   string
}

func NewClient(opt *Options) *client {
	opt.init()
	return &client{opt: opt, queryURL: opt.QueryURL, apiKey: opt.APIKey}
}

func (c *client) Query(query string) (*Response, error) {

	u, err := url.Parse(c.queryURL)
	if err != nil {
		return nil, errors.Join(err, errors.New("invalid query url"))
	}

	fmt.Printf("query url: %s\n", c.queryURL)

	req, err := http.NewRequest(http.MethodPost, c.queryURL, strings.NewReader(query))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "text/plain")
	req.Header.Add("Accept", "application/json")

	client := &http.Client{}
	if c.apiKey != "" && u.Scheme == "https" {
		req.Header.Add("X-API-Key", c.apiKey)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	fmt.Println(query)

	if resp.StatusCode != 200 {
		return nil, errors.Join(errors.New(resp.Status), errors.New("query failed"))
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
