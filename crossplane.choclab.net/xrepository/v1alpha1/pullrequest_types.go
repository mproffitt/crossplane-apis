// Copyright 2024 The crossplane-apis Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// Repository type metadata.
var (
	PullRequestKind      = "PullRequest"
	PullRequestGroupKind = schema.GroupKind{
		Group: XRDGroup,
		Kind:  PullRequestKind,
	}.String()
	PullRequestKindAPIVersion   = PullRequestKind + "." + GroupVersion.String()
	PullRequestGroupVersionKind = GroupVersion.WithKind(PullRequestKind)
)

// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +genclient
// +genclient:nonNamespaced
//
// +kubebuilder:resource:scope=Cluster,categories=crossplane
// +kubebuilder:subresource:status
// +kubebuilder:resource:shortName=pr
// +crossbuilder:generate:xrd:claimNames:kind=PullRequestClaim,plural=pullrequestclaims
// +crossbuilder:generate:xrd:defaultCompositionRef:name=github-pull-request
type PullRequest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PullRequestSpec   `json:"spec"`
	Status PullRequestStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
type PullRequestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []PullRequest `json:"items"`
}

// Use the `spec` struct to contain parameters you might not want to share
// when nesting XRDs - these will usually be parameters that may be defined
// in a parent.

type PullRequestSpec struct {
	xpv1.ResourceSpec `json:",inline"`

	PullRequestParameters `json:",inline"`
}

type PullRequestParameters struct {
	// The message to start the XRD development with
	// +required
	Message string `json:"message"`
}

type PullRequestStatus struct {
	xpv1.ConditionedStatus `json:",inline"`

	// Message will be output here on the first revision
	// +optional
	Hello string `json:"hello,omitempty"`
}
