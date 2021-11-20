/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"fmt"
	"strings"

	"github.com/dave/dst"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astbuilder"
)

// PackageImport represents an import of a name from a package
type PackageImport struct {
	packageReference PackageReference
	name             string
}

var _ fmt.Stringer = &PackageImport{}

// NewPackageImport creates a new package import from a reference
func NewPackageImport(packageReference PackageReference) PackageImport {
	return PackageImport{
		packageReference: packageReference,
	}
}

// WithName creates a new package reference with a friendly name
func (pi PackageImport) WithName(name string) PackageImport {
	pi.name = name
	return pi
}

func (pi PackageImport) AsImportSpec() *dst.ImportSpec {
	var name *dst.Ident
	if pi.name != "" {
		name = dst.NewIdent(pi.name)
	}

	return &dst.ImportSpec{
		Name: name,
		Path: astbuilder.StringLiteral(pi.packageReference.PackagePath()),
	}
}

// PackageName is the package name of the package reference
func (pi PackageImport) PackageName() string {
	if pi.HasExplicitName() {
		return pi.name
	}

	return pi.packageReference.PackageName()
}

// HasExplicitName returns true if this package import has an explicitly defined name
func (pi PackageImport) HasExplicitName() bool {
	return pi.name != ""
}

// Equals returns true if the passed package reference references the same package, false otherwise
func (pi PackageImport) Equals(ref PackageImport) bool {
	packagesEqual := pi.packageReference.Equals(ref.packageReference)
	namesEqual := pi.name == ref.name

	return packagesEqual && namesEqual
}

func (pi PackageImport) String() string {
	if len(pi.name) > 0 {
		return fmt.Sprintf("%s %s", pi.name, pi.packageReference)
	}

	return pi.packageReference.String()
}

// TODO: There's an assumption here that this package is a local package, or at least a package that has a format
// TODO: similar to one
// ServiceNameForImport extracts a name for the service for use to disambiguate imports
// E.g. for microsoft.batch/v201700401, extract "batch"
//      for microsoft.storage/v20200101 extract "storage"
//      for microsoft.storsimple.1200 extract "storsimple1200" and so on
func (pi PackageImport) ServiceNameForImport() string {
	pathBits := strings.Split(pi.packageReference.PackagePath(), "/")
	index := len(pathBits) - 1
	if index > 0 {
		index--
	}

	result := strings.Replace(pathBits[index], ".", "", -1)
	return result
}

// Create a versioned name based on the service for use to disambiguate imports
// E.g. for microsoft.batch/v201700401, extract "batchv201700401"
//      for microsoft.storage/v20200101 extract "storagev20200101" and so on
func (pi PackageImport) VersionedNameForImport() string {
	service := pi.ServiceNameForImport()
	return service + pi.packageReference.PackageName()
}
