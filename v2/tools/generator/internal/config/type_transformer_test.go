/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package config_test

import (
	"testing"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/config"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/test"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

func Test_TransformByGroup_CorrectlySelectsTypes(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	transformer := config.TypeTransformer{
		TypeMatcher: config.TypeMatcher{
			Group: newFieldMatcher("role"),
		},
		Target: &config.TransformTarget{
			Name: newFieldMatcher("int"),
		},
	}
	err := transformer.Initialize(test.MakeLocalPackageReference)
	g.Expect(err).To(BeNil())

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

	transformer := config.TypeTransformer{
		TypeMatcher: config.TypeMatcher{
			Version: newFieldMatcher("v2019*"),
		},
		Target: &config.TransformTarget{
			Name: newFieldMatcher("int"),
		},
	}
	err := transformer.Initialize(test.MakeLocalPackageReference)
	g.Expect(err).To(BeNil())

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

	transformer := config.TypeTransformer{
		TypeMatcher: config.TypeMatcher{
			Name: newFieldMatcher("p*"),
		},
		Target: &config.TransformTarget{
			Name: newFieldMatcher("int"),
		},
	}
	err := transformer.Initialize(test.MakeLocalPackageReference)
	g.Expect(err).To(BeNil())

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

	transformer := config.TypeTransformer{
		TypeMatcher: config.TypeMatcher{
			Name: newFieldMatcher("tutor"),
		},
		Target: &config.TransformTarget{
			Group:   newFieldMatcher("role"),
			Version: newFieldMatcher("2019-01-01"),
			Name:    newFieldMatcher("student"),
		},
	}
	err := transformer.Initialize(test.MakeLocalPackageReference)
	g.Expect(err).To(BeNil())

	// Tutor should be student
	g.Expect(transformer.TransformTypeName(tutor2019)).To(Equal(student2019))
}

func Test_TransformTypeName_WhenConfiguredWithMap_ReturnsExpectedMapType(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	transformer := config.TypeTransformer{
		Property: newFieldMatcher("tutor"),
		Target: &config.TransformTarget{
			Map: &config.MapType{
				Key: config.TransformTarget{
					Name: newFieldMatcher("string"),
				},
				Value: config.TransformTarget{
					Name: newFieldMatcher("int"),
				},
			},
		},
	}
	err := transformer.Initialize(test.MakeLocalPackageReference)
	g.Expect(err).To(BeNil())

	expected := astmodel.NewMapType(
		astmodel.StringType,
		astmodel.IntType)

	g.Expect(transformer.TransformTypeName(tutor2019)).To(Equal(expected))
}

func Test_TransformTypeName_WhenConfiguredWithEnum_ReturnsExpectedEnumType(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	transformer := config.TypeTransformer{
		Property: newFieldMatcher("tutor"),
		Target: &config.TransformTarget{
			Enum: &config.EnumType{
				Base: "string",
				Values: []string{
					"alpha",
					"beta",
					"preview",
				},
			},
		},
	}
	err := transformer.Initialize(test.MakeLocalPackageReference)
	g.Expect(err).To(BeNil())

	expected := astmodel.NewEnumType(
		astmodel.StringType,
		astmodel.MakeEnumValue("Alpha", "alpha"),
		astmodel.MakeEnumValue("Beta", "beta"),
		astmodel.MakeEnumValue("Preview", "preview"))

	g.Expect(transformer.TransformTypeName(tutor2019)).To(Equal(expected))
}

func Test_TransformTypeName_WhenEnumMissingBase_ReturnsExpectedError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	transformer := config.TypeTransformer{
		Property: newFieldMatcher("tutor"),
		Target: &config.TransformTarget{
			Enum: &config.EnumType{
				Values: []string{
					"alpha",
					"beta",
					"preview",
				},
			},
		},
	}
	err := transformer.Initialize(test.MakeLocalPackageReference)
	g.Expect(err).ToNot(BeNil())
	g.Expect(err.Error()).To(ContainSubstring("requires a base type"))
}

func Test_TransformTypeName_WhenEnumHasInvalidBase_ReturnsExpectedError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	transformer := config.TypeTransformer{
		Property: newFieldMatcher("tutor"),
		Target: &config.TransformTarget{
			Enum: &config.EnumType{
				Base: "flag",
				Values: []string{
					"alpha",
					"beta",
					"preview",
				},
			},
		},
	}
	err := transformer.Initialize(test.MakeLocalPackageReference)
	g.Expect(err).ToNot(BeNil())
	g.Expect(err.Error()).To(ContainSubstring("unknown primitive type"))
}

