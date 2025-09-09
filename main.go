package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	pref "github.com/diverse-inc/jp_prefecture"
	"github.com/slack-go/slack"

	"github.com/owlinux1000/city-league-cancel-dtector/client"
	"github.com/owlinux1000/city-league-cancel-dtector/config"
)

var (
	NotFoundPrefectureName = errors.New("not found prefecture name")
)

func RealMain(args []string) error {

	ctx := context.Background()
	env, err := config.LoadEnv(ctx)
	if err != nil {
		return err
	}

	cfg, err := config.LoadConfig(ctx, env.ConfigPath)
	if err != nil {
		return err
	}

	client, err := client.NewClient(ctx, cfg.Endpoint)
	if err != nil {
		return err
	}

	slackClient := slack.New(env.SlackToken)

	prefID := []string{}
	for _, prefecture := range cfg.Prefecture {
		prefInfo, ok := pref.FindByRoma(prefecture)
		if !ok {
			return fmt.Errorf("%w: %s", NotFoundPrefectureName, prefecture)
		}
		prefIDStr := strconv.Itoa(prefInfo.Code())
		prefID = append(prefID, prefIDStr)
	}

	values := url.Values{
		"prefecture[]": prefID,
		"event_type[]": []string{"3:2"}, // city league
		"offset":       []string{"0"},
		"accepting":    []string{"true"},
		"order":        []string{"1"},
	}

	resp, err := client.EventSearch(ctx, values)
	if err != nil {
		return err
	}

	if resp.Code != http.StatusOK {
		slackClient.PostMessage(
			"#feed_test",
			slack.MsgOptionText(
				resp.Message,
				false,
			),
		)
		return nil
	}

	eventDetailBaseURL := cfg.Endpoint + "/event/detail"
	for _, event := range resp.Events {
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

		slackClient.PostMessage(
			"#feed_test",
			slack.MsgOptionText(
				fmt.Sprintf(
					":eyes: 2025/%s (%s) %s\nURL: %s\nAddress: %s",
					event.EventDate,
					event.EventDateWeek,
					event.ShopName,
					eventURL,					
					event.Address,
				),
				false,
			),
		)
	}

	return nil
}

func main() {
	exitStatus := 0
	if err := RealMain(os.Args); err != nil {
		fmt.Println(err.Error())
		exitStatus = 1
	}
	os.Exit(exitStatus)
}
