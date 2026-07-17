package config

import (
	"fmt"
	"os"
	"flag"

	"gopkg.in/yaml.v3"
)

func Load() (*Config, error) {

	configPath := flag.String(
		"config",
		"./configs/config.yaml",
		"path to config yaml",
	)

	flag.Parse()

	data, err := os.ReadFile(*configPath)
	if err != nil {
		return nil, fmt.Errorf("read config: %w", err)
	}

	var cfg Config

	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("parse config: %w", err)
	}

	return &cfg, nil
}