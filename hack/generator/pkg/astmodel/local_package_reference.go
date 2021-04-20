/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"fmt"
	"strings"
	"unicode"
)

// LocalPackageReference specifies a local package name or reference
type LocalPackageReference struct {
	localPathPrefix string
	group           string
	version         string
}

var _ PackageReference = LocalPackageReference{}
var _ fmt.Stringer = LocalPackageReference{}

const generatorVersionPrefix string = "v1alpha1api"

// MakeLocalPackageReference Creates a new local package reference from a group and version
func MakeLocalPackageReference(prefix string, group string, version string) LocalPackageReference {
	return LocalPackageReference{
		localPathPrefix: prefix,
		group:           group,
		version:         version,
	}
}

// CreateLocalPackageNameFromVersion transforms a version string (2018-06-01) into a package
// name (v1alpha1api20180601)
func CreateLocalPackageNameFromVersion(version string) string {
	return generatorVersionPrefix + sanitizePackageName(version)
}

// sanitizePackageName removes all non-alphanum characters and converts to lower case
func sanitizePackageName(input string) string {
	var builder []rune = make([]rune, 0, len(input))

	for _, r := range input {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			builder = append(builder, unicode.ToLower(rune(r)))
		}
	}

	return string(builder)
}

// AsLocalPackage returns this instance and true
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

// PackageName returns the package name of this reference
func (pr LocalPackageReference) PackageName() string {
	return pr.version
}

// PackagePath returns the fully qualified package path
func (pr LocalPackageReference) PackagePath() string {
	url := pr.localPathPrefix + "/" + pr.group + "/" + pr.version
	return url
}

// Equals returns true if the passed package reference references the same package, false otherwise
func (pr LocalPackageReference) Equals(ref PackageReference) bool {
	if ref == nil {
		return false
	}

	if other, ok := ref.AsLocalPackage(); ok {
		return pr.localPathPrefix == other.localPathPrefix &&
			pr.version == other.version &&
			pr.group == other.group
	}

	return false
}

// String returns the string representation of the package reference
func (pr LocalPackageReference) String() string {
	return fmt.Sprintf("local:%s/%s", pr.group, pr.version)
}

// IsPreview returns true if this package reference is a preview
func (pr LocalPackageReference) IsPreview() bool {
	return containsPreviewVersionLabel(strings.ToLower(pr.version))
}

// IsLocalPackageReference returns true if the supplied reference is a local one
func IsLocalPackageReference(ref PackageReference) bool {
	_, ok := ref.(LocalPackageReference)
	return ok
}
