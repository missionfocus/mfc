package cmd

import (
	"fmt"
	"git.missionfocus.com/open-source/mf-vault/vault"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"time"
)

var credentialsPath string
var profileName string
var silent bool

func init() {
	rootCmd.AddCommand(awsCmd)

	defaultCredentialsPath := filepath.Join(os.Getenv("HOME"), ".aws", "credentials")
	awsCmd.PersistentFlags().StringVarP(&credentialsPath, "credentials", "c", defaultCredentialsPath, "path to AWS credentials file")
	awsCmd.PersistentFlags().StringVarP(&profileName, "profile", "p", "vault", "name of the profile")
	awsCmd.PersistentFlags().BoolVarP(&silent, "silent", "s", false, "update AWS credentials with no output to stdout")
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
		stsSecret := vault.NewSTSSecret(secret)
		if err := stsSecret.ToProfile(credentialsPath, profileName); err != nil {
			fatal(err)
		}

		if silent {
			return
		}

		fmt.Printf("AWS profile `%s` updated with credentials for IAM role `%s` of account `%s`.\n", profileName, role, account)
		fmt.Printf("These credentials are valid for: %s\n", (time.Second * time.Duration(secret.LeaseDuration)).String())

		loginUrl, err := stsSecret.GenerateLoginUrl(account)
		if err != nil {
			fatal(err)
		}

		fmt.Printf("Console login URL (valid for 15 minutes):\n\n%s\n", loginUrl.String())
	},
}
