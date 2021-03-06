
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsApiGatewayGatewayResponse describes a AwsApiGatewayGatewayResponse resource
type AwsApiGatewayGatewayResponse struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsApiGatewayGatewayResponseSpec	`json:"spec"`
}


// AwsApiGatewayGatewayResponseSpec is the spec for a AwsApiGatewayGatewayResponse Resource
type AwsApiGatewayGatewayResponseSpec struct {
	StatusCode	string	`json:"status_code"`
	ResponseTemplates	map[string]string	`json:"response_templates"`
	ResponseParameters	map[string]string	`json:"response_parameters"`
	RestApiId	string	`json:"rest_api_id"`
	ResponseType	string	`json:"response_type"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsApiGatewayGatewayResponseList is a list of AwsApiGatewayGatewayResponse resources
type AwsApiGatewayGatewayResponseList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsApiGatewayGatewayResponse	`json:"items"`
}

