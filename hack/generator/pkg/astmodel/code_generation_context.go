/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import "fmt"

// CodeGenerationContext stores context about the location code-generation is occurring.
// This is required because some things (such as specific field types) are impacted by the context
// in which the field declaration occurs - for example in a file with two conflicting package references
// a disambiguation must occur and field types must ensure they correctly refer to the disambiguated types
type CodeGenerationContext struct {
	packageImports map[PackageReference]PackageImport
	currentPackage *PackageReference
}

// New CodeGenerationContext creates a new immutable code generation context
func NewCodeGenerationContext(currentPackage *PackageReference, packageImports map[PackageImport]struct{}) *CodeGenerationContext {
	packageImportsMap := make(map[PackageReference]PackageImport)
	for imp := range packageImports {
		packageImportsMap[imp.PackageReference] = imp
	}

	return &CodeGenerationContext{currentPackage: currentPackage, packageImports: packageImportsMap}
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
func (codeGenContext *CodeGenerationContext) GetImportedPackageName(reference *PackageReference) (string, error) {
	packageImport, ok := codeGenContext.packageImports[*reference]
	if !ok {
		return "", fmt.Errorf("package %s not imported", reference)
	}

	return packageImport.PackageName(), nil
}
