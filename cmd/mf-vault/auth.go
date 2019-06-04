package mf_vault

import (
	"git.missionfocus.com/open-source/mf-vault/pkg/vault"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	rootCmd.AddCommand(authCmd)
	authCmd.AddCommand(authApproleCmd)
}

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Authenticate with Vault.",
}

var authApproleCmd = &cobra.Command{
	Use:   "approle",
	Short: "Authenticate with Vault's AppRole engine.",
	Run: func(cmd *cobra.Command, args []string) {
		var (
			roleID   = os.Getenv("VAULT_ROLE_ID")
			secretID = os.Getenv("VAULT_SECRET_ID")
		)
		if roleID == "" || secretID == "" {
			check(errors.New("both VAULT_ROLE_ID and VAULT_SECRET_ID must be set to use AppRole authentication"))
		}

		client, err := getClient()
		check(err)
		v := vault.New(client)
		token, err := v.AuthApprole(roleID, secretID)
		check(err)
		check(writeToken(token))
	},
}
