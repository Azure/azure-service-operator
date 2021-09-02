package functions

import (
	"github.com/dave/dst"

	"github.com/Azure/azure-service-operator/hack/generator/pkg/astmodel"
)

// ObjectFunction is a simple helper that implements the Function interface. It is intended for use for functions
// that only need information about the object they are operating on
type ObjectFunction struct {
	name             string
	o                *astmodel.ObjectType
	idFactory        astmodel.IdentifierFactory
	asFunc           ObjectFunctionHandler
	requiredPackages *astmodel.PackageReferenceSet
}

// NewObjectFunction creates a new object function
func NewObjectFunction(
	name string,
	objectType *astmodel.ObjectType,
	requiredPackages *astmodel.PackageReferenceSet,
	idFactory astmodel.IdentifierFactory,
	asFunc ObjectFunctionHandler) *ObjectFunction {
	return &ObjectFunction{
		name:             name,
		o:                objectType,
		asFunc:           asFunc,
		requiredPackages: requiredPackages,
		idFactory:        idFactory,
	}
}

type ObjectFunctionHandler func(f *ObjectFunction, codeGenerationContext *astmodel.CodeGenerationContext, receiver astmodel.TypeName, methodName string) *dst.FuncDecl

// Name returns the unique name of this function
// (You can't have two functions with the same name on the same object or resource)
func (k *ObjectFunction) Name() string {
	return k.name
}

func (k *ObjectFunction) IdFactory() astmodel.IdentifierFactory {
	return k.idFactory
}

// RequiredPackageReferences returns the set of required packages for this function
func (k *ObjectFunction) RequiredPackageReferences() *astmodel.PackageReferenceSet {
	return k.requiredPackages
}

// References returns the set of types to which this function refers.
// SHOULD include any types which this function references but its receiver doesn't.
// SHOULD NOT include the receiver of this function.
func (k *ObjectFunction) References() astmodel.TypeNameSet {
	return nil
}

// AsFunc renders the current instance as a Go abstract syntax tree
func (k *ObjectFunction) AsFunc(codeGenerationContext *astmodel.CodeGenerationContext, receiver astmodel.TypeName) *dst.FuncDecl {
	return k.asFunc(k, codeGenerationContext, receiver, k.name)
}

// Equals checks if this function is equal to the passed in function
func (k *ObjectFunction) Equals(f astmodel.Function) bool {
	typedF, ok := f.(*ObjectFunction)
	if !ok {
		return false
	}

	// TODO: We're not actually checking function structure here
	return k.o.Equals(typedF.o) && k.name == typedF.name
}
