package config

import (
	"context"
	"os"

	"github.com/sethvargo/go-envconfig"
	"gopkg.in/yaml.v2"
)

func LoadEnv(ctx context.Context) (*Env, error) {
	var env Env
	if err := envconfig.Process(ctx, &env); err != nil {
		return nil, err
	}
	return &env, nil
}

func LoadConfig(ctx context.Context, path string) (*Config, error) {
	var config Config

	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return nil, err
	}
	err = yaml.NewDecoder(file).Decode(&config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}