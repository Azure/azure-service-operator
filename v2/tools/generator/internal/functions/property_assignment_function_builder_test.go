// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package functions

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/config"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/conversions"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/test"
)

func TestPropertyAssignmentFunctionBuilder_ManualStrategyDelegatesPropertyBagAccess(t *testing.T) {
	t.Parallel()

	for _, missing := range []bool{false, true} {
		for _, direction := range []conversions.Direction{conversions.ConvertFrom, conversions.ConvertTo} {
			name := direction.SelectString("from", "to")
			if missing {
				name += "PropertyBag"
			}

			t.Run(name, func(t *testing.T) {
				t.Parallel()
				g := NewGomegaWithT(t)

				current := test.CreateObjectDefinition(
					test.Pkg2020, "Person",
					test.FullNameProperty, test.FamilyNameProperty, test.PropertyBagProperty,
				)
				next := test.CreateObjectDefinition(
					test.Pkg2021, "Person",
					test.FamilyNameProperty, test.PropertyBagProperty,
				)
				if !missing {
					object, ok := astmodel.AsObjectType(next.Type())
					g.Expect(ok).To(BeTrue())
					next = next.WithType(object.WithProperty(test.FullNameProperty))
				}

				cfg := config.NewObjectModelConfiguration()
				g.Expect(cfg.ModifyProperty(current.Name(), test.FullNameProperty.PropertyName(),
					func(pc *config.PropertyConfiguration) error {
						pc.ConversionStrategy.Set(config.ConversionStrategyManual)
						return nil
					})).To(Succeed())

				ctx := conversions.NewPropertyConversionContext(
					conversions.AssignPropertiesMethodPrefix,
					astmodel.MakeTypeDefinitionSetFromDefinitions(current, next),
					astmodel.NewIdentifierFactory(),
				).WithConfiguration(cfg)

				builder := NewPropertyAssignmentFunctionBuilder(current, next, direction)
				fn, err := builder.Build(ctx)
				g.Expect(err).NotTo(HaveOccurred())
				if missing {
					g.Expect(fn.conversions).NotTo(HaveKey(string(test.FullNameProperty.PropertyName())))
				} else {
					g.Expect(fn.conversions).To(HaveKey(string(test.FullNameProperty.PropertyName())))
				}
				g.Expect(fn.conversions).To(HaveKey(string(test.FamilyNameProperty.PropertyName())))
			})
		}
	}
}

func TestPropertyAssignmentFunctionBuilder_ManualConversionStillInitializesFromStatus(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	spec := test.CreateObjectDefinition(test.Pkg2020, "Person", test.FullNameProperty)
	status := test.CreateObjectDefinition(test.Pkg2020, "Person_STATUS", test.FullNameProperty)
	cfg := config.NewObjectModelConfiguration()
	g.Expect(cfg.ModifyProperty(spec.Name(), test.FullNameProperty.PropertyName(),
		func(pc *config.PropertyConfiguration) error {
			pc.ConversionStrategy.Set(config.ConversionStrategyManual)
			return nil
		})).To(Succeed())

	ctx := conversions.NewPropertyConversionContext(
		conversions.InitializationMethodPrefix,
		astmodel.MakeTypeDefinitionSetFromDefinitions(spec, status),
		astmodel.NewIdentifierFactory(),
	).WithConfiguration(cfg)

	builder := NewPropertyAssignmentFunctionBuilder(spec, status, conversions.ConvertFrom)
	fn, err := builder.Build(ctx)
	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(fn.conversions).To(HaveKey(string(test.FullNameProperty.PropertyName())))
}
