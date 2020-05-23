package config

import (
	"github.com/kelseyhightower/envconfig"
	"log"
)

type Configuration struct {
	Database Database `json:"database" required:"true"`
	API      Server   `json:"api" required:"true"`
	GRPC     Server   `json:"grpc" required:"true"`
}

func MustLoad() Configuration {
	configuration := Configuration{}
	err := envconfig.Process("GMS", &configuration)
	if err != nil {
		log.Panicf("can't load configuration: %v", err)
	}
	return configuration
}
