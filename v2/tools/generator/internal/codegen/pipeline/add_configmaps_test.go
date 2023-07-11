/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/config"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/test"
)

func TestAddConfigMaps_AddsSpecWithRequiredConfigMaps(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// Define a test resource
	spec := test.CreateSpec(
		test.Pkg2020,
		"Person",
		test.FullNameProperty,
		test.FamilyNameProperty,
		test.KnownAsProperty,
		test.RestrictedNameProperty)
	status := test.CreateStatus(test.Pkg2020, "Person")
	resource := test.CreateResource(test.Pkg2020, "Person", spec, status)

	defs := make(astmodel.TypeDefinitionSet)
	defs.AddAll(resource, status, spec)

	omc := config.NewObjectModelConfiguration()
	g.Expect(
		omc.ModifyProperty(
			spec.Name(),
			test.FullNameProperty.PropertyName(),
			func(pc *config.PropertyConfiguration) error {
				pc.ImportConfigMapMode.Set(config.ImportConfigMapModeRequired)
				return nil
			})).
		To(Succeed())
	g.Expect(
		omc.ModifyProperty(
			spec.Name(),
			test.FamilyNameProperty.PropertyName(),
			func(pc *config.PropertyConfiguration) error {
				pc.ImportConfigMapMode.Set(config.ImportConfigMapModeOptional)
				return nil
			})).
		To(Succeed())
	g.Expect(
		omc.ModifyProperty(
			spec.Name(),
			test.RestrictedNameProperty.PropertyName(),
			func(pc *config.PropertyConfiguration) error {
				pc.ImportConfigMapMode.Set(config.ImportConfigMapModeOptional)
				return nil
			})).
		To(Succeed())

	configuration := config.NewConfiguration()
	configuration.ObjectModelConfiguration = omc

	addConfigMaps := AddConfigMaps(configuration)

	// Don't need a context when testing
	state := NewState().WithDefinitions(defs)
	finalState, err := addConfigMaps.Run(context.TODO(), state)

	g.Expect(err).To(Succeed())

	test.AssertPackagesGenerateExpectedCode(t, finalState.Definitions())
}
