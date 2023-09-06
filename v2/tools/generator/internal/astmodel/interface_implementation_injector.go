/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

// InterfaceImplementationInjector is a utility for injecting interface implementations into resources and objects
type InterfaceImplementationInjector struct {
	// visitor is used to do the actual injection
	visitor TypeVisitor[any]
}

// NewInterfaceImplementationInjector creates a new interface implementation injector for modifying resources & objects
func NewInterfaceImplementationInjector() *InterfaceImplementationInjector {
	result := &InterfaceImplementationInjector{}

	result.visitor = TypeVisitorBuilder[any]{
		VisitObjectType:   result.injectInterfaceImplementationIntoObject,
		VisitResourceType: result.injectInterfaceImplementationIntoResource,
	}.Build()

	return result
}

// Inject modifies the passed type definition by injecting the passed interface implementation
func (injector *InterfaceImplementationInjector) Inject(
	def TypeDefinition, implementations ...*InterfaceImplementation) (TypeDefinition, error) {
	result := def

	for _, impl := range implementations {
		var err error
		result, err = injector.visitor.VisitDefinition(result, impl)
		if err != nil {
			return TypeDefinition{}, err
		}
	}

	return result, nil
}

// injectInterfaceImplementationIntoObject takes the interface implementationprovided as a context and includes it on
// the provided object type
func (_ *InterfaceImplementationInjector) injectInterfaceImplementationIntoObject(
	_ *TypeVisitor[any], ot *ObjectType, ctx interface{}) (Type, error) {
	impl := ctx.(*InterfaceImplementation)
	return ot.WithInterface(impl), nil
}

// injectInterfaceImplementationIntoResource takes the interface implementation provided as a context and includes it on
// the provided resource type
func (_ *InterfaceImplementationInjector) injectInterfaceImplementationIntoResource(
	_ *TypeVisitor[any], rt *ResourceType, ctx interface{}) (Type, error) {
	impl := ctx.(*InterfaceImplementation)
	return rt.WithInterface(impl), nil
}
