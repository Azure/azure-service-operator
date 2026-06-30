/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"strings"

	"github.com/Azure/azure-service-operator/v2/internal/set"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/config"
)

// partitionID is a unique identifier for a partition of resources in the report.
// It is used to group resources into sections of the report, and to create links between sections of the report.
type partitionID string

const (
	partitionIDPrerelease partitionID = "prerelease" // Prerelease resources haven't yet been released
	partitionIDLatest     partitionID = "latest"     // Latest resources are the most recent recommended versions
	partitionIDOther      partitionID = "other"
	partitionIDReleased   partitionID = "released"
	partitionIDDeprecated partitionID = "deprecated" // Deprecated resources will be removed in a future release
)

var (
	masterIndexPartitions = []partitionID{
		partitionIDPrerelease,
		partitionIDLatest, // Only one of Latest or Released will be populated at any one time
		partitionIDReleased,
	}

	groupIndexPartitions = []partitionID{
		partitionIDPrerelease,
		partitionIDLatest,
		partitionIDOther,
		partitionIDReleased,
		partitionIDDeprecated,
	}
)

// resourceVersionsReportPartition represents a set of resources that have been partitioned into groups based
// on their version.
type resourceVersionsReportPartition map[partitionID]set.Set[resourceVersionsReportItem]

// partitionResources takes a set of resources and partitions them into groups based on their version status.
func partitionResources(
	resources set.Set[resourceVersionsReportItem],
	currentRelease string,
) resourceVersionsReportPartition {
	result := make(resourceVersionsReportPartition, 5)

	available := resources.Copy()

	createPartition := func(
		id partitionID,
		members set.Set[resourceVersionsReportItem],
	) {
		result[id] = members
		available = available.Except(members)
	}

	createPartition(
		partitionIDDeprecated,
		available.Where(result.isDeprecatedResource),
	)

	createPartition(
		partitionIDPrerelease,
		available.Where(result.isUnreleasedResource(currentRelease)),
	)

	recommended := result.findRecommendedVersions(available)

	// If all versions are recommended, we just create a Released partition,
	// otherwise we divide into Latest and Other partitions
	if len(recommended) == len(available) {
		createPartition(
			partitionIDReleased,
			available,
		)
	} else {
		createPartition(
			partitionIDLatest,
			recommended,
		)

		// available now contains only non-recommended versions, so we put them in the Other partition
		createPartition(
			partitionIDOther,
			available,
		)
	}

	return result
}

// isDeprecatedResource returns true if the type definition is for a deprecated resource.
func (resourceVersionsReportPartition) isDeprecatedResource(item resourceVersionsReportItem) bool {
	pkg := item.name.InternalPackageReference()
	grp, ver := pkg.GroupVersion()

	// Handcrafted versions are never deprecated
	// (reusing the regex from config to ensure consistency)
	if config.VersionRegex.MatchString(ver) {
		return false
	}

	// v1api style versions are deprecated unless the group is using VersionMigrationMode Legacy
	if astmodel.VersionMigrationModeForGroup(grp) == astmodel.VersionMigrationModeLegacy {
		return false
	}

	return strings.HasPrefix(ver, astmodel.GeneratorVersion)
}

// isUnreleasedResource returns true if the type definition is for an unreleased resource
func (resourceVersionsReportPartition) isUnreleasedResource(
	currentRelease string,
) func(item resourceVersionsReportItem) bool {
	return func(item resourceVersionsReportItem) bool {
		// if item.supportedFrom == currentRelease {
		// 	return false
		// }

		result := astmodel.ComparePathAndVersion(item.supportedFrom, currentRelease)
		return result >= 0
	}
}

// findRecommendedVersions selects a single version of each resource to recommend.
// Stable versions are preferred over preview.
// Later versions are preferred over earlier.
func (resourceVersionsReportPartition) findRecommendedVersions(
	available set.Set[resourceVersionsReportItem],
) set.Set[resourceVersionsReportItem] {
	index := make(map[string]resourceVersionsReportItem, len(available))

	for _, item := range available.Values() {
		known, ok := index[item.name.Name()]
		if !ok {
			// First time we've seen this resource, use this version
			index[item.name.Name()] = item
			continue
		}

		knownVersion := known.name.InternalPackageReference()
		itemVersion := item.name.InternalPackageReference()

		if knownVersion.IsPreview() && !itemVersion.IsPreview() {
			// Prefer stable versions over preview, so replace the preview version
			index[item.name.Name()] = item
			continue
		}

		if !knownVersion.IsPreview() && itemVersion.IsPreview() {
			// Prefer stable versions over preview, so keep what we have
			continue
		}

		// Both are either preview or stable, so compare versions
		if astmodel.ComparePathAndVersion(itemVersion.Version(), knownVersion.Version()) > 0 {
			// Prefer later versions
			index[item.name.Name()] = item
		}
	}

	result := set.Make[resourceVersionsReportItem]()

	for _, item := range index {
		result.Add(item)
	}

	return result
}
