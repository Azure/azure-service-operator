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
	"github.com/Azure/azure-service-operator/hack/generator/pkg/codegen/storage"
	"github.com/Azure/azure-service-operator/hack/generator/pkg/functions"
	"github.com/Azure/azure-service-operator/hack/generator/pkg/test"
)

func TestInjectOriginalVersionProperty_InjectsIntoSpec(t *testing.T) {
	g := NewGomegaWithT(t)

	// Define a test resource
	spec := test.CreateSpec(pkg2020, "Person", fullNameProperty, familyNameProperty, knownAsProperty)
	status := test.CreateStatus(pkg2020, "Person")
	resource := test.CreateResource(pkg2020, "Person", spec, status)

	types := make(astmodel.Types)
	types.AddAll(resource, status, spec)

	injectOriginalProperty := InjectOriginalVersionProperty()

	// Don't need a context when testing
	finalTypes, err := injectOriginalProperty.Run(context.TODO(), types)

	g.Expect(err).To(Succeed())

	test.AssertPackagesGenerateExpectedCode(t, finalTypes, t.Name())
}

func TestInjectOriginalVersionProperty_WhenOriginalVersionFunctionFound_DoesNotInjectIntoSpec(t *testing.T) {
	g := NewGomegaWithT(t)

	idFactory := astmodel.NewIdentifierFactory()
	fnInjector := storage.NewFunctionInjector()

	// Define a test resource
	spec := test.CreateSpec(pkg2020, "Person", fullNameProperty, familyNameProperty, knownAsProperty)
	status := test.CreateStatus(pkg2020, "Person")
	resource := test.CreateResource(pkg2020, "Person", spec, status)

	spec, err := fnInjector.Inject(spec, functions.NewOriginalVersionFunction(idFactory))
	g.Expect(err).To(Succeed())

	types := make(astmodel.Types)
	types.AddAll(resource, status, spec)

	injectOriginalProperty := InjectOriginalVersionProperty()

	// Don't need a context when testing
	finalTypes, err := injectOriginalProperty.Run(context.TODO(), types)

	g.Expect(err).To(Succeed())

	test.AssertPackagesGenerateExpectedCode(t, finalTypes, t.Name())
}
