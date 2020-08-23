/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package codegen

import (
	"context"
	"testing"

	"github.com/Azure/k8s-infra/hack/generator/pkg/astmodel"
	"github.com/pkg/errors"

	. "github.com/onsi/gomega"
)

func TestFindsAnytypes(t *testing.T) {
	g := NewGomegaWithT(t)
	p1 := astmodel.MakeLocalPackageReference("horo.logy", "v20200730")
	p2 := astmodel.MakeLocalPackageReference("road.train", "v20200730")
	p3 := astmodel.MakeLocalPackageReference("wah.wah", "v20200730")

	defs := make(astmodel.Types)
	add := func(p astmodel.PackageReference, n string, t astmodel.Type) {
		defs.Add(astmodel.MakeTypeDefinition(astmodel.MakeTypeName(p, n), t))
	}
	// A couple of types in the same package...
	add(p1, "A", astmodel.AnyType)
	add(p1, "B", astmodel.NewObjectType().WithProperties(
		astmodel.NewPropertyDefinition("Field1", "field1", astmodel.BoolType),
		astmodel.NewPropertyDefinition("Field2", "field2", astmodel.AnyType),
	))
	// One in another...
	add(p2, "A", astmodel.NewMapType(astmodel.StringType, astmodel.AnyType))
	// One that's fine.
	add(p3, "C", astmodel.NewArrayType(astmodel.IntType))

	results, err := checkForAnyType(nil).Action(context.Background(), defs)

	g.Expect(results).To(HaveLen(0))
	g.Expect(err).To(MatchError("AnyTypes found - add exclusions for: horo.logy/v20200730, road.train/v20200730"))
}

func TestIgnoresExpectedAnyTypePackages(t *testing.T) {
	g := NewGomegaWithT(t)
	p1 := astmodel.MakeLocalPackageReference("horo.logy", "v20200730")
	p2 := astmodel.MakeLocalPackageReference("road.train", "v20200730")
	p3 := astmodel.MakeLocalPackageReference("wah.wah", "v20200730")

	defs := make(astmodel.Types)
	add := func(p astmodel.PackageReference, n string, t astmodel.Type) {
		defs.Add(astmodel.MakeTypeDefinition(astmodel.MakeTypeName(p, n), t))
	}
	// A couple of types in the same package...
	add(p1, "A", astmodel.AnyType)
	add(p1, "B", astmodel.NewObjectType().WithProperties(
		astmodel.NewPropertyDefinition("Field1", "field1", astmodel.BoolType),
		astmodel.NewPropertyDefinition("Field2", "field2", astmodel.AnyType),
	))
	// One in another...
	add(p2, "A", astmodel.NewMapType(astmodel.StringType, astmodel.AnyType))
	// One that's fine.
	add(p3, "C", astmodel.NewArrayType(astmodel.IntType))

	exclusions := []string{"horo.logy/v20200730", "road.train/v20200730"}
	results, err := checkForAnyType(exclusions).Action(context.Background(), defs)
	g.Expect(err).To(BeNil())

	expected := make(astmodel.Types)
	expected.Add(astmodel.MakeTypeDefinition(
		astmodel.MakeTypeName(p3, "C"), astmodel.NewArrayType(astmodel.IntType),
	))
	g.Expect(results).To(Equal(expected))
}

func TestComplainsAboutUnneededExclusions(t *testing.T) {
	g := NewGomegaWithT(t)
	p1 := astmodel.MakeLocalPackageReference("horo.logy", "v20200730")
	p2 := astmodel.MakeLocalPackageReference("road.train", "v20200730")
	p3 := astmodel.MakeLocalPackageReference("wah.wah", "v20200730")

	defs := make(astmodel.Types)
	add := func(p astmodel.PackageReference, n string, t astmodel.Type) {
		defs.Add(astmodel.MakeTypeDefinition(astmodel.MakeTypeName(p, n), t))
	}
	// A couple of types in the same package...
	add(p1, "A", astmodel.AnyType)
	add(p1, "B", astmodel.NewObjectType().WithProperties(
		astmodel.NewPropertyDefinition("Field1", "field1", astmodel.BoolType),
		astmodel.NewPropertyDefinition("Field2", "field2", astmodel.AnyType),
	))
	// One in another...
	add(p2, "A", astmodel.NewMapType(astmodel.StringType, astmodel.AnyType))
	// One that's fine.
	add(p3, "C", astmodel.NewArrayType(astmodel.IntType))

	exclusions := []string{
		"people.vultures/20200821",
		"horo.logy/v20200730",
		"gamma.knife/v20200821",
		"road.train/v20200730",
	}
	results, err := checkForAnyType(exclusions).Action(context.Background(), defs)
	g.Expect(results).To(HaveLen(0))
	g.Expect(errors.Cause(err)).To(MatchError("no AnyTypes found in: gamma.knife/v20200821, people.vultures/20200821"))
}
