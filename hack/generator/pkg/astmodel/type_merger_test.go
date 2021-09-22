/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"testing"

	. "github.com/onsi/gomega"
	"github.com/pkg/errors"
)

func TestCanMergeSameTypes(t *testing.T) {
	g := NewGomegaWithT(t)

	merger := NewTypeMerger(func(_ interface{}, l, r Type) (Type, error) {
		return nil, errors.New("reached fallback")
	})

	merger.Add(func(l, r *PrimitiveType) (Type, error) {
		return StringType, nil
	})

	result, err := merger.Merge(IntType, BoolType)
	g.Expect(err).To(Not(HaveOccurred()))
	g.Expect(result).To(Equal(StringType))
}

func TestCanMergeDifferentTypes(t *testing.T) {
	g := NewGomegaWithT(t)

	merger := NewTypeMerger(func(_ interface{}, l, r Type) (Type, error) {
		return BoolType, nil
	})

	merger.Add(func(l *ObjectType, r *PrimitiveType) (Type, error) {
		return StringType, nil
	})

	result, err := merger.Merge(NewObjectType(), IntType) // shouldn't hit fallback
	g.Expect(err).To(Not(HaveOccurred()))
	g.Expect(result).To(Equal(StringType))

	result, err = merger.Merge(IntType, NewObjectType()) // should hit fallback
	g.Expect(err).To(Not(HaveOccurred()))
	g.Expect(result).To(Equal(BoolType))
}

func TestCanMergeWithGenericTypeArgument(t *testing.T) {
	g := NewGomegaWithT(t)

	merger := NewTypeMerger(func(_ interface{}, l, r Type) (Type, error) {
		return BoolType, nil
	})

	merger.Add(func(l *ObjectType, r Type) (Type, error) {
		return StringType, nil
	})

	result, err := merger.Merge(NewObjectType(), IntType) // shouldn't hit fallback
	g.Expect(err).To(Not(HaveOccurred()))
	g.Expect(result).To(Equal(StringType))

	result, err = merger.Merge(IntType, NewObjectType()) // should hit fallback
	g.Expect(err).To(Not(HaveOccurred()))
	g.Expect(result).To(Equal(BoolType))
}

func TestCanMergeWithUnorderedMerger(t *testing.T) {
	g := NewGomegaWithT(t)

	merger := NewTypeMerger(func(_ interface{}, l, r Type) (Type, error) {
		return BoolType, nil
	})

	merger.AddUnordered(func(l *ObjectType, r Type) (Type, error) {
		return StringType, nil
	})

	result, err := merger.Merge(NewObjectType(), IntType) // shouldn't hit fallback
	g.Expect(err).To(Not(HaveOccurred()))
	g.Expect(result).To(Equal(StringType))

	result, err = merger.Merge(IntType, NewObjectType()) // shouldn't hit fallback either
	g.Expect(err).To(Not(HaveOccurred()))
	g.Expect(result).To(Equal(StringType))
}

var leftFallback MergerFunc = func(ctx interface{}, left, right Type) (Type, error) { return left, nil }

func TestAddPanicsWhenPassedANonFunction(t *testing.T) {
	g := NewGomegaWithT(t)

	merger := NewTypeMerger(leftFallback)

	g.Expect(func() { merger.Add(123) }).To(PanicWith("merger must be a function"))
}

func TestMergerFuncMustTakeTwoOrThreeArguments(t *testing.T) {
	g := NewGomegaWithT(t)

	merger := NewTypeMerger(leftFallback)

	msg := "merger must take take arguments of type (left [some Type], right [some Type]) or (ctx X, left [some Type], right [some Type])"

	g.Expect(func() { merger.Add(func() (Type, error) { return nil, nil }) }).To(PanicWith(msg))
	g.Expect(func() { merger.Add(func(x Type) (Type, error) { return x, nil }) }).To(PanicWith(msg))
	g.Expect(func() { merger.Add(func(x, _, _, _ Type) (Type, error) { return x, nil }) }).To(PanicWith(msg))
}

func TestMergerFuncMustTakeTypesAssignableToTypeAsArguments(t *testing.T) {
	g := NewGomegaWithT(t)

	merger := NewTypeMerger(leftFallback)

	msg := "merger must take take arguments of type (left [some Type], right [some Type]) or (ctx X, left [some Type], right [some Type])"

	// left side wrong
	g.Expect(func() { merger.Add(func(_ int, _ Type) (Type, error) { return nil, nil }) }).To(PanicWith(msg))
	g.Expect(func() { merger.Add(func(_ int, _ EnumType) (Type, error) { return nil, nil }) }).To(PanicWith(msg))
	g.Expect(func() { merger.Add(func(_ interface{}, x int, _ Type) (Type, error) { return nil, nil }) }).To(PanicWith(msg))
	g.Expect(func() { merger.Add(func(_ interface{}, _ int, _ EnumType) (Type, error) { return nil, nil }) }).To(PanicWith(msg))

	// right side wrong
	g.Expect(func() { merger.Add(func(_ Type, _ int) (Type, error) { return nil, nil }) }).To(PanicWith(msg))
	g.Expect(func() { merger.Add(func(_ EnumType, _ int) (Type, error) { return nil, nil }) }).To(PanicWith(msg))
	g.Expect(func() { merger.Add(func(_ interface{}, _ Type, _ int) (Type, error) { return nil, nil }) }).To(PanicWith(msg))
	g.Expect(func() { merger.Add(func(_ interface{}, _ EnumType, _ int) (Type, error) { return nil, nil }) }).To(PanicWith(msg))
}

func TestFuncMustTakeTypesAssignableToTypeAsArguments(t *testing.T) {
	g := NewGomegaWithT(t)

	merger := NewTypeMerger(leftFallback)

	msg := "merger must return (Type, error)"

	g.Expect(func() { merger.Add(func(_, _ Type) Type { return nil }) }).To(PanicWith(msg))
	g.Expect(func() { merger.Add(func(_ interface{}, _, _ Type) Type { return nil }) }).To(PanicWith(msg))
}

func TestMergeReturnsNonNilSide(t *testing.T) {
	g := NewGomegaWithT(t)

	merger := NewTypeMerger(leftFallback)

	var ctx interface{} = nil

	// left side
	result, err := merger.Merge(StringType, nil)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(result).To(Equal(StringType))

	result, err = merger.MergeWithContext(ctx, StringType, nil)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(result).To(Equal(StringType))

	// right side
	result, err = merger.Merge(nil, StringType)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(result).To(Equal(StringType))

	result, err = merger.MergeWithContext(ctx, nil, StringType)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(result).To(Equal(StringType))
}
