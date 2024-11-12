/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package config_test

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/config"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/test"
)

func Test_TransformByGroup_CorrectlySelectsTypes(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	matcher := config.NewFieldMatcher("role")
	matcher2 := config.NewFieldMatcher("int")
	transformer := config.TypeTransformer{
		TypeMatcher: config.TypeMatcher{
			Group: matcher,
		},
		Target: &config.TransformResult{
			Name: matcher2,
		},
	}

	// Roles should be selected
	g.Expect(transformer.TransformTypeName(student2019)).To(Equal(astmodel.IntType))
	g.Expect(transformer.TransformTypeName(tutor2019)).To(Equal(astmodel.IntType))

	// Party and Plays should not be selected
	g.Expect(transformer.TransformTypeName(person2020)).To(BeNil())
	g.Expect(transformer.TransformTypeName(post2019)).To(BeNil())
}

func Test_TransformByVersion_CorrectlySelectsTypes(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	matcher := config.NewFieldMatcher("v2019*")
	matcher2 := config.NewFieldMatcher("int")
	transformer := config.TypeTransformer{
		TypeMatcher: config.TypeMatcher{
			Version: matcher,
		},
		Target: &config.TransformResult{
			Name: matcher2,
		},
	}

	// 2019 versions should be transformed
	g.Expect(transformer.TransformTypeName(student2019)).To(Equal(astmodel.IntType))
	g.Expect(transformer.TransformTypeName(tutor2019)).To(Equal(astmodel.IntType))
	g.Expect(transformer.TransformTypeName(post2019)).To(Equal(astmodel.IntType))

	// other versions should not
	g.Expect(transformer.TransformTypeName(person2020)).To(BeNil())
}

func Test_TransformByName_CorrectlySelectsTypes(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	matcher := config.NewFieldMatcher("p*")
	matcher2 := config.NewFieldMatcher("int")
	transformer := config.TypeTransformer{
		TypeMatcher: config.TypeMatcher{
			Name: matcher,
		},
		Target: &config.TransformResult{
			Name: matcher2,
		},
	}

	// Names starting with p should be transformed
	g.Expect(transformer.TransformTypeName(post2019)).To(Equal(astmodel.IntType))
	g.Expect(transformer.TransformTypeName(person2020)).To(Equal(astmodel.IntType))

	// other versions should not
	g.Expect(transformer.TransformTypeName(student2019)).To(BeNil())
	g.Expect(transformer.TransformTypeName(tutor2019)).To(BeNil())
}

func Test_TransformCanTransform_ToComplexType(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	matcher := config.NewFieldMatcher("tutor")
	matcher2 := config.NewFieldMatcher("role")
	matcher3 := config.NewFieldMatcher("2019-01-01")
	matcher4 := config.NewFieldMatcher("student")
	transformer := config.TypeTransformer{
		TypeMatcher: config.TypeMatcher{
			Name: matcher,
		},
		Target: &config.TransformResult{
			Group:   matcher2,
			Version: matcher3,
			Name:    matcher4,
		},
	}

	// Tutor should be student
	g.Expect(transformer.TransformTypeName(tutor2019)).To(Equal(student2019))
}

func Test_TransformTypeName_WhenConfiguredWithMap_ReturnsExpectedMapType(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	matcher := config.NewFieldMatcher("tutor")
	matcher2 := config.NewFieldMatcher("string")
	matcher3 := config.NewFieldMatcher("int")
	transformer := config.TypeTransformer{
		Property: matcher,
		Target: &config.TransformResult{
			Map: &config.MapResult{
				Key: config.TransformResult{
					Name: matcher2,
				},
				Value: config.TransformResult{
					Name: matcher3,
				},
			},
		},
	}

	expected := astmodel.NewMapType(
		astmodel.StringType,
		astmodel.IntType)

	g.Expect(transformer.TransformTypeName(tutor2019)).To(Equal(expected))
}

func Test_TransformTypeName_WhenConfiguredWithEnum_ReturnsExpectedEnumType(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	matcher := config.NewFieldMatcher("tutor")
	transformer := config.TypeTransformer{
		Property: matcher,
		Target: &config.TransformResult{
			Enum: &config.EnumResult{
				Base: "string",
				Values: []string{
					"alpha",
					"beta",
					"preview",
				},
			},
		},
	}

	expected := astmodel.NewEnumType(
		astmodel.StringType,
		astmodel.MakeEnumValue("Alpha", "\"alpha\""),
		astmodel.MakeEnumValue("Beta", "\"beta\""),
		astmodel.MakeEnumValue("Preview", "\"preview\""))

	g.Expect(transformer.TransformTypeName(tutor2019)).To(Equal(expected))
}

