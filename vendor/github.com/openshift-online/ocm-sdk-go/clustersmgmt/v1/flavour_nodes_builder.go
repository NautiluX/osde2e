/*
Copyright (c) 2019 Red Hat, Inc.

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

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1 // github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1

// FlavourNodesBuilder contains the data and logic needed to build 'flavour_nodes' objects.
//
// Counts of different classes of nodes inside a flavour.
type FlavourNodesBuilder struct {
	compute *int
	infra   *int
	master  *int
}

// NewFlavourNodes creates a new builder of 'flavour_nodes' objects.
func NewFlavourNodes() *FlavourNodesBuilder {
	return new(FlavourNodesBuilder)
}

// Compute sets the value of the 'compute' attribute
// to the given value.
//
//
func (b *FlavourNodesBuilder) Compute(value int) *FlavourNodesBuilder {
	b.compute = &value
	return b
}

// Infra sets the value of the 'infra' attribute
// to the given value.
//
//
func (b *FlavourNodesBuilder) Infra(value int) *FlavourNodesBuilder {
	b.infra = &value
	return b
}

// Master sets the value of the 'master' attribute
// to the given value.
//
//
func (b *FlavourNodesBuilder) Master(value int) *FlavourNodesBuilder {
	b.master = &value
	return b
}

// Copy copies the attributes of the given object into this builder, discarding any previous values.
func (b *FlavourNodesBuilder) Copy(object *FlavourNodes) *FlavourNodesBuilder {
	if object == nil {
		return b
	}
	b.compute = object.compute
	b.infra = object.infra
	b.master = object.master
	return b
}

// Build creates a 'flavour_nodes' object using the configuration stored in the builder.
func (b *FlavourNodesBuilder) Build() (object *FlavourNodes, err error) {
	object = new(FlavourNodes)
	if b.compute != nil {
		object.compute = b.compute
	}
	if b.infra != nil {
		object.infra = b.infra
	}
	if b.master != nil {
		object.master = b.master
	}
	return
}
