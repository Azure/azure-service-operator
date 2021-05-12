/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"testing"

	. "github.com/onsi/gomega"
)

func Test_ReferenceGraph_Gives_Correct_Depth(t *testing.T) {
	g := NewGomegaWithT(t)

	pr := makeTestLocalPackageReference("group", "package")

	name := func(name string) TypeName {
		return MakeTypeName(pr, name)
	}

	names := func(ns ...string) TypeNameSet {
		result := make(TypeNameSet)
		for _, n := range ns {
			result.Add(name(n))
		}

		return result
	}

	// if there are two paths to reach ‘c’, it should get the lowest depth assigned:

	// r > a (1)
	//  \ /
	//   v
	//   b (1/2)
	//   |
	//   v
	//   c (2/3)

	// the result should not be dependent on iteration order!

	references := map[TypeName]TypeNameSet{
		name("r"): names("a", "b"),
		name("a"): names("b"),
		name("b"): names("c"),
	}

	roots := names("r")

	graph := NewReferenceGraph(roots, references)

	result := graph.Connected()

	g.Expect(result[name("a")]).To(Equal(1))
	g.Expect(result[name("b")]).To(Equal(1))
	g.Expect(result[name("c")]).To(Equal(2)) // not 3!
}
