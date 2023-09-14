/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"testing"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/config"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/test"

	. "github.com/onsi/gomega"
)

func TestGolden_CreateBackwardCompatibilityTypes(t *testing.T) {

	t.Parallel()
	g := NewGomegaWithT(t)

	// Arrange - create a resource with a status and spec

	spec := test.CreateSpec(
		test.Pkg2020,
		"Person",
		test.FullNameProperty,
		test.KnownAsProperty,
		test.FamilyNameProperty,
	)
	status := test.CreateStatus(
		test.Pkg2020,
		"Person",
		test.FullNameProperty,
		test.KnownAsProperty,
		test.FamilyNameProperty,
		test.StatusProperty,
	)

	person := test.CreateResource(test.Pkg2020, "Person", spec, status)

	// Arrange - set up configuration where we expect this resource needs a compatibility type

	cfg := config.NewObjectModelConfiguration()
	g.Expect(cfg.ModifyType(
		person.Name(),
		func(t *config.TypeConfiguration) error {
			t.SupportedFrom.Set("v2.0.0-beta.1")
			return nil
		})).To(Succeed())

	// Arrange - Create our initial state
	defs := make(astmodel.TypeDefinitionSet)
	defs.AddAll(person, spec, status)

	initialState, err := RunTestPipeline(
		NewState(defs))
	g.Expect(err).To(Succeed())

	// Act

	finalState, err := RunTestPipeline(
		initialState,
		CreateTypesForBackwardCompatibility("v2.0.0-", cfg))
	g.Expect(err).To(Succeed())

	// Assert

	g.Expect(finalState).To(Not(BeNil()))
	g.Expect(finalState.Definitions()).To(HaveLen(6))

	test.AssertPackagesGenerateExpectedCode(t, finalState.Definitions())
}
