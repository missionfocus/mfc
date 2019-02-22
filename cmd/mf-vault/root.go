package mf_vault

import (
	"fmt"
	"github.com/hashicorp/vault/api"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"path/filepath"
)

var (
	credentialsPath string
	profileName     string
	tokenFilePath   string
	silent          bool
)

func init() {
	defaultCredentialsPath := filepath.Join(os.Getenv("HOME"), ".aws", "credentials")
	rootCmd.PersistentFlags().StringVarP(&tokenFilePath, "token-file", "t", filepath.Join(os.Getenv("HOME"), ".vault-token"), "path to vault token file")
	rootCmd.PersistentFlags().StringVarP(&credentialsPath, "credentials", "c", defaultCredentialsPath, "path to AWS credentials file")
	rootCmd.PersistentFlags().BoolVarP(&silent, "silent", "s", false, "suppress output to stdout")
}

// Do not modify this variable, it will be set at build time.
var version string

var rootCmd = &cobra.Command{
	Use:     "mf-vault",
	Version: func() string {
		if version == "" {
			return "next"
		}
		return version
	}(),
	Short:   "CLI for interacting with the Mission Focus Vault",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func check(err error) {
	if err != nil {
		fmt.Println("Fatal: " + err.Error())
		os.Exit(1)
	}
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

func silentPrint(str string) {
	if !silent {
		fmt.Print(str)
	}
}

func silentPrintf(format string, a ...interface{}) {
	if !silent {
		fmt.Printf(format, a...)
	}
}