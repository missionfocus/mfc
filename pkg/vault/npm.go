package vault

import (
	"bytes"
	"github.com/hashicorp/vault/api"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
)

type NPMSecret struct {
	Registry string
	Token    string
}

func NewNPMSecret(secret *api.Secret) *NPMSecret {
	data := secret.Data["data"].(map[string]interface{})
	return &NPMSecret{
		Registry: data["registry"].(string),
		Token:    data["token"].(string),
	}
}

var authRe = regexp.MustCompile(`(?m)^//(.+):_authToken=(.+)$`)

func (s *NPMSecret) UpdateNpmrc(path string) error {
	f, err := openOrCreate(path, 0700, 0600)
	if err != nil {
		return err
	}
	content, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}

	// First case, the registry already exists
	for _, match := range authRe.FindAllSubmatch(content, -1) {
		if !bytes.Equal(match[1], []byte(s.Registry)) {
			continue
		}

		updated := bytes.Replace(match[0], match[2], []byte(s.Token), 1)
		content = bytes.Replace(content, match[0], updated, 1)

		return ioutil.WriteFile(path, content, 0600)
	}

	// Other case, it doesn't exist, so append it
	content = append(content, []byte("\n//"+s.Registry+":_authToken="+s.Token)...)
	return ioutil.WriteFile(path, content, 0600)
}

func openOrCreate(path string, dirPerm, filePerm os.FileMode) (*os.File, error) {
	if exists, err := fileExists(path); exists {
		return os.Open(path)
	} else if err != nil {
		return nil, err
	}

	dir, _ := filepath.Split(path)
	if err := os.MkdirAll(dir, dirPerm); err != nil {
		return nil, err
	}

	return os.OpenFile(path, os.O_RDWR|os.O_CREATE, filePerm)
}

func fileExists(path string) (bool, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false, nil
	} else if err != nil {
		return false, err
	} else {
		return true, nil
	}
}
