package client

import "context"

type ClientAdapter interface {
	EventSearch(ctx context.Context, params *EventSearchParams) (*EventSearchResponse, error)
	EventURL(event *Event) string
}
