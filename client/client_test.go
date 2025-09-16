package client

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/owlinux1000/city-league-finder/config"
)

func mockEventSearchServer(t *testing.T, resp *EventSearchResponse) *httptest.Server {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.URL.Path != "/event_search" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		params := r.URL.Query()

		if slices.Compare(params["prefecture[]"], []string{"13", "14"}) == 0 {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			if err := json.NewEncoder(w).Encode(resp); err != nil {
				t.Fatal(err)
			}
			return
		}

		w.WriteHeader(http.StatusNotFound)
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			t.Fatal(err)
		}
	}))
	return ts
}
func TestNewEventSearch(t *testing.T) {
	tests := []struct {
		name          string
		params        EventSearchParams
		expected      EventSearchResponse
		expectedError error
	}{
		{
			name: "success with open league",
			params: EventSearchParams{
				Prefecture: []string{"13", "14"},
				LeagueType: []string{"1"},
				Offset:     "0",
				Accepting:  "true",
				Order:      "1",
				EventType:  []string{"3:2"},
			},
			expected: EventSearchResponse{
				Code:       200,
				EventCount: 1,
				Events: []Event{
					{
						ID:              878,
						DateID:          1551135,
						ShopID:          10459,
						EventDateParams: "20250923",
						EventDate:       "09/23",
						EventStartedAt:  "11:00",
						EventEndedAt:    "",
						PrefectureName:  "東京都",
						Venue:           "",
						LeagueName:      "オープン",
					},
				},
			},
			expectedError: nil,
		},
		{
			name: "not found result",
			params: EventSearchParams{
				Prefecture: []string{"1"},
				LeagueType: []string{"1"},
				Offset:     "0",
				Accepting:  "true",
				Order:      "1",
				EventType:  []string{"3:2"},
			},
			expected: EventSearchResponse{
				Code:    400,
				Message: "イベント情報が存在しません。",
			},
			expectedError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := mockEventSearchServer(t, &tt.expected)
			defer ts.Close()

			ctx := context.Background()
			cl, err := NewClient(ctx, ts.URL)
			if err != nil {
				t.Fatal(err)
			}

			resp, err := cl.EventSearch(ctx, &tt.params)
			if tt.expectedError != nil {
				assert.ErrorIs(t, err, tt.expectedError)
			} else {
				assert.Equal(t, tt.expected, *resp)
			}
		})
	}
}

func TestNewEventSearchParams(t *testing.T) {
	tests := []struct {
		name          string
		config        *config.Config
		expected      *EventSearchParams
		expectedError error
	}{
		{
			name: "success",
			config: &config.Config{
				Prefecture: []string{"Tokyo", "Osaka"},
				LeagueKind: []string{"open", "junior"},
				Offset:     0,
				Order:      1,
				Accepting:  true,
			},
			expected: &EventSearchParams{
				Prefecture: []string{"13", "27"},
				LeagueType: []string{"1", "2"},
				Offset:     "0",
				Order:      "1",
				Accepting:  "true",
				EventType:  []string{"3:2"},
			},
			expectedError: nil,
		},
		{
			name: "failure with invalid prefecture name",
			config: &config.Config{
				Prefecture: []string{"Tokyo", "invalid"},
				LeagueKind: []string{"open", "junior"},
				Offset:     0,
				Order:      1,
				Accepting:  true,
			},
			expected:      nil,
			expectedError: ErrNotFoundPrefectureName,
		},
		{
			name: "failure with invalid league type",
			config: &config.Config{
				Prefecture: []string{"Tokyo"},
				LeagueKind: []string{"open", "junior", "invalid"},
				Offset:     0,
				Order:      1,
				Accepting:  true,
			},
			expected:      nil,
			expectedError: ErrInvalidLeagueType,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			params, err := NewEventSearchParams(tt.config)
			if tt.expectedError != nil {
				assert.ErrorIs(t, err, tt.expectedError)
			} else {
				assert.Equal(t, tt.expected, params)
			}
		})
	}
}
