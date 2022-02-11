/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"testing"

	. "github.com/onsi/gomega"
)

var (
	oneTypeName = MakeTypeName(GenRuntimeReference, "One")
	twoTypeName = MakeTypeName(GenRuntimeReference, "Two")
)

/*
 * Empty Set tests
 */

func Test_TypeNameSet_WhenEmpty_HasLengthZero(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	emptySet := NewTypeNameSet()
	g.Expect(len(emptySet)).To(Equal(0))
}

/*
 * Add() Tests
 */

func Test_TypeNameSet_AfterAddingFirstItem_ContainsItem(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	set := NewTypeNameSet()
	set.Add(oneTypeName)
	g.Expect(set.Contains(oneTypeName)).To(BeTrue())
}

func Test_TypeNameSet_AfterAddingSecondItem_ContainsItem(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	set := NewTypeNameSet()
	set.Add(oneTypeName)
	set.Add(twoTypeName)
	g.Expect(set.Contains(twoTypeName)).To(BeTrue())
}
