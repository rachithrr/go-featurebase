package gofb

type Options struct {
	Host     string
	Port     string
	QueryURL string
	APIKey   string
}

func (opt *Options) init() {
	if opt.Host == "" {
		opt.Host = "localhost"
	}
	if opt.Port == "" {
		opt.Port = "10101"
	}
	if opt.QueryURL == "" {
		opt.QueryURL = "http://localhost:10101/sql"
	}
	if opt.APIKey == "" {
		opt.APIKey = "default_api_key"
	}
}
}
