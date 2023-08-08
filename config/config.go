package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

var conf Config

func NewConfig() Config {
	err := envconfig.Process("", &conf)
	if err != nil {
		log.Fatal(err.Error())
	}
	return conf
}

type Config struct {
	Pubsub PubsubConfig
}

// env: PUBSUB_XXX
type PubsubConfig struct {
	ProjectID string `split_words:"true"`
}
