/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"runtime/debug"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var Version string

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display the version",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {

		buildInfo, ok := debug.ReadBuildInfo()
		if ok && Version == "" {
			Version = buildInfo.Main.Version
		}

		s := fmt.Sprintln("city-league-finder:")
		s += fmt.Sprintf("  Version: \t%s\n", Version)
		s += fmt.Sprintf("  Go version: \t%s\n", buildInfo.GoVersion)
		for _, setting := range buildInfo.Settings {
			switch setting.Key {
			case "vcs.revision":
				s += fmt.Sprintf("  Git commit: \t%s\n", setting.Value)
			case "vcs.time":
				format, err := time.Parse(time.RFC3339, setting.Value)
				if err != nil {
					return err
				}
				s += fmt.Sprintf("  Built: \t%s\n", format.Format(time.ANSIC))
			}
		}
		fmt.Println(strings.TrimSuffix(s, "\n"))

		return nil
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)

}
