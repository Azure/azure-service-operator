/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package resolver_test

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/google/uuid"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	batch "github.com/Azure/azure-service-operator/v2/api/batch/v1api20210101"
	"github.com/Azure/azure-service-operator/v2/internal/resolver"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
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

	expectedARMID := fmt.Sprintf("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/%s", resourceGroupName)

	location, err := hierarchy.Location()
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(a.Spec.Location).To(Equal(to.Ptr(location)))
	g.Expect(hierarchy.AzureName()).To(Equal(resourceGroupName))
	g.Expect(hierarchy.FullyQualifiedARMID("00000000-0000-0000-0000-000000000000")).To(Equal(expectedARMID))
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
	g.Expect(hierarchy.FullyQualifiedARMID("00000000-0000-0000-0000-000000000000")).To(Equal(expectedARMID))
}

func Test_ResourceHierarchy_ResourceGroup_TopLevelResource(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	resourceGroupName := "myrg"
	name := "myresource"

	a, b := createResourceGroupRootedResource(resourceGroupName, name)
	hierarchy := resolver.ResourceHierarchy{a, b}

	expectedARMID := fmt.Sprintf("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/%s/providers/Microsoft.Batch/batchAccounts/%s", resourceGroupName, name)

	rg, err := hierarchy.ResourceGroup()
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(rg).To(Equal(resourceGroupName))
	g.Expect(hierarchy.AzureName()).To(Equal(name))
	g.Expect(hierarchy.FullyQualifiedARMID("00000000-0000-0000-0000-000000000000")).To(Equal(expectedARMID))
}

func Test_ResourceHierarchy_ResourceGroup_NestedResource(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	resourceGroupName := "myrg"
	resourceName := "myresource"
	childResourceName := "mychildresource"

	hierarchy := createDeeplyNestedResource(resourceGroupName, resourceName, childResourceName)

	expectedARMID := fmt.Sprintf(
		"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/%s/providers/Microsoft.Storage/storageAccounts/%s/blobServices/%s",
		resourceGroupName,
		resourceName,
		hierarchy[2].AzureName())

	rg, err := hierarchy.ResourceGroup()
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(rg).To(Equal(resourceGroupName))
	g.Expect(hierarchy.AzureName()).To(Equal(hierarchy[2].AzureName()))
	g.Expect(hierarchy.FullyQualifiedARMID("00000000-0000-0000-0000-000000000000")).To(Equal(expectedARMID))
}

func Test_ResourceHierarchy_ResourceGroup_NestedResource_MatchSubscriptionWithOwner(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	resourceGroupName := "myrg"
	resourceName := "myresource"
	childResourceName := "mychildresource"

	hierarchy := createDeeplyNestedResource(resourceGroupName, resourceName, childResourceName)

	uuid := uuid.New().String()
	_, err := hierarchy.FullyQualifiedARMID(uuid)
	g.Expect(err).To(MatchError(fmt.Sprintf("resource subscription \"%s\" does not match parent subscription \"00000000-0000-0000-0000-000000000000\"", uuid)))
}

func Test_ResourceHierarchy_ExtensionOnResourceGroup(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	resourceGroupName := "myrg"
	extensionName := "myextension"

	a, b := createExtensionResourceOnResourceGroup(resourceGroupName, extensionName)
	hierarchy := resolver.ResourceHierarchy{a, b}

	expectedARMID := fmt.Sprintf(
		"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/%s/providers/Microsoft.SimpleExtension/simpleExtensions/%s",
		resourceGroupName,
		extensionName)

	g.Expect(hierarchy.ResourceGroup()).To(Equal(resourceGroupName))
	g.Expect(hierarchy.FullyQualifiedARMID("00000000-0000-0000-0000-000000000000")).To(Equal(expectedARMID))
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

	g.Expect(hierarchy.FullyQualifiedARMID("00000000-0000-0000-0000-000000000000")).To(Equal(expectedARMID))
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
		"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/%s/providers/Microsoft.Batch/batchAccounts/%s/providers/Microsoft.SimpleExtension/simpleExtensions/%s",
		resourceGroupName,
		resourceName,
		extensionName)

	g.Expect(hierarchy.ResourceGroup()).To(Equal(resourceGroupName))
	g.Expect(hierarchy.FullyQualifiedARMID("00000000-0000-0000-0000-000000000000")).To(Equal(expectedARMID))
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
		"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/%s/providers/Microsoft.Storage/storageAccounts/%s/blobServices/%s/providers/Microsoft.SimpleExtension/simpleExtensions/%s",
		resourceGroupName,
		resourceName,
		hierarchy[2].AzureName(),
		extensionName)

	g.Expect(hierarchy.ResourceGroup()).To(Equal(resourceGroupName))
	g.Expect(hierarchy.FullyQualifiedARMID("00000000-0000-0000-0000-000000000000")).To(Equal(expectedARMID))
	g.Expect(hierarchy.AzureName()).To(Equal(extensionName))
}

