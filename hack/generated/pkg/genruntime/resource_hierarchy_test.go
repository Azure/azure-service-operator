/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package genruntime_test

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/hack/generated/pkg/genruntime"
)

func Test_ResourceHierarchy_ResourceGroupOnly(t *testing.T) {
	g := NewGomegaWithT(t)

	resourceGroupName := "myrg"

	a := createResourceGroup(resourceGroupName)
	hierarchy := genruntime.ResourceHierarchy{a}

	// This is expected to fail
	_, err := hierarchy.ResourceGroup()
	g.Expect(err).To(HaveOccurred())

	location, err := hierarchy.Location()
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(a.Spec.Location).To(Equal(location))
	g.Expect(hierarchy.FullAzureName()).To(Equal(resourceGroupName))
}

func Test_ResourceHierarchy_ResourceGroup_TopLevelResource(t *testing.T) {
	g := NewGomegaWithT(t)

	resourceGroupName := "myrg"
	name := "myresource"

	a, b := createResourceGroupRootedResource(resourceGroupName, name)
	hierarchy := genruntime.ResourceHierarchy{a, b}

	rg, err := hierarchy.ResourceGroup()
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(rg).To(Equal(resourceGroupName))
	g.Expect(hierarchy.FullAzureName()).To(Equal(name))
}

func Test_ResourceHierarchy_ResourceGroup_NestedResource(t *testing.T) {
	g := NewGomegaWithT(t)

	resourceGroupName := "myrg"
	resourceName := "myresource"
	childResourceName := "mychildresource"

	hierarchy := createDeeplyNestedResource(resourceGroupName, resourceName, childResourceName)

	rg, err := hierarchy.ResourceGroup()
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(rg).To(Equal(resourceGroupName))
	g.Expect(hierarchy.FullAzureName()).To(Equal(fmt.Sprintf("%s/%s", hierarchy[1].AzureName(), hierarchy[2].AzureName())))
}

func Test_ResourceHierarchy_ScopeOnHierarchyWithoutExtensionResourceErrors(t *testing.T) {
	g := NewGomegaWithT(t)

	resourceGroupName := "myrg"
	name := "myresource"

	a, b := createResourceGroupRootedResource(resourceGroupName, name)
	hierarchy := genruntime.ResourceHierarchy{a, b}

	_, err := hierarchy.Scope()
	g.Expect(err).To(HaveOccurred())
	g.Expect(err.Error()).To(ContainSubstring("myresource is not of kind \"extension\""))
}

func Test_ResourceHierarchy_ExtensionOnResourceGroup(t *testing.T) {
	g := NewGomegaWithT(t)

	resourceGroupName := "myrg"
	extensionName := "myextension"

	a, b := createExtensionResourceOnResourceGroup(resourceGroupName, extensionName)
	hierarchy := genruntime.ResourceHierarchy{a, b}

	g.Expect(hierarchy.ResourceGroup()).To(Equal(resourceGroupName))

	g.Expect(hierarchy.Scope()).To(Equal(""))
	g.Expect(hierarchy.FullAzureName()).To(Equal(extensionName))
}

func Test_ResourceHierarchy_ExtensionOnResourceInResourceGroup(t *testing.T) {
	g := NewGomegaWithT(t)

	resourceGroupName := "myrg"
	resourceName := "myresource"
	extensionName := "myextension"

	hierarchy := createExtensionResourceOnResourceInResourceGroup(resourceGroupName, resourceName, extensionName)

	g.Expect(hierarchy.ResourceGroup()).To(Equal(resourceGroupName))
	g.Expect(hierarchy.Scope()).To(Equal(fmt.Sprintf("%s/%s", hierarchy[1].GetType(), resourceName)))
	g.Expect(hierarchy.FullAzureName()).To(Equal(extensionName))
}

func Test_ResourceHierarchy_ExtensionOnDeepHierarchy(t *testing.T) {
	g := NewGomegaWithT(t)

	resourceGroupName := "myrg"
	resourceName := "myresource"
	childResourceName := "mychildresource"
	extensionName := "myextension"

	hierarchy := createExtensionResourceOnDeepHierarchyInResourceGroup(resourceGroupName, resourceName, childResourceName, extensionName)

	g.Expect(hierarchy.ResourceGroup()).To(Equal(resourceGroupName))
	g.Expect(hierarchy.Scope()).To(Equal(fmt.Sprintf("%s/%s/%s/%s", hierarchy[1].GetType(), resourceName, "blobServices", hierarchy[2].AzureName())))
	g.Expect(hierarchy.FullAzureName()).To(Equal(extensionName))
}