func Test_TransformTypeName_WhenEnumMissingBase_ReturnsExpectedError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	matcher := config.NewFieldMatcher("tutor")
	transformer := config.TypeTransformer{
		Property: matcher,
		Target: &config.TransformResult{
			Enum: &config.EnumResult{
				Values: []string{
					"alpha",
					"beta",
					"preview",
				},
			},
		},
	}

	_, err := transformer.TransformTypeName(tutor2019)
	g.Expect(err).ToNot(BeNil())
	g.Expect(err.Error()).To(ContainSubstring("requires a base type"))
}

func Test_TransformTypeName_WhenEnumHasInvalidBase_ReturnsExpectedError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	matcher := config.NewFieldMatcher("tutor")
	transformer := config.TypeTransformer{
		Property: matcher,
		Target: &config.TransformResult{
			Enum: &config.EnumResult{
				Base: "flag",
				Values: []string{
					"alpha",
					"beta",
					"preview",
				},
			},
		},
	}
	_, err := transformer.TransformTypeName(tutor2019)
	g.Expect(err).ToNot(BeNil())
	g.Expect(err.Error()).To(ContainSubstring("unknown primitive type"))
}

func Test_TransformCanTransform_ToNestedMapType(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	matcher := config.NewFieldMatcher("tutor")
	matcher2 := config.NewFieldMatcher("string")
	matcher3 := config.NewFieldMatcher("int")
	matcher4 := config.NewFieldMatcher("float")
	transformer := config.TypeTransformer{
		TypeMatcher: config.TypeMatcher{
			Name: matcher,
		},
		Target: &config.TransformResult{
			Map: &config.MapResult{
				Key: config.TransformResult{
					Name: matcher2,
				},
				Value: config.TransformResult{
					Map: &config.MapResult{
						Key: config.TransformResult{
							Name: matcher3,
						},
						Value: config.TransformResult{
							Name: matcher4,
						},
					},
				},
			},
		},
	}

	expected := astmodel.NewMapType(
		astmodel.StringType,
		astmodel.NewMapType(
			astmodel.IntType,
			astmodel.FloatType))

	g.Expect(transformer.TransformTypeName(tutor2019)).To(Equal(expected))
}

func Test_TransformWithMissingMapKey_ReportsError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	matcher := config.NewFieldMatcher("tutor")
	matcher2 := config.NewFieldMatcher("string")
	matcher3 := config.NewFieldMatcher("int")
	transformer := config.TypeTransformer{
		TypeMatcher: config.TypeMatcher{
			Name: matcher,
		},
		Target: &config.TransformResult{
			Map: &config.MapResult{
				Key: config.TransformResult{
					Name: matcher2,
				},
				Value: config.TransformResult{
					Map: &config.MapResult{
						Value: config.TransformResult{
							Name: matcher3,
						},
					},
				},
			},
		},
	}

	_, err := transformer.TransformTypeName(tutor2019)
	g.Expect(err).To(Not(BeNil()))
	g.Expect(err.Error()).To(ContainSubstring("no result transformation specified"))
	g.Expect(err.Error()).To(ContainSubstring("target/map/value/map/key"))
}

func Test_TransformWithMissingTargetTypeAndNoRename_ReportsError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	matcher := config.NewFieldMatcher("tutor")
	transformer := config.TypeTransformer{
		TypeMatcher: config.TypeMatcher{
			Name: matcher,
		},
	}

	_, err := transformer.TransformTypeName(tutor2019)
	g.Expect(err).To(Not(BeNil()))
	g.Expect(err.Error()).To(ContainSubstring("transformer must either rename or modify"))
}

func Test_TransformWithRemoveButNoProperty_ReportsError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	transformer := config.TypeTransformer{
		Remove: true,
	}

	_, err := transformer.TransformTypeName(tutor2019)
	g.Expect(err).To(Not(BeNil()))
	g.Expect(err).To(MatchError("remove is only usable with property transforms"))
}

func Test_TransformWithRemoveAndTarget_ReportsError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	matcher := config.NewFieldMatcher("hat")
	matcher2 := config.NewFieldMatcher("int")
	transformer := config.TypeTransformer{
		Property: matcher,
		Target: &config.TransformResult{
			Name: matcher2,
		},
		Remove: true,
	}

	_, err := transformer.TransformTypeName(tutor2019)
	g.Expect(err).To(Not(BeNil()))
	g.Expect(err).To(MatchError("remove and target can't both be set"))
}

