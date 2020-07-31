package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"git.missionfocus.com/ours/code/tools/mfc/pkg/vault"
	"github.com/spf13/cobra"
)

func init() {
	vaultCmd.AddCommand(vaultCryptCmd)

	vaultCryptCmd.PersistentFlags().BoolVarP(&vaultCryptDecrypt, "decrypt", "d", false, "decrypt input ciphertext to plaintext")
	vaultCryptCmd.PersistentFlags().StringVarP(&vaultCryptRecipient, "recipient", "r", "", "recipient")
}

var (
	vaultCryptDecrypt   bool
	vaultCryptRecipient string
)

const vaultCryptExample = `
  echo "String to encrypt" | mfc vault crypt -r <recipient> # Encrypt string for recipient specified.
  echo "vault:v1:encryptedString" | mfc vault crypt -d      # Decrypt vault encrypted string intended for myself.`

var vaultCryptCmd = &cobra.Command{
	Use:     "crypt [file]",
	Short:   "Encrypt and decrypt secrets",
	Example: vaultCryptExample,
	Args:    cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		v, err := getVaultClientWithToken()
		check(err)
		crypt := vault.NewCryptClient(v)

		r := os.Stdin
		if len(args) > 0 {
			f, err := os.Open(args[0])
			check(err)
			defer f.Close()
			r = f
		}

		in, err := ioutil.ReadAll(r)
		check(err)
		txt := string(in)

		if vaultCryptDecrypt {
			pt, err := crypt.De(vaultCryptRecipient, txt)
			check(err)
			fmt.Print(pt)
		} else {
			ct, err := crypt.En(vaultCryptRecipient, txt)
			check(err)
			fmt.Print(ct)
		}
	},
}
