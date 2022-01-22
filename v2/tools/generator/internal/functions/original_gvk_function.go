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

// OriginalGVKFunction implements a function to return the GVK for the type on which it is called
// We put these on our resource types, giving us a way to obtain the right type of instance when the reconciler is
// working with ARM. The code differs slightly depending on whether we're injecting into an API or storage variant.
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
	idFactory                  astmodel.IdentifierFactory
	hasOriginalVersionFunction bool
	hasOriginalVersionProperty bool
}

// Ensure OriginalGVKFunction properly implements Function
var _ astmodel.Function = &OriginalGVKFunction{}

type OriginalVersionKind string

const (
	ReadProperty = OriginalVersionKind("property")
	ReadFunction = OriginalVersionKind("function")
)

// NewOriginalGVKFunction creates a new OriginalGVKFunction
func NewOriginalGVKFunction(originalVersion OriginalVersionKind, idFactory astmodel.IdentifierFactory) *OriginalGVKFunction {
	result := &OriginalGVKFunction{
		idFactory: idFactory,
	}

	result.hasOriginalVersionFunction = originalVersion == ReadFunction
	result.hasOriginalVersionProperty = originalVersion == ReadProperty

	return result
}

// Name returns the name of this function, which is always OriginalGVK()
func (o *OriginalGVKFunction) Name() string {
	return "OriginalGVK"
}

// RequiredPackageReferences returns the set of packages required by OriginalGVK()
func (o *OriginalGVKFunction) RequiredPackageReferences() *astmodel.PackageReferenceSet {
	return astmodel.NewPackageReferenceSet(astmodel.APIMachinerySchemaReference)
}

// References shows that OriginalGVK() references no other generated types
func (o *OriginalGVKFunction) References() astmodel.TypeNameSet {
	return astmodel.NewTypeNameSet()
}

// AsFunc returns the generated code for the OriginalGVK() function
func (o *OriginalGVKFunction) AsFunc(
	generationContext *astmodel.CodeGenerationContext, receiver astmodel.TypeName) *dst.FuncDecl {
	gvkType := astmodel.GroupVersionKindType.AsType(generationContext)
	groupVersionPackageGlobal := dst.NewIdent("GroupVersion")

	receiverName := o.idFactory.CreateReceiver(receiver.Name())

	spec := astbuilder.Selector(dst.NewIdent(receiverName), "Spec")

	builder := astbuilder.NewCompositeLiteralBuilder(gvkType)
	builder.AddField("Group", astbuilder.Selector(groupVersionPackageGlobal, "Group"))

	if o.hasOriginalVersionProperty {
		builder.AddField("Version", astbuilder.Selector(spec, "OriginalVersion"))
	} else if o.hasOriginalVersionFunction {
		builder.AddField("Version", astbuilder.CallExpr(spec, "OriginalVersion"))
	}

	builder.AddField("Kind", astbuilder.StringLiteral(receiver.Name()))
	initGVK := builder.Build()

	funcDetails := &astbuilder.FuncDetails{
		ReceiverIdent: receiverName,
		ReceiverType:  astbuilder.Dereference(receiver.AsType(generationContext)),
		Name:          o.Name(),
		Body:          astbuilder.Statements(astbuilder.Returns(astbuilder.AddrOf(initGVK))),
	}

	funcDetails.AddComments("returns a GroupValueKind for the original API version used to create the resource")
	funcDetails.AddReturn(astbuilder.Dereference(gvkType))

	return funcDetails.DefineFunc()
}

// Equals returns true if the passed function is equal to us, or false otherwise
func (o *OriginalGVKFunction) Equals(f astmodel.Function, _ astmodel.EqualityOverrides) bool {
	_, ok := f.(*OriginalGVKFunction)
	// Equality is just based on Type for now
	return ok
}
