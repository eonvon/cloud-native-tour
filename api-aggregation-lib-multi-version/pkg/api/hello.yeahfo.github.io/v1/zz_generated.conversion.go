//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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
// Code generated by conversion-gen. DO NOT EDIT.

package v1

import (
	helloyeahfogithubio "github.com/yeahfo/cloud-native-tour/api-aggregation-lib-multi-version/pkg/api/hello.yeahfo.github.io"
	helloyeahfogithubiov1 "github.com/yeahfo/cloud-native-tour/api/hello.yeahfo.github.io/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*helloyeahfogithubiov1.FooList)(nil), (*helloyeahfogithubio.FooList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_FooList_To_helloyeahfogithubio_FooList(a.(*helloyeahfogithubiov1.FooList), b.(*helloyeahfogithubio.FooList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*helloyeahfogithubio.FooList)(nil), (*helloyeahfogithubiov1.FooList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_helloyeahfogithubio_FooList_To_v1_FooList(a.(*helloyeahfogithubio.FooList), b.(*helloyeahfogithubiov1.FooList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*helloyeahfogithubio.FooSpec)(nil), (*helloyeahfogithubiov1.FooSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_helloyeahfogithubio_FooSpec_To_v1_FooSpec(a.(*helloyeahfogithubio.FooSpec), b.(*helloyeahfogithubiov1.FooSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*helloyeahfogithubio.Foo)(nil), (*helloyeahfogithubiov1.Foo)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_helloyeahfogithubio_Foo_To_v1_Foo(a.(*helloyeahfogithubio.Foo), b.(*helloyeahfogithubiov1.Foo), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*helloyeahfogithubiov1.FooSpec)(nil), (*helloyeahfogithubio.FooSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_FooSpec_To_helloyeahfogithubio_FooSpec(a.(*helloyeahfogithubiov1.FooSpec), b.(*helloyeahfogithubio.FooSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*helloyeahfogithubiov1.Foo)(nil), (*helloyeahfogithubio.Foo)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_Foo_To_helloyeahfogithubio_Foo(a.(*helloyeahfogithubiov1.Foo), b.(*helloyeahfogithubio.Foo), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1_Foo_To_helloyeahfogithubio_Foo(in *helloyeahfogithubiov1.Foo, out *helloyeahfogithubio.Foo, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1_FooSpec_To_helloyeahfogithubio_FooSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

func autoConvert_helloyeahfogithubio_Foo_To_v1_Foo(in *helloyeahfogithubio.Foo, out *helloyeahfogithubiov1.Foo, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_helloyeahfogithubio_FooSpec_To_v1_FooSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	// WARNING: in.Status requires manual conversion: does not exist in peer-type
	return nil
}

func autoConvert_v1_FooList_To_helloyeahfogithubio_FooList(in *helloyeahfogithubiov1.FooList, out *helloyeahfogithubio.FooList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]helloyeahfogithubio.Foo, len(*in))
		for i := range *in {
			if err := Convert_v1_Foo_To_helloyeahfogithubio_Foo(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1_FooList_To_helloyeahfogithubio_FooList is an autogenerated conversion function.
func Convert_v1_FooList_To_helloyeahfogithubio_FooList(in *helloyeahfogithubiov1.FooList, out *helloyeahfogithubio.FooList, s conversion.Scope) error {
	return autoConvert_v1_FooList_To_helloyeahfogithubio_FooList(in, out, s)
}

func autoConvert_helloyeahfogithubio_FooList_To_v1_FooList(in *helloyeahfogithubio.FooList, out *helloyeahfogithubiov1.FooList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]helloyeahfogithubiov1.Foo, len(*in))
		for i := range *in {
			if err := Convert_helloyeahfogithubio_Foo_To_v1_Foo(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_helloyeahfogithubio_FooList_To_v1_FooList is an autogenerated conversion function.
func Convert_helloyeahfogithubio_FooList_To_v1_FooList(in *helloyeahfogithubio.FooList, out *helloyeahfogithubiov1.FooList, s conversion.Scope) error {
	return autoConvert_helloyeahfogithubio_FooList_To_v1_FooList(in, out, s)
}

func autoConvert_v1_FooSpec_To_helloyeahfogithubio_FooSpec(in *helloyeahfogithubiov1.FooSpec, out *helloyeahfogithubio.FooSpec, s conversion.Scope) error {
	// WARNING: in.Msg requires manual conversion: does not exist in peer-type
	// WARNING: in.Description requires manual conversion: does not exist in peer-type
	return nil
}

func autoConvert_helloyeahfogithubio_FooSpec_To_v1_FooSpec(in *helloyeahfogithubio.FooSpec, out *helloyeahfogithubiov1.FooSpec, s conversion.Scope) error {
	// WARNING: in.Image requires manual conversion: does not exist in peer-type
	// WARNING: in.Config requires manual conversion: does not exist in peer-type
	return nil
}
