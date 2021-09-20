/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"strings"
	"testing"

	"github.com/sebdah/goldie/v2"

	"github.com/Azure/azure-service-operator/hack/generator/pkg/astmodel"
	"github.com/Azure/azure-service-operator/hack/generator/pkg/test"
)

func TestGolden_ReportResourceVersions(t *testing.T) {
	g := goldie.New(t)

	person2020 := test.CreateResource(
		test.Pkg2020,
		"Person",
		test.CreateSpec(test.Pkg2020, "Person"),
		test.CreateStatus(test.Pkg2020, "Person"))

	address2020 := test.CreateResource(
		test.Pkg2020,
		"Address",
		test.CreateSpec(test.Pkg2020, "Address"),
		test.CreateStatus(test.Pkg2020, "Address"))

	person2021 := test.CreateResource(
		test.Pkg2021,
		"Person",
		test.CreateSpec(test.Pkg2021, "Person"),
		test.CreateStatus(test.Pkg2021, "Person"))

	address2021 := test.CreateResource(
		test.Pkg2021,
		"Address",
		test.CreateSpec(test.Pkg2021, "Address"),
		test.CreateStatus(test.Pkg2021, "Address"))

	types := make(astmodel.Types)
	types.AddAll(person2020, address2020, person2021, address2021)

	report := NewResourceVersionsReport(types)

	var buffer strings.Builder
	report.WriteToBuffer(&buffer)

	g.Assert(t, t.Name(), []byte(buffer.String()))
}
