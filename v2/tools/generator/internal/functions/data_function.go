/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package functions

import (
	"github.com/dave/dst"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

type DataFunction[T any] struct {
	name      string
	idFactory astmodel.IdentifierFactory

	asFunc DataFunctionHandler[T]
	data   T

	requiredPackages *astmodel.PackageReferenceSet
	referencedTypes  astmodel.TypeNameSet
}

var _ astmodel.Function = &DataFunction[string]{}

// NewDataFunction creates a new data function
func NewDataFunction[T any](
	name string,
	data T,
	idFactory astmodel.IdentifierFactory,
	asFunc DataFunctionHandler[T],
	requiredPackages ...astmodel.PackageReference,
) *DataFunction[T] {
	packages := astmodel.NewPackageReferenceSet(requiredPackages...)

	return &DataFunction[T]{
		name:             name,
		requiredPackages: packages,
		data:             data,
		asFunc:           asFunc,
		referencedTypes:  astmodel.NewTypeNameSet(),
		idFactory:        idFactory,
	}
}

type DataFunctionHandler[T any] func(
	f *DataFunction[T],
	codeGenerationContext *astmodel.CodeGenerationContext,
	receiver astmodel.TypeName,
	methodName string,
) (*dst.FuncDecl, error)

// Name returns the unique name of this function
// (You can't have two functions with the same name on the same object or resource)
func (fn *DataFunction[T]) Name() string {
	return fn.name
}

func (fn *DataFunction[T]) IDFactory() astmodel.IdentifierFactory {
	return fn.idFactory
}

// RequiredPackageReferences returns the set of required packages for this function
func (fn *DataFunction[T]) RequiredPackageReferences() *astmodel.PackageReferenceSet {
	return fn.requiredPackages
}

func (fn *DataFunction[T]) Data() T {
	return fn.data
}

// References returns the set of types to which this function refers.
// SHOULD include any types which this function references but its receiver doesn't.
// SHOULD NOT include the receiver of this function.
func (fn *DataFunction[T]) References() astmodel.TypeNameSet {
	return fn.referencedTypes
}

// AsFunc renders the current instance as a Go abstract syntax tree
func (fn *DataFunction[T]) AsFunc(
	codeGenerationContext *astmodel.CodeGenerationContext,
	receiver astmodel.InternalTypeName,
) (*dst.FuncDecl, error) {
	return fn.asFunc(fn, codeGenerationContext, receiver, fn.name)
}

// Equals checks if this function is equal to the passed in function
func (fn *DataFunction[T]) Equals(f astmodel.Function, _ astmodel.EqualityOverrides) bool {
	typedF, ok := f.(*DataFunction[T])
	if !ok {
		return false
	}

	// TODO: We're not actually checking function structure here
	// - ensure overrides is used if/when we do so
	return fn.name == typedF.name
}

// AddPackageReference adds one or more required package references
func (fn *DataFunction[T]) AddPackageReference(refs ...astmodel.PackageReference) {
	for _, ref := range refs {
		fn.requiredPackages.AddReference(ref)
	}
}

// AddReferencedTypes adds one or more types that are required
// Their packages are automatically included as well.
func (fn *DataFunction[T]) AddReferencedTypes(types ...astmodel.TypeName) {
	for _, t := range types {
		fn.referencedTypes.Add(t)
		fn.requiredPackages.AddReference(t.PackageReference())
	}
}
