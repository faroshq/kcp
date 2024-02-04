/*
Copyright 2022 The KCP Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	corev1alpha1 "github.com/kcp-dev/kcp/sdk/apis/core/v1alpha1"
	conditionsv1alpha1 "github.com/kcp-dev/kcp/sdk/apis/third_party/conditions/apis/conditions/v1alpha1"
	"github.com/kcp-dev/kcp/sdk/apis/third_party/conditions/util/conditions"
)

// WorkspaceRootRequest describes the request for root workspace. It is used to create
// a root workspace for a user.
//
// +crd
// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories=experimental
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`,description="The current phase (e.g. Scheduling, Initializing, Ready, Deleting)"
// +kubebuilder:printcolumn:name="URL",type=string,JSONPath=`.status.URL`,description="URL to access the logical cluster"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
type WorkspaceRootRequest struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ObjectMeta `json:"metadata,omitempty"`
	// +optional
	Spec WorkspaceRootRequestSpec `json:"spec,omitempty"`
	// +optional
	Status WorkspaceRootRequestStatus `json:"status,omitempty"`
}

// WorkspaceRootRequestSpec is the specification of the WorkspaceRootRequest resource.
type WorkspaceRootRequestSpec struct {
	// Path to be used to create
	// +optional
	Path string `json:"path,omitempty"`
}

// WorkspaceRootRequestStatus communicates the observed state of the RootWorkspaceRequest status.
type WorkspaceRootRequestStatus struct {
	// url is the address under which the Kubernetes-cluster-like endpoint
	// can be found. This URL can be used to access the logical cluster with standard Kubernetes
	// client libraries and command line tools.
	//
	// +kubebuilder:format:uri
	URL string `json:"URL,omitempty"`

	// Phase of the logical cluster (Initializing, Ready).
	//
	// +kubebuilder:default=Scheduling
	Phase corev1alpha1.LogicalClusterPhaseType `json:"phase,omitempty"`

	// Current processing state of the LogicalCluster.
	// +optional
	Conditions conditionsv1alpha1.Conditions `json:"conditions,omitempty"`
}

func (in *WorkspaceRootRequest) SetConditions(c conditionsv1alpha1.Conditions) {
	in.Status.Conditions = c
}

func (in *WorkspaceRootRequest) GetConditions() conditionsv1alpha1.Conditions {
	return in.Status.Conditions
}

var _ conditions.Getter = &WorkspaceRootRequest{}
var _ conditions.Setter = &WorkspaceRootRequest{}

// WorkspaceRootRequestList is a list of WorkspaceRootRequest
//
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type WorkspaceRootRequestList struct {
	v1.TypeMeta `json:",inline"`
	v1.ListMeta `json:"metadata"`

	Items []WorkspaceRootRequest `json:"items"`
}
