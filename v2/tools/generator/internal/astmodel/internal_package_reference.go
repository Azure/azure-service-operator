/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

// InternalPackageReference describes a package reference that points to a local package (either a storage package
// or a standard one). It can be used to abstract across the exact package type (storage vs local)
type InternalPackageReference interface {
	PackageReference

	// IsPreview returns true if this package reference has a suffix indicating it's a preview
	// release, false otherwise
	IsPreview() bool

	// Group returns the group to which this package belongs.
	Group() string

	// PackagePath returns the fully qualified package path
	PackagePath() string

	// FolderPath returns the relative path to this package on disk.
	FolderPath() string

	// LocalPathPrefix returns the prefix (everything up to the group name)
	LocalPathPrefix() string

	// GroupVersion returns the group and version of this reference
	GroupVersion() (string, string)

	// Version returns the version of this reference
	Version() string
}
