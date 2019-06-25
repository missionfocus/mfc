package mf_vault

import (
	"fmt"
	"git.missionfocus.com/open-source/mf-vault/pkg/vault"
	"github.com/spf13/cobra"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func init() {
	rootCmd.AddCommand(sshCmd)
	sshCmd.AddCommand(sshSignUserCmd)
	sshCmd.AddCommand(sshSignHostCmd)
	sshCmd.AddCommand(sshCACmd)
}

const sshDefaultEngine = "ssh-signer"

var sshCmd = &cobra.Command{
	Use:   "ssh",
	Short: "Performs operations related to SSH.",
}

var sshSignUserCmd = &cobra.Command{
	Use:   "sign",
	Short: "Sign client SSH key",
	Run: func(cmd *cobra.Command, args []string) {
		keyPath := filepath.Join(homeDir, ".ssh", "id_rsa.pub")
		signedKeyPath := filepath.Join(homeDir, ".ssh", "id_rsa-cert.pub")

		signPubKey(keyPath, signedKeyPath, "user")

	},
}

var sshSignHostCmd = &cobra.Command{
	Use:   "sign-host",
	Short: "Sign host SSH key",
	Run: func(cmd *cobra.Command, args []string) {
		keyPath := filepath.Join("/etc/ssh/", "ssh_host_rsa_key.pub")
		signedKeyPath := filepath.Join("/etc/ssh/", "ssh_host_rsa_key-cert.pub")

		signPubKey(keyPath, signedKeyPath, "host")

	},
}

func signPubKey(keyPath string, signedKeyPath string, usage string) {
	client, clientError := getClientWithToken()
	check(clientError)

	v := vault.New(client)

	keyBytes, keyReadError := ioutil.ReadFile(keyPath)
	check(keyReadError)

	key := string(keyBytes)
	keyTrimmed := strings.TrimRight(key, "\r\n")
	trimmedKeyBytes := []byte(keyTrimmed)
	secret, signError := v.SSHSignPubKey(trimmedKeyBytes, usage)
	check(signError)

	data := secret.Data
	signedKey := data["signed_key"].(string)
	signedKeyBytes := []byte(signedKey)
	writeSignedKeyError := ioutil.WriteFile(signedKeyPath, signedKeyBytes, 0644)
	check(writeSignedKeyError)

	silentPrintf("Signed public key written to: %s\n", signedKeyPath)
}

var sshCACmd = &cobra.Command{
	Use: "ca",
	Short: "Print the public key of the Vault SSH Signer CA.",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := getClientWithToken()
		check(err)
		v := vault.New(client)

		pubKey, err := v.SSHCA(sshDefaultEngine)
		check(err)
		fmt.Println(pubKey)
	},
}