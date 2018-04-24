
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsIamInstanceProfile struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsIamInstanceProfileSpec	`json:"spec"`
}

type AwsIamInstanceProfileList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsIamInstanceProfile	`json:"items"`
}

type AwsIamInstanceProfileSpec struct {
	NamePrefix	string	`json:"name_prefix"`
	Path	string	`json:"path"`
}