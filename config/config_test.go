package config

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadEnv(t *testing.T) {
	tests := []struct {
		name          string
		env           map[string]string
		expected      *Env
		expectedError error
	}{
		{
			name: "success with the required environment variables",
			env: map[string]string{
				"CONFIG_PATH": "/path/to/config.yaml",
				"SLACK_TOKEN": "censord",
			},
			expected: &Env{
				ConfigPath: "/path/to/config.yaml",
				SlackToken: "censord",
			},
			expectedError: nil,
		},
		{
			name: "failure with missing CONFIG_PATH",
			env: map[string]string{
				"SLACK_TOKEN": "censord",
			},
			expected:      nil,
			expectedError: ErrFailedToLoadEnv,
		},
		{
			name: "failure with missing SLACK_TOKEN",
			env: map[string]string{
				"CONFIG_PATH": "/path/to/config.yaml",
			},
			expected:      nil,
			expectedError: ErrFailedToLoadEnv,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for k, v := range tt.env {
				os.Setenv(k, v)
			}

			ctx := context.Background()
			env, err := LoadEnv(ctx)

			if tt.expectedError != nil {
				assert.ErrorIs(t, err, tt.expectedError)
			} else {
				assert.Equal(t, tt.expected, env)
			}

			for k := range tt.env {
				os.Unsetenv(k)
			}

		})
	}
}

func TestLoadConfig(t *testing.T) {
	tests := []struct {
		name          string
		path          string
		expected      *Config
		expectedError error
	}{
		{
			name: "success with the required config",
			path: "testdata/config.yaml",
			expected: &Config{
				Endpoint:   "https://players.pokemon-card.com",
				Prefecture: []string{"Tokyo", "Osaka"},
				LeagueKind: []string{"open", "junior"},
				Offset:     0,
				Order:      1,
				Accepting:  true,
				SlackConfig: SlackConfig{
					MemberID: "test",
					Channel:  "#test",
				},
			},
		},
		{
			name:          "failure with non existing config",
			path:          "/path/to/non_existing_config.yaml",
			expected:      nil,
			expectedError: ErrFailedToOpenFile,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			ctx := context.Background()
			env, err := LoadConfig(ctx, tt.path)

			if tt.expectedError != nil {
				assert.ErrorIs(t, err, tt.expectedError)
			} else {
				assert.Equal(t, tt.expected, env)
			}
		})
	}
}
