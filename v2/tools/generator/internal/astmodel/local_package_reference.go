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
	localPathPrefix  string
	group            string
	apiVersion       string
	generatorVersion string
}

var (
	_ PackageReference = LocalPackageReference{}
	_ fmt.Stringer     = LocalPackageReference{}
)

const GeneratorVersion string = "v1beta"

// MakeLocalPackageReference Creates a new local package reference from a group and version
func MakeLocalPackageReference(prefix string, group string, versionPrefix string, version string) LocalPackageReference {
	return LocalPackageReference{
		localPathPrefix:  prefix,
		group:            group,
		generatorVersion: versionPrefix,
		apiVersion:       sanitizePackageName(version),
	}
}

// LocalPathPrefix returns the prefix (everything up to the group name)
func (pr LocalPackageReference) LocalPathPrefix() string {
	return pr.localPathPrefix
}

// Group returns the group of this local reference
func (pr LocalPackageReference) Group() string {
	return pr.group
}

// Version returns the version of this local reference
func (pr LocalPackageReference) Version() string {
	return pr.generatorVersion + pr.apiVersion
}

// PackageName returns the package name of this reference
func (pr LocalPackageReference) PackageName() string {
	return pr.Version()
}

// PackagePath returns the fully qualified package path
func (pr LocalPackageReference) PackagePath() string {
	url := pr.localPathPrefix + "/" + pr.group + "/" + pr.PackageName()
	return url
}

// Equals returns true if the passed package reference references the same package, false otherwise
func (pr LocalPackageReference) Equals(ref PackageReference) bool {
	if ref == nil {
		return false
	}

	if other, ok := ref.(LocalPackageReference); ok {
		return pr.localPathPrefix == other.localPathPrefix &&
			pr.generatorVersion == other.generatorVersion &&
			pr.apiVersion == other.apiVersion &&
			pr.group == other.group
	}

	return false
}

// String returns the string representation of the package reference
func (pr LocalPackageReference) String() string {
	return pr.PackagePath()
}

// IsPreview returns true if this package reference is a preview
// We don't check the version prefix (which contains the version of the generator) as that may contain alpha or beta
// even if the ARM version is not preview.
func (pr LocalPackageReference) IsPreview() bool {
	return containsPreviewVersionLabel(strings.ToLower(pr.apiVersion))
}

// WithVersionPrefix returns a new LocalPackageReference with a different version prefix
func (pr LocalPackageReference) WithVersionPrefix(prefix string) LocalPackageReference {
	pr.generatorVersion = prefix
	return pr
}

// HasVersionPrefix returns true if we have the specified version prefix, false otherwise.
func (pr LocalPackageReference) HasVersionPrefix(prefix string) bool {
	return pr.generatorVersion == prefix
}

func (pr LocalPackageReference) GeneratorVersion() string {
	return pr.generatorVersion
}

func (pr LocalPackageReference) ApiVersion() string {
	return pr.apiVersion
}

// ExpandedApiVersion returns the version in 'yyyy-mm-dd(-preview)' form,
// instead of 'yyyymmdd(preview)'.
func (pr LocalPackageReference) ExpandedApiVersion() string {
	av := pr.apiVersion
	result := av[0:4] + "-" + av[4:6] + "-" + av[6:8]
	if len(av) > 8 {
		return result + "-" + av[8:]
	}

	return result
}

// IsLocalPackageReference returns true if the supplied reference is a local one
func IsLocalPackageReference(ref PackageReference) bool {
	_, ok := ref.(LocalPackageReference)
	return ok
}

// TryGroupVersion returns the group and version of this local reference.
func (pr LocalPackageReference) TryGroupVersion() (string, string, bool) {
	return pr.group, pr.Version(), true
}

// GroupVersion returns the group and version of this local reference.
func (pr LocalPackageReference) GroupVersion() (string, string) {
	return pr.group, pr.Version()
}

// sanitizePackageName removes all non-alphanumeric characters and converts to lower case
func sanitizePackageName(input string) string {
	var builder = make([]rune, 0, len(input))

	for _, r := range input {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			builder = append(builder, unicode.ToLower(r))
		}
	}

	return string(builder)
}
