// Workflow defines a DAG of Tasks with SBS constraints.
package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type TaskRef struct {
	Kind string `json:"kind"`
	Name string `json:"name"`
}

type DAGNode struct {
	ID           string    `json:"id"`
	TaskRef      TaskRef   `json:"taskRef"`
	Dependencies []string  `json:"dependencies"`
	Retries      int       `json:"retries"`
}

type WorkflowSpec struct {
	ClusterRef    ObjectRef    `json:"clusterRef"`
	DAG           DAG          `json:"dag"`
	Execution     ExecutionConfig `json:"execution"`
	Deterministic DeterministicConfig `json:"deterministic"`
}

type ObjectRef struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}

type DAG struct {
	Nodes []DAGNode `json:"nodes"`
}

type ExecutionConfig struct {
	MaxParallel int          `json:"maxParallel"`
	RetryPolicy RetryPolicy `json:"retryPolicy"`
}

type RetryPolicy struct {
	MaxAttempts int   `json:"maxAttempts"`
	BackoffMs   int64 `json:"backoffMs"`
}

type DeterministicConfig struct {
	RequiredTickSync bool `json:"requiredTickSync"`
	Replayable       bool `json:"replayable"`
}

type WorkflowStatus struct {
	Phase       string   `json:"phase"`
	Completed   int      `json:"completed"`
	Failed      int      `json:"failed"`
	Running     int      `json:"running"`
	Conditions  []metav1.Condition `json:"conditions"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced
type Workflow struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec   WorkflowSpec   `json:"spec"`
	Status WorkflowStatus `json:"status"`
}

func (r *Workflow) Hub() {}

// +kubebuilder:object:root=true
type WorkflowList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []Workflow `json:"items"`
}

// Task defines a runnable unit with sandbox constraints.
type TaskSpec struct {
	Binary    BinarySpec    `json:"binary"`
	Sandbox   SandboxSpec   `json:"sandbox"`
	Checkpoint CheckpointSpec `json:"checkpoint"`
}

type BinarySpec struct {
	Image       string `json:"image"`
	Entrypoint  string `json:"entrypoint"`
	SHA256      string `json:"sha256"`
}

type SandboxSpec struct {
	CPULimit       string   `json:"cpuLimit"`
	MemoryLimit    string   `json:"memoryLimit"`
	AllowedSyscalls []string `json:"allowedSyscalls"`
	BlockedSyscalls []string `json:"blockedSyscalls"`
}

type CheckpointSpec struct {
	Enabled       bool  `json:"enabled"`
	IntervalTicks int64 `json:"intervalTicks"`
}

type TaskStatus struct {
	Phase   string `json:"phase"`
	NodeID  string `json:"nodeID"`
	Message string `json:"message"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced
type Task struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec   TaskSpec   `json:"spec"`
	Status TaskStatus `json:"status"`
}

func (r *Task) Hub() {}

// +kubebuilder:object:root=true
type TaskList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []Task `json:"items"`
}

// Policy defines SBS invariant constraints.
type PolicySpec struct {
	Enforcement string         `json:"enforcement"`
	SBS         SBSSpec       `json:"sbs"`
	Network     NetworkPolicy `json:"network"`
}

type SBSSpec struct {
	Invariants      []Invariant     `json:"invariants"`
	AllowedSyscalls []string       `json:"allowedSyscalls"`
	BlockedSyscalls []string       `json:"blockedSyscalls"`
}

type Invariant struct {
	Name       string `json:"name"`
	Expression string `json:"expression"`
	Severity   string `json:"severity"`
}

type NetworkPolicy struct {
	Mode string `json:"mode"`
}

type PolicyStatus struct {
	Phase      string `json:"phase"`
	EnforcedAt int64  `json:"enforcedAt"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced
type Policy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec   PolicySpec   `json:"spec"`
	Status PolicyStatus `json:"status"`
}

func (r *Policy) Hub() {}

// +kubebuilder:object:root=true
type PolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []Policy `json:"items"`
}
