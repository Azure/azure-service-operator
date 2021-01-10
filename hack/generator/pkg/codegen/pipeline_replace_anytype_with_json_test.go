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

func TestReplacingAnyTypes(t *testing.T) {
	g := NewGomegaWithT(t)
	p1 := astmodel.MakeLocalPackageReference("horo.logy", "v20200730")
	aName := astmodel.MakeTypeName(p1, "A")
	bName := astmodel.MakeTypeName(p1, "B")

	defs := make(astmodel.Types)
	defs.Add(astmodel.MakeTypeDefinition(aName, astmodel.AnyType))
	defs.Add(astmodel.MakeTypeDefinition(
		bName,
		astmodel.NewObjectType().WithProperties(
			astmodel.NewPropertyDefinition("Field1", "field1", astmodel.BoolType),
			astmodel.NewPropertyDefinition("Field2", "field2", astmodel.AnyType),
		),
	))

	results, err := replaceAnyTypeWithJSON().action(context.Background(), defs)

	g.Expect(err).To(BeNil())

	a := results[aName]
	expectedType := astmodel.MakeTypeName(
		astmodel.MakeExternalPackageReference("k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"),
		"JSON",
	)
	g.Expect(a.Type()).To(Equal(expectedType))

	bDef := results[bName]
	bProp, found := bDef.Type().(*astmodel.ObjectType).Property("Field2")
	g.Expect(found).To(BeTrue())
	g.Expect(bProp.PropertyType()).To(Equal(expectedType))
}

func TestReplacingMapMapInterface(t *testing.T) {
	// We want to replace map[string]map[string]interface{} with
	// map[string]JSON, rather than the right one, since
	// controller-gen can't handle it at the moment.
	g := NewGomegaWithT(t)
	p1 := astmodel.MakeLocalPackageReference("horo.logy", "v20200730")
	aName := astmodel.MakeTypeName(p1, "A")

	defs := make(astmodel.Types)
	defs.Add(astmodel.MakeTypeDefinition(
		aName,
		astmodel.NewObjectType().WithProperties(
			astmodel.NewPropertyDefinition("Field1", "field1", astmodel.BoolType),
			astmodel.NewPropertyDefinition("Maps", "maps", astmodel.NewMapType(
				astmodel.StringType,
				astmodel.NewMapType(
					astmodel.StringType,
					astmodel.AnyType,
				),
			)),
		),
	))

	results, err := replaceAnyTypeWithJSON().action(context.Background(), defs)

	g.Expect(err).To(BeNil())

	// A should be a map[string]JSON.
	expectedType := astmodel.NewMapType(
		astmodel.StringType,
		astmodel.MakeTypeName(
			astmodel.MakeExternalPackageReference("k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"),
			"JSON",
		),
	)

	aDef := results[aName]
	aProp, found := aDef.Type().(*astmodel.ObjectType).Property("Maps")
	g.Expect(found).To(BeTrue())
	g.Expect(aProp.PropertyType()).To(Equal(expectedType))
}
