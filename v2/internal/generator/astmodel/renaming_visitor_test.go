/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestRenamingVisitor_RenamesTypeAndReferences(t *testing.T) {
	g := NewGomegaWithT(t)

	badObject := NewTestObject("BadName")

	prop := NewPropertyDefinition("Ref", "ref", badObject.Name())
	otherObject := NewTestObject("Container", prop)

	types := make(Types)
	types.AddAll(badObject, otherObject)

	newName := badObject.Name().WithName("GoodName")
	renames := map[TypeName]TypeName{
		badObject.Name(): newName,
	}
	renamer := NewRenamingVisitor(renames)
	result, err := renamer.RenameAll(types)

	expectedRenamedTypeName := badObject.Name().WithName("GoodName")
	expectedOtherObject := otherObject.Type().(*ObjectType).WithProperty(prop.WithType(expectedRenamedTypeName))
	expectedResult := make(Types)

	expectedResult.Add(badObject.WithName(expectedRenamedTypeName))
	expectedResult.Add(otherObject.WithType(expectedOtherObject))

	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(result).To(Equal(expectedResult))
}

func TestRenamingVisitor_RewritesResourceOwner(t *testing.T) {
	g := NewGomegaWithT(t)

	badObject := NewTestObject("BadName")
	badName := badObject.Name()

	childSpecDef := NewTestObject("ChildSpec")
	childStatusDef := NewTestObject("ChildStatus")
	childResource := NewResourceType(childSpecDef.Name(), childStatusDef.Name()).
		WithOwner(&badName)
	childDef := MakeTypeDefinition(
		MakeTypeName(badObject.name.PackageReference, "ChildResource"),
		childResource,
	)

	types := make(Types)
	types.AddAll(badObject, childDef)

	newName := badObject.Name().WithName("GoodName")
	renames := map[TypeName]TypeName{
		badObject.Name(): newName,
	}
	renamer := NewRenamingVisitor(renames)
	result, err := renamer.RenameAll(types)

	expectedRenamedTypeName := badObject.Name().WithName("GoodName")

	expectedChildDef := MakeTypeDefinition(
		childDef.Name(),
		NewResourceType(childSpecDef.Name(), childStatusDef.Name()).
			WithOwner(&expectedRenamedTypeName),
	)
	expectedResult := make(Types)

	expectedResult.Add(badObject.WithName(expectedRenamedTypeName))
	expectedResult.Add(expectedChildDef)

	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(result).To(Equal(expectedResult))
}
