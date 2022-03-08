/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package storage

import (
	"fmt"

	"github.com/pkg/errors"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

// GroupConversionGraphBuilder is used to construct a group conversion graph with all the required conversions
// to/from/between storage variants of the packages
type GroupConversionGraphBuilder struct {
	group         string                                                  // Name of the group needing conversions
	versionPrefix string                                                  // Prefix expected on core LocalPackageReferences
	references    *astmodel.PackageReferenceSet                           // Set of all API package references that make up this group
	links         map[astmodel.PackageReference]astmodel.PackageReference // A collection of links that make up the graph
}

// NewGroupConversionGraphBuilder creates a new builder for a specific group
func NewGroupConversionGraphBuilder(group string, versionPrefix string) *GroupConversionGraphBuilder {
	return &GroupConversionGraphBuilder{
		group:         group,
		versionPrefix: versionPrefix,
		references:    astmodel.NewPackageReferenceSet(),
		links:         make(map[astmodel.PackageReference]astmodel.PackageReference),
	}
}

// Add includes the supplied package reference(s) in the conversion graph for this group
func (b *GroupConversionGraphBuilder) Add(refs ...astmodel.PackageReference) {
	for _, ref := range refs {
		b.references.AddReference(ref)
	}
}

// Build connects all the provided API definitions together into a single conversion graph
func (b *GroupConversionGraphBuilder) Build() (*GroupConversionGraph, error) {
	stages := []func([]astmodel.PackageReference){
		b.apiReferencesConvertToStorage,
		b.compatibilityReferencesConvertForward,
		b.previewReferencesConvertBackward,
		b.nonPreviewReferencesConvertForward,
	}

	toProcess := b.references.AsSlice()
	astmodel.SortPackageReferencesByPathAndVersion(toProcess)

	for _, s := range stages {
		s(toProcess)
		toProcess = b.withoutLinkedReferences(toProcess)
	}

	// Expect to have only the hub reference left
	if len(toProcess) != 1 {
		return nil, errors.Errorf(
			"expected to have linked all references in group %q, but have %d left",
			b.group,
			len(toProcess))
	}

	result := &GroupConversionGraph{
		group: b.group,
		links: b.links,
	}

	return result, nil
}

// compatibilityReferencesConvertForward links any compatibility references forward to the following version
func (b *GroupConversionGraphBuilder) compatibilityReferencesConvertForward(refs []astmodel.PackageReference) {
	for i, ref := range refs {
		if !b.isCompatibilityPackage(ref) {
			continue
		}

		// Safe to use i+1 because compatibility references will always be preceded by a reference to the original
		// storage package reference
		b.links[ref] = refs[i+1]
	}
}

// apiReferencesConvertToStorage links each API type to the associated storage package
func (b *GroupConversionGraphBuilder) apiReferencesConvertToStorage(refs []astmodel.PackageReference) {
	for _, ref := range refs {
		if s, ok := ref.(astmodel.StoragePackageReference); ok {
			b.links[s.Local()] = s
		}
	}
}

// previewReferencesConvertBackward links each preview version to the immediately prior version, no matter whether it's
// preview or GA.
func (b *GroupConversionGraphBuilder) previewReferencesConvertBackward(refs []astmodel.PackageReference) {
	for i, ref := range refs {
		if i == 0 || !ref.IsPreview() {
			continue
		}

		b.links[ref] = refs[i-1]
	}
}

// nonPreviewReferencesConvertForward links each version with the immediately following version.
// By the time we run this stage, we should only have non-preview (aka GA) releases left
func (b *GroupConversionGraphBuilder) nonPreviewReferencesConvertForward(refs []astmodel.PackageReference) {
	for i, ref := range refs {
		if i == 0 {
			continue
		}

		prior := refs[i-1]
		b.links[prior] = ref
	}
}

// withoutLinkedReferences returns a new slice of references, omitting any that are already linked into our graph
func (b *GroupConversionGraphBuilder) withoutLinkedReferences(
	refs []astmodel.PackageReference) []astmodel.PackageReference {
	var result []astmodel.PackageReference
	for _, ref := range refs {
		if _, ok := b.links[ref]; !ok {
			result = append(result, ref)
		}
	}

	return result
}

func (b *GroupConversionGraphBuilder) isCompatibilityPackage(ref astmodel.PackageReference) bool {
	switch r := ref.(type) {
	case astmodel.ExternalPackageReference:
		return false
	case astmodel.LocalPackageReference:
		return !r.HasVersionPrefix(b.versionPrefix)
	case astmodel.StoragePackageReference:
		return b.isCompatibilityPackage(r.Local())
	default:
		msg := fmt.Sprintf(
			"unexpected PackageReference implementation %T",
			ref)
		panic(msg)
	}
}
