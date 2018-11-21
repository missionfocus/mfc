package cmd

import (
	"fmt"
	"github.com/hashicorp/vault/api"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"path/filepath"
)

var tokenFilePath string

func init() {
	rootCmd.PersistentFlags().StringVarP(&tokenFilePath, "token-file", "t", filepath.Join(os.Getenv("HOME"), ".vault-token"), "path to vault token file")
}

var rootCmd = &cobra.Command{
	Use:   "mf-vault",
	Version: "0.2.0",
	Short: "CLI for interacting with the Mission Focus Vault",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fatal(err)
	}
}

func fatal(err error) {
	fmt.Println(err)
	os.Exit(1)
}

func getClient() (*api.Client, error) {
	client, err := api.NewClient(&api.Config{
		Address: os.Getenv("VAULT_ADDR"),
	})
	if err != nil {
		return nil, err
	}

	return client, nil
}

func getClientWithToken() (*api.Client, error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	if client.Token() == "" {
		token, err := ioutil.ReadFile(tokenFilePath)
		if err != nil {
			return nil, err
		}

		client.SetToken(string(token))
	}

	return client, nil
}
