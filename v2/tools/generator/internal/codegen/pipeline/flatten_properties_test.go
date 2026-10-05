/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/go-logr/logr"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/test"
)

var placeholderPackage = test.MakeLocalPackageReference("group", "version")

func TestDuplicateNamesAreCaughtAndRenamed(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	prop := astmodel.NewPropertyDefinition("Duplicate", "dupe", astmodel.StringType)

	innerObj := astmodel.NewObjectType().WithProperties(prop)
	innerObjProp := astmodel.NewPropertyDefinition("Inner", "inner", innerObj).SetFlatten(true)

	objType := astmodel.NewObjectType().WithProperties(prop, innerObjProp)

	defs := make(astmodel.TypeDefinitionSet)
	defs.Add(astmodel.MakeTypeDefinition(astmodel.MakeInternalTypeName(placeholderPackage, "ObjType"), objType))

	state := NewState(defs)
	stage := FlattenProperties(logr.Discard())

	result, err := stage.Run(context.Background(), state)

	// We don't fail but flattening does not occur, and flatten is set to false
	g.Expect(err).ToNot(HaveOccurred())

	// should have a renamed property which is flattened-from "inner"
	newName := astmodel.PropertyName("InnerDuplicate")
	newJSONName := "inner_duplicate"
	newObjType := astmodel.NewObjectType().
		WithProperties(
			prop,
			prop.WithName(newName).WithJSONName(newJSONName).AddFlattenedFrom("Inner"),
		)
	expectedDefs := make(astmodel.TypeDefinitionSet)
	expectedDefs.Add(astmodel.MakeTypeDefinition(astmodel.MakeInternalTypeName(placeholderPackage, "ObjType"), newObjType))

	g.Expect(result.Definitions()).To(Equal(expectedDefs))
}

func TestFlatteningWorks(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	inner2Obj := astmodel.NewObjectType().WithProperties(
		astmodel.NewPropertyDefinition("x", "x", astmodel.StringType),
	)

	innerObj := astmodel.NewObjectType().WithProperties(
		astmodel.NewPropertyDefinition("inner2", "inner2", inner2Obj).SetFlatten(true),
		astmodel.NewPropertyDefinition("y", "y", astmodel.IntType),
	)

	objType := astmodel.NewObjectType().WithProperties(
		astmodel.NewPropertyDefinition("inner", "inner", innerObj).SetFlatten(true),
		astmodel.NewPropertyDefinition("z", "z", astmodel.IntType),
	)

	defs := make(astmodel.TypeDefinitionSet)
	defs.Add(astmodel.MakeTypeDefinition(astmodel.MakeInternalTypeName(placeholderPackage, "objType"), objType))

	state := NewState(defs)
	stage := FlattenProperties(logr.Discard())

	result, err := stage.Run(context.Background(), state)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(result.Definitions()).To(HaveLen(1))

	var it astmodel.Type
	for _, single := range result.Definitions() {
		it = single.Type()
		break
	}

	ot := it.(*astmodel.ObjectType)
	g.Expect(ot.Properties()).To(HaveLen(3))

	xProp, ok := ot.Property("x")
	g.Expect(ok).To(BeTrue())
	yProp, ok := ot.Property("y")
	g.Expect(ok).To(BeTrue())
	zProp, ok := ot.Property("z")
	g.Expect(ok).To(BeTrue())

	g.Expect(xProp.FlattenedFrom()).To(Equal([]astmodel.PropertyName{"inner", "inner2", "x"}))
	g.Expect(yProp.FlattenedFrom()).To(Equal([]astmodel.PropertyName{"inner", "y"}))
	g.Expect(zProp.FlattenedFrom()).To(Equal([]astmodel.PropertyName{"z"}))
}

// TestFlattenProperties_FlattensSelectedProperties verifies that properties explicitly marked for flattening
// are inlined while other properties retain their original shape.
func TestFlattenProperties_FlattensSelectedProperties(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	inner := test.CreateObjectDefinition(
		test.Pkg2020,
		"Inner",
		astmodel.NewPropertyDefinition("Color", "color", astmodel.OptionalStringType),
		astmodel.NewPropertyDefinition("Name", "name", astmodel.OptionalStringType),
	)

	outerName := astmodel.MakeInternalTypeName(test.Pkg2020, "Test")
	outer := astmodel.MakeTypeDefinition(
		outerName,
		astmodel.NewObjectType().WithProperties(
			astmodel.NewPropertyDefinition("InnerFlat", "innerFlat", astmodel.NewOptionalType(inner.Name())).
				SetFlatten(true),
			astmodel.NewPropertyDefinition("InnerNotFlat", "innerNotFlat", astmodel.NewOptionalType(inner.Name())),
		),
	)

	defs := make(astmodel.TypeDefinitionSet)
	defs.AddAll(inner, outer)

	state, err := RunTestPipeline(
		NewState(defs),
		FlattenProperties(logr.Discard()),
	)
	g.Expect(err).ToNot(HaveOccurred())

	result := state.Definitions().MustGetDefinition(outerName)
	test.AssertPropertyExistsWithType(t, result.Type(), "Color", astmodel.OptionalStringType)
	test.AssertPropertyExistsWithType(t, result.Type(), "Name", astmodel.OptionalStringType)
	test.AssertPropertyExistsWithType(t, result.Type(), "InnerNotFlat", astmodel.NewOptionalType(inner.Name()))
	test.AssertPropertyCount(t, result.Type(), 3)

	test.AssertPackagesGenerateExpectedCode(t, state.Definitions())
}
