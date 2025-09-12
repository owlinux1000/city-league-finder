package config

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/creasty/defaults"
	"github.com/sethvargo/go-envconfig"
	"gopkg.in/yaml.v2"
)

var (
	FailedToLoadEnv      = errors.New("failed to load env")
	FailedToOpenFile     = errors.New("failed to open file")
	FailedToDecodeStruct = errors.New("failed to decode struct")
)

func LoadEnv(ctx context.Context) (*Env, error) {
	var env Env
	if err := envconfig.Process(ctx, &env); err != nil {
		return nil, fmt.Errorf("%w: %w", FailedToLoadEnv, err)
	}
	return &env, nil
}

func LoadConfig(ctx context.Context, path string) (*Config, error) {
	var config Config

	file, err := os.Open(path)
	defer file.Close()
	
	if err != nil {
		return nil, FailedToOpenFile
	}

	if err = yaml.NewDecoder(file).Decode(&config); err != nil {
		return nil, FailedToDecodeStruct
	}

	if err := defaults.Set(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
