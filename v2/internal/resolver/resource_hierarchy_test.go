/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package resolver_test

import (
	"fmt"
	"testing"

	"github.com/Azure/go-autorest/autorest/to"
	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/v2/internal/resolver"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func Test_ResourceHierarchy_ResourceGroupOnly(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	resourceGroupName := "myrg"

	a := createResourceGroup(resourceGroupName)
	hierarchy := resolver.ResourceHierarchy{a}

	// This is expected to fail
	_, err := hierarchy.ResourceGroup()
	g.Expect(err).To(HaveOccurred())

	expectedARMID := fmt.Sprintf("/subscriptions/1234/resourceGroups/%s", resourceGroupName)

	location, err := hierarchy.Location()
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(a.Spec.Location).To(Equal(to.StringPtr(location)))
	g.Expect(hierarchy.AzureName()).To(Equal(resourceGroupName))
	g.Expect(hierarchy.FullyQualifiedARMID("1234")).To(Equal(expectedARMID))
}

func Test_ResourceHierarchy_TenantScopeResourceOnly(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	subscriptionName := "mysub"

	a := createSubscription(subscriptionName)
	hierarchy := resolver.ResourceHierarchy{a}

	// This is expected to fail
	_, err := hierarchy.ResourceGroup()
	g.Expect(err).To(HaveOccurred())
	_, err = hierarchy.Location()
	g.Expect(err).To(HaveOccurred())

	expectedARMID := fmt.Sprintf("/providers/Microsoft.Subscription/aliases/%s", subscriptionName)

	g.Expect(hierarchy.AzureName()).To(Equal(subscriptionName))
	g.Expect(hierarchy.FullyQualifiedARMID("1234")).To(Equal(expectedARMID))
}

func Test_ResourceHierarchy_ResourceGroup_TopLevelResource(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	resourceGroupName := "myrg"
	name := "myresource"

	a, b := createResourceGroupRootedResource(resourceGroupName, name)
	hierarchy := resolver.ResourceHierarchy{a, b}

	expectedARMID := fmt.Sprintf("/subscriptions/1234/resourceGroups/%s/providers/Microsoft.Batch/batchAccounts/%s", resourceGroupName, name)

	rg, err := hierarchy.ResourceGroup()
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(rg).To(Equal(resourceGroupName))
	g.Expect(hierarchy.AzureName()).To(Equal(name))
	g.Expect(hierarchy.FullyQualifiedARMID("1234")).To(Equal(expectedARMID))
}

func Test_ResourceHierarchy_ResourceGroup_NestedResource(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	resourceGroupName := "myrg"
	resourceName := "myresource"
	childResourceName := "mychildresource"

	hierarchy := createDeeplyNestedResource(resourceGroupName, resourceName, childResourceName)

	expectedARMID := fmt.Sprintf(
		"/subscriptions/1234/resourceGroups/%s/providers/Microsoft.Storage/storageAccounts/%s/blobServices/%s",
		resourceGroupName,
		resourceName,
		hierarchy[2].AzureName())

	rg, err := hierarchy.ResourceGroup()
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(rg).To(Equal(resourceGroupName))
	g.Expect(hierarchy.AzureName()).To(Equal(hierarchy[2].AzureName()))
	g.Expect(hierarchy.FullyQualifiedARMID("1234")).To(Equal(expectedARMID))
}

func Test_ResourceHierarchy_ExtensionOnResourceGroup(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	resourceGroupName := "myrg"
	extensionName := "myextension"

	a, b := createExtensionResourceOnResourceGroup(resourceGroupName, extensionName)
	hierarchy := resolver.ResourceHierarchy{a, b}

	expectedARMID := fmt.Sprintf(
		"/subscriptions/1234/resourceGroups/%s/providers/Microsoft.SimpleExtension/simpleExtensions/%s",
		resourceGroupName,
		extensionName)

	g.Expect(hierarchy.ResourceGroup()).To(Equal(resourceGroupName))
	g.Expect(hierarchy.FullyQualifiedARMID("1234")).To(Equal(expectedARMID))
	g.Expect(hierarchy.AzureName()).To(Equal(extensionName))
}

