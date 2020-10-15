/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import "fmt"

const (
	LocalPathPrefix = "github.com/Azure/k8s-infra/hack/generated/apis/" // TODO: From config?
)

// LocalPackageReference specifies a local package name or reference
type LocalPackageReference struct {
	group   string
	version string
}

var _ PackageReference = LocalPackageReference{}
var _ fmt.Stringer = LocalPackageReference{}

// MakeLocalPackageReference Creates a new local package reference from a group and version
func MakeLocalPackageReference(group string, version string) LocalPackageReference {
	return LocalPackageReference{group: group, version: version}
}

// IsLocalPackage returns true
func (pr LocalPackageReference) AsLocalPackage() (LocalPackageReference, bool) {
	return pr, true
}

// Group returns the group of this local reference
func (pr LocalPackageReference) Group() string {
	return pr.group
}

// Version returns the version of this local reference
func (pr LocalPackageReference) Version() string {
	return pr.version
}

// Package returns the package name of this reference
func (pr LocalPackageReference) PackageName() string {
	return pr.version
}

// PackagePath returns the fully qualified package path
func (pr LocalPackageReference) PackagePath() string {
	url := LocalPathPrefix + pr.group + "/" + pr.version
	return url
}

// Equals returns true if the passed package reference references the same package, false otherwise
func (pr LocalPackageReference) Equals(ref PackageReference) bool {
	if other, ok := ref.(LocalPackageReference); ok {
		return pr.version == other.version &&
			pr.group == other.group
	}

	return false
}

// String returns the string representation of the package reference
func (pr LocalPackageReference) String() string {
	return pr.PackagePath()
}

// IsLocalPackageReference() returns true if the supplied reference is a local one
func IsLocalPackageReference(ref PackageReference) bool {
	_, ok := ref.(LocalPackageReference)
	return ok
}
