/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package functions

import (
	"github.com/dave/dst"
	"github.com/rotisserie/eris"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astbuilder"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

// createBodyReturningLiteralString creates a function body which returns a single literal string.
func createBodyReturningLiteralString(
	result string,
	comment string,
	receiverTypeEnum ReceiverType,
) ObjectFunctionHandler {
	return createBodyReturningValue(
		astbuilder.StringLiteral(result),
		astmodel.StringType,
		comment,
		receiverTypeEnum,
	)
}

// resourceParameterIdent returns the identifier a webhook function uses for the resource, whether as a
// parameter (the validators) or as the local the defaulter type-asserts obj into. It renames the
// identifier when the receiver has already taken it, as happens for a kind called AlertRuleResource.
func resourceParameterIdent(receiverIdent string, idFactory astmodel.IdentifierFactory) string {
	knownLocals := astmodel.NewKnownLocalsSet(idFactory)
	knownLocals.Add(receiverIdent)

	return knownLocals.CreateLocal("resource", "", "Obj")
}

func createBodyReturningValue(
	result dst.Expr,
	returnType astmodel.Type,
	comment string,
	receiverTypeEnum ReceiverType,
) ObjectFunctionHandler {
	return func(
		k *ObjectFunction,
		codeGenerationContext *astmodel.CodeGenerationContext,
		receiver astmodel.TypeName,
		methodName string,
	) (*dst.FuncDecl, error) {
		receiverIdent := k.IDFactory().CreateReceiver(receiver.Name())

		// Support both ptr and non-ptr receivers
		var receiverType astmodel.Type
		if receiverTypeEnum == ReceiverTypePtr {
			receiverType = astmodel.NewOptionalType(receiver)
		} else {
			receiverType = receiver
		}

		receiverExpr, err := receiverType.AsTypeExpr(codeGenerationContext)
		if err != nil {
			return nil, eris.Wrapf(err, "creating receiver expression for %s", receiverType)
		}

		fn := &astbuilder.FuncDetails{
			Name:          methodName,
			ReceiverIdent: receiverIdent,
			ReceiverType:  receiverExpr,
			Params:        nil,
			Body:          astbuilder.Statements(astbuilder.Returns(result)),
		}

		fn.AddComments(comment)
		returnTypeExpr, err := returnType.AsTypeExpr(codeGenerationContext)
		if err != nil {
			return nil, eris.Wrapf(err, "creating type expression for %s", returnType)
		}

		fn.AddReturn(returnTypeExpr)

		return fn.DefineFunc(), nil
	}
}