func Test_ResourceHierarchy_ExtensionOnTenantScopeResource(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	subscriptionName := "mysub"
	extensionName := "myextension"

	hierarchy := createExtensionResourceOnTenantScopeResource(subscriptionName, extensionName)

	// This is expected to fail
	_, err := hierarchy.ResourceGroup()
	g.Expect(err).To(HaveOccurred())
	_, err = hierarchy.Location()
	g.Expect(err).To(HaveOccurred())

	// TODO: Confirm that this is actually the right URL
	expectedARMID := fmt.Sprintf(
		"/providers/Microsoft.Subscription/aliases/%s/providers/Microsoft.SimpleExtension/simpleExtensions/%s",
		subscriptionName,
		extensionName)

	g.Expect(hierarchy.FullyQualifiedARMID("1234")).To(Equal(expectedARMID))
	g.Expect(hierarchy.AzureName()).To(Equal(extensionName))
}

func Test_ResourceHierarchy_ExtensionOnResourceInResourceGroup(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	resourceGroupName := "myrg"
	resourceName := "myresource"
	extensionName := "myextension"

	hierarchy := createExtensionResourceOnResourceInResourceGroup(resourceGroupName, resourceName, extensionName)

	expectedARMID := fmt.Sprintf(
		"/subscriptions/1234/resourceGroups/%s/providers/Microsoft.Batch/batchAccounts/%s/providers/Microsoft.SimpleExtension/simpleExtensions/%s",
		resourceGroupName,
		resourceName,
		extensionName)

	g.Expect(hierarchy.ResourceGroup()).To(Equal(resourceGroupName))
	g.Expect(hierarchy.FullyQualifiedARMID("1234")).To(Equal(expectedARMID))
	g.Expect(hierarchy.AzureName()).To(Equal(extensionName))
}

func Test_ResourceHierarchy_ExtensionOnDeepHierarchy(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	resourceGroupName := "myrg"
	resourceName := "myresource"
	childResourceName := "mychildresource"
	extensionName := "myextension"

	hierarchy := createExtensionResourceOnDeepHierarchyInResourceGroup(resourceGroupName, resourceName, childResourceName, extensionName)

	expectedARMID := fmt.Sprintf(
		"/subscriptions/1234/resourceGroups/%s/providers/Microsoft.Storage/storageAccounts/%s/blobServices/%s/providers/Microsoft.SimpleExtension/simpleExtensions/%s",
		resourceGroupName,
		resourceName,
		hierarchy[2].AzureName(),
		extensionName)

	g.Expect(hierarchy.ResourceGroup()).To(Equal(resourceGroupName))
	g.Expect(hierarchy.FullyQualifiedARMID("1234")).To(Equal(expectedARMID))
	g.Expect(hierarchy.AzureName()).To(Equal(extensionName))
}

func Test_ResourceHierarchy_ChildResourceIDOverride_DoesNotImpactResourceItself(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	resourceGroupName := "myrg"

	a := createResourceGroup(resourceGroupName)
	genruntime.SetChildResourceIDOverride(a, "/a/b/c")
	hierarchy := resolver.ResourceHierarchy{a}

	expectedARMID := fmt.Sprintf("/subscriptions/1234/resourceGroups/%s", resourceGroupName)

	location, err := hierarchy.Location()
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(a.Spec.Location).To(Equal(to.StringPtr(location)))
	g.Expect(hierarchy.AzureName()).To(Equal(resourceGroupName))
	g.Expect(hierarchy.FullyQualifiedARMID("1234")).To(Equal(expectedARMID))
}

func Test_ResourceHierarchy_ChildResourceIDOverride_ImpactsExtensionResource(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	resourceGroupName := "myrg"
	extensionName := "myextension"

	a, b := createExtensionResourceOnResourceGroup(resourceGroupName, extensionName)
	genruntime.SetChildResourceIDOverride(a, "/a/b/c")
	hierarchy := resolver.ResourceHierarchy{a, b}

	// This is expected to fail
	_, err := hierarchy.ResourceGroup()
	g.Expect(err).To(HaveOccurred())
	_, err = hierarchy.Location()
	g.Expect(err).To(HaveOccurred())

	expectedARMID := fmt.Sprintf(
		"/a/b/c/providers/Microsoft.SimpleExtension/simpleExtensions/%s",
		extensionName)

	g.Expect(hierarchy.FullyQualifiedARMID("1234")).To(Equal(expectedARMID))
	g.Expect(hierarchy.AzureName()).To(Equal(extensionName))
}
