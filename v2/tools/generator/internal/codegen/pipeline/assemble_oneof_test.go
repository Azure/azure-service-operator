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

func Test_OneOfAssembler_AssembleOneOfs_ReturnsExpectedResult(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	idFactory := astmodel.NewIdentifierFactory()

	greek := createTestRootOneOf("Greek", test.FullNameProperty)
	alpha := createTestLeafOneOf("Alpha", greek)
	beta := createTestLeafOneOf("Beta", greek)
	gamma := createTestLeafOneOf("Gamma", greek)
	delta := createTestLeafOneOf("Delta", greek, test.StatusProperty)
	epsilon := createTestLeafOneOf("Epsilon", delta, test.KnownAsProperty) // child of delta

	defs := make(astmodel.TypeDefinitionSet, 6)
	defs.AddAll(greek, alpha, beta, gamma, delta, epsilon)

	assembler := newOneOfAssembler(defs, idFactory)
	result, err := assembler.assembleOneOfs()
	g.Expect(err).To(BeNil())

	g.Expect(result).To(HaveLen(6))

	// Check that the root oneOf is still there
	g.Expect(result.Contains(greek.Name())).To(BeTrue())

	test.AssertDefinitionHasExpectedShape(t, greek.Name().Name(), result.MustGetDefinition(greek.Name()))
	test.AssertDefinitionHasExpectedShape(t, alpha.Name().Name(), result.MustGetDefinition(alpha.Name()))
	test.AssertDefinitionHasExpectedShape(t, beta.Name().Name(), result.MustGetDefinition(beta.Name()))
	test.AssertDefinitionHasExpectedShape(t, gamma.Name().Name(), result.MustGetDefinition(gamma.Name()))
	test.AssertDefinitionHasExpectedShape(t, delta.Name().Name(), result.MustGetDefinition(delta.Name()))
	test.AssertDefinitionHasExpectedShape(t, epsilon.Name().Name(), result.MustGetDefinition(epsilon.Name()))
}

func createTestRootOneOf(
	name string,
	commonProperties ...*astmodel.PropertyDefinition,
) astmodel.TypeDefinition {
	typeName := astmodel.MakeInternalTypeName(test.Pkg2020, name)
	oneOf := astmodel.NewOneOfType(name).
		WithDiscriminatorProperty("discriminator")

	if len(commonProperties) > 0 {
		obj := astmodel.NewObjectType().WithProperties(commonProperties...)
		oneOf = oneOf.WithAdditionalPropertyObject(obj)
	}

	return astmodel.MakeTypeDefinition(typeName, oneOf)
}

func createTestLeafOneOf(
	name string,
	root astmodel.TypeDefinition,
	additionalProperties ...*astmodel.PropertyDefinition,
) astmodel.TypeDefinition {
	typeName := astmodel.MakeInternalTypeName(test.Pkg2020, name)
	oneOf := astmodel.NewOneOfType(name, root.Name(), root.Name()).
		WithDiscriminatorValue(name)

	if len(additionalProperties) > 0 {
		obj := astmodel.NewObjectType().WithProperties(additionalProperties...)
		oneOf = oneOf.WithAdditionalPropertyObject(obj)
	}

	return astmodel.MakeTypeDefinition(typeName, oneOf)
}
