/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package storage

import "github.com/Azure/azure-service-operator/hack/generator/pkg/astmodel"

// PropertyInjector is a utility for injecting property definitions into resources and objects
type PropertyInjector struct {
	// visitor is used to do the actual injection
	visitor astmodel.TypeVisitor
}

// NewPropertyInjector creates a new property injector for modifying resources and objects
func NewPropertyInjector() *PropertyInjector {
	result := &PropertyInjector{}

	result.visitor = astmodel.TypeVisitorBuilder{
		VisitObjectType:   result.injectPropertyIntoObject,
		VisitResourceType: result.injectPropertyIntoResource,
	}.Build()

	return result
}

// Inject modifies the passed type definition by injecting the passed property
func (pi *PropertyInjector) Inject(def astmodel.TypeDefinition, prop *astmodel.PropertyDefinition) (astmodel.TypeDefinition, error) {
	return pi.visitor.VisitDefinition(def, prop)
}

// injectPropertyIntoObject takes the property provided as a context and includes it on the provided object type
func (pi *PropertyInjector) injectPropertyIntoObject(
	_ *astmodel.TypeVisitor, ot *astmodel.ObjectType, ctx interface{}) (astmodel.Type, error) {
	prop := ctx.(*astmodel.PropertyDefinition)
	return ot.WithProperty(prop), nil
}

// injectPropertyIntoResource takes the property  provided as a context and includes it on the provided resource type
func (pi *PropertyInjector) injectPropertyIntoResource(
	_ *astmodel.TypeVisitor, rt *astmodel.ResourceType, ctx interface{}) (astmodel.Type, error) {
	prop := ctx.(*astmodel.PropertyDefinition)
	return rt.WithProperty(prop), nil
}
