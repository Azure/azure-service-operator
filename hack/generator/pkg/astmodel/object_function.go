package astmodel

import (
	"github.com/dave/dst"
)

// ObjectFunction is a simple helper that implements the Function interface. It is intended for use for functions
// that only need information about the object they are operating on
// TODO: Relocate this into the functions package
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

type ObjectFunctionHandler func(f *ObjectFunction, codeGenerationContext *CodeGenerationContext, receiver TypeName, methodName string) *dst.FuncDecl

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
