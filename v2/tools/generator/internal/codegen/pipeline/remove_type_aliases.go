/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"
	"fmt"

	kerrors "k8s.io/apimachinery/pkg/util/errors"

	"github.com/pkg/errors"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

// RemoveTypeAliasesStageID is the unique identifier for this pipeline stage
const RemoveTypeAliasesStageID = "removeAliases"

// RemoveTypeAliases creates a pipeline stage removing type aliases
func RemoveTypeAliases() *Stage {
	return NewLegacyStage(
		RemoveTypeAliasesStageID,
		"Remove type aliases",
		func(ctx context.Context, definitions astmodel.TypeDefinitionSet) (astmodel.TypeDefinitionSet, error) {
			simplifyAliases := func(this *astmodel.TypeVisitor, it astmodel.TypeName, ctx interface{}) (astmodel.Type, error) {
				return resolveTypeName(this, it, definitions)
			}

			visitor := astmodel.TypeVisitorBuilder{
				VisitTypeName: simplifyAliases,
			}.Build()

			result := make(astmodel.TypeDefinitionSet)

			var errs []error
			for _, typeDef := range definitions {
				visitedType, err := visitor.Visit(typeDef.Type(), nil)
				if err != nil {
					errs = append(errs, err)
				} else {
					result.Add(typeDef.WithType(visitedType))
				}
			}

			if len(errs) > 0 {
				return nil, kerrors.NewAggregate(errs)
			}

			return result, nil
		})
}

func resolveTypeName(visitor *astmodel.TypeVisitor, name astmodel.TypeName, definitions astmodel.TypeDefinitionSet) (astmodel.Type, error) {
	// Don't try to remove external refs
	if _, _, ok := name.PackageReference().TryGroupVersion(); !ok {
		return name, nil
	}

	def, ok := definitions[name]
	if !ok {
		return nil, errors.Errorf("couldn't find definition for type name %s", name)
	}

	// If this typeName definition has a type of object, enum, validated, flagged, resource, or resourceList
	// it's okay. Everything else we want to pull up one level to remove the alias
	switch concreteType := def.Type().(type) {
	case *astmodel.ObjectType:
		return def.Name(), nil // must remain named for controller-gen
	case *astmodel.EnumType:
		return def.Name(), nil // must remain named so there is somewhere to put validations
	case *astmodel.ResourceType:
		return def.Name(), nil // must remain named for controller-gen
	case *astmodel.ValidatedType:
		return def.Name(), nil // must remain named so there is somewhere to put validations
	case *astmodel.FlaggedType:
		return def.Name(), nil // must remain named as it is just wrapping objectType (and objectType remains named)
	case *astmodel.InterfaceType:
		return def.Name(), nil // must remain named
	case astmodel.TypeName:
		// We need to resolve further because this type is an alias
		return resolveTypeName(visitor, concreteType, definitions)
	case *astmodel.PrimitiveType:
		return visitor.Visit(concreteType, nil)
	case *astmodel.OptionalType:
		return visitor.Visit(concreteType, nil)
	case *astmodel.ArrayType:
		return visitor.Visit(concreteType, nil)
	case *astmodel.MapType:
		return visitor.Visit(concreteType, nil)
	case *astmodel.ErroredType:
		return visitor.Visit(concreteType, nil)
	default:
		panic(fmt.Sprintf("Don't know how to resolve type %T for typeName %s", concreteType, name))
	}
}
