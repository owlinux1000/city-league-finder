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
	if err := v.Struct(config); err != nil {
		return err
	}

	for _, kind := range config.Notifier {
		switch kind {
		case "slack":
			if config.SlackConfig == (SlackConfig{}) {
				return fmt.Errorf("%w: %s", ErrNotFoundNotifierConfig, kind)
			}
			if err := v.Struct(config.SlackConfig); err != nil {
				return err
			}
		case "discord":
			if config.DiscordConfig == (DiscordConfig{}) {
				return fmt.Errorf("%w: %s", ErrNotFoundNotifierConfig, kind)
			}
			if err := v.Struct(config.DiscordConfig); err != nil {
				return err
			}
		}
	}

	return v.Struct(config)
}
