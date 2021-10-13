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
	types.Add(badObject)
	types.Add(otherObject)

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
