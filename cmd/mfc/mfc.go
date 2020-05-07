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

var mfcCmd = &cobra.Command{
	Use: "mfc",
	Version: func() string {
		if version == "" {
			return "next"
		}
		return version
	}(),
	Short: "Mission Focus Control CLI",
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
	pw, err := terminal.ReadPassword(syscall.Stdin)
	return string(pw), err
}

func mustParseURL(str string) (u *url.URL) {
	u, err := url.Parse(str)
	check(err)
	return
}
