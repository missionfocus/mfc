package mf_vault

import (
	"bufio"
	"fmt"
	"git.missionfocus.com/open-source/mf-vault/pkg/vault"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	rootCmd.AddCommand(authCmd)
	authCmd.AddCommand(authApproleCmd)
	authCmd.AddCommand(authLDAPCmd)
	authCmd.AddCommand(authTokenCmd)
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

var authLDAPCmd = &cobra.Command{
	Use:   "ldap [username] [password]",
	Short: "Authenticate to Vault using LDAP credentials",
	Args:  cobra.MaximumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		switch len(args) {
		case 0:
			scanner := bufio.NewScanner(os.Stdin)
			fmt.Print("Enter your LDAP username: ")
			scanner.Scan()
			args = append(args, scanner.Text())
			fallthrough
		case 1:
			pw, err := securePrompt("Enter your password (will be hidden): ")
			check(err)
			args = append(args, pw)
		}

		client, err := getClient()
		check(err)
		v := vault.New(client)
		token, err := v.AuthLDAP(args[0], args[1])
		check(err)
		check(writeToken(token))
		fmt.Printf("\nLogged in to Vault as %s.\n", args[0])
	},
}

var authTokenCmd = &cobra.Command{
	Use:   "token <token>",
	Short: "Authenticate to Vault using a raw token",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		check(writeToken(args[0]))
	},
}
