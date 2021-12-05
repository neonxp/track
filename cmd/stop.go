package cmd

import (
	"strconv"

	"github.com/spf13/cobra"

	"github.com/neonxp/track/internal/tracker"
)

// stopCmd represents the stop command
var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop activity",
	Long:  `Stop working on activity`,
	Run: func(cmd *cobra.Command, args []string) {
		activities := tr.List(false)
		if len(args) != 0 {
			id, err := strconv.Atoi(args[0])
			if err != nil {
				cmd.PrintErr("First argument must be activity id, got %s.\n", args[0])
				return
			}
			activities = []*tracker.Activity{tr.Activity(id)}
		}
		for _, activity := range activities {
			if err := tr.Stop(activity.ID); err != nil {
				cmd.PrintErr(err)
				return
			}
			cmd.Printf("Stopped activity \"%s\".\n", activity.Title)
		}
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
}
