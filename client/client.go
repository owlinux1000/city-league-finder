package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	pref "github.com/diverse-inc/jp_prefecture"

	"github.com/owlinux1000/city-league-cancel-detector/config"
)

var (
	LeagueType map[string]string = map[string]string{
		"open":   "1",
		"junior": "2",
		"senior": "3",
		"master": "4",
	}
)

func NewClient(ctx context.Context, endpoint string) (*Client, error) {
	client := &http.Client{
		Transport: http.DefaultTransport,
	}

	return &Client{
		client:   client,
		Endpoint: endpoint,
	}, nil
}

func (c *Client) EventSearch(ctx context.Context, params *EventSearchParams) (*EventSearchResponse, error) {

	query := url.Values{
		"prefecture[]":  params.Prefecture,
		"event_type[]":  params.EventType,
		"league_type[]": params.LeagueType,
		"offset":        []string{params.Offset},
		"accepting":     []string{params.Accepting},
		"order":         []string{params.Order},
	}

	resp, err := c.get(ctx, eventSearchPath, query)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	var eventSearchResponse EventSearchResponse
	err = json.Unmarshal(body, &eventSearchResponse)
	if err != nil {
		return nil, err
	}

	return &eventSearchResponse, nil
}

func (c *Client) get(ctx context.Context, path string, query url.Values) (*http.Response, error) {
	url := c.Endpoint + path
	if query != nil {
		url += "?" + query.Encode()
	}

	fmt.Println(url)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	return c.client.Do(req)
}

func (c *Client) EventURL(event *Event) string {
	eventDetailBaseURL := c.Endpoint + "/event/detail"
	elems := []string{
		eventDetailBaseURL,
		strconv.Itoa(event.EventHoldingID),
		"1",
		strconv.Itoa(event.ShopID),
		event.EventDateParams,
		strconv.Itoa(event.DateID),
	}
	eventURL := strings.Join(
		elems,
		"/",
	)
	return eventURL
}

func NewEventSearchParams(cfg *config.Config) (*EventSearchParams, error) {

	prefID := []string{}
	for _, prefecture := range cfg.Prefecture {
		prefInfo, ok := pref.FindByRoma(prefecture)
		if !ok {
			return nil, fmt.Errorf("%w: %s", ErrNotFoundPrefectureName, prefecture)
		}
		prefIDStr := strconv.Itoa(prefInfo.Code())
		prefID = append(prefID, prefIDStr)
	}

	leagueType := []string{""}
	for _, league := range cfg.LeagueKind {
		l, ok := LeagueType[league]
		if !ok {
			return nil, fmt.Errorf("%w: %s", ErrInvalidLeagueType, league)
		}
		leagueType = append(leagueType, l)
	}

	return &EventSearchParams{
		Prefecture: prefID,
		Offset:     strconv.Itoa(cfg.Offset),
		Accepting:  strconv.FormatBool(cfg.Accepting),
		Order:      strconv.Itoa(cfg.Order),
		EventType:  []string{"3:2"}, // City League
		LeagueType: leagueType,
	}, nil
}
