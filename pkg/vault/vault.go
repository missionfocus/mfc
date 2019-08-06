package vault

import (
	"github.com/hashicorp/vault/api"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type Vault interface {
	AuthLDAP(username string, password string) (string, error)
	AuthApprole(roleID string, secretID string) (string, error)

	AWSReadSTS(account string, role string, ttl string) (*api.Secret, error)
	AWSListRoles(account string) ([]string, error)

	KvListAll(key string) []string
	KvReadAws(key string) (*STSSecret, error)
	KvGpgImport(key string, private bool) ([]byte, error)
	KvNPMAuth(key string) (*NPMSecret, error)

	PkiCreateFiles(secret *api.Secret, path string) error

	SSHSignPubKey(publicKeyBytes []byte, usage string) (*api.Secret, error)
	SSHCA(path string) (string, error)
}

type vault struct {
	*api.Client
}

func New(client *api.Client) Vault {
	return &vault{client}
}

func (v *vault) AWSReadSTS(account string, role string, ttl string) (*api.Secret, error) {
	secret, err := v.Logical().Write(strings.Join([]string{account, "sts", role}, "/"), map[string]interface{}{
		"ttl": ttl,
	})
	if err != nil {
		return nil, err
	}

	return secret, nil
}

func (v *vault) AWSListRoles(account string) ([]string, error) {
	endpoint := path.Join(account, "roles")
	secret, err := v.Logical().List(endpoint)
	if err != nil {
		return nil, err
	}

	roles := make([]string, len(secret.Data["keys"].([]interface{})))
	for i, iface := range secret.Data["keys"].([]interface{}) {
		roles[i] = iface.(string)
	}
	return roles, nil
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

func (v *vault) AuthApprole(roleID string, secretID string) (string, error) {
	endpoint := path.Join("auth", "approle", "login")
	data := map[string]interface{}{
		"role_id":   roleID,
		"secret_id": secretID,
	}

	secret, err := v.Logical().Write(endpoint, data)
	if err != nil {
		return "", err
	}

	return secret.Auth.ClientToken, nil
}

func (v *vault) KvListAll(key string) []string {
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

func (v *vault) KvReadAws(key string) (*STSSecret, error) {
	secret, err := v.Logical().Read(key)
	if err != nil {
		return nil, err
	}
	return NewSTSSecret(secret), nil
}

func (v *vault) SSHSignPubKey(publicKeyBytes []byte, usage string) (*api.Secret, error) {
	if usage == "user" {
		sshClient := v.Client.SSHWithMountPoint("ssh-signer")
		data := map[string]interface{}{"public_key": string(publicKeyBytes)}
		return sshClient.SignKey("user-key", data)
	} else {
		sshClient := v.Client.SSHWithMountPoint("ssh-signer")
		data := map[string]interface{}{
			"public_key": string(publicKeyBytes),
			"cert_type":  "host",
		}
		return sshClient.SignKey("host-key", data)
	}
}

func (v *vault) KvGpgImport(key string, private bool) (out []byte, err error) {
	secret, err := v.Logical().Read(key)
	if err != nil {
		return nil, err
	}

	gpgSecret, err := NewPGPSecret(secret)
	if err != nil {
		return nil, err
	}

	if private {
		out, err = gpgSecret.ImportPrivate()
	} else {
		out, err = gpgSecret.ImportPublic()
	}

	return
}

func (v *vault) PkiCreateFiles(secret *api.Secret, path string) error {
	pkiSecret := NewPKISecret(secret)

	fullchain, err := os.OpenFile(filepath.Join(path, "fullchain.pem"), os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer fullchain.Close()
	if err := pkiSecret.WriteChain(fullchain); err != nil {
		return err
	}

	privkey, err := os.OpenFile(filepath.Join(path, "privkey.pem"), os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer privkey.Close()
	if err := pkiSecret.WritePrivateKey(privkey); err != nil {
		return err
	}

	return nil
}

func (v *vault) KvNPMAuth(key string) (*NPMSecret, error) {
	secret, err := v.Logical().Read(key)
	if err != nil {
		return nil, err
	}
	return NewNPMSecret(secret), nil
}

func (v *vault) SSHCA(p string) (string, error) {
	secret, err := v.Logical().Read(path.Join(p, "config", "ca"))
	if err != nil {
		return "", err
	}
	return secret.Data["public_key"].(string), nil
}
