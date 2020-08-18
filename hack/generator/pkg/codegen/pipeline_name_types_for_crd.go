/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package codegen

import (
	"context"

	"github.com/Azure/k8s-infra/hack/generator/pkg/astmodel"
)

// nameTypesForCRD - for CRDs all inner enums and objects must be named, so we do it here
func nameTypesForCRD(idFactory astmodel.IdentifierFactory) PipelineStage {

	return MakePipelineStage(
		"nameTypes",
		"Name inner types for CRD",
		func(ctx context.Context, types astmodel.Types) (astmodel.Types, error) {

			result := make(astmodel.Types)

			// this is a little bit of a hack, better way to do it?
			getDescription := func(typeName astmodel.TypeName) []string {
				if typeDef, ok := types[typeName]; ok {
					return typeDef.Description()
				}

				return []string{}
			}

			for typeName, typeDef := range types {

				newDefs := nameInnerTypes(typeDef, idFactory, getDescription)
				for _, newDef := range newDefs {
					result.Add(newDef)
				}

				if _, ok := result[typeName]; !ok {
					// if we didn’t regenerate the “input” type in nameInnerTypes then it won’t
					// have been added to the output; do it here
					result.Add(typeDef)
				}
			}

			return result, nil
		})
}

func nameInnerTypes(
	def astmodel.TypeDefinition,
	idFactory astmodel.IdentifierFactory,
	getDescription func(typeName astmodel.TypeName) []string) []astmodel.TypeDefinition {

	var resultTypes []astmodel.TypeDefinition

	// for this visitor, we will pass around the name hint as 'ctx' parameter
	visitor := astmodel.MakeTypeVisitor()

	visitor.VisitEnumType = func(this *astmodel.TypeVisitor, it *astmodel.EnumType, ctx interface{}) astmodel.Type {
		nameHint := ctx.(string)

		enumName := astmodel.MakeTypeName(def.Name().PackageReference, idFactory.CreateEnumIdentifier(nameHint))

		namedEnum := astmodel.MakeTypeDefinition(enumName, it)
		namedEnum = namedEnum.WithDescription(getDescription(enumName))

		resultTypes = append(resultTypes, namedEnum)

		return namedEnum.Name()
	}

	visitor.VisitObjectType = func(this *astmodel.TypeVisitor, it *astmodel.ObjectType, ctx interface{}) astmodel.Type {
		nameHint := ctx.(string)

		var props []*astmodel.PropertyDefinition
		// first map the inner types:
		for _, prop := range it.Properties() {
			newPropType := this.Visit(prop.PropertyType(), nameHint+"_"+string(prop.PropertyName()))
			props = append(props, prop.WithType(newPropType))
		}

		objectName := astmodel.MakeTypeName(def.Name().PackageReference, nameHint)

		namedObjectType := astmodel.MakeTypeDefinition(objectName, it.WithProperties(props...))
		namedObjectType = namedObjectType.WithDescription(getDescription(objectName))

		resultTypes = append(resultTypes, namedObjectType)

		return namedObjectType.Name()
	}

	visitor.VisitResourceType = func(this *astmodel.TypeVisitor, it *astmodel.ResourceType, ctx interface{}) astmodel.Type {
		nameHint := ctx.(string)

		spec := this.Visit(it.SpecType(), nameHint+"_Spec")

		var status astmodel.Type
		if it.StatusType() != nil {
			status = this.Visit(it.StatusType(), nameHint+"_Status")
		}

		resourceName := astmodel.MakeTypeName(def.Name().PackageReference, nameHint)

		// TODO: Should we have some better "clone" sort of thing in resource?
		newResource := astmodel.NewResourceType(spec, status).WithOwner(it.Owner())
		resource := astmodel.MakeTypeDefinition(resourceName, newResource)
		resource = resource.WithDescription(getDescription(resourceName))

		resultTypes = append(resultTypes, resource)

		return resource.Name()
	}

	_ = visitor.Visit(def.Type(), def.Name().Name())

	return resultTypes
}
