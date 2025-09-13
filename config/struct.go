package config

import (
	_ "github.com/creasty/defaults"
	_ "gopkg.in/yaml.v2"
)

type Env struct {
	ConfigPath string `env:"CONFIG_PATH,required"`
	SlackToken string `env:"SLACK_TOKEN,required"`
}

type Config struct {
	Endpoint    string      `yaml:"endpoint"`
	Prefecture  []string    `yaml:"prefecture"`
	LeagueKind  []string    `yaml:"league"`
	Offset      int         `yaml:"offset" default:"0"`
	Accepting   bool        `yaml:"accepting" default:"true"`
	Order       int         `yaml:"order" default:"1"`
	SlackConfig SlackConfig `yaml:"slack"`
}

type SlackConfig struct {
	Channel  string `yaml:"channel"`
	MemberID string `yaml:"memberID"`
}
