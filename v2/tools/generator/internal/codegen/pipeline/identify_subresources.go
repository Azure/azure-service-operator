/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"
	"strings"

	"github.com/pkg/errors"
	"k8s.io/klog/v2"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

const IdentifySubResourcesStageID = "identifySubResources"

// IdentifySubResources replaces cross resource references with genruntime.ResourceReference.
func IdentifySubResources() Stage {
	return MakeLegacyStage(
		IdentifySubResourcesStageID,
		"Replace sub-resources with references",
		func(ctx context.Context, definitions astmodel.Types) (astmodel.Types, error) {
			result := make(astmodel.Types)

			var subResourceErrs []error

			specs := astmodel.NewTypeNameSet()
			for _, def := range definitions {
				if rt, ok := astmodel.AsResourceType(def.Type()); ok {
					if tn, ok := astmodel.AsTypeName(rt.SpecType()); ok {
						specs.Add(tn)
					}
				}
			}

			visitor := MakeSubResourceTypeVisitor()
			for tn, def := range definitions {
				if strings.HasSuffix(def.Name().Name(), "_Status") {
					result.Add(def) // skip status types
					continue
				}

				isSpec := specs.Contains(tn)

				t, err := visitor.Visit(def.Type(), isSpec)
				if err != nil {
					return nil, errors.Wrapf(err, "visiting %q", def.Name())
				}

				result.Add(def.WithType(t))
			}

			if len(subResourceErrs) > 0 {
				return nil, errors.New("\n" + toBulletedErrList(subResourceErrs))
			}

			return result, nil
		})
}

func LooksLikeResource(o *astmodel.ObjectType) bool {
	// NB: See swagger_type_extrator.go:isMarkedAsARMResource
	_, hasType := o.Property("Type")
	_, hasName := o.Property("Name")
	_, hasId := o.Property("Id")

	return hasType && hasName && hasId
}

func MakeSubResourceTypeVisitor() astmodel.TypeVisitor {
	return astmodel.TypeVisitorBuilder{
		VisitObjectType: func(this *astmodel.TypeVisitor, it *astmodel.ObjectType, ctx interface{}) (astmodel.Type, error) {
			isSpec := ctx.(bool)
			if !isSpec && LooksLikeResource(it) {
				klog.Infof("Replaced sub-resource")
				return astmodel.ResourceReferenceType, nil
			}

			return astmodel.IdentityVisitOfObjectType(this, it, false)
		},
	}.Build()
}
