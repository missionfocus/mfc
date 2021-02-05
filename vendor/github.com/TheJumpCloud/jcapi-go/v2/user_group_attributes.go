/*
 * JumpCloud APIs
 *
 *  JumpCloud's V2 API. This set of endpoints allows JumpCloud customers to manage objects, groupings and mappings and interact with the JumpCloud Graph.
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package v2

type UserGroupAttributes struct {

	PosixGroups []UserGroupAttributesPosixGroups `json:"posixGroups,omitempty"`

	SambaEnabled bool `json:"sambaEnabled,omitempty"`
}