func Test_TransformCanTransform_ToNestedMapType(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	transformer := config.TypeTransformer{
		TypeMatcher: config.TypeMatcher{
			Name: newFieldMatcher("tutor"),
		},
		Target: &config.TransformTarget{
			Map: &config.MapType{
				Key: config.TransformTarget{
					Name: newFieldMatcher("string"),
				},
				Value: config.TransformTarget{
					Map: &config.MapType{
						Key: config.TransformTarget{
							Name: newFieldMatcher("int"),
						},
						Value: config.TransformTarget{
							Name: newFieldMatcher("float"),
						},
					},
				},
			},
		},
	}
	err := transformer.Initialize(test.MakeLocalPackageReference)
	g.Expect(err).To(BeNil())

	expected := astmodel.NewMapType(
		astmodel.StringType,
		astmodel.NewMapType(
			astmodel.IntType,
			astmodel.FloatType))

	g.Expect(transformer.TransformTypeName(tutor2019)).To(Equal(expected))
}

func Test_TransformWithMissingMapValue_ReportsError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	transformer := config.TypeTransformer{
		TypeMatcher: config.TypeMatcher{
			Name: newFieldMatcher("tutor"),
		},
		Target: &config.TransformTarget{
			Map: &config.MapType{
				Key: config.TransformTarget{
					Name: newFieldMatcher("string"),
				},
				Value: config.TransformTarget{
					Map: &config.MapType{
						Value: config.TransformTarget{
							Name: newFieldMatcher("int"),
						},
					},
				},
			},
		},
	}
	err := transformer.Initialize(test.MakeLocalPackageReference)
	g.Expect(err).To(Not(BeNil()))
	g.Expect(err.Error()).To(ContainSubstring("no target type found in target/map/value/map/key"))
}

func Test_TransformWithMissingTargetType_ReportsError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	transformer := config.TypeTransformer{
		TypeMatcher: config.TypeMatcher{
			Name: newFieldMatcher("tutor"),
		},
	}

	err := transformer.Initialize(test.MakeLocalPackageReference)
	g.Expect(err).To(Not(BeNil()))
	g.Expect(err.Error()).To(ContainSubstring("no target type and remove is not set"))
}

func Test_TransformWithRemoveButNoProperty_ReportsError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	transformer := config.TypeTransformer{
		Remove: true,
	}

	err := transformer.Initialize(test.MakeLocalPackageReference)
	g.Expect(err).To(Not(BeNil()))
	g.Expect(err).To(MatchError("remove is only usable with property matches"))
}

func Test_TransformWithRemoveAndTarget_ReportsError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	transformer := config.TypeTransformer{
		Property: newFieldMatcher("hat"),
		Target: &config.TransformTarget{
			Name: newFieldMatcher("int"),
		},
		Remove: true,
	}

	err := transformer.Initialize(test.MakeLocalPackageReference)
	g.Expect(err).To(Not(BeNil()))
	g.Expect(err).To(MatchError("remove and target can't both be set"))
}

func Test_TransformWithBothNameAndMapTargets_ReportsError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	transformer := config.TypeTransformer{
		TypeMatcher: config.TypeMatcher{
			Name: newFieldMatcher("tutor"),
		},
		Target: &config.TransformTarget{
			Name: newFieldMatcher("int"),
			Map: &config.MapType{
				Key: config.TransformTarget{
					Name: newFieldMatcher("string"),
				},
				Value: config.TransformTarget{
					Name: newFieldMatcher("string"),
				},
			},
		},
	}

	err := transformer.Initialize(test.MakeLocalPackageReference)
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

	transformer := config.TypeTransformer{
		TypeMatcher: config.TypeMatcher{
			Name: newFieldMatcher("tutor"),
		},
		Target: &config.TransformTarget{
			Name: newFieldMatcher("int"),
			Enum: &config.EnumType{
				Values: []string{
					"alpha",
					"beta",
					"preview",
				},
			},
		},
	}

	err := transformer.Initialize(test.MakeLocalPackageReference)
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

	transformer := config.TypeTransformer{
		TypeMatcher: config.TypeMatcher{
			Name: newFieldMatcher("tutor"),
		},
		Target: &config.TransformTarget{
			Map: &config.MapType{
				Key: config.TransformTarget{
					Name: newFieldMatcher("string"),
				},
				Value: config.TransformTarget{
					Name: newFieldMatcher("string"),
				},
			},
			Enum: &config.EnumType{
				Values: []string{
					"alpha",
					"beta",
					"preview",
				},
			},
		},
	}

	err := transformer.Initialize(test.MakeLocalPackageReference)
	g.Expect(err).To(Not(BeNil()))

	// Check contents of error message to ensure it mentions both targets, don't need exact string match
	g.Expect(err.Error()).To(SatisfyAll(
		ContainSubstring("cannot specify both"),
		ContainSubstring("Enum transformation"),
		ContainSubstring("Map transformation")))
}

