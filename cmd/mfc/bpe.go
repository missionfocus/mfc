package main

import (
	"git.missionfocus.com/ours/code/tools/mfc/pkg/bpe"
	"git.missionfocus.com/ours/code/tools/mfc/pkg/vault"
	"github.com/spf13/cobra"
)

func init() {
	mfcCmd.AddCommand(bpeCmd)
	bpeCmd.AddCommand(bpeVelocityReportCmd)
}

var bpeCmd = &cobra.Command{
	Use:     "BusinessProcessEngineering",
	Short:   "bpe <cmd>",
	Aliases: []string{"bpe"},
}

const gitlabVelocityExample = `
    mfc gitlab check velocity  ""		# Generate documentation for mfc in markdown format
 	mfc gitlab check velocity "GDAC 18" # Gets the velocity for the GDAC 18 milestone`

//bpeVelocityReportCmd returns the weight a mfm completed during a milestone. TODO iterations once implemented
var bpeVelocityReportCmd = &cobra.Command{
	Use:     "velocity",
	Short:   "Check each MFM velocity for a given milestone and iteration",
	Args:    cobra.ExactArgs(1),
	Example: gitlabVelocityExample,
	Run: func(cmd *cobra.Command, args []string) {
		vClient, err := getVaultClientWithToken()
		check(err)
		v := vault.New(vClient)

		client, err := getGitLabClient(v)
		check(err)

		check(bpe.VelocityReport(client, args[0], ""))
	},
}
