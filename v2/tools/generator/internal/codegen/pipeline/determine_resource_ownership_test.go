/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"testing"

	"github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

var pr astmodel.PackageReference = astmodel.MakeLocalPackageReference("prefix", "group", "v", "20000101")

func Test_FindChildren_ResourceDoesNotOwnItself(t *testing.T) {
	t.Parallel()

	g := gomega.NewWithT(t)

	resources := make(astmodel.TypeDefinitionSet)

	ownerType := astmodel.NewResourceType(nil, nil).WithARMURI("/resources/owner")
	ownerName := astmodel.MakeInternalTypeName(pr, "Owner")
	resources.Add(astmodel.MakeTypeDefinition(ownerName, ownerType))

	children := findChildren(ownerType, ownerName, resources)
	g.Expect(children).To(gomega.BeEmpty())
}

func Test_FindChildren_ResourceOwnsChild(t *testing.T) {
	t.Parallel()

	g := gomega.NewWithT(t)

	resources := make(astmodel.TypeDefinitionSet)

	ownerType := astmodel.NewResourceType(nil, nil).WithARMURI("/resources/owner")
	ownerName := astmodel.MakeInternalTypeName(pr, "Owner")
	resources.Add(astmodel.MakeTypeDefinition(ownerName, ownerType))

	childType := astmodel.NewResourceType(nil, nil).WithARMURI("/resources/owner/subresources/child")
	childName := astmodel.MakeInternalTypeName(pr, "Child")
	resources.Add(astmodel.MakeTypeDefinition(childName, childType))

	children := findChildren(ownerType, ownerName, resources)
	g.Expect(children).To(gomega.ConsistOf(childName))
}

func Test_FindChildren_ResourceOwnsChildWhenNameParametersAreDifferent(t *testing.T) {
	t.Parallel()

	g := gomega.NewWithT(t)

	resources := make(astmodel.TypeDefinitionSet)

	ownerType := astmodel.NewResourceType(nil, nil).WithARMURI("/resources/{name}")
	ownerName := astmodel.MakeInternalTypeName(pr, "Owner")
	resources.Add(astmodel.MakeTypeDefinition(ownerName, ownerType))

	childType := astmodel.NewResourceType(nil, nil).WithARMURI("/resources/{otherName}/subresources/child")
	childName := astmodel.MakeInternalTypeName(pr, "Child")
	resources.Add(astmodel.MakeTypeDefinition(childName, childType))

	children := findChildren(ownerType, ownerName, resources)
	g.Expect(children).To(gomega.ConsistOf(childName))
}

func Test_FindChildren_ResourceOwnsChildWhenNameIsDefault(t *testing.T) {
	t.Parallel()

	g := gomega.NewWithT(t)

	resources := make(astmodel.TypeDefinitionSet)

	ownerType := astmodel.NewResourceType(nil, nil).WithARMURI("/resources/default")
	ownerName := astmodel.MakeInternalTypeName(pr, "Owner")
	resources.Add(astmodel.MakeTypeDefinition(ownerName, ownerType))

	childType := astmodel.NewResourceType(nil, nil).WithARMURI("/resources/default/subresources/child")
	childName := astmodel.MakeInternalTypeName(pr, "Child")
	resources.Add(astmodel.MakeTypeDefinition(childName, childType))

	children := findChildren(ownerType, ownerName, resources)
	g.Expect(children).To(gomega.ConsistOf(childName))
}

func Test_FindChildren_ResourceDoesNotOwnGrandChild(t *testing.T) {
	t.Parallel()

	g := gomega.NewWithT(t)

	resources := make(astmodel.TypeDefinitionSet)

	ownerType := astmodel.NewResourceType(nil, nil).WithARMURI("/resources/owner")
	ownerName := astmodel.MakeInternalTypeName(pr, "Owner")
	resources.Add(astmodel.MakeTypeDefinition(ownerName, ownerType))

	grandChildType := astmodel.NewResourceType(nil, nil).WithARMURI("/resources/owner/subresources/child/subsubresources/grandchild")
	grandChildName := astmodel.MakeInternalTypeName(pr, "GrandChild")
	resources.Add(astmodel.MakeTypeDefinition(grandChildName, grandChildType))

	children := findChildren(ownerType, ownerName, resources)
	g.Expect(children).To(gomega.BeEmpty())
}

func Test_FindChildren_ResourceDoesNotOwnExtendedVersionOfName(t *testing.T) {
	t.Parallel()

	g := gomega.NewWithT(t)

	resources := make(astmodel.TypeDefinitionSet)

	ownerType := astmodel.NewResourceType(nil, nil).WithARMURI("/resources/owner")
	ownerName := astmodel.MakeInternalTypeName(pr, "Owner")
	resources.Add(astmodel.MakeTypeDefinition(ownerName, ownerType))

	grandChildType := astmodel.NewResourceType(nil, nil).WithARMURI("/resources/ownerLonger")
	grandChildName := astmodel.MakeInternalTypeName(pr, "GrandChild")
	resources.Add(astmodel.MakeTypeDefinition(grandChildName, grandChildType))

	children := findChildren(ownerType, ownerName, resources)
	g.Expect(children).To(gomega.BeEmpty())
}
