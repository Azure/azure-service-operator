/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

const UserAssignedIdentityTypeDescription = "Information about the user assigned identity for the resource"

const HandleUserAssignedIdentitiesStageID = "handleUserAssignedIdentities"

func HandleUserAssignedIdentities() *Stage {
	stage := NewStage(
		HandleUserAssignedIdentitiesStageID,
		"Transform UserAssignedIdentities on spec types be resource references with the expected shape",
		func(ctx context.Context, state *State) (*State, error) {

			transformer := newUserAssignedIdentityTransformer()

			updatedDefs := make(astmodel.TypeDefinitionSet)
			for _, def := range state.Definitions() {
				updatedDef, err := transformer.visitor.VisitDefinition(def, def.Name())
				if err != nil {
					return nil, err
				}

				updatedDefs.Add(updatedDef)
			}

			// Now add all of the new types we need to
			updatedDefs.AddTypes(transformer.typesToAdd)

			return state.WithDefinitions(updatedDefs), nil
		})

	return stage
}

type userAssignedIdentityTransformer struct {
	visitor    astmodel.TypeVisitor
	typesToAdd astmodel.TypeDefinitionSet
}

func newUserAssignedIdentityTransformer() *userAssignedIdentityTransformer {
	result := &userAssignedIdentityTransformer{
		typesToAdd: make(astmodel.TypeDefinitionSet),
	}

	visitor := astmodel.TypeVisitorBuilder{
		VisitObjectType: result.transformUserAssignedIdentityProperty,
	}.Build()
	result.visitor = visitor

	return result
}

func (t *userAssignedIdentityTransformer) transformUserAssignedIdentityProperty(this *astmodel.TypeVisitor, it *astmodel.ObjectType, ctx interface{}) (astmodel.Type, error) {
	name := ctx.(astmodel.InternalTypeName)

	// Doesn't apply to status types
	if name.IsStatus() {
		return astmodel.IdentityVisitOfObjectType(this, it, ctx)
	}

	prop, ok := it.Properties().Get(astmodel.UserAssignedIdentitiesProperty)
	if !ok {
		return astmodel.IdentityVisitOfObjectType(this, it, ctx)
	}

	mapType, ok := astmodel.AsMapType(prop.PropertyType())
	if !ok {
		return astmodel.IdentityVisitOfObjectType(this, it, ctx)
	}

	// The map should be of string -> Object, TypeName, or AnyType.
	primitiveKeyType, ok := astmodel.AsPrimitiveType(mapType.KeyType())
	if !ok || primitiveKeyType != astmodel.StringType {
		return astmodel.IdentityVisitOfObjectType(this, it, ctx)
	}

	_, isValueTypeName := astmodel.AsTypeName(mapType.ValueType())
	_, isValueObject := astmodel.AsObjectType(mapType.ValueType())
	isValueAny := astmodel.TypeEquals(mapType.ValueType(), astmodel.AnyType)
	isValueMapOfAny := astmodel.TypeEquals(mapType.ValueType(), astmodel.NewMapType(astmodel.StringType, astmodel.AnyType))
	if !isValueTypeName && !isValueObject && !isValueAny && !isValueMapOfAny {
		return astmodel.IdentityVisitOfObjectType(this, it, ctx)
	}

	// Replace the map with an array of structs. This is required because
	// the key of the map is an ARM ID which we would like to eventually
	// generate a genruntime.ResourceReference, but you can't have a struct be a map
	// key in CRDs, so we have to special case this property transform it to the correct shape
	// during ARM serialization

	userAssignedIdentityDef := newUserAssignedIdentityDefinition(name.PackageReference())
	err := t.typesToAdd.AddAllowDuplicates(userAssignedIdentityDef)
	if err != nil {
		return nil, err
	}

	prop = prop.WithType(astmodel.NewArrayType(userAssignedIdentityDef.Name()))
	it = it.WithProperty(prop)

	return astmodel.IdentityVisitOfObjectType(this, it, ctx)
}

func newUserAssignedIdentityDefinition(pr astmodel.PackageReference) astmodel.TypeDefinition {
	name := astmodel.MakeInternalTypeName(pr, astmodel.UserAssignedIdentitiesTypeName)

	prop := astmodel.NewPropertyDefinition("Reference", "reference", astmodel.ResourceReferenceType)

	// This type is special but we have to set an ARMReferenceTag anyway for downstream consumers to be happy.
	prop = prop.WithTag(astmodel.ARMReferenceTag, "Reference")
	t := astmodel.NewObjectType().WithProperty(prop)

	return astmodel.MakeTypeDefinition(name, t).WithDescription(UserAssignedIdentityTypeDescription)
}
