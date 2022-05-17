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

	person2020 := test.CreateSpec(test.Pkg2020, "Person", test.FullNameProperty, test.KnownAsProperty, test.FamilyNameProperty)
	person2021 := test.CreateSpec(test.Pkg2021, "Person", test.FullNameProperty, test.KnownAsProperty, test.FamilyNameProperty)
	person2022 := test.CreateSpec(test.Pkg2022, "Person", test.FullNameProperty, test.KnownAsProperty, test.FamilyNameProperty)

	cfg := config.NewObjectModelConfiguration()

	defs := make(astmodel.TypeDefinitionSet)
	defs.AddAll(person2020, person2021, person2022)

	initialState, err := RunTestPipeline(
		NewState(defs))
	g.Expect(err).To(Succeed())

	finalState, err := RunTestPipeline(
		initialState,
		CreateTypesForBackwardCompatibility("vPrefix", cfg))
	g.Expect(err).To(Succeed())

	test.AssertPackagesGenerateExpectedCode(t, finalState.Definitions())
}
