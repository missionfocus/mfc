package main

import (
	"git.missionfocus.com/ours/code/tools/mfc/pkg/vault"
	"github.com/spf13/cobra"
)

func init() {
	vaultCmd.AddCommand(vaultMacroCmd)
	vaultMacroCmd.AddCommand(vaultMinioMacro)
}

var vaultMacroCmd = &cobra.Command{
	Use:   "macro",
	Short: "Perform a high-level operation",
}

const MinioProfileName = "minio"

const vaultMinioExample = `
  mfc vault macro minio # Configures ~/.aws/credentials for access to minio`

var vaultMinioMacro = &cobra.Command{
	Use:     "minio",
	Short:   "Configure AWS credentials to use Jackie as an S3 implementation",
	Example: vaultMinioExample,
	Run: func(cmd *cobra.Command, args []string) {
		client, err := getVaultClientWithToken()
		check(err)

		v := vault.New(client)
		credentialsSecret, err := v.KvReadAws(vault.MinioBasePath)
		check(err)
		check(credentialsSecret.ToProfile(mfcAWSCredentialsPath, MinioProfileName))
		silentPrintf("AWS profile `%s` updated with the credentials read from `%s`.\n", MinioProfileName, vault.MinioBasePath)

		endpointSecret, err := client.Logical().Read(vault.MinioBasePath)
		check(err)
		data := endpointSecret.Data["data"].(map[string]interface{})
		endpoint := data["endpoint"].(string)
		silentPrintf("Configure your S3 client to use the following endpoint:\n\n\t%s\n", endpoint)
	},
}
