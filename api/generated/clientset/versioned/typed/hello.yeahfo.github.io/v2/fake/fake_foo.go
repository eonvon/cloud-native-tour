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

package fake

import (
	helloyeahfogithubiov2 "github.com/yeahfo/cloud-native-tour/api/generated/applyconfiguration/hello.yeahfo.github.io/v2"
	typedhelloyeahfogithubiov2 "github.com/yeahfo/cloud-native-tour/api/generated/clientset/versioned/typed/hello.yeahfo.github.io/v2"
	v2 "github.com/yeahfo/cloud-native-tour/api/hello.yeahfo.github.io/v2"
	gentype "k8s.io/client-go/gentype"
)

// fakeFoos implements FooInterface
type fakeFoos struct {
	*gentype.FakeClientWithListAndApply[*v2.Foo, *v2.FooList, *helloyeahfogithubiov2.FooApplyConfiguration]
	Fake *FakeHelloV2
}

func newFakeFoos(fake *FakeHelloV2, namespace string) typedhelloyeahfogithubiov2.FooInterface {
	return &fakeFoos{
		gentype.NewFakeClientWithListAndApply[*v2.Foo, *v2.FooList, *helloyeahfogithubiov2.FooApplyConfiguration](
			fake.Fake,
			namespace,
			v2.SchemeGroupVersion.WithResource("foos"),
			v2.SchemeGroupVersion.WithKind("Foo"),
			func() *v2.Foo { return &v2.Foo{} },
			func() *v2.FooList { return &v2.FooList{} },
			func(dst, src *v2.FooList) { dst.ListMeta = src.ListMeta },
			func(list *v2.FooList) []*v2.Foo { return gentype.ToPointerSlice(list.Items) },
			func(list *v2.FooList, items []*v2.Foo) { list.Items = gentype.FromPointerSlice(items) },
		),
		fake,
	}
}
