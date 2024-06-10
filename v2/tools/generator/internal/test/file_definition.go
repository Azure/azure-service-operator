/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package test

import (
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

type goSourceFileFactory func(pkg astmodel.InternalPackageReference,
	definitions []astmodel.TypeDefinition,
	packages map[astmodel.InternalPackageReference]*astmodel.PackageDefinition,
) astmodel.GoSourceFile

var _ goSourceFileFactory = createFileDefinition

// createFileDefinition creates a code file containing the passed definitions.
// pkg is the package we're generating
// definitions is the set of type definitions to include in the file.
// packages is a map of all other packages being generated (to allow for cross-package references).
func createFileDefinition(
	pkg astmodel.InternalPackageReference,
	definitions []astmodel.TypeDefinition,
	packages map[astmodel.InternalPackageReference]*astmodel.PackageDefinition,
) astmodel.GoSourceFile {
	defs := make([]astmodel.TypeDefinition, 0, len(definitions))
	for _, def := range definitions {
		if def.Name().InternalPackageReference() == pkg {
			defs = append(defs, def)
		}
	}

	// put all definitions in one file, regardless.
	// the package reference isn't really used here.
	fileDef := astmodel.NewFileDefinition(pkg, defs, packages)
	return fileDef
}

var _ goSourceFileFactory = createTestFileDefinition

// createTestFileDefinition creates a test file containing tests for the passed definitions.
// pkg is the package we're generating
// definitions is the set of type definitions from which we select definitions to include in the file.
// packages is a map of all other packages being generated (to allow for cross-package references).
func createTestFileDefinition(
	pkg astmodel.InternalPackageReference,
	definitions []astmodel.TypeDefinition,
	packages map[astmodel.InternalPackageReference]*astmodel.PackageDefinition,
) astmodel.GoSourceFile {
	defs := make([]astmodel.TypeDefinition, 0, len(definitions))
	for _, def := range definitions {
		if def.Name().InternalPackageReference() == pkg {
			defs = append(defs, def)
		}
	}

	// put all definitions in one file, regardless.
	// the package reference isn't really used here.
	fileDef := astmodel.NewTestFileDefinition(pkg, defs, packages)
	return fileDef
}

func createSetOfPackages(
	definitions []astmodel.TypeDefinition,
) map[astmodel.InternalPackageReference]*astmodel.PackageDefinition {
	// Create a set of package definitions, one for each package required by the supplied definitions
	packages := make(map[astmodel.InternalPackageReference]*astmodel.PackageDefinition, 2)
	for _, def := range definitions {
		pkg := def.Name().InternalPackageReference()
		pkgDef, ok := packages[pkg]
		if !ok {
			// Need to create a new package definition
			pkgDef = astmodel.NewPackageDefinition(pkg)
			packages[pkg] = pkgDef
		}

		pkgDef.AddDefinition(def)
	}

	return packages
}
