package cmd

import (
	"errors"
	"fmt"
	"time"

	"github.com/limeiralucas/chrono-cli/pkg/app"
	"github.com/limeiralucas/chrono-cli/pkg/config"
	toggl_provider "github.com/limeiralucas/chrono-cli/pkg/provider/toggl"
	"github.com/limeiralucas/chrono-cli/pkg/util"
	"github.com/spf13/cobra"
)

var togglCmd = &cobra.Command{
	Use:   "toggl",
	Short: "Access toggl time entries",
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all Time Entries",
	RunE: (func(cmd *cobra.Command, args []string) error {
		config, err := config.ReadConfig("./config.json")
		if err != nil {
			return err
		}

		week, err := cmd.Flags().GetInt8("week")
		if err != nil {
			return err
		}
		if week > 0 {
			return errors.New("week flag must be less or equal to 0")
		}
		weekDate := util.AddWeek(time.Now().UTC(), week)
		weekStart, weekEnd := util.GetWeekStartAndEnd(weekDate)

		provider := toggl_provider.NewTimeEntryProvider(config.Token)
		service := app.NewTimeEntryService(&provider)

		entries, err := service.List(weekStart, weekEnd)
		if err != nil {
			return err
		}

		for _, entry := range entries {
			fmt.Printf("%s: %s - %s\n", entry.Description, entry.StartDate, entry.EndDate)
		}

		return nil
	}),
}

func init() {
	listCmd.Flags().Int8P("week", "w", 0, "Week interval. Ex.: -1 (last week), -2 (two weeks ago)")

	togglCmd.AddCommand(listCmd)
}
