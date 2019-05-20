package vault

import (
	"encoding/base64"
	"github.com/hashicorp/vault/api"
	"io/ioutil"
	"regexp"
)

type NPMSecret struct {
	Username string
	Password string
}

func NewNPMSecret(secret *api.Secret) *NPMSecret {
	return &NPMSecret{
		Username: secret.Data["username"].(string),
		Password: secret.Data["password"].(string),
	}
}

func (s *NPMSecret) Base64() string {
	return base64.StdEncoding.EncodeToString([]byte(s.Username + ":" + s.Password))
}

var authRe = regexp.MustCompile(`(?m)^_auth=[0-9A-z+/=]+$`)

func (s *NPMSecret) UpdateNpmrc(path string) error {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	authStr := "_auth=" + s.Base64()
	if authRe.Match(content) {
		content = authRe.ReplaceAll(content, []byte(authStr))
	} else {
		if content[len(content)-1] != '\n' {
			content = append(content, '\n')
		}
		content = append(content, []byte(authStr+"\n")...)
	}
	return ioutil.WriteFile(path, content, 0600)
}
