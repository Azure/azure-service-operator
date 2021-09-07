/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"
	"testing"

	"github.com/Azure/azure-service-operator/hack/generator/pkg/astmodel"
	"github.com/Azure/azure-service-operator/hack/generator/pkg/test"

	"github.com/pkg/errors"

	. "github.com/onsi/gomega"
)

func TestFindsAnyTypes(t *testing.T) {
	g := NewWithT(t)
	p1 := test.MakeLocalPackageReference("horo.logy", "v20200730")
	p2 := test.MakeLocalPackageReference("road.train", "v20200730")
	p3 := test.MakeLocalPackageReference("wah.wah", "v20200730")

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

	state := NewState().WithTypes(defs)
	stage := FilterOutDefinitionsUsingAnyType(nil)
	finalState, err := stage.Run(context.Background(), state)

	g.Expect(finalState).To(BeNil())
	g.Expect(err).To(MatchError("AnyTypes found - add exclusions for: horo.logy/v20200730, road.train/v20200730"))
}

func TestIgnoresExpectedAnyTypePackages(t *testing.T) {
	g := NewWithT(t)
	p1 := test.MakeLocalPackageReference("horo.logy", "v20200730")
	p2 := test.MakeLocalPackageReference("road.train", "v20200730")
	p3 := test.MakeLocalPackageReference("wah.wah", "v20200730")

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

	state := NewState().WithTypes(defs)
	finalState, err := FilterOutDefinitionsUsingAnyType(exclusions).action(context.Background(), state)
	g.Expect(err).To(BeNil())

	expected := make(astmodel.Types)
	expected.Add(astmodel.MakeTypeDefinition(
		astmodel.MakeTypeName(p3, "C"), astmodel.NewArrayType(astmodel.IntType),
	))
	g.Expect(finalState.Types()).To(Equal(expected))
}

func TestComplainsAboutUnneededExclusions(t *testing.T) {
	g := NewWithT(t)
	p1 := test.MakeLocalPackageReference("horo.logy", "v20200730")
	p2 := test.MakeLocalPackageReference("road.train", "v20200730")
	p3 := test.MakeLocalPackageReference("wah.wah", "v20200730")

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

	state := NewState().WithTypes(defs)
	stage := FilterOutDefinitionsUsingAnyType(exclusions)
	finalState, err := stage.Run(context.Background(), state)
	g.Expect(finalState).To(BeNil())
	g.Expect(errors.Cause(err)).To(MatchError("no AnyTypes found in: gamma.knife/v20200821, people.vultures/20200821"))
}
