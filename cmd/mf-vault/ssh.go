package mf_vault

import (
	"git.missionfocus.com/open-source/mf-vault/pkg/vault"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func init() {
	rootCmd.AddCommand(sshCmd)
	sshCmd.AddCommand(sshSignCmd)
}

var sshCmd = &cobra.Command{
	Use:   "ssh",
	Short: "Performs operations related to SSH.",
}

var sshSignCmd = &cobra.Command{
	Use:   "sign",
	Short: "Sign client SSH key",
	Run: func(cmd *cobra.Command, args []string) {

		client, clientError := getClientWithToken()
		check(clientError)

		v := vault.New(client)

		keyPath := filepath.Join(os.Getenv("HOME"), ".ssh", "id_rsa.pub")
		keyBytes, keyReadError := ioutil.ReadFile(keyPath)
		check(keyReadError)

		key := string(keyBytes)
		keyTrimmed := strings.TrimRight(key, "\r\n")
		trimmedKeyBytes := []byte(keyTrimmed)
		secret, signError := v.SSHSignUserKey(trimmedKeyBytes)
		check(signError)

		data := secret.Data
		signedKey := data["signed_key"].(string)
		signedKeyBytes := []byte(signedKey)
		signedKeyPath := filepath.Join(os.Getenv("HOME"), ".ssh", "id_rsa-cert.pub")
		writeSignedKeyError := ioutil.WriteFile(signedKeyPath, signedKeyBytes, 0644)
		check(writeSignedKeyError)

		silentPrintf("Signed public key written to: %s\n", signedKeyPath)

	},
}
