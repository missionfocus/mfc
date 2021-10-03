package main

import (
	"io/ioutil"
	"os"

	"github.com/missionfocus/mfc/pkg/vault"
	"github.com/hashicorp/vault/api"
	"github.com/spf13/cobra"
)

func init() {
	mfcCmd.AddCommand(vaultCmd)
}

var vaultCmd = &cobra.Command{
	Use:   "vault",
	Short: "Interact with Vault",
}

/** Vault utility functions */

func getVaultClient() (*api.Client, error) {
	addr := os.Getenv("VAULT_ADDR")
	if addr == "" {
		addr = vault.DefaultAddr
	}
	client, err := api.NewClient(&api.Config{
		Address: addr,
	})
	if err != nil {
		return nil, err
	}
	return client, nil
}

func getVaultClientWithToken() (*api.Client, error) {
	client, err := getVaultClient()
	if err != nil {
		return nil, err
	}

	if client.Token() == "" {
		token, err := ioutil.ReadFile(mfcVaultTokenFile)
		if err != nil {
			return nil, err
		}

		client.SetToken(string(token))
	}

	return client, nil
}

func writeVaultToken(token string) error {
	return ioutil.WriteFile(mfcVaultTokenFile, []byte(token), 0600)
}
