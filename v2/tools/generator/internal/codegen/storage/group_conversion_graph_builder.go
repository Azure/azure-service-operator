/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package storage

import (
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

// GroupConversionGraphBuilder is used to construct a group conversion graph with all the required conversions
// to/from/between storage variants of the packages
type GroupConversionGraphBuilder struct {
	group      string                        // Name of the group needing conversions
	references *astmodel.PackageReferenceSet // Set of all API package references that make up this group
}

// NewGroupConversionGraphBuilder creates a new builder for a specific group
func NewGroupConversionGraphBuilder(group string) *GroupConversionGraphBuilder {
	return &GroupConversionGraphBuilder{
		group:      group,
		references: astmodel.NewPackageReferenceSet(),
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

	// Our graph is a map of package references, keyed by the upstream reference
	links := make(map[astmodel.PackageReference]astmodel.PackageReference)

	// Create links for all our API references (based on their storage references)
	// and make a list of the storage references for linking
	var storageReferences []astmodel.PackageReference
	for _, ref := range b.references.AsSlice() {
		if sr, ok := ref.(astmodel.StoragePackageReference); ok {
			// Create a link from the API version to the storage version
			links[sr.Local()] = sr
			// Append to our list of all storage references
			storageReferences = append(storageReferences, sr)
		}
	}

	// Create links between the storage versions
	astmodel.SortPackageReferencesByPathAndVersion(storageReferences)

	// For each Preview API version, the link goes from the associated storage variant to the immediately prior
	// storage variant, no matter whether it's preview or GA.
	for i, ref := range storageReferences {
		if i == 0 || !ref.IsPreview() {
			continue
		}

		links[storageReferences[i]] = storageReferences[i-1]
	}

	// For each GA (non-Preview) API version, the link goes from the associated storage variant to the next GA storage
	// variant, skipping any preview releases in between (if any)
	var gaRelease astmodel.PackageReference
	for index, ref := range storageReferences {
		if index == 0 {
			// Always treat the very earliest version as GA
			gaRelease = storageReferences[0]
			continue
		}

		if !ref.IsPreview() {
			// Found a GA release, link to the prior GA release
			nextGARelease := storageReferences[index]
			links[gaRelease] = nextGARelease
			gaRelease = nextGARelease
		}
	}

	result := &GroupConversionGraph{
		group: b.group,
		links: links,
	}

	return result, nil
}
