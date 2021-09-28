/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package functions

import (
	"fmt"
	"strings"

	"github.com/dave/dst"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astbuilder"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

type ReceiverType string

const (
	ReceiverTypePtr    = ReceiverType("ptr")
	ReceiverTypeStruct = ReceiverType("struct")
)

// NewGetTypeFunction returns a function "GetType" that returns a static string value for the ResourceType of a resource
func NewGetTypeFunction(
	armType string,
	idFactory astmodel.IdentifierFactory,
	receiverType ReceiverType) astmodel.Function {

	comment := fmt.Sprintf("returns the ARM Type of the resource. This is always %q", strings.Trim(armType, "\""))
	result := NewObjectFunction("Get"+astmodel.TypeProperty, idFactory, newStaticStringReturnFunctionBody(armType, comment, receiverType))
	result.AddPackageReference(astmodel.GenRuntimeReference)
	return result
}

func newStaticStringReturnFunctionBody(
	result string,
	comment string,
	receiverTypeEnum ReceiverType) func(k *ObjectFunction, codeGenerationContext *astmodel.CodeGenerationContext, receiver astmodel.TypeName, methodName string) *dst.FuncDecl {

	return func(k *ObjectFunction, codeGenerationContext *astmodel.CodeGenerationContext, receiver astmodel.TypeName, methodName string) *dst.FuncDecl {
		receiverIdent := k.IdFactory().CreateIdentifier(receiver.Name(), astmodel.NotExported)

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
			Body: astbuilder.Statements(
				astbuilder.Returns(astbuilder.TextLiteral(result))),
		}

		fn.AddComments(comment)
		fn.AddReturns("string")

		return fn.DefineFunc()
	}
}
