package main

import (
	"os"

	"git.missionfocus.com/ours/code/tools/mfc/pkg/tmetric"
	"git.missionfocus.com/ours/code/tools/mfc/pkg/vault"
	"github.com/spf13/cobra"
)

func init() {
	mfcCmd.AddCommand(tmetricCmd)
	tmetricCmd.AddCommand(tmetricSetTokenCmd)
	tmetricCmd.AddCommand(tmetricPerformanceCmd)
	tmetricCmd.AddCommand(tmetricHoursCommand)
	tmetricCmd.AddCommand(tmetricScannerCommand)

	tmetricPerformanceCmd.Flags().StringVarP(&tmetricFormat, "format", "f", "md", "output format to use for performance records")
	tmetricPerformanceCmd.Flags().StringVarP(&tmetricStartDate, "start-date", "d", "", "start date from which to query time entries")
	tmetricPerformanceCmd.Flags().StringVarP(&tmetricEndDate, "end-date", "e", "", "end date from which to query time entries")
	tmetricHoursCommand.Flags().StringVarP(&tmetricPerson, "person", "p", "", "Insert MFM to search")
}

var tmetricCmd = &cobra.Command{
	Use:     "tmetric",
	Short:   "Interact with TMetric",
	Aliases: []string{"tm"},
}

var tmetricSetTokenCmd = &cobra.Command{
	Use:   "set-token <token>",
	Short: "Sets the TMetric API token that will be used to authenticate with TMetric",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client, err := getVaultClientWithToken()
		check(err)
		v := vault.New(client)

		_, err = v.KVUserWrite("tmetric", map[string]interface{}{
			"data": map[string]interface{}{
				"token": args[0],
			},
		})
		check(err)
	},
}

var (
	tmetricFormat    string
	tmetricStartDate string
	tmetricEndDate   string
)

// tmetricPerformanceCmd prints MFM's performance records.
var tmetricPerformanceCmd = &cobra.Command{
	Use:     "performance",
	Short:   "Overview of all each individual's performance",
	Aliases: []string{"perf"},
	Run: func(cmd *cobra.Command, args []string) {
		vaultAPIClient, err := getVaultClientWithToken()
		check(err)
		vaultClient := vault.New(vaultAPIClient)

		glClient, err := getGitLabClient(vaultClient)
		check(err)

		check(tmetric.GetReports(glClient, vaultClient, os.Stdout, tmetricStartDate, tmetricEndDate, tmetricFormat))
	},
}

var (
	tmetricPerson  string
	gitlabIssueURL string
)

var tmetricHoursCommand = &cobra.Command{
	Use:     "hours",
	Short:   "Summarize an Individual's hours",
	Aliases: []string{"hrs"},
	Run: func(cmd *cobra.Command, args []string) {
		vaultAPIClient, err := getVaultClientWithToken()
		check(err)
		vaultClient := vault.New(vaultAPIClient)

		check(tmetric.GetPersonHoursSummary(vaultClient, os.Stdout, tmetricPerson))
	},
}

var tmetricScannerCommand = &cobra.Command{
	Use:     "scanner",
	Short:   "Pipeline TMetric Scan",
	Aliases: []string{"scan"},
	Run: func(cmd *cobra.Command, args []string) {
		vaultAPIClient, err := getVaultClientWithToken()
		check(err)
		vaultClient := vault.New(vaultAPIClient)

		check(tmetric.Scanner(vaultClient, os.Stdout))
	},
}
