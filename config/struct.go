package config

import (
	_ "github.com/creasty/defaults"
	_ "github.com/go-playground/validator/v10"
	_ "gopkg.in/yaml.v2"
)

type Env struct {
	ConfigPath string `env:"CONFIG_PATH,required"`
	SlackToken string `env:"SLACK_TOKEN"`
}

type Config struct {
	Endpoint      string        `yaml:"endpoint" default:"https://players.pokemon-card.com"`
	Prefecture    []string      `yaml:"prefecture" validate:"required,unique"`
	LeagueKind    []string      `yaml:"league" validate:"required,unique,dive,oneof=open junior senior master"`
	Offset        int           `yaml:"offset" default:"0" `
	Accepting     bool          `yaml:"accepting" default:"true"`
	Order         int           `yaml:"order" default:"1"`
	Notifier      []string      `yaml:"notifier" validate:"required,unique,dive,oneof=slack discord"`
	SlackConfig   SlackConfig   `yaml:"slack" validate:"-"`
	DiscordConfig DiscordConfig `yaml:"discord" validate:"-"`
}

type SlackConfig struct {
	Channel  string `yaml:"channel" validate:"required"`
	MemberID string `yaml:"memberID"`
}

type DiscordConfig struct {
	Webhook  string `yaml:"webhook" validate:"required"`
	MemberID string `yaml:"memberID"`
}
