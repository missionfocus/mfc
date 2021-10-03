package vault

import (
	"path"

	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/vault/api"
)

const (
	NPMBasePath   = "secret/data/ci/shared/npm"
	MinioBasePath = "secret/data/ci/shared/minio"
	UserBasePath  = "secret/data/user"
)

type KVItem struct {
	Key  string                 `yaml:"key"`
	Data map[string]interface{} `yaml:"data,omitempty"`
}

func (v *vault) KvListAll(key string) []string {
	tree := NewKVTree(v.Client, key)
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

func (v *vault) KvNPMAuth(key string) (*NPMSecret, error) {
	secret, err := v.Logical().Read(key)
	if err != nil {
		return nil, err
	}
	return NewNPMSecret(secret), nil
}

func (v *vault) KVGetAll(key string) ([]KVItem, []*TreeNode) {
	items := make([]KVItem, 0)
	errNodes := make([]*TreeNode, 0)

	NewKVTree(v.Client, key).Traverse(func(node *TreeNode) {
		if node.Err != nil {
			errNodes = append(errNodes, node)
		}

		// Non-leaf node, don't add
		if node.Key[len(node.Key)-1] == '/' {
			return
		}

		items = append(items, KVItem{Key: node.Key, Data: node.Secret.Data})
	})

	return items, errNodes
}

func (v *vault) KVPutAll(items []KVItem) error {
	me := &multierror.Error{}
	for _, item := range items {
		data := map[string]interface{}{"data": item.Data}
		if _, err := v.Logical().Write(item.Key, data); err != nil {
			me = multierror.Append(me, err)
		}
	}
	return me.ErrorOrNil()
}

func (v *vault) KVUserGet(key string) (*api.Secret, error) {
	self, err := v.lookupSelf()
	if err != nil {
		return nil, err
	}

	return v.Logical().Read(path.Join(UserBasePath, self, key))
}

func (v *vault) KVUserWrite(key string, data map[string]interface{}) (*api.Secret, error) {
	self, err := v.lookupSelf()
	if err != nil {
		return nil, err
	}

	return v.Logical().Write(path.Join(UserBasePath, self, key), data)
}

func (v *vault) lookupSelf() (string, error) {
	tok, err := v.Auth().Token().LookupSelf()
	if err != nil {
		return "", err
	}
	return tok.Data["display_name"].(string), nil
}
