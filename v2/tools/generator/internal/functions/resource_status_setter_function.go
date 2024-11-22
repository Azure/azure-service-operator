/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package functions

import (
	"fmt"

	"github.com/dave/dst"
	"github.com/rotisserie/eris"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astbuilder"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

// ResourceStatusSetterFunction is a function that sets the status on a resource
type ResourceStatusSetterFunction struct {
	nameOfStatusType string
	idFactory        astmodel.IdentifierFactory
}

var _ astmodel.Function = &ResourceStatusSetterFunction{}

// NewResourceStatusSetterFunction creates a new ResourceStatusSetterFunction
// resource is the resource type that will host the function
// idFactory is an IdentifierFactory for creating local variable names
func NewResourceStatusSetterFunction(
	resource *astmodel.ResourceType,
	idFactory astmodel.IdentifierFactory,
) *ResourceStatusSetterFunction {
	statusTypeName, ok := astmodel.AsTypeName(resource.StatusType())
	if !ok {
		panic(fmt.Sprintf("expected Status to be a TypeName but found %T", resource.StatusType()))
	}

	return &ResourceStatusSetterFunction{
		nameOfStatusType: statusTypeName.Name(),
		idFactory:        idFactory,
	}
}

// Name returns the unique name of this function
func (fn ResourceStatusSetterFunction) Name() string {
	return "SetStatus"
}

// RequiredPackageReferences returns the set of package references required by this function
func (fn ResourceStatusSetterFunction) RequiredPackageReferences() *astmodel.PackageReferenceSet {
	return astmodel.NewPackageReferenceSet(astmodel.ErisReference, astmodel.GenRuntimeReference)
}

// References returns the set of other types required by this function
func (fn ResourceStatusSetterFunction) References() astmodel.TypeNameSet {
	return astmodel.NewTypeNameSet(astmodel.ConvertibleStatusInterfaceType)
}

// AsFunc generates the required function declaration
func (fn ResourceStatusSetterFunction) AsFunc(
	codeGenerationContext *astmodel.CodeGenerationContext,
	receiver astmodel.InternalTypeName,
) (*dst.FuncDecl, error) {
	receiverIdent := fn.idFactory.CreateReceiver(receiver.Name())
	receiverType := astmodel.NewOptionalType(receiver)
	receiverTypeExpr, err := receiverType.AsTypeExpr(codeGenerationContext)
	if err != nil {
		return nil, eris.Wrap(err, "creating receiver type expression")
	}

	statusLocal := "st"
	statusParameter := "status"
	errorsPackage := codeGenerationContext.MustGetImportedPackageName(astmodel.ErisReference)

	// <receiver>.Status = st
	assignFromStatus := astbuilder.SimpleAssignment(
		astbuilder.Selector(dst.NewIdent(receiverIdent), "Status"), astbuilder.Dereference(dst.NewIdent(statusLocal)))

	// if st, ok := status.(<type>); ok {
	//     <receiver>.Status = st
	//     return nil
	// }
	simplePath := astbuilder.IfType(
		dst.NewIdent(statusParameter),
		astbuilder.PointerTo(dst.NewIdent(fn.nameOfStatusType)),
		statusLocal,
		assignFromStatus, astbuilder.Returns(astbuilder.Nil()))
	astbuilder.AddComment(&simplePath.Decorations().Start, "// If we have exactly the right type of status, assign it")
	simplePath.Decorations().Before = dst.NewLine

	// var st <nameOfStatusType>
	declareLocal := astbuilder.LocalVariableDeclaration(
		statusLocal,
		dst.NewIdent(fn.nameOfStatusType),
		"// Convert status to required version")
	declareLocal.Decorations().Before = dst.EmptyLine

	// err := status.ConvertStatusTo(&st)
	convert := astbuilder.ShortDeclaration(
		"err",
		astbuilder.CallQualifiedFunc(
			statusParameter,
			"ConvertStatusTo",
			astbuilder.AddrOf(dst.NewIdent(statusLocal))))

	// if err != nil {
	//     return errors.Wrap(err, "failed to convert status")
	// }
	returnIfErr := astbuilder.ReturnIfNotNil(dst.NewIdent("err"),
		astbuilder.WrappedErrorf(errorsPackage, "failed to convert status"))

	// <receiver>.Status = st
	assignFromLocal := astbuilder.SimpleAssignment(
		astbuilder.Selector(dst.NewIdent(receiverIdent), "Status"), dst.NewIdent(statusLocal))
	assignFromLocal.Decorations().Before = dst.EmptyLine

	returnNil := astbuilder.Returns(astbuilder.Nil())

	builder := &astbuilder.FuncDetails{
		ReceiverIdent: receiverIdent,
		ReceiverType:  receiverTypeExpr,
		Name:          "SetStatus",
		Body: astbuilder.Statements(
			simplePath,
			declareLocal,
			convert,
			returnIfErr,
			assignFromLocal,
			returnNil),
	}

	convertibleStatusInterfaceTypeExpr, err := astmodel.ConvertibleStatusInterfaceType.AsTypeExpr(codeGenerationContext)
	if err != nil {
		return nil, eris.Wrap(err, "creating convertible status interface type expression")
	}

	builder.AddParameter(statusParameter, convertibleStatusInterfaceTypeExpr)

	errorTypeExpr, err := astmodel.ErrorType.AsTypeExpr(codeGenerationContext)
	if err != nil {
		return nil, eris.Wrap(err, "creating error type expression")
	}

	builder.AddReturn(errorTypeExpr)

	builder.AddComments("sets the status of this resource")

	return builder.DefineFunc(), nil
}

// Equals returns true if the passed function is equal
func (fn ResourceStatusSetterFunction) Equals(other astmodel.Function, _ astmodel.EqualityOverrides) bool {
	otherFn, ok := other.(*ResourceStatusSetterFunction)
	if !ok {
		return false
	}

	return fn.nameOfStatusType == otherFn.nameOfStatusType
}
