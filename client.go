package gofb

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
)

type Client struct {
	opt *Options
}

func NewClient(opt *Options) *Client {
	opt.init()
	return &Client{opt: opt}
}

func (c *Client) Query(query string) (string, error) {

	resp, err := http.Post("http://"+c.opt.Host+":"+c.opt.Port, "application/x-www-form-urlencoded", bytes.NewBuffer([]byte(query)))
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", errors.New(resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
