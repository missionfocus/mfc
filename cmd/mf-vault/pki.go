package mf_vault

import (
	"encoding/json"
	"git.missionfocus.com/open-source/mf-vault/pkg/vault"
	"github.com/hashicorp/vault/api"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	rootCmd.AddCommand(pkiCmd)
	pkiCmd.AddCommand(pkiCreateFilesCmd)

	wd, _ := os.Getwd()
	pkiCmd.PersistentFlags().StringVar(&pkiCreateFilesDir, "dir", wd, "directory to create files in")
}

var pkiCmd = &cobra.Command{
	Use:   "pki",
	Short: "Interact with Vault's PKI engine.",
}

var pkiCreateFilesDir string

var pkiCreateFilesCmd = &cobra.Command{
	Use:   "create-files [filename]",
	Short: "Parses a Vault PKI secret in JSON format and creates certificate files.",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		r := os.Stdin

		if len(args) == 1 {
			f, err := os.Open(args[0])
			check(err)
			defer f.Close()
			r = f
		}

		var secret api.Secret
		err := json.NewDecoder(r).Decode(&secret)
		check(err)

		client, err := getClientWithToken()
		check(err)
		v := vault.New(client)
		check(v.PkiCreateFiles(&secret, pkiCreateFilesDir))
	},
}
