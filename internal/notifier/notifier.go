package notifier

import (
	"github.com/owlinux1000/city-league-finder/config"
	"github.com/owlinux1000/city-league-finder/internal/notifier/discord"
	"github.com/owlinux1000/city-league-finder/internal/notifier/model"
	"github.com/owlinux1000/city-league-finder/internal/notifier/slack"
)

func new(env *config.Env, cfg *config.NotifierConfig) (model.Notifier, error) {
	switch {
	case cfg.Slack != nil && cfg.Slack.Enabled:
		return slack.New(env.SlackToken, cfg.Slack), nil
	case cfg.Discord != nil && cfg.Discord.Enabled:
		return discord.New(cfg.Discord), nil
	default:
		return nil, ErrNotFoundNotifierConfig
	}
}

func NewNotifiers(env *config.Env, cfg *config.NotifierConfig) ([]model.Notifier, error) {
	notifiers := []model.Notifier{}
	if cfg.Slack != nil && cfg.Slack.Enabled {
		notifier, err := new(env, cfg)
		if err != nil {
			return []model.Notifier{}, err
		}
		notifiers = append(notifiers, notifier)
	}
	if cfg.Discord != nil && cfg.Discord.Enabled {
		notifier, err := new(env, cfg)
		if err != nil {
			return []model.Notifier{}, err
		}
		notifiers = append(notifiers, notifier)
	}
	return notifiers, nil
}
