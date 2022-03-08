/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package storage

import (
	"fmt"

	"github.com/pkg/errors"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/config"
)

// ConversionGraphBuilder is used to construct a conversion graph with all the required conversions to/from/between
// the storage variants of the packages. It uses a separate GroupConversionGraphBuilder for each distinct group
type ConversionGraphBuilder struct {
	configuration *config.ObjectModelConfiguration
	versionPrefix string
	subBuilders   map[string]*GroupConversionGraphBuilder
}

// NewConversionGraphBuilder creates a new builder for all our required conversion graphs
func NewConversionGraphBuilder(
	configuration *config.ObjectModelConfiguration,
	versionPrefix string) *ConversionGraphBuilder {
	return &ConversionGraphBuilder{
		configuration: configuration,
		versionPrefix: versionPrefix,
		subBuilders:   make(map[string]*GroupConversionGraphBuilder),
	}
}

// Add includes the supplied package reference in the conversion graph
func (b *ConversionGraphBuilder) Add(refs ...astmodel.PackageReference) {
	for _, ref := range refs {
		subBuilder := b.getSubBuilder(ref)
		subBuilder.Add(ref)
	}
}

// AddAll includes all the supplied package references in the conversion graph
func (b *ConversionGraphBuilder) AddAll(set *astmodel.PackageReferenceSet) {
	for _, ref := range set.AsSlice() {
		b.Add(ref)
	}
}

// Build connects all the provided API definitions together into a single conversion graph
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
		configuration: b.configuration,
		subGraphs:     subgraphs,
	}

	return result, nil
}

// getSubBuilder finds the relevant builder for the group of the provided reference, creating one if necessary
func (b *ConversionGraphBuilder) getSubBuilder(ref astmodel.PackageReference) *GroupConversionGraphBuilder {
	// Expect to get either a local or a storage reference, not an external one
	group, _, ok := ref.GroupVersion()
	if !ok {
		panic(fmt.Sprintf("cannot use external package reference %s with a conversion graph", ref))
	}

	subBuilder, ok := b.subBuilders[group]
	if !ok {
		subBuilder = NewGroupConversionGraphBuilder(group, b.versionPrefix)
		b.subBuilders[group] = subBuilder
	}

	return subBuilder
}
