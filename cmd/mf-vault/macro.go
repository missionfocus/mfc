package mf_vault

import (
	"git.missionfocus.com/open-source/mf-vault/pkg/vault"
	"github.com/spf13/cobra"
	"path"
)

func init() {
	rootCmd.AddCommand(macroCmd)
	macroCmd.AddCommand(minioMacro)
}

var macroCmd = &cobra.Command{
	Use:   "macro",
	Short: "Perform a high-level operation",
}

const MinioBasePath = "secret/minio"
const MinioProfileName = "minio"

var minioMacro = &cobra.Command{
	Use:   "minio",
	Short: "Configure AWS credentials to use Jackie as an S3 implementation",
	Run: func(cmd *cobra.Command, args []string) {
		key := path.Join(MinioBasePath, "aws-credentials")

		client, err := getClientWithToken()
		check(err)

		v := vault.New(client)
		credentialsSecret, err := v.KvReadAws(key)
		check(err)
		check(credentialsSecret.ToProfile(credentialsPath, MinioProfileName))
		silentPrintf("AWS profile `%s` updated with the credentials read from `%s`.\n", MinioProfileName, key)

		endpointSecret, err := client.Logical().Read(path.Join(MinioBasePath, "end-point"))
		check(err)
		endpoint := endpointSecret.Data["end-point"].(string)
		silentPrintf("Configure your S3 client to use the following endpoint:\n\n\t%s\n", endpoint)
	},
}
