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
	groupName   string
	packageName string
}

var _ PackageReference = LocalPackageReference{}
var _ fmt.Stringer = LocalPackageReference{}

// MakeLocalPackageReference Creates a new local package reference from a group and package name
func MakeLocalPackageReference(groupName string, packageName string) LocalPackageReference {
	return LocalPackageReference{groupName: groupName, packageName: packageName}
}

// IsLocalPackage returns true
func (pr LocalPackageReference) IsLocalPackage() bool {
	return true
}

// Group returns the group of this reference
func (pr LocalPackageReference) Group() (string, error) {
	return pr.groupName, nil
}

// Package returns the package name of this reference
func (pr LocalPackageReference) Package() string {
	return pr.packageName
}

// PackagePath returns the fully qualified package path
func (pr LocalPackageReference) PackagePath() string {
	url := LocalPathPrefix + pr.groupName + "/" + pr.packageName
	return url
}

// Equals returns true if the passed package reference references the same package, false otherwise
func (pr LocalPackageReference) Equals(ref PackageReference) bool {
	if other, ok := ref.(LocalPackageReference); ok {
		return pr.packageName == other.packageName &&
			pr.groupName == other.groupName
	}

	return false
}

// String returns the string representation of the package reference
func (pr LocalPackageReference) String() string {
	return pr.PackagePath()
}
