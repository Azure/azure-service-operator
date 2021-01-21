/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"github.com/pkg/errors"
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

	ref := MakeTypeName(makeTestLocalPackageReference(group, version), name)
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

	ref := MakeTypeName(makeTestLocalPackageReference(group, version), name)
	objectType := NewObjectType().WithProperties(fullName, familyName, knownAs)
	objectDefinition := MakeTypeDefinition(ref, objectType).WithDescription(description)

	g.Expect(objectDefinition.Description()).To(Equal(description))
}

/*
 * AsAst() Tests
 */

func Test_TypeDefinitionAsAst_GivenValidStruct_ReturnsNonNilResult(t *testing.T) {
	g := NewGomegaWithT(t)

	ref := MakeTypeName(makeTestLocalPackageReference("group", "2020-01-01"), "name")
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

/*
 * ApplyObjectTransformation() tests
 */

func TestApplyObjectTransformation_GivenObjectAndTransformation_AppliesTransformation(t *testing.T) {
	g := NewGomegaWithT(t)

	ref := MakeTypeName(makeTestLocalPackageReference("group", "2020-01-01"), "name")
	original := MakeTypeDefinition(ref, NewObjectType())
	property := NewStringPropertyDefinition("FullName")

	transformed, err := original.ApplyObjectTransformation(func(objectType *ObjectType) (Type, error) {
		return objectType.WithProperty(property), nil
	})

	g.Expect(err).To(BeNil())

	ot, ok := transformed.Type().(*ObjectType)
	g.Expect(ok).To(BeTrue())
	g.Expect(ot).NotTo(BeNil())

	prop, ok := ot.Property("FullName")
	g.Expect(ok).To(BeTrue())
	g.Expect(prop).NotTo(BeNil())

}

func TestApplyObjectTransformation_GivenObjectAndTransformationReturningError_ReturnsError(t *testing.T) {
	g := NewGomegaWithT(t)

	ref := MakeTypeName(makeTestLocalPackageReference("group", "2020-01-01"), "name")
	original := MakeTypeDefinition(ref, NewObjectType())

	_, err := original.ApplyObjectTransformation(func(objectType *ObjectType) (Type, error) {
		return nil, errors.New("failed")
	})

	g.Expect(err).NotTo(BeNil())
}

func TestApplyObjectTransformation_GivenNonObjectAndTransformation_ReturnsError(t *testing.T) {
	g := NewGomegaWithT(t)

	ref := MakeTypeName(makeTestLocalPackageReference("group", "2020-01-01"), "name")
	original := MakeTypeDefinition(ref, StringType)
	property := NewStringPropertyDefinition("FullName")

	_, err := original.ApplyObjectTransformation(func(objectType *ObjectType) (Type, error) {
		return objectType.WithProperty(property), nil
	})

	g.Expect(err).NotTo(BeNil())
}

/*
 * ApplyObjectTransformations() tests
 */

var (
	fullNameProperty = NewStringPropertyDefinition("FullName")
	knownAsProperty  = NewStringPropertyDefinition("KnownAs")

	injectFullName = func(objectType *ObjectType) (*ObjectType, error) {
		return objectType.WithProperty(fullNameProperty), nil
	}

	injectKnownAs = func(objectType *ObjectType) (*ObjectType, error) {
		return objectType.WithProperty(knownAsProperty), nil
	}

	failingTransform = func(objectType *ObjectType) (*ObjectType, error) {
		return nil, errors.New("bang")
	}
)

func TestApplyObjectTransformations_GivenObjectAndTransformations_AppliesTransformations(t *testing.T) {
	g := NewGomegaWithT(t)

	ref := MakeTypeName(makeTestLocalPackageReference("group", "2020-01-01"), "name")
	original := MakeTypeDefinition(ref, NewObjectType())

	transformed, err := original.ApplyObjectTransformations(injectFullName, injectKnownAs)

	g.Expect(err).To(BeNil())

	ot, ok := transformed.Type().(*ObjectType)
	g.Expect(ok).To(BeTrue())
	g.Expect(ot).NotTo(BeNil())

	fn, ok := ot.Property("FullName")
	g.Expect(ok).To(BeTrue())
	g.Expect(fn).NotTo(BeNil())

	ka, ok := ot.Property("KnownAs")
	g.Expect(ok).To(BeTrue())
	g.Expect(ka).NotTo(BeNil())
}

func TestApplyObjectTransformations_GivenObjectAndFirstTransformationReturningError_ReturnsErrorWithoutCallingSecondTransformation(t *testing.T) {
	g := NewGomegaWithT(t)

	ref := MakeTypeName(makeTestLocalPackageReference("group", "2020-01-01"), "name")
	original := MakeTypeDefinition(ref, NewObjectType())

	_, err := original.ApplyObjectTransformations(failingTransform, injectKnownAs)

	g.Expect(err).NotTo(BeNil())
}

func TestApplyObjectTransformations_GivenObjectAndSecondTransformationReturningError_ReturnsErrorFromSecondTransformation(t *testing.T) {
	g := NewGomegaWithT(t)

	ref := MakeTypeName(makeTestLocalPackageReference("group", "2020-01-01"), "name")
	original := MakeTypeDefinition(ref, NewObjectType())

	_, err := original.ApplyObjectTransformations(injectFullName, failingTransform)

	g.Expect(err).NotTo(BeNil())
}

func TestApplyObjectTransformations_GivenNonObjectAndTransformations_ReturnsError(t *testing.T) {
	g := NewGomegaWithT(t)

	ref := MakeTypeName(makeTestLocalPackageReference("group", "2020-01-01"), "name")
	original := MakeTypeDefinition(ref, StringType)

	_, err := original.ApplyObjectTransformations(injectFullName, injectKnownAs)

	g.Expect(err).NotTo(BeNil())
}
