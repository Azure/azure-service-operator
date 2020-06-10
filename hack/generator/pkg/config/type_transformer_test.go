/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package config_test

import (
	"github.com/Azure/k8s-infra/hack/generator/pkg/config"
	"testing"

	"github.com/Azure/k8s-infra/hack/generator/pkg/astmodel"
	. "github.com/onsi/gomega"
)

func Test_TransformByGroup_CorrectlySelectsTypes(t *testing.T) {
	g := NewGomegaWithT(t)

	transformer := config.TypeTransformer{
		TypeMatcher: config.TypeMatcher{Group: "role"},
		Target: config.TransformTarget{
			Name: "int",
		},
	}
	err := transformer.Initialize()
	g.Expect(err).To(BeNil())

	// Roles should be selected
	g.Expect(transformer.TransformTypeName(student2019)).To(Equal(astmodel.IntType))
	g.Expect(transformer.TransformTypeName(tutor2019)).To(Equal(astmodel.IntType))

	// Party and Plays should not be selected
	g.Expect(transformer.TransformTypeName(person2020)).To(BeNil())
	g.Expect(transformer.TransformTypeName(post2019)).To(BeNil())
}

func Test_TransformByVersion_CorrectlySelectsTypes(t *testing.T) {
	g := NewGomegaWithT(t)

	transformer := config.TypeTransformer{
		TypeMatcher: config.TypeMatcher{Version: "2019-*"},
		Target: config.TransformTarget{
			Name: "int",
		},
	}
	err := transformer.Initialize()
	g.Expect(err).To(BeNil())

	// 2019 versions should be transformed
	g.Expect(transformer.TransformTypeName(student2019)).To(Equal(astmodel.IntType))
	g.Expect(transformer.TransformTypeName(tutor2019)).To(Equal(astmodel.IntType))
	g.Expect(transformer.TransformTypeName(post2019)).To(Equal(astmodel.IntType))

	// other versions should not
	g.Expect(transformer.TransformTypeName(person2020)).To(BeNil())
}

func Test_TransformByName_CorrectlySelectsTypes(t *testing.T) {
	g := NewGomegaWithT(t)

	transformer := config.TypeTransformer{
		TypeMatcher: config.TypeMatcher{Name: "p*"},
		Target: config.TransformTarget{
			Name: "int",
		},
	}
	err := transformer.Initialize()
	g.Expect(err).To(BeNil())

	// Names starting with p should be transformed
	g.Expect(transformer.TransformTypeName(post2019)).To(Equal(astmodel.IntType))
	g.Expect(transformer.TransformTypeName(person2020)).To(Equal(astmodel.IntType))

	// other versions should not
	g.Expect(transformer.TransformTypeName(student2019)).To(BeNil())
	g.Expect(transformer.TransformTypeName(tutor2019)).To(BeNil())
}

func Test_TransformCanTransform_ToComplexType(t *testing.T) {
	g := NewGomegaWithT(t)

	transformer := config.TypeTransformer{
		TypeMatcher: config.TypeMatcher{Name: "tutor"},
		Target: config.TransformTarget{
			PackagePath: "github.com/Azure/k8s-infra/hack/generator/apis/role/2019-01-01",
			Name:        "student",
		},
	}
	err := transformer.Initialize()
	g.Expect(err).To(BeNil())

	// Tutor should be student
	g.Expect(transformer.TransformTypeName(tutor2019)).To(Equal(student2019))
}
