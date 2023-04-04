package cmd

import (
	"errors"
	"fmt"
	"time"

	"github.com/limeiralucas/chrono-cli/pkg/app"
	"github.com/limeiralucas/chrono-cli/pkg/config"
	"github.com/limeiralucas/chrono-cli/pkg/domain"
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

		elapsedTime, err := service.ElapsedTimeByDay(weekStart, weekEnd)
		if err != nil {
			return err
		}

		fmt.Printf("Week: %s - %s\n\n", weekStart.Format("02/01"), weekEnd.Format("02/01"))
		for day, elapsed := range elapsedTime {
			fmt.Printf("%s: %.2f\n", day, elapsed)
		}

		return nil
	}),
}

var reportCmd = &cobra.Command{
	Use:   "report",
	Short: "Time report of the week",
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

		timeReport, err := service.TimeReport(weekStart, weekEnd)
		if err != nil {
			return err
		}

		fmt.Printf("\nWeek: %s - %s\n", weekStart.Format("02/01"), weekEnd.Format("02/01"))
		for day, dayReport := range timeReport {
			other := []*domain.TimeEntry{}

			dayTotalDuration := float32(0)
			tagTotalDuration := float32(0)
			reportStr := ""
			entriesDetails := ""

			for tag, entries := range dayReport {
				if tag == "[other]" {
					other = append(other, entries...)
				} else {
					for _, entry := range entries {
						entriesDetails += fmt.Sprintf("  * %s\n", entry)

						durationInHours := entry.DurationInHours()
						if durationInHours > 0 {
							tagTotalDuration += entry.DurationInHours()
						}
					}
					dayTotalDuration += tagTotalDuration

					reportStr += fmt.Sprintf("\n- %s (%.2f)\n", tag, tagTotalDuration)
					reportStr += fmt.Sprint(entriesDetails)
				}

				entriesDetails = ""
				tagTotalDuration = 0.0
			}

			if len(other) > 0 {
				for _, entry := range other {
					entriesDetails += fmt.Sprintf("  * %s\n", entry)

					durationInHours := entry.DurationInHours()
					if durationInHours > 0 {
						tagTotalDuration += entry.DurationInHours()
					}
				}
				dayTotalDuration += tagTotalDuration

				reportStr += fmt.Sprintf("\n- [other] (%.2f)\n", tagTotalDuration)
				reportStr += fmt.Sprint(entriesDetails)
			}

			fmt.Printf("\n[[ %s ]] (%.2f)\n", day, dayTotalDuration)
			fmt.Print(reportStr)
		}

		return nil
	}),
}

func init() {
	listCmd.Flags().Int8P("week", "w", 0, "Week interval. Ex.: -1 (last week), -2 (two weeks ago)")
	reportCmd.Flags().Int8P("week", "w", 0, "Week interval. Ex.: -1 (last week), -2 (two weeks ago)")

	togglCmd.AddCommand(listCmd)
	togglCmd.AddCommand(reportCmd)
}
