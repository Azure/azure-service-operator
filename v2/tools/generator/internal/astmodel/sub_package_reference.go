/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"fmt"
	"path"
)

type SubPackageReference struct {
	parent PackageReference
	name   string
}

var _ PackageReference = SubPackageReference{}
var _ LocalLikePackageReference = SubPackageReference{}

var _ fmt.Stringer = SubPackageReference{}

// MakeSubPackageReference creates a new SubPackageReference, representing a nested but distinct package.
// name is the name of the sub-package.
// parent is the parent package.
func MakeSubPackageReference(
	name string,
	parent PackageReference,
) SubPackageReference {
	return SubPackageReference{
		parent: parent,
		name:   name,
	}
}

// PackageName returns the name of the package.
func (s SubPackageReference) PackageName() string {
	return s.name
}

// PackagePath returns the fully qualified package path.
func (s SubPackageReference) PackagePath() string {
	return path.Join(s.parent.PackagePath(), s.name)
}

// Equals returns true if the passed package reference is a sub-package reference with the same name and an equal
// parent.
func (s SubPackageReference) Equals(ref PackageReference) bool {
	other, ok := ref.(SubPackageReference)
	if !ok {
		return false
	}

	return s.name == other.name && s.parent.Equals(other.parent)
}

// String returns the string representation of the package reference, and implements fmt.Stringer.
func (s SubPackageReference) String() string {
	return path.Join(s.parent.String(), s.name)
}

// IsPreview returns true if the package reference is a preview version.
func (s SubPackageReference) IsPreview() bool {
	return s.parent.IsPreview()
}

// TryGroupVersion returns the group and version of the package reference, if it has them.
// Subpackages have the same group/version as their parent.
func (s SubPackageReference) TryGroupVersion() (string, string, bool) {
	return s.parent.TryGroupVersion()
}

// GroupVersion returns the group and version of the package reference.
// Subpackages have the same group/version as their parent.
func (s SubPackageReference) GroupVersion() (string, string) {
	return s.parent.GroupVersion()
}

// Group returns the group of the package reference.
// Subpackages have the same group as their parent.
func (s SubPackageReference) Group() string {
	return s.parent.Group()
}

// Parent returns the parent package reference.
func (s SubPackageReference) Parent() PackageReference {
	return s.parent
}

func (s SubPackageReference) LocalPathPrefix() string {
	if lpr, ok := s.parent.(LocalLikePackageReference); ok {
		return lpr.LocalPathPrefix()
	}

	panic("SubPackageReference parent is not a LocalLikePackageReference")
}

func (s SubPackageReference) Version() string {
	if lpr, ok := s.parent.(LocalLikePackageReference); ok {
		return lpr.Version()
	}

	panic("SubPackageReference parent is not a LocalLikePackageReference")
}

// ImportAlias returns the import alias to use for this package reference.
func (s SubPackageReference) ImportAlias(style PackageImportStyle) string {
	base := s.parent.ImportAlias(style)
	switch style {
	case VersionOnly:
		return base + s.name[0:1]
	case GroupOnly:
		return base
	case GroupAndVersion:
		return base + s.name[0:1]
	default:
		panic(fmt.Sprintf("didn't expect PackageImportStyle %q", style))
	}
}

// Base returns the parent of this subpackge for DerivedPackageReference.
func (s SubPackageReference) Base() PackageReference {
	return s.parent
}
