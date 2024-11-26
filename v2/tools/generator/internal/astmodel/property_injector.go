/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import "github.com/rotisserie/eris"

// PropertyInjector is a utility for injecting property definitions into resources and objects
type PropertyInjector struct {
	// visitor is used to do the actual injection
	visitor TypeVisitor[*PropertyDefinition]
}

// NewPropertyInjector creates a new property injector for modifying resources and objects
func NewPropertyInjector() *PropertyInjector {
	result := &PropertyInjector{}

	result.visitor = TypeVisitorBuilder[*PropertyDefinition]{
		VisitObjectType: result.injectPropertyIntoObject,
	}.Build()

	return result
}

// Inject modifies the passed type definition by injecting the passed property
func (pi *PropertyInjector) Inject(def TypeDefinition, prop *PropertyDefinition) (TypeDefinition, error) {
	result, err := pi.visitor.VisitDefinition(def, prop)
	if err != nil {
		return TypeDefinition{}, eris.Wrapf(err, "failed to inject property %q into %q", prop.PropertyName(), def.Name())
	}

	return result, nil
}

// injectPropertyIntoObject takes the property provided as a context and includes it on the provided object type
func (pi *PropertyInjector) injectPropertyIntoObject(
	_ *TypeVisitor[*PropertyDefinition], ot *ObjectType, prop *PropertyDefinition,
) (Type, error) {
	// Ensure that we don't already have a property with the same name
	if _, ok := ot.Property(prop.PropertyName()); ok {
		return nil, eris.Errorf("already has property named %q", prop.PropertyName())
	}

	return ot.WithProperty(prop), nil
}
