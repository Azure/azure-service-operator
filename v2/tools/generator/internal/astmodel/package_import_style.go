/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

// PackageImportStyle is used to determine how to define a unique package import alias in a given context
type PackageImportStyle string

const (
	VersionOnly     PackageImportStyle = "VersionOnly"     // Only need the version of the package to be unique
	GroupOnly       PackageImportStyle = "GroupOnly"       // Only need the group of the package to be unique
	GroupAndVersion PackageImportStyle = "GroupAndVersion" // Need both group and version of the package to be unique
)
