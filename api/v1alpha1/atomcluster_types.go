// Package api — ATOMFederationOS Kubernetes API types.
package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ATOMClusterSpec defines the desired state of ATOMCluster.
type ATOMClusterSpec struct {
	Version      string         `json:"version"`
	ClusterID    string         `json:"clusterID"`
	Nodes        []NodeSpec     `json:"nodes"`
	Deterministic DeterministicConfig `json:"deterministic"`
}

type DeterministicConfig struct {
	TickIntervalMs int64   `json:"tickIntervalMs"`
	QuorumRatio    float64 `json:"quorumRatio"`
	EnableLockstep bool    `json:"enableLockstep"`
}

type NodeSpec struct {
	NodeID string   `json:"nodeID"`
	Host   string   `json:"host"`
	Port   int      `json:"port"`
	Roles  []string `json:"roles"`
}

// ATOMClusterStatus defines the observed state of ATOMCluster.
type ATOMClusterStatus struct {
	Phase       string            `json:"phase"`
	ReadyNodes  int              `json:"readyNodes"`
	Conditions  []metav1.Condition `json:"conditions"`
	LastTick    int64            `json:"lastTick"`
	ClusterHash string           `json:"clusterHash"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
type ATOMCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec   ATOMClusterSpec   `json:"spec"`
	Status ATOMClusterStatus `json:"status"`
}

func (r *ATOMCluster) Hub() {}

// +kubebuilder:object:root=true
type ATOMClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []ATOMCluster `json:"items"`
}