func Test_TransformWithBothNameAndMapTargets_ReportsError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	matcher := config.NewFieldMatcher("tutor")
	matcher2 := config.NewFieldMatcher("int")
	matcher3 := config.NewFieldMatcher("string")
	matcher4 := config.NewFieldMatcher("string")
	transformer := config.TypeTransformer{
		TypeMatcher: config.TypeMatcher{
			Name: matcher,
		},
		Target: &config.TransformResult{
			Name: matcher2,
			Map: &config.MapResult{
				Key: config.TransformResult{
					Name: matcher3,
				},
				Value: config.TransformResult{
					Name: matcher4,
				},
			},
		},
	}

	_, err := transformer.TransformTypeName(tutor2019)
	g.Expect(err).To(Not(BeNil()))

	// Check contents of error message to ensure it mentions both targets, don't need exact string match
	g.Expect(err.Error()).To(SatisfyAll(
		ContainSubstring("cannot specify both"),
		ContainSubstring("Map transformation"),
		ContainSubstring("Name transformation")))
}

func Test_TransformWithBothNameAndEnumTargets_ReportsError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	matcher := config.NewFieldMatcher("tutor")
	matcher2 := config.NewFieldMatcher("int")
	transformer := config.TypeTransformer{
		TypeMatcher: config.TypeMatcher{
			Name: matcher,
		},
		Target: &config.TransformResult{
			Name: matcher2,
			Enum: &config.EnumResult{
				Values: []string{
					"alpha",
					"beta",
					"preview",
				},
			},
		},
	}

	_, err := transformer.TransformTypeName(tutor2019)
	g.Expect(err).To(Not(BeNil()))

	// Check contents of error message to ensure it mentions both targets, don't need exact string match
	g.Expect(err.Error()).To(SatisfyAll(
		ContainSubstring("cannot specify both"),
		ContainSubstring("Enum transformation"),
		ContainSubstring("Name transformation")))
}

func Test_TransformWithBothMapAndEnumTargets_ReportsError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	matcher := config.NewFieldMatcher("tutor")
	matcher2 := config.NewFieldMatcher("string")
	matcher3 := config.NewFieldMatcher("string")
	transformer := config.TypeTransformer{
		TypeMatcher: config.TypeMatcher{
			Name: matcher,
		},
		Target: &config.TransformResult{
			Map: &config.MapResult{
				Key: config.TransformResult{
					Name: matcher2,
				},
				Value: config.TransformResult{
					Name: matcher3,
				},
			},
			Enum: &config.EnumResult{
				Values: []string{
					"alpha",
					"beta",
					"preview",
				},
			},
		},
	}

	_, err := transformer.TransformTypeName(tutor2019)
	g.Expect(err).To(Not(BeNil()))

	// Check contents of error message to ensure it mentions both targets, don't need exact string match
	g.Expect(err.Error()).To(SatisfyAll(
		ContainSubstring("cannot specify both"),
		ContainSubstring("Enum transformation"),
		ContainSubstring("Map transformation")))
}

func Test_TypeTransformer_WhenTransformingTypeName_ReturnsExpectedTypeName(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name     string
		original astmodel.InternalTypeName
		newName  string
		expected astmodel.InternalTypeName
	}{
		{
			"Rename student to tutor",
			student2019,
			"tutor",
			tutor2019,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(
			c.name,
			func(t *testing.T) {
				t.Parallel()
				g := NewGomegaWithT(t)

				matcher := config.NewFieldMatcher("role")
				matcher2 := config.NewFieldMatcher(c.newName)
				transformer := config.TypeTransformer{
					TypeMatcher: config.TypeMatcher{
						Group: matcher,
					},
					Target: &config.TransformResult{
						Name: matcher2,
					},
				}

				actual, err := transformer.TransformTypeName(c.original)
				g.Expect(actual).To(Equal(c.expected))
				g.Expect(err).To(Succeed())
			})
	}
}

func Test_TransformCanTransformProperty(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	matcher := config.NewFieldMatcher("*")
	matcher2 := config.NewFieldMatcher("foo")
	matcher3 := config.NewFieldMatcher("string")
	transformer := config.TypeTransformer{
		TypeMatcher: config.TypeMatcher{
			Name: matcher,
		},
		Property: matcher2,
		Target: &config.TransformResult{
			Name: matcher3,
		},
	}

	typeName := student2019
	prop := astmodel.NewPropertyDefinition("foo", "foo", astmodel.IntType)
	typeDef := astmodel.NewObjectType().WithProperties(prop)

	result, err := transformer.TransformProperty(typeName, typeDef)
	g.Expect(result).ToNot(BeNil())
	g.Expect(err).ToNot(HaveOccurred())

	fooProp, ok := result.NewType.Property("foo")
	g.Expect(ok).To(BeTrue())
	g.Expect(fooProp.PropertyType()).To(Equal(astmodel.StringType))
}

