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
	sshSignCmd.PersistentFlags().StringVarP(&sshSignKeyPath, "public-key", "a", filepath.Join(os.Getenv("HOME"), ".ssh", "id_rsa.pub"), "Path used to read SSH public key")
	sshSignCmd.PersistentFlags().StringVarP(&sshSignSignedKeyPath, "signed-public-key", "b", filepath.Join(os.Getenv("HOME"), ".ssh", "signed-cert.pub"), "Path to write signed certificate")
}

var sshCmd = &cobra.Command{
	Use:   "ssh",
	Short: "Performs operations related to SSH.",
}

var (
	sshSignKeyPath       string
	sshSignSignedKeyPath string
)

var sshSignCmd = &cobra.Command{
	Use:   "sign",
	Short: "Sign client SSH key",
	Run: func(cmd *cobra.Command, args []string) {

		client, clientError := getClientWithToken()
		check(clientError)

		v := vault.New(client)

		keyBytes, keyReadError := ioutil.ReadFile(sshSignKeyPath)
		check(keyReadError)

		key := string(keyBytes)
		keyTrimmed := strings.TrimRight(key, "\r\n")
		trimmedKeyBytes := []byte(keyTrimmed)
		secret, signError := v.SSHSignUserKey(trimmedKeyBytes)
		check(signError)

		data := secret.Data
		signedKey := data["signed_key"].(string)
		signedKeyBytes := []byte(signedKey)
		writeSignedKeyError := ioutil.WriteFile(sshSignSignedKeyPath, signedKeyBytes, 0644)
		check(writeSignedKeyError)

		silentPrintf("Signed public key written to: %s\n", sshSignSignedKeyPath)

	},
}
