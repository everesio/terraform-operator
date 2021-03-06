
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsLoadBalancerPolicy describes a AwsLoadBalancerPolicy resource
type AwsLoadBalancerPolicy struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsLoadBalancerPolicySpec	`json:"spec"`
}


// AwsLoadBalancerPolicySpec is the spec for a AwsLoadBalancerPolicy Resource
type AwsLoadBalancerPolicySpec struct {
	LoadBalancerName	string	`json:"load_balancer_name"`
	PolicyName	string	`json:"policy_name"`
	PolicyTypeName	string	`json:"policy_type_name"`
	PolicyAttribute	AwsLoadBalancerPolicyPolicyAttribute	`json:"policy_attribute"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsLoadBalancerPolicyList is a list of AwsLoadBalancerPolicy resources
type AwsLoadBalancerPolicyList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsLoadBalancerPolicy	`json:"items"`
}


// AwsLoadBalancerPolicyPolicyAttribute is a AwsLoadBalancerPolicyPolicyAttribute
type AwsLoadBalancerPolicyPolicyAttribute struct {
	Name	string	`json:"name"`
	Value	string	`json:"value"`
}
