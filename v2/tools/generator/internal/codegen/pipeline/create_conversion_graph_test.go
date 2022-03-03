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

func TestCreateConversionGraph(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	person2020 := test.CreateSpec(test.Pkg2020, "Person", test.FullNameProperty, test.KnownAsProperty, test.FamilyNameProperty)
	person2021 := test.CreateSpec(test.Pkg2021, "Person", test.FullNameProperty, test.KnownAsProperty, test.FamilyNameProperty)
	person2022 := test.CreateSpec(test.Pkg2022, "Person", test.FullNameProperty, test.KnownAsProperty, test.FamilyNameProperty)

	person2020s := test.CreateSpec(test.Pkg2020s, "Person", test.FullNameProperty, test.KnownAsProperty, test.FamilyNameProperty)
	person2021s := test.CreateSpec(test.Pkg2021s, "Person", test.FullNameProperty, test.KnownAsProperty, test.FamilyNameProperty)
	person2022s := test.CreateSpec(test.Pkg2022s, "Person", test.FullNameProperty, test.KnownAsProperty, test.FamilyNameProperty)

	defs := make(astmodel.TypeDefinitionSet)
	defs.AddAll(person2020, person2021, person2022)
	defs.AddAll(person2020s, person2021s, person2022s)

	initialState := NewState().WithDefinitions(defs)
	cfg := config.NewConfiguration()
	stage := CreateConversionGraph(cfg)
	finalState, err := stage.Run(context.TODO(), initialState)
	g.Expect(err).To(Succeed())
	g.Expect(finalState.Definitions()).To(Equal(defs))
	g.Expect(finalState.ConversionGraph()).NotTo(BeNil())

	graph := finalState.ConversionGraph()

	// Expect to have a link from Pkg2020 to a matching storage version
	storage2020, ok := graph.LookupTransition(test.Pkg2020)
	g.Expect(ok).To(BeTrue())
	g.Expect(storage2020.String()).To(ContainSubstring(test.Pkg2020.Version()))

	// Expect to have a link from Pkg2021 to a matching storage version
	storage2021, ok := graph.LookupTransition(test.Pkg2021)
	g.Expect(ok).To(BeTrue())
	g.Expect(storage2021.String()).To(ContainSubstring(test.Pkg2021.Version()))

	// Expect to have a link from Pkg2022 to a matching storage version
	storage2022, ok := graph.LookupTransition(test.Pkg2022)
	g.Expect(ok).To(BeTrue())
	g.Expect(storage2022.String()).To(ContainSubstring(test.Pkg2022.Version()))

	// Expect to have a link from Storage2020 to Storage2021
	linkedFrom2020, ok := graph.LookupTransition(storage2020)
	g.Expect(ok).To(BeTrue())
	g.Expect(linkedFrom2020).To(Equal(storage2021))

	// Expect to have a link from Storage2021 version Storage2022
	linkedFrom2021, ok := graph.LookupTransition(storage2021)
	g.Expect(ok).To(BeTrue())
	g.Expect(linkedFrom2021).To(Equal(storage2022))

	// Expect NOT to have a link from Storage2022
	_, ok = graph.LookupTransition(storage2022)
	g.Expect(ok).To(BeFalse())

	// Finally, check that the five links we've verified are the only ones there
	// We check this last, so that a failure doesn't block more detailed checks
	g.Expect(graph.TransitionCount()).To(Equal(5))
}
