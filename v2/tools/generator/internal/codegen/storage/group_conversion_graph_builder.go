/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package storage

import (
	"github.com/pkg/errors"

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
	links := make(map[astmodel.PackageReference]astmodel.PackageReference)

	// Get all original API package references in sorted order by release version
	sortedApiReferences := b.references.AsSlice()
	sortedStorageReferences := make([]astmodel.PackageReference, len(sortedApiReferences))
	astmodel.SortPackageReferencesByPathAndVersion(sortedApiReferences)

	// For each original API reference, create a storage variant reference to match
	for index, ref := range sortedApiReferences {
		localRef, ok := ref.(astmodel.LocalPackageReference)
		if !ok {
			// Shouldn't have any non-local references, if we do, abort
			return nil, errors.Errorf("expected all API references to be local references, but %s was not", ref)
		}

		storageRef := astmodel.MakeStoragePackageReference(localRef)
		sortedStorageReferences[index] = storageRef
		links[localRef] = storageRef
	}

	// For each Preview API version, the link goes from the associated storage variant to the immediately prior
	// storage variant, no matter whether it's preview or GA.
	// We use the API reference to work out whether it's preview or not, because the suffix used for storage
	// variants means IsPreview() always returns true
	for i, ref := range sortedApiReferences {
		if i == 0 || !ref.IsPreview() {
			continue
		}

		links[sortedStorageReferences[i]] = sortedStorageReferences[i-1]
	}

	// For each GA (non-Preview) API version, the link goes from the associated storage variant to the next GA storage
	// variant, skipping any preview releases in between (if any)
	var gaRelease astmodel.PackageReference
	for index, ref := range sortedApiReferences {
		if index == 0 {
			// Always treat the very earliest version as GA
			gaRelease = sortedStorageReferences[index]
			continue
		}

		if !ref.IsPreview() {
			// Found a GA release, link to the prior GA release
			nextGARelease := sortedStorageReferences[index]
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
