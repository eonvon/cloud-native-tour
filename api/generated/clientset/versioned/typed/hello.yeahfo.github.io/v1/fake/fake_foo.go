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
	helloeonvongithubiov1 "github.com/eonvon/cloud-native-tour/api/generated/applyconfiguration/hello.eonvon.github.io/v1"
	typedhelloeonvongithubiov1 "github.com/eonvon/cloud-native-tour/api/generated/clientset/versioned/typed/hello.eonvon.github.io/v1"
	v1 "github.com/eonvon/cloud-native-tour/api/hello.eonvon.github.io/v1"
	gentype "k8s.io/client-go/gentype"
)

// fakeFoos implements FooInterface
type fakeFoos struct {
	*gentype.FakeClientWithListAndApply[*v1.Foo, *v1.FooList, *helloeonvongithubiov1.FooApplyConfiguration]
	Fake *FakeHelloV1
}

func newFakeFoos(fake *FakeHelloV1, namespace string) typedhelloeonvongithubiov1.FooInterface {
	return &fakeFoos{
		gentype.NewFakeClientWithListAndApply[*v1.Foo, *v1.FooList, *helloeonvongithubiov1.FooApplyConfiguration](
			fake.Fake,
			namespace,
			v1.SchemeGroupVersion.WithResource("foos"),
			v1.SchemeGroupVersion.WithKind("Foo"),
			func() *v1.Foo { return &v1.Foo{} },
			func() *v1.FooList { return &v1.FooList{} },
			func(dst, src *v1.FooList) { dst.ListMeta = src.ListMeta },
			func(list *v1.FooList) []*v1.Foo { return gentype.ToPointerSlice(list.Items) },
			func(list *v1.FooList, items []*v1.Foo) { list.Items = gentype.FromPointerSlice(items) },
		),
		fake,
	}
}
