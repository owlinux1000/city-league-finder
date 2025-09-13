package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/sync/errgroup"

	"github.com/owlinux1000/city-league-cancel-detector/client"
	"github.com/owlinux1000/city-league-cancel-detector/config"
	"github.com/owlinux1000/city-league-cancel-detector/internal/notifier"
	"github.com/owlinux1000/city-league-cancel-detector/internal/notifier/model"
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

	notifiers := map[string]model.Notifier{}
	for _, kind := range cfg.Notifier {
		notifiers[kind], err = notifier.New(kind, env, cfg)
		if err != nil {
			return err
		}
	}

	message := ""
	for _, event := range resp.Events {
		eventURL := cl.EventURL(&event)
		message += fmt.Sprintf(
			"2025/%s (%s) %s\nURL: %s\nAddress: %s\n\n",
			event.EventDate,
			event.EventDateWeek,
			event.ShopName,
			eventURL,
			event.Address,
		)
	}

	g := new(errgroup.Group)
	for _, notifier := range notifiers {
		g.Go(
			func() error {
				if err := notifier.PostMessage(message); err != nil {
					return err
				}
				return nil
			},
		)
		if err := g.Wait(); err != nil {
			return err
		}
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
