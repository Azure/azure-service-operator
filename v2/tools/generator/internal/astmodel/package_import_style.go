/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

type PackageImportStyle string

const (
	VersionOnly     PackageImportStyle = "VersionOnly"
	GroupOnly       PackageImportStyle = "GroupOnly"
	GroupAndVersion PackageImportStyle = "GroupAndVersion"
)
