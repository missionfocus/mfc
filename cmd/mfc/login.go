package main

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(loginCmd)
}

var loginCmd = &cobra.Command{
	Use:        "login [username] [password]",
	Short:      "Log in to Vault",
	Deprecated: `use "auth ldap" or "auth token"`,
	Args:       cobra.MaximumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		authLDAPCmd.Run(cmd, args)
	},
}
