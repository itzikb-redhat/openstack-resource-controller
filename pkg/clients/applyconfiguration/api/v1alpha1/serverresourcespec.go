/*
Copyright 2024 The ORC Authors.

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

package v1alpha1

import (
	v1alpha1 "github.com/k-orc/openstack-resource-controller/api/v1alpha1"
)

// ServerResourceSpecApplyConfiguration represents a declarative configuration of the ServerResourceSpec type for use
// with apply.
type ServerResourceSpecApplyConfiguration struct {
	Name   *v1alpha1.OpenStackName `json:"name,omitempty"`
	Image  *v1alpha1.ORCNameRef    `json:"image,omitempty"`
	Flavor *v1alpha1.ORCNameRef    `json:"flavor,omitempty"`
}

// ServerResourceSpecApplyConfiguration constructs a declarative configuration of the ServerResourceSpec type for use with
// apply.
func ServerResourceSpec() *ServerResourceSpecApplyConfiguration {
	return &ServerResourceSpecApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *ServerResourceSpecApplyConfiguration) WithName(value v1alpha1.OpenStackName) *ServerResourceSpecApplyConfiguration {
	b.Name = &value
	return b
}

// WithImage sets the Image field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Image field is set to the value of the last call.
func (b *ServerResourceSpecApplyConfiguration) WithImage(value v1alpha1.ORCNameRef) *ServerResourceSpecApplyConfiguration {
	b.Image = &value
	return b
}

// WithFlavor sets the Flavor field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Flavor field is set to the value of the last call.
func (b *ServerResourceSpecApplyConfiguration) WithFlavor(value v1alpha1.ORCNameRef) *ServerResourceSpecApplyConfiguration {
	b.Flavor = &value
	return b
}
