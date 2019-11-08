## mf-vault kv

Interact with Vault's Key/Value engine

### Synopsis

Interact with Vault's Key/Value engine

### Options

```
  -h, --help   help for kv
```

### Options inherited from parent commands

```
      --aws-creds-file string   path to AWS credentials file (default "/Users/ccday/.aws/credentials")
      --silent                  suppress output to stdout
      --token-file string       path to vault token file (default "/Users/ccday/.vault-token")
```

### SEE ALSO

* [mf-vault](mf-vault.md)	 - CLI for interacting with the Mission Focus Vault
* [mf-vault kv aws](mf-vault_kv_aws.md)	 - Read the secret at `path` as AWS credentials
* [mf-vault kv getall](mf-vault_kv_getall.md)	 - Recursively gets the data for all keys under the specified path as YAML
* [mf-vault kv gpg](mf-vault_kv_gpg.md)	 - Interact with GPG keys stored in Vault
* [mf-vault kv listall](mf-vault_kv_listall.md)	 - Lists all keys under the specified K/V engine key. Key must end with `/`
* [mf-vault kv npm](mf-vault_kv_npm.md)	 - Interact with NPM configuration stored in Vault
* [mf-vault kv putall](mf-vault_kv_putall.md)	 - Puts all keys in the specified YAML file into the KV engine

###### Auto generated by spf13/cobra on 8-Nov-2019