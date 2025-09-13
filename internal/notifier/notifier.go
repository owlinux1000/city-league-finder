package notifier

import (
	"github.com/owlinux1000/city-league-finder/config"
	"github.com/owlinux1000/city-league-finder/internal/notifier/discord"
	"github.com/owlinux1000/city-league-finder/internal/notifier/model"
	"github.com/owlinux1000/city-league-finder/internal/notifier/slack"
)

func New(kind string, env *config.Env, cfg *config.Config) (model.Notifier, error) {
	switch kind {
	case "slack":
		return slack.New(env.SlackToken, &cfg.SlackConfig), nil
	case "discord":
		return discord.New(&cfg.DiscordConfig), nil
	default:
		return nil, ErrInvalidNotifierKind
	}
}
