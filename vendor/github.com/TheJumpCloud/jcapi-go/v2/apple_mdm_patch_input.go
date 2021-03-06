/*
 * JumpCloud APIs
 *
 *  JumpCloud's V2 API. This set of endpoints allows JumpCloud customers to manage objects, groupings and mappings and interact with the JumpCloud Graph.
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package v2

type AppleMdmPatchInput struct {

	// A signed certificate obtained from Apple after providing Apple with the plist file provided on POST.
	AppleSignedCert string `json:"appleSignedCert,omitempty"`

	// A new name for the Apple MDM configuration.
	Name string `json:"name,omitempty"`
}
