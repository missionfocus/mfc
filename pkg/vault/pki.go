package vault

import (
	"encoding/json"
	"io"
	"path"

	"github.com/pkg/errors"
)

const (
	DefaultPKIEnginePath = "pki"
	DefaultPKIEngineRole = "missionfocus-dot-com"
)

type PKIIssueOptions struct {
	RoleName          string
	CommonName        string
	AltNames          string
	IPSANs            string
	URISANs           string
	OtherSANs         string
	TTL               string
	Format            string
	PrivateKeyFormat  string
	ExcludeCNFromSANs bool
}

type PKIIssueSecret struct {
	Certificate    string `json:"certificate"`
	IssuingCA      string `json:"issuing_ca"`
	CAChain        string `json:"ca_chain"`
	PrivateKey     string `json:"private_key"`
	PrivateKeyType string `json:"private_key_type"`
	SerialNumber   string `json:"serial_number"`
}

func NewPKIIssueSecret(data map[string]interface{}) (*PKIIssueSecret, error) {
	var secret PKIIssueSecret
	m, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(m, &secret)
	return &secret, err
}

func (s *PKIIssueSecret) WriteChain(w io.Writer) error {
	if _, err := io.WriteString(w, s.Certificate+"\n"); err != nil {
		return errors.Wrap(err, "failed to write certificate")
	}
	if _, err := io.WriteString(w, s.IssuingCA); err != nil {
		return errors.Wrap(err, "failed to write issuing CA certificate")
	}
	return nil
}

func (s *PKIIssueSecret) WritePrivateKey(w io.Writer) error {
	_, err := io.WriteString(w, s.PrivateKey)
	if err != nil {
		return errors.Wrap(err, "failed to write private key")
	}
	return nil
}

func (s *PKIIssueSecret) WriteCertificate(w io.Writer) error {
	_, err := io.WriteString(w, s.Certificate)
	if err != nil {
		return errors.Wrap(err, "failed to write certificate")
	}
	return nil
}

func (s *PKIIssueSecret) WriteJSON(w io.Writer) error {
	return json.NewEncoder(w).Encode(s)
}

func (v *vault) PKIIssue(options *PKIIssueOptions) (*PKIIssueSecret, error) {
	secret, err := v.Logical().Write(path.Join(DefaultPKIEnginePath, "issue", options.RoleName), map[string]interface{}{
		"common_name":          options.CommonName,
		"alt_names":            options.AltNames,
		"ip_sans":              options.IPSANs,
		"uri_sans":             options.URISANs,
		"other_sans":           options.OtherSANs,
		"ttl":                  options.TTL,
		"format":               options.Format,
		"private_key_format":   options.PrivateKeyFormat,
		"exclude_cn_from_sans": options.ExcludeCNFromSANs,
	})
	if err != nil {
		return nil, err
	}
	return NewPKIIssueSecret(secret.Data)
}

func (v *vault) PKIGetCACert(enginePath string) (string, error) {
	secret, err := v.Logical().Read(path.Join(enginePath, "cert", "ca"))
	if err != nil {
		return "", err
	}
	return secret.Data["certificate"].(string), nil
}
