/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package storage

import (
	"fmt"
	"sort"

	"github.com/pkg/errors"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

// ResourceConversionGraphBuilder is used to construct a group conversion graph with all the required conversions
// to/from/between storage variants of the packages
type ResourceConversionGraphBuilder struct {
	name          string                                  // Name of the resources needing conversions
	versionPrefix string                                  // Prefix expected on core LocalPackageReferences
	references    astmodel.TypeNameSet                    // Set of all Type Names that make up this group
	links         map[astmodel.TypeName]astmodel.TypeName // A collection of links that make up the graph
}

// NewResourceConversionGraphBuilder creates a new builder for a specific resource/type
func NewResourceConversionGraphBuilder(name string, versionPrefix string) *ResourceConversionGraphBuilder {
	return &ResourceConversionGraphBuilder{
		name:          name,
		versionPrefix: versionPrefix,
		references:    astmodel.NewTypeNameSet(),
		links:         make(map[astmodel.TypeName]astmodel.TypeName),
	}
}

// Add includes the supplied package reference(s) in the conversion graph for this group
func (b *ResourceConversionGraphBuilder) Add(names ...astmodel.TypeName) {
	for _, name := range names {
		b.references.Add(name)
	}
}

// Build connects all the provided API definitions together into a single conversion graph
func (b *ResourceConversionGraphBuilder) Build() (*ResourceConversionGraph, error) {
	stages := []func([]astmodel.TypeName){
		b.apiReferencesConvertToStorage,
		b.compatibilityReferencesConvertForward,
		b.previewReferencesConvertBackward,
		b.nonPreviewReferencesConvertForward,
	}

	toProcess := make([]astmodel.TypeName, 0, len(b.references))
	for name := range b.references {
		toProcess = append(toProcess, name)
	}

	sort.Slice(toProcess, func(i, j int) bool {
		return astmodel.ComparePathAndVersion(
			toProcess[i].PackageReference().ImportPath(),
			toProcess[j].PackageReference().ImportPath())
	})

	for _, s := range stages {
		s(toProcess)
		toProcess = b.withoutLinkedNames(toProcess)
	}

	// Expect to have only the hub reference left
	if len(toProcess) != 1 {
		return nil, errors.Errorf(
			"expected to have linked all references in with name %q, but have %d left",
			b.name,
			len(toProcess))
	}

	result := &ResourceConversionGraph{
		name:  b.name,
		links: b.links,
	}

	return result, nil
}

// compatibilityReferencesConvertForward links any compatibility references forward to the following version
func (b *ResourceConversionGraphBuilder) compatibilityReferencesConvertForward(names []astmodel.TypeName) {
	for i, name := range names {
		if !b.isCompatibilityPackage(name.PackageReference()) {
			continue
		}

		// Safe to use i+1 because compatibility references will always be preceded by a reference to the original
		// storage package reference
		b.links[name] = names[i+1]
	}
}

// apiReferencesConvertToStorage links each API type to the associated storage package
func (b *ResourceConversionGraphBuilder) apiReferencesConvertToStorage(names []astmodel.TypeName) {
	for _, name := range names {
		if s, ok := name.PackageReference().(astmodel.DerivedPackageReference); ok {
			n := name.WithPackageReference(s.Base())
			b.links[n] = name
		}
	}
}

// previewReferencesConvertBackward links each preview version to the immediately prior version, no matter whether it's
// preview or GA.
func (b *ResourceConversionGraphBuilder) previewReferencesConvertBackward(names []astmodel.TypeName) {
	for i, name := range names {
		if i == 0 || !name.PackageReference().IsPreview() {
			continue
		}

		b.links[name] = names[i-1]
	}
}

// nonPreviewReferencesConvertForward links each version with the immediately following version.
// By the time we run this stage, we should only have non-preview (aka GA) releases left
func (b *ResourceConversionGraphBuilder) nonPreviewReferencesConvertForward(names []astmodel.TypeName) {
	for i, name := range names {
		// Links are created from the current index to the next;
		// if we're at the end of the sequence, there's nothing to do.
		if i+1 >= len(names) {
			break
		}

		b.links[name] = names[i+1]
	}
}

// withoutLinkedNames returns a new slice of references, omitting any that are already linked into our graph
func (b *ResourceConversionGraphBuilder) withoutLinkedNames(
	names []astmodel.TypeName,
) []astmodel.TypeName {
	result := make([]astmodel.TypeName, 0, len(names))
	for _, ref := range names {
		if _, ok := b.links[ref]; !ok {
			result = append(result, ref)
		}
	}

	return result
}

func (b *ResourceConversionGraphBuilder) isCompatibilityPackage(ref astmodel.PackageReference) bool {
	switch r := ref.(type) {
	case astmodel.ExternalPackageReference:
		return false
	case astmodel.LocalPackageReference:
		return !r.HasVersionPrefix(b.versionPrefix)
	case astmodel.StoragePackageReference:
		return b.isCompatibilityPackage(r.Local())
	case astmodel.SubPackageReference:
		return b.isCompatibilityPackage(r.Parent())
	default:
		msg := fmt.Sprintf(
			"unexpected PackageReference implementation %T",
			ref)
		panic(msg)
	}
}
