/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/test"
	"testing"

	. "github.com/onsi/gomega"
)

func Test_OneOfAssembler_AssembleOneOfs_ReturnsExpectedResult(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	greek := createTestRootOneOf("Greek")
	alpha := createTestLeafOneOf("Alpha", greek)
	beta := createTestLeafOneOf("Beta", greek)
	gamma := createTestLeafOneOf("Gamma", greek)
	delta := createTestLeafOneOf("Delta", greek)

	defs := make(astmodel.TypeDefinitionSet, 5)
	defs.AddAll(greek, alpha, beta, gamma, delta)

	assembler := newOneOfAssembler(defs)
	result := assembler.assembleOneOfs()

	g.Expect(result).To(HaveLen(5))

	// Check that the root oneOf is still there
	g.Expect(result.Contains(greek.Name())).To(BeTrue())
}

func createTestRootOneOf(name string) astmodel.TypeDefinition {
	typeName := astmodel.MakeTypeName(test.Pkg2020, name)
	oneOf := astmodel.NewOneOfType(name).
		WithDiscriminatorProperty("discriminator")
	return astmodel.MakeTypeDefinition(typeName, oneOf)
}

func createTestLeafOneOf(name string, root astmodel.TypeDefinition) astmodel.TypeDefinition {
	typeName := astmodel.MakeTypeName(test.Pkg2020, name)
	oneOf := astmodel.NewOneOfType(name, root.Name(), root.Name()).
		WithDiscriminatorValue(name)
	return astmodel.MakeTypeDefinition(typeName, oneOf)
}
