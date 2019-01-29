package vault

import (
	"github.com/hashicorp/vault/api"
	"path"
	"strings"
)

type TreeNode struct {
	Key      string
	Secret   *api.Secret
	Children []*TreeNode
	Err      error
}

type Tree interface {
	Root() *TreeNode
	Traverse(visit func(node *TreeNode))
}

type kvTree struct {
	client  *api.Client
	rootKey string
	root    *TreeNode
}

// Creates a new Tree for a KV engine.
func NewKvTree(client *api.Client, rootKey string) Tree {
	return &kvTree{client, rootKey, nil}
}

// Returns the key/value subtree of the specified key for a KV engine.
func (t *kvTree) Root() *TreeNode {
	if t.root == nil {
		t.root = t.buildSubtree(t.rootKey)
	}
	return t.root
}

// Calls visit on each node of the KV tree.
func (t *kvTree) Traverse(visit func(node *TreeNode)) {
	t.traverse(visit, t.Root())
}

// Traverse recursive helper.
func (t *kvTree) traverse(visit func(node *TreeNode), node *TreeNode) {
	visit(node)
	for _, child := range node.Children {
		t.traverse(visit, child)
	}
}

// Calls the Vault API to recursively build the KV subtree at key.
func (t *kvTree) buildSubtree(key string) *TreeNode {
	// Base case, leaf node.
	if rune(key[len(key)-1]) != '/' {
		return &TreeNode{Key: key}
	}

	split := clean(strings.Split(key, "/"))
	secret, err := t.client.Logical().List(path.Join(split[0], "metadata", strings.Join(split[1:], "/")))
	if err != nil {
		return &TreeNode{Key: key, Err: err}
	}

	var children []*TreeNode
	keys, ok := secret.Data["keys"].([]interface{})
	if ok {
		children = make([]*TreeNode, len(keys))
		for i, k := range keys {
			children[i] = t.buildSubtree(strings.Join(append(split, k.(string)), "/"))
		}
	}

	return &TreeNode{Key: key, Secret: secret, Children: children}
}

// Removes empty values from a string slice.
func clean(slc []string) []string {
	cleaned := make([]string, 0)
	for _, str := range slc {
		if str != "" {
			cleaned = append(cleaned, str)
		}
	}
	return cleaned
}
