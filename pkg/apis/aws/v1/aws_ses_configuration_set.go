
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSesConfigurationSet struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsSesConfigurationSetSpec	`json:"spec"`
}

type AwsSesConfigurationSetList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsSesConfigurationSet	`json:"items"`
}

type AwsSesConfigurationSetSpec struct {
	Name	string	`json:"name"`
}
