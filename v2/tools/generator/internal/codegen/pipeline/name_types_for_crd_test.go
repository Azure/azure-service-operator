/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/test"
	. "github.com/onsi/gomega"
	"testing"
)

func Test_NewNameHint(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name             string
		typeName         astmodel.TypeName
		expectedBaseName string
		expectedSuffix   string
		expectedName     string
	}{
		{
			"Simple TypeName",
			astmodel.MakeTypeName(test.Pkg2020, "Person"),
			"Person",
			"",
			"Person",
		},
		{
			"Simple TypeName",
			astmodel.MakeTypeName(test.Pkg2020, "Person"+astmodel.StatusSuffix),
			"Person",
			"STATUS",
			"Person_STATUS",
		},
		{
			"Simple TypeName",
			astmodel.MakeTypeName(test.Pkg2020, "Person"+astmodel.SpecSuffix),
			"Person",
			"Spec",
			"Person_Spec",
		},
		{
			"Simple TypeName",
			astmodel.MakeTypeName(test.Pkg2020, "Person"+astmodel.ArmSuffix),
			"Person",
			"ARM",
			"Person_ARM",
		},
		{
			"Simple TypeName",
			astmodel.MakeTypeName(test.Pkg2020, "Person"+astmodel.StatusSuffix+astmodel.ArmSuffix),
			"Person",
			"STATUS_ARM",
			"Person_STATUS_ARM",
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			g := NewGomegaWithT(t)

			actual := newNameHint(c.typeName)
			g.Expect(actual.baseName).To(Equal(c.expectedBaseName))
			g.Expect(actual.suffix).To(Equal(c.expectedSuffix))

			name := actual.AsTypeName(test.Pkg2020)
			g.Expect(name.Name()).To(Equal(c.expectedName))
		})
	}
}
