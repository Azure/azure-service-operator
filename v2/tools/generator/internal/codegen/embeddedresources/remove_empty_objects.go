/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package embeddedresources

import (
	"github.com/pkg/errors"
	kerrors "k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/klog/v2"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

func RemoveEmptyObjects(definitions astmodel.TypeDefinitionSet) (astmodel.TypeDefinitionSet, error) {
	result := definitions
	for {
		toRemove := findEmptyObjectTypes(result)
		if len(toRemove) == 0 {
			break
		}

		var err error
		result, err = removeReferencesToTypes(result, toRemove)
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

func findEmptyObjectTypes(definitions astmodel.TypeDefinitionSet) astmodel.TypeNameSet {
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

		klog.V(4).Infof("Removing %q as it has no properties", def.Name())
		result.Add(def.Name())
	}

	return result
}

func removeReferencesToTypes(definitions astmodel.TypeDefinitionSet, toRemove astmodel.TypeNameSet) (astmodel.TypeDefinitionSet, error) {
	result := make(astmodel.TypeDefinitionSet)
	visitor := makeRemovedTypeVisitor(toRemove)

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
	typeName *astmodel.TypeName
}

func makeRemovedTypeVisitor(toRemove astmodel.TypeNameSet) astmodel.TypeVisitor {
	// This is basically copied from IdentityVisitOfObjectType, but since it has/needs a per-property context we can't use that
	removeReferencesToEmptyTypes := func(this *astmodel.TypeVisitor, it *astmodel.ObjectType, ctx interface{}) (astmodel.Type, error) {
		// just map the property types
		var errs []error
		var newProps []*astmodel.PropertyDefinition
		it.Properties().ForEach(func(prop *astmodel.PropertyDefinition) {
			ctx := &visitorCtx{}
			p, err := this.Visit(prop.PropertyType(), ctx)
			if err != nil {
				errs = append(errs, err)
			} else if ctx.typeName == nil || !toRemove.Contains(*ctx.typeName) {
				newProps = append(newProps, prop.WithType(p))
			} else if toRemove.Contains(*ctx.typeName) {
				klog.V(4).Infof("Removing property %q (referencing %q) as the type has no properties", prop.PropertyName(), *ctx.typeName)
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
			} else if ctx.typeName != nil && !toRemove.Contains(*ctx.typeName) {
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

	typeNameVisitWithSafetyCheck := func(this *astmodel.TypeVisitor, it astmodel.TypeName, ctx interface{}) (astmodel.Type, error) {
		typedCtx, ok := ctx.(*visitorCtx)
		if ok {
			// Safety check that we're not overwriting typeName
			if typedCtx.typeName != nil {
				return nil, errors.Errorf("would've overwritten ctx.typeName %q", typedCtx.typeName)
			}
			typedCtx.typeName = &it
		}
		return astmodel.IdentityVisitOfTypeName(this, it, ctx)
	}

	visitor := astmodel.TypeVisitorBuilder{
		VisitObjectType: removeReferencesToEmptyTypes,
		VisitTypeName:   typeNameVisitWithSafetyCheck,
	}.Build()

	return visitor
}
