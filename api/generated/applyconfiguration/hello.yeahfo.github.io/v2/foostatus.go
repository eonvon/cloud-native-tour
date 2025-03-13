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
)

// FooStatusApplyConfiguration represents a declarative configuration of the FooStatus type for use
// with apply.
type FooStatusApplyConfiguration struct {
	Phase      *helloeonvongithubiov2.FooPhase  `json:"phase,omitempty"`
	Conditions []FooConditionApplyConfiguration `json:"conditions,omitempty"`
}

// FooStatusApplyConfiguration constructs a declarative configuration of the FooStatus type for use with
// apply.
func FooStatus() *FooStatusApplyConfiguration {
	return &FooStatusApplyConfiguration{}
}

// WithPhase sets the Phase field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Phase field is set to the value of the last call.
func (b *FooStatusApplyConfiguration) WithPhase(value helloeonvongithubiov2.FooPhase) *FooStatusApplyConfiguration {
	b.Phase = &value
	return b
}

// WithConditions adds the given value to the Conditions field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Conditions field.
func (b *FooStatusApplyConfiguration) WithConditions(values ...*FooConditionApplyConfiguration) *FooStatusApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithConditions")
		}
		b.Conditions = append(b.Conditions, *values[i])
	}
	return b
}
