/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"
)

// Common values for testing
var (
	propertyName        = PropertyName("FullName")
	propertyType        = StringType
	propertyJsonName    = "family-name"
	propertyDescription = "description"
)

/*
 * NewPropertyDefinition() tests
 */

func Test_NewPropertyDefinition_GivenValues_ReturnsInstanceWithExpectedFields(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	property := NewPropertyDefinition(propertyName, propertyJsonName, propertyType)

	g.Expect(property.propertyName).To(Equal(propertyName))
	g.Expect(property.propertyType).To(Equal(propertyType))
	g.Expect(property.tags["json"]).To(Equal([]string{propertyJsonName, "omitempty"}))
	g.Expect(property.description).To(BeEmpty())
}

func Test_NewPropertyDefinition_GivenValues_ReturnsInstanceWithExpectedGetters(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	property := NewPropertyDefinition(propertyName, propertyJsonName, propertyType)

	g.Expect(property.PropertyName()).To(Equal(propertyName))
	g.Expect(property.PropertyType()).To(Equal(propertyType))
}

/*
 * Tag tests
 */
func Test_PropertyDefinition_TagsAdded_TagsAreRenderedAsExpected(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	original := NewPropertyDefinition(propertyName, propertyJsonName, propertyType)
	updated := original.WithTag("key", "value")

	g.Expect(updated).NotTo(Equal(original))
	g.Expect(updated.renderedTags()).To(Equal(fmt.Sprintf("json:\"%s,omitempty\" key:\"value\"", propertyJsonName)))
}

func Test_PropertyDefinition_TagsAdded_TagsCanBeRetrieved(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	original := NewPropertyDefinition(propertyName, propertyJsonName, propertyType)
	updated := original.WithTag("key", "value")

	g.Expect(updated).NotTo(Equal(original))

	tags, ok := updated.Tag("key")
	g.Expect(ok).To(BeTrue())
	g.Expect(tags).To(HaveLen(1))
	g.Expect(tags[0]).To(Equal("value"))
}

func Test_PropertyDefinition_TagDoesntExist_TagsReturnsFalse(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	original := NewPropertyDefinition(propertyName, propertyJsonName, propertyType)

	_, ok := original.Tag("key")
	g.Expect(ok).To(BeFalse())
}

func Test_PropertyDefinition_MultipleTagsAdded_TagsAreRenderedCommaSeparated(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	original := NewPropertyDefinition(propertyName, propertyJsonName, propertyType)
	updated := original.WithTag("key", "value").WithTag("key", "value2").WithTag("key", "value3")

	g.Expect(updated.renderedTags()).To(Equal(fmt.Sprintf("json:\"%s,omitempty\" key:\"value,value2,value3\"", propertyJsonName)))
}

func Test_PropertyDefinition_ExistingTagAdded_TagsAreNotDuplicated(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	original := NewPropertyDefinition(propertyName, propertyJsonName, propertyType)
	updated := original.WithTag("json", propertyJsonName)

	g.Expect(updated).To(Equal(original))
	g.Expect(updated.renderedTags()).To(Equal(fmt.Sprintf("json:\"%s,omitempty\"", propertyJsonName)))
}

func Test_PropertyDefinition_TagKeyRemoved_TagIsNotRendered(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	original := NewPropertyDefinition(propertyName, propertyJsonName, propertyType)
	updated := original.WithoutTag("json", "")

	g.Expect(updated).NotTo(Equal(original))
	g.Expect(updated.renderedTags()).To(Equal(""))
}

func Test_PropertyDefinition_LastTagValueRemoved_TagIsNotRendered(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	original := NewPropertyDefinition(propertyName, propertyJsonName, propertyType)
	updated := original.WithoutTag("json", propertyJsonName).WithoutTag("json", "omitempty")

	g.Expect(updated).NotTo(Equal(original))
	g.Expect(updated.renderedTags()).To(Equal(""))
}

func Test_PropertyDefinition_TagValueRemoved_RemainingTagsAreRenderedAsExpected(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	original := NewPropertyDefinition(propertyName, propertyJsonName, propertyType).
		WithTag("key", "value1").
		WithTag("key", "value2").
		WithTag("key", "value3")
	updated := original.WithoutTag("key", "value2")

	g.Expect(updated).ToNot(Equal(original))
	g.Expect(updated.renderedTags()).To(Equal(fmt.Sprintf("json:\"%s,omitempty\" key:\"value1,value3\"", propertyJsonName)))
}

func Test_PropertyDefinition_NonExistentTagKeyRemoved_TagsAreRenderedAsExpected(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	original := NewPropertyDefinition(propertyName, propertyJsonName, propertyType)
	updated := original.WithoutTag("doesntexist", "")

	g.Expect(updated).To(Equal(original))
	g.Expect(updated.renderedTags()).To(Equal(fmt.Sprintf("json:\"%s,omitempty\"", propertyJsonName)))
}

