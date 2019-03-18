package mf_vault

import (
	"fmt"
	"git.missionfocus.com/open-source/mf-vault/pkg/vault"
	"github.com/spf13/cobra"
	"time"
)

func init() {
	rootCmd.AddCommand(awsCmd)
	awsCmd.PersistentFlags().StringVarP(&awsProfileName, "profile", "p", "", "name of the profile")
	awsCmd.PersistentFlags().StringVarP(&awsTtl, "ttl", "l", "3600s", "requested TTL of the STS token")
}

var (
	awsTtl         string
	awsProfileName string
)

var awsCmd = &cobra.Command{
	Use:   "aws <account> <role>",
	Short: "Manages AWS Credentials",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		account := args[0]
		role := args[1]
		if awsProfileName == "" {
			awsProfileName = account
		}

		client, err := getClientWithToken()
		check(err)
		v := vault.New(client)

		secret, err := v.AwsReadSTS(account, role, awsTtl)
		check(err)
		stsSecret := vault.NewSTSSecret(secret)
		check(stsSecret.ToProfile(credentialsPath, awsProfileName))

		if silent {
			return
		}

		fmt.Printf("AWS profile `%s` updated with credentials for IAM role `%s` of account `%s`.\n", awsProfileName, role, account)
		fmt.Printf("These credentials are valid for: %s\n", (time.Second * time.Duration(secret.LeaseDuration)).String())

		loginUrl, err := stsSecret.GenerateLoginUrl(account)
		check(err)
		fmt.Printf("Console login URL (valid for 15 minutes):\n\n%s\n", loginUrl.String())
	},
}
