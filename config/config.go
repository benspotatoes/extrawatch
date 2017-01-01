package config

import "github.com/koding/multiconfig"

type Extrawatch struct {
	HostPort string `default:"localhost:8000"`
	Db       string `default:"host=localhost port=5432 dbname=extrawatch user=extrawatch sslmode=disable"`
	Cors     bool   `default:"true"`
}

func NewConfig() *Extrawatch {
	config := &Extrawatch{}
	multi := multiconfig.New()
	multi.MustLoad(config)
	return config
}
