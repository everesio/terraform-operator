
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsGuarddutyDetector struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsGuarddutyDetectorSpec	`json:"spec"`
}

type AwsGuarddutyDetectorList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsGuarddutyDetector	`json:"items"`
}

type AwsGuarddutyDetectorSpec struct {
	Enable	bool	`json:"enable"`
}
