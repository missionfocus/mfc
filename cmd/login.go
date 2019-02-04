package cmd

import (
	"fmt"
	"git.missionfocus.com/open-source/mf-vault/vault"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
	"io/ioutil"
)

func init() {
	rootCmd.AddCommand(loginCmd)
}

var loginCmd = &cobra.Command{
	Use:   "login <username> [password]",
	Short: "Login to Vault via LDAP username/password",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		username := args[0]
		var password string

		if len(args) == 1 {
			fmt.Print("Password (will be hidden): ")
			input, err := terminal.ReadPassword(0)
			fmt.Print("\n")
			if err != nil {
				check(err)
			}
			password = string(input)
		} else {
			password = args[1]
		}

		client, err := getClient()
		if err != nil {
			check(err)
		}

		v := vault.New(client)
		token, err := v.AuthLDAP(username, password)
		if err != nil {
			check(err)
		}

		ioutil.WriteFile(tokenFilePath, []byte(token), 0600)
		fmt.Printf("Logged in to Vault as %s.\n", username)
	},
}
