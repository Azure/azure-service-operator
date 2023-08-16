/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

// InterfaceInjector is a utility for injecting interface implementations into resources and objects
type InterfaceInjector struct {
	// visitor is used to do the actual injection
	visitor TypeVisitor[any]
}

// NewInterfaceInjector creates a new interface injector for modifying resources and objects
func NewInterfaceInjector() *InterfaceInjector {
	result := &InterfaceInjector{}

	result.visitor = TypeVisitorBuilder[any]{
		VisitObjectType:   result.injectInterfaceIntoObject,
		VisitResourceType: result.injectInterfaceIntoResource,
	}.Build()

	return result
}

// Inject modifies the passed type definition by injecting the passed function
func (i *InterfaceInjector) Inject(def TypeDefinition, implementation *InterfaceImplementation) (TypeDefinition, error) {
	result, err := i.visitor.VisitDefinition(def, implementation)
	if err != nil {
		return TypeDefinition{}, err
	}
	return result, nil
}

// injectFunctionIntoObject takes the function provided as a context and includes it on the
// provided object type
func (i *InterfaceInjector) injectInterfaceIntoObject(
	_ *TypeVisitor[any], ot *ObjectType, ctx any) (Type, error) {
	implementation := ctx.(*InterfaceImplementation)
	return ot.WithInterface(implementation), nil
}

// injectFunctionIntoResource takes the function provided as a context and includes it on the
// provided resource type
func (i *InterfaceInjector) injectInterfaceIntoResource(
	_ *TypeVisitor[any], rt *ResourceType, ctx any) (Type, error) {
	fn := ctx.(*InterfaceImplementation)
	return rt.WithInterface(fn), nil
}
