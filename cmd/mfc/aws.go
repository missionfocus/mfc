package mfc

import (
	"fmt"
	"git.missionfocus.com/ours/code/tools/mfc/pkg/vault"
	"github.com/pkg/browser"
	"github.com/spf13/cobra"
	"time"
)

func init() {
	rootCmd.AddCommand(awsCmd)
	awsCmd.AddCommand(awsIssueCmd)
	awsCmd.AddCommand(awsListRolesCmd)

	awsIssueCmd.PersistentFlags().StringVarP(&awsIssueProfileName, "profile", "p", "", "name of the profile")
	awsIssueCmd.PersistentFlags().StringVarP(&awsIssueTTL, "ttl", "l", "3600s", "requested TTL of the STS token")
	awsIssueCmd.PersistentFlags().BoolVarP(&awsIssueAutoOpenURL, "open", "o", false, "automatically open the AWS console")
}

var (
	awsIssueTTL         string
	awsIssueProfileName string
	awsIssueAutoOpenURL bool
)

var awsCmd = &cobra.Command{
	Use:   "aws",
	Short: "Interact Vault's AWS secrets engine",
}

var awsIssueCmd = &cobra.Command{
	Use:   "issue <account> <role>",
	Short: "Issue AWS credentials for the specified account and role",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		account := args[0]
		role := args[1]
		if awsIssueProfileName == "" {
			awsIssueProfileName = account
		}

		client, err := getClientWithToken()
		check(err)
		v := vault.New(client)

		secret, err := v.AWSReadSTS(account, role, awsIssueTTL)
		check(err)
		stsSecret := vault.NewSTSSecret(secret)
		check(stsSecret.ToProfile(credentialsPath, awsIssueProfileName))

		if silent {
			return
		}

		fmt.Printf("AWS profile `%s` updated with credentials for IAM role `%s` of account `%s`.\n", awsIssueProfileName, role, account)
		fmt.Printf("These credentials are valid for: %s\n", (time.Second * time.Duration(secret.LeaseDuration)).String())

		loginURL, err := stsSecret.GenerateLoginUrl(account)
		check(err)
		fmt.Printf("Console login URL (valid for 15 minutes):\n\n%s\n", loginURL.String())
		if awsIssueAutoOpenURL {
			check(browser.OpenURL(loginURL.String()))
		}
	},
}

var awsListRolesCmd = &cobra.Command{
	Use:   "list-roles <account>",
	Short: "List available roles for the specified account",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client, err := getClientWithToken()
		check(err)
		v := vault.New(client)

		roles, err := v.AWSListRoles(args[0])
		check(err)
		for _, role := range roles {
			fmt.Println(role)
		}
	},
}
