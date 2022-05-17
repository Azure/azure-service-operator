/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package test

import (
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

func CreateFileDefinition(definitions ...astmodel.TypeDefinition) *astmodel.FileDefinition {
	packages := make(map[astmodel.PackageReference]*astmodel.PackageDefinition)

	ref := definitions[0].Name().PackageReference
	group, version := ref.GroupVersion()
	pkgDefinition := astmodel.NewPackageDefinition(group, version)
	for _, def := range definitions {
		pkgDefinition.AddDefinition(def)
	}

	packages[ref] = pkgDefinition

	// put all definitions in one file, regardless.
	// the package reference isn't really used here.
	fileDef := astmodel.NewFileDefinition(ref, definitions, packages)
	return fileDef
}

func CreateTestFileDefinition(definitions ...astmodel.TypeDefinition) *astmodel.TestFileDefinition {
	packages := make(map[astmodel.PackageReference]*astmodel.PackageDefinition)

	// Use the package reference of the first definition for the whole file
	ref := definitions[0].Name().PackageReference

	group, version := ref.GroupVersion()
	pkgDefinition := astmodel.NewPackageDefinition(group, version)
	for _, def := range definitions {
		pkgDefinition.AddDefinition(def)
	}

	packages[ref] = pkgDefinition

	// put all definitions in one file, regardless.
	// the package reference isn't really used here.
	fileDef := astmodel.NewTestFileDefinition(ref, definitions, packages)
	return fileDef
}
