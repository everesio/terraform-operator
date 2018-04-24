
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDxLag struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsDxLagSpec	`json:"spec"`
}

type AwsDxLagList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsDxLag	`json:"items"`
}

type AwsDxLagSpec struct {
	Name	string	`json:"name"`
	ConnectionsBandwidth	string	`json:"connections_bandwidth"`
	Location	string	`json:"location"`
	ForceDestroy	bool	`json:"force_destroy"`
	Tags	map[string]interface{}	`json:"tags"`
}