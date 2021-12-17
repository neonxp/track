package cmd

import (
	"strconv"
	"strings"

	"github.com/spf13/afero"
	"github.com/spf13/cobra"

	"github.com/neonxp/track/internal/tracker"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start activity",
	Long:  `Start new timespan on activity`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.PrintErr("First argument must be activity id\n")
			return
		}
		id, err := strconv.Atoi(args[0])
		if err != nil {
			cmd.PrintErr("First argument must be activity id, got %s\n", args[0])
			return
		}
		comment := strings.Join(args[1:], " ")

		fs := afero.NewOsFs()
		tr, err := tracker.New(fs)
		if err != nil {
			cmd.PrintErr(err)
			return
		}
		if err := tr.Start(id, comment); err != nil {
			cmd.PrintErr(err)
			return
		}

		activity := tr.Activity(id)

		cmd.Printf("Started new span for activity \"%s\".\n", activity.Title)
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
