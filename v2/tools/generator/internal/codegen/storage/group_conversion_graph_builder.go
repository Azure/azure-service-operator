/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package storage

import (
	"github.com/rotisserie/eris"
	"golang.org/x/exp/slices"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/config"
)

// GroupConversionGraphBuilder is used to construct a conversion graph with all the required conversions to/from/between
// the storage variants of the packages. It uses a separate ResourceConversionGraphBuilder for each distinct resource/type
type GroupConversionGraphBuilder struct {
	group         string                                                                  // Common group of the resources needing conversions
	configuration *config.ObjectModelConfiguration                                        // Configuration used to look up renames
	subBuilders   map[string]*ResourceConversionGraphBuilder                              // Nested builders, one for each resource, keyed by resource name
	packages      *astmodel.PackageReferenceSet                                           // Set of all storage packages in this group
	links         map[astmodel.InternalPackageReference]astmodel.InternalPackageReference // A collection of links that make up the graph (storage packages only)
}

// NewGroupConversionGraphBuilder creates a new builder for all our required conversion graphs
func NewGroupConversionGraphBuilder(
	group string,
	configuration *config.ObjectModelConfiguration,
) *GroupConversionGraphBuilder {
	return &GroupConversionGraphBuilder{
		group:         group,
		configuration: configuration,
		subBuilders:   make(map[string]*ResourceConversionGraphBuilder),
		packages:      astmodel.NewPackageReferenceSet(),
		links:         make(map[astmodel.InternalPackageReference]astmodel.InternalPackageReference),
	}
}

// Add includes the supplied type names in the conversion graph
func (b *GroupConversionGraphBuilder) Add(names ...astmodel.InternalTypeName) {
	for _, name := range names {
		subBuilder := b.getSubBuilder(name)
		subBuilder.Add(name)

		b.packages.AddReference(name.InternalPackageReference())
	}
}

// AddAll includes all the supplied types names in the conversion graph
func (b *GroupConversionGraphBuilder) AddAll(names astmodel.InternalTypeNameSet) {
	for name := range names {
		b.Add(name)
	}
}

// Build connects all the provided API definitions together into a single conversion graph
func (b *GroupConversionGraphBuilder) Build() (*GroupConversionGraph, error) {
	subGraphs := make(map[string]*ResourceConversionGraph, len(b.subBuilders))
	for group, builder := range b.subBuilders {
		subgraph, err := builder.Build()
		if err != nil {
			return nil, eris.Wrapf(err, "building subgraph for group %s", group)
		}

		subGraphs[group] = subgraph
	}

	stages := []func([]astmodel.InternalPackageReference){
		b.apiReferencesConvertToStorage,
		b.compatibilityReferencesConvertToOriginalPackage,
		b.previewReferencesConvertBackward,
		b.nonPreviewReferencesConvertForward,
	}

	// order the storage packages by version so that we link them in the correct order
	packages := make([]astmodel.InternalPackageReference, 0, b.packages.Length())
	for ref := range b.packages.All() {
		packages = append(packages, ref.(astmodel.InternalPackageReference))
	}
	slices.SortFunc(
		packages,
		func(left astmodel.InternalPackageReference, right astmodel.InternalPackageReference) int {
			return astmodel.ComparePathAndVersion(left.ImportPath(), right.ImportPath())
		})

	for _, s := range stages {
		s(packages)
		packages = b.withoutLinkedRefs(packages)
	}

	result := &GroupConversionGraph{
		group:         b.group,
		subGraphs:     subGraphs,
		links:         b.links,
		configuration: b.configuration,
	}

	return result, nil
}

// apiReferencesConvertToStorage links each API package to the associated storage package
func (b *GroupConversionGraphBuilder) apiReferencesConvertToStorage(refs []astmodel.InternalPackageReference) {
	for _, ref := range refs {
		if s, ok := ref.(astmodel.DerivedPackageReference); ok {
			b.links[s.Base()] = ref
		}
	}
}

// compatibilityReferencesConvertToOriginalPackage links any compatibility references to the original if present
func (b *GroupConversionGraphBuilder) compatibilityReferencesConvertToOriginalPackage(refs []astmodel.InternalPackageReference) {
	for i, ref := range refs {
		// Last package can't be linked forward
		if i+1 >= len(refs) {
			break
		}

		if !isCompatibilityPackage(ref) {
			continue
		}

		next := refs[i+1]
		if next.HasAPIVersion(ref.APIVersion()) {
			b.links[ref] = next
		}
	}
}

// previewReferencesConvertBackward links each preview version to the immediately prior version, no matter whether it's
// preview or GA.
func (b *GroupConversionGraphBuilder) previewReferencesConvertBackward(refs []astmodel.InternalPackageReference) {
	for i, ref := range refs {
		if i == 0 || !ref.IsPreview() {
			continue
		}

		b.links[ref] = refs[i-1]
	}
}

// nonPreviewReferencesConvertForward links each version with the immediately following version.
// By the time we run this stage, we should only have non-preview (aka GA) releases left
func (b *GroupConversionGraphBuilder) nonPreviewReferencesConvertForward(refs []astmodel.InternalPackageReference) {
	for i, ref := range refs {
		// Links are created from the current index to the next;
		// if we're at the end of the sequence, there's nothing to do.
		if i+1 >= len(refs) {
			break
		}

		b.links[ref] = refs[i+1]
	}
}

// withoutLinkedNames returns a new slice of references, omitting any that are already linked into our graph
func (b *GroupConversionGraphBuilder) withoutLinkedRefs(
	refs []astmodel.InternalPackageReference,
) []astmodel.InternalPackageReference {
	result := make([]astmodel.InternalPackageReference, 0, len(refs))
	for _, ref := range refs {
		if _, ok := b.links[ref]; !ok {
			result = append(result, ref)
		}
	}

	return result
}

// getSubBuilder finds the relevant builder for the resource/type of the provided reference, creating one if necessary
func (b *GroupConversionGraphBuilder) getSubBuilder(name astmodel.TypeName) *ResourceConversionGraphBuilder {
	// Expect to get either a local or a storage reference, not an external one
	n := name.Name()
	subBuilder, ok := b.subBuilders[n]
	if !ok {
		subBuilder = NewResourceConversionGraphBuilder(n)
		b.subBuilders[n] = subBuilder
	}

	return subBuilder
}
