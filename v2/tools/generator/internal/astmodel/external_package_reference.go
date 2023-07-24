/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"fmt"
	"strings"
)

// ExternalPackageReference indicates a package to be imported from an external source, such as from GitHub or the go
// standard library.
type ExternalPackageReference struct {
	packagePath string
}

var _ PackageReference = ExternalPackageReference{}
var _ fmt.Stringer = ExternalPackageReference{}

// MakeExternalPackageReference creates a new package reference from a path
func MakeExternalPackageReference(packagePath string) ExternalPackageReference {
	return ExternalPackageReference{packagePath: packagePath}
}

// PackageName returns the package name of this reference
func (pr ExternalPackageReference) PackageName() string {
	l := strings.Split(pr.packagePath, "/")
	return l[len(l)-1]
}

// PackagePath returns the fully qualified package path
func (pr ExternalPackageReference) PackagePath() string {
	return pr.packagePath
}

// ImportPath returns the path to use when importing this package
func (pr ExternalPackageReference) ImportPath() string {
	return pr.packagePath
}

// Equals returns true if the passed package reference references the same package, false otherwise
func (pr ExternalPackageReference) Equals(ref PackageReference) bool {
	if other, ok := ref.(ExternalPackageReference); ok {
		return pr.packagePath == other.packagePath
	}

	return false
}

// IsPreview returns false because external references are never previews
func (pr ExternalPackageReference) IsPreview() bool {
	return false
}

// String returns the string representation of the package reference
func (pr ExternalPackageReference) String() string {
	return pr.packagePath
}

// TryGroupVersion returns the group and version of this external reference.
func (pr ExternalPackageReference) TryGroupVersion() (string, string, bool) {
	return "", "", false
}

// GroupVersion triggers a panic because external references don't have a group or version
func (pr ExternalPackageReference) GroupVersion() (string, string) {
	panic(fmt.Sprintf("external package reference %s doesn't have a group or version", pr))
}

// ImportAlias returns the import alias to use for this package reference
func (pr ExternalPackageReference) ImportAlias(style PackageImportStyle) string {
	msg := fmt.Sprintf("cannot create import alias for external package reference %s", pr.packagePath)
	panic(msg)
}

// Group triggers a panic because external references don't have a group
func (pr ExternalPackageReference) Group() string {
	panic(fmt.Sprintf("external package reference %s doesn't have a group", pr))
}
