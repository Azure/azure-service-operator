/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"errors"
	"fmt"
	"strings"
)

// LibraryPackageReference indicates a library package that needs to be imported
type LibraryPackageReference struct {
	packagePath string
}

var _ PackageReference = LibraryPackageReference{}
var _ fmt.Stringer = LibraryPackageReference{}

// MakeLibraryPackageReference creates a new package reference from a path
func MakeLibraryPackageReference(packagePath string) LibraryPackageReference {
	return LibraryPackageReference{packagePath: packagePath}
}

// IsLocalPackage returns false to indicate that library packages are not local
func (pr LibraryPackageReference) IsLocalPackage() bool {
	return false
}

// Group returns an error because it's invalid for library packages
func (pr LibraryPackageReference) Group() (string, error) {
	return "", errors.New("Cannot return Group() for a library package")
}

// Package returns the package name of this reference
func (pr LibraryPackageReference) Package() string {
	l := strings.Split(pr.packagePath, "/")
	return l[len(l)-1]
}

// PackagePath returns the fully qualified package path
func (pr LibraryPackageReference) PackagePath() string {
	return pr.packagePath
}

// Equals returns true if the passed package reference references the same package, false otherwise
func (pr LibraryPackageReference) Equals(ref PackageReference) bool {
	if other, ok := ref.(LibraryPackageReference); ok {
		return pr.packagePath == other.packagePath
	}

	return false
}

// String returns the string representation of the package reference
func (pr LibraryPackageReference) String() string {
	return pr.packagePath
}
