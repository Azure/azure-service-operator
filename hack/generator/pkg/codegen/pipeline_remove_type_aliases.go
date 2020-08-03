/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package codegen

import (
	"context"
	"fmt"

	"github.com/Azure/k8s-infra/hack/generator/pkg/astmodel"
	"k8s.io/klog/v2"
)

// removeTypeAliases creates a pipeline stage removing type aliases
func removeTypeAliases() PipelineStage {
	return PipelineStage{
		"Remove type aliases",
		func(ctx context.Context, types astmodel.Types) (astmodel.Types, error) {
			visitor := astmodel.MakeTypeVisitor()
			var result = make(astmodel.Types)

			visitor.VisitTypeName = func(this *astmodel.TypeVisitor, it astmodel.TypeName, ctx interface{}) astmodel.Type {
				resolvedName := resolveTypeName(this, it, types)
				return resolvedName
			}

			for _, typeDef := range types {
				newType := visitor.Visit(typeDef.Type(), nil)
				result.Add(typeDef.WithType(newType))
			}

			return result, nil
		}}
}

func resolveTypeName(visitor *astmodel.TypeVisitor, name astmodel.TypeName, types astmodel.Types) astmodel.Type {
	def, ok := types[name]
	if !ok {
		panic(fmt.Sprintf("Couldn't find type for type name %s", name))
	}

	// If this typeName definition has a type of object, enum, or resource
	// it's okay. Everything else we want to pull up one level to remove the alias
	switch concreteType := def.Type().(type) {
	case *astmodel.ObjectType:
		return def.Name()
	case *astmodel.EnumType:
		return def.Name()
	case *astmodel.ResourceType:
		return def.Name()
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
