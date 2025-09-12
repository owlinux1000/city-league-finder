package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/slack-go/slack"

	"github.com/owlinux1000/city-league-cancel-detector/client"
	"github.com/owlinux1000/city-league-cancel-detector/config"
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

	cl, err := client.NewClient(ctx, cfg.Endpoint)
	if err != nil {
		return err
	}

	slackClient := slack.New(env.SlackToken)

	params, err := client.NewEventSearchParams(cfg)
	if err != nil {
		return err
	}

	resp, err := cl.EventSearch(ctx, params)
	if err != nil {
		return err
	}

	if resp.Code != http.StatusOK {
		log.Println(resp.Message)
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
			cfg.SlackConfig.Channel,
			slack.MsgOptionText(
				fmt.Sprintf(
					"<@%s>\n:eyes: 2025/%s (%s) %s\nURL: %s\nAddress: %s",
					cfg.SlackConfig.MemberID,
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
