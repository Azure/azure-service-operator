/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

const (
	LocalPathPrefix       = "github.com/Azure/k8s-infra/hack/generated/apis/" // TODO: From config?
	genRuntimePathPrefix  = "github.com/Azure/k8s-infra/hack/generated/pkg/genruntime"
	GenRuntimePackageName = "genruntime"
)

var MetaV1PackageReference = MakePackageReference("k8s.io/apimachinery/pkg/apis/meta/v1")

// PackageReference indicates which package a type belongs to
type PackageReference struct {
	packagePath string
}

// MakeLocalPackageReference Creates a new local package reference from a group and package name
func MakeLocalPackageReference(groupName string, packageName string) PackageReference {
	url := LocalPathPrefix + groupName + "/" + packageName
	return PackageReference{packagePath: url}
}

// MakePackageReference creates a new package reference from a path
func MakePackageReference(packagePath string) PackageReference {
	return PackageReference{packagePath: packagePath}
}

// MakeGenRuntimePackageReference creates a new package reference for the genruntime package
func MakeGenRuntimePackageReference() PackageReference {
	url := genRuntimePathPrefix
	return PackageReference{packagePath: url}
}

func (pr PackageReference) IsLocalPackage() bool {
	return strings.HasPrefix(pr.packagePath, LocalPathPrefix)
}

func (pr PackageReference) stripLocalPackagePrefix() (string, error) {
	if !pr.IsLocalPackage() {
		return "", errors.Errorf("cannot strip local package prefix from non-local package %v", pr.packagePath)
	}

	return strings.Replace(pr.packagePath, LocalPathPrefix, "", -1), nil
}

// GroupAndPackage gets the group and package for this package reference if applicable,
// or an error if not
func (pr PackageReference) GroupAndPackage() (string, string, error) {
	groupAndVersion, err := pr.stripLocalPackagePrefix()
	if err != nil {
		return "", "", err
	}

	result := strings.Split(groupAndVersion, "/")
	return result[0], result[1], nil
}

// PackagePath returns the fully qualified package path
func (pr PackageReference) PackagePath() string {
	return pr.packagePath
}

// PackageName is the package name of the package reference
func (pr PackageReference) PackageName() string {
	l := strings.Split(pr.packagePath, "/")
	return l[len(l)-1]
}

// Equals returns true if the passed package reference references the same package, false otherwise
func (pr PackageReference) Equals(ref PackageReference) bool {
	return pr.packagePath == ref.packagePath
}

// String returns the string representation of the package reference
func (pr PackageReference) String() string {
	return pr.packagePath
}

// Ensure we implement Stringer
var _ fmt.Stringer = PackageReference{}
