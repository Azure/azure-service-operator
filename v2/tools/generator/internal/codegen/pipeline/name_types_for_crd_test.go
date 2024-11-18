/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/test"
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
			astmodel.MakeInternalTypeName(test.Pkg2020, "Person"),
			"Person",
			"",
			"Person",
		},
		{
			"Simple TypeName",
			astmodel.MakeInternalTypeName(test.Pkg2020, "Person"+astmodel.StatusSuffix),
			"Person",
			"STATUS",
			"Person_STATUS",
		},
		{
			"Simple TypeName",
			astmodel.MakeInternalTypeName(test.Pkg2020, "Person"+astmodel.SpecSuffix),
			"Person",
			"Spec",
			"Person_Spec",
		},
		{
			"Simple TypeName",
			astmodel.MakeInternalTypeName(test.Pkg2020, "Person"+astmodel.StatusSuffix),
			"Person",
			"STATUS",
			"Person_STATUS",
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			g := NewGomegaWithT(t)

			actual := newNameHint(c.typeName)
			g.Expect(actual.baseName).To(Equal(c.expectedBaseName))
			g.Expect(actual.suffix).To(Equal(c.expectedSuffix))

			name := actual.AsTypeName(test.Pkg2020)
			g.Expect(name.Name()).To(Equal(c.expectedName))
		})
	}
}
