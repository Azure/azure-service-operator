/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package functions

import (
	"go/token"

	"github.com/dave/dst"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astbuilder"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

// GetConditionsFunction returns a function declaration containing the implementation of the GetConditions() function.
//
//	func (r *<receiver>) GetConditions() genruntime.Conditions {
//	    return r.Status.Conditions
//	}
func GetConditionsFunction(k *ResourceFunction, codeGenerationContext *astmodel.CodeGenerationContext, receiver astmodel.TypeName, methodName string) *dst.FuncDecl {
	receiverIdent := k.IdFactory().CreateReceiver(receiver.Name())
	receiverType := receiver.AsType(codeGenerationContext)

	status := astbuilder.Selector(dst.NewIdent(receiverIdent), "Status")

	fn := &astbuilder.FuncDetails{
		Name:          methodName,
		ReceiverIdent: receiverIdent,
		ReceiverType:  astbuilder.PointerTo(receiverType),
		Body: []dst.Stmt{
			astbuilder.Returns(astbuilder.Selector(status, astmodel.ConditionsProperty)),
		},
	}

	fn.AddComments("returns the conditions of the resource")
	fn.AddReturn(astmodel.ConditionsType.AsType(codeGenerationContext))

	return fn.DefineFunc()
}

// SetConditionsFunction returns a function declaration containing the implementation of the SetConditions() function.
//
//	func (r *<receiver>) SetConditions(conditions genruntime.Conditions) {
//	    r.Status.Conditions = conditions
//	}
func SetConditionsFunction(k *ResourceFunction, codeGenerationContext *astmodel.CodeGenerationContext, receiver astmodel.TypeName, methodName string) *dst.FuncDecl {
	conditionsParameterName := k.IdFactory().CreateIdentifier(astmodel.ConditionsProperty, astmodel.NotExported)

	receiverIdent := k.IdFactory().CreateReceiver(receiver.Name())
	receiverType := receiver.AsType(codeGenerationContext)
	status := astbuilder.Selector(dst.NewIdent(receiverIdent), "Status")

	fn := &astbuilder.FuncDetails{
		Name:          methodName,
		ReceiverIdent: receiverIdent,
		ReceiverType:  astbuilder.PointerTo(receiverType),
		Body: []dst.Stmt{
			astbuilder.QualifiedAssignment(status, "Conditions", token.ASSIGN, dst.NewIdent(conditionsParameterName)),
		},
	}

	fn.AddParameter(
		conditionsParameterName,
		astmodel.ConditionsType.AsType(codeGenerationContext))
	fn.AddComments("sets the conditions on the resource status")

	return fn.DefineFunc()
}
