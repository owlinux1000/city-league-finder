package discord

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/owlinux1000/city-league-cancel-detector/config"
	"github.com/owlinux1000/city-league-cancel-detector/internal/notifier/model"
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
	fmt.Println(string(jsonData))

	resp, err := http.Post(n.config.Webhook, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fmt.Println(resp.StatusCode)

	if resp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return fmt.Errorf("%w: %s", err, string(body))
	}

	return nil
}
