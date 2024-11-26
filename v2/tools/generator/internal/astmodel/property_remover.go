/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import "github.com/rotisserie/eris"

// PropertyRemover is a utility for removing property definitions from objects
type PropertyRemover struct {
	// visitor is used to do the actual injection
	visitor TypeVisitor[PropertyName]
}

// NewPropertyRemover creates a new property remover for modifying objects
func NewPropertyRemover() *PropertyRemover {
	result := &PropertyRemover{}

	result.visitor = TypeVisitorBuilder[PropertyName]{
		VisitObjectType: result.removePropertyFromObject,
	}.Build()

	return result
}

// Remove modifies the passed type definition by removing the passed property
func (pi *PropertyRemover) Remove(def TypeDefinition, name PropertyName) (TypeDefinition, error) {
	result, err := pi.visitor.VisitDefinition(def, name)
	if err != nil {
		return TypeDefinition{}, eris.Wrapf(err, "failed to remove property %q from %q", name, def.Name())
	}

	return result, nil
}

// injectPropertyIntoObject takes the property provided as a context and includes it on the provided object type
func (pi *PropertyRemover) removePropertyFromObject(
	_ *TypeVisitor[PropertyName], ot *ObjectType, name PropertyName,
) (Type, error) {
	return ot.WithoutProperty(name), nil
}
