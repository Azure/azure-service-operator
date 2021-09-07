/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/hack/generator/pkg/astmodel"
	"github.com/Azure/azure-service-operator/hack/generator/pkg/test"
)

func TestInjectHubFunction_WhenResourceIsStorageVersion_GeneratesExpectedFile(t *testing.T) {
	g := NewWithT(t)

	idFactory := astmodel.NewIdentifierFactory()

	// Define a test resource
	spec := test.CreateSpec(test.Pkg2020, "Person", test.FullNameProperty, test.FamilyNameProperty, test.KnownAsProperty)
	status := test.CreateStatus(test.Pkg2020, "Person")
	resource := test.CreateResource(test.Pkg2020, "Person", spec, status)

	resource = resource.WithType(
		resource.Type().(*astmodel.ResourceType).MarkAsStorageVersion())

	types := make(astmodel.Types)
	types.AddAll(resource, status, spec)

	injectHubFunction := InjectHubFunction(idFactory)

	// Don't need a context when testing
	state := NewState().WithTypes(types)
	finalState, err := injectHubFunction.Run(context.TODO(), state)

	g.Expect(err).To(Succeed())

	test.AssertPackagesGenerateExpectedCode(t, finalState.Types())
}

func TestInjectHubFunction_WhenResourceIsNotStorageVersion_GeneratesExpectedFile(t *testing.T) {
	g := NewWithT(t)

	idFactory := astmodel.NewIdentifierFactory()

	// Define a test resource
	spec := test.CreateSpec(test.Pkg2020, "Person", test.FullNameProperty, test.FamilyNameProperty, test.KnownAsProperty)
	status := test.CreateStatus(test.Pkg2020, "Person")
	resource := test.CreateResource(test.Pkg2020, "Person", spec, status)

	types := make(astmodel.Types)
	types.AddAll(resource, status, spec)

	injectHubFunction := InjectHubFunction(idFactory)

	// Don't need a context when testing
	state := NewState().WithTypes(types)
	finalState, err := injectHubFunction.Run(context.TODO(), state)

	g.Expect(err).To(Succeed())

	test.AssertPackagesGenerateExpectedCode(t, finalState.Types())
}
