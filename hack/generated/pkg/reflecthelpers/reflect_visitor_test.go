/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package reflecthelpers

import (
	"reflect"
	"testing"

	. "github.com/onsi/gomega"
)

type PrimitiveStruct struct {
	Int1    int
	Int2    int32
	String1 string
}

type PrimitivePtrStruct struct {
	Int1    *int
	Int2    *int32
	String1 string
}

func Test_ReflectVisitor_VisitsSimpleStruct(t *testing.T) {
	g := NewGomegaWithT(t)

	var visitedTypes []reflect.Kind
	visitor := NewReflectVisitor()
	visitor.VisitPrimitive = func(this *ReflectVisitor, it interface{}, ctx interface{}) error {
		visitedTypes = append(visitedTypes, reflect.TypeOf(it).Kind())

		return nil
	}

	s := PrimitiveStruct{
		Int1:    5,
		Int2:    7,
		String1: "Hello reflection",
	}

	err := visitor.Visit(s, nil)
	g.Expect(err).NotTo(HaveOccurred())

	g.Expect(visitedTypes).To(HaveLen(3))
	g.Expect(visitedTypes[0]).To(Equal(reflect.Int))
	g.Expect(visitedTypes[1]).To(Equal(reflect.Int32))
	g.Expect(visitedTypes[2]).To(Equal(reflect.String))
}

func Test_ReflectVisitor_VisitsPtrToSimpleStruct(t *testing.T) {
	g := NewGomegaWithT(t)

	var visitedTypes []reflect.Kind
	visitor := NewReflectVisitor()
	visitor.VisitPrimitive = func(this *ReflectVisitor, it interface{}, ctx interface{}) error {
		visitedTypes = append(visitedTypes, reflect.TypeOf(it).Kind())

		return nil
	}

	s := &PrimitiveStruct{
		Int1:    5,
		Int2:    7,
		String1: "Hello reflection",
	}

	err := visitor.Visit(s, nil)
	g.Expect(err).NotTo(HaveOccurred())

	g.Expect(visitedTypes).To(HaveLen(3))
	g.Expect(visitedTypes[0]).To(Equal(reflect.Int))
	g.Expect(visitedTypes[1]).To(Equal(reflect.Int32))
	g.Expect(visitedTypes[2]).To(Equal(reflect.String))
}

func Test_ReflectVisitor_VisitsStructWithNilField(t *testing.T) {
	g := NewGomegaWithT(t)

	var visitedTypes []reflect.Kind
	visitor := NewReflectVisitor()
	visitor.VisitPrimitive = func(this *ReflectVisitor, it interface{}, ctx interface{}) error {
		visitedTypes = append(visitedTypes, reflect.TypeOf(it).Kind())

		return nil
	}

	num := int32(7)
	s := PrimitivePtrStruct{
		Int1:    nil,
		Int2:    &num,
		String1: "Hello reflection",
	}

	err := visitor.Visit(s, nil)
	g.Expect(err).NotTo(HaveOccurred())

	g.Expect(visitedTypes).To(HaveLen(2))
	g.Expect(visitedTypes[0]).To(Equal(reflect.Int32))
	g.Expect(visitedTypes[1]).To(Equal(reflect.String))
}

func Test_ReflectVisitor_VisitsSlice(t *testing.T) {
	g := NewGomegaWithT(t)

	var visitedTypes []reflect.Kind
	visitor := NewReflectVisitor()
	visitor.VisitPrimitive = func(this *ReflectVisitor, it interface{}, ctx interface{}) error {
		visitedTypes = append(visitedTypes, reflect.TypeOf(it).Kind())

		return nil
	}

	num := int32(7)
	slice := []*PrimitivePtrStruct{
		{
			Int1:    nil,
			Int2:    &num,
			String1: "Hello reflection",
		},
		{
			Int1:    nil,
			Int2:    &num,
			String1: "Hello reflection",
		},
		{
			Int1:    nil,
			Int2:    &num,
			String1: "Hello reflection",
		},
		nil,
	}

	err := visitor.Visit(slice, nil)
	g.Expect(err).NotTo(HaveOccurred())

	g.Expect(visitedTypes).To(HaveLen(6))
	for i := 0; i < 6; i += 2 {
		g.Expect(visitedTypes[i]).To(Equal(reflect.Int32))
		g.Expect(visitedTypes[i+1]).To(Equal(reflect.String))
	}
}

func Test_ReflectVisitor_VisitsMap(t *testing.T) {
	g := NewGomegaWithT(t)

	var visitedTypes []reflect.Kind
	visitor := NewReflectVisitor()
	visitor.VisitPrimitive = func(this *ReflectVisitor, it interface{}, ctx interface{}) error {
		visitedTypes = append(visitedTypes, reflect.TypeOf(it).Kind())

		return nil
	}

	num := int32(7)
	m := map[string]*PrimitivePtrStruct{
		"test": {
			Int1:    nil,
			Int2:    &num,
			String1: "Hello reflection",
		},
		"foo": {
			Int1:    nil,
			Int2:    &num,
			String1: "Hello reflection",
		},
		"bar": {
			Int1:    nil,
			Int2:    &num,
			String1: "Hello reflection",
		},
	}

	err := visitor.Visit(m, nil)
	g.Expect(err).NotTo(HaveOccurred())

	g.Expect(visitedTypes).To(HaveLen(9))
	for i := 0; i < 6; i += 3 {
		g.Expect(visitedTypes[i]).To(Equal(reflect.String))
		g.Expect(visitedTypes[i+1]).To(Equal(reflect.Int32))
		g.Expect(visitedTypes[i+2]).To(Equal(reflect.String))
	}
}
