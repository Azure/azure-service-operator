/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package jsonast

import (
	"testing"

	"github.com/Azure/k8s-infra/hack/generator/pkg/astmodel"

	. "github.com/onsi/gomega"
)

func Test_WithSingleFilter_FiltersExpectedTypes(t *testing.T) {
	g := NewGomegaWithT(t)
	person := astmodel.NewStructDefinition(astmodel.NewStructReference("person", "group", "2020-01-01"))
	post := astmodel.NewStructDefinition(astmodel.NewStructReference("post", "group", "2019-01-01"))
	student := astmodel.NewStructDefinition(astmodel.NewStructReference("student", "group", "2019-01-01"))

	filter := TypeFilter{Action: IncludeType, Version: "2019*"}
	config := NewExportConfiguration(&filter)

	g.Expect(config.ShouldExport(person)).To(Equal(Export))
	g.Expect(config.ShouldExport(post)).To(Equal(Export))
	g.Expect(config.ShouldExport(student)).To(Equal(Export))
}

func Test_WithMultipleFilters_FiltersExpectedTypes(t *testing.T) {
	g := NewGomegaWithT(t)
	person := astmodel.NewStructDefinition(astmodel.NewStructReference("person", "group", "2020-01-01"))
	post := astmodel.NewStructDefinition(astmodel.NewStructReference("post", "group", "2019-01-01"))
	student := astmodel.NewStructDefinition(astmodel.NewStructReference("student", "group", "2019-01-01"))
	address := astmodel.NewStructDefinition(astmodel.NewStructReference("address", "group", "2020-01-01"))

	versionFilter := TypeFilter{
		Action:  IncludeType,
		Version: "2019*"}
	nameFilter := TypeFilter{
		Action: IncludeType,
		Name:   "*ss"}
	config := NewExportConfiguration(&versionFilter, &nameFilter)

	g.Expect(config.ShouldExport(person)).To(Equal(Export))
	g.Expect(config.ShouldExport(post)).To(Equal(Export))
	g.Expect(config.ShouldExport(student)).To(Equal(Export))
	g.Expect(config.ShouldExport(address)).To(Equal(Export))
}

func Test_WithMultipleFilters_GivesPrecedenceToEarlierFilters(t *testing.T) {
	g := NewGomegaWithT(t)

	person2019 := astmodel.NewStructDefinition(astmodel.NewStructReference("person", "group", "2019-01-01"))
	student2019 := astmodel.NewStructDefinition(astmodel.NewStructReference("student", "group", "2019-01-01"))

	person2020 := astmodel.NewStructDefinition(astmodel.NewStructReference("person", "group", "2020-01-01"))
	professor2020 := astmodel.NewStructDefinition(astmodel.NewStructReference("professor", "group", "2020-01-01"))
	tutor2020 := astmodel.NewStructDefinition(astmodel.NewStructReference("tutor", "group", "2020-01-01"))
	student2020 := astmodel.NewStructDefinition(astmodel.NewStructReference("student", "group", "2020-01-01"))

	alwaysExportPerson := TypeFilter{
		Action: IncludeType,
		Name:   "person"}
	exclude2019 := TypeFilter{
		Action:  ExcludeType,
		Version: "2019-01-01"}
	config := NewExportConfiguration(&alwaysExportPerson, &exclude2019)

	g.Expect(config.ShouldExport(person2019)).To(Equal(Export))
	g.Expect(config.ShouldExport(student2019)).To(Equal(Skip))

	g.Expect(config.ShouldExport(person2020)).To(Equal(Export))
	g.Expect(config.ShouldExport(professor2020)).To(Equal(Export))
	g.Expect(config.ShouldExport(tutor2020)).To(Equal(Export))
	g.Expect(config.ShouldExport(student2020)).To(Equal(Export))
}
