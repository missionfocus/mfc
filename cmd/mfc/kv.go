package mfc

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
	rootCmd.AddCommand(kvCmd)

	kvCmd.AddCommand(kvListAllCmd)
	kvCmd.AddCommand(kvAwsCmd)
	kvCmd.AddCommand(kvGpgCmd)
	kvCmd.AddCommand(kvNPMCmd)
	kvCmd.AddCommand(kvGetAllCmd)
	kvCmd.AddCommand(kvPutAllCmd)

	kvAwsCmd.PersistentFlags().StringVarP(&kvAwsProfileName, "profile", "p", "vault", "name of the profile")

	kvGpgImportCmd.PersistentFlags().BoolVar(&kvGpgImportPrivate, "private", false, "import the pair's private key")
	kvGpgCmd.AddCommand(kvGpgImportCmd)

	kvNPMAuthCmd.PersistentFlags().BoolVar(&kvNPMStdout, "stdout", false, "write the NPM auth token to stdout instead of .npmrc")
	kvNPMAuthCmd.PersistentFlags().StringVarP(&kvNPMRcPath, "path", "p", filepath.Join(homeDir(), ".npmrc"), "path to .npmrc")
	kvNPMCmd.AddCommand(kvNPMAuthCmd)
}

var kvCmd = &cobra.Command{
	Use:   "kv",
	Short: "Interact with Vault's Key/Value engine",
}

var kvListAllCmd = &cobra.Command{
	Use:     "listall <key>",
	Short:   "Lists all keys under the specified K/V engine key. Key must end with `/`",
	Example: "  listall secret/\tLists all keys under the default K/V secrets engine",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client, err := getClientWithToken()
		check(err)

		v := vault.New(client)
		keys := v.KvListAll(args[0])
		for _, key := range keys {
			fmt.Println(key)
		}
	},
}

var kvGetAllCmd = &cobra.Command{
	Use:   "getall <key>",
	Short: "Recursively gets the data for all keys under the specified path as YAML",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client, err := getClientWithToken()
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

var kvPutAllCmd = &cobra.Command{
	Use:   "putall <file>",
	Short: "Puts all keys in the specified YAML file into the KV engine",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client, err := getClientWithToken()
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

var kvAwsProfileName string

var kvAwsCmd = &cobra.Command{
	Use:   "aws <path>",
	Short: "Read the secret at `path` as AWS credentials",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		key := args[0]

		client, err := getClientWithToken()
		check(err)

		v := vault.New(client)
		secret, err := v.KvReadAws(key)
		check(err)

		check(secret.ToProfile(credentialsPath, kvAwsProfileName))
		silentPrintf("AWS profile `%s` updated with the credentials read from `%s`.\n", kvAwsProfileName, key)
	},
}

var kvGpgCmd = &cobra.Command{
	Use:   "gpg",
	Short: "Interact with GPG keys stored in Vault",
}

var kvGpgImportPrivate bool

var kvGpgImportCmd = &cobra.Command{
	Use:   "import <key>",
	Short: "Import the GPG key at the specified KV engine key",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		key := args[0]

		client, err := getClientWithToken()
		check(err)

		v := vault.New(client)
		out, err := v.KvGpgImport(key, kvGpgImportPrivate)
		silentPrint(string(out))
		check(err)
	},
}

var kvNPMCmd = &cobra.Command{
	Use:   "npm",
	Short: "Interact with NPM configuration stored in Vault",
}

var (
	kvNPMRcPath string
	kvNPMStdout bool
)

const npmBasePath = "secret/data/ci/shared/npm"

var kvNPMAuthCmd = &cobra.Command{
	Use:   "auth <registry>",
	Short: "Update .npmrc with authentication data from Vault",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client, err := getClientWithToken()
		check(err)
		v := vault.New(client)

		secret, err := v.KvNPMAuth(path.Join(npmBasePath, args[0]))
		check(err)

		if kvNPMStdout {
			fmt.Print(secret.Token)
			return
		}
		check(secret.UpdateNpmrc(kvNPMRcPath))
	},
}
