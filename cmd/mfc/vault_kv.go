package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"

	"github.com/missionfocus/mfc/pkg/vault"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

func init() {
	vaultCmd.AddCommand(vaultKVCmd)
	vaultKVCmd.AddCommand(vaultKVListAllCmd)
	vaultKVCmd.AddCommand(vaultKVAwsCmd)
	vaultKVCmd.AddCommand(vaultKVGPGCmd)
	vaultKVCmd.AddCommand(vaultKVNPMCmd)
	vaultKVCmd.AddCommand(vaultKVGetAllCmd)
	vaultKVCmd.AddCommand(vaultKVPutAllCmd)
	vaultKVCmd.AddCommand(vaultKVUserCmd)
	vaultKVNPMCmd.AddCommand(kvNPMAuthCmd)

	vaultKVUserCmd.AddCommand(vaultKVUserGetCmd)
	vaultKVUserCmd.AddCommand(vaultKVUserWriteCmd)

	vaultKVAwsCmd.PersistentFlags().StringVarP(&vaultKVAwsProfileName, "profile", "p", "vault", "name of the profile")

	kvNPMAuthCmd.PersistentFlags().BoolVar(&vautlKVNPMStdout, "stdout", false, "write the NPM auth token to stdout instead of .npmrc")
	kvNPMAuthCmd.PersistentFlags().StringVarP(&vaultKVNPMRCPath, "path", "p", filepath.Join(homeDir(), ".npmrc"), "path to .npmrc")
}

var vaultKVCmd = &cobra.Command{
	Use:   "kv",
	Short: "Interact with Vault's Key/Value engine",
}

const vaultKVListAllExample = `
	mfc vault kv listall /secret/user/ldap-<user>/        # Lists all accessible keys under specified K/V engine key`

var vaultKVListAllCmd = &cobra.Command{
	Use:     "listall <key>",
	Short:   "Lists all keys under the specified K/V engine key. Key must end with `/`",
	Example: vaultKVListAllExample,
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

const vaultKVGetAllExample = `
	mfc vault kv getall /secret/data/user/ldap-<username>/ # Gets recursive listing of available data`

var vaultKVGetAllCmd = &cobra.Command{
	Use:     "getall <key>",
	Short:   "Recursively gets the data for all keys under the specified path as YAML",
	Example: vaultKVGetAllExample,
	Args:    cobra.ExactArgs(1),
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

const vaultPutAllExample = `
	mfc vault kv putall ./key.yml # Gets JSON data from a file`

var vaultKVPutAllCmd = &cobra.Command{
	Use:     "putall <file>",
	Short:   "Puts all keys in the specified YAML file into the KV engine",
	Example: vaultPutAllExample,
	Args:    cobra.ExactArgs(1),
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

const vaultKVAwsExample = `
	mf-vault aws issue missionfocus engineer # Sets npm fontawsome credentials
	mf-vault aws issue sandbox engineer      # Sets npm fontawsome credentials`

var vaultKVAwsCmd = &cobra.Command{
	Use:     "aws <path>",
	Short:   "Read the secret at `path` as AWS credentials",
	Example: vaultKVAwsExample,
	Args:    cobra.ExactArgs(1),
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

var vaultKVNPMCmd = &cobra.Command{
	Use:   "npm",
	Short: "Interact with NPM configuration stored in Vault",
}

var (
	vaultKVNPMRCPath string
	vautlKVNPMStdout bool
)

const vaultNPMAuthExample = `
	mfc vault kv npm auth fontawsome -p $HOME/.npmrc  # Sets npm fontawsome credentials
	mfc vault kv npm auth nexus -p $HOME/.npmrc       # Sets npm fontawsome credentials
	mfc vault kv npm auth fortawsome  -p $HOME/.npmrc # Sets fortawsome credentials`

var kvNPMAuthCmd = &cobra.Command{
	Use:     "auth <registry>",
	Short:   "Update .npmrc with authentication data from Vault",
	Example: vaultNPMAuthExample,
	Args:    cobra.ExactArgs(1),
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

var vaultKVUserCmd = &cobra.Command{
	Use:   "user",
	Short: "Interact with the user's personal namespace in Vault",
}

const vaultKVUserGetExample = `
	mfc vault kv user get gitlab | jq -r '.token' # Gets value for a key`

var vaultKVUserGetCmd = &cobra.Command{
	Use:     "get <key>",
	Short:   "Get the value at the specified key",
	Example: vaultKVUserGetExample,
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client, err := getVaultClientWithToken()
		check(err)
		v := vault.New(client)

		secret, err := v.KVUserGet(args[0])
		check(err)
		if secret == nil {
			check(fmt.Errorf("unknown key %s", args[0]))
		}

		check(json.NewEncoder(os.Stdout).Encode(secret.Data["data"]))
	},
}

var vaultKVUserWriteExample = `
	echo '{ "data": { "mysecret": "some secret data!" } }' | mfc vault kv user write mykey # Writes JSON data from stdin
	mfc vault kv user write mykey mydata.json                                              # Writes JSON data from a file`

var vaultKVUserWriteCmd = &cobra.Command{
	Use:     "write <key> [file]",
	Short:   "Write JSON data to the specified key",
	Args:    cobra.RangeArgs(1, 2),
	Example: vaultKVUserWriteExample,
	Run: func(cmd *cobra.Command, args []string) {
		client, err := getVaultClientWithToken()
		check(err)
		v := vault.New(client)

		r := os.Stdin
		if len(args) > 0 {
			f, err := os.Open(args[1])
			check(err)
			defer f.Close()
			r = f
		}

		data := make(map[string]interface{})
		check(json.NewDecoder(r).Decode(&data))

		_, err = v.KVUserWrite(args[0], data)
		check(err)
	},
}
