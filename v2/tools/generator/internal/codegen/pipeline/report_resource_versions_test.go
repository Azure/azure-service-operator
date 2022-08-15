/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"strings"
	"testing"

	. "github.com/onsi/gomega"
	"github.com/sebdah/goldie/v2"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/config"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/test"
)

func TestGolden_ReportResourceVersions(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	gold := goldie.New(t)

	person2020desc := []string{
		"This is an older version",
		"of the Person resource",
	}
	person2020 := test.CreateResource(
		test.Pkg2020,
		"Person",
		test.CreateSpec(test.Pkg2020, "Person"),
		test.CreateStatus(test.Pkg2020, "Person")).
		WithDescription(person2020desc)

	address2020 := test.CreateResource(
		test.Pkg2020,
		"Address",
		test.CreateSpec(test.Pkg2020, "Address"),
		test.CreateStatus(test.Pkg2020, "Address"))

	person2021desc := []string{
		"This is a newer version",
		"of the Person resource",
	}
	person2021 := test.CreateResource(
		test.Pkg2021,
		"Person",
		test.CreateSpec(test.Pkg2021, "Person"),
		test.CreateStatus(test.Pkg2021, "Person")).
		WithDescription(person2021desc)

	address2021 := test.CreateResource(
		test.Pkg2021,
		"Address",
		test.CreateSpec(test.Pkg2021, "Address"),
		test.CreateStatus(test.Pkg2021, "Address"))

	defs := make(astmodel.TypeDefinitionSet)
	defs.AddAll(person2020, address2020, person2021, address2021)

	// utility function used to configure a which ASO version from which a resource was supported
	supportedFrom := func(from string) func(tc *config.TypeConfiguration) error {
		return func(tc *config.TypeConfiguration) error {
			tc.SetSupportedFrom(from)
			return nil
		}
	}

	cfg := config.NewConfiguration()
	cfg.RootURL = "https://github.com/Azure/azure-service-operator/tree/main/v2"
	// TODO: As these are the test resources, SampleLink would not render in the test as we're not constructing just a link anymore.
	// TODO: We walk through the samples directory to look for actual present sample and then generate a link for them.
	cfg.SamplesPath = "../../../../../config/samples"

	omc := cfg.ObjectModelConfiguration
	g.Expect(omc.ModifyType(person2020.Name(), supportedFrom("beta.0"))).To(Succeed())
	g.Expect(omc.ModifyType(address2020.Name(), supportedFrom("beta.0"))).To(Succeed())
	g.Expect(omc.ModifyType(person2021.Name(), supportedFrom("beta.2"))).To(Succeed())
	g.Expect(omc.ModifyType(address2021.Name(), supportedFrom("beta.2"))).To(Succeed())

	srr := cfg.SupportedResourcesReport
	srr.Introduction = "These are the resources with Azure Service Operator support."

	report := NewResourceVersionsReport(defs, cfg)

	var buffer strings.Builder
	g.Expect(report.WriteToBuffer(
		&buffer)).
		To(Succeed())

	gold.Assert(t, t.Name(), []byte(buffer.String()))
}
