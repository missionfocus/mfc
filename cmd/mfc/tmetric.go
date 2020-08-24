package main

import (
	"os"

	"git.missionfocus.com/ours/code/tools/mfc/pkg/bpe"
	"git.missionfocus.com/ours/code/tools/mfc/pkg/vault"
	"github.com/spf13/cobra"
)

func init() {
	mfcCmd.AddCommand(tmetricCmd)
	tmetricCmd.AddCommand(tmetricSetTokenCmd)
	tmetricCmd.AddCommand(tmetricHoursCommand)
	tmetricCmd.AddCommand(tmetricValidateTMetricCommand)

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
	tmetricPerson    string
	gitlabIssueURL   string
)

var tmetricHoursCommand = &cobra.Command{
	Use:     "hours",
	Short:   "Summarize an Individual's hours",
	Aliases: []string{"hrs"},
	Run: func(cmd *cobra.Command, args []string) {
		vaultAPIClient, err := getVaultClientWithToken()
		check(err)
		vaultClient := vault.New(vaultAPIClient)

		check(bpe.GetPersonHoursSummary(vaultClient, os.Stdout, tmetricPerson))
	},
}

var tmetricValidateTMetricCommand = &cobra.Command{
	Use:     "validate",
	Short:   "Pipeline TMetric Time Validator",
	Aliases: []string{"v"},
	Run: func(cmd *cobra.Command, args []string) {
		vaultAPIClient, err := getVaultClientWithToken()
		check(err)
		vaultClient := vault.New(vaultAPIClient)

		check(bpe.ValidateTMetricTime(vaultClient, os.Stdout))
	},
}
