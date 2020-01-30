package mf_vault

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"git.missionfocus.com/open-source/mf-vault/pkg/vault"
	"github.com/hashicorp/vault/api"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(pkiCmd)
	pkiCmd.AddCommand(pkiCreateFilesCmd)
	pkiCmd.AddCommand(pkiIssueCmd)
	pkiCmd.AddCommand(pkiCACmd)

	wd, _ := os.Getwd()
	pkiCmd.PersistentFlags().StringVar(&pkiCreateFilesDir, "dir", wd, "directory to create files in")

	pkiIssueCmd.PersistentFlags().StringVar(&pkiIssueTTL, "ttl", "", "ttl of the issued cert")
	pkiIssueCmd.PersistentFlags().StringVar(&pkiIssueFormat, "format", "pem", "format of the returned data, one of: pem, der, pem_bundle")
	pkiIssueCmd.PersistentFlags().StringVarP(&pkiIssueWrite, "write", "w", "", "location to write files")
}

var pkiCmd = &cobra.Command{
	Use:   "pki",
	Short: "Interact with Vault's PKI engine",
}

var pkiCreateFilesDir string

var pkiCreateFilesCmd = &cobra.Command{
	Use:   "create-files [filename]",
	Short: "Parse a Vault PKI secret in JSON format and create certificate files",
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
		check(v.PKICreateFiles(&secret, pkiCreateFilesDir))
	},
}

var (
	pkiIssueTTL    string
	pkiIssueFormat string
	pkiIssueWrite  string
)

var pkiIssueCmd = &cobra.Command{
	Use:   "issue <common name>",
	Short: "Issue a new certificate signed by the Vault CA for the specified CN",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client, err := getClientWithToken()
		check(err)
		v := vault.New(client)
		secret, err := v.PKIIssue(&vault.PKIIssueOptions{
			RoleName:   vault.DefaultPKIEngineRole,
			CommonName: args[0],
			TTL:        pkiIssueTTL,
			Format:     pkiIssueFormat,
		})
		check(err)

		if pkiIssueWrite == "" {
			check(secret.WriteJSON(os.Stdout))
			return
		}

		chain, err := os.OpenFile(filepath.Join(pkiIssueWrite, "certificate.pem"), os.O_CREATE|os.O_WRONLY, 0600)
		check(err)
		defer chain.Close()
		check(secret.WriteCertificate(chain))

		priv, err := os.OpenFile(filepath.Join(pkiIssueWrite, "privkey.pem"), os.O_CREATE|os.O_WRONLY, 0600)
		check(err)
		defer priv.Close()
		check(secret.WritePrivateKey(priv))
	},
}

var pkiCACmd = &cobra.Command{
	Use: "ca",
	Short: "Get the CA certificate of the Vault CA",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := getClientWithToken()
		check(err)
		v := vault.New(client)

		cert, err := v.PKIGetCACert(vault.DefaultPKIEnginePath)
		check (err)
		fmt.Print(cert)
	},
}
