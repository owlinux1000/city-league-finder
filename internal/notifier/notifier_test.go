package notifier

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/owlinux1000/city-league-finder/config"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name          string
		env           *config.Env
		cfg           *config.NotifierConfig
		expectedError error
	}{
		{
			name: "success with slack",
			env: &config.Env{
				SlackToken: "test",
			},
			cfg: &config.NotifierConfig{
				Slack: &config.SlackConfig{
					MemberID: "test",
					Channel:  "#test",
					Enabled:  true,
				},
			},
			expectedError: nil,
		},
		{
			name: "success with discord",
			env:  &config.Env{},
			cfg: &config.NotifierConfig{
				Discord: &config.DiscordConfig{
					MemberID: "test",
					Webhook:  "https://example.com",
					Enabled:  true,
				},
			},
			expectedError: nil,
		},
		{
			name:          "failure with invalid notifier kind",
			env:           &config.Env{},
			cfg:           &config.NotifierConfig{},
			expectedError: ErrNotFoundNotifierConfig,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := new(tt.env, tt.cfg)
			if tt.expectedError != nil {
				assert.ErrorIs(t, err, tt.expectedError)
			}
		})
	}
}
