// Package config contains config related logic
package config

import (
	"os"

	"go.yaml.in/yaml/v4"
)

type BucketConfig struct {
	DefaultMaxSize     string
	DefaultMaxFileSize string
	AllowedMIMETypes   []string
}

type Config struct {
	Port             int
	StorageDirectory string
	BucketConfig     BucketConfig
	Database         string
}

func LoadConfig() (*Config, error) {
	b, err := os.ReadFile("config.yaml")
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := yaml.Load(b, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
