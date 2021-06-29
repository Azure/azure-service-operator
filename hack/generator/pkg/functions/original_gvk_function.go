/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package functions

import (
	"github.com/dave/dst"

	"github.com/Azure/azure-service-operator/hack/generator/pkg/astbuilder"
	"github.com/Azure/azure-service-operator/hack/generator/pkg/astmodel"
)

// OriginalGVKFunction implements a function to return the GVK for the type on which it is called
// We put these on the *storage* versions of our resource types, giving us a way to obtain the right type of instance
// when the reconciler is working with ARM.
//
// func (resource *SomeResource) OriginalGVK() scheme.GroupVersionKind {
//     return scheme.GroupVersionKind{
//         Group: GroupVersion.Group,
//         Version: resource.Spec.OriginalVersion,
//         Kind: "SomeResource",
//     }
// }
//
type OriginalGVKFunction struct {
	idFactory astmodel.IdentifierFactory
}

// Ensure OriginalGVKFunction properly implements Function
var _ astmodel.Function = &OriginalGVKFunction{}

// NewOriginalGVKFunction creates a new OriginalGVKFunction
func NewOriginalGVKFunction(idFactory astmodel.IdentifierFactory) *OriginalGVKFunction {
	return &OriginalGVKFunction{
		idFactory: idFactory,
	}
}

// Name returns the name of this function, which is always OriginalGVK()
func (o OriginalGVKFunction) Name() string {
	return "OriginalGVK"
}

// RequiredPackageReferences returns the set of packages required by OriginalGVK()
func (o OriginalGVKFunction) RequiredPackageReferences() *astmodel.PackageReferenceSet {
	return astmodel.NewPackageReferenceSet(astmodel.APIMachinerySchemaReference)
}

// References shows that OriginalGVK() references no other generated types
func (o OriginalGVKFunction) References() astmodel.TypeNameSet {
	return astmodel.NewTypeNameSet()
}

// AsFunc returns the generated code for the OriginalGVK() function
func (o OriginalGVKFunction) AsFunc(
	generationContext *astmodel.CodeGenerationContext, receiver astmodel.TypeName) *dst.FuncDecl {
	gvkType := astmodel.GroupVersionKindTypeName.AsType(generationContext)
	groupVersionPackageGlobal := dst.NewIdent("GroupVersion")

	receiverName := o.idFactory.CreateIdentifier(receiver.Name(), astmodel.NotExported)

	spec := astbuilder.Selector(dst.NewIdent(receiverName), "Spec")

	builder := astbuilder.NewCompositeLiteralDetails(gvkType)
	builder.AddField("Group", astbuilder.Selector(groupVersionPackageGlobal, "Group"))
	builder.AddField("Version", astbuilder.Selector(spec, "OriginalVersion"))
	builder.AddField("Kind", astbuilder.StringLiteral(receiver.Name()))
	initGVK := builder.Build()

	funcDetails := &astbuilder.FuncDetails{
		ReceiverIdent: receiverName,
		ReceiverType:  astbuilder.Dereference(receiver.AsType(generationContext)),
		Name:          o.Name(),
		Body:          astbuilder.Statements(astbuilder.Returns(astbuilder.AddrOf(initGVK))),
	}

	funcDetails.AddReturn(astbuilder.Dereference(gvkType))

	return funcDetails.DefineFunc()
}

// Equals returns true if the passed function is equal to us, or false otherwise
func (o OriginalGVKFunction) Equals(f astmodel.Function) bool {
	_, ok := f.(*OriginalGVKFunction)
	// Equality is just based on Type for now
	return ok
}
