
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCognitoIdentityPoolRolesAttachment describes a AwsCognitoIdentityPoolRolesAttachment resource
type AwsCognitoIdentityPoolRolesAttachment struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsCognitoIdentityPoolRolesAttachmentSpec	`json:"spec"`
}


// AwsCognitoIdentityPoolRolesAttachmentSpec is the spec for a AwsCognitoIdentityPoolRolesAttachment Resource
type AwsCognitoIdentityPoolRolesAttachmentSpec struct {
	IdentityPoolId	string	`json:"identity_pool_id"`
	RoleMapping	AwsCognitoIdentityPoolRolesAttachmentRoleMapping	`json:"role_mapping"`
	Roles	map[string]AwsCognitoIdentityPoolRolesAttachmentRoles	`json:"roles"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCognitoIdentityPoolRolesAttachmentList is a list of AwsCognitoIdentityPoolRolesAttachment resources
type AwsCognitoIdentityPoolRolesAttachmentList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsCognitoIdentityPoolRolesAttachment	`json:"items"`
}


// AwsCognitoIdentityPoolRolesAttachmentMappingRule is a AwsCognitoIdentityPoolRolesAttachmentMappingRule
type AwsCognitoIdentityPoolRolesAttachmentMappingRule struct {
	Claim	string	`json:"claim"`
	MatchType	string	`json:"match_type"`
	RoleArn	string	`json:"role_arn"`
	Value	string	`json:"value"`
}

// AwsCognitoIdentityPoolRolesAttachmentRoleMapping is a AwsCognitoIdentityPoolRolesAttachmentRoleMapping
type AwsCognitoIdentityPoolRolesAttachmentRoleMapping struct {
	AmbiguousRoleResolution	string	`json:"ambiguous_role_resolution"`
	MappingRule	[]AwsCognitoIdentityPoolRolesAttachmentMappingRule	`json:"mapping_rule"`
	Type	string	`json:"type"`
	IdentityProvider	string	`json:"identity_provider"`
}

// AwsCognitoIdentityPoolRolesAttachmentRoles is a AwsCognitoIdentityPoolRolesAttachmentRoles
type AwsCognitoIdentityPoolRolesAttachmentRoles struct {
	Authenticated	string	`json:"authenticated"`
	Unauthenticated	string	`json:"unauthenticated"`
}
