/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/hack/generator/pkg/astmodel"
)

var placeholderPackage = astmodel.MakeLocalPackageReference("prefix", "group", "version")

func TestDuplicateNamesAreCaught(t *testing.T) {
	g := NewGomegaWithT(t)

	prop := astmodel.NewPropertyDefinition("duplicate", "dupe", astmodel.StringType)

	innerObj := astmodel.NewObjectType().WithProperties(prop)
	innerObjProp := astmodel.NewPropertyDefinition("inner", "inner", innerObj).SetFlatten(true)

	objType := astmodel.NewObjectType().WithProperties(prop, innerObjProp)

	types := make(astmodel.Types)
	types.Add(astmodel.MakeTypeDefinition(astmodel.MakeTypeName(placeholderPackage, "objType"), objType))

	result, err := applyPropertyFlattening(context.Background(), types)

	// We don't fail but flattening does not occur, and flatten is set to false
	g.Expect(err).ToNot(HaveOccurred())

	// identical to objType above but the "inner" prop is not set to flatten
	newObjType := astmodel.NewObjectType().WithProperties(prop, innerObjProp.SetFlatten(false))
	expectedTypes := make(astmodel.Types)
	expectedTypes.Add(astmodel.MakeTypeDefinition(astmodel.MakeTypeName(placeholderPackage, "objType"), newObjType))

	g.Expect(result).To(Equal(expectedTypes))
}

func TestFlatteningWorks(t *testing.T) {
	g := NewGomegaWithT(t)

	inner2Obj := astmodel.NewObjectType().WithProperties(
		astmodel.NewPropertyDefinition("x", "x", astmodel.StringType))

	innerObj := astmodel.NewObjectType().WithProperties(
		astmodel.NewPropertyDefinition("inner2", "inner2", inner2Obj).SetFlatten(true),
		astmodel.NewPropertyDefinition("y", "y", astmodel.IntType))

	objType := astmodel.NewObjectType().WithProperties(
		astmodel.NewPropertyDefinition("inner", "inner", innerObj).SetFlatten(true),
		astmodel.NewPropertyDefinition("z", "z", astmodel.IntType))

	types := make(astmodel.Types)
	types.Add(astmodel.MakeTypeDefinition(astmodel.MakeTypeName(placeholderPackage, "objType"), objType))

	result, err := applyPropertyFlattening(context.Background(), types)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(result).To(HaveLen(1))

	var it astmodel.Type
	for _, single := range result {
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

	g.Expect(xProp.FlattenedFrom()).To(Equal([]astmodel.PropertyName{"inner", "inner2"}))
	g.Expect(yProp.FlattenedFrom()).To(Equal([]astmodel.PropertyName{"inner"}))
	g.Expect(zProp.FlattenedFrom()).To(Equal([]astmodel.PropertyName{}))
}
