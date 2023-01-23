package cmd

import (
	"fmt"
	"strconv"

	"github.com/limeiralucas/chrono-cli/internal/pkg/time"
	"github.com/spf13/cobra"
)

var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "A command that converts milliseconds to hour",
	Args:  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		milliseconds, err := strconv.ParseInt(args[0], 10, 0)

		if err != nil {
			panic(err)
		}

		hours, err := time.MilliToHour(int(milliseconds))

		if err != nil {
			panic(err)
		}

		fmt.Println(hours)
	},
}
