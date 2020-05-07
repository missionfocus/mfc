package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"git.missionfocus.com/ours/code/tools/mfc/pkg/vault"
	"github.com/spf13/cobra"
)

func init() {
	cryptCmd.PersistentFlags().BoolVarP(&cryptDecrypt, "decrypt", "d", false, "decrypt input ciphertext to plaintext")
	cryptCmd.PersistentFlags().StringVarP(&cryptRecipient, "recipient", "r", "", "recipient")

	rootCmd.AddCommand(cryptCmd)
}

var (
	cryptDecrypt   bool
	cryptRecipient string
)

var cryptCmd = &cobra.Command{
	Use:   "crypt [file]",
	Short: "Encrypt and decrypt secrets",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		v, err := getClientWithToken()
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

		if cryptDecrypt {
			pt, err := crypt.De(cryptRecipient, txt)
			check(err)
			fmt.Print(pt)
		} else {
			ct, err := crypt.En(cryptRecipient, txt)
			check(err)
			fmt.Print(ct)
		}
	},
}
