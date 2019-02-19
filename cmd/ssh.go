package cmd

import (
	"fmt"
	"git.missionfocus.com/open-source/mf-vault/vault"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func init() {
	rootCmd.AddCommand(sshCmd)
	sshCmd.PersistentFlags().StringVarP(&keyPath, "public-key", "a", filepath.Join(os.Getenv("HOME"), ".ssh", "id_rsa.pub"), "Path used to read SSH public key")
	sshCmd.PersistentFlags().StringVarP(&signedKeyPath, "signed-public-key", "b", filepath.Join(os.Getenv("HOME"), ".ssh", "signed-cert.pub"), "Path to write signed certificate")
}

var (
	keyPath       string
	signedKeyPath string
)

var sshCmd = &cobra.Command{
	Use:   "ssh",
	Short: "Sign client SSH key",
	Run: func(cmd *cobra.Command, args []string) {

		client, clientError := getClientWithToken()
		check(clientError)

		v := vault.New(client)

		sshKeyBytes, sshKeyReadError := ioutil.ReadFile(sshKeyPath)
		check(sshKeyReadError)

		sshKey := string(sshKeyBytes)
		sshKeyTrimmed := strings.TrimRight(sshKey, "\r\n")
		trimmedSSHKeyBytes := []byte(sshKeyTrimmed)
		secret, signError := v.SignUserKey(trimmedSSHKeyBytes)
		check(signError)

		data := secret.Data
		signedKey := data["signed_key"].(string)
		signedKeyBytes := []byte(signedKey)
		writeSignedKeyError := ioutil.WriteFile(signedKeyPath, signedKeyBytes, 0644)
		check(writeSignedKeyError)

		fmt.Printf("Signed public key written to: %s\n", signedKeyPath)

	},
}
