/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package storage

import (
	"fmt"

	"github.com/rotisserie/eris"
	"golang.org/x/exp/slices"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

// ResourceConversionGraphBuilder is used to construct a group conversion graph with all the required conversions
// to/from/between storage variants of the packages
type ResourceConversionGraphBuilder struct {
	name       string                       // Name of the resources needing conversions
	references astmodel.InternalTypeNameSet // Set of all Type Names that make up this group
	links      astmodel.TypeAssociation     // A collection of links that make up the graph
}

// NewResourceConversionGraphBuilder creates a new builder for a specific resource/type
func NewResourceConversionGraphBuilder(name string) *ResourceConversionGraphBuilder {
	return &ResourceConversionGraphBuilder{
		name:       name,
		references: astmodel.NewInternalTypeNameSet(),
		links:      make(astmodel.TypeAssociation),
	}
}

// Add includes the supplied package reference(s) in the conversion graph for this group
func (b *ResourceConversionGraphBuilder) Add(names ...astmodel.InternalTypeName) {
	for _, name := range names {
		b.references.Add(name)
	}
}

// Build connects all the provided API definitions together into a single conversion graph
func (b *ResourceConversionGraphBuilder) Build() (*ResourceConversionGraph, error) {
	stages := []func([]astmodel.InternalTypeName){
		b.apiReferencesConvertToStorage,
		b.legacyV1apiReferencesConvertToNewStyle,
		b.previewReferencesConvertBackward,
		b.nonPreviewReferencesConvertForward,
	}

	toProcess := make([]astmodel.InternalTypeName, 0, len(b.references))
	for name := range b.references {
		toProcess = append(toProcess, name)
	}

	slices.SortFunc(
		toProcess,
		func(i astmodel.InternalTypeName, j astmodel.InternalTypeName) int {
			return astmodel.ComparePathAndVersion(
				i.PackageReference().ImportPath(),
				j.PackageReference().ImportPath())
		})

	for _, s := range stages {
		s(toProcess)
		toProcess = b.withoutLinkedNames(toProcess)
	}

	// Expect to have only the hub reference left
	if len(toProcess) != 1 {
		return nil, eris.Errorf(
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

// legacyV1apiReferencesConvertToNewStyle links any legacy v1api package references to their new-style v equivalent, if present
func (b *ResourceConversionGraphBuilder) legacyV1apiReferencesConvertToNewStyle(names []astmodel.InternalTypeName) {
	allNames := make(map[string]astmodel.InternalTypeName, 0)
	for _, name := range names {
		allNames[name.String()] = name
	}

	for _, name := range names {
		newStyleRef, ok := asNewStylePackageReference(name.InternalPackageReference())
		if !ok {
			continue
		}

		newStyleName := name.WithPackageReference(newStyleRef)
		if _, exists := allNames[newStyleName.String()]; exists {
			b.links[name] = newStyleName
		}
	}
}

// apiReferencesConvertToStorage links each API type to the associated storage package
func (b *ResourceConversionGraphBuilder) apiReferencesConvertToStorage(names []astmodel.InternalTypeName) {
	for _, name := range names {
		if s, ok := name.InternalPackageReference().(astmodel.DerivedPackageReference); ok {
			n := name.WithPackageReference(s.Base())
			b.links[n] = name
		}
	}
}

// previewReferencesConvertBackward links each preview version to the immediately prior version, no matter whether it's
// preview or GA.
func (b *ResourceConversionGraphBuilder) previewReferencesConvertBackward(names []astmodel.InternalTypeName) {
	for i, name := range names {
		if i == 0 || !name.InternalPackageReference().IsPreview() {
			continue
		}

		b.links[name] = names[i-1]
	}
}

// nonPreviewReferencesConvertForward links each version with the immediately following version.
// By the time we run this stage, we should only have non-preview (aka GA) releases left
func (b *ResourceConversionGraphBuilder) nonPreviewReferencesConvertForward(names []astmodel.InternalTypeName) {
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
	names []astmodel.InternalTypeName,
) []astmodel.InternalTypeName {
	result := make([]astmodel.InternalTypeName, 0, len(names))
	for _, ref := range names {
		if _, ok := b.links[ref]; !ok {
			result = append(result, ref)
		}
	}

	return result
}

func isCompatibilityPackage(ref astmodel.PackageReference) bool {
	switch r := ref.(type) {
	case astmodel.ExternalPackageReference:
		return false
	case astmodel.LocalPackageReference:
		return r.HasVersionPrefix("v1api")
	case astmodel.SubPackageReference:
		return isCompatibilityPackage(r.Parent())
	default:
		msg := fmt.Sprintf(
			"unexpected PackageReference implementation %T",
			ref)
		panic(msg)
	}
}

// asNewStylePackageReference returns the new-style package reference for the supplied reference, if it is a legacy
// reference, or false if it isn't.
// Correctly handles nested package references, returning a new nested reference with the new-style reference at the
// leaf if appropriate.
func asNewStylePackageReference(
	ref astmodel.PackageReference,
) (astmodel.InternalPackageReference, bool) {
	switch r := ref.(type) {
	case astmodel.ExternalPackageReference:
		return nil, false

	case astmodel.LocalPackageReference:
		if r.HasVersionPrefix("v1api") {
			return r.WithVersionPrefix("v"), true
		}

		return nil, false

	case astmodel.SubPackageReference:
		newParent, ok := asNewStylePackageReference(r.Parent())
		if !ok {
			return nil, false
		}

		return astmodel.MakeSubPackageReference(r.PackageName(), newParent), true

	default:
		msg := fmt.Sprintf(
			"unexpected PackageReference implementation %T",
			ref)
		panic(msg)
	}
}
