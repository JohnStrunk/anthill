package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Credentials defines gluster node secret credentials
type Credentials struct {
	SecretName      string `json:"secretName"`
	SecretNamespace string `json:"secreteNamespace"`
}

// GlusterNodeExternal defines external details of gluster nodes
type GlusterNodeExternal struct {
	Address string      `json:"address"`
	Creds   Credentials `json:"credentials"`
}

// StorageDevice defines storage details of gluster nodes
type StorageDevice struct {
	Device  string   `json:"device"`
	PVCName string   `json:"pvcName"`
	Tags    []string `json:"tags"`
}

// GlusterNodeSpec defines the desired state of GlusterNode
type GlusterNodeSpec struct {
        Name         string                 `json:"name"`
        Cluster      string                 `json:"cluster"`
        Zone         string                 `json:"zone"`
	DesiredState string                 `json:"desiredState"`
	ExternalInfo GlusterNodeExternal    `json:"external"`
	Storage      []StorageDevice        `json:"storage"`
	Affinity     *corev1.NodeAffinity   `json:"nodeAffinity"`
}

// GlusterNodeStatus defines the observed state of GlusterNode
type GlusterNodeStatus struct {
	CurrentState string `json:"currentState"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GlusterNode is the Schema for the glusternodes API
// +k8s:openapi-gen=true
type GlusterNode struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GlusterNodeSpec   `json:"spec,omitempty"`
	Status GlusterNodeStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GlusterNodeList contains a list of GlusterNode
type GlusterNodeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GlusterNode `json:"items"`
}

func init() {
	SchemeBuilder.Register(&GlusterNode{}, &GlusterNodeList{})
}
