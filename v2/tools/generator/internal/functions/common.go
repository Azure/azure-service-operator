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

// createBodyReturningLiteralString creates a function body which returns a single literal string.
func createBodyReturningLiteralString(
	result string,
	comment string,
	receiverTypeEnum ReceiverType,
) func(k *ObjectFunction, codeGenerationContext *astmodel.CodeGenerationContext, receiver astmodel.TypeName, methodName string) *dst.FuncDecl {
	return createBodyReturningValue(
		astbuilder.StringLiteral(result),
		astmodel.StringType,
		comment,
		receiverTypeEnum)
}

func createBodyReturningValue(
	result dst.Expr,
	returnType astmodel.Type,
	comment string,
	receiverTypeEnum ReceiverType,
) func(k *ObjectFunction, codeGenerationContext *astmodel.CodeGenerationContext, receiver astmodel.TypeName, methodName string) *dst.FuncDecl {
	return func(k *ObjectFunction, codeGenerationContext *astmodel.CodeGenerationContext, receiver astmodel.TypeName, methodName string) *dst.FuncDecl {
		receiverIdent := k.IdFactory().CreateReceiver(receiver.Name())

		// Support both ptr and non-ptr receivers
		var receiverType astmodel.Type
		if receiverTypeEnum == ReceiverTypePtr {
			receiverType = astmodel.NewOptionalType(receiver)
		} else {
			receiverType = receiver
		}

		fn := &astbuilder.FuncDetails{
			Name:          methodName,
			ReceiverIdent: receiverIdent,
			ReceiverType:  receiverType.AsType(codeGenerationContext),
			Params:        nil,
			Body:          astbuilder.Statements(astbuilder.Returns(result)),
		}

		fn.AddComments(comment)
		fn.AddReturn(returnType.AsType(codeGenerationContext))

		return fn.DefineFunc()
	}
}
