/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"testing"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"

	. "github.com/onsi/gomega"
)

const packagePath = "test.package/v1"

func TestConnectionChecker_Avoids_Cycles(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	makeName := func(name string) astmodel.TypeName {
		return astmodel.MakeInternalTypeName(astmodel.MakeExternalPackageReference(packagePath), name)
	}

	makeSet := func(names ...string) astmodel.TypeNameSet {
		var typeNames []astmodel.TypeName
		for _, n := range names {
			typeNames = append(typeNames, makeName(n))
		}
		return astmodel.NewTypeNameSet(typeNames...)
	}

	roots := makeSet("res1", "res2")
	references := map[astmodel.TypeName]astmodel.TypeNameSet{
		makeName("G1"):   makeSet("G2"),
		makeName("G2"):   makeSet("A"),
		makeName("res1"): makeSet("A"),
		makeName("A"):    makeSet("B", "C"),
		makeName("B"):    nil,
		makeName("C"):    makeSet("D"),
		makeName("D"):    makeSet("A"), // cyclic
	}

	graph := astmodel.MakeReferenceGraph(roots, references)
	connectedSet := graph.Connected()

	names := astmodel.NewTypeNameSet()
	for name := range connectedSet {
		names.Add(name)
	}

	g.Expect(names).To(Equal(makeSet("res1", "res2", "A", "B", "C", "D")))
}
