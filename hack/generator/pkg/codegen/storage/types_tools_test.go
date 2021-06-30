/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package storage

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/hack/generator/pkg/astmodel"
	"github.com/Azure/azure-service-operator/hack/generator/pkg/test"
)

var (
	testGroup = "microsoft.person"

	testPackage = test.MakeLocalPackageReference(testGroup, "v20200101")

	fullNameProperty = astmodel.NewPropertyDefinition("FullName", "fullName", astmodel.StringType).
				WithDescription("As would be used to address mail")

	familyNameProperty = astmodel.NewPropertyDefinition("FamilyName", "familyName", astmodel.StringType).
				WithDescription("Shared name of the family")

	knownAsProperty = astmodel.NewPropertyDefinition("KnownAs", "knownAs", astmodel.StringType).
			WithDescription("How the person is generally known")
)

func TestFindSpecTypes(t *testing.T) {
	g := NewGomegaWithT(t)

	// Define a test resource
	spec := test.CreateSpec(testPackage, "Person", fullNameProperty, familyNameProperty, knownAsProperty)
	status := test.CreateStatus(testPackage, "Person")
	resource := test.CreateResource(testPackage, "Person", spec, status)

	types := make(astmodel.Types)
	types.AddAll(resource, status, spec)

	specs := FindSpecTypes(types)

	g.Expect(specs).To(HaveLen(1))
	g.Expect(specs.Contains(spec.Name())).To(BeTrue())
}

func TestFindStatusTypes(t *testing.T) {
	g := NewGomegaWithT(t)

	// Define a test resource
	spec := test.CreateSpec(testPackage, "Person", fullNameProperty, familyNameProperty, knownAsProperty)
	status := test.CreateStatus(testPackage, "Person")
	resource := test.CreateResource(testPackage, "Person", spec, status)

	types := make(astmodel.Types)
	types.AddAll(resource, status, spec)

	statuses := FindStatusTypes(types)

	g.Expect(statuses).To(HaveLen(1))
	g.Expect(statuses.Contains(status.Name())).To(BeTrue())
}
