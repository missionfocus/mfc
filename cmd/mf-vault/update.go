package mf_vault

import (
	"fmt"
	"git.missionfocus.com/open-source/mf-vault/pkg/autoupdate"
	"github.com/spf13/cobra"
	"net/url"
	"os"
	"path/filepath"
)

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.PersistentFlags().BoolVar(&updateForce, "force", false, "ignore checks for other package managers")
}

const projectURL = "https://git.missionfocus.com/api/v4/projects/394"

var updateForce bool

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update the mf-vault binary to the latest release",
	Run: func(cmd *cobra.Command, args []string) {
		if isInstalledFromBrew() && !updateForce {
			fmt.Print("It looks like mf-vault was installed using Homebrew.\nPlease use \"brew update && brew upgrade mf-vault\" or re-run this command with the --force flag.\n")
			os.Exit(1)
		}

		updater := &autoupdate.GitLabUpdater{ProjectURL: mustParseURL(projectURL),}

		fmt.Println("Checking for update...")
		nextVer, err := updater.Check(version)
		check(err)
		if nextVer == "" {
			fmt.Println("Up to date.")
			os.Exit(0)
		}

		fmt.Printf("Updating to v%s...", nextVer)
		check(updater.Update())
		fmt.Print(" done.\n")
	},
}

func isInstalledFromBrew() bool {
	exe, err := os.Executable()
	check(err)
	exe, err = filepath.EvalSymlinks(exe)
	check(err)
	match, err := filepath.Match("/usr/local/Cellar/*", exe)
	check(err)
	return match
}

func mustParseURL(str string) (u *url.URL) {
	u, err := url.Parse(str)
	check(err)
	return
}