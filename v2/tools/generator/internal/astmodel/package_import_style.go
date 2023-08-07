/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

// PackageImportStyle is used to determine how to define a unique package import alias in a given context
type PackageImportStyle string

const (
	// VersionOnly is used when using just the version of the package is sufficiently unique
	VersionOnly PackageImportStyle = "VersionOnly"
	// GroupOnly is used when using just the group of the package sufficiently unique
	GroupOnly PackageImportStyle = "GroupOnly"
	// GroupAndVersion is used when both the group and version of the package are required for it to be unique
	GroupAndVersion PackageImportStyle = "GroupAndVersion"
)
