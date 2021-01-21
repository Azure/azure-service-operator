/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package config_test

import (
	"testing"

	"github.com/Azure/k8s-infra/hack/generator/pkg/astmodel"
	"github.com/Azure/k8s-infra/hack/generator/pkg/config"

	. "github.com/onsi/gomega"
)

// TODO: in a common test function rather than in a bunch of test modules?
func makeTestLocalPackageReference(group string, version string) astmodel.LocalPackageReference {
	return astmodel.MakeLocalPackageReference("github.com/Azure/k8s-infra/hack/generated", group, version)
}

// Shared test values:
var package2019 = makeTestLocalPackageReference("group", "2019-01-01")
var person2019TypeName = astmodel.MakeTypeName(package2019, "person")
var post2019TypeName = astmodel.MakeTypeName(package2019, "post")
var student2019TypeName = astmodel.MakeTypeName(package2019, "student")

var package2020 = makeTestLocalPackageReference("group", "2020-01-01")
var address2020TypeName = astmodel.MakeTypeName(package2020, "address")
var person2020TypeName = astmodel.MakeTypeName(package2020, "person")
var professor2020TypeName = astmodel.MakeTypeName(package2020, "professor")
var student2020TypeName = astmodel.MakeTypeName(package2020, "student")
var tutor2020TypeName = astmodel.MakeTypeName(package2020, "tutor")

func Test_WithSingleFilter_FiltersExpectedTypes(t *testing.T) {
	g := NewGomegaWithT(t)
	person := person2020
	post := post2019
	student := student2019

	filter := config.ExportFilter{Action: config.ExportFilterInclude, TypeMatcher: config.TypeMatcher{Version: "2019*"}}
	c := config.NewConfiguration()
	c = c.WithExportFilters(&filter)

	g.Expect(c.ShouldExport(person)).To(Equal(config.Export))
	g.Expect(c.ShouldExport(post)).To(Equal(config.Export))
	g.Expect(c.ShouldExport(student)).To(Equal(config.Export))
}

func Test_WithMultipleFilters_FiltersExpectedTypes(t *testing.T) {
	g := NewGomegaWithT(t)
	person := person2020TypeName
	post := post2019TypeName
	student := student2019TypeName
	address := address2020TypeName

	versionFilter := config.ExportFilter{
		Action:      config.ExportFilterInclude,
		TypeMatcher: config.TypeMatcher{Version: "2019*"},
	}
	nameFilter := config.ExportFilter{
		Action:      config.ExportFilterInclude,
		TypeMatcher: config.TypeMatcher{Name: "*ss"},
	}
	c := config.NewConfiguration()
	c = c.WithExportFilters(&versionFilter, &nameFilter)

	g.Expect(c.ShouldExport(person)).To(Equal(config.Export))
	g.Expect(c.ShouldExport(post)).To(Equal(config.Export))
	g.Expect(c.ShouldExport(student)).To(Equal(config.Export))
	g.Expect(c.ShouldExport(address)).To(Equal(config.Export))
}

func Test_WithMultipleFilters_GivesPrecedenceToEarlierFilters(t *testing.T) {
	g := NewGomegaWithT(t)

	alwaysExportPerson := config.ExportFilter{
		Action:      config.ExportFilterInclude,
		TypeMatcher: config.TypeMatcher{Name: "person"}}
	exclude2019 := config.ExportFilter{
		Action:      config.ExportFilterExclude,
		TypeMatcher: config.TypeMatcher{Version: "2019-01-01"}}
	c := config.NewConfiguration()
	c = c.WithExportFilters(&alwaysExportPerson, &exclude2019)

	g.Expect(c.ShouldExport(person2019TypeName)).To(Equal(config.Export))
	g.Expect(c.ShouldExport(student2019TypeName)).To(Equal(config.Skip))

	g.Expect(c.ShouldExport(person2020TypeName)).To(Equal(config.Export))
	g.Expect(c.ShouldExport(professor2020TypeName)).To(Equal(config.Export))
	g.Expect(c.ShouldExport(tutor2020TypeName)).To(Equal(config.Export))
	g.Expect(c.ShouldExport(student2020TypeName)).To(Equal(config.Export))
}
