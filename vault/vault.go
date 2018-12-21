package vault

import (
	"github.com/hashicorp/vault/api"
	"strings"
)

type Vault interface {
	ReadSTS(account string, role string) (*api.Secret, error)
	AuthLDAP(username string, password string) (string, error)
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
