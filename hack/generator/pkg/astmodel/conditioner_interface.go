/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"go/token"

	"github.com/dave/dst"

	"github.com/Azure/azure-service-operator/hack/generator/pkg/astbuilder"
)

const (
	ConditionsProperty = "Conditions"
)

// NewConditionerInterfaceImpl creates an InterfaceImplementation with GetConditions() and
// SetConditions() methods, implementing the genruntime.Conditioner interface.
func NewConditionerInterfaceImpl(
	idFactory IdentifierFactory,
	resource *ResourceType) (*InterfaceImplementation, error) {

	getConditions := &resourceFunction{
		name:             "Get" + ConditionsProperty,
		resource:         resource,
		idFactory:        idFactory,
		asFunc:           getConditionsFunction,
		requiredPackages: NewPackageReferenceSet(GenRuntimeConditionsReference),
	}

	setConditions := &resourceFunction{
		name:             "Set" + ConditionsProperty,
		resource:         resource,
		idFactory:        idFactory,
		asFunc:           setConditionsFunction,
		requiredPackages: NewPackageReferenceSet(GenRuntimeConditionsReference),
	}

	result := NewInterfaceImplementation(
		ConditionerTypeName,
		getConditions,
		setConditions)

	return result, nil
}

// getConditionsFunction returns a function declaration containing the implementation of the GetConditions() function.
//
// func (r *<receiver>)GetConditions() genruntime.Conditions {
//     return r.Status.Conditions
// }
func getConditionsFunction(k *resourceFunction, codeGenerationContext *CodeGenerationContext, receiver TypeName, methodName string) *dst.FuncDecl {
	receiverIdent := k.idFactory.CreateIdentifier(receiver.Name(), NotExported)
	receiverType := receiver.AsType(codeGenerationContext)

	status := astbuilder.Selector(dst.NewIdent(receiverIdent), "Status")

	fn := &astbuilder.FuncDetails{
		Name:          methodName,
		ReceiverIdent: receiverIdent,
		ReceiverType: &dst.StarExpr{
			X: receiverType,
		},
		Body: []dst.Stmt{
			astbuilder.Returns(astbuilder.Selector(status, ConditionsProperty)),
		},
	}

	fn.AddComments("returns the conditions of the resource")
	fn.AddReturn(ConditionsTypeName.AsType(codeGenerationContext))

	return fn.DefineFunc()
}

// setConditionsFunction returns a function declaration containing the implementation of the SetConditions() function.
//
// func (r *<receiver>)SetConditions(conditions genruntime.Conditions) {
//     r.Status.Conditions = conditions
// }
func setConditionsFunction(k *resourceFunction, codeGenerationContext *CodeGenerationContext, receiver TypeName, methodName string) *dst.FuncDecl {
	conditionsParameterName := k.idFactory.CreateIdentifier(ConditionsProperty, NotExported)

	receiverIdent := k.idFactory.CreateIdentifier(receiver.Name(), NotExported)
	receiverType := receiver.AsType(codeGenerationContext)
	status := astbuilder.Selector(dst.NewIdent(receiverIdent), "Status")

	fn := &astbuilder.FuncDetails{
		Name:          methodName,
		ReceiverIdent: receiverIdent,
		ReceiverType: &dst.StarExpr{
			X: receiverType,
		},
		Body: []dst.Stmt{
			astbuilder.QualifiedAssignment(status, "Conditions", token.ASSIGN, dst.NewIdent(conditionsParameterName)),
		},
	}

	fn.AddParameter(
		conditionsParameterName,
		ConditionsTypeName.AsType(codeGenerationContext))
	fn.AddComments("sets the conditions on the resource status")

	return fn.DefineFunc()
}
