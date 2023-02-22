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

// GroupConversionGraphBuilder is used to construct a conversion graph with all the required conversions to/from/between
// the storage variants of the packages. It uses a separate ResourceConversionGraphBuilder for each distinct group
type GroupConversionGraphBuilder struct {
	group         string // Common group of the resources needing conversions
	configuration *config.ObjectModelConfiguration
	versionPrefix string
	subBuilders   map[string]*ResourceConversionGraphBuilder
}

// NewGroupConversionGraphBuilder creates a new builder for all our required conversion graphs
func NewGroupConversionGraphBuilder(
	group string,
	configuration *config.ObjectModelConfiguration,
	versionPrefix string,
) *GroupConversionGraphBuilder {
	return &GroupConversionGraphBuilder{
		group:         group,
		configuration: configuration,
		versionPrefix: versionPrefix,
		subBuilders:   make(map[string]*ResourceConversionGraphBuilder),
	}
}

// Add includes the supplied type names in the conversion graph
func (b *GroupConversionGraphBuilder) Add(names ...astmodel.TypeName) {
	for _, name := range names {
		subBuilder := b.getSubBuilder(name)
		subBuilder.Add(name)
	}
}

// AddAll includes all the supplied types names in the conversion graph
func (b *GroupConversionGraphBuilder) AddAll(names astmodel.TypeNameSet) {
	for name := range names {
		b.Add(name)
	}
}

// Build connects all the provided API definitions together into a single conversion graph
func (b *GroupConversionGraphBuilder) Build() (*GroupConversionGraph, error) {
	subgraphs := make(map[string]*ResourceConversionGraph, len(b.subBuilders))
	for group, builder := range b.subBuilders {
		subgraph, err := builder.Build()
		if err != nil {
			return nil, errors.Wrapf(err, "building subgraph for group %s", group)
		}

		subgraphs[group] = subgraph
	}

	result := &GroupConversionGraph{
		group:     b.group,
		subGraphs: subgraphs,
	}

	return result, nil
}

// getSubBuilder finds the relevant builder for the group of the provided reference, creating one if necessary
func (b *GroupConversionGraphBuilder) getSubBuilder(name astmodel.TypeName) *ResourceConversionGraphBuilder {
	// Expect to get either a local or a storage reference, not an external one
	n := name.Name()
	subBuilder, ok := b.subBuilders[n]
	if !ok {
		subBuilder = NewResourceConversionGraphBuilder(n, b.versionPrefix)
		b.subBuilders[n] = subBuilder
	}

	return subBuilder
}
