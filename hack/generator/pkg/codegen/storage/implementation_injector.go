/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package storage

import "github.com/Azure/azure-service-operator/hack/generator/pkg/astmodel"

// ImplementationInjector is a utility for injecting interface implementations into resources and objects
type ImplementationInjector struct {
	// visitor is used to do the actual injection
	visitor astmodel.TypeVisitor
}

// NewImplementationInjector creates a new implementation injector for modifying resources and objects
func NewImplementationInjector() *ImplementationInjector {
	result := &ImplementationInjector{}

	result.visitor = astmodel.TypeVisitorBuilder{
		VisitObjectType:   result.injectImplementationIntoObject,
		VisitResourceType: result.injectImplementationIntoResource,
	}.Build()

	return result
}

// Inject modifies the passed type definition by injecting the passed implementation
func (ii *ImplementationInjector) Inject(
	def astmodel.TypeDefinition, implementation *astmodel.InterfaceImplementation) (astmodel.TypeDefinition, error) {
	return ii.visitor.VisitDefinition(def, implementation)
}

// injectImplementationIntoObject takes the function provided as a context and includes it on the provided object type
func (ii *ImplementationInjector) injectImplementationIntoObject(
	_ *astmodel.TypeVisitor, ot *astmodel.ObjectType, ctx interface{}) (astmodel.Type, error) {
	implementation := ctx.(*astmodel.InterfaceImplementation)
	return ot.WithInterface(implementation), nil
}

// injectImplementationIntoResource takes the implementation provided as a context and includes it on the provided resource type
func (ii *ImplementationInjector) injectImplementationIntoResource(
	_ *astmodel.TypeVisitor, rt *astmodel.ResourceType, ctx interface{}) (astmodel.Type, error) {
	implementation := ctx.(*astmodel.InterfaceImplementation)
	return rt.WithInterface(implementation), nil
}
