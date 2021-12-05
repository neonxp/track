package cmd

import (
	"strings"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add activity",
	Long:  `Add new activity that we can track`,
	Run: func(cmd *cobra.Command, args []string) {
		titles := []string{}
		tags := []string{}
		contexts := []string{}
		for _, s := range args {
			if len(s) == 0 {
				continue
			}
			if s[0:1] == "#" {
				if len(s[1:]) >= 1 {
					tags = append(tags, s[1:])
					continue
				}
			}
			if s[0:1] == "@" {
				if len(s[1:]) >= 1 {
					contexts = append(contexts, s[1:])
					continue
				}
			}
			titles = append(titles, s)
		}
		title := strings.Join(titles, " ")
		activityID, err := tr.Add(title, tags, contexts)
		if err != nil {
			cmd.PrintErr(err)
			return
		}
		cmd.Printf("Activity #%d added! Now you can start it.\n", activityID)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
