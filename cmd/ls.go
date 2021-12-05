package cmd

import (
	"strings"

	"github.com/spf13/cobra"
)

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List activities",
	Long:  `List started (or all by -a flag) activities`,
	Run: func(cmd *cobra.Command, args []string) {
		all, err := cmd.Flags().GetBool("all")
		if err != nil {
			cmd.PrintErr(err)
			return
		}
		verbose, err := cmd.Flags().GetBool("verbose")
		if err != nil {
			cmd.PrintErr(err)
			return
		}
		activities := tr.List(all)
		if len(activities) == 0 {
			cmd.Printf("There is no activities.\n")
			return
		}
		if all {
			cmd.Printf("Activities:\n")
		} else {
			cmd.Printf("Started activities:\n")
		}
		for _, activity := range activities {
			cmd.Printf("%d. %s\n", activity.ID, activity.Title)
			if len(activity.Tags) > 0 {
				cmd.Printf("\tTags: %v\n", activity.Tags)
			}
			if len(activity.Context) > 0 {
				cmd.Printf("\tContexts: %v\n", activity.Context)
			}
			cmd.Printf("\t%d timespans\n", len(activity.Spans))
			for i, span := range activity.Spans {
				if !verbose && i < len(activity.Spans)-1 {
					continue
				}
				stop := "now"
				if span.Stop != nil {
					stop = span.Stop.Format("15:04:05 2.1.2006")
				}
				if strings.Trim(span.Comment, " ") != "" {
					cmd.Printf("\t%s — %s: \"%s\"\n", span.Start.Format("15:04:05 2.1.2006"), stop, span.Comment)
				} else {
					cmd.Printf("\t%s — %s\n", span.Start.Format("15:04:05 2.1.2006"), stop)
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)
	lsCmd.Flags().BoolP("all", "a", false, "List all activities. Only started by default")
	lsCmd.Flags().BoolP("verbose", "v", false, "List all timespans. Only last by default")
}
