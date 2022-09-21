/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package reporting

import (
	"bytes"
	"github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/test"
	"github.com/sebdah/goldie/v2"
	"testing"
)

func Test_TypeCatalogReport_GivenTypes_ShowsExpectedDetails(t *testing.T) {
	golden := goldie.New(t)
	g := gomega.NewWithT(t)

	defs := createDefinitionSet()

	var content bytes.Buffer
	rpt := NewTypeCatalogReport(defs)
	g.Expect(rpt.WriteTo(&content)).To(gomega.Succeed())

	golden.Assert(t, t.Name(), content.Bytes())
}

func Test_TypeCatalogReport_GivenTypes_WhenInlined_ShowsExpectedDetails(t *testing.T) {
	golden := goldie.New(t)
	g := gomega.NewWithT(t)

	defs := createDefinitionSet()

	var content bytes.Buffer
	rpt := NewTypeCatalogReport(defs)
	rpt.InlineTypes()

	g.Expect(rpt.WriteTo(&content)).To(gomega.Succeed())
	golden.Assert(t, t.Name(), content.Bytes())
}

func createDefinitionSet() astmodel.TypeDefinitionSet {
	testSpec := test.CreateSpec(
		test.Pkg2020,
		"TestResource",
		test.FullNameProperty,
		test.FamilyNameProperty,
		test.KnownAsProperty)

	testStatus := test.CreateStatus(
		test.Pkg2020,
		"TestResource")

	testResource := test.CreateResource(
		test.Pkg2020,
		"TestResource",
		testSpec,
		testStatus)

	defs := make(astmodel.TypeDefinitionSet)
	defs.Add(testResource)
	defs.Add(testSpec)
	defs.Add(testStatus)
	return defs
}
