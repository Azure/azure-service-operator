/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package testcases

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/test"
)

func TestGolden_JSONSerializationTestCase_AsFunc(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	idFactory := astmodel.NewIdentifierFactory()

	testCaseInjector := astmodel.NewTestCaseInjector()

	currentSpec := test.CreateSpec(
		test.Pkg2020,
		"Person",
		test.FullNameProperty,
		test.KnownAsProperty,
		test.FamilyNameProperty,
	)

	container, ok := astmodel.AsPropertyContainer(currentSpec.Type())
	g.Expect(ok).To(BeTrue())

	testcase := NewJSONSerializationTestCase(currentSpec.Name(), container, false, idFactory)

	currentSpec, err := testCaseInjector.Inject(currentSpec, testcase)
	g.Expect(err).To(Succeed())

	test.AssertSingleTypeDefinitionGeneratesExpectedCode(t, "person", currentSpec, test.IncludeTestFiles(), test.ExcludeCodeFiles())
}

func TestGolden_JSONSerializationTestCase_UInt64(t *testing.T) {
	t.Parallel()
	for _, useRapid := range []bool{false, true} {
		name := "uint64_gopter"
		if useRapid {
			name = "uint64_rapid"
		}
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)
			idFactory := astmodel.NewIdentifierFactory()
			spec := test.CreateSpec(test.Pkg2020, "Limits",
				astmodel.NewPropertyDefinition("Limit", "limit", astmodel.UInt64Type),
				astmodel.NewPropertyDefinition("OptionalLimit", "optionalLimit", astmodel.NewOptionalType(astmodel.UInt64Type)))
			container, ok := astmodel.AsPropertyContainer(spec.Type())
			g.Expect(ok).To(BeTrue())
			var testcase astmodel.TestCase = NewJSONSerializationTestCase(spec.Name(), container, false, idFactory)
			if useRapid {
				testcase = NewRapidJSONSerializationTestCase(spec.Name(), container, false, idFactory)
			}
			spec, err := astmodel.NewTestCaseInjector().Inject(spec, testcase)
			g.Expect(err).To(Succeed())
			test.AssertSingleTypeDefinitionGeneratesExpectedCode(t, name, spec, test.IncludeTestFiles(), test.ExcludeCodeFiles())
		})
	}
}
