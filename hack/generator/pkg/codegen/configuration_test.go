/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package codegen

import (
	"testing"

	"github.com/Azure/k8s-infra/hack/generator/pkg/jsonast"

	"github.com/Azure/k8s-infra/hack/generator/pkg/astmodel"

	. "github.com/onsi/gomega"
)

// Shared test values:
var package2019 = astmodel.NewLocalPackageReference("group", "2019-01-01")
var person2019 = astmodel.NewStructDefinition(astmodel.NewTypeName(package2019, "person"), astmodel.EmptyStructType, false)
var post2019 = astmodel.NewStructDefinition(astmodel.NewTypeName(package2019, "post"), astmodel.EmptyStructType, false)
var student2019 = astmodel.NewStructDefinition(astmodel.NewTypeName(package2019, "student"), astmodel.EmptyStructType, false)

var package2020 = astmodel.NewLocalPackageReference("group", "2020-01-01")
var address2020 = astmodel.NewStructDefinition(astmodel.NewTypeName(package2020, "address"), astmodel.EmptyStructType, false)
var person2020 = astmodel.NewStructDefinition(astmodel.NewTypeName(package2020, "person"), astmodel.EmptyStructType, false)
var professor2020 = astmodel.NewStructDefinition(astmodel.NewTypeName(package2020, "professor"), astmodel.EmptyStructType, false)
var student2020 = astmodel.NewStructDefinition(astmodel.NewTypeName(package2020, "student"), astmodel.EmptyStructType, false)
var tutor2020 = astmodel.NewStructDefinition(astmodel.NewTypeName(package2020, "tutor"), astmodel.EmptyStructType, false)

func Test_WithSingleFilter_FiltersExpectedTypes(t *testing.T) {
	g := NewGomegaWithT(t)
	person := person2020
	post := post2019
	student := student2019

	filter := jsonast.TypeFilter{Action: jsonast.IncludeType, Version: "2019*"}
	config := NewConfiguration(&filter)

	g.Expect(config.ShouldExport(person)).To(Equal(Export))
	g.Expect(config.ShouldExport(post)).To(Equal(Export))
	g.Expect(config.ShouldExport(student)).To(Equal(Export))
}

func Test_WithMultipleFilters_FiltersExpectedTypes(t *testing.T) {
	g := NewGomegaWithT(t)
	person := person2020
	post := post2019
	student := student2019
	address := address2020

	versionFilter := jsonast.TypeFilter{
		Action:  jsonast.IncludeType,
		Version: "2019*"}
	nameFilter := jsonast.TypeFilter{
		Action: jsonast.IncludeType,
		Name:   "*ss"}
	config := NewConfiguration(&versionFilter, &nameFilter)

	g.Expect(config.ShouldExport(person)).To(Equal(Export))
	g.Expect(config.ShouldExport(post)).To(Equal(Export))
	g.Expect(config.ShouldExport(student)).To(Equal(Export))
	g.Expect(config.ShouldExport(address)).To(Equal(Export))
}

func Test_WithMultipleFilters_GivesPrecedenceToEarlierFilters(t *testing.T) {
	g := NewGomegaWithT(t)

	alwaysExportPerson := jsonast.TypeFilter{
		Action: jsonast.IncludeType,
		Name:   "person"}
	exclude2019 := jsonast.TypeFilter{
		Action:  jsonast.ExcludeType,
		Version: "2019-01-01"}
	config := NewConfiguration(&alwaysExportPerson, &exclude2019)

	g.Expect(config.ShouldExport(person2019)).To(Equal(Export))
	g.Expect(config.ShouldExport(student2019)).To(Equal(Skip))

	g.Expect(config.ShouldExport(person2020)).To(Equal(Export))
	g.Expect(config.ShouldExport(professor2020)).To(Equal(Export))
	g.Expect(config.ShouldExport(tutor2020)).To(Equal(Export))
	g.Expect(config.ShouldExport(student2020)).To(Equal(Export))
}
