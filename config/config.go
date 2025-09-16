package config

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/creasty/defaults"
	"github.com/go-playground/validator/v10"
	"github.com/sethvargo/go-envconfig"
	"gopkg.in/yaml.v2"
)

var (
	ErrFailedToLoadEnv      = errors.New("failed to load env")
	ErrFailedToOpenFile     = errors.New("failed to open file")
	ErrFailedToDecodeStruct = errors.New("failed to decode struct")
)

func LoadEnv(ctx context.Context) (*Env, error) {
	var env Env
	if err := envconfig.Process(ctx, &env); err != nil {
		return nil, fmt.Errorf("%w: %w", ErrFailedToLoadEnv, err)
	}
	return &env, nil
}

func LoadConfig(ctx context.Context, path string) (*Config, error) {
	var config Config

	file, err := os.Open(path)
	if err != nil {
		return nil, ErrFailedToOpenFile
	}
	defer file.Close()

	if err = yaml.NewDecoder(file).Decode(&config); err != nil {
		return nil, ErrFailedToDecodeStruct
	}

	if err := validate(&config); err != nil {
		return nil, err
	}

	if err := defaults.Set(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

func validate(config *Config) error {
	v := validator.New()

	switch {
	case config.NotifierConfig.Slack != nil && config.NotifierConfig.Slack.Enabled:
		if err := v.Struct(config.NotifierConfig.Slack); err != nil {
			return err
		}
	case config.NotifierConfig.Discord != nil && config.NotifierConfig.Discord.Enabled:
		if err := v.Struct(config.NotifierConfig.Discord); err != nil {
			return err
		}
	default:
		return ErrNotFoundNotifierConfig
	}

	return v.Struct(config)
}
