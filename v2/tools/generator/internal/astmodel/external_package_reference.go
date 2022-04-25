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

// GroupVersion returns the group and version of this local reference.
func (pr ExternalPackageReference) GroupVersion() (string, string, bool) {
	return "", "", false
}
