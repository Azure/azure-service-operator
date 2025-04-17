/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package functions

import (
	"go/token"

	"github.com/dave/dst"
	"github.com/rotisserie/eris"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astbuilder"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

// GetConditionsFunction returns a function declaration containing the implementation of the GetConditions() function.
//
//	func (r *<receiver>) GetConditions() genruntime.Conditions {
//	    return r.Status.Conditions
//	}
func GetConditionsFunction(
	k *ResourceFunction,
	codeGenerationContext *astmodel.CodeGenerationContext,
	receiver astmodel.TypeName,
	methodName string,
) (*dst.FuncDecl, error) {
	receiverIdent := k.IDFactory().CreateReceiver(receiver.Name())
	receiverExpr, err := receiver.AsTypeExpr(codeGenerationContext)
	if err != nil {
		return nil, eris.Wrapf(err, "creating type expression for %s", receiver)
	}

	status := astbuilder.Selector(dst.NewIdent(receiverIdent), "Status")

	fn := &astbuilder.FuncDetails{
		Name:          methodName,
		ReceiverIdent: receiverIdent,
		ReceiverType:  astbuilder.PointerTo(receiverExpr),
		Body: astbuilder.Statements(
			astbuilder.Returns(astbuilder.Selector(status, astmodel.ConditionsProperty)),
		),
	}

	fn.AddComments("returns the conditions of the resource")
	conditionsTypeExpr, err := astmodel.ConditionsType.AsTypeExpr(codeGenerationContext)
	if err != nil {
		return nil, eris.Wrapf(err, "creating type expression for %s", astmodel.ConditionsType)
	}

	fn.AddReturn(conditionsTypeExpr)

	return fn.DefineFunc(), nil
}

// SetConditionsFunction returns a function declaration containing the implementation of the SetConditions() function.
//
//	func (r *<receiver>) SetConditions(conditions genruntime.Conditions) {
//	    r.Status.Conditions = conditions
//	}
func SetConditionsFunction(
	k *ResourceFunction,
	codeGenerationContext *astmodel.CodeGenerationContext,
	receiver astmodel.TypeName,
	methodName string,
) (*dst.FuncDecl, error) {
	conditionsParameterName := k.IDFactory().CreateIdentifier(astmodel.ConditionsProperty, astmodel.NotExported)

	receiverIdent := k.IDFactory().CreateReceiver(receiver.Name())
	receiverExpr, err := receiver.AsTypeExpr(codeGenerationContext)
	if err != nil {
		return nil, eris.Wrapf(err, "creating type expression for %s", receiver)
	}

	status := astbuilder.Selector(dst.NewIdent(receiverIdent), "Status")

	fn := &astbuilder.FuncDetails{
		Name:          methodName,
		ReceiverIdent: receiverIdent,
		ReceiverType:  astbuilder.PointerTo(receiverExpr),
		Body: astbuilder.Statements(
			astbuilder.QualifiedAssignment(status, "Conditions", token.ASSIGN, dst.NewIdent(conditionsParameterName)),
		),
	}

	conditionsTypeExpr, err := astmodel.ConditionsType.AsTypeExpr(codeGenerationContext)
	if err != nil {
		return nil, eris.Wrapf(err, "unable to render type expression for %s", astmodel.ConditionsType)
	}

	fn.AddParameter(
		conditionsParameterName,
		conditionsTypeExpr)
	fn.AddComments("sets the conditions on the resource status")

	return fn.DefineFunc(), nil
}
