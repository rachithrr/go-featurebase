package gofb

type Client struct {
	opt *Options
}

func NewClient(opt *Options) *Client {
	opt.init()
	return &Client{opt: opt}
}

func (c *Client) Query(query string) (string, error) {
	return "", nil
}
