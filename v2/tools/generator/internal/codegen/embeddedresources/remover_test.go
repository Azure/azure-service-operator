/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package embeddedresources

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/config"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/test"
)

func TestFindResourcesEmbeddedInParent_WhenParentTypeMissing_ReturnsTypoSuggestion(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	pkg := test.MakeLocalPackageReference("network", "2020-01-01")
	childName := astmodel.MakeInternalTypeName(pkg, "Subnet")
	parentName := astmodel.MakeInternalTypeName(pkg, "VirtualNetwork")

	defs := make(astmodel.TypeDefinitionSet)
	defs.Add(astmodel.MakeTypeDefinition(childName, astmodel.NewObjectType().WithIsResource(true)))
	defs.Add(astmodel.MakeTypeDefinition(parentName, astmodel.NewObjectType().WithIsResource(true)))

	cfg := config.NewConfiguration()
	err := cfg.ObjectModelConfiguration.ModifyType(
		childName,
		func(tc *config.TypeConfiguration) error {
			tc.ResourceEmbeddedInParent.Set("VirtalNetwork")
			return nil
		},
	)
	g.Expect(err).To(Succeed())

	_, err = findResourcesEmbeddedInParent(cfg, defs)
	g.Expect(err).To(HaveOccurred())
	g.Expect(err.Error()).To(ContainSubstring("did you mean VirtualNetwork"))
}
