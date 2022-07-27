/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package interfaces

import (
	"github.com/dave/dst"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astbuilder"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/functions"
)

// TODO: Remove this?
func AddTenantResourceInterface(
	resourceDef astmodel.TypeDefinition,
	idFactory astmodel.IdentifierFactory) (astmodel.TypeDefinition, error) {

	tenantFunction := functions.NewObjectFunction("Tenant", idFactory, newTenantFunction())
	tenantResourceImpl := astmodel.NewInterfaceImplementation(astmodel.TenantResourceType, tenantFunction)

	r := resourceDef.Type().(*astmodel.ResourceType)
	r = r.WithInterface(tenantResourceImpl)

	return resourceDef.WithType(r), nil
}

// newTenantFunction creates a function "Tenant" that returns empty string.
// This function and its associated interface are only used as a marker that a particular type is tenant scoped
func newTenantFunction() func(k *functions.ObjectFunction, codeGenerationContext *astmodel.CodeGenerationContext, receiver astmodel.TypeName, methodName string) *dst.FuncDecl {
	return func(k *functions.ObjectFunction, codeGenerationContext *astmodel.CodeGenerationContext, receiver astmodel.TypeName, methodName string) *dst.FuncDecl {
		receiverIdent := k.IdFactory().CreateReceiver(receiver.Name())
		receiverType := astmodel.NewOptionalType(receiver)

		fn := &astbuilder.FuncDetails{
			Name:          methodName,
			ReceiverIdent: receiverIdent,
			ReceiverType:  receiverType.AsType(codeGenerationContext),
			Params:        nil,
		}

		fn.AddComments("is a marker function indicating that this resource is scoped at the tenant level")

		return fn.DefineFunc()
	}
}