func Test_TransformCanTransformProperty_Wildcard(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	matcher := config.NewFieldMatcher("*")
	matcher2 := config.NewFieldMatcher("foo*")
	matcher3 := config.NewFieldMatcher("string")
	transformer := config.TypeTransformer{
		TypeMatcher: config.TypeMatcher{
			Name: matcher,
		},
		Property: matcher2,
		Target: &config.TransformResult{
			Name: matcher3,
		},
	}

	typeName := student2019
	props := []*astmodel.PropertyDefinition{
		astmodel.NewPropertyDefinition("foo1", "foo1", astmodel.IntType),
		astmodel.NewPropertyDefinition("foo2", "foo2", astmodel.BoolType),
		astmodel.NewPropertyDefinition("other", "other", astmodel.FloatType),
	}
	typeDef := astmodel.NewObjectType().WithProperties(props...)

	result, err := transformer.TransformProperty(typeName, typeDef)
	g.Expect(result).ToNot(BeNil())
	g.Expect(err).ToNot(HaveOccurred())

	foo1Prop, ok := result.NewType.Property("foo1")
	g.Expect(ok).To(BeTrue())
	g.Expect(foo1Prop.PropertyType()).To(Equal(astmodel.StringType))

	foo2Prop, ok := result.NewType.Property("foo2")
	g.Expect(ok).To(BeTrue())
	g.Expect(foo2Prop.PropertyType()).To(Equal(astmodel.StringType))

	otherProp, ok := result.NewType.Property("other")
	g.Expect(ok).To(BeTrue())
	g.Expect(otherProp.PropertyType()).To(Equal(astmodel.FloatType))
}

func Test_TransformDoesNotTransformPropertyIfTypeDoesNotMatch(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	matcher := config.NewFieldMatcher("*")
	matcher2 := config.NewFieldMatcher("string")
	matcher3 := config.NewFieldMatcher("foo")
	matcher4 := config.NewFieldMatcher("int")
	transformer := config.TypeTransformer{
		TypeMatcher: config.TypeMatcher{
			Name: matcher,
		},
		IfType: &config.TransformSelector{
			Name: matcher2,
		},
		Property: matcher3,
		Target: &config.TransformResult{
			Name: matcher4,
		},
	}

	typeName := student2019
	prop := astmodel.NewPropertyDefinition("foo", "foo", astmodel.IntType)
	typeDef := astmodel.NewObjectType().WithProperties(prop)

	result, err := transformer.TransformProperty(typeName, typeDef)
	g.Expect(result).To(BeNil()) // as ifType does not match
	g.Expect(err).ToNot(HaveOccurred())
}

func TestTransformProperty_DoesTransformProperty_IfTypeDoesMatch(t *testing.T) {
	t.Parallel()

	matcher := config.NewFieldMatcher("*")
	matcher2 := config.NewFieldMatcher("int")
	matcher3 := config.NewFieldMatcher("foo")
	matcher4 := config.NewFieldMatcher("string")
	transformIntToString := config.TypeTransformer{
		TypeMatcher: config.TypeMatcher{
			Name: matcher,
		},
		IfType: &config.TransformSelector{
			Name: matcher2,
		},
		Property: matcher3,
		Target: &config.TransformResult{
			Name: matcher4,
		},
	}

	matcher5 := config.NewFieldMatcher("*")
	matcher6 := config.NewFieldMatcher("int")
	matcher7 := config.NewFieldMatcher("foo")
	matcher8 := config.NewFieldMatcher("string")
	transformOptionalIntToOptionalString := config.TypeTransformer{
		TypeMatcher: config.TypeMatcher{
			Name: matcher5,
		},
		IfType: &config.TransformSelector{
			Name:     matcher6,
			Optional: true,
		},
		Property: matcher7,
		Target: &config.TransformResult{
			Name:     matcher8,
			Optional: true,
		},
	}

	intProperty := astmodel.NewPropertyDefinition("foo", "foo", astmodel.IntType)
	objectWithInt := astmodel.NewObjectType().WithProperties(intProperty)

	optionalIntProperty := astmodel.NewPropertyDefinition("foo", "foo", astmodel.OptionalIntType)
	objectWithOptionalInt := astmodel.NewObjectType().WithProperties(optionalIntProperty)

	cases := []struct {
		name              string
		transformer       config.TypeTransformer
		subject           *astmodel.ObjectType
		propertyToInspect astmodel.PropertyName
		expectedType      astmodel.Type
	}{
		{
			"Int to string",
			transformIntToString,
			objectWithInt,
			"foo",
			astmodel.StringType,
		},
		{
			"optional Int to optional string",
			transformOptionalIntToOptionalString,
			objectWithOptionalInt,
			"foo",
			astmodel.OptionalStringType,
		},
	}

	for _, c := range cases {
		c := c

		t.Run(
			c.name,
			func(t *testing.T) {
				t.Parallel()
				g := NewGomegaWithT(t)

				result, err := c.transformer.TransformProperty(student2019, c.subject)
				g.Expect(result).To(Not(BeNil()))
				g.Expect(err).ToNot(HaveOccurred())

				prop, ok := result.NewType.Property(c.propertyToInspect)
				g.Expect(ok).To(BeTrue())
				g.Expect(prop.PropertyType()).To(Equal(c.expectedType))
			})
	}
}