func Test_PropertyDefinition_NonExistentTagValueRemoved_TagsAreRenderedAsExpected(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	original := NewPropertyDefinition(propertyName, propertyJsonName, propertyType)
	updated := original.WithoutTag("doesntexist", "val")

	g.Expect(updated).To(Equal(original))
	g.Expect(updated.renderedTags()).To(Equal(fmt.Sprintf("json:\"%s,omitempty\"", propertyJsonName)))
}

/*
 * WithDescription() tests
 */

func Test_PropertyDefinitionWithDescription_GivenDescription_SetsField(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	property := NewPropertyDefinition(propertyName, propertyJsonName, propertyType).WithDescription(propertyDescription)

	g.Expect(property.description).To(Equal(propertyDescription))
}

func Test_PropertyDefinitionWithDescription_GivenDescription_ReturnsDifferentReference(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	original := NewPropertyDefinition(propertyName, propertyJsonName, propertyType)
	updated := original.WithDescription(propertyDescription)

	g.Expect(updated).NotTo(Equal(original))
}

func Test_PropertyDefinitionWithDescription_GivenDescription_DoesNotModifyOriginal(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	original := NewPropertyDefinition(propertyName, propertyJsonName, propertyType)
	updated := original.WithDescription(propertyDescription)

	g.Expect(updated.description).NotTo(Equal(original.description))
}

func Test_PropertyDefinitionWithDescription_GivenEmptyDescription_ReturnsDifferentReference(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	original := NewPropertyDefinition(propertyName, propertyJsonName, propertyType).WithDescription(propertyDescription)
	updated := original.WithDescription("")

	g.Expect(updated).NotTo(Equal(original))
}

func Test_PropertyDefinitionWithNoDescription_GivenEmptyDescription_ReturnsSameReference(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	original := NewPropertyDefinition(propertyName, propertyJsonName, propertyType)
	updated := original.WithDescription("")

	g.Expect(updated).To(Equal(original))
}

func Test_PropertyDefinitionWithDescription_GivenSameDescription_ReturnsSameReference(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	original := NewPropertyDefinition(propertyName, propertyJsonName, propertyType).WithDescription(propertyDescription)
	updated := original.WithDescription(propertyDescription)

	g.Expect(updated).To(Equal(original))
}

/*
 * WithType() tests
 */

func Test_PropertyDefinitionWithType_GivenNewType_SetsFieldOnResult(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	original := NewPropertyDefinition(propertyName, propertyJsonName, propertyType)
	updated := original.WithType(IntType)

	g.Expect(updated.propertyType).To(Equal(IntType))
}

func Test_PropertyDefinitionWithType_GivenNewType_DoesNotModifyOriginal(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	original := NewPropertyDefinition(propertyName, propertyJsonName, propertyType)
	_ = original.WithType(IntType)

	g.Expect(original.propertyType).To(Equal(propertyType))
}

func Test_PropertyDefinitionWithType_GivenNewType_ReturnsDifferentReference(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	original := NewPropertyDefinition(propertyName, propertyJsonName, propertyType)
	updated := original.WithType(IntType)

	g.Expect(updated).NotTo(Equal(original))
}

func Test_PropertyDefinitionWithType_GivenSameType_ReturnsExistingReference(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	original := NewPropertyDefinition(propertyName, propertyJsonName, propertyType)
	updated := original.WithType(propertyType)

	g.Expect(updated).To(BeIdenticalTo(original))
}

/*
 * MakeRequired() Tests
 */

func Test_PropertyDefinitionMakeRequired_WhenOptional_ReturnsDifferentReference(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	original := NewPropertyDefinition(propertyName, propertyJsonName, propertyType).MakeOptional()
	updated := original.MakeRequired()

	g.Expect(updated).NotTo(BeIdenticalTo(original))
	g.Expect(updated.renderedTags()).To(Equal(fmt.Sprintf("json:\"%s,omitempty\"", propertyJsonName)))
}

func TestPropertyDefinitionMakeRequired_WhenOptional_ReturnsTypeWithIsRequiredTrue(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	original := NewPropertyDefinition(propertyName, propertyJsonName, propertyType).MakeOptional()
	updated := original.MakeRequired()

	g.Expect(updated.IsRequired()).To(BeTrue())
}

func Test_PropertyDefinitionMakeRequired_WhenRequired_ReturnsExistingReference(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	original := NewPropertyDefinition(propertyName, propertyJsonName, propertyType).MakeRequired()
	updated := original.MakeRequired()

	g.Expect(updated).To(BeIdenticalTo(original))
}

/*
 * MakeOptional() Tests
 */