func Test_ResourceHierarchy_OwnerARMIDResourceGroup(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	resourceGroupName := "myrg"
	resourceName := "myresource"

	rootARMID := fmt.Sprintf("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/%s", resourceGroupName)
	resource := createResourceGroupARMIDRootedResource(rootARMID, resourceName)
	hierarchy := resolver.ResourceHierarchy{resource}

	expectedARMID := fmt.Sprintf(
		"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/%s/providers/Microsoft.Batch/batchAccounts/%s",
		resourceGroupName,
		resourceName)

	g.Expect(hierarchy.ResourceGroup()).To(Equal(resourceGroupName))
	g.Expect(hierarchy.FullyQualifiedARMID("00000000-0000-0000-0000-000000000000")).To(Equal(expectedARMID))
	g.Expect(hierarchy.AzureName()).To(Equal(resourceName))
}

func Test_ResourceHierarchy_OwnerARMIDParent(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	resourceGroupName := "myrg"
	parentName := "parent"
	resourceName := "myresource"

	rootARMID := fmt.Sprintf("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/%s/providers/Microsoft.Storage/storageAccounts/%s", resourceGroupName, parentName)
	resource := createChildResourceOwnedByARMID(rootARMID, resourceName)
	hierarchy := resolver.ResourceHierarchy{resource}

	expectedARMID := fmt.Sprintf(
		"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/%s/providers/Microsoft.Storage/storageAccounts/%s/blobServices/default",
		resourceGroupName,
		parentName)

	g.Expect(hierarchy.ResourceGroup()).To(Equal(resourceGroupName))
	g.Expect(hierarchy.FullyQualifiedARMID("00000000-0000-0000-0000-000000000000")).To(Equal(expectedARMID))
	g.Expect(hierarchy.AzureName()).To(Equal("default"))
}

func Test_ResourceHierarchy_OwnerARMIDWrongProvider_ReturnsError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	resourceGroupName := "myrg"
	parentName := "parent"
	resourceName := "myresource"

	rootARMID := fmt.Sprintf("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/%s/providers/Microsoft.Qux/storageAccounts/%s", resourceGroupName, parentName)
	resource := createChildResourceOwnedByARMID(rootARMID, resourceName)
	hierarchy := resolver.ResourceHierarchy{resource}

	_, err := hierarchy.FullyQualifiedARMID("00000000-0000-0000-0000-000000000000")
	g.Expect(err).To(MatchError("expected owner ARM ID to be from provider \"Microsoft.Storage\", but was \"Microsoft.Qux\""))
}

func Test_ResourceHierarchy_OwnerARMIDWrongType_ReturnsError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	resourceGroupName := "myrg"
	parentName := "parent"
	resourceName := "myresource"

	rootARMID := fmt.Sprintf("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/%s/providers/Microsoft.Storage/cats/%s", resourceGroupName, parentName)
	resource := createChildResourceOwnedByARMID(rootARMID, resourceName)
	hierarchy := resolver.ResourceHierarchy{resource}

	_, err := hierarchy.FullyQualifiedARMID("00000000-0000-0000-0000-000000000000")
	g.Expect(err).To(MatchError("expected owner ARM ID to be of type \"Microsoft.Storage/storageAccounts\", but was \"Microsoft.Storage/cats\""))
}

func Test_ResourceHierarchy_OwnerARMIDAnotherWrongType_ReturnsError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	resourceGroupName := "myrg"
	parentName := "parent"
	resourceName := "myresource"

	rootARMID := fmt.Sprintf("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/%s/providers/Microsoft.Storage/cats/%s", resourceGroupName, parentName)
	resource := createResourceGroupARMIDRootedResource(rootARMID, resourceName)
	hierarchy := resolver.ResourceHierarchy{resource}

	_, err := hierarchy.FullyQualifiedARMID("00000000-0000-0000-0000-000000000000")
	g.Expect(err).To(MatchError("expected owner ARM ID to be for a resource group, but was \"Microsoft.Storage/cats\""))
}

