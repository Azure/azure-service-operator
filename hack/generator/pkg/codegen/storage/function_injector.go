/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package storage

import "github.com/Azure/azure-service-operator/hack/generator/pkg/astmodel"

// FunctionInjector is a utility for injecting function definitions into resources and objects
type FunctionInjector struct {
	// visitor is used to do the actual injection
	visitor astmodel.TypeVisitor
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

// Inject modifies the passed type definition by injecting the passed function(s)
func (fi *FunctionInjector) Inject(
	def astmodel.TypeDefinition, functions ...astmodel.Function) (astmodel.TypeDefinition, error) {
	return fi.visitor.VisitDefinition(def, functions)
}

// injectFunctionIntoObject takes the function provided as a context and includes it on the
// provided object type
func (_ *FunctionInjector) injectFunctionIntoObject(
	_ *astmodel.TypeVisitor, ot *astmodel.ObjectType, ctx interface{}) (astmodel.Type, error) {
	functions := ctx.([]astmodel.Function)

	result := ot
	for _, fn := range functions {
		result = result.WithFunction(fn)
	}

	return result, nil
}

// injectFunctionIntoResource takes the function provided as a context and includes it on the
// provided resource type
func (_ *FunctionInjector) injectFunctionIntoResource(
	_ *astmodel.TypeVisitor, rt *astmodel.ResourceType, ctx interface{}) (astmodel.Type, error) {
	functions := ctx.([]astmodel.Function)

	result := rt
	for _, fn := range functions {
		result = result.WithFunction(fn)
	}

	return result, nil
}
