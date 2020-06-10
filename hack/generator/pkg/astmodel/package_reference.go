/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"fmt"
	"strings"
)

const (
	localPathPrefix = "github.com/Azure/k8s-infra/hack/generator/apis/"
)

// PackageReference indicates which package
// a struct belongs to.
type PackageReference struct {
	packagePath string
}

// NewLocalPackageReference Creates a new local package reference from a group and package name
func NewLocalPackageReference(groupName string, packageName string) *PackageReference {
	url := localPathPrefix + groupName + "/" + packageName
	return &PackageReference{packagePath: url}
}

// NewPackageReference creates a new package reference from a path
func NewPackageReference(packagePath string) *PackageReference {
	return &PackageReference{packagePath: packagePath}
}

func (pr *PackageReference) IsLocalPackage() bool {
	return strings.HasPrefix(pr.packagePath, localPathPrefix)
}

func (pr *PackageReference) stripLocalPackagePrefix() (string, error) {
	if !pr.IsLocalPackage() {
		return "", fmt.Errorf("cannot strip local package prefix from non-local package %v", pr.packagePath)
	}

	return strings.Replace(pr.packagePath, localPathPrefix, "", -1), nil
}

// GroupAndPackage gets the group and package for this package reference if applicable,
// or an error if not
func (pr *PackageReference) GroupAndPackage() (string, string, error) {
	groupAndVersion, err := pr.stripLocalPackagePrefix()
	if err != nil {
		return "", "", err
	}

	result := strings.Split(groupAndVersion, "/")
	return result[0], result[1], nil
}

// PackagePath returns the fully qualified package path
func (pr *PackageReference) PackagePath() string {
	return pr.packagePath
}

// PackageName is the package name of the package reference
func (pr *PackageReference) PackageName() string {
	l := strings.Split(pr.packagePath, "/")
	return l[len(l)-1]
}

// Equals returns true if the passed package reference references the same package, false otherwise
func (pr *PackageReference) Equals(ref *PackageReference) bool {
	return pr.packagePath == ref.packagePath
}

// String returns the string representation of the package reference
func (pr *PackageReference) String() string {
	return pr.packagePath
}

// Ensure we implement Stringer
var _ fmt.Stringer = &PackageReference{}
