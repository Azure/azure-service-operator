/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"fmt"

	"github.com/Azure/k8s-infra/hack/generator/pkg/astbuilder"
	"github.com/dave/dst"
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

// HasExplicitName() returns true if this package import has an explicitly defined name
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
		return fmt.Sprintf("%v %v", pi.name, pi.packageReference)
	}

	return pi.packageReference.String()
}
