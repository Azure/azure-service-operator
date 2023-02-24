/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/config"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/test"
)

func TestCreateConversionGraph(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	cfg := config.NewConfiguration()

	person2020 := test.CreateSpec(test.Pkg2020, "Person", test.FullNameProperty, test.KnownAsProperty, test.FamilyNameProperty)
	person2021 := test.CreateSpec(test.Pkg2021, "Person", test.FullNameProperty, test.KnownAsProperty, test.FamilyNameProperty)
	person2022 := test.CreateSpec(test.Pkg2022, "Person", test.FullNameProperty, test.KnownAsProperty, test.FamilyNameProperty)

	defs := make(astmodel.TypeDefinitionSet)
	defs.AddAll(person2020, person2021, person2022)

	initialState, err := RunTestPipeline(
		NewState(defs),
		CreateStorageTypes())
	g.Expect(err).To(Succeed())

	finalState, err := RunTestPipeline(
		initialState,
		CreateConversionGraph(cfg, "v"))
	g.Expect(err).To(Succeed())

	g.Expect(finalState.Definitions()).To(HaveLen(6))
	g.Expect(finalState.ConversionGraph()).NotTo(BeNil())

	graph := finalState.ConversionGraph()

	// Expect to have a link from Pkg2020 to a matching storage version
	storage2020 := graph.LookupTransition(person2020.Name())
	g.Expect(storage2020.String()).To(ContainSubstring(test.Pkg2020.Version()))

	// Expect to have a link from Pkg2021 to a matching storage version
	storage2021 := graph.LookupTransition(person2021.Name())
	g.Expect(storage2021.String()).To(ContainSubstring(test.Pkg2021.Version()))

	// Expect to have a link from Pkg2022 to a matching storage version
	storage2022 := graph.LookupTransition(person2022.Name())
	g.Expect(storage2022.String()).To(ContainSubstring(test.Pkg2022.Version()))

	// Expect to have a link from Storage2020 to Storage2021
	linkedFrom2020 := graph.LookupTransition(storage2020)
	g.Expect(linkedFrom2020).To(Equal(storage2021))

	// Expect to have a link from Storage2021 version Storage2022
	linkedFrom2021 := graph.LookupTransition(storage2021)
	g.Expect(linkedFrom2021).To(Equal(storage2022))

	// Expect NOT to have a link from Storage2022
	missing := graph.LookupTransition(storage2022)
	g.Expect(missing.IsEmpty()).To(BeTrue())

	// Finally, check that the five links we've verified are the only ones there
	// We check this last, so that a failure doesn't block more detailed checks
	g.Expect(graph.TransitionCount()).To(Equal(5))
}
