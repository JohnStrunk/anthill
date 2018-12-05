package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// GlusterStorageTarget defines a storage target
type GlusterStorageTarget struct {
	Name    string      `json:"name"`
	Address []string    `json:"address"`
	Creds   Credentials `json:"credentials"`
}

// GlusterClusterReplicationDetails defines replication details
type GlusterClusterReplicationDetails struct {
	Creds   Credentials            `json:"credentials"`
	Targets []GlusterStorageTarget `json:"targets"`
}

// GlusterNodeThreshold defines threshold details for node
type GlusterNodeThreshold struct {
	Nodes          int    `json:"nodes"`
	MinNodes       int    `json:"minNodes"`
	MaxNodes       int    `json:"maxNodes"`
	FreeStorageMin string `json:"freeStorageMin"`
	FreeStorageMax string `json:"freeStorageMax"`
}

// GlusterNodeStorageDetails defines storage class details
type GlusterNodeStorageDetails struct {
	StorageClassName string `json:"storageClassName"`
	Capacity         string `json:"capacity"`
}

// GlusterNodeTemplate defines a gluster node's template
type GlusterNodeTemplate struct {
	Name       string                    `json:"name"`
	Zone       string                    `json:"zone"`
	Thresholds GlusterNodeThreshold      `json:"thresholds"`
	Affinity   *corev1.NodeAffinity      `json:"nodeAffinity"`
	Storage    GlusterNodeStorageDetails `json:"storage"`
}

// GlusterClusterSpec defines the desired state of GlusterCluster
type GlusterClusterSpec struct {
	Name          string                           `json:"name"`
	Options       map[string]string                `json:"clusterOptions"`
	Drivers       []string                         `json:"drivers"`
	Creds         Credentials                      `json:"glusterCA"`
	Replication   GlusterClusterReplicationDetails `json:"replication"`
	NodeTemplates []GlusterNodeTemplate            `json:"nodeTemplates"`
}

// GlusterClusterStatus defines the observed state of GlusterCluster
type GlusterClusterStatus struct {
	State string `json:"state"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GlusterCluster is the Schema for the glusterclusters API
// +k8s:openapi-gen=true
type GlusterCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GlusterClusterSpec   `json:"spec,omitempty"`
	Status GlusterClusterStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GlusterClusterList contains a list of GlusterCluster
type GlusterClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GlusterCluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&GlusterCluster{}, &GlusterClusterList{})
}
