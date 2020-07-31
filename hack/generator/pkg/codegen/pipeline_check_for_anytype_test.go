/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package codegen

import (
	"context"
	"testing"

	"github.com/Azure/k8s-infra/hack/generator/pkg/astmodel"

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

	results, err := checkForAnyType().Action(context.Background(), defs)

	g.Expect(results).To(HaveLen(0))
	g.Expect(err).To(MatchError("AnyTypes found - add exclusions for: horo.logy/v20200730, road.train/v20200730"))
}
