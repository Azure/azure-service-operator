/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"
	"testing"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/test"

	. "github.com/onsi/gomega"
)

func TestReplacingAnyTypes(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	p1 := test.MakeLocalPackageReference("horo.logy", "v20200730")
	aName := astmodel.MakeInternalTypeName(p1, "A")
	bName := astmodel.MakeInternalTypeName(p1, "B")

	defs := make(astmodel.TypeDefinitionSet)
	defs.Add(astmodel.MakeTypeDefinition(aName, astmodel.AnyType))
	defs.Add(astmodel.MakeTypeDefinition(
		bName,
		astmodel.NewObjectType().WithProperties(
			astmodel.NewPropertyDefinition("Field1", "field1", astmodel.BoolType),
			astmodel.NewPropertyDefinition("Field2", "field2", astmodel.AnyType),
		),
	))

	state := NewState().WithDefinitions(defs)
	finalState, err := ReplaceAnyTypeWithJSON().action(context.Background(), state)
	g.Expect(err).To(BeNil())

	finalDefs := finalState.Definitions()
	a := finalDefs[aName]
	expectedType := astmodel.MakeInternalTypeName(
		astmodel.MakeExternalPackageReference("k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"),
		"JSON",
	)
	g.Expect(a.Type()).To(Equal(expectedType))

	bDef := finalDefs[bName]
	bProp, found := bDef.Type().(*astmodel.ObjectType).Property("Field2")
	g.Expect(found).To(BeTrue())
	g.Expect(bProp.PropertyType()).To(Equal(expectedType))
}

func TestReplacingMapMapInterface(t *testing.T) {
	// We want to replace map[string]map[string]interface{} with
	// map[string]JSON, rather than the right one, since
	// controller-gen can't handle it at the moment.
	t.Parallel()
	g := NewGomegaWithT(t)
	p1 := test.MakeLocalPackageReference("horo.logy", "v20200730")
	aName := astmodel.MakeInternalTypeName(p1, "A")

	defs := make(astmodel.TypeDefinitionSet)
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

	state := NewState().WithDefinitions(defs)
	finalState, err := ReplaceAnyTypeWithJSON().action(context.Background(), state)

	g.Expect(err).To(BeNil())

	// A should be a map[string]JSON.
	expectedType := astmodel.NewMapType(
		astmodel.StringType,
		astmodel.MakeInternalTypeName(
			astmodel.MakeExternalPackageReference("k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"),
			"JSON",
		),
	)

	finalDefinitions := finalState.Definitions()
	aDef := finalDefinitions[aName]
	aProp, found := aDef.Type().(*astmodel.ObjectType).Property("Maps")
	g.Expect(found).To(BeTrue())
	g.Expect(aProp.PropertyType()).To(Equal(expectedType))
}
