/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package functions

import (
	"github.com/dave/dst"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

// ObjectFunction is a simple helper that implements the Function interface. It is intended for use for functions
// that only need information about the object they are operating on
type ObjectFunction struct {
	name             string
	idFactory        astmodel.IdentifierFactory
	asFunc           ObjectFunctionHandler
	requiredPackages *astmodel.PackageReferenceSet
	referencedTypes  astmodel.TypeNameSet[astmodel.TypeName]
}

var _ astmodel.Function = &ObjectFunction{}

// NewObjectFunction creates a new object function
func NewObjectFunction(
	name string,
	idFactory astmodel.IdentifierFactory,
	asFunc ObjectFunctionHandler) *ObjectFunction {
	return &ObjectFunction{
		name:             name,
		asFunc:           asFunc,
		requiredPackages: astmodel.NewPackageReferenceSet(),
		referencedTypes:  astmodel.NewTypeNameSet[astmodel.TypeName](),
		idFactory:        idFactory,
	}
}

type ObjectFunctionHandler func(f *ObjectFunction, codeGenerationContext *astmodel.CodeGenerationContext, receiver astmodel.TypeName, methodName string) *dst.FuncDecl

// Name returns the unique name of this function
// (You can't have two functions with the same name on the same object or resource)
func (fn *ObjectFunction) Name() string {
	return fn.name
}

func (fn *ObjectFunction) IdFactory() astmodel.IdentifierFactory {
	return fn.idFactory
}

// RequiredPackageReferences returns the set of required packages for this function
func (fn *ObjectFunction) RequiredPackageReferences() *astmodel.PackageReferenceSet {
	return fn.requiredPackages
}

// References returns the set of types to which this function refers.
// SHOULD include any types which this function references but its receiver doesn't.
// SHOULD NOT include the receiver of this function.
func (fn *ObjectFunction) References() astmodel.TypeNameSet[astmodel.TypeName] {
	return fn.referencedTypes
}

// AsFunc renders the current instance as a Go abstract syntax tree
func (fn *ObjectFunction) AsFunc(codeGenerationContext *astmodel.CodeGenerationContext, receiver astmodel.TypeName) *dst.FuncDecl {
	return fn.asFunc(fn, codeGenerationContext, receiver, fn.name)
}

// Equals checks if this function is equal to the passed in function
func (fn *ObjectFunction) Equals(f astmodel.Function, _ astmodel.EqualityOverrides) bool {
	typedF, ok := f.(*ObjectFunction)
	if !ok {
		return false
	}

	// TODO: We're not actually checking function structure here
	// - ensure overrides is used if/when we do so
	return fn.name == typedF.name
}

// AddPackageReference adds one or more required package references
func (fn *ObjectFunction) AddPackageReference(refs ...astmodel.PackageReference) {
	for _, ref := range refs {
		fn.requiredPackages.AddReference(ref)
	}
}

// AddReferencedTypes adds one or more types that are required
// Their packages are automatically included as well.
func (fn *ObjectFunction) AddReferencedTypes(types ...astmodel.TypeName) {
	for _, t := range types {
		fn.referencedTypes.Add(t)
		fn.requiredPackages.AddReference(t.PackageReference())
	}
}
