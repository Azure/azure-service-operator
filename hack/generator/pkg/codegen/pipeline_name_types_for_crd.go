/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package codegen

import (
	"context"

	"github.com/Azure/k8s-infra/hack/generator/pkg/astmodel"
)

// nameTypesForCRD - for CRDs all inner enums and structs must be named, so we do it here
func nameTypesForCRD(idFactory astmodel.IdentifierFactory) PipelineStage {

	return PipelineStage{
		Name: "expand inner types",
		Action: func(ctx context.Context, types Types) (Types, error) {

			result := make(Types)

			// this is a little bit of a hack, better way to do it?
			getDescription := func(typeName *astmodel.TypeName) *string {
				if typeDef, ok := types[*typeName]; ok {
					return typeDef.Description()
				}

				return nil
			}

			for _, typeDef := range types {
				newDefs := handleType(typeDef, idFactory, getDescription)
				for _, newDef := range newDefs {
					result[*newDef.Name()] = newDef
				}
			}

			return result, nil
		},
	}
}

func handleType(def astmodel.TypeDefinition, idFactory astmodel.IdentifierFactory, getDescription func(*astmodel.TypeName) *string) []astmodel.TypeDefinition {

	var resultTypes []astmodel.TypeDefinition

	// for this visitor, we will pass around the name hint as 'ctx' parameter
	visitor := astmodel.MakeTypeVisitor()

	visitor.VisitEnumType = func(this *astmodel.TypeVisitor, it *astmodel.EnumType, ctx interface{}) astmodel.Type {
		nameHint := ctx.(string)
		enumName := astmodel.NewTypeName(def.Name().PackageReference, idFactory.CreateEnumIdentifier(nameHint))

		namedEnum := astmodel.MakeTypeDefinition(enumName, it)
		namedEnum = namedEnum.WithDescription(getDescription(enumName))

		resultTypes = append(resultTypes, namedEnum)

		return namedEnum.Name()
	}

	visitor.VisitStructType = func(this *astmodel.TypeVisitor, it *astmodel.StructType, ctx interface{}) astmodel.Type {
		nameHint := ctx.(string)

		var props []*astmodel.PropertyDefinition
		// first map the inner types:
		for _, prop := range it.Properties() {
			newPropType := this.Visit(prop.PropertyType(), nameHint+string(prop.PropertyName()))
			props = append(props, prop.WithType(newPropType))
		}

		structName := astmodel.NewTypeName(def.Name().PackageReference, nameHint)

		namedStruct := astmodel.MakeTypeDefinition(structName, it.WithProperties(props...))
		namedStruct = namedStruct.WithDescription(getDescription(structName))

		resultTypes = append(resultTypes, namedStruct)

		return namedStruct.Name()
	}

	visitor.VisitResourceType = func(this *astmodel.TypeVisitor, it *astmodel.ResourceType, ctx interface{}) astmodel.Type {
		nameHint := ctx.(string)

		spec := this.Visit(it.SpecType(), nameHint+"Spec")

		var status astmodel.Type // the type is very important here, it must be a nil(Type) if status isn’t set, not a nil(*TypeName)
		if it.StatusType() != nil {
			status = this.Visit(it.StatusType(), nameHint+"Status")
		}

		resourceName := astmodel.NewTypeName(def.Name().PackageReference, nameHint)
		
		resource := astmodel.MakeTypeDefinition(resourceName, astmodel.NewResourceType(spec, status))
		resource = resource.WithDescription(getDescription(resourceName))

		resultTypes = append(resultTypes, resource)

		return resource.Name()
	}

	_ = visitor.Visit(def.Type(), def.Name().Name())

	return resultTypes
}
