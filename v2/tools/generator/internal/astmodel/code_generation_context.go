/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"fmt"

	"github.com/pkg/errors"
)

// CodeGenerationContext stores context about the location code-generation is occurring.
// This is required because some things (such as specific field types) are impacted by the context
// in which the field declaration occurs - for example in a file with two conflicting package references
// a disambiguation must occur and field types must ensure they correctly refer to the disambiguated types
type CodeGenerationContext struct {
	packageImports *PackageImportSet
	currentPackage InternalPackageReference
	usedImports    *PackageImportSet

	generatedPackages map[InternalPackageReference]*PackageDefinition
}

var _ ReadonlyTypeDefinitions = &CodeGenerationContext{}

// NewCodeGenerationContext creates a new immutable code generation context
func NewCodeGenerationContext(
	currentPackage InternalPackageReference,
	packageImports *PackageImportSet,
	generatedPackages map[InternalPackageReference]*PackageDefinition,
) *CodeGenerationContext {
	imports := NewPackageImportSet()
	imports.Merge(packageImports)

	return &CodeGenerationContext{
		currentPackage:    currentPackage,
		packageImports:    imports,
		generatedPackages: generatedPackages,
		usedImports:       NewPackageImportSet(),
	}
}

// CurrentPackage returns the current package being generated
func (ctx *CodeGenerationContext) CurrentPackage() InternalPackageReference {
	return ctx.currentPackage
}

// PackageImports returns the set of package imports in the current context
func (ctx *CodeGenerationContext) PackageImports() *PackageImportSet {
	// return a copy to ensure immutability
	result := NewPackageImportSet()
	result.Merge(ctx.packageImports)
	return result
}

// UsedPackageImports returns the set of package imports that have been used by the generated code
func (ctx *CodeGenerationContext) UsedPackageImports() *PackageImportSet {
	// return a copy to ensure immutability
	result := NewPackageImportSet()
	result.Merge(ctx.usedImports)
	return result
}

// GetImportedPackageName gets the imported packages name or an error if the package was not imported
func (ctx *CodeGenerationContext) GetImportedPackageName(reference PackageReference) (string, error) {
	packageImport, ok := ctx.packageImports.ImportFor(reference)
	if !ok {
		return "", errors.Errorf("package %s not imported", reference)
	}

	ctx.usedImports.AddImport(packageImport)
	return packageImport.PackageName(), nil
}

// MustGetImportedPackageName gets the imported packages name or panics if the package was not imported
// Use this when you're absolutely positively sure the package will be there already
func (ctx *CodeGenerationContext) MustGetImportedPackageName(reference PackageReference) string {
	result, err := ctx.GetImportedPackageName(reference)
	if err != nil {
		panic(err)
	}

	return result
}

// GetGeneratedPackage gets a reference to the PackageDefinition referred to by the provided reference
func (ctx *CodeGenerationContext) GetGeneratedPackage(reference InternalPackageReference) (*PackageDefinition, error) {
	// Make sure that we're actually importing that package -- don't want to allow references to things we aren't importing
	_, err := ctx.GetImportedPackageName(reference)
	if !reference.Equals(ctx.currentPackage) && err != nil {
		return nil, err
	}

	packageDef, ok := ctx.generatedPackages[reference]
	if !ok {
		return nil, errors.Errorf("%s not imported", reference)
	}
	return packageDef, nil
}

// GetDefinition looks up a particular type definition in a package available in this context
func (ctx *CodeGenerationContext) GetDefinition(name InternalTypeName) (TypeDefinition, error) {
	pkg, err := ctx.GetGeneratedPackage(name.InternalPackageReference())
	if err != nil {
		return TypeDefinition{}, err
	}

	return pkg.GetDefinition(name)
}

// MustGetDefinition looks up a particular type definition in a package available in this context.
// If it cannot be found, MustGetDefinition panics.
func (ctx *CodeGenerationContext) MustGetDefinition(name InternalTypeName) TypeDefinition {
	result, err := ctx.GetDefinition(name)
	if err != nil {
		panic(err)
	}

	return result
}

// GetDefinitionsInPackage returns the actual definitions from a specific package
func (ctx *CodeGenerationContext) GetDefinitionsInPackage(packageRef InternalPackageReference) (TypeDefinitionSet, bool) {
	def, ok := ctx.generatedPackages[packageRef]
	if !ok {
		// Package reference not found
		return nil, false
	}

	return def.definitions, ok
}

// GetDefinitionsInCurrentPackage returns the actual definitions from a specific package
func (ctx *CodeGenerationContext) GetDefinitionsInCurrentPackage() TypeDefinitionSet {
	def, ok := ctx.GetDefinitionsInPackage(ctx.currentPackage)
	if !ok {
		msg := fmt.Sprintf("Should always have definitions for the current package %s", ctx.currentPackage)
		panic(msg)
	}

	return def
}

// GetAllReachableDefinitions returns the actual definitions from a specific package
func (ctx *CodeGenerationContext) GetAllReachableDefinitions() TypeDefinitionSet {
	// Since we modify result, we make sure we're working with a copy of the set
	result := ctx.GetDefinitionsInCurrentPackage().Copy()
	for _, pkgImport := range ctx.packageImports.AsSlice() {
		ipr, ok := pkgImport.packageReference.(InternalPackageReference)
		if !ok {
			// Skip non-internal references
			continue
		}
		defs, found := ctx.GetDefinitionsInPackage(ipr)
		if found {
			for k, v := range defs {
				result[k] = v
			}
		}
	}

	return result
}
