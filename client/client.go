package client

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)


func NewClient(ctx context.Context, endpoint string) (*Client, error) {
	client := &http.Client{
		Transport: http.DefaultTransport,
	}

	return &Client{
		client: client,
		Endpoint: endpoint,
	}, nil
}

func (c *Client) EventSearch(ctx context.Context, query url.Values) (*EventSearchResponse, error) {
	resp, err := c.get(ctx, eventSearchPath, query)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
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
	
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	return c.client.Do(req)
}