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

package main

import (
	"github.com/mproffitt/crossplane-apis/crossplane.choclab.net/xrepository/v1alpha1"

	xkcl "github.com/crossplane-contrib/function-kcl/input/v1beta1"
	xapiextv1 "github.com/crossplane/crossplane/apis/apiextensions/v1"
	"github.com/mproffitt/crossbuilder/pkg/generate/composition/build"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type builder struct{}

var (
	Builder          = builder{}
	TemplateBasePath string
)

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
			"owner":    "mproffitt",
			"provider": "github",
		})

	build.SetBasePath(TemplateBasePath)

	// Add pipeline steps here
	c.NewPipelineStep("step-kcl-do-something").
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
					Source: "oci://ghcr.io/mproffitt/kcl-test:0.0.1-e74de4b",
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
