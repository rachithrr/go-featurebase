package gofb

type Options struct {
	Host string
	Port string
}

func (opt *Options) init() {
	if opt.Host == "" {
		opt.Host = "localhost"
	}
	if opt.Port == "" {
		opt.Port = "10101"
	}

}
