package slack

import (
	"fmt"

	"github.com/slack-go/slack"

	"github.com/owlinux1000/city-league-cancel-detector/config"
	"github.com/owlinux1000/city-league-cancel-detector/internal/notifier/model"
)

func New(token string, config *config.SlackConfig) model.Notifier {
	return &Notifier{
		client: slack.New(token),
		config: config,
	}
}

func (n *Notifier) PostMessage(message string) error {
	if n.config.MemberID != "" {
		message = fmt.Sprintf("<@%s> %s", n.config.MemberID, message)
	}
	_, _, err := n.client.PostMessage(
		n.config.Channel,
		slack.MsgOptionText(
			message,
			false,
		),
	)
	return err
}
