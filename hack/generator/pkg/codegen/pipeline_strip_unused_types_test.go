/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package codegen

import (
	"testing"

	"github.com/Azure/k8s-infra/hack/generator/pkg/astmodel"

	. "github.com/onsi/gomega"
)

const packagePath = "test.package/v1"

func TestConnectionChecker_Avoids_Cycles(t *testing.T) {
	g := NewGomegaWithT(t)
	makeName := func(name string) astmodel.TypeName {
		return astmodel.MakeTypeName(astmodel.MakeLibraryPackageReference(packagePath), name)
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

	graph := astmodel.NewReferenceGraph(roots, references)
	connectedSet := graph.Connected()

	var names astmodel.TypeNameSet
	for name := range connectedSet {
		names = names.Add(name)
	}

	g.Expect(names).To(Equal(makeSet("res1", "res2", "A", "B", "C", "D")))
}
