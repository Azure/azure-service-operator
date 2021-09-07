/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package config_test

import (
	"testing"

	"github.com/Azure/azure-service-operator/hack/generator/pkg/config"

	"github.com/Azure/azure-service-operator/hack/generator/pkg/astmodel"
	. "github.com/onsi/gomega"
)

// Shared test values -- note that these are used by type_transformer_test.go too
var person2020 = astmodel.MakeTypeName(makeTestLocalPackageReference("party", "2020-01-01"), "person")
var post2019 = astmodel.MakeTypeName(makeTestLocalPackageReference("thing", "2019-01-01"), "post")
var student2019 = astmodel.MakeTypeName(makeTestLocalPackageReference("role", "2019-01-01"), "student")
var tutor2019 = astmodel.MakeTypeName(makeTestLocalPackageReference("role", "2019-01-01"), "tutor")

func Test_FilterByGroup_CorrectlySelectsStructs(t *testing.T) {
	g := NewWithT(t)

	filter := config.TypeMatcher{Group: "role"}
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
	g := NewWithT(t)

	filter := config.TypeMatcher{Version: "2019-*"}
	err := filter.Initialize()
	g.Expect(err).To(BeNil())

	// Version from 2019 should be selected
	g.Expect(filter.AppliesToType(post2019)).To(BeTrue())
	g.Expect(filter.AppliesToType(student2019)).To(BeTrue())

	// Version not from 2019 should not be selected
	g.Expect(filter.AppliesToType(person2020)).To(BeFalse())
}

func Test_FilterByName_CorrectlySelectsStructs(t *testing.T) {
	g := NewWithT(t)

	filter := config.TypeMatcher{Name: "p*"}
	err := filter.Initialize()
	g.Expect(err).To(BeNil())

	// Name starts with "p" should be selected
	g.Expect(filter.AppliesToType(person2020)).To(BeTrue())
	g.Expect(filter.AppliesToType(post2019)).To(BeTrue())

	// Name does not start with "p" should not be selected
	g.Expect(filter.AppliesToType(student2019)).To(BeFalse())
}

func Test_FilterByMultipleConditions_CorrectlySelectsStructs(t *testing.T) {
	g := NewWithT(t)

	filter := config.TypeMatcher{Name: "p*", Version: "2019-*"}
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
	g := NewWithT(t)

	filter := config.TypeMatcher{Group: "ROLE", Name: "TuToR"}
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
	g := NewWithT(t)

	filter := config.TypeMatcher{Name: "p*;*t"}
	err := filter.Initialize()
	g.Expect(err).To(BeNil())

	// Selected
	g.Expect(filter.AppliesToType(person2020)).To(BeTrue())
	g.Expect(filter.AppliesToType(post2019)).To(BeTrue())
	g.Expect(filter.AppliesToType(student2019)).To(BeTrue())

	// Not selected
	g.Expect(filter.AppliesToType(tutor2019)).To(BeFalse())
}
