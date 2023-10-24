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

// Shared test values -- note that these are used by type_transformer_test.go too
var (
	person2020  = astmodel.MakeInternalTypeName(test.MakeLocalPackageReference("party", "2020-01-01"), "person")
	post2019    = astmodel.MakeInternalTypeName(test.MakeLocalPackageReference("thing", "2019-01-01"), "post")
	student2019 = astmodel.MakeInternalTypeName(test.MakeLocalPackageReference("role", "2019-01-01"), "student")
	tutor2019   = astmodel.MakeInternalTypeName(test.MakeLocalPackageReference("role", "2019-01-01"), "tutor")
)

func Test_FilterByGroup_CorrectlySelectsStructs(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	filter := config.TypeMatcher{
		Group: config.NewFieldMatcher("role"),
	}
	err := filter.Initialize()
	g.Expect(err).To(BeNil())

	// Roles should be selected
	g.Expect(filter.AppliesToType(student2019)).To(BeTrue())
	g.Expect(filter.AppliesToType(tutor2019)).To(BeTrue())

	// Party and Plays should not be selected
	g.Expect(filter.AppliesToType(person2020)).To(BeFalse())
	g.Expect(filter.AppliesToType(post2019)).To(BeFalse())
}

func Test_FilterByVersion_CorrectlySelectsStructs(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	filter := config.TypeMatcher{
		Version: config.NewFieldMatcher("v2019*"),
	}
	err := filter.Initialize()
	g.Expect(err).To(BeNil())

	// Version from 2019 should be selected
	g.Expect(filter.AppliesToType(post2019)).To(BeTrue())
	g.Expect(filter.AppliesToType(student2019)).To(BeTrue())

	// Version not from 2019 should not be selected
	g.Expect(filter.AppliesToType(person2020)).To(BeFalse())
}

func Test_FilterByName_CorrectlySelectsStructs(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	filter := config.TypeMatcher{
		Name: config.NewFieldMatcher("p*"),
	}
	err := filter.Initialize()
	g.Expect(err).To(BeNil())

	// Name starts with "p" should be selected
	g.Expect(filter.AppliesToType(person2020)).To(BeTrue())
	g.Expect(filter.AppliesToType(post2019)).To(BeTrue())

	// Name does not start with "p" should not be selected
	g.Expect(filter.AppliesToType(student2019)).To(BeFalse())
}

func Test_FilterByMultipleConditions_CorrectlySelectsStructs(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	filter := config.TypeMatcher{
		Name:    config.NewFieldMatcher("p*"),
		Version: config.NewFieldMatcher("v2019*"),
	}
	err := filter.Initialize()
	g.Expect(err).To(BeNil())

	// Version not selected by filter
	g.Expect(filter.AppliesToType(person2020)).To(BeFalse())

	// Both name and version selected by filter
	g.Expect(filter.AppliesToType(post2019)).To(BeTrue())

	// Name not selected by filter
	g.Expect(filter.AppliesToType(student2019)).To(BeFalse())
}

func Test_FiltersAreCaseInsensitive(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	filter := config.TypeMatcher{
		Group: config.NewFieldMatcher("ROLE"),
		Name:  config.NewFieldMatcher("TuToR"),
	}
	err := filter.Initialize()
	g.Expect(err).To(BeNil())

	// Tutor
	g.Expect(filter.AppliesToType(tutor2019)).To(BeTrue())

	// Others should not be selected
	g.Expect(filter.AppliesToType(person2020)).To(BeFalse())
	g.Expect(filter.AppliesToType(post2019)).To(BeFalse())
	g.Expect(filter.AppliesToType(student2019)).To(BeFalse())
}

func Test_FilterByMultipleWildcards_CorrectlySelectsStructs(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	filter := config.TypeMatcher{
		Name: config.NewFieldMatcher("p*;*t"),
	}
	err := filter.Initialize()
	g.Expect(err).To(BeNil())

	// Selected
	g.Expect(filter.AppliesToType(person2020)).To(BeTrue())
	g.Expect(filter.AppliesToType(post2019)).To(BeTrue())
	g.Expect(filter.AppliesToType(student2019)).To(BeTrue())

	// Not selected
	g.Expect(filter.AppliesToType(tutor2019)).To(BeFalse())
}
