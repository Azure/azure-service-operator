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

func TestAddOperatorSpec_AddsSpecWithConfiguredSecrets(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// Define a test resource
	spec := test.CreateSpec(test.Pkg2020, "Person", test.FullNameProperty, test.FamilyNameProperty, test.KnownAsProperty)
	status := test.CreateStatus(test.Pkg2020, "Person")
	resource := test.CreateResource(test.Pkg2020, "Person", spec, status)

	defs := make(astmodel.TypeDefinitionSet)
	defs.AddAll(resource, status, spec)

	idFactory := astmodel.NewIdentifierFactory()
	omc := config.NewObjectModelConfiguration()
	g.Expect(
		omc.ModifyType(
			resource.Name(),
			func(tc *config.TypeConfiguration) error {
				tc.SetAzureGeneratedSecrets([]string{"key1"})
				return nil
			})).
		To(Succeed())

	configuration := config.NewConfiguration()
	configuration.ObjectModelConfiguration = omc

	addOperatorSpec := AddOperatorSpec(configuration, idFactory)

	// Don't need a context when testing
	state := NewState().WithDefinitions(defs)
	finalState, err := addOperatorSpec.Run(context.TODO(), state)

	g.Expect(err).To(Succeed())

	test.AssertPackagesGenerateExpectedCode(t, finalState.Definitions())
}
