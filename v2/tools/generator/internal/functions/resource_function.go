/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package functions

import (
	"github.com/dave/dst"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

type ResourceFunctionHandler func(f *ResourceFunction, codeGenerationContext *astmodel.CodeGenerationContext, receiver astmodel.InternalTypeName, methodName string) *dst.FuncDecl

// ResourceFunction is a simple helper that implements the Function interface. It is intended for use for functions
// that only need information about the resource they are operating on
type ResourceFunction struct {
	name             string
	resource         *astmodel.ResourceType
	idFactory        astmodel.IdentifierFactory
	asFunc           ResourceFunctionHandler
	requiredPackages *astmodel.PackageReferenceSet
}

var _ astmodel.Function = &ResourceFunction{}

// NewResourceFunction creates a new resource function
func NewResourceFunction(
	name string,
	resource *astmodel.ResourceType,
	idFactory astmodel.IdentifierFactory,
	asFunc ResourceFunctionHandler,
	requiredPackages *astmodel.PackageReferenceSet) *ResourceFunction {
	return &ResourceFunction{
		name:             name,
		resource:         resource,
		idFactory:        idFactory,
		asFunc:           asFunc,
		requiredPackages: requiredPackages,
	}
}

// Name returns the unique name of this function
// (You can't have two functions with the same name on the same object or resource)
func (fn *ResourceFunction) Name() string {
	return fn.name
}

func (fn *ResourceFunction) IdFactory() astmodel.IdentifierFactory {
	return fn.idFactory
}

// RequiredPackageReferences returns the set of required packages for this function
func (fn *ResourceFunction) RequiredPackageReferences() *astmodel.PackageReferenceSet {
	return fn.requiredPackages
}

// References returns the set of types to which this function refers.
// SHOULD include any types which this function references but its receiver doesn't.
// SHOULD NOT include the receiver of this function.
func (fn *ResourceFunction) References() astmodel.TypeNameSet[astmodel.TypeName] {
	return nil
}

// Resource returns the resource this function applies to
func (fn *ResourceFunction) Resource() *astmodel.ResourceType {
	return fn.resource
}

// AsFunc renders the current instance as a Go abstract syntax tree
func (fn *ResourceFunction) AsFunc(codeGenerationContext *astmodel.CodeGenerationContext, receiver astmodel.InternalTypeName) *dst.FuncDecl {
	return fn.asFunc(fn, codeGenerationContext, receiver, fn.name)
}

// Equals determines if this Function is equal to another one
func (fn *ResourceFunction) Equals(f astmodel.Function, overrides astmodel.EqualityOverrides) bool {
	typedF, ok := f.(*ResourceFunction)
	if !ok {
		return false
	}

	// TODO: We're not actually checking function structure here
	return fn.resource.Equals(typedF.resource, overrides) && fn.name == typedF.name
}
