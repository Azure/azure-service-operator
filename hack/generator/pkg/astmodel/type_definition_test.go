/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"testing"

	. "github.com/onsi/gomega"
)

/*
 * MakeTypeDefinition() Tests
 */

func Test_MakeTypeDefinition_GivenValues_InitializesProperties(t *testing.T) {
	g := NewGomegaWithT(t)

	const name = "demo"
	const group = "group"
	const version = "2020-01-01"

	ref := MakeTypeName(MakeLocalPackageReference(group, version), name)
	objectType := NewObjectType().WithProperties(fullName, familyName, knownAs)
	objectDefinition := MakeTypeDefinition(ref, objectType)

	g.Expect(objectDefinition.Name().name).To(Equal(name))
	g.Expect(objectDefinition.Type()).To(Equal(objectType))

	localRef, ok := objectDefinition.Name().PackageReference.AsLocalPackage()
	g.Expect(ok).To(BeTrue())

	g.Expect(localRef.Group()).To(Equal(group))
	g.Expect(localRef.Version()).To(Equal(version))
	g.Expect(objectDefinition.Type().(*ObjectType).properties).To(HaveLen(3))
}

/*
 * WithDescription() tests
 */

func Test_TypeDefinitionWithDescription_GivenDescription_ReturnsExpected(t *testing.T) {
	g := NewGomegaWithT(t)

	const name = "demo"
	const group = "group"
	const version = "2020-01-01"

	description := []string{"This is my test description"}

	ref := MakeTypeName(MakeLocalPackageReference(group, version), name)
	objectType := NewObjectType().WithProperties(fullName, familyName, knownAs)
	objectDefinition := MakeTypeDefinition(ref, objectType).WithDescription(description)

	g.Expect(objectDefinition.Description()).To(Equal(description))
}

/*
 * AsAst() Tests
 */

func Test_TypeDefinitionAsAst_GivenValidStruct_ReturnsNonNilResult(t *testing.T) {
	g := NewGomegaWithT(t)

	ref := MakeTypeName(MakeLocalPackageReference("group", "2020-01-01"), "name")
	definition := MakeTypeDefinition(ref, NewObjectType())
	node := definition.AsDeclarations(nil)

	g.Expect(node).NotTo(BeNil())
}

func createStringProperty(name string, description string) *PropertyDefinition {
	return NewPropertyDefinition(PropertyName(name), name, StringType).WithDescription(description)
}

func createIntProperty(name string, description string) *PropertyDefinition {
	return NewPropertyDefinition(PropertyName(name), name, IntType).WithDescription(description)
}