func TestPropertyDefinitionMakeOptional_WhenRequired_ReturnsDifferentReference(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	original := NewPropertyDefinition(propertyName, propertyJsonName, propertyType).MakeRequired()
	updated := original.MakeOptional()

	g.Expect(updated).NotTo(BeIdenticalTo(original))
	g.Expect(updated.renderedTags()).To(Equal(fmt.Sprintf("json:\"%s,omitempty\"", propertyJsonName)))
}

func TestPropertyDefinitionMakeOptional_WhenRequired_ReturnsTypeWithIsRequiredFalse(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	original := NewPropertyDefinition(propertyName, propertyJsonName, propertyType).MakeRequired()
	updated := original.MakeOptional()

	g.Expect(updated.IsRequired()).NotTo(BeTrue())
}

func Test_PropertyDefinitionMakeOptional_WhenOptional_ReturnsExistingReference(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	original := NewPropertyDefinition(propertyName, propertyJsonName, propertyType).MakeOptional()
	updated := original.MakeOptional()

	g.Expect(updated).To(BeIdenticalTo(original))
}

func Test_PropertyDefinitionMakeOptional_WhenRequiredFalse_ReturnsSelf(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	original := NewPropertyDefinition(propertyName, propertyJsonName, propertyType)
	updated := original.MakeOptional()

	g.Expect(updated).To(BeIdenticalTo(original))
}

/*
 * MakeTypeRequired() Tests
 */

func Test_PropertyDefinitionMakeTypeRequired_WhenOptional_ReturnsDifferentReference(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	original := NewPropertyDefinition(propertyName, propertyJsonName, propertyType).MakeTypeOptional()
	updated := original.MakeTypeRequired()

	g.Expect(updated).NotTo(BeIdenticalTo(original))
	g.Expect(updated.renderedTags()).To(Equal(fmt.Sprintf("json:\"%s,omitempty\"", propertyJsonName)))
}

func TestPropertyDefinitionMakeTypeRequired_WhenOptional_ReturnsNonOptionalType(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	original := NewPropertyDefinition(propertyName, propertyJsonName, propertyType).MakeTypeOptional()
	updated := original.MakeTypeRequired()

	g.Expect(updated.PropertyType()).To(Equal(propertyType))
}

func Test_PropertyDefinitionMakeTypeRequired_WhenRequired_ReturnsExistingReference(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	original := NewPropertyDefinition(propertyName, propertyJsonName, propertyType).MakeTypeRequired()
	updated := original.MakeTypeRequired()

	g.Expect(updated).To(BeIdenticalTo(original))
}

func Test_PropertyDefinitionMakeTypeRequired_PropertyTypeArrayAndMap(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name         string
		propertyType Type
	}{
		// Expect equal to self
		{"required array property returns self", NewArrayType(propertyType)},
		{"required map property returns self", NewMapType(propertyType, propertyType)},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)
			original := NewPropertyDefinition(propertyName, propertyJsonName, c.propertyType)
			updated := original.MakeTypeRequired()

			g.Expect(updated).To(BeIdenticalTo(original))
			//g.Expect(updated.IsRequired()).To(BeFalse())
			//g.Expect(updated.propertyType).To(BeIdenticalTo(original.propertyType))

			g.Expect(updated.renderedTags()).To(Equal(fmt.Sprintf("json:\"%s,omitempty\"", propertyJsonName)))
		})
	}
}

/*
 * MakeTypeOptional() Tests
 */

func TestPropertyDefinitionMakeTypeOptional_WhenRequired_ReturnsDifferentReference(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	original := NewPropertyDefinition(propertyName, propertyJsonName, propertyType).MakeTypeRequired()
	updated := original.MakeTypeOptional()

	g.Expect(updated).NotTo(BeIdenticalTo(original))
	g.Expect(updated.renderedTags()).To(Equal(fmt.Sprintf("json:\"%s,omitempty\"", propertyJsonName)))
}

func TestPropertyDefinitionMakeTypeOptional_WhenRequired_ReturnsOptionalType(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	original := NewPropertyDefinition(propertyName, propertyJsonName, propertyType).MakeTypeRequired()
	updated := original.MakeTypeOptional()

	g.Expect(updated.PropertyType()).To(BeAssignableToTypeOf(&OptionalType{}))
}

func Test_PropertyDefinitionMakeTypeOptional_WhenOptional_ReturnsExistingReference(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	original := NewPropertyDefinition(propertyName, propertyJsonName, propertyType).MakeTypeOptional()
	updated := original.MakeTypeOptional()

	g.Expect(updated).To(BeIdenticalTo(original))
}

