package config

import (
	"log"
	"os"
)

type ServerConfig struct {
	GoogleCloudProject string
	HalodotapiURL      string
	CryptumAuth        string
}

func NewServerConfig() *ServerConfig {
	cfg := &ServerConfig{}
	// Configuration parameters from environment variables.
	cfg.GoogleCloudProject = os.Getenv("GCP_PROJECT")
	if cfg.GoogleCloudProject == "" {
		log.Fatalln("GCP_PROJECT is not set")
	}
	cfg.HalodotapiURL = os.Getenv("HALODOTAPI_URL")
	if cfg.HalodotapiURL == "" {
		log.Fatalln("HALODOTAPI_URL_PREFIX is not set")
	}
	cfg.CryptumAuth = os.Getenv("CRYPTUM_AUTH")
	if cfg.CryptumAuth == "" {
		log.Fatalln("CRYPTUM_AUTH is not set")
	}

	return cfg
}
