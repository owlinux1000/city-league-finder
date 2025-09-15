/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"

	"github.com/owlinux1000/city-league-finder/client"
	"github.com/owlinux1000/city-league-finder/config"
	"github.com/owlinux1000/city-league-finder/internal/notifier"
	"github.com/owlinux1000/city-league-finder/internal/notifier/model"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run this tool",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {

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

		notifiers := map[string]model.Notifier{}
		for _, kind := range cfg.Notifier {
			notifiers[kind], err = notifier.New(kind, env, cfg)
			if err != nil {
				return err
			}
		}

		if err := run(ctx, cl, cfg, notifiers); err != nil {
			return err
		}

		return nil
	},
}

func init() {
	runCmd.Flags().StringP("config", "c", "", "config file path")
	rootCmd.AddCommand(runCmd)
}

func run(ctx context.Context, cl *client.Client, cfg *config.Config, notifiers map[string]model.Notifier) error {
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

	message := ""
	for _, event := range resp.Events {
		eventURL := cl.EventURL(&event)

		date, err := time.Parse("20060102", event.EventDateParams)
		if err != nil {
			return err
		}

		message += fmt.Sprintf(
			"%s (%s) %s\nURL: %s\nAddress: %s\n\n",
			date.Format("2006/01/02"),
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
