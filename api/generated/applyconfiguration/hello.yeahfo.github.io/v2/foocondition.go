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
// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v2

import (
	helloeonvongithubiov2 "github.com/eonvon/cloud-native-tour/api/hello.eonvon.github.io/v2"
	v1 "k8s.io/api/core/v1"
)

// FooConditionApplyConfiguration represents a declarative configuration of the FooCondition type for use
// with apply.
type FooConditionApplyConfiguration struct {
	Type   *helloeonvongithubiov2.FooConditionType `json:"type,omitempty"`
	Status *v1.ConditionStatus                     `json:"status,omitempty"`
}

// FooConditionApplyConfiguration constructs a declarative configuration of the FooCondition type for use with
// apply.
func FooCondition() *FooConditionApplyConfiguration {
	return &FooConditionApplyConfiguration{}
}

// WithType sets the Type field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Type field is set to the value of the last call.
func (b *FooConditionApplyConfiguration) WithType(value helloeonvongithubiov2.FooConditionType) *FooConditionApplyConfiguration {
	b.Type = &value
	return b
}

// WithStatus sets the Status field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Status field is set to the value of the last call.
func (b *FooConditionApplyConfiguration) WithStatus(value v1.ConditionStatus) *FooConditionApplyConfiguration {
	b.Status = &value
	return b
}
