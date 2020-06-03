package main

import (
	"fmt"
	"git.missionfocus.com/ours/code/tools/mfc/pkg/vault"
	"github.com/pkg/browser"
	"github.com/spf13/cobra"
	"time"
)

func init() {
	vaultCmd.AddCommand(vaultAWSCmd)
	vaultAWSCmd.AddCommand(vaultAWSIssueCmd)
	vaultAWSCmd.AddCommand(vaultAWSListRolesCmd)

	vaultAWSIssueCmd.PersistentFlags().StringVarP(&vaultAWSIssueProfileName, "profile", "p", "", "name of the profile")
	vaultAWSIssueCmd.PersistentFlags().StringVarP(&vaultAWSIssueTTL, "ttl", "l", "3600s", "requested TTL of the STS token")
	vaultAWSIssueCmd.PersistentFlags().BoolVarP(&vaultAWSIssueAutoOpenURL, "open", "o", false, "automatically open the AWS console")
}

var (
	vaultAWSIssueTTL         string
	vaultAWSIssueProfileName string
	vaultAWSIssueAutoOpenURL bool
)

var vaultAWSCmd = &cobra.Command{
	Use:   "aws",
	Short: "Interact Vault's AWS secrets engine",
}

var vaultAWSIssueCmd = &cobra.Command{
	Use:   "issue <account> <role>",
	Short: "Issue AWS credentials for the specified account and role",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		account := args[0]
		role := args[1]
		if vaultAWSIssueProfileName == "" {
			vaultAWSIssueProfileName = account
		}

		client, err := getVaultClientWithToken()
		check(err)
		v := vault.New(client)

		secret, err := v.AWSReadSTS(account, role, vaultAWSIssueTTL)
		check(err)
		stsSecret := vault.NewSTSSecret(secret)
		check(stsSecret.ToProfile(mfcAWSCredentialsPath, vaultAWSIssueProfileName))

		if mfcSilent {
			return
		}

		fmt.Printf("AWS profile `%s` updated with credentials for IAM role `%s` of account `%s`.\n", vaultAWSIssueProfileName, role, account)
		fmt.Printf("These credentials are valid for: %s\n", (time.Second * time.Duration(secret.LeaseDuration)).String())

		loginURL, err := stsSecret.GenerateLoginUrl(account)
		check(err)
		fmt.Printf("Console login URL (valid for 15 minutes):\n\n%s\n", loginURL.String())
		if vaultAWSIssueAutoOpenURL {
			check(browser.OpenURL(loginURL.String()))
		}
	},
}

var vaultAWSListRolesCmd = &cobra.Command{
	Use:   "list-roles <account>",
	Short: "List available roles for the specified account",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client, err := getVaultClientWithToken()
		check(err)
		v := vault.New(client)

		roles, err := v.AWSListRoles(args[0])
		check(err)
		for _, role := range roles {
			fmt.Println(role)
		}
	},
}