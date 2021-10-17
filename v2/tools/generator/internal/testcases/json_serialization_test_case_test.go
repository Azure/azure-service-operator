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
	g := NewGomegaWithT(t)

	idFactory := astmodel.NewIdentifierFactory()

	testCaseInjector := astmodel.NewTestCaseInjector()

	currentSpec := test.CreateSpec(
		test.Pkg2020,
		"Person",
		test.FullNameProperty,
		test.KnownAsProperty,
		test.FamilyNameProperty)

	container, ok := astmodel.AsPropertyContainer(currentSpec.Type())
	g.Expect(ok).To(BeTrue())

	testcase := NewJSONSerializationTestCase(currentSpec.Name(), container, idFactory)

	currentSpec, err := testCaseInjector.Inject(currentSpec, testcase)
	g.Expect(err).To(Succeed())

	test.AssertSingleTypeDefinitionGeneratesExpectedCode(t, "person", currentSpec, test.IncludeTestFiles(), test.ExcludeCodeFiles())
}
