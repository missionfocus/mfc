package mf_vault

import (
	"fmt"
	"github.com/hashicorp/vault/api"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
	"io/ioutil"
	"os"
	"path/filepath"
	"syscall"
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

	rootCmd.PersistentFlags().StringVar(&tokenFilePath, "token-file", defaultTokenPath, "path to vault token file")
	rootCmd.PersistentFlags().StringVar(&credentialsPath, "aws-creds-file", defaultCredentialsPath, "path to AWS credentials file")
	rootCmd.PersistentFlags().BoolVar(&silent, "silent", false, "suppress output to stdout")
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

var cachedHomeDir string

func homeDir() string {
	if cachedHomeDir == "" {
		home, err := homedir.Dir()
		check(err)
		cachedHomeDir = home
	}
	return cachedHomeDir
}

func securePrompt(prompt string) (string, error) {
	fmt.Print(prompt)
	pw, err := terminal.ReadPassword(int(syscall.Stdin))
	return string(pw), err
}

func writeToken(token string) error {
	return ioutil.WriteFile(tokenFilePath, []byte(token), 0600)
}
