/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"fmt"
	"path"
	"strings"
	"unicode"
)

// LocalPackageReference specifies a local package name or reference
type LocalPackageReference struct {
	localPathPrefix  string
	group            string
	apiVersion       string
	generatorVersion string
	// version == apiVersion + generatorVersion
	// cached to save on allocations
	version string
}

var (
	_ PackageReference         = LocalPackageReference{}
	_ fmt.Stringer             = LocalPackageReference{}
	_ InternalPackageReference = LocalPackageReference{}
)

const GeneratorVersion string = "v1api"

// MakeLocalPackageReference Creates a new local package reference from a group and version
func MakeLocalPackageReference(prefix string, group string, versionPrefix string, version string) LocalPackageReference {
	return LocalPackageReference{
		localPathPrefix:  prefix,
		group:            group,
		generatorVersion: versionPrefix,
		apiVersion:       version,
		version:          versionPrefix + sanitizePackageName(version),
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
	return pr.version
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

// ImportPath returns the path to use when importing this package
func (pr LocalPackageReference) ImportPath() string {
	return path.Join(pr.localPathPrefix, pr.group, pr.version)
}

// FolderPath returns the relative path to this package on disk.
func (pr LocalPackageReference) FolderPath() string {
	return path.Join(pr.group, pr.version)
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
	return ContainsPreviewVersionLabel(strings.ToLower(pr.apiVersion))
}

// WithVersionPrefix returns a new LocalPackageReference with a different version prefix
func (pr LocalPackageReference) WithVersionPrefix(prefix string) LocalPackageReference {
	pr.generatorVersion = prefix
	pr.version = prefix + sanitizePackageName(pr.apiVersion)
	return pr
}

// HasVersionPrefix returns true if we have the specified version prefix, false otherwise.
func (pr LocalPackageReference) HasVersionPrefix(prefix string) bool {
	return pr.generatorVersion == prefix
}

// GeneratorVersion returns the part of the package name refering to the version of the generator
func (pr LocalPackageReference) GeneratorVersion() string {
	return pr.generatorVersion
}

// ApiVersion returns the API version of this reference, separate from the generator version
func (pr LocalPackageReference) ApiVersion() string {
	return pr.apiVersion
}

// HasApiVersion returns true if this reference has the specified API version
func (pr LocalPackageReference) HasApiVersion(ver string) bool {
	return strings.EqualFold(pr.apiVersion, ver)
}

// IsLocalPackageReference returns true if the supplied reference is a local one
func IsLocalPackageReference(ref PackageReference) bool {
	_, ok := ref.(LocalPackageReference)
	return ok
}

// GroupVersion returns the group and version of this local reference.
func (pr LocalPackageReference) GroupVersion() (string, string) {
	return pr.group, pr.Version()
}

// ImportAlias returns the import alias to use for this package reference
func (pr LocalPackageReference) ImportAlias(style PackageImportStyle) string {
	switch style {
	case VersionOnly:
		return fmt.Sprintf(
			"%s%s",
			pr.simplifiedGeneratorVersion(pr.generatorVersion),
			pr.simplifiedApiVersion(pr.apiVersion))
	case GroupOnly:
		return pr.group
	case GroupAndVersion:
		return fmt.Sprintf(
			"%s_%s%s",
			pr.group,
			pr.simplifiedGeneratorVersion(pr.generatorVersion),
			pr.simplifiedApiVersion(pr.apiVersion))
	default:
		panic(fmt.Sprintf("didn't expect PackageImportStyle %q", style))
	}
}

var apiVersionSimplifier = strings.NewReplacer(
	"alpha", "a",
	"beta", "b",
	"preview", "p",
	"-", "",
)

func (pr LocalPackageReference) simplifiedApiVersion(version string) string {
	return strings.ToLower(apiVersionSimplifier.Replace(version))
}

var generatorVersionSimplifier = strings.NewReplacer(
	"v1alpha1api", "alpha",
	"v1beta1api", "beta",
	"v1api", "v",
)

func (pr LocalPackageReference) simplifiedGeneratorVersion(version string) string {
	return generatorVersionSimplifier.Replace(version)
}

// sanitizePackageName removes all non-alphanumeric characters and converts to lower case
func sanitizePackageName(input string) string {
	builder := make([]rune, 0, len(input))

	for _, r := range input {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			builder = append(builder, unicode.ToLower(r))
		}
	}

	return string(builder)
}
