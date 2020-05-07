package main

import (
	"fmt"
	"git.missionfocus.com/ours/code/tools/mfc/pkg/vault"
	"github.com/spf13/cobra"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func init() {
	vaultCmd.AddCommand(vaultSSHCmd)
	vaultSSHCmd.AddCommand(vaultSSHSignUserCmd)
	vaultSSHCmd.AddCommand(vaultSSHSignHostCmd)
	vaultSSHCmd.AddCommand(vaultSSHCACmd)
}

var vaultSSHCmd = &cobra.Command{
	Use:   "ssh",
	Short: "Interact with Vault's SSH engine",
}

var vaultSSHSignUserCmd = &cobra.Command{
	Use:   "sign",
	Short: "Sign client SSH key",
	Run: func(cmd *cobra.Command, args []string) {
		keyPath := filepath.Join(homeDir(), ".ssh", "id_rsa.pub")
		signedKeyPath := filepath.Join(homeDir(), ".ssh", "id_rsa-cert.pub")

		vaultSSHSignPubKey(keyPath, signedKeyPath, "user")
	},
}

var vaultSSHSignHostCmd = &cobra.Command{
	Use:   "sign-host",
	Short: "Sign host SSH key",
	Run: func(cmd *cobra.Command, args []string) {
		keyPath := filepath.Join("/etc/ssh/", "ssh_host_rsa_key.pub")
		signedKeyPath := filepath.Join("/etc/ssh/", "ssh_host_rsa_key-cert.pub")

		vaultSSHSignPubKey(keyPath, signedKeyPath, "host")
	},
}

func vaultSSHSignPubKey(keyPath string, signedKeyPath string, usage string) {
	client, clientError := getVaultClientWithToken()
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

var vaultSSHCACmd = &cobra.Command{
	Use:   "ca",
	Short: "Print the public key of the Vault SSH Signer CA",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := getVaultClientWithToken()
		check(err)
		v := vault.New(client)

		pubKey, err := v.SSHCA(vault.SSHDefaultEngine)
		check(err)
		fmt.Println(pubKey)
	},
}
