/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"testing"

	. "github.com/onsi/gomega"
)

/*
 * NewFileDefinition tests
 */

func Test_NewFileDefinition_GivenValues_InitializesFields(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	person := NewTestObject(
		"Person",
		NewStringPropertyDefinition("fullName"),
		NewStringPropertyDefinition("knownAs"),
		NewStringPropertyDefinition("familyName"),
	)
	file := NewFileDefinition(person.Name().PackageReference(), []TypeDefinition{person}, nil)

	g.Expect(file.packageReference).To(Equal(person.Name().PackageReference()))
	g.Expect(file.definitions).To(HaveLen(1))
}

/*
 * calcRanks() tests
 */

func Test_CalcRanks_GivenMultipleRoots_AssignsRankZeroToAll(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// +---------+   +---------+   +---------+   +---------+
	// |         |   |         |   |         |   |         |
	// |  root1  |   |  root2  |   |  root3  |   |  root4  |
	// |         |   |         |   |         |   |         |
	// +---------+   +---------+   +---------+   +---------+

	root1 := NewTestObject("r1")
	root2 := NewTestObject("b")
	root3 := NewTestObject("c")
	root4 := NewTestObject("d")

	ranks := calcRanks([]TypeDefinition{root1, root2, root3, root4})

	g.Expect(ranks[root1.Name()]).To(Equal(0))
	g.Expect(ranks[root2.Name()]).To(Equal(0))
	g.Expect(ranks[root3.Name()]).To(Equal(0))
	g.Expect(ranks[root4.Name()]).To(Equal(0))
}

func Test_CalcRanks_GivenLinearDependencies_AssignsRanksInOrder(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// +---------+
	// |         |
	// |  rank0  |
	// |         |
	// +---+-----+
	//     |
	//     |
	//     v
	// +---+-----+
	// |         |
	// |  rank1  |
	// |         |
	// +---+-----+
	//     |
	//     |
	//     v
	// +---+-----+
	// |         |
	// |  rank2  |
	// |         |
	// +---+-----+
	//     |
	//     |
	//     v
	// +---+-----+
	// |         |
	// |  rank3  |
	// |         |
	// +---------+

	rank3 := NewTestObject("d")
	referenceToRank3 := NewPropertyDefinition("f3", "f3", rank3.Name())

	rank2 := NewTestObject("c", referenceToRank3)
	referenceToRank2 := NewPropertyDefinition("f2", "f2", rank2.Name())

	rank1 := NewTestObject("b", referenceToRank2)
	referenceToRank1 := NewPropertyDefinition("f1", "f1", rank1.Name())

	rank0 := NewTestObject("a", referenceToRank1)

	ranks := calcRanks([]TypeDefinition{rank0, rank1, rank2, rank3})

	g.Expect(ranks[rank0.Name()]).To(Equal(0))
	g.Expect(ranks[rank1.Name()]).To(Equal(1))
	g.Expect(ranks[rank2.Name()]).To(Equal(2))
	g.Expect(ranks[rank3.Name()]).To(Equal(3))
}

func Test_CalcRanks_GivenDiamondDependencies_AssignRanksInOrder(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	//         +---------+
	//         |         |
	//         |   top   |
	//         |         |
	//         ++-------++
	//          |       |
	//          |       |
	//          v       v
	// +--------++     ++--------+
	// |         |     |         |
	// |  left   |     |  right  |
	// |         |     |         |
	// +--------++     ++--------+
	//          |       |
	//          |       |
	//          v       v
	//         ++-------++
	//         |         |
	//         | bottom  |
	//         |         |
	//         +---------+

	bottom := NewTestObject("bottom")
	referenceToBottom := NewPropertyDefinition("b", "b", bottom.Name())

	left := NewTestObject("l", referenceToBottom)
	referenceToLeft := NewPropertyDefinition("l", "l", left.Name())

	right := NewTestObject("r", referenceToBottom)
	referenceToRight := NewPropertyDefinition("r", "r", right.Name())

	top := NewTestObject("a", referenceToLeft, referenceToRight)

	ranks := calcRanks([]TypeDefinition{top, left, right, bottom})

	g.Expect(ranks[top.Name()]).To(Equal(0))
	g.Expect(ranks[right.Name()]).To(Equal(1))
	g.Expect(ranks[left.Name()]).To(Equal(1))
	g.Expect(ranks[bottom.Name()]).To(Equal(2))
}

func Test_CalcRanks_GivenDiamondWithBar_AssignRanksInOrder(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	//         +---------+
	//         |         |
	//         |   top   |
	//         |         |
	//         ++-------++
	//          |       |
	//          |       |
	//          v       v
	// +--------++     ++--------+
	// |         |     |         |
	// |  left   +---->+  right  |
	// |         |     |         |
	// +--------++     ++--------+
	//          |       |
	//          |       |
	//          v       v
	//         ++-------++
	//         |         |
	//         | bottom  |
	//         |         |
	//         +---------+

	bottom := NewTestObject("bottom")
	referenceToBottom := NewPropertyDefinition("b", "b", bottom.Name())

	right := NewTestObject("r", referenceToBottom)
	referenceToRight := NewPropertyDefinition("r", "r", right.Name())

	left := NewTestObject("l", referenceToBottom, referenceToRight)
	referenceToLeft := NewPropertyDefinition("l", "l", left.Name())

	top := NewTestObject("a", referenceToLeft, referenceToRight)

	ranks := calcRanks([]TypeDefinition{top, left, right, bottom})

	g.Expect(ranks[top.Name()]).To(Equal(0))
	g.Expect(ranks[right.Name()]).To(Equal(1))
	g.Expect(ranks[left.Name()]).To(Equal(1))
	g.Expect(ranks[bottom.Name()]).To(Equal(2))
}

func Test_CalcRanks_GivenDiamondWithReverseBar_AssignRanksInOrder(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	//         +---------+
	//         |         |
	//         |   top   |
	//         |         |
	//         ++-------++
	//          |       |
	//          |       |
	//          v       v
	// +--------++     ++--------+
	// |         |     |         |
	// |  left   +<----+  right  |
	// |         |     |         |
	// +--------++     ++--------+
	//          |       |
	//          |       |
	//          v       v
	//         ++-------++
	//         |         |
	//         | bottom  |
	//         |         |
	//         +---------+

	bottom := NewTestObject("bottom")

	referenceToBottom := NewPropertyDefinition("b", "b", bottom.Name())
	left := NewTestObject("l", referenceToBottom)

	referenceToLeft := NewPropertyDefinition("l", "l", left.Name())
	right := NewTestObject("r", referenceToBottom, referenceToLeft)

	referenceToRight := NewPropertyDefinition("r", "r", right.Name())
	top := NewTestObject("a", referenceToLeft, referenceToRight)

	ranks := calcRanks([]TypeDefinition{top, left, right, bottom})

	g.Expect(ranks[top.Name()]).To(Equal(0))
	g.Expect(ranks[right.Name()]).To(Equal(1))
	g.Expect(ranks[left.Name()]).To(Equal(1))
	g.Expect(ranks[bottom.Name()]).To(Equal(2))
}

/*
 * Supporting methods
 */

func NewTestObject(name string, fields ...*PropertyDefinition) TypeDefinition {
	ref := MakeInternalTypeName(makeTestLocalPackageReference("group", "2020-01-01"), name)
	return MakeTypeDefinition(ref, NewObjectType().WithProperties(fields...))
}

func NewStringPropertyDefinition(name string) *PropertyDefinition {
	return NewPropertyDefinition(PropertyName(name), name, StringType)
}
