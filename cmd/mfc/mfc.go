package main

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"syscall"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
)

func init() {
	defaultCredentialsPath := filepath.Join(homeDir(), ".aws", "credentials")
	defaultTokenPath := filepath.Join(homeDir(), ".vault-token")

	mfcCmd.PersistentFlags().StringVar(&mfcVaultTokenFile, "token-file", defaultTokenPath, "path to vault token file")
	mfcCmd.PersistentFlags().StringVar(&mfcAWSCredentialsPath, "aws-creds-file", defaultCredentialsPath, "path to AWS credentials file")
	mfcCmd.PersistentFlags().BoolVar(&mfcSilent, "silent", false, "suppress output to stdout")
}

// Globals
var (
	// Do not modify this variable, it will be set at build time.
	version string
)

var (
	mfcAWSCredentialsPath string
	mfcVaultTokenFile     string
	mfcSilent             bool
)

const mfcExample = `
  mfc config completion -h       # Set up mfc auto completion for your shell
  mfc vault auth ldap <username> # Renew vault-token through LDAP credentials
  mfc vault aws -h               # See options for AWS 
  mfc vault macro minio          # Set up minio authentication
  mfc vault crypt -h             # See usage for encrypting, and decrypting messages`

var mfcCmd = &cobra.Command{
	Use: "mfc",
	Version: func() string {
		if version == "" {
			return "next"
		}
		return version
	}(),
	Short:   "Mission Focus Control CLI",
	Example: mfcExample,
}

func main() {
	if err := mfcCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

/* Common utility functions */

func check(err error) {
	if err != nil {
		fmt.Println("Fatal: " + err.Error())
		os.Exit(1)
	}
}

func silentPrint(str string) {
	if !mfcSilent {
		fmt.Print(str)
	}
}

func silentPrintf(format string, a ...interface{}) {
	if !mfcSilent {
		fmt.Printf(format, a...)
	}
}

func homeDir() string {
	home, err := homedir.Dir()
	check(err)
	return home
}

func securePrompt(prompt string) (string, error) {
	fmt.Print(prompt)
	pw, err := terminal.ReadPassword(int(syscall.Stdin)) // Do not remove this int cast, it's required for Windows
	return string(pw), err
}

func mustParseURL(str string) (u *url.URL) {
	u, err := url.Parse(str)
	check(err)
	return
}
