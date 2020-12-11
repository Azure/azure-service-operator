/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"github.com/pkg/errors"
)

// CodeGenerationContext stores context about the location code-generation is occurring.
// This is required because some things (such as specific field types) are impacted by the context
// in which the field declaration occurs - for example in a file with two conflicting package references
// a disambiguation must occur and field types must ensure they correctly refer to the disambiguated types
type CodeGenerationContext struct {
	packageImports *PackageImportSet
	currentPackage PackageReference
	usedImports    *PackageImportSet

	generatedPackages map[PackageReference]*PackageDefinition
}

// NewCodeGenerationContext creates a new immutable code generation context
func NewCodeGenerationContext(
	currentPackage PackageReference,
	packageImports *PackageImportSet,
	generatedPackages map[PackageReference]*PackageDefinition) *CodeGenerationContext {

	imports := NewPackageImportSet()
	imports.Merge(packageImports)

	return &CodeGenerationContext{
		currentPackage:    currentPackage,
		packageImports:    imports,
		generatedPackages: generatedPackages,
		usedImports:       NewPackageImportSet()}
}

// CurrentPackage returns the current package being generated
func (codeGenContext *CodeGenerationContext) CurrentPackage() PackageReference {
	return codeGenContext.currentPackage
}

// PackageImports returns the set of package imports in the current context
func (codeGenContext *CodeGenerationContext) PackageImports() *PackageImportSet {
	// return a copy to ensure immutability
	result := NewPackageImportSet()
	result.Merge(codeGenContext.packageImports)
	return result
}

// UsedImports returns the set of package imports that have been used by the generated code
func (codeGenContext *CodeGenerationContext) UsedPackageImports() *PackageImportSet {
	// return a copy to ensure immutability
	result := NewPackageImportSet()
	result.Merge(codeGenContext.usedImports)
	return result
}

// GetImportedPackageName gets the imported packages name or an error if the package was not imported
func (codeGenContext *CodeGenerationContext) GetImportedPackageName(reference PackageReference) (string, error) {
	packageImport, ok := codeGenContext.packageImports.ImportFor(reference)
	if !ok {
		return "", errors.Errorf("package %s not imported", reference)
	}

	codeGenContext.usedImports.AddImport(packageImport)
	return packageImport.PackageName(), nil
}

// MustGetImportedPackageName gets the imported packages name or panics if the package was not imported
// Use this when you're absolutely positively sure the package will be there already
func (codeGenContext *CodeGenerationContext) MustGetImportedPackageName(reference PackageReference) string {
	result, err := codeGenContext.GetImportedPackageName(reference)
	if err != nil {
		panic(err)
	}

	return result
}

// GetGeneratedPackage gets a reference to the PackageDefinition referred to by the provided reference
func (codeGenContext *CodeGenerationContext) GetGeneratedPackage(reference PackageReference) (*PackageDefinition, error) {
	// Make sure that we're actually importing that package -- don't want to allow references to things we aren't importing
	_, err := codeGenContext.GetImportedPackageName(reference)
	if !reference.Equals(codeGenContext.currentPackage) && err != nil {
		return nil, err
	}

	packageDef, ok := codeGenContext.generatedPackages[reference]
	if !ok {
		return nil, errors.Errorf("%v not imported", reference)
	}
	return packageDef, nil
}

// GetImportedDefinition looks up a particular type definition in a package referenced by this context
func (codeGenContext *CodeGenerationContext) GetImportedDefinition(typeName TypeName) (TypeDefinition, error) {
	pkg, err := codeGenContext.GetGeneratedPackage(typeName.PackageReference)

	if err != nil {
		return TypeDefinition{}, err
	}

	return pkg.GetDefinition(typeName)
}

// GetTypesInPackage returns the actual definitions from a specific package
func (codeGenContext *CodeGenerationContext) GetTypesInPackage(ref PackageReference) (Types, bool) {
	def, ok := codeGenContext.generatedPackages[ref]
	if !ok {
		// Package reference not found
		return nil, false
	}

	return def.definitions, ok
}

// GetTypesInCurrentPackage returns the actual definitions from a specific package
func (codeGenContext *CodeGenerationContext) GetTypesInCurrentPackage() Types {
	def, ok := codeGenContext.GetTypesInPackage(codeGenContext.currentPackage)
	if !ok {
		panic("Should always have definitions for the current package")
	}

	return def
}
