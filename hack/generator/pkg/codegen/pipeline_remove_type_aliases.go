/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package codegen

import (
	"context"
	"fmt"

	kerrors "k8s.io/apimachinery/pkg/util/errors"

	"github.com/Azure/k8s-infra/hack/generator/pkg/astmodel"
	"k8s.io/klog/v2"
)

// removeTypeAliases creates a pipeline stage removing type aliases
func removeTypeAliases() PipelineStage {
	return MakePipelineStage(
		"removeAliases",
		"Remove type aliases",
		func(ctx context.Context, types astmodel.Types) (astmodel.Types, error) {
			visitor := astmodel.MakeTypeVisitor()
			var result = make(astmodel.Types)

			visitor.VisitTypeName = func(this *astmodel.TypeVisitor, it astmodel.TypeName, ctx interface{}) (astmodel.Type, error) {
				return resolveTypeName(this, it, types)
			}

			var errs []error
			for _, typeDef := range types {
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

func resolveTypeName(visitor *astmodel.TypeVisitor, name astmodel.TypeName, types astmodel.Types) (astmodel.Type, error) {
	def, ok := types[name]
	if !ok {
		panic(fmt.Sprintf("Couldn't find type for type name %s", name))
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
	case astmodel.ValidatedType:
		return def.Name(), nil // must remain named so there is somewhere to put validations
	case *astmodel.FlaggedType:
		return def.Name(), nil // must remain named as it is just wrapping objectType (and objectType remains named)
	case astmodel.TypeName:
		// We need to resolve further because this type is an alias
		klog.V(3).Infof("Found type alias %s, replacing it with %s", name, concreteType)
		return resolveTypeName(visitor, concreteType, types)
	case *astmodel.PrimitiveType:
		return visitor.Visit(concreteType, nil)
	case *astmodel.OptionalType:
		return visitor.Visit(concreteType, nil)
	case *astmodel.ArrayType:
		return visitor.Visit(concreteType, nil)
	case *astmodel.MapType:
		return visitor.Visit(concreteType, nil)
	default:
		panic(fmt.Sprintf("Don't know how to resolve type %T for typeName %s", concreteType, name))
	}
}
