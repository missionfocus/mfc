package vault

import (
	"github.com/hashicorp/vault/api"
	"strings"
)

type Vault interface {
	ReadSTS(account string, role string) (*STSSecret, error)
}

type vault struct {
	*api.Client
}

func New(client *api.Client) Vault {
	return &vault{client}
}

func (v *vault) ReadSTS(account string, role string) (*STSSecret, error) {
	secret, err := v.Logical().Read(strings.Join([]string{account, "sts", role}, "/"))
	if err != nil {
		return nil, err
	}

	return &STSSecret{
		AccessKeyID:     secret.Data["access_key"].(string),
		SecretAccessKey: secret.Data["secret_key"].(string),
		SecurityToken:   secret.Data["security_token"].(string),
	}, nil
}
