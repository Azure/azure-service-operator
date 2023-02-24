/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package storage

import (
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
	versionPrefix string,
) *ConversionGraphBuilder {
	return &ConversionGraphBuilder{
		configuration: configuration,
		versionPrefix: versionPrefix,
		subBuilders:   make(map[string]*GroupConversionGraphBuilder),
	}
}

// Add includes the supplied TypeNames in the conversion graph
func (b *ConversionGraphBuilder) Add(names ...astmodel.TypeName) {
	for _, name := range names {
		subBuilder := b.getSubBuilder(name)
		subBuilder.Add(name)
	}
}

// AddAll includes the TypeNames in the supplied set in conversion graph
func (b *ConversionGraphBuilder) AddAll(set astmodel.TypeNameSet) {
	for name := range set {
		b.Add(name)
	}
}

// Build connects all the provided API definitions together into a single conversion graph
func (b *ConversionGraphBuilder) Build() (*ConversionGraph, error) {
	subgraphs := make(map[string]*GroupConversionGraph, len(b.subBuilders))
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
func (b *ConversionGraphBuilder) getSubBuilder(name astmodel.TypeName) *GroupConversionGraphBuilder {
	// Expect to get either a local or a storage reference, not an external one
	group, _ := name.PackageReference.GroupVersion()
	subBuilder, ok := b.subBuilders[group]
	if !ok {
		subBuilder = NewGroupConversionGraphBuilder(group, b.configuration, b.versionPrefix)
		b.subBuilders[group] = subBuilder
	}

	return subBuilder
}
