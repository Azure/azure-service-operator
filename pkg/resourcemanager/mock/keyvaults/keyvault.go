// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package keyvaults

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"

	"github.com/Azure/azure-sdk-for-go/services/keyvault/mgmt/2018-02-14/keyvault"
	"github.com/Azure/azure-service-operator/api/v1alpha1"
	pkghelpers "github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/mock/helpers"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/to"
	"k8s.io/apimachinery/pkg/runtime"
)

type keyVaultResource struct {
	resourceGroupName string
	vaultName         string
	KeyVault          keyvault.Vault
}

type MockKeyVaultManager struct {
	keyVaultResources []keyVaultResource
}

func findKeyVault(res []keyVaultResource, predicate func(keyVaultResource) bool) (int, keyVaultResource) {
	for index, r := range res {
		if predicate(r) {
			return index, r
		}
	}
	return -1, keyVaultResource{}
}

// CreateVault creates a new key vault
func (manager *MockKeyVaultManager) CreateVault(ctx context.Context, instance *v1alpha1.KeyVault, tags map[string]*string) (keyvault.Vault, error) {
	vaultName := instance.Name
	groupName := instance.Spec.ResourceGroup
	location := instance.Spec.Location

	v := keyvault.Vault{
		Response:   helpers.GetRestResponse(http.StatusOK),
		Properties: &keyvault.VaultProperties{},
		ID:         to.StringPtr(pkghelpers.RandomString(10)),
		Name:       to.StringPtr(vaultName),
		Location:   to.StringPtr(location),
	}

	_, err := manager.GetVault(ctx, groupName, vaultName)
	if err != nil {
		manager.keyVaultResources = append(manager.keyVaultResources, keyVaultResource{
			resourceGroupName: groupName,
			vaultName:         vaultName,
			KeyVault:          v,
		})
	}

	return v, nil
}

// CreateVaultWithAccessPolicies creates a new key vault
func (manager *MockKeyVaultManager) CreateVaultWithAccessPolicies(ctx context.Context, groupName string, vaultName string, location string, clientID string, sku azurev1alpha1.KeyVaultSku) (keyvault.Vault, error) {
	v := keyvault.Vault{
		Response:   helpers.GetRestResponse(http.StatusOK),
		Properties: &keyvault.VaultProperties{},
		ID:         to.StringPtr(pkghelpers.RandomString(10)),
		Name:       to.StringPtr(vaultName),
		Location:   to.StringPtr(location),
	}

	_, err := manager.GetVault(ctx, groupName, vaultName)
	if err != nil {
		manager.keyVaultResources = append(manager.keyVaultResources, keyVaultResource{
			resourceGroupName: groupName,
			vaultName:         vaultName,
			KeyVault:          v,
		})
	}

	return v, nil
}

// DeleteVault removes the resource group named by env var
func (manager *MockKeyVaultManager) DeleteVault(ctx context.Context, groupName string, vaultName string) (result autorest.Response, err error) {
	vaults := manager.keyVaultResources

	index, _ := findKeyVault(vaults, func(g keyVaultResource) bool {
		return g.resourceGroupName == groupName &&
			g.vaultName == vaultName
	})

	if index == -1 {
		return helpers.GetRestResponse(http.StatusNotFound), errors.New("key vault not found")
	}

	manager.keyVaultResources = append(vaults[:index], vaults[index+1:]...)

	return helpers.GetRestResponse(http.StatusOK), nil
}

// Returns an existing keyvault instance
func (manager *MockKeyVaultManager) GetVault(ctx context.Context, groupName string, vaultName string) (result keyvault.Vault, err error) {
	vaults := manager.keyVaultResources

	index, v := findKeyVault(vaults, func(g keyVaultResource) bool {
		return g.resourceGroupName == groupName &&
			g.vaultName == vaultName
	})

	if index == -1 {
		return keyvault.Vault{
			Response: helpers.GetRestResponse(http.StatusNotFound),
		}, errors.New("key vault not found")
	}

	return v.KeyVault, nil
}

func (manager *MockKeyVaultManager) convert(obj runtime.Object) (*v1alpha1.KeyVault, error) {
	local, ok := obj.(*v1alpha1.KeyVault)
	if !ok {
		return nil, fmt.Errorf("failed type assertion on kind: %s", obj.GetObjectKind().GroupVersionKind().String())
	}
	return local, nil
}

func (manager *MockKeyVaultManager) Ensure(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {

	instance, err := manager.convert(obj)
	if err != nil {
		return true, err
	}

	tags := map[string]*string{}
	_, _ = manager.CreateVault(
		ctx,
		instance,
		tags,
	)

	instance.Status.Provisioned = true

	return true, nil
}
func (manager *MockKeyVaultManager) Delete(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {

	instance, err := manager.convert(obj)
	if err != nil {
		return true, err
	}

	_, _ = manager.DeleteVault(ctx, instance.Spec.ResourceGroup, instance.Name)

	return false, nil
}
func (manager *MockKeyVaultManager) GetParents(runtime.Object) ([]resourcemanager.KubeParent, error) {
	return []resourcemanager.KubeParent{}, nil
}