func Test_ResourceHierarchy_OwnerARMIDDistinctSubscription_ReturnsSuccess(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	resourceGroupName := "myrg"
	resourceName := "myresource"

	uuid := uuid.New().String()
	rootARMID := fmt.Sprintf("/subscriptions/%s/resourceGroups/%s", uuid, resourceGroupName)
	resource := createResourceGroupARMIDRootedResource(rootARMID, resourceName)
	hierarchy := resolver.ResourceHierarchy{resource}

	_, err := hierarchy.FullyQualifiedARMID("00000000-0000-0000-0000-000000000000")
	g.Expect(err).To(BeNil())
}

func Test_ResourceHierarchy_OwnerARMIDWithExtension(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	resourceGroupName := "myrg"
	resourceName := "myresource"
	extensionName := "myextension"

	rootARMID := fmt.Sprintf("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/%s/providers/Microsoft.Storage/storageAccounts/%s", resourceGroupName, resourceName)
	resource := createSimpleExtensionResourceOwnedByARMID(extensionName, rootARMID)
	hierarchy := resolver.ResourceHierarchy{resource}

	expectedARMID := fmt.Sprintf(
		"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/%s/providers/Microsoft.Storage/storageAccounts/%s/providers/Microsoft.SimpleExtension/simpleExtensions/%s",
		resourceGroupName,
		resourceName,
		extensionName)

	g.Expect(hierarchy.ResourceGroup()).To(Equal(resourceGroupName))
	g.Expect(hierarchy.FullyQualifiedARMID("00000000-0000-0000-0000-000000000000")).To(Equal(expectedARMID))
	g.Expect(hierarchy.AzureName()).To(Equal(extensionName))
}

func Test_ResourceHierarchy_OwnerARMIDWithExtensionOnDeepHierarchy(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	resourceGroupName := "myrg"
	resourceName := "myresource"
	childResourceName := "mychildresource"
	extensionName := "myextension"

	rootARMID := fmt.Sprintf("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/%s/providers/Microsoft.Storage/storageAccounts/%s", resourceGroupName, resourceName)
	hierarchy := createExtensionResourceOnDeepHierarchyOwnedByARMID(rootARMID, childResourceName, extensionName)

	expectedARMID := fmt.Sprintf(
		"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/%s/providers/Microsoft.Storage/storageAccounts/%s/blobServices/%s/providers/Microsoft.SimpleExtension/simpleExtensions/%s",
		resourceGroupName,
		resourceName,
		hierarchy[0].AzureName(),
		extensionName)

	g.Expect(hierarchy.ResourceGroup()).To(Equal(resourceGroupName))
	g.Expect(hierarchy.FullyQualifiedARMID("00000000-0000-0000-0000-000000000000")).To(Equal(expectedARMID))
	g.Expect(hierarchy.AzureName()).To(Equal(extensionName))
}

func Test_ResourceHierarchy_ChildResourceIDOverride_DoesNotImpactResourceItself(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	resourceGroupName := "myrg"

	a := createResourceGroup(resourceGroupName)
	genruntime.SetChildResourceIDOverride(a, "/a/b/c")
	hierarchy := resolver.ResourceHierarchy{a}

	expectedARMID := fmt.Sprintf("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/%s", resourceGroupName)

	location, err := hierarchy.Location()
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(a.Spec.Location).To(Equal(to.Ptr(location)))
	g.Expect(hierarchy.AzureName()).To(Equal(resourceGroupName))
	g.Expect(hierarchy.FullyQualifiedARMID("00000000-0000-0000-0000-000000000000")).To(Equal(expectedARMID))
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

	g.Expect(hierarchy.FullyQualifiedARMID("00000000-0000-0000-0000-000000000000")).To(Equal(expectedARMID))
	g.Expect(hierarchy.AzureName()).To(Equal(extensionName))
}

func Test_ResourceHierarchy_GivenTopLevelResourceMissingAzureName_ReturnsError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	resourceGroupName := "myrg"
	a := createResourceGroup(resourceGroupName)

	name := "myresource"
	b := &batch.BatchAccount{
		TypeMeta: metav1.TypeMeta{
			Kind:       "BatchAccount",
			APIVersion: batch.GroupVersion.String(),
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: testNamespace,
		},
		Spec: batch.BatchAccount_Spec{
			Owner: &genruntime.KnownResourceReference{
				Name: resourceGroupName,
			},
		},
	}

	hierarchy := resolver.ResourceHierarchy{a, b}
	_, err := hierarchy.FullyQualifiedARMID("00000000-0000-0000-0000-000000000000")

	g.Expect(err).To(HaveOccurred())
	g.Expect(err.Error()).To(ContainSubstring("empty AzureName"))
}
