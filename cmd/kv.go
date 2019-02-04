package cmd

import (
	"fmt"
	"git.missionfocus.com/open-source/mf-vault/vault"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(kvCmd)
	kvCmd.AddCommand(kvListAllCmd)
	kvCmd.AddCommand(kvAwsCmd)
	kvAwsCmd.PersistentFlags().StringVarP(&profileName, "profile", "p", "vault", "name of the profile")
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
			check(err)
		}
		v := vault.New(client)
		keys := v.KvListAll(args[0])
		for _, key := range keys {
			fmt.Println(key)
		}
	},
}

var kvAwsCmd = &cobra.Command{
	Use: "aws <path>",
	Short: "Reads the secret at path as AWS credentials.",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		key := args[0]

		client, err := getClientWithToken()
		check(err)

		v := vault.New(client)
		secret, err := v.KvReadAws(key)
		check(err)

		check(secret.ToProfile(credentialsPath, profileName))
		silentPrint(fmt.Sprintf("AWS profile `%s` updated with the credentials read from `%s`.\n", profileName, key))
	},
}
