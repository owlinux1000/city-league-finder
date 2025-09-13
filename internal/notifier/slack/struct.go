package slack

import (
	"github.com/slack-go/slack"

	"github.com/owlinux1000/city-league-finder/config"
)

type Notifier struct {
	client *slack.Client
	config *config.SlackConfig
}
