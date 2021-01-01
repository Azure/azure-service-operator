// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

// +build all azureloadbalancer

package controllers

import azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"

func TestAzureContainerRegistry(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()

	rgName := tc.resourceGroupName
	rgLocation := tc.resourceGroupLocation
	azureContainerRegistryName := GenerateAlphaNumTestResourceName("azurecontainerregistry")

	azureContainerRegistryInstance := &azurev1alpha1.AzureContainerRegistry{
		ObjectMeta: metav1.ObjectMeta{
			Name:      azureContainerRegistryName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.AzureContainerRegistrySpec{
			AdminUserEnabled: false,
			Location:         tc.resourceGroupLocation,
			ResourceGroup:    tc.resourceGroupName,
			Sku: azurev1alpha1.ContainerRegistrySku{
				Name: "Standard",
			},
		},
	}

	// Create the Keyvault object and expect the Reconcile to be created
	EnsureInstance(ctx, t, tc, azureContainerRegistryInstance)

	// verify key vault exists in Azure
	assert.Eventually(func() bool {
		result, _ := tc.containerRegistryManager.GetRegistry(ctx, tc.resourceGroupName, azureContainerRegistryInstance.Name)
		return result.Response.StatusCode == http.StatusOK
	}, tc.timeout, tc.retry, "wait for containerRegistryInstance to be ready in azure")

	EnsureDelete(ctx, t, tc, azureContainerRegistryInstance)

	assert.Eventually(func() bool {

		result, _ := tc.containerRegistryManager.GetRegistry(ctx, tc.resourceGroupName, azureContainerRegistryInstance.Name)
		return result.Response.StatusCode == http.StatusNotFound
	}, tc.timeout, tc.retry, "wait for containerRegistryInstance to be gone from azure")

}
