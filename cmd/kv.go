package cmd

import (
	"fmt"
	"git.missionfocus.com/open-source/mf-vault/vault"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(kvCmd)
	kvCmd.AddCommand(kvListAllCmd)
}

var kvCmd = &cobra.Command{
	Use:   "kv",
	Short: "Manages Key/Value secrets.",
}

var kvListAllCmd = &cobra.Command{
	Use:   "listall <key>",
	Short: "Lists all keys under the specified K/V engine key. Key must end with `/`.",
	Example: "  listall secret/\tLists all keys under the default K/V secrets engine.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client, err := getClientWithToken()
		if err != nil {
			fatal(err)
		}
		v := vault.New(client)
		keys := v.ListAllKV(args[0])
		for _, key := range keys {
			fmt.Println(key)
		}
	},
}
