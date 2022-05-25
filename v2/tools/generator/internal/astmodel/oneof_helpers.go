/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import "fmt"

type PropertyNameAndType struct {
	PropertyName PropertyName
	TypeName     TypeName // the name of the type inside the pointer type
}

func resolveOneOfMemberToObjectType(t Type, definitions TypeDefinitionSet) (TypeName, *ObjectType) {
	// OneOfs are expected to contain properties that are:
	// pointer to typename to objectType

	propType, err := definitions.FullyResolve(t)
	if err != nil {
		panic(err) // type should not contain unresolvable references at this point
	}

	optionalType, ok := propType.(*OptionalType)
	if !ok {
		panic(fmt.Sprintf("OneOf contained non-optional type %s", propType.String()))
	}

	typeName, ok := optionalType.Element().(TypeName)
	if !ok {
		panic("Expected OneOf to have pointer to TypeName")
	}

	resolvedInnerOptional, err := definitions.FullyResolve(typeName)
	if err != nil {
		panic(err) // definitions should not contain unresolvable references at this point
	}

	propObjType, ok := AsObjectType(resolvedInnerOptional)
	if !ok {
		panic(fmt.Sprintf("OneOf contained non-object type %s", propType.String()))
	}

	return typeName, propObjType
}

func getDiscriminatorMapping(
	oneOf *ObjectType,
	propName PropertyName,
	definitions TypeDefinitionSet,
) map[string]PropertyNameAndType {
	props := oneOf.Properties()
	result := make(map[string]PropertyNameAndType, len(props))
	for _, prop := range props {
		propObjTypeName, propObjType := resolveOneOfMemberToObjectType(prop.PropertyType(), definitions)

		potentialDiscriminatorProp, ok := propObjType.Property(propName)
		if !ok {
			return nil
		}

		// Unwrap the property type if it's optional
		propType := potentialDiscriminatorProp.PropertyType()
		var optionalType *OptionalType
		if optionalType, ok = propType.(*OptionalType); ok {
			propType = optionalType.Element()
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
			// not entirely true since the options could all have distinct sets, but this is good-
			// enough for now
		}

		enumValue := enumOptions[0].Value
		if _, ok := result[enumValue]; ok {
			return nil // if values are not distinct for each member, cannot be used as discriminator
		}

		result[enumValue] = PropertyNameAndType{prop.PropertyName(), propObjTypeName}
	}

	return result
}

func DetermineDiscriminantAndValues(oneOf *ObjectType, definitions TypeDefinitionSet) (string, map[string]PropertyNameAndType) {
	// grab out the first member of the OneOf
	var firstMember *ObjectType
	for _, prop := range oneOf.Properties() {
		_, firstMember = resolveOneOfMemberToObjectType(prop.PropertyType(), definitions)
		break
	}

	// try to find a discriminator property out of the properties on the first member
	for _, prop := range firstMember.Properties() {
		mapping := getDiscriminatorMapping(oneOf, prop.PropertyName(), definitions)
		if mapping != nil {
			jsonTag, ok := prop.Tag("json")
			if !ok {
				// in reality every property will have a JSON tag
				panic("discriminator property had no JSON tag")
			}

			// first part of JSON tag is the JSON name
			return jsonTag[0], mapping
		}
	}

	panic("unable to determine a discriminator property for oneOf type")
}
