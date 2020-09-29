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
	packageImports map[PackageReference]PackageImport
	currentPackage PackageReference

	generatedPackages map[PackageReference]*PackageDefinition
}

// New CodeGenerationContext creates a new immutable code generation context
func NewCodeGenerationContext(
	currentPackage PackageReference,
	packageImports map[PackageImport]struct{},
	generatedPackages map[PackageReference]*PackageDefinition) *CodeGenerationContext {

	packageImportsMap := make(map[PackageReference]PackageImport)
	for imp := range packageImports {
		packageImportsMap[imp.PackageReference] = imp
	}

	return &CodeGenerationContext{
		currentPackage:    currentPackage,
		packageImports:    packageImportsMap,
		generatedPackages: generatedPackages}
}

// CurrentPackage returns the current package being generated
func (codeGenContext *CodeGenerationContext) CurrentPackage() PackageReference {
	return codeGenContext.currentPackage
}

// PackageImports returns the set of package references in the current context
func (codeGenContext *CodeGenerationContext) PackageImports() map[PackageReference]PackageImport {
	// return a copy of the map to ensure immutability
	result := make(map[PackageReference]PackageImport)

	for key, value := range codeGenContext.packageImports {
		result[key] = value
	}
	return result
}

// GetImportedPackageName gets the imported packages name or an error if the package was not imported
func (codeGenContext *CodeGenerationContext) GetImportedPackageName(reference PackageReference) (string, error) {
	packageImport, ok := codeGenContext.packageImports[reference]
	if !ok {
		return "", errors.Errorf("package %s not imported", reference)
	}

	return packageImport.PackageName(), nil
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
