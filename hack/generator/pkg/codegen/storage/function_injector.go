/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package storage

import "github.com/Azure/azure-service-operator/hack/generator/pkg/astmodel"

// FunctionInjector is a utility for injecting function definitions into resources and objects
type FunctionInjector struct {
	visitor astmodel.TypeVisitor // Visitor used to achieve the required modification
}

// NewFunctionInjector creates a new function injector for modifying resources and objects
func NewFunctionInjector() *FunctionInjector {
	result := &FunctionInjector{}

	result.visitor = astmodel.TypeVisitorBuilder{
		VisitObjectType:   result.injectFunctionIntoObject,
		VisitResourceType: result.injectFunctionIntoResource,
	}.Build()

	return result
}

// Inject modifies the passed type definition by injecting the passed function
func (fi *FunctionInjector) Inject(def astmodel.TypeDefinition, fn astmodel.Function) (astmodel.TypeDefinition, error) {
	return fi.visitor.VisitDefinition(def, fn)
}

// injectFunctionIntoObject takes the function provided as a context and includes it on the
// provided object type
func (_ FunctionInjector) injectFunctionIntoObject(
	_ *astmodel.TypeVisitor, ot *astmodel.ObjectType, ctx interface{}) (astmodel.Type, error) {
	fn := ctx.(astmodel.Function)
	return ot.WithFunction(fn), nil
}

// injectFunctionIntoResource takes the function provided as a context and includes it on the
// provided resource type
func (_ FunctionInjector) injectFunctionIntoResource(
	_ *astmodel.TypeVisitor, rt *astmodel.ResourceType, ctx interface{}) (astmodel.Type, error) {
	fn := ctx.(astmodel.Function)
	return rt.WithFunction(fn), nil
}
