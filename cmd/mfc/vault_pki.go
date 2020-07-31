package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"git.missionfocus.com/ours/code/tools/mfc/pkg/vault"
	"github.com/hashicorp/vault/api"
	"github.com/spf13/cobra"
)

func init() {
	vaultCmd.AddCommand(vaultPKICmd)
	vaultPKICmd.AddCommand(vaultPKICreateFilesCmd)
	vaultPKICmd.AddCommand(vaultPKIIssueCmd)
	vaultPKICmd.AddCommand(vaultPKICACmd)

	wd, _ := os.Getwd()

	vaultPKICmd.PersistentFlags().StringVar(&vaultPKICreateFilesDir, "dir", wd, "directory to create files in")
	vaultPKIIssueCmd.PersistentFlags().StringVar(&vaultPKIIssueTTL, "ttl", "", "ttl of the issued cert")
	vaultPKIIssueCmd.PersistentFlags().StringVar(&vaultPKIIssueFormat, "format", "pem", "format of the returned data, one of: pem, der, pem_bundle")
	vaultPKIIssueCmd.PersistentFlags().StringVarP(&vaultPKIIssueWrite, "write", "w", "", "location to write files")
}

var vaultPKICmd = &cobra.Command{
	Use:   "pki",
	Short: "Interact with Vault's PKI engine",
}

var vaultPKICreateFilesDir string

const vaultPKICreateFileExample = `
  mfc vault pki create-files pki.json # Reads in a json formatted file`

var vaultPKICreateFilesCmd = &cobra.Command{
	Use:     "create-files [filename]",
	Short:   "Parse a Vault PKI secret in JSON format and create certificate files",
	Example: vaultPKICreateFileExample,
	Args:    cobra.MaximumNArgs(1),
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

		client, err := getVaultClientWithToken()
		check(err)
		v := vault.New(client)
		check(v.PKICreateFiles(&secret, vaultPKICreateFilesDir))
	},
}

var (
	vaultPKIIssueTTL    string
	vaultPKIIssueFormat string
	vaultPKIIssueWrite  string
)

const vaultPKIIssueExample = `
  mfc vault pki issue <example>.missionfocus.com --ttl 31556952 -w ./ # Generate pki certificate for example host with a ttl of one year to current dir`

var vaultPKIIssueCmd = &cobra.Command{
	Use:     "issue <common name>",
	Short:   "Issue a new certificate signed by the Vault CA for the specified CN",
	Example: vaultPKIIssueExample,
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client, err := getVaultClientWithToken()
		check(err)
		v := vault.New(client)
		secret, err := v.PKIIssue(&vault.PKIIssueOptions{
			RoleName:   vault.DefaultPKIEngineRole,
			CommonName: args[0],
			TTL:        vaultPKIIssueTTL,
			Format:     vaultPKIIssueFormat,
		})
		check(err)

		if vaultPKIIssueWrite == "" {
			check(secret.WriteJSON(os.Stdout))
			return
		}

		chain, err := os.OpenFile(filepath.Join(vaultPKIIssueWrite, "certificate.pem"), os.O_CREATE|os.O_WRONLY, 0600)
		check(err)
		defer chain.Close()
		check(secret.WriteCertificate(chain))

		priv, err := os.OpenFile(filepath.Join(vaultPKIIssueWrite, "privkey.pem"), os.O_CREATE|os.O_WRONLY, 0600)
		check(err)
		defer priv.Close()
		check(secret.WritePrivateKey(priv))
	},
}

const vaultPKICAExample = `
  mfc vault pki ca --dir ./ # Get the vault CA certificate`

var vaultPKICACmd = &cobra.Command{
	Use:     "ca",
	Short:   "Get the CA certificate of the Vault CA",
	Example: vaultPKICAExample,
	Run: func(cmd *cobra.Command, args []string) {
		client, err := getVaultClientWithToken()
		check(err)
		v := vault.New(client)

		cert, err := v.PKIGetCACert(vault.DefaultPKIEnginePath)
		check(err)
		fmt.Print(cert)
	},
}
