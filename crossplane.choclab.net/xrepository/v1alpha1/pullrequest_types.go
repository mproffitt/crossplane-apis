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
// +crossbuilder:generate:xrd:claimNames:kind=PullRequestClaim,plural=PullRequestClaims
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
}

type PullRequestStatus struct {
	xpv1.ConditionedStatus `json:",inline"`

	// Example status to get things going
	//
	// +optional
	Bob string `json:"bob,omitempty"`
}
