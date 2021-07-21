/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package storage

import (
	"fmt"

	"github.com/pkg/errors"

	"github.com/Azure/azure-service-operator/hack/generator/pkg/astmodel"
)

// ConversionGraphBuilder is used to construct a conversion graph with all the required conversions to/from/between
// the storage variants of the packages. It uses a separate GroupConversionGraphBuilder for each distinct group
type ConversionGraphBuilder struct {
	subBuilders map[string]*GroupConversionGraphBuilder
}

// NewConversionGraphBuilder creates a new builder for all our required conversion graphs
func NewConversionGraphBuilder() *ConversionGraphBuilder {
	return &ConversionGraphBuilder{
		subBuilders: make(map[string]*GroupConversionGraphBuilder),
	}
}

// Add includes the supplied package reference in the conversion graph
func (b *ConversionGraphBuilder) Add(ref astmodel.PackageReference) {
	subBuilder := b.getSubBuilder(ref)
	subBuilder.Add(ref)
}

// AddAll includes all the supplied package references in the conversion graph
func (b *ConversionGraphBuilder) AddAll(set *astmodel.PackageReferenceSet) {
	for _, ref := range set.AsSlice() {
		b.Add(ref)
	}
}

// Build connects all the provided API types together into a single conversion graph
func (b *ConversionGraphBuilder) Build() (*ConversionGraph, error) {

	subgraphs := make(map[string]*GroupConversionGraph)
	for group, builder := range b.subBuilders {
		subgraph, err := builder.Build()
		if err != nil {
			return nil, errors.Wrapf(err, "building subgraph for group %s", group)
		}

		subgraphs[group] = subgraph
	}

	result := &ConversionGraph{
		subGraphs: subgraphs,
	}

	return result, nil
}

// getSubBuilder finds the relevant builder for the group of the provided reference, creating one if necessary
func (b *ConversionGraphBuilder) getSubBuilder(ref astmodel.PackageReference) *GroupConversionGraphBuilder {
	// Expect to get either a local or a storage reference, not an external one
	local, ok := ref.AsLocalPackage()
	if !ok {
		panic(fmt.Sprintf("cannot use external package reference %s with a conversion graph", ref))
	}

	group := local.Group()
	subBuilder, ok := b.subBuilders[group]
	if !ok {
		subBuilder = NewGroupConversionGraphBuilder(group)
		b.subBuilders[group] = subBuilder
	}

	return subBuilder
}
