/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"github.com/dave/dst"
)

// Function represents something that is an (unnamed) Go function
type Function interface {
	// Name is the unique name of this function
	// (You can't have two functions with the same name on the same object or resource)
	Name() string

	RequiredPackageReferences() *PackageReferenceSet

	// References returns the set of types to which this function refers.
	// SHOULD include any types which this function references but its receiver doesn't.
	// SHOULD NOT include the receiver of this function.
	References() TypeNameSet

	// AsFunc renders the current instance as a Go abstract syntax tree
	AsFunc(codeGenerationContext *CodeGenerationContext, receiver TypeName) *dst.FuncDecl

	// Equals determines if this Function is equal to another one
	Equals(f Function) bool
}

var _ Function = &ObjectFunction{}

type ObjectFunctionHandler func(f *ObjectFunction, codeGenerationContext *CodeGenerationContext, receiver TypeName, methodName string) *dst.FuncDecl

// ObjectFunction is a simple helper that implements the Function interface. It is intended for use for functions
// that only need information about the object they are operating on
type ObjectFunction struct {
	name             string
	o                *ObjectType
	idFactory        IdentifierFactory
	asFunc           ObjectFunctionHandler
	requiredPackages *PackageReferenceSet
}

// NewObjectFunction creates a new object function
func NewObjectFunction(
	name string,
	objectType *ObjectType,
	requiredPackages *PackageReferenceSet,
	idFactory IdentifierFactory,
	asFunc ObjectFunctionHandler) *ObjectFunction {
	return &ObjectFunction{
		name:             name,
		o:                objectType,
		asFunc:           asFunc,
		requiredPackages: requiredPackages,
		idFactory:        idFactory,
	}
}

// Name returns the unique name of this function
// (You can't have two functions with the same name on the same object or resource)
func (k *ObjectFunction) Name() string {
	return k.name
}

func (k *ObjectFunction) IdFactory() IdentifierFactory {
	return k.idFactory
}

// RequiredPackageReferences returns the set of required packages for this function
func (k *ObjectFunction) RequiredPackageReferences() *PackageReferenceSet {
	return k.requiredPackages
}

// References returns the set of types to which this function refers.
// SHOULD include any types which this function references but its receiver doesn't.
// SHOULD NOT include the receiver of this function.
func (k *ObjectFunction) References() TypeNameSet {
	return nil
}

// AsFunc renders the current instance as a Go abstract syntax tree
func (k *ObjectFunction) AsFunc(codeGenerationContext *CodeGenerationContext, receiver TypeName) *dst.FuncDecl {
	return k.asFunc(k, codeGenerationContext, receiver, k.name)
}

// Equals checks if this function is equal to the passed in function
func (k *ObjectFunction) Equals(f Function) bool {
	typedF, ok := f.(*ObjectFunction)
	if !ok {
		return false
	}

	// TODO: We're not actually checking function structure here
	return k.o.Equals(typedF.o) && k.name == typedF.name
}

type resourceFunctionHandler func(f *resourceFunction, codeGenerationContext *CodeGenerationContext, receiver TypeName, methodName string) *dst.FuncDecl

// resourceFunction is a simple helper that implements the Function interface. It is intended for use for functions
// that only need information about the resource they are operating on
type resourceFunction struct {
	name             string
	resource         *ResourceType
	idFactory        IdentifierFactory
	asFunc           resourceFunctionHandler
	requiredPackages *PackageReferenceSet
}

// Name returns the unique name of this function
// (You can't have two functions with the same name on the same object or resource)
func (r *resourceFunction) Name() string {
	return r.name
}

// RequiredPackageReferences returns the set of required packages for this function
func (r *resourceFunction) RequiredPackageReferences() *PackageReferenceSet {
	return r.requiredPackages
}

// References returns the set of types to which this function refers.
// SHOULD include any types which this function references but its receiver doesn't.
// SHOULD NOT include the receiver of this function.
func (r *resourceFunction) References() TypeNameSet {
	return nil
}

// AsFunc renders the current instance as a Go abstract syntax tree
func (r *resourceFunction) AsFunc(codeGenerationContext *CodeGenerationContext, receiver TypeName) *dst.FuncDecl {
	return r.asFunc(r, codeGenerationContext, receiver, r.name)
}

// Equals determines if this Function is equal to another one
func (r *resourceFunction) Equals(f Function) bool {
	typedF, ok := f.(*resourceFunction)
	if !ok {
		return false
	}

	// TODO: We're not actually checking function structure here
	return r.resource.Equals(typedF.resource) && r.name == typedF.name
}
