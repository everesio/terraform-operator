
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDirectoryServiceConditionalForwarder describes a AwsDirectoryServiceConditionalForwarder resource
type AwsDirectoryServiceConditionalForwarder struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsDirectoryServiceConditionalForwarderSpec	`json:"spec"`
}


// AwsDirectoryServiceConditionalForwarderSpec is the spec for a AwsDirectoryServiceConditionalForwarder Resource
type AwsDirectoryServiceConditionalForwarderSpec struct {
	DirectoryId	string	`json:"directory_id"`
	DnsIps	[]string	`json:"dns_ips"`
	RemoteDomainName	string	`json:"remote_domain_name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDirectoryServiceConditionalForwarderList is a list of AwsDirectoryServiceConditionalForwarder resources
type AwsDirectoryServiceConditionalForwarderList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsDirectoryServiceConditionalForwarder	`json:"items"`
}

