/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

// InternalPackageReference describes a package reference that points to a local package (either a storage package
// or a standard one). It can be used to abstract across the exact package type (storage vs local)
type InternalPackageReference interface {
	PackageReference

	// LocalPathPrefix returns the prefix (everything up to the group name)
	LocalPathPrefix() string
	Group() string
	Version() string
	PackageName() string
	PackagePath() string
}
