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

	cfg := config.NewObjectModelConfiguration()
	g.Expect(
		cfg.ModifyType(
			person2020.Name(),
			func(tc *config.TypeConfiguration) error {
				tc.SetSupportedFrom("beta.0")
				return nil
			})).
		To(Succeed())
	g.Expect(
		cfg.ModifyType(
			address2020.Name(),
			func(tc *config.TypeConfiguration) error {
				tc.SetSupportedFrom("beta.0")
				return nil
			})).
		To(Succeed())
	g.Expect(
		cfg.ModifyType(
			person2021.Name(),
			func(tc *config.TypeConfiguration) error {
				tc.SetSupportedFrom("beta.2")
				return nil
			})).
		To(Succeed())
	g.Expect(
		cfg.ModifyType(
			address2021.Name(),
			func(tc *config.TypeConfiguration) error {
				tc.SetSupportedFrom("beta.2")
				return nil
			})).
		To(Succeed())

	report := NewResourceVersionsReport(defs, cfg)

	var buffer strings.Builder
	g.Expect(report.WriteToBuffer(
		&buffer, "https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples")).
		To(Succeed())

	gold.Assert(t, t.Name(), []byte(buffer.String()))
}
