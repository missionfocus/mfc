package cmd

import (
	"fmt"
	"git.missionfocus.com/open-source/mf-vault/vault"
	"github.com/spf13/cobra"
	"time"
)

func init() {
	rootCmd.AddCommand(awsCmd)
	awsCmd.PersistentFlags().StringVarP(&profileName, "profile", "p", "vault", "name of the profile")
}

var awsCmd = &cobra.Command{
	Use:   "aws <account> <role>",
	Short: "Manages AWS Credentials",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		account := args[0]
		role := args[1]

		client, err := getClientWithToken()
		if err != nil {
			check(err)
		}
		v := vault.New(client)

		secret, err := v.ReadSTS(account, role)
		if err != nil {
			check(err)
		}
		stsSecret := vault.NewSTSSecret(secret)
		if err := stsSecret.ToProfile(credentialsPath, profileName); err != nil {
			check(err)
		}

		if silent {
			return
		}

		fmt.Printf("AWS profile `%s` updated with credentials for IAM role `%s` of account `%s`.\n", profileName, role, account)
		fmt.Printf("These credentials are valid for: %s\n", (time.Second * time.Duration(secret.LeaseDuration)).String())

		loginUrl, err := stsSecret.GenerateLoginUrl(account)
		if err != nil {
			check(err)
		}

		fmt.Printf("Console login URL (valid for 15 minutes):\n\n%s\n", loginUrl.String())
	},
}
