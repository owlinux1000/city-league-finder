package discord

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/owlinux1000/city-league-finder/config"
	"github.com/owlinux1000/city-league-finder/internal/notifier/model"
)

func New(config *config.DiscordConfig) model.Notifier {
	return &Notifier{
		config: config,
	}
}

func (n *Notifier) PostMessage(message string) error {
	if n.config.MemberID != "" {
		message = fmt.Sprintf("<@!%s> %s", n.config.MemberID, message)
	}

	body := map[string]any{
		"content": message,
		"allowed_mentions": map[string][]string{
			"parse": {"users"},
		},
	}

	jsonData, err := json.Marshal(body)
	if err != nil {
		return err
	}

	resp, err := http.Post(n.config.Webhook, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return fmt.Errorf("%w: %s", err, string(body))
	}

	return nil
}
