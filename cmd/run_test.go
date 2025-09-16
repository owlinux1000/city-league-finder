package cmd

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/owlinux1000/city-league-finder/client"
	mocksClient "github.com/owlinux1000/city-league-finder/client/mocks"
	"github.com/owlinux1000/city-league-finder/config"
	"github.com/owlinux1000/city-league-finder/internal/notifier/model"
	mocksNotifier "github.com/owlinux1000/city-league-finder/internal/notifier/model/mocks"
)

func TestRun(t *testing.T) {
	tests := []struct {
		name          string
		env           *config.Env
		cfg           *config.Config
		notifiers     []model.Notifier
		expectedError error
		setup         func(cl *mocksClient.MockClientAdapter, env *config.Env, cfg *config.Config) []model.Notifier
	}{
		{
			name: "success",
			env: &config.Env{
				ConfigPath: "/path/to/config.yaml",
				SlackToken: "test",
			},
			cfg: &config.Config{
				Endpoint:   "https://example.com",
				Prefecture: []string{"Tokyo", "Osaka"},
				LeagueKind: []string{"open", "junior"},
				Offset:     0,
				Order:      1,
				Accepting:  true,
				NotifierConfig: config.NotifierConfig{
					Slack: &config.SlackConfig{
						MemberID: "test",
						Channel:  "#test",
						Enabled:  true,
					},
					Discord: &config.DiscordConfig{
						Webhook: "https://example.com",
						Enabled: true,
					},
				},
			},
			notifiers:     []model.Notifier{},
			expectedError: nil,
			setup: func(cl *mocksClient.MockClientAdapter, env *config.Env, cfg *config.Config) []model.Notifier {
				resp := client.EventSearchResponse{
					Code:       200,
					EventCount: 1,
					Events: []client.Event{
						{
							ID:              1,
							DateID:          1,
							ShopID:          1,
							ShopName:        "test",
							EventDateParams: "20250923",
							EventDate:       "09/23",
							EventDateWeek:   "日",
							EventStartedAt:  "11:00",
							EventEndedAt:    "",
							Address:         "東京都",
							PrefectureName:  "東京都",
							Venue:           "",
							LeagueName:      "オープン",
						},
					},
				}

				params, err := client.NewEventSearchParams(cfg)
				if err != nil {
					t.Fatal(err)
				}
				cl.EXPECT().EventSearch(mock.Anything, params).Return(&resp, nil)

				eventURLs := []string{
					"https://example.com/event/detail/1/1/1/20250923/1",
				}

				message := ""
				for i, event := range resp.Events {
					cl.EXPECT().EventURL(&event).Return(eventURLs[i])
					date, err := time.Parse("20060102", resp.Events[0].EventDateParams)
					if err != nil {
						t.Fatal(err)
					}
					message += fmt.Sprintf(
						"%s (%s) %s\nURL: %s\nAddress: %s\n\n",
						date.Format("2006/01/02"),
						resp.Events[0].EventDateWeek,
						resp.Events[0].ShopName,
						eventURLs[i],
						resp.Events[0].Address,
					)
				}

				notifiers := []model.Notifier{}
				mocksNotifier := mocksNotifier.NewMockNotifier(t)
				mocksNotifier.EXPECT().PostMessage(message).Return(nil)
				notifiers = append(notifiers, mocksNotifier, mocksNotifier)
				return notifiers
			},
		},
		{
			name: "not found notifier",
			env: &config.Env{
				ConfigPath: "/path/to/config.yaml",
			},
			cfg: &config.Config{
				Endpoint:       "https://example.com",
				NotifierConfig: config.NotifierConfig{},
			},
			expectedError: ErrNotFoundNotifierConfig,
		},
		{
			name: "not found events",
			env: &config.Env{
				ConfigPath: "/path/to/config.yaml",
				SlackToken: "test",
			},
			cfg: &config.Config{
				Endpoint:   "https://example.com",
				Prefecture: []string{"Tokyo", "Osaka"},
				LeagueKind: []string{"open", "junior"},
				Offset:     0,
				Order:      1,
				Accepting:  true,
				NotifierConfig: config.NotifierConfig{
					Slack: &config.SlackConfig{
						MemberID: "test",
						Channel:  "#test",
						Enabled:  true,
					},
					Discord: &config.DiscordConfig{
						MemberID: "test",
						Webhook:  "https://example.com",
						Enabled:  true,
					},
				},
			},
			expectedError: nil,
			setup: func(cl *mocksClient.MockClientAdapter, env *config.Env, cfg *config.Config) []model.Notifier {
				resp := client.EventSearchResponse{
					Code:    404,
					Message: "イベント情報が存在しません。",
				}

				params, err := client.NewEventSearchParams(cfg)
				if err != nil {
					t.Fatal(err)
				}
				cl.EXPECT().EventSearch(mock.Anything, params).Return(&resp, nil)
				return []model.Notifier{}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cl := &mocksClient.MockClientAdapter{}
			if tt.setup != nil {
				tt.notifiers = tt.setup(cl, tt.env, tt.cfg)
			}

			ctx := context.Background()
			err := run(ctx, cl, tt.cfg, tt.notifiers)
			if tt.expectedError != nil {
				assert.ErrorIs(t, err, tt.expectedError)
			}
		})
	}

}
