package mf_vault

import (
	"fmt"
	"github.com/hashicorp/vault/api"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"path/filepath"
)

const defaultVaultAddr = "https://vault.missionfocus.com"

var (
	credentialsPath string
	tokenFilePath   string
	silent          bool
)

func init() {
	defaultCredentialsPath := filepath.Join(homeDir(), ".aws", "credentials")
	defaultTokenPath := filepath.Join(homeDir(), ".vault-token")

	rootCmd.PersistentFlags().StringVarP(&tokenFilePath, "token-file", "t", defaultTokenPath, "path to vault token file")
	rootCmd.PersistentFlags().StringVarP(&credentialsPath, "credentials", "c", defaultCredentialsPath, "path to AWS credentials file")
	rootCmd.PersistentFlags().BoolVarP(&silent, "silent", "s", false, "suppress output to stdout")
}

// Do not modify this variable, it will be set at build time.
var version string

var rootCmd = &cobra.Command{
	Use: "mf-vault",
	Version: func() string {
		if version == "" {
			return "next"
		}
		return version
	}(),
	Short: "CLI for interacting with the Mission Focus Vault",
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
	addr := os.Getenv("VAULT_ADDR")
	if addr == "" {
		addr = defaultVaultAddr
	}
	client, err := api.NewClient(&api.Config{
		Address: addr,
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

func homeDir() string {
	home, err := homedir.Dir()
	check(err)
	return home
}
