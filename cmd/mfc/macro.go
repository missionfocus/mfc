package main

import (
	"git.missionfocus.com/ours/code/tools/mfc/pkg/vault"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(macroCmd)
	macroCmd.AddCommand(minioMacro)
}

var macroCmd = &cobra.Command{
	Use:   "macro",
	Short: "Perform a high-level operation",
}

const MinioBasePath = "secret/data/ci/shared/minio"
const MinioProfileName = "minio"

var minioMacro = &cobra.Command{
	Use:   "minio",
	Short: "Configure AWS credentials to use Jackie as an S3 implementation",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := getClientWithToken()
		check(err)

		v := vault.New(client)
		credentialsSecret, err := v.KvReadAws(MinioBasePath)
		check(err)
		check(credentialsSecret.ToProfile(credentialsPath, MinioProfileName))
		silentPrintf("AWS profile `%s` updated with the credentials read from `%s`.\n", MinioProfileName, MinioBasePath)

		endpointSecret, err := client.Logical().Read(MinioBasePath)
		check(err)
		data := endpointSecret.Data["data"].(map[string]interface{})
		endpoint := data["endpoint"].(string)
		silentPrintf("Configure your S3 client to use the following endpoint:\n\n\t%s\n", endpoint)
	},
}
