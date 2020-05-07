package main

import (
	"fmt"
	"git.missionfocus.com/ours/code/tools/mfc/pkg/vault"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
)

func init() {
	vaultCmd.AddCommand(vaultKVCmd)
	vaultKVCmd.AddCommand(vaultKVListAllCmd)
	vaultKVCmd.AddCommand(vaultKVAwsCmd)
	vaultKVCmd.AddCommand(vaultKVGPGCmd)
	vaultKVCmd.AddCommand(vaultKVNPMCmd)
	vaultKVCmd.AddCommand(vaultKVGetAllCmd)
	vaultKVCmd.AddCommand(vaultKVPutAllCmd)
	vaultKVGPGCmd.AddCommand(vaultKVGpgImportCmd)
	vaultKVNPMCmd.AddCommand(kvNPMAuthCmd)

	vaultKVAwsCmd.PersistentFlags().StringVarP(&vaultKVAwsProfileName, "profile", "p", "vault", "name of the profile")
	vaultKVGpgImportCmd.PersistentFlags().BoolVar(&vaultKVGpgImportPrivate, "private", false, "import the pair's private key")

	kvNPMAuthCmd.PersistentFlags().BoolVar(&vautlKVNPMStdout, "stdout", false, "write the NPM auth token to stdout instead of .npmrc")
	kvNPMAuthCmd.PersistentFlags().StringVarP(&vaultKVNPMRCPath, "path", "p", filepath.Join(homeDir(), ".npmrc"), "path to .npmrc")
}

var vaultKVCmd = &cobra.Command{
	Use:   "kv",
	Short: "Interact with Vault's Key/Value engine",
}

var vaultKVListAllCmd = &cobra.Command{
	Use:     "listall <key>",
	Short:   "Lists all keys under the specified K/V engine key. Key must end with `/`",
	Example: "  listall secret/\tLists all keys under the default K/V secrets engine",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client, err := getVaultClientWithToken()
		check(err)

		v := vault.New(client)
		keys := v.KvListAll(args[0])
		for _, key := range keys {
			fmt.Println(key)
		}
	},
}

var vaultKVGetAllCmd = &cobra.Command{
	Use:   "getall <key>",
	Short: "Recursively gets the data for all keys under the specified path as YAML",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client, err := getVaultClientWithToken()
		check(err)
		v := vault.New(client)

		items, errNodes := v.KVGetAll(args[0])
		y, err := yaml.Marshal(items)
		check(err)
		fmt.Print(string(y))

		if len(errNodes) > 0 {
			fmt.Fprintln(os.Stderr, "\nerror: could not get the following keys:")
			for _, node := range errNodes {
				_, _ = fmt.Fprintln(os.Stderr, node.Key)
			}
		}
	},
}

var vaultKVPutAllCmd = &cobra.Command{
	Use:   "putall <file>",
	Short: "Puts all keys in the specified YAML file into the KV engine",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client, err := getVaultClientWithToken()
		check(err)
		v := vault.New(client)

		content, err := ioutil.ReadFile(args[0])
		check(err)
		var items []vault.KVItem
		check(yaml.Unmarshal(content, &items))
		if err := v.KVPutAll(items); err != nil {
			fmt.Fprintln(os.Stderr, "error(s): writing keys:")
			fmt.Fprint(os.Stderr, err)
		}
	},
}

var vaultKVAwsProfileName string

var vaultKVAwsCmd = &cobra.Command{
	Use:   "aws <path>",
	Short: "Read the secret at `path` as AWS credentials",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		key := args[0]

		client, err := getVaultClientWithToken()
		check(err)

		v := vault.New(client)
		secret, err := v.KvReadAws(key)
		check(err)

		check(secret.ToProfile(mfcAWSCredentialsPath, vaultKVAwsProfileName))
		silentPrintf("AWS profile `%s` updated with the credentials read from `%s`.\n", vaultKVAwsProfileName, key)
	},
}

var vaultKVGPGCmd = &cobra.Command{
	Use:   "gpg",
	Short: "Interact with GPG keys stored in Vault",
}

var vaultKVGpgImportPrivate bool

var vaultKVGpgImportCmd = &cobra.Command{
	Use:   "import <key>",
	Short: "Import the GPG key at the specified KV engine key",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		key := args[0]

		client, err := getVaultClientWithToken()
		check(err)

		v := vault.New(client)
		out, err := v.KvGpgImport(key, vaultKVGpgImportPrivate)
		silentPrint(string(out))
		check(err)
	},
}

var vaultKVNPMCmd = &cobra.Command{
	Use:   "npm",
	Short: "Interact with NPM configuration stored in Vault",
}

var (
	vaultKVNPMRCPath string
	vautlKVNPMStdout bool
)

var kvNPMAuthCmd = &cobra.Command{
	Use:   "auth <registry>",
	Short: "Update .npmrc with authentication data from Vault",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client, err := getVaultClientWithToken()
		check(err)
		v := vault.New(client)

		secret, err := v.KvNPMAuth(path.Join(vault.NPMBasePath, args[0]))
		check(err)

		if vautlKVNPMStdout {
			fmt.Print(secret.Token)
			return
		}
		check(secret.UpdateNpmrc(vaultKVNPMRCPath))
	},
}
