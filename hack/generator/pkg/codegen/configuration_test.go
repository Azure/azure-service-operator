/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package codegen

import (
	"github.com/Azure/k8s-infra/hack/generator/pkg/jsonast"
	"testing"

	"github.com/Azure/k8s-infra/hack/generator/pkg/astmodel"

	. "github.com/onsi/gomega"
)

// Shared test values:
var person2019 = astmodel.NewStructDefinition(astmodel.NewStructReference("person", "group", "2019-01-01", false), astmodel.EmptyStructType)
var post2019 = astmodel.NewStructDefinition(astmodel.NewStructReference("post", "group", "2019-01-01", false), astmodel.EmptyStructType)
var student2019 = astmodel.NewStructDefinition(astmodel.NewStructReference("student", "group", "2019-01-01", false), astmodel.EmptyStructType)

var address2020 = astmodel.NewStructDefinition(astmodel.NewStructReference("address", "group", "2020-01-01", false), astmodel.EmptyStructType)
var person2020 = astmodel.NewStructDefinition(astmodel.NewStructReference("person", "group", "2020-01-01", false), astmodel.EmptyStructType)
var professor2020 = astmodel.NewStructDefinition(astmodel.NewStructReference("professor", "group", "2020-01-01", false), astmodel.EmptyStructType)
var student2020 = astmodel.NewStructDefinition(astmodel.NewStructReference("student", "group", "2020-01-01", false), astmodel.EmptyStructType)
var tutor2020 = astmodel.NewStructDefinition(astmodel.NewStructReference("tutor", "group", "2020-01-01", false), astmodel.EmptyStructType)

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
