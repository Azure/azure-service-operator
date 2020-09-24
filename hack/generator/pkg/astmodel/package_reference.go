/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

const (
	genRuntimePathPrefix  = "github.com/Azure/k8s-infra/hack/generated/pkg/genruntime"
	GenRuntimePackageName = "genruntime"
	GroupSuffix           = ".infra.azure.com"
)

var MetaV1PackageReference = MakeLibraryPackageReference("k8s.io/apimachinery/pkg/apis/meta/v1")

type PackageReference interface {
	// IsLocalPackage returns a valud indicating whether this is a local package
	IsLocalPackage() bool
	// Group returns the group of this reference (but only if it is local)
	Group() (string, error)
	// Package returns the package name of this reference
	Package() string
	// PackagePath returns the fully qualified package path
	PackagePath() string
	// Equals returns true if the passed package reference references the same package, false otherwise
	Equals(ref PackageReference) bool
	// String returns the string representation of the package reference
	String() string
}

// MakeGenRuntimePackageReference creates a new package reference for the genruntime package
func MakeGenRuntimePackageReference() PackageReference {
	return MakeLibraryPackageReference(genRuntimePathPrefix)
}
