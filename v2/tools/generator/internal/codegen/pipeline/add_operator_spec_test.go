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

	types := make(astmodel.TypeDefinitionSet)
	types.AddAll(resource, status, spec)

	idFactory := astmodel.NewIdentifierFactory()
	omc, tc := config.CreateTestObjectModelConfiguration(resource.Name())
	tc.SetAzureGeneratedSecrets([]string{"key1"})

	configuration := config.NewConfiguration()
	configuration.ObjectModelConfiguration = omc

	addOperatorSpec := AddOperatorSpec(configuration, idFactory)

	// Don't need a context when testing
	state := NewState().WithTypes(types)
	finalState, err := addOperatorSpec.Run(context.TODO(), state)

	g.Expect(err).To(Succeed())

	test.AssertPackagesGenerateExpectedCode(t, finalState.Types())
}
