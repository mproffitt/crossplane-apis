package main

import (
	"github.com/mproffitt/crossplane-apis/crossplane.choclab.net/xrepository/v1alpha1"

	xkcl "github.com/crossplane-contrib/function-kcl/input/v1beta1"
	xapiextv1 "github.com/crossplane/crossplane/apis/apiextensions/v1"
	"github.com/mproffitt/crossbuilder/pkg/generate/composition/build"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type builder struct{}

var Builder = builder{}
var TemplateBasePath string

func (b *builder) GetCompositeTypeRef() build.ObjectKindReference {
	return build.ObjectKindReference{
		GroupVersionKind: v1alpha1.PullRequestGroupVersionKind,
		Object:           &v1alpha1.PullRequest{},
	}
}

func (b *builder) Build(c build.CompositionSkeleton) {
	c.WithName("github-pull-request").
		WithMode(xapiextv1.CompositionModePipeline).
		WithLabels(map[string]string{
			// Add labels for uniquely identifying this composition
			"component": "xrepository",
			"provider":  "github",
			"type":      "pull-request",
			"owner":     "choclab",
		})

	build.SetBasePath(TemplateBasePath)

	// Add pipeline steps here
	c.NewPipelineStep("step-kcl-create-pr").
		WithFunctionRef(xapiextv1.FunctionReference{
			Name: "function-kcl",
		}).
		WithInput(build.ObjectKindReference{
			Object: &xkcl.KCLInput{
				TypeMeta: metav1.TypeMeta{
					APIVersion: "krm.kcl.dev/v1alpha1",
					Kind:       "KCLInput",
				},
				Spec: xkcl.RunSpec{
					Source: "oci://ghcr.io/mproffitt/github-pull-request:0.0.1",
				},
			},
		})

	// Add the auto-ready function at the end
	// This ensures the XR is marked ready when all
	//   created MRs are ready
	c.NewPipelineStep("function-auto-ready").
		WithFunctionRef(xapiextv1.FunctionReference{
			Name: "function-auto-ready",
		})

}
