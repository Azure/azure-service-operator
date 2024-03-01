/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"testing"

	"github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/config"
)

var pr astmodel.InternalPackageReference = astmodel.MakeLocalPackageReference("prefix", "group", "v", "20000101")

func Test_FindChildren_ResourceDoesNotOwnItself(t *testing.T) {
	t.Parallel()

	g := gomega.NewWithT(t)

	cfg := config.NewConfiguration()
	resources := make(astmodel.TypeDefinitionSet)

	ownerType := astmodel.NewResourceType(nil, nil).WithARMURI("/resources/owner")
	ownerName := astmodel.MakeInternalTypeName(pr, "Owner")
	owner := astmodel.MakeTypeDefinition(ownerName, ownerType)
	resources.Add(owner)

	stage := newOwnershipStage(cfg, resources)

	children := stage.findChildren(owner)
	g.Expect(children).To(gomega.BeEmpty())
}

func Test_FindChildren_ResourceOwnsChild(t *testing.T) {
	t.Parallel()

	g := gomega.NewWithT(t)

	cfg := config.NewConfiguration()
	resources := make(astmodel.TypeDefinitionSet)

	ownerType := astmodel.NewResourceType(nil, nil).WithARMURI("/resources/owner")
	ownerName := astmodel.MakeInternalTypeName(pr, "Owner")
	owner := astmodel.MakeTypeDefinition(ownerName, ownerType)
	resources.Add(owner)

	childType := astmodel.NewResourceType(nil, nil).WithARMURI("/resources/owner/subresources/child")
	childName := astmodel.MakeInternalTypeName(pr, "Child")
	child := astmodel.MakeTypeDefinition(childName, childType)
	resources.Add(child)

	stage := newOwnershipStage(cfg, resources)
	children := stage.findChildren(owner)
	g.Expect(children).To(gomega.ConsistOf(childName))
}

func Test_FindChildren_ResourceOwnsChildWhenNameParametersAreDifferent(t *testing.T) {
	t.Parallel()

	g := gomega.NewWithT(t)

	cfg := config.NewConfiguration()
	resources := make(astmodel.TypeDefinitionSet)

	ownerType := astmodel.NewResourceType(nil, nil).WithARMURI("/resources/{name}")
	ownerName := astmodel.MakeInternalTypeName(pr, "Owner")
	owner := astmodel.MakeTypeDefinition(ownerName, ownerType)
	resources.Add(owner)

	childType := astmodel.NewResourceType(nil, nil).WithARMURI("/resources/{otherName}/subresources/child")
	childName := astmodel.MakeInternalTypeName(pr, "Child")
	child := astmodel.MakeTypeDefinition(childName, childType)
	resources.Add(child)

	stage := newOwnershipStage(cfg, resources)
	children := stage.findChildren(owner)
	g.Expect(children).To(gomega.ConsistOf(childName))
}

func Test_FindChildren_ResourceOwnsChildWhenNameIsDefault(t *testing.T) {
	t.Parallel()

	g := gomega.NewWithT(t)

	cfg := config.NewConfiguration()
	resources := make(astmodel.TypeDefinitionSet)

	ownerType := astmodel.NewResourceType(nil, nil).WithARMURI("/resources/default")
	ownerName := astmodel.MakeInternalTypeName(pr, "Owner")
	owner := astmodel.MakeTypeDefinition(ownerName, ownerType)
	resources.Add(owner)

	childType := astmodel.NewResourceType(nil, nil).WithARMURI("/resources/default/subresources/child")
	childName := astmodel.MakeInternalTypeName(pr, "Child")
	child := astmodel.MakeTypeDefinition(childName, childType)
	resources.Add(child)

	stage := newOwnershipStage(cfg, resources)
	children := stage.findChildren(owner)
	g.Expect(children).To(gomega.ConsistOf(childName))
}

func Test_FindChildren_ResourceDoesNotOwnGrandChild(t *testing.T) {
	t.Parallel()

	g := gomega.NewWithT(t)

	cfg := config.NewConfiguration()
	resources := make(astmodel.TypeDefinitionSet)

	ownerType := astmodel.NewResourceType(nil, nil).WithARMURI("/resources/owner")
	ownerName := astmodel.MakeInternalTypeName(pr, "Owner")
	owner := astmodel.MakeTypeDefinition(ownerName, ownerType)
	resources.Add(owner)

	grandChildType := astmodel.NewResourceType(nil, nil).WithARMURI("/resources/owner/subresources/child/subsubresources/grandchild")
	grandChildName := astmodel.MakeInternalTypeName(pr, "GrandChild")
	grandChild := astmodel.MakeTypeDefinition(grandChildName, grandChildType)
	resources.Add(grandChild)

	stage := newOwnershipStage(cfg, resources)
	children := stage.findChildren(owner)
	g.Expect(children).To(gomega.BeEmpty())
}

func Test_FindChildren_ResourceDoesNotOwnExtendedVersionOfName(t *testing.T) {
	t.Parallel()

	g := gomega.NewWithT(t)

	cfg := config.NewConfiguration()
	resources := make(astmodel.TypeDefinitionSet)

	ownerType := astmodel.NewResourceType(nil, nil).WithARMURI("/resources/owner")
	ownerName := astmodel.MakeInternalTypeName(pr, "Owner")
	owner := astmodel.MakeTypeDefinition(ownerName, ownerType)
	resources.Add(owner)

	grandChildType := astmodel.NewResourceType(nil, nil).WithARMURI("/resources/ownerLonger")
	grandChildName := astmodel.MakeInternalTypeName(pr, "GrandChild")
	grandChild := astmodel.MakeTypeDefinition(grandChildName, grandChildType)
	resources.Add(grandChild)

	stage := newOwnershipStage(cfg, resources)
	children := stage.findChildren(owner)
	g.Expect(children).To(gomega.BeEmpty())
}
