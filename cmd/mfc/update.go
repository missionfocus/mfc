package main

import (
	"fmt"

	"git.missionfocus.com/ours/code/tools/mfc/pkg/autoupdate"
	"github.com/spf13/cobra"
)

func init() {
	mfcCmd.AddCommand(updateCmd)
}

const projectURL = "https://git.missionfocus.com/api/v4/projects/394"

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update the mf-vault binary to the latest release",
	Run: func(cmd *cobra.Command, args []string) {
		updater := &autoupdate.GitLabUpdater{ProjectURL: mustParseURL(projectURL)}

		fmt.Println("Checking for update...")
		nextVer, err := updater.Check(version)
		check(err)
		if nextVer == "" {
			fmt.Println("Up to date.")
			return
		}

		fmt.Printf("Updating to v%s...", nextVer)
		check(updater.Update())
		fmt.Print(" done.\n")
	},
}
