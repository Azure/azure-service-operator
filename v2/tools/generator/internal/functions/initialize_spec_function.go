/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package functions

import (
	"fmt"

	"github.com/dave/dst"
	"github.com/pkg/errors"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astbuilder"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

// NewInitializeSpecFunction creates a function that initializes a resource's spec from its status.
// Used as a part of implementing the ImportableResource interface.
// def is the resource definition that we're generating the function for.
// specInitializeFunction is the name of the function on the spec that initializes it from the status.
// idFactory is used to create identifiers for the function.
func NewInitializeSpecFunction(
	def astmodel.TypeDefinition,
	specInitializeFunction string,
	idFactory astmodel.IdentifierFactory) (astmodel.Function, error) {
	rsrc, ok := astmodel.AsResourceType(def.Type())
	if !ok {
		return nil, errors.Errorf("expected %q to be a resource", def.Name())
	}

	statusType, ok := astmodel.AsTypeName(rsrc.StatusType())
	if !ok {
		return nil, errors.Errorf("expected %q to be a TypeName", rsrc.StatusType())
	}

	requiredPackages := astmodel.NewPackageReferenceSet(
		astmodel.GenRuntimeReference,
		astmodel.FmtReference,
	)

	createFn := func(
		fn *ResourceFunction,
		codeGenerationContext *astmodel.CodeGenerationContext,
		receiver astmodel.TypeName,
		methodName string,
	) *dst.FuncDecl {
		fmtPackage := codeGenerationContext.MustGetImportedPackageName(astmodel.FmtReference)

		receiverType := receiver.AsType(codeGenerationContext)
		receiverName := idFactory.CreateReceiver(receiver.Name())

		knownLocals := astmodel.NewKnownLocalsSet(idFactory)
		knownLocals.Add(receiverName)

		statusParam := knownLocals.CreateLocal("status", "", "Source")
		statusLocal := knownLocals.CreateLocal("s", "")

		// return <receiver>.Spec.Initialize_From_Status(s)
		returnConversion := astbuilder.Returns(
			astbuilder.CallExpr(
				astbuilder.Selector(dst.NewIdent(receiverName), "Spec"),
				specInitializeFunction,
				dst.NewIdent(statusLocal)))

		// if s, ok := fromStatus.(<type of status>); ok {
		//   return receiver.Spec.InitializeFromStatus(s)
		// }
		initialize := astbuilder.IfType(
			dst.NewIdent(statusParam),
			astbuilder.Dereference(statusType.AsType(codeGenerationContext)),
			statusLocal,
			returnConversion)

		// return fmt.Errorf("expected Status of type <type of status> but received %T instead", fromStatus)
		returnError := astbuilder.Returns(
			astbuilder.FormatError(
				fmtPackage,
				fmt.Sprintf("expected Status of type %s but received %%T instead", statusType.Name()),
				dst.NewIdent(statusParam)))
		returnError.Decorations().Before = dst.EmptyLine

		funcDetails := astbuilder.FuncDetails{
			Name:          methodName,
			ReceiverIdent: receiverName,
			ReceiverType:  astbuilder.Dereference(receiverType),
			Body: astbuilder.Statements(
				initialize,
				returnError),
		}

		funcDetails.AddComments("initializes the spec for this resource from the given status")
		funcDetails.AddParameter(statusParam, astmodel.ConvertibleStatusInterfaceType.AsType(codeGenerationContext))
		funcDetails.AddReturns("error")

		return funcDetails.DefineFunc()
	}

	return NewResourceFunction(
		"InitializeSpec",
		rsrc,
		idFactory,
		createFn,
		requiredPackages), nil
}
