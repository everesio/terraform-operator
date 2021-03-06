
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsAmiLaunchPermission describes a AwsAmiLaunchPermission resource
type AwsAmiLaunchPermission struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsAmiLaunchPermissionSpec	`json:"spec"`
}


// AwsAmiLaunchPermissionSpec is the spec for a AwsAmiLaunchPermission Resource
type AwsAmiLaunchPermissionSpec struct {
	ImageId	string	`json:"image_id"`
	AccountId	string	`json:"account_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsAmiLaunchPermissionList is a list of AwsAmiLaunchPermission resources
type AwsAmiLaunchPermissionList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsAmiLaunchPermission	`json:"items"`
}