func Test_PropertyDefinitionMakeTypeOptional_PropertyTypeArrayAndMap(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name                  string
		propertyType          Type
		propertyRequiredFirst bool
	}{
		// Expect equal to self
		{"optional array property returns self", NewArrayType(propertyType), false},
		{"optional map property returns self", NewMapType(propertyType, propertyType), false},
		{"required array property returns self", NewArrayType(propertyType), true},
		{"required map property returns self", NewMapType(propertyType, propertyType), true},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			if c.propertyRequiredFirst {
				original := NewPropertyDefinition(propertyName, propertyJsonName, c.propertyType)
				required := original.MakeTypeRequired()
				updated := required.MakeTypeOptional()

				g.Expect(updated).To(BeIdenticalTo(original))
				g.Expect(updated.IsRequired()).To(BeFalse())
				g.Expect(updated.renderedTags()).To(Equal(fmt.Sprintf("json:\"%s,omitempty\"", propertyJsonName)))
			} else {
				original := NewPropertyDefinition(propertyName, propertyJsonName, c.propertyType)
				updated := original.MakeTypeOptional()

				g.Expect(updated).To(Equal(original))
				g.Expect(updated.renderedTags()).To(Equal(fmt.Sprintf("json:\"%s,omitempty\"", propertyJsonName)))
			}
		})
	}
}

/*
 * AsAst() Tests
 */

func Test_PropertyDefinitionAsAst_GivenValidField_ReturnsNonNilResult(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	original := NewPropertyDefinition(propertyName, propertyJsonName, propertyType).
		MakeRequired().
		WithDescription(propertyDescription)

	node := original.AsField(nil)

	g.Expect(node).NotTo(BeNil())
}

/*
 * Equals Tests
 */

func TestPropertyDefinition_Equals_WhenGivenPropertyDefinition_ReturnsExpectedResult(t *testing.T) {
	t.Parallel()

	strProperty := createStringProperty("FullName", "Full Legal Name")
	otherStrProperty := createStringProperty("FullName", "Full Legal Name")

	intProperty := createIntProperty("Age", "Age at last birthday")

	differentName := createStringProperty("Name", "Full Legal Name")
	differentType := createIntProperty("FullName", "Full Legal Name")
	differentTags := createStringProperty("FullName", "Full Legal Name").WithTag("a", "b")
	differentDescription := createStringProperty("FullName", "The whole thing")
	differentValidation := createStringProperty("FullName", "Full Legal Name").MakeRequired()

	cases := []struct {
		name          string
		thisProperty  *PropertyDefinition
		otherProperty *PropertyDefinition
		expected      bool
	}{
		// Expect equal to self
		{"Equal to self", strProperty, strProperty, true},
		{"Equal to self", intProperty, intProperty, true},
		// Expect equal to same
		{"Equal to same", strProperty, otherStrProperty, true},
		{"Equal to same", otherStrProperty, strProperty, true},
		// Expect not-equal when properties are different
		{"Not-equal if names are different", strProperty, differentName, false},
		{"Not-equal if types are different", strProperty, differentType, false},
		{"Not-equal if descriptions are different", strProperty, differentDescription, false},
		{"Not-equal if tags are different", strProperty, differentTags, false},
		{"Not-equal if validations are different", strProperty, differentValidation, false},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			areEqual := c.thisProperty.Equals(c.otherProperty, EqualityOverrides{})

			g.Expect(areEqual).To(Equal(c.expected))
		})
	}
}

func TestSettingSameFlattenValueDoesNotAllocateNewPropertyDefinition(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	strProperty := createStringProperty("FullName", "Full Legal Name")
	strPropertyFlatten := strProperty.SetFlatten(true)

	// safety check
	g.Expect(strProperty).ToNot(Equal(strPropertyFlatten))

	// actual asserts
	g.Expect(strProperty.SetFlatten(false)).To(BeIdenticalTo(strProperty))
	g.Expect(strPropertyFlatten.SetFlatten(true)).To(BeIdenticalTo(strPropertyFlatten))
}

func TestPropertyDefinitionsWithDifferentFlattenSettingsAreNotEqual(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	strProperty := createStringProperty("FullName", "Full Legal Name")
	strPropertyFlatten := strProperty.SetFlatten(true)
	g.Expect(strProperty.Equals(strPropertyFlatten, EqualityOverrides{})).To(BeFalse())
}

func TestPropertyDefinitionsWithDifferentFlattenedFromSettingsAreNotEqual(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	strProperty := createStringProperty("FullName", "Full Legal Name")
	strPropertyFlatten1 := strProperty.AddFlattenedFrom("foo")
	strPropertyFlatten2 := strProperty.AddFlattenedFrom("bar")
	g.Expect(strProperty.Equals(strPropertyFlatten1, EqualityOverrides{})).To(BeFalse())
	g.Expect(strProperty.Equals(strPropertyFlatten2, EqualityOverrides{})).To(BeFalse())
	g.Expect(strPropertyFlatten1.Equals(strPropertyFlatten2, EqualityOverrides{})).To(BeFalse())
}
