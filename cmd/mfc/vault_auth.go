package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"git.missionfocus.com/ours/code/tools/mfc/pkg/vault"
	"github.com/spf13/cobra"
)

func init() {
	vaultCmd.AddCommand(vaultAuthCmd)
	vaultAuthCmd.AddCommand(vaultAuthApproleCmd)
	vaultAuthCmd.AddCommand(vaultAuthLDAPCmd)
	vaultAuthCmd.AddCommand(vaultAuthTokenCmd)
	vaultAuthCmd.AddCommand(vaultAuthRADIUSCmd)
}

var vaultAuthCmd = &cobra.Command{
	Use:     "auth",
	Short:   "Authenticate with Vault",
	Example: vaultAuthExample,
}

var AppRoleCredentialsError = errors.New("both VAULT_ROLE_ID and VAULT_SECRET_ID must be set or passed as arguments to use AppRole authentication")

var vaultAuthApproleCmd = &cobra.Command{
	Use:   "approle [role id] [secret id]",
	Short: "Authenticate with Vault using AppRole RoleID/SecretID",
	Args:  cobra.MaximumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		var (
			roleID   = os.Getenv("VAULT_ROLE_ID")
			secretID = os.Getenv("VAULT_SECRET_ID")
		)

		if roleID == "" {
			if len(args) < 1 {
				check(AppRoleCredentialsError)
			}

			roleID = args[0]
		}

		if secretID == "" {
			if len(args) < 2 {
				check(AppRoleCredentialsError)
			}

			secretID = args[1]
		}

		client, err := getVaultClient()
		check(err)
		v := vault.New(client)
		token, err := v.AuthApprole(roleID, secretID)
		check(err)
		check(writeVaultToken(token))
	},
}

const vaultAuthExample = `
  mfc vault auth ldap                       # Authenticate with ldap with prompts for username & password
  mfc vault auth ldap <username>            # Authenticate with ldap by passing in username`

var vaultAuthLDAPCmd = &cobra.Command{
	Use:     "ldap [username] [password]",
	Short:   "Authenticate to Vault using LDAP credentials",
	Example: vaultAuthExample,
	Args:    cobra.MaximumNArgs(2),
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

		client, err := getVaultClient()
		check(err)
		v := vault.New(client)
		token, err := v.AuthLDAP(args[0], args[1])
		check(err)
		check(writeVaultToken(token))
		fmt.Printf("\nLogged in to Vault as %s.\n", args[0])
	},
}

var vaultAuthTokenCmd = &cobra.Command{
	Use:   "token [token]",
	Short: "Authenticate to Vault using a raw token",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var token string
		if len(args) == 0 {
			t, err := securePrompt("Enter your Vault token (will be hidden): ")
			check(err)
			token = t
		} else {
			token = args[0]
		}
		check(writeVaultToken(token))
	},
}

var vaultAuthRADIUSCmd = &cobra.Command{
	Use:   "radius [username] [password] [mfa token]",
	Short: "Authenticate to Vault using RADIUS",
	Args:  cobra.MaximumNArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		switch len(args) {
		case 0:
			scanner := bufio.NewScanner(os.Stdin)
			fmt.Print("Enter your RADIUS username: ")
			scanner.Scan()
			args = append(args, scanner.Text())
			fallthrough
		case 1:
			pw, err := securePrompt("Enter your password (will be hidden): ")
			check(err)
			args = append(args, pw)
			fmt.Print("\n")
			fallthrough
		case 2:
			tok, err := securePrompt("Enter your MFA token (will be hidden): ")
			check(err)
			args = append(args, tok)
		}

		client, err := getVaultClient()
		check(err)
		v := vault.New(client)
		token, err := v.AuthRADIUS(args[0], args[1], args[2])
		check(err)
		check(writeVaultToken(token))
		fmt.Printf("\nLogged in to Vault as %s.\n", args[0])
	},
}
