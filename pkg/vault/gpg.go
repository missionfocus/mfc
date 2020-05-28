package vault

import (
	"bytes"
	"encoding/base64"
	"git.missionfocus.com/ours/code/libraries/go/gpg"
	"github.com/hashicorp/vault/api"
	"github.com/pkg/errors"
)

// Secret that represents a PGP keypair.
type PGPSecret struct {
	Passphrase string
	PrivateKey string
	PublicKey  string
}

// Create a PGP Secret from a generic Vault secret.
func NewPGPSecret(secret *api.Secret) (*PGPSecret, error) {
	data := secret.Data["data"].(map[string]interface{})
	rawPassphrase, ok := data["passphrase"]
	if !ok {
		return nil, errors.New("failed to get passphrase: key does not exist")
	}
	passphrase, ok := rawPassphrase.(string)
	if !ok {
		return nil, errors.New("failed to get passphrase: value must be of type string")
	}

	private, err := getB64EncodedValue(data, "private")
	if err != nil {
		return nil, errors.Wrap(err, "failed to get private key")
	}

	public, err := getB64EncodedValue(data, "public")
	if err != nil {
		return nil, errors.Wrap(err, "failed to get public key")
	}

	return &PGPSecret{
		Passphrase: passphrase,
		PrivateKey: private,
		PublicKey:  public,
	}, nil
}

func getB64EncodedValue(m map[string]interface{}, key string) (string, error) {
	raw, ok := m[key]
	if !ok {
		return "", errors.New("key does not exist")
	}

	encoded, ok := raw.(string)
	if !ok {
		return "", errors.New("value must be of type string")
	}

	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return "", err
	}

	return string(decoded), nil
}

// Imports the secret's private key.
func (s *PGPSecret) ImportPrivate() ([]byte, error) {
	return gpg.ImportWithPassphrase(bytes.NewBufferString(s.PrivateKey), s.Passphrase)
}

// Imports the secret's public key.
func (s *PGPSecret) ImportPublic() ([]byte, error) {
	return gpg.Import(bytes.NewBufferString(s.PublicKey))
}
