/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package functions

import (
	"github.com/dave/dst"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astbuilder"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

// OriginalVersionFunction implements a function to return the version for the type on which it is called
// We put these on the *API* versions of our spec types, giving us a way to obtain the right type of instance when the
// reconciler is working with ARM.
//
//	func (spec *SomeSpec) OriginalVersion() string {
//	    return GroupVersion.Version
//	}
type OriginalVersionFunction struct {
	idFactory astmodel.IdentifierFactory
}

// Ensure OriginalVersionFunction properly implements ValueFunction
var _ astmodel.ValueFunction = &OriginalVersionFunction{}

// NewOriginalVersionFunction creates a new OriginalVersionFunction
func NewOriginalVersionFunction(idFactory astmodel.IdentifierFactory) *OriginalVersionFunction {
	return &OriginalVersionFunction{
		idFactory: idFactory,
	}
}

// Name returns the name of this function, which is always OriginalVersion()
func (o *OriginalVersionFunction) Name() string {
	return "OriginalVersion"
}

// RequiredPackageReferences returns the set of packages required by OriginalVersion()
func (o *OriginalVersionFunction) RequiredPackageReferences() *astmodel.PackageReferenceSet {
	return astmodel.NewPackageReferenceSet()
}

// References shows that OriginalVersion() references no other generated types
func (o *OriginalVersionFunction) References() astmodel.TypeNameSet {
	return astmodel.NewTypeNameSet()
}

// AsFunc returns the generated code for the OriginalVersion() function
func (o *OriginalVersionFunction) AsFunc(
	generationContext *astmodel.CodeGenerationContext, receiver astmodel.TypeName) *dst.FuncDecl {
	groupVersionPackageGlobal := dst.NewIdent("GroupVersion")

	receiverName := o.idFactory.CreateReceiver(receiver.Name())

	returnVersion := astbuilder.Returns(
		astbuilder.Selector(groupVersionPackageGlobal, "Version"))

	funcDetails := &astbuilder.FuncDetails{
		ReceiverIdent: receiverName,
		ReceiverType:  astbuilder.Dereference(receiver.AsType(generationContext)),
		Name:          o.Name(),
		Body:          astbuilder.Statements(returnVersion),
	}

	funcDetails.AddComments("returns the original API version used to create the resource.")
	funcDetails.AddReturn(dst.NewIdent("string"))

	return funcDetails.DefineFunc()
}

// Equals returns true if the passed function is equal to us, or false otherwise
func (o *OriginalVersionFunction) Equals(f astmodel.Function, _ astmodel.EqualityOverrides) bool {
	_, ok := f.(*OriginalVersionFunction)
	// Equality is just based on Type for now
	return ok
}

// ReturnType indicates that this function returns a string
func (o *OriginalVersionFunction) ReturnType() astmodel.Type {
	return astmodel.StringType
}
