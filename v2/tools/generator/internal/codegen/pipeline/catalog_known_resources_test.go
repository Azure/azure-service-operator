/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/test"
)

func TestCatalogKnownResources_IndexesByARMType(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	oldPkg := test.MakeLocalPackageReference("cdn", "2023-05-01")
	newPkg := test.MakeLocalPackageReference("cdn", "2026-07-01")
	const domainARMType = "Microsoft.Cdn/profiles/customDomains"
	const endpointDomainARMType = "Microsoft.Cdn/profiles/endpoints/customDomains"
	makeResource := func(pkg astmodel.InternalPackageReference, name string, armType string) astmodel.TypeDefinition {
		resource := makeTestResource(pkg, name)
		rt := astmodel.MustBeResourceType(resource.Type()).WithARMType(armType)
		return resource.WithType(rt)
	}

	oldDomain := makeResource(oldPkg, "Profiles_CustomDomain", domainARMType)
	newDomain := makeResource(newPkg, "RenamedDomain", domainARMType)
	endpointDomain := makeResource(newPkg, "Profiles_CustomDomain", endpointDomainARMType)
	defs := make(astmodel.TypeDefinitionSet)
	defs.AddAll(oldDomain, newDomain, endpointDomain, test.CreateSpec(oldPkg, "NotAResource"))

	state, err := CatalogKnownResources().Run(context.Background(), NewState(defs))
	g.Expect(err).NotTo(HaveOccurred())
	catalog, err := GetStateData[map[string]astmodel.TypeNameSet](state, AllKnownResources)
	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(catalog).To(Equal(map[string]astmodel.TypeNameSet{
		domainARMType:         astmodel.NewTypeNameSet(oldDomain.Name(), newDomain.Name()),
		endpointDomainARMType: astmodel.NewTypeNameSet(endpointDomain.Name()),
	}))
}
