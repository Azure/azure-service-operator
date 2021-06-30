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

func TestInjectOriginalGVKFunction(t *testing.T) {
	g := NewGomegaWithT(t)

	idFactory := astmodel.NewIdentifierFactory()

	// Define a test resource
	spec := test.CreateSpec(pkg2020, "Person", fullNameProperty, familyNameProperty, knownAsProperty)
	status := test.CreateStatus(pkg2020, "Person")
	resource := test.CreateResource(pkg2020, "Person", spec, status)

	types := make(astmodel.Types)
	types.AddAll(resource, status, spec)

	injectOriginalVersion := InjectOriginalGVKFunction(idFactory)

	// Don't need a context when testing
	finalTypes, err := injectOriginalVersion.Run(context.TODO(), types)

	g.Expect(err).To(Succeed())

	test.AssertPackagesGenerateExpectedCode(t, finalTypes, t.Name())
}
