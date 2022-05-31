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
			Group: config.NewFieldMatcher("role"),
		},
		Target: &config.TransformTarget{
			Name: "int",
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
			Version: config.NewFieldMatcher("v2019*"),
		},
		Target: &config.TransformTarget{
			Name: "int",
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
			Name: config.NewFieldMatcher("p*"),
		},
		Target: &config.TransformTarget{
			Name: "int",
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
			Name: config.NewFieldMatcher("tutor"),
		},
		Target: &config.TransformTarget{
			Group:   "role",
			Version: "2019-01-01",
			Name:    "student",
		},
	}
	err := transformer.Initialize(test.MakeLocalPackageReference)
	g.Expect(err).To(BeNil())

	// Tutor should be student
	g.Expect(transformer.TransformTypeName(tutor2019)).To(Equal(student2019))
}

func Test_TransformCanTransform_ToNestedMapType(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	transformer := config.TypeTransformer{
		TypeMatcher: config.TypeMatcher{
			Name: config.NewFieldMatcher("tutor"),
		},
		Target: &config.TransformTarget{
			Map: &config.MapType{
				Key: config.TransformTarget{
					Name: "string",
				},
				Value: config.TransformTarget{
					Map: &config.MapType{
						Key: config.TransformTarget{
							Name: "int",
						},
						Value: config.TransformTarget{
							Name: "float",
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
			Name: config.NewFieldMatcher("tutor"),
		},
		Target: &config.TransformTarget{
			Map: &config.MapType{
				Key: config.TransformTarget{
					Name: "string",
				},
				Value: config.TransformTarget{
					Map: &config.MapType{
						Value: config.TransformTarget{
							Name: "int",
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
			Name: config.NewFieldMatcher("tutor"),
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
		Property: config.NewFieldMatcher("hat"),
		Target: &config.TransformTarget{
			Name: "int",
		},
		Remove: true,
	}

	err := transformer.Initialize(test.MakeLocalPackageReference)
	g.Expect(err).To(Not(BeNil()))
	g.Expect(err).To(MatchError("remove and target can't both be set"))
}

func Test_TransformWithMultipleTargets_ReportsError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	transformer := config.TypeTransformer{
		TypeMatcher: config.TypeMatcher{
			Name: config.NewFieldMatcher("tutor"),
		},
		Target: &config.TransformTarget{
			Name: "int",
			Map: &config.MapType{
				Key: config.TransformTarget{
					Name: "string",
				},
				Value: config.TransformTarget{
					Name: "string",
				},
			},
		},
	}

	err := transformer.Initialize(test.MakeLocalPackageReference)
	g.Expect(err).To(Not(BeNil()))
	g.Expect(err.Error()).To(ContainSubstring("multiple target types defined"))
}

func Test_TransformWithNonExistentPrimitive_ReportsError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	transformer := config.TypeTransformer{
		TypeMatcher: config.TypeMatcher{
			Name: config.NewFieldMatcher("tutor"),
		},
		Target: &config.TransformTarget{
			Name: "nemo",
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
			Name: config.NewFieldMatcher("tutor"),
		},
		IfType: &config.TransformTarget{
			Name: "from",
		},
		Target: &config.TransformTarget{
			Name: "to",
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
			Name: config.NewFieldMatcher("*"),
		},
		Property: config.NewFieldMatcher("foo"),
		Target: &config.TransformTarget{
			Name: "string",
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
			Name: config.NewFieldMatcher("*"),
		},
		Property: config.NewFieldMatcher("foo*"),
		Target: &config.TransformTarget{
			Name: "string",
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
			Name: config.NewFieldMatcher("*"),
		},
		IfType: &config.TransformTarget{
			Name: "string",
		},
		Property: config.NewFieldMatcher("foo"),
		Target: &config.TransformTarget{
			Name: "int",
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

func Test_TransformDoesTransformPropertyIfTypeDoesMatch(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	transformer := config.TypeTransformer{
		TypeMatcher: config.TypeMatcher{
			Name: config.NewFieldMatcher("*"),
		},
		IfType: &config.TransformTarget{
			Name: "int",
		},
		Property: config.NewFieldMatcher("foo"),
		Target: &config.TransformTarget{
			Name: "string",
		},
	}

	err := transformer.Initialize(test.MakeLocalPackageReference)
	g.Expect(err).To(BeNil())

	typeName := student2019
	prop := astmodel.NewPropertyDefinition("foo", "foo", astmodel.IntType)
	typeDef := astmodel.NewObjectType().WithProperties(prop)

	result := transformer.TransformProperty(typeName, typeDef)
	g.Expect(result).To(Not(BeNil()))

	fooProp, ok := result.NewType.Property("foo")
	g.Expect(ok).To(BeTrue())
	g.Expect(fooProp.PropertyType()).To(Equal(astmodel.StringType))
}

func Test_TransformCanRemoveProperty(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	transformer := config.TypeTransformer{
		TypeMatcher: config.TypeMatcher{
			Name: config.NewFieldMatcher("*"),
		},
		IfType: &config.TransformTarget{
			Name: "int",
		},
		Property: config.NewFieldMatcher("foo"),
		Remove:   true,
	}

	err := transformer.Initialize(test.MakeLocalPackageReference)
	g.Expect(err).To(BeNil())

	typeName := student2019
	prop := astmodel.NewPropertyDefinition("foo", "foo", astmodel.IntType)
	typeDef := astmodel.NewObjectType().WithProperties(prop)

	result := transformer.TransformProperty(typeName, typeDef)
	g.Expect(result).To(Not(BeNil()))

	_, ok := result.NewType.Property("foo")
	g.Expect(ok).To(BeFalse())
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
