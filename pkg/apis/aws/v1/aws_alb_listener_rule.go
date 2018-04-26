
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsAlbListenerRule describes a AwsAlbListenerRule resource
type AwsAlbListenerRule struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsAlbListenerRuleSpec	`json:"spec"`
}


// AwsAlbListenerRuleSpec is the spec for a AwsAlbListenerRule Resource
type AwsAlbListenerRuleSpec struct {
	ListenerArn	string	`json:"listener_arn"`
	Action	[]Action	`json:"action"`
	Condition	Condition	`json:"condition"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsAlbListenerRuleList is a list of AwsAlbListenerRule resources
type AwsAlbListenerRuleList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsAlbListenerRule	`json:"items"`
}


// Action is a Action
type Action struct {
	TargetGroupArn	string	`json:"target_group_arn"`
	Type	string	`json:"type"`
}

// Condition is a Condition
type Condition struct {
	Values	[]string	`json:"values"`
	Field	string	`json:"field"`
}
