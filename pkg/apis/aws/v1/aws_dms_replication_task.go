
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDmsReplicationTask describes a AwsDmsReplicationTask resource
type AwsDmsReplicationTask struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsDmsReplicationTaskSpec	`json:"spec"`
}


// AwsDmsReplicationTaskSpec is the spec for a AwsDmsReplicationTask Resource
type AwsDmsReplicationTaskSpec struct {
	Tags	map[string]string	`json:"tags"`
	CdcStartTime	string	`json:"cdc_start_time"`
	MigrationType	string	`json:"migration_type"`
	ReplicationTaskId	string	`json:"replication_task_id"`
	ReplicationTaskSettings	string	`json:"replication_task_settings"`
	SourceEndpointArn	string	`json:"source_endpoint_arn"`
	TableMappings	string	`json:"table_mappings"`
	TargetEndpointArn	string	`json:"target_endpoint_arn"`
	ReplicationInstanceArn	string	`json:"replication_instance_arn"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDmsReplicationTaskList is a list of AwsDmsReplicationTask resources
type AwsDmsReplicationTaskList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsDmsReplicationTask	`json:"items"`
}

