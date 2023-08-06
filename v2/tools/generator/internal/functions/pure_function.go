/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package functions

import (
	"github.com/dave/dst"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

// StandaloneFunction is a simple helper that implements the Function interface. It is intended for use
// for pure functions
type PureFunction struct {
	name             string
	idFactory        astmodel.IdentifierFactory
	asFunc           PureFunctionHandler
	requiredPackages *astmodel.PackageReferenceSet
	referencedTypes  astmodel.TypeNameSet
}

var _ astmodel.Function = &PureFunction{}

// NewPureFunction creates a new pure function
func NewPureFunction(
	name string,
	idFactory astmodel.IdentifierFactory,
	asFunc PureFunctionHandler) *PureFunction {
	return &PureFunction{
		name:             name,
		asFunc:           asFunc,
		requiredPackages: astmodel.NewPackageReferenceSet(),
		referencedTypes:  astmodel.NewTypeNameSet(),
		idFactory:        idFactory,
	}
}

type PureFunctionHandler func(f *PureFunction, codeGenerationContext *astmodel.CodeGenerationContext, methodName string) *dst.FuncDecl

// Name returns the unique name of this function
func (fn *PureFunction) Name() string {
	return fn.name
}

func (fn *PureFunction) IdFactory() astmodel.IdentifierFactory {
	return fn.idFactory
}

// RequiredPackageReferences returns the set of required packages for this function
func (fn *PureFunction) RequiredPackageReferences() *astmodel.PackageReferenceSet {
	return fn.requiredPackages
}

// References returns the set of types to which this function refers.
// SHOULD include any types which this function references but its receiver doesn't.
// SHOULD NOT include the receiver of this function.
func (fn *PureFunction) References() astmodel.TypeNameSet {
	return fn.referencedTypes
}

// AsFunc renders the current instance as a Go abstract syntax tree
func (fn *PureFunction) AsFunc(codeGenerationContext *astmodel.CodeGenerationContext, _ astmodel.TypeName) *dst.FuncDecl {
	return fn.asFunc(fn, codeGenerationContext, fn.name)
}

// Equals checks if this function is equal to the passed in function
func (fn *PureFunction) Equals(f astmodel.Function, _ astmodel.EqualityOverrides) bool {
	typedF, ok := f.(*PureFunction)
	if !ok {
		return false
	}

	// TODO: We're not actually checking function structure here
	// - ensure overrides is used if/when we do so
	return fn.name == typedF.name
}

// AddPackageReference adds one or more required package references
func (fn *PureFunction) AddPackageReference(refs ...astmodel.PackageReference) {
	for _, ref := range refs {
		fn.requiredPackages.AddReference(ref)
	}
}

// AddReferencedTypes adds one or more types that are required
// Their packages are automatically included as well.
func (fn *PureFunction) AddReferencedTypes(types ...astmodel.TypeName) {
	for _, t := range types {
		fn.referencedTypes.Add(t)
		fn.requiredPackages.AddReference(t.PackageReference())
	}
}
