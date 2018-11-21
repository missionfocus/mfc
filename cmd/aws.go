package cmd

import (
	"fmt"
	"git.missionfocus.com/devops/mf-vault/vault"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

var credentialsPath string
var profileName string

func init() {
	rootCmd.AddCommand(awsCmd)

	defaultCredentialsPath := filepath.Join(os.Getenv("HOME"), ".aws", "credentials")
	awsCmd.PersistentFlags().StringVarP(&credentialsPath, "credentials", "c", defaultCredentialsPath, "path to AWS credentials file")
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
			fatal(err)
		}
		v := vault.New(client)

		secret, err := v.ReadSTS(account, role)
		if err != nil {
			fatal(err)
		}

		if err := secret.ToProfile(credentialsPath, profileName); err != nil {
			fatal(err)
		}

		fmt.Printf("AWS profile `%s` updated with credentials for IAM role `%s` of account `%s`.\n", profileName, role, account)

		loginUrl, err := secret.GenerateLoginUrl(account)
		if err != nil {
			fatal(err)
		}

		fmt.Printf("Console login URL (valid for 15 minutes):\n\n%s\n", loginUrl.String())
	},
}
