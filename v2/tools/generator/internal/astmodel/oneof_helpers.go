/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"github.com/pkg/errors"
)

type PropertyNameAndType struct {
	PropertyName PropertyName
	TypeName     TypeName // the name of the type inside the pointer type
}

func resolveOneOfMemberToObjectType(
	t Type,
	definitions TypeDefinitionSet,
) (TypeName, *ObjectType, error) {
	// OneOfs are expected to contain properties that are:
	// pointer to typename to objectType

	tn, ok := AsTypeName(t)
	if !ok {
		return EmptyTypeName,
			nil,
			errors.Errorf("expected oneOf member to be a TypeName, instead was %s", DebugDescription(t))
	}

	propType, err := definitions.FullyResolve(tn)
	if err != nil {
		return EmptyTypeName,
			nil,
			errors.Wrapf(err, "unable to resolve oneOf member type %s", tn)
	}

	propObjType, ok := AsObjectType(propType)
	if !ok {
		return EmptyTypeName,
			nil,
			errors.Errorf("OneOf %s referenced non-object type %s", t, DebugDescription(propType))
	}

	return tn, propObjType, nil
}

func getDiscriminatorMapping(
	oneOf *ObjectType,
	propName PropertyName,
	definitions TypeDefinitionSet,
) map[string]PropertyNameAndType {
	props := oneOf.Properties().Copy()
	result := make(map[string]PropertyNameAndType, len(props))
	for _, prop := range props {
		propObjTypeName, propObjType, err := resolveOneOfMemberToObjectType(prop.PropertyType(), definitions)
		if err != nil {
			panic(err)
		}

		potentialDiscriminatorProp, found := propObjType.Property(propName)
		if !found {
			return nil
		}

		// Unwrap the property type if it's optional
		propType := potentialDiscriminatorProp.PropertyType()
		if opt, ok := AsOptionalType(propType); ok {
			propType = opt.Element()
		}

		potentialDiscriminatorType, err := definitions.FullyResolve(propType)
		if err != nil {
			panic(err) // should not be unresolvable
		}

		enumType, ok := AsEnumType(potentialDiscriminatorType)
		if !ok {
			return nil // if not an enum type cannot be used as discriminator
		}

		enumOptions := enumType.Options()
		if len(enumOptions) != 1 {
			return nil // if enum type has more than one value, cannot be used as discriminator
			// not entirely true since the options could all have distinct sets, but this is good enough for now
		}

		enumValue := enumOptions[0].Value
		if _, ok := result[enumValue]; ok {
			return nil // if values are not distinct for each member, cannot be used as discriminator
		}

		result[enumValue] = PropertyNameAndType{prop.PropertyName(), propObjTypeName}
	}

	return result
}

func DetermineDiscriminantAndValues(
	oneOf *ObjectType,
	definitions TypeDefinitionSet,
) (string, map[string]PropertyNameAndType, error) {
	// grab out the first member of the OneOf
	firstProp := oneOf.Properties().First()
	_, firstMember, err := resolveOneOfMemberToObjectType(firstProp.PropertyType(), definitions)
	if err != nil {
		return "", nil, errors.Wrap(err, "unable to resolve first member of OneOf")
	}

	// try to find a discriminator property out of the properties on the first member
	for _, prop := range firstMember.Properties().Copy() {
		mapping := getDiscriminatorMapping(oneOf, prop.PropertyName(), definitions)
		if mapping != nil {
			jsonTag, ok := prop.Tag("json")
			if !ok {
				// in reality every property will have a JSON tag
				panic("discriminator property had no JSON tag")
			}

			// first part of JSON tag is the JSON name
			return jsonTag[0], mapping, nil
		}
	}

	return "", nil, errors.Errorf("unable to determine a discriminator property for oneOf type")
}
