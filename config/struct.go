package config

type Env struct {
	ConfigPath string `env:"CONFIG_PATH,required"`
	SlackToken string `env:"SLACK_TOKEN,required"`
}

type Config struct {
	Endpoint string `yaml:"endpoint"`
	Prefecture []string `yaml:"prefecture"`
}