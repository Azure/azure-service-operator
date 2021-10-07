/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package functions

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/v2/internal/generator/astmodel"
	"github.com/Azure/azure-service-operator/v2/internal/generator/test"
)

func TestGolden_OriginalVersionFunction_GeneratesExpectedCode(t *testing.T) {
	g := NewGomegaWithT(t)
	idFactory := astmodel.NewIdentifierFactory()

	emptyDef := test.CreateObjectDefinition(test.Pkg2020, "Demo")
	injector := astmodel.NewFunctionInjector()

	originalVersionFunction := NewOriginalVersionFunction(idFactory)
	demoDef, err := injector.Inject(emptyDef, originalVersionFunction)
	g.Expect(err).To(Succeed())

	/*
	 * When verifying the golden file, check for what's changed
	 */
	test.AssertSingleTypeDefinitionGeneratesExpectedCode(t, "OriginalVersionFunction", demoDef, test.DiffWith(emptyDef))
}
