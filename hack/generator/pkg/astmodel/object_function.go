package astmodel

import (
	"github.com/dave/dst"
)

// objectFunction is a simple helper that implements the Function interface. It is intended for use for functions
// that only need information about the object they are operating on
// TODO: Reduce use of this in favour of the alternative in the functions package
type objectFunction struct {
	name             string
	o                *ObjectType
	idFactory        IdentifierFactory
	asFunc           objectFunctionHandler
	requiredPackages *PackageReferenceSet
}

var _ Function = &objectFunction{}

type objectFunctionHandler func(f *objectFunction, codeGenerationContext *CodeGenerationContext, receiver TypeName, methodName string) *dst.FuncDecl

// Name returns the unique name of this function
// (You can't have two functions with the same name on the same object or resource)
func (k *objectFunction) Name() string {
	return k.name
}

func (k *objectFunction) IdFactory() IdentifierFactory {
	return k.idFactory
}

// RequiredPackageReferences returns the set of required packages for this function
func (k *objectFunction) RequiredPackageReferences() *PackageReferenceSet {
	return k.requiredPackages
}

// References returns the set of types to which this function refers.
// SHOULD include any types which this function references but its receiver doesn't.
// SHOULD NOT include the receiver of this function.
func (k *objectFunction) References() TypeNameSet {
	return nil
}

// AsFunc renders the current instance as a Go abstract syntax tree
func (k *objectFunction) AsFunc(codeGenerationContext *CodeGenerationContext, receiver TypeName) *dst.FuncDecl {
	return k.asFunc(k, codeGenerationContext, receiver, k.name)
}

// Equals checks if this function is equal to the passed in function
func (k *objectFunction) Equals(f Function) bool {
	typedF, ok := f.(*objectFunction)
	if !ok {
		return false
	}

	// TODO: We're not actually checking function structure here
	return k.o.Equals(typedF.o) && k.name == typedF.name
}
