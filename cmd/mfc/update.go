package main

import (
	"fmt"
	"net/http"
	"time"

	"git.missionfocus.com/ours/code/tools/mfc/pkg/autoupdate"
	"github.com/spf13/cobra"
)

func init() {
	mfcCmd.AddCommand(updateCmd)
}

const (
	manifestURL = "http://public.missionfocus.com/mfc/manifest.yaml"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update the mfc binary to the latest release",
	Run: func(cmd *cobra.Command, args []string) {
		updater := autoupdate.ManifestUpdater{
			ManifestURL: manifestURL,
			HTTPClient:  &http.Client{Timeout: 10 * time.Second},
		}

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
