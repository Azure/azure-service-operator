/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

// InternalPackageReference describes a package reference that points to a local package (either a storage package
// or a standard one). It can be used to abstract across the exact package type (storage vs local)
type InternalPackageReference interface {
	PackageReference

	// Group returns the group to which this package belongs.
	Group() string

	// FolderPath returns the relative path to this package on disk.
	FolderPath() string

	// LocalPathPrefix returns the prefix (everything up to the group name)
	LocalPathPrefix() string

	// TryGroupVersion returns the group and version of this reference.
	// Returns true if the reference has a group and version, false otherwise.
	TryGroupVersion() (string, string, bool)

	// GroupVersion returns the group and version of this reference
	GroupVersion() (string, string)

	// GroupVersion returns the version of this reference
	Version() string
}
