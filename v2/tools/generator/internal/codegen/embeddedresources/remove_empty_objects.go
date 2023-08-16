/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package embeddedresources

import (
	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	kerrors "k8s.io/apimachinery/pkg/util/errors"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

func RemoveEmptyObjects(
	definitions astmodel.TypeDefinitionSet,
	log logr.Logger,
) (astmodel.TypeDefinitionSet, error) {
	result := definitions
	for {
		toRemove := findEmptyObjectTypes(result, log)
		if len(toRemove) == 0 {
			break
		}

		var err error
		result, err = removeReferencesToTypes(result, toRemove, log)
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

func findEmptyObjectTypes(
	definitions astmodel.TypeDefinitionSet,
	log logr.Logger,
) astmodel.TypeNameSet {
	result := astmodel.NewTypeNameSet()

	for _, def := range definitions {
		ot, ok := astmodel.AsObjectType(def.Type())
		if !ok {
			continue
		}

		// If there's still something "in" the object then we don't want to remove it
		if !ot.Properties().IsEmpty() || len(ot.EmbeddedProperties()) != 0 {
			continue
		}

		if astmodel.DoNotPrune.IsOn(def.Type()) {
			continue
		}

		log.V(1).Info(
			"Removing empty type",
			"name", def.Name())

		result.Add(def.Name())
	}

	return result
}

func removeReferencesToTypes(
	definitions astmodel.TypeDefinitionSet,
	toRemove astmodel.TypeNameSet,
	log logr.Logger,
) (astmodel.TypeDefinitionSet, error) {
	result := make(astmodel.TypeDefinitionSet)
	visitor := makeRemovedTypeVisitor(toRemove, log)

	for _, def := range definitions {
		if toRemove.Contains(def.Name()) {
			continue
		}

		updatedDef, err := visitor.VisitDefinition(def, nil)
		if err != nil {
			return nil, errors.Wrapf(err, "visiting definition %q", def.Name())
		}
		result.Add(updatedDef)
	}

	return result, nil
}

type visitorCtx struct {
	typeName astmodel.TypeName
}

func makeRemovedTypeVisitor(
	toRemove astmodel.TypeNameSet,
	log logr.Logger,
) astmodel.TypeVisitor[*visitorCtx] {
	// This is basically copied from IdentityVisitOfObjectType, but since it has/needs a per-property context we can't use that
	removeReferencesToEmptyTypes := func(
		this *astmodel.TypeVisitor[*visitorCtx],
		it *astmodel.ObjectType,
		ctx *visitorCtx,
	) (astmodel.Type, error) {
		// just map the property types
		var errs []error
		var newProps []*astmodel.PropertyDefinition
		it.Properties().ForEach(func(prop *astmodel.PropertyDefinition) {
			ctx := &visitorCtx{}
			p, err := this.Visit(prop.PropertyType(), ctx)
			if err != nil {
				errs = append(errs, err)
			} else if ctx.typeName == nil || !toRemove.Contains(ctx.typeName) {
				newProps = append(newProps, prop.WithType(p))
			} else if toRemove.Contains(ctx.typeName) {
				log.V(1).Info(
					"Removing reference to empty type",
					"property", prop.PropertyName(),
					"referencing", ctx.typeName)
			}
		})

		if len(errs) > 0 {
			return nil, kerrors.NewAggregate(errs)
		}

		// map the embedded types too
		var newEmbeddedProps []*astmodel.PropertyDefinition
		for _, prop := range it.EmbeddedProperties() {
			ctx := &visitorCtx{}
			p, err := this.Visit(prop.PropertyType(), ctx)
			if err != nil {
				errs = append(errs, err)
			} else if ctx.typeName != nil && !toRemove.Contains(ctx.typeName) {
				newEmbeddedProps = append(newEmbeddedProps, prop.WithType(p))
			}
		}

		if len(errs) > 0 {
			return nil, kerrors.NewAggregate(errs)
		}

		result := it.WithoutProperties()
		result = result.WithProperties(newProps...)
		result = result.WithoutEmbeddedProperties()
		result, err := result.WithEmbeddedProperties(newEmbeddedProps...)
		if err != nil {
			return nil, err
		}

		return result, nil
	}

	typeNameVisitWithSafetyCheck := func(
		this *astmodel.TypeVisitor[*visitorCtx],
		it astmodel.TypeName,
		ctx *visitorCtx,
	) (astmodel.Type, error) {
		// Safety check that we're not overwriting typeName
		if ctx != nil {
			if ctx.typeName != nil {
				return nil, errors.Errorf("would've overwritten ctx.typeName %q", ctx.typeName)
			}

			ctx.typeName = it
		}

		return astmodel.IdentityVisitOfTypeName(this, it, ctx)
	}

	visitor := astmodel.TypeVisitorBuilder[*visitorCtx]{
		VisitObjectType: removeReferencesToEmptyTypes,
		VisitTypeName:   typeNameVisitWithSafetyCheck,
	}.Build()

	return visitor
}
