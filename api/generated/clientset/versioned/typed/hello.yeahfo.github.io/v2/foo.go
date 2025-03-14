/*
Copyright The Kubernetes Authors.

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
// Code generated by client-gen. DO NOT EDIT.

package v2

import (
	context "context"

	applyconfigurationhelloeonvongithubiov2 "github.com/eonvon/cloud-native-tour/api/generated/applyconfiguration/hello.eonvon.github.io/v2"
	scheme "github.com/eonvon/cloud-native-tour/api/generated/clientset/versioned/scheme"
	helloeonvongithubiov2 "github.com/eonvon/cloud-native-tour/api/hello.eonvon.github.io/v2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// FoosGetter has a method to return a FooInterface.
// A group's client should implement this interface.
type FoosGetter interface {
	Foos(namespace string) FooInterface
}

// FooInterface has methods to work with Foo resources.
type FooInterface interface {
	Create(ctx context.Context, foo *helloeonvongithubiov2.Foo, opts v1.CreateOptions) (*helloeonvongithubiov2.Foo, error)
	Update(ctx context.Context, foo *helloeonvongithubiov2.Foo, opts v1.UpdateOptions) (*helloeonvongithubiov2.Foo, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, foo *helloeonvongithubiov2.Foo, opts v1.UpdateOptions) (*helloeonvongithubiov2.Foo, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*helloeonvongithubiov2.Foo, error)
	List(ctx context.Context, opts v1.ListOptions) (*helloeonvongithubiov2.FooList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *helloeonvongithubiov2.Foo, err error)
	Apply(ctx context.Context, foo *applyconfigurationhelloeonvongithubiov2.FooApplyConfiguration, opts v1.ApplyOptions) (result *helloeonvongithubiov2.Foo, err error)
	// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
	ApplyStatus(ctx context.Context, foo *applyconfigurationhelloeonvongithubiov2.FooApplyConfiguration, opts v1.ApplyOptions) (result *helloeonvongithubiov2.Foo, err error)
	FooExpansion
}

// foos implements FooInterface
type foos struct {
	*gentype.ClientWithListAndApply[*helloeonvongithubiov2.Foo, *helloeonvongithubiov2.FooList, *applyconfigurationhelloeonvongithubiov2.FooApplyConfiguration]
}

// newFoos returns a Foos
func newFoos(c *HelloV2Client, namespace string) *foos {
	return &foos{
		gentype.NewClientWithListAndApply[*helloeonvongithubiov2.Foo, *helloeonvongithubiov2.FooList, *applyconfigurationhelloeonvongithubiov2.FooApplyConfiguration](
			"foos",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *helloeonvongithubiov2.Foo { return &helloeonvongithubiov2.Foo{} },
			func() *helloeonvongithubiov2.FooList { return &helloeonvongithubiov2.FooList{} },
			gentype.PrefersProtobuf[*helloeonvongithubiov2.Foo](),
		),
	}
}
