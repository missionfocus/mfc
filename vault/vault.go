package vault

import (
	"github.com/hashicorp/vault/api"
	"strings"
)

type Vault interface {
	ReadSTS(account string, role string) (*api.Secret, error)
	AuthLDAP(username string, password string) (string, error)
	ListAllKV(key string) []string
}

type vault struct {
	*api.Client
}

func New(client *api.Client) Vault {
	return &vault{client}
}

func (v *vault) ReadSTS(account string, role string) (*api.Secret, error) {
	secret, err := v.Logical().Read(strings.Join([]string{account, "sts", role}, "/"))
	if err != nil {
		return nil, err
	}

	return secret, nil
}

func (v *vault) AuthLDAP(username string, password string) (string, error) {
	endpoint := strings.Join([]string{"auth", "ldap", "login", username}, "/")
	data := map[string]interface{}{
		"password": password,
	}

	secret, err := v.Logical().Write(endpoint, data)
	if err != nil {
		return "", err
	}

	return secret.Auth.ClientToken, nil
}

func (v *vault) ListAllKV(key string) []string {
	tree := NewKvTree(v.Client, key)
	keys := make([]string, 0)

	tree.Traverse(func(node *TreeNode) {
		if node.Err != nil {
			keys = append(keys, "Error: Could not list key: "+node.Err.Error())
			return
		}
		keys = append(keys, node.Key)
	})

	return keys
}
