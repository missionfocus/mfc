package vault

import (
	"github.com/hashicorp/vault/api"
	"github.com/pkg/errors"
	"io"
	"time"
)

type PKISecret struct {
	Certificate    string
	Expiration     time.Time
	IssuingCa      string
	PrivateKey     string
	PrivateKeyType string
	SerialNumber   string
}

func NewPKISecret(secret *api.Secret) *PKISecret {
	var pkiSecret PKISecret

	if cert, ok := secret.Data["certificate"]; ok {
		pkiSecret.Certificate = cert.(string)
	}

	if exp, ok := secret.Data["expiration"]; ok {
		pkiSecret.Expiration = time.Unix(int64(exp.(float64)), 0)
	}

	if iss, ok := secret.Data["issuing_ca"]; ok {
		pkiSecret.IssuingCa = iss.(string)
	}

	if priv, ok := secret.Data["private_key"]; ok {
		pkiSecret.PrivateKey = priv.(string)
	}

	if keyType, ok := secret.Data["private_key_type"]; ok {
		pkiSecret.PrivateKeyType = keyType.(string)
	}

	if sn, ok := secret.Data["serial_number"]; ok {
		pkiSecret.SerialNumber = sn.(string)
	}

	return &pkiSecret
}

func (s *PKISecret) WriteChain(w io.Writer) error {
	if _, err := io.WriteString(w, s.Certificate+"\n"); err != nil {
		return errors.Wrap(err, "failed to write certificate")
	}
	if _, err := io.WriteString(w, s.IssuingCa); err != nil {
		return errors.Wrap(err, "failed to write issuing CA certificate")
	}
	return nil
}

func (s *PKISecret) WritePrivateKey(w io.Writer) error {
	_, err := io.WriteString(w, s.PrivateKey)
	if err != nil {
		return errors.Wrap(err, "failed to write private key")
	}
	return nil
}
