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
	v1 "github.com/yeahfo/cloud-native-tour/kube-aggregator/pkg/apis/apiregistration/v1"
	apiregistrationv1 "github.com/yeahfo/cloud-native-tour/kube-aggregator/pkg/client/clientset_generated/clientset/typed/apiregistration/v1"
	gentype "k8s.io/client-go/gentype"
)

// fakeAPIServices implements APIServiceInterface
type fakeAPIServices struct {
	*gentype.FakeClientWithList[*v1.APIService, *v1.APIServiceList]
	Fake *FakeApiregistrationV1
}

func newFakeAPIServices(fake *FakeApiregistrationV1) apiregistrationv1.APIServiceInterface {
	return &fakeAPIServices{
		gentype.NewFakeClientWithList[*v1.APIService, *v1.APIServiceList](
			fake.Fake,
			"",
			v1.SchemeGroupVersion.WithResource("apiservices"),
			v1.SchemeGroupVersion.WithKind("APIService"),
			func() *v1.APIService { return &v1.APIService{} },
			func() *v1.APIServiceList { return &v1.APIServiceList{} },
			func(dst, src *v1.APIServiceList) { dst.ListMeta = src.ListMeta },
			func(list *v1.APIServiceList) []*v1.APIService { return gentype.ToPointerSlice(list.Items) },
			func(list *v1.APIServiceList, items []*v1.APIService) { list.Items = gentype.FromPointerSlice(items) },
		),
		fake,
	}
}