func Test_TransformWithNonExistentPrimitive_ReportsError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	transformer := config.TypeTransformer{
		TypeMatcher: config.TypeMatcher{
			Name: newFieldMatcher("tutor"),
		},
		Target: &config.TransformTarget{
			Name: newFieldMatcher("nemo"),
		},
	}

	err := transformer.Initialize(test.MakeLocalPackageReference)
	g.Expect(err).To(Not(BeNil()))
	g.Expect(err.Error()).To(ContainSubstring("unknown primitive type transformation target: nemo"))
}

func Test_TransformWithIfTypeAndNoProperty_ReportsError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	transformer := config.TypeTransformer{
		TypeMatcher: config.TypeMatcher{
			Name: newFieldMatcher("tutor"),
		},
		IfType: &config.TransformTarget{
			Name: newFieldMatcher("from"),
		},
		Target: &config.TransformTarget{
			Name: newFieldMatcher("to"),
		},
	}

	err := transformer.Initialize(test.MakeLocalPackageReference)
	g.Expect(err).To(Not(BeNil()))
	g.Expect(err.Error()).To(ContainSubstring("ifType is only usable with property matches"))
}

func Test_TransformCanTransformProperty(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	transformer := config.TypeTransformer{
		TypeMatcher: config.TypeMatcher{
			Name: newFieldMatcher("*"),
		},
		Property: newFieldMatcher("foo"),
		Target: &config.TransformTarget{
			Name: newFieldMatcher("string"),
		},
	}

	err := transformer.Initialize(test.MakeLocalPackageReference)
	g.Expect(err).To(BeNil())

	typeName := student2019
	prop := astmodel.NewPropertyDefinition("foo", "foo", astmodel.IntType)
	typeDef := astmodel.NewObjectType().WithProperties(prop)

	result := transformer.TransformProperty(typeName, typeDef)
	g.Expect(result).ToNot(BeNil())

	fooProp, ok := result.NewType.Property("foo")
	g.Expect(ok).To(BeTrue())
	g.Expect(fooProp.PropertyType()).To(Equal(astmodel.StringType))
}

func Test_TransformCanTransformProperty_Wildcard(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	transformer := config.TypeTransformer{
		TypeMatcher: config.TypeMatcher{
			Name: newFieldMatcher("*"),
		},
		Property: newFieldMatcher("foo*"),
		Target: &config.TransformTarget{
			Name: newFieldMatcher("string"),
		},
	}

	err := transformer.Initialize(test.MakeLocalPackageReference)
	g.Expect(err).To(BeNil())

	typeName := student2019
	props := []*astmodel.PropertyDefinition{
		astmodel.NewPropertyDefinition("foo1", "foo1", astmodel.IntType),
		astmodel.NewPropertyDefinition("foo2", "foo2", astmodel.BoolType),
		astmodel.NewPropertyDefinition("other", "other", astmodel.FloatType),
	}
	typeDef := astmodel.NewObjectType().WithProperties(props...)

	result := transformer.TransformProperty(typeName, typeDef)
	g.Expect(result).ToNot(BeNil())

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

	transformer := config.TypeTransformer{
		TypeMatcher: config.TypeMatcher{
			Name: newFieldMatcher("*"),
		},
		IfType: &config.TransformTarget{
			Name: newFieldMatcher("string"),
		},
		Property: newFieldMatcher("foo"),
		Target: &config.TransformTarget{
			Name: newFieldMatcher("int"),
		},
	}

	err := transformer.Initialize(test.MakeLocalPackageReference)
	g.Expect(err).To(BeNil())

	typeName := student2019
	prop := astmodel.NewPropertyDefinition("foo", "foo", astmodel.IntType)
	typeDef := astmodel.NewObjectType().WithProperties(prop)

	result := transformer.TransformProperty(typeName, typeDef)
	g.Expect(result).To(BeNil()) // as ifType does not match
}

