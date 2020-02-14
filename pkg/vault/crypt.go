package vault

import (
	"encoding/base64"
	"path"

	"github.com/hashicorp/vault/api"
)

const CryptTransitEngineName = "transit"

type CryptClient struct {
	v *api.Client
}

func NewCryptClient(v *api.Client) *CryptClient {
	return &CryptClient{v: v}
}

func (c *CryptClient) En(recipient string, pt string) (string, error) {
	name, err := c.keyName(recipient)
	if err != nil {
		return "", err
	}

	sec, err := c.v.Logical().Write(path.Join(CryptTransitEngineName, "encrypt", name), map[string]interface{}{
		"plaintext": base64.StdEncoding.EncodeToString([]byte(pt)),
	})
	if err != nil {
		return "", err
	}
	return sec.Data["ciphertext"].(string), nil
}

func (c *CryptClient) De(recipient string, ct string) (string, error) {
	name, err := c.keyName(recipient)
	if err != nil {
		return "", err
	}

	sec, err := c.v.Logical().Write(path.Join(CryptTransitEngineName, "decrypt", name), map[string]interface{}{
		"ciphertext": ct,
	})
	if err != nil {
		return "", err
	}

	dec, err := base64.StdEncoding.DecodeString(sec.Data["plaintext"].(string))
	if err != nil {
		return "", err
	}
	return string(dec), nil
}

func (c *CryptClient) keyName(recipient string) (string, error) {
	if recipient == "" {
		tok, err := c.v.Auth().Token().LookupSelf()
		if err != nil {
			return "", err
		}
		return tok.Data["display_name"].(string), nil
	}
	return "ldap-" + recipient, nil
}
