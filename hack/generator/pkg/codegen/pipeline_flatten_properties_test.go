/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package codegen

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
	g.Expect(result).To(BeNil())
	g.Expect(err).To(MatchError("visit of type of \"prefix/group/version/objType\" failed: Flattening caused duplicate property name \"duplicate\""))
}

func TestFlatteningWorks(t *testing.T) {
	g := NewGomegaWithT(t)

	innerObj := astmodel.NewObjectType().WithProperties(
		astmodel.NewPropertyDefinition("x", "x", astmodel.StringType),
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
}
