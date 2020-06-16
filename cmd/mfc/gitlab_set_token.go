package main

import (
	"git.missionfocus.com/ours/code/tools/mfc/pkg/vault"
	"github.com/spf13/cobra"
)

func init() {
	gitlabCmd.AddCommand(gitLabSetTokenCmd)
}

var gitLabSetTokenCmd = &cobra.Command{
	Use:   "set-token <token>",
	Short: "Sets the GitLab Personal Access Token that will be used to authenticate with GitLab",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client, err := getVaultClientWithToken()
		check(err)
		v := vault.New(client)

		_, err = v.KVUserWrite("gitlab", map[string]interface{}{
			"data": map[string]interface{}{
				"token": args[0],
			},
		})
		check(err)
	},
}
