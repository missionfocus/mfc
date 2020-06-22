package main

import (
	"git.missionfocus.com/ours/code/tools/mfc/pkg/tmetric"
	"git.missionfocus.com/ours/code/tools/mfc/pkg/vault"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	mfcCmd.AddCommand(tmetricCmd)
	tmetricCmd.AddCommand(tmetricSetTokenCmd)
	tmetricCmd.AddCommand(tmPerform)
	tmetricCmd.AddCommand(tmFind)
	tmetricCmd.AddCommand(tmScan)

	tmPerform.Flags().StringVarP(&tmetricFormat, "format", "f", "md", "output format to use for performance records")
	tmPerform.Flags().StringVarP(&tmetricStartDate, "start-date", "d", "", "start date from which to query time entries")
	tmPerform.Flags().StringVarP(&tmetricEndDate, "end-date", "e", "", "end date from which to query time entries")
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

/*
 * Subcommands
 */

var (
	tmetricFormat    string
	tmetricStartDate string
	tmetricEndDate   string
)

// Prints Individual's performance records.
var tmPerform = &cobra.Command{
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

// For pipeline to run each night to find errors in TMetric times.
var tmScan = &cobra.Command{
	Use:     "scan",
	Short:   "Nightly TMetric Scan",
	Aliases: []string{"scan"},
	Run: func(cmd *cobra.Command, args []string) {
		vaultAPIClient, err := getVaultClientWithToken()
		check(err)
		vaultClient := vault.New(vaultAPIClient)

		check(tmetric.ScanAll(vaultClient, os.Stdout))
	},
}
