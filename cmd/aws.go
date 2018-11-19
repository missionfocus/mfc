package cmd

import (
	"fmt"
	"git.missionfocus.com/devops/mf-vault/vault"
	"github.com/hashicorp/vault/api"
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
	Use:   "aws [account] [role]",
	Short: "Manages AWS Credentials",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		account := args[0]
		role := args[1]

		vClient, err := api.NewClient(&api.Config{
			Address: os.Getenv("VAULT_ADDR"),
		})
		if err != nil {
			fatal(err)
		}

		vClient.SetToken(os.Getenv("VAULT_TOKEN"))
		v := vault.New(vClient)

		secret, err := v.ReadSTS(account, role)
		if err != nil {
			fatal(err)
		}

		if err := secret.ToProfile(credentialsPath, profileName); err != nil {
			fatal(err)
		}

		fmt.Printf("AWS profile `%s` updated with Vault credentials.\n", profileName)

		loginUrl, err := secret.GenerateLoginUrl(account)
		if err != nil {
			fatal(err)
		}

		fmt.Printf("Temporary console login URL: %s\n", loginUrl.String())
	},
}