func TestTransformProperty_DoesTransformProperty_IfTypeDoesMatch(t *testing.T) {
	t.Parallel()

	transformIntToString := config.TypeTransformer{
		TypeMatcher: config.TypeMatcher{
			Name: newFieldMatcher("*"),
		},
		IfType: &config.TransformTarget{
			Name: newFieldMatcher("int"),
		},
		Property: newFieldMatcher("foo"),
		Target: &config.TransformTarget{
			Name: newFieldMatcher("string"),
		},
	}

	transformOptionalIntToOptionalString := config.TypeTransformer{
		TypeMatcher: config.TypeMatcher{
			Name: newFieldMatcher("*"),
		},
		IfType: &config.TransformTarget{
			Name:     newFieldMatcher("int"),
			Optional: true,
		},
		Property: newFieldMatcher("foo"),
		Target: &config.TransformTarget{
			Name:     newFieldMatcher("string"),
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

				g.Expect(c.transformer.Initialize(test.MakeLocalPackageReference)).To(Succeed())

				result := c.transformer.TransformProperty(student2019, c.subject)
				g.Expect(result).To(Not(BeNil()))

				prop, ok := result.NewType.Property(c.propertyToInspect)
				g.Expect(ok).To(BeTrue())
				g.Expect(prop.PropertyType()).To(Equal(c.expectedType))
			})
	}
}

func TestTransformProperty_CanRemoveProperty(t *testing.T) {
	t.Parallel()

	removeIntProperty := config.TypeTransformer{
		TypeMatcher: config.TypeMatcher{
			Name: newFieldMatcher("*"),
		},
		IfType: &config.TransformTarget{
			Name: newFieldMatcher("int"),
		},
		Property: newFieldMatcher("foo"),
		Remove:   true,
	}

	intProperty := astmodel.NewPropertyDefinition("foo", "foo", astmodel.IntType)
	objectWithIntProperty := astmodel.NewObjectType().WithProperties(intProperty)

	removeCopyProperty := config.TypeTransformer{
		TypeMatcher: config.TypeMatcher{
			Name: newFieldMatcher("*"),
		},
		IfType: &config.TransformTarget{
			Group:    newFieldMatcher("deploymenttemplate"),
			Version:  newFieldMatcher("2019-04-01"),
			Name:     newFieldMatcher("ResourceCopy"),
			Optional: true,
		},
		Property: newFieldMatcher("Copy"),
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

				g.Expect(c.transformer.Initialize(test.MakeLocalPackageReference)).To(Succeed())

				result := c.transformer.TransformProperty(student2019, c.subject)
				g.Expect(result).To(Not(BeNil()))

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

func TestTypeTarget_AppliesToType_ReturnsExpectedResult(t *testing.T) {
	t.Parallel()

	mapTarget := &config.TransformTarget{
		Map: &config.MapType{
			Key: config.TransformTarget{
				Name: newFieldMatcher("string"),
			},
			Value: config.TransformTarget{
				Name: newFieldMatcher("any"),
			},
		},
	}

	mapType := astmodel.NewMapType(astmodel.StringType, astmodel.AnyType)

	nameTarget := &config.TransformTarget{
		Group:    newFieldMatcher("definitions"),
		Version:  newFieldMatcher("v1"),
		Name:     newFieldMatcher("ResourceCopy"),
		Optional: true,
	}

	nameTargetWithWildcardVersion := &config.TransformTarget{
		Group:    newFieldMatcher("definitions"),
		Version:  newFieldMatcher("*"),
		Name:     newFieldMatcher("ResourceCopy"),
		Optional: true,
	}

	nameType := astmodel.NewOptionalType(
		astmodel.MakeInternalTypeName(
			test.MakeLocalPackageReference("definitions", "v1"),
			"ResourceCopy"))

	cases := []struct {
		name        string
		target      *config.TransformTarget
		subject     astmodel.Type
		expectation bool
	}{
		{"Matches map[string]any", mapTarget, mapType, true},
		{"Matches with exact details", nameTarget, nameType, true},
		{"Matches with version wildcard", nameTargetWithWildcardVersion, nameType, true},
	}

	for _, c := range cases {
		c := c
		t.Run(
			c.name,
			func(t *testing.T) {
				t.Parallel()
				g := NewGomegaWithT(t)

				g.Expect(c.target.AppliesToType(c.subject)).To(Equal(c.expectation))
			})
	}
}

func newFieldMatcher(field string) config.FieldMatcher {
	matcher := config.NewFieldMatcher(field)
	return matcher
}