func TestTransformProperty_CanRemoveProperty(t *testing.T) {
	t.Parallel()

	matcher := config.NewFieldMatcher("*")
	matcher2 := config.NewFieldMatcher("int")
	matcher3 := config.NewFieldMatcher("foo")
	removeIntProperty := config.TypeTransformer{
		TypeMatcher: config.TypeMatcher{
			Name: matcher,
		},
		IfType: &config.TransformSelector{
			Name: matcher2,
		},
		Property: matcher3,
		Remove:   true,
	}

	intProperty := astmodel.NewPropertyDefinition("foo", "foo", astmodel.IntType)
	objectWithIntProperty := astmodel.NewObjectType().WithProperties(intProperty)

	matcher4 := config.NewFieldMatcher("*")
	matcher5 := config.NewFieldMatcher("deploymenttemplate")
	matcher6 := config.NewFieldMatcher("2019-04-01")
	matcher7 := config.NewFieldMatcher("ResourceCopy")
	matcher8 := config.NewFieldMatcher("Copy")
	removeCopyProperty := config.TypeTransformer{
		TypeMatcher: config.TypeMatcher{
			Name: matcher4,
		},
		IfType: &config.TransformSelector{
			Group:    matcher5,
			Version:  matcher6,
			Name:     matcher7,
			Optional: true,
		},
		Property: matcher8,
		Remove:   true,
	}

	resourceCopyType := astmodel.MakeInternalTypeName(
		test.MakeLocalPackageReference("deploymenttemplate", "2019-04-01"),
		"ResourceCopy")
	copyProperty := astmodel.NewPropertyDefinition("Copy", "copy", astmodel.NewOptionalType(resourceCopyType))
	objectWithCopyProperty := astmodel.NewObjectType().WithProperties(copyProperty)

	cases := []struct {
		name              string
		transformer       config.TypeTransformer
		subject           *astmodel.ObjectType
		propertyToInspect astmodel.PropertyName
	}{
		{
			"Removes int property",
			removeIntProperty,
			objectWithIntProperty,
			"foo",
		},
		{
			"Removes Copy property",
			removeCopyProperty,
			objectWithCopyProperty,
			"Copy",
		},
	}

	for _, c := range cases {
		c := c

		t.Run(
			c.name,
			func(t *testing.T) {
				t.Parallel()
				g := NewGomegaWithT(t)

				result, err := c.transformer.TransformProperty(student2019, c.subject)
				g.Expect(result).To(Not(BeNil()))
				g.Expect(err).ToNot(HaveOccurred())

				_, ok := result.NewType.Property(c.propertyToInspect)
				g.Expect(ok).To(BeFalse())
			})
	}
}

func Test_TransformResult_String(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	result := config.PropertyTransformResult{
		TypeName:        student2019,
		Property:        "HairColour",
		NewPropertyType: astmodel.StringType,
		Because:         "string is better",
	}
	g.Expect(result.String()).To(Equal(test.MakeLocalPackageReference("role", "2019-01-01").PackagePath() + "/student.HairColour -> string because string is better"))
}

func Test_TransformResult_StringRemove(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	result := config.PropertyTransformResult{
		TypeName: student2019,
		Property: "HairColour",
		Removed:  true,
		Because:  "it's irrelevant",
	}
	g.Expect(result.String()).To(Equal(test.MakeLocalPackageReference("role", "2019-01-01").PackagePath() + "/student.HairColour removed because it's irrelevant"))
}
