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
	t.Parallel()
	g := NewGomegaWithT(t)

	badObject := NewTestObject("BadName")

	prop := NewPropertyDefinition("Ref", "ref", badObject.Name())
	otherObject := NewTestObject("Container", prop)

	defs := make(TypeDefinitionSet)
	defs.AddAll(badObject, otherObject)

	newName := badObject.Name().WithName("GoodName")
	renames := map[InternalTypeName]InternalTypeName{
		badObject.Name(): newName,
	}
	renamer := NewRenamingVisitor(renames)
	result, err := renamer.RenameAll(defs)

	expectedRenamedTypeName := badObject.Name().WithName("GoodName")
	expectedOtherObject := otherObject.Type().(*ObjectType).WithProperty(prop.WithType(expectedRenamedTypeName))
	expectedResult := make(TypeDefinitionSet)

	expectedResult.Add(badObject.WithName(expectedRenamedTypeName))
	expectedResult.Add(otherObject.WithType(expectedOtherObject))

	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(result).To(Equal(expectedResult))
}

func TestRenamingVisitor_RewritesResourceOwner(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	badObject := NewTestObject("BadName")
	badName := badObject.Name()

	childSpecDef := NewTestObject("ChildSpec")
	childStatusDef := NewTestObject("ChildStatus")
	childResource := NewResourceType(childSpecDef.Name(), childStatusDef.Name()).
		WithOwner(badName)
	childDef := MakeTypeDefinition(
		MakeInternalTypeName(badObject.name.InternalPackageReference(), "ChildResource"),
		childResource,
	)

	defs := make(TypeDefinitionSet)
	defs.AddAll(badObject, childDef)

	newName := badObject.Name().WithName("GoodName")
	renames := map[InternalTypeName]InternalTypeName{
		badObject.Name(): newName,
	}
	renamer := NewRenamingVisitor(renames)
	result, err := renamer.RenameAll(defs)

	expectedRenamedTypeName := badObject.Name().WithName("GoodName")

	expectedChildDef := MakeTypeDefinition(
		childDef.Name(),
		childResource.WithOwner(expectedRenamedTypeName),
	)
	expectedResult := make(TypeDefinitionSet)

	expectedResult.Add(badObject.WithName(expectedRenamedTypeName))
	expectedResult.Add(expectedChildDef)

	g.Expect(err).ToNot(HaveOccurred())

	// Expect the parent resource to be renamed
	g.Expect(result[newName]).To(Equal(badObject.WithName(newName)))

	// Expect the child resource to have a new owner
	g.Expect(result[expectedChildDef.Name()]).To(Equal(expectedChildDef))

	g.Expect(result).To(Equal(expectedResult))
}
