package config

import (
	"log"
	"os"
)

type ServerConfig struct {
	GoogleCloudProject string
	AutocodeURL        string
	AutocodeAuth       string
}

func NewServerConfig() *ServerConfig {
	cfg := &ServerConfig{}
	// Configuration parameters from environment variables.
	cfg.GoogleCloudProject = os.Getenv("GCP_PROJECT")
	if cfg.GoogleCloudProject == "" {
		log.Fatalln("GCP_PROJECT is not set")
	}
	cfg.AutocodeURL = os.Getenv("AUTOCODE_URL")
	if cfg.AutocodeURL == "" {
		log.Fatalln("AUTOCODE_URL is not set")
	}
	cfg.AutocodeAuth = os.Getenv("AUTOCODE_AUTH")
	if cfg.AutocodeAuth == "" {
		log.Fatalln("AUTOCODE_AUTH is not set")
	}

	return cfg
}
