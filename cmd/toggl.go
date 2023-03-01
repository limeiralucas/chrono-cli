package cmd

import (
	"fmt"

	"github.com/limeiralucas/chrono-cli/pkg/app"
	"github.com/limeiralucas/chrono-cli/pkg/config"
	toggl_provider "github.com/limeiralucas/chrono-cli/pkg/provider/toggl"
	"github.com/spf13/cobra"
)

var togglCmd = &cobra.Command{
	Use:   "toggl",
	Short: "Access toggl time entries",
}

func init() {
	togglCmd.AddCommand(&cobra.Command{
		Use:   "list",
		Short: "List all Time Entries",
		Run: (func(cmd *cobra.Command, args []string) {
			config, err := config.ReadConfig("./config.json")
			if err != nil {
				panic(err)
			}

			provider := toggl_provider.NewTimeEntryProvider(config.Token)
			service := app.NewTimeEntryService(&provider)

			entries, err := service.List()
			if err != nil {
				panic(err)
			}

			for _, entry := range entries {
				fmt.Printf("%s: %s - %s\n", entry.Description, entry.StartDate, entry.EndDate)
			}
		}),
	})
}
