package mf_vault

import (
	"fmt"
	"git.missionfocus.com/open-source/mf-vault/pkg/vault"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
	"io/ioutil"
	"syscall"
)

func init() {
	rootCmd.AddCommand(loginCmd)
}

var loginCmd = &cobra.Command{
	Use:   "login [username] [password]",
	Short: "Log in to Vault",
	Args:  cobra.MaximumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		var (
			username string
			password string
		)

		switch len(args) {
		case 0:
			token, err := securePrompt("Enter your Vault token (will be hidden): ")
			check(err)
			check(writeToken(token))
			fmt.Printf("\nToken written to %s\n", tokenFilePath)
			return
		case 1:
			pw, err := securePrompt("Enter your password (will be hidden): ")
			check(err)
			password = pw
		case 2:
			password = args[1]
		}

		username = args[0]
		client, err := getClient()
		check(err)
		v := vault.New(client)
		token, err := v.AuthLDAP(username, password)
		check(err)
		check(writeToken(token))
		fmt.Printf("\nLogged in to Vault as %s.\n", username)
	},
}

func securePrompt(prompt string) (string, error) {
	fmt.Print(prompt)
	pw, err := terminal.ReadPassword(int(syscall.Stdin))
	return string(pw), err
}

func writeToken(token string) error {
	return ioutil.WriteFile(tokenFilePath, []byte(token), 0600)
}
