/*

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package keyvaults

import (
	"context"
	"fmt"
	"log"

	auth "github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	"github.com/Azure/azure-sdk-for-go/services/keyvault/mgmt/2018-02-14/keyvault"
	"github.com/Azure/azure-service-operator/api/v1alpha1"
	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/iam"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/go-logr/logr"
	uuid "github.com/satori/go.uuid"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
)

type azureKeyVaultManager struct {
	Log    logr.Logger
	Scheme *runtime.Scheme
}

func NewAzureKeyVaultManager(log logr.Logger, scheme *runtime.Scheme) *azureKeyVaultManager {
	return &azureKeyVaultManager{
		Log:    log,
		Scheme: scheme,
	}
}

func getVaultsClient() keyvault.VaultsClient {
	vaultsClient := keyvault.NewVaultsClient(config.SubscriptionID())
	a, err := iam.GetResourceManagementAuthorizer()
	if err != nil {
		log.Fatalf("failed to initialize authorizer: %v\n", err)
	}
	vaultsClient.Authorizer = a
	vaultsClient.AddToUserAgent(config.UserAgent())
	return vaultsClient
}

func getObjectID(ctx context.Context, tenantID string, clientID string) *string {
	appclient := auth.NewApplicationsClient(tenantID)
	a, err := iam.GetGraphAuthorizer()
	if err != nil {
		return nil
	}
	appclient.Authorizer = a
	appclient.AddToUserAgent(config.UserAgent())

	result, err := appclient.GetServicePrincipalsIDByAppID(ctx, clientID)
	if err != nil {
		return nil
	}
	return result.Value
}

// CreateVault creates a new key vault
func (k *azureKeyVaultManager) CreateVault(ctx context.Context, groupName string, vaultName string, location string, tags map[string]*string) (keyvault.Vault, error) {
	vaultsClient := getVaultsClient()
	id, err := uuid.FromString(config.TenantID())
	if err != nil {
		return keyvault.Vault{}, err
	}

	// Check if keyvault name is valid
	vaultNameCheck := keyvault.VaultCheckNameAvailabilityParameters{
		Name: to.StringPtr(vaultName),
		Type: to.StringPtr("Microsoft.KeyVault/vaults"),
	}
	result, err := vaultsClient.CheckNameAvailability(ctx, vaultNameCheck)
	if err != nil {
		k.Log.Info("CheckNameAvailability returned error")
		return keyvault.Vault{}, err
	}
	k.Log.Info("debug", "reason", result.Reason)
	if result.Reason == keyvault.Reason("Invalid") || result.Reason == keyvault.AccountNameInvalid {
		// Check for "Invalid" is to overcome a bug in KeyVault API which returns "Invalid" instead of defined "AccountNameInvalid"
		k.Log.Info("Name is invalid")
		return keyvault.Vault{}, fmt.Errorf("AccountNameInvalid")
	} else if result.Reason == keyvault.AlreadyExists {
		k.Log.Info("Keyvault with same name already exists elsewhere")
		return keyvault.Vault{}, fmt.Errorf("AlreadyExists")
	}

	params := keyvault.VaultCreateOrUpdateParameters{
		Properties: &keyvault.VaultProperties{
			TenantID:       &id,
			AccessPolicies: &[]keyvault.AccessPolicyEntry{},
			Sku: &keyvault.Sku{
				Family: to.StringPtr("A"),
				Name:   keyvault.Standard,
			},
		},
		Location: to.StringPtr(location),
		Tags:     tags,
	}

	k.Log.Info(fmt.Sprintf("creating keyvault '%s' in resource group '%s' and location: %v", vaultName, groupName, location))
	future, err := vaultsClient.CreateOrUpdate(ctx, groupName, vaultName, params)

	return future.Result(vaultsClient)
}

// CreateVaultWithAccessPolicies creates a new key vault and provides access policies to the specified user
func (k *azureKeyVaultManager) CreateVaultWithAccessPolicies(ctx context.Context, groupName string, vaultName string, location string, clientID string) (keyvault.Vault, error) {
	vaultsClient := getVaultsClient()
	id, err := uuid.FromString(config.TenantID())
	if err != nil {
		return keyvault.Vault{}, err
	}

	// Check if keyvault name is valid
	vaultNameCheck := keyvault.VaultCheckNameAvailabilityParameters{
		Name: to.StringPtr(vaultName),
		Type: to.StringPtr("Microsoft.KeyVault/vaults"),
	}
	result, err := vaultsClient.CheckNameAvailability(ctx, vaultNameCheck)
	if err != nil {
		k.Log.Info("CheckNameAvailability returned error")
		return keyvault.Vault{}, err
	}
	if result.Reason == keyvault.Reason("Invalid") || result.Reason == keyvault.AccountNameInvalid {
		k.Log.Info("Name is invalid")
		return keyvault.Vault{}, fmt.Errorf("AccountNameInvalid")
	} else if result.Reason == keyvault.AlreadyExists {
		k.Log.Info("Keyvault with same name already exists elsewhere")
		return keyvault.Vault{}, fmt.Errorf("AlreadyExists")
	}

	apList := []keyvault.AccessPolicyEntry{}
	ap := keyvault.AccessPolicyEntry{
		TenantID: &id,
		Permissions: &keyvault.Permissions{
			Keys: &[]keyvault.KeyPermissions{
				keyvault.KeyPermissionsCreate,
			},
			Secrets: &[]keyvault.SecretPermissions{
				keyvault.SecretPermissionsSet,
				keyvault.SecretPermissionsGet,
				keyvault.SecretPermissionsDelete,
				keyvault.SecretPermissionsList,
			},
		},
	}
	if clientID != "" {
		if objID := getObjectID(ctx, config.TenantID(), clientID); objID != nil {
			ap.ObjectID = objID
			apList = append(apList, ap)
		}

	}

	params := keyvault.VaultCreateOrUpdateParameters{
		Properties: &keyvault.VaultProperties{
			TenantID:       &id,
			AccessPolicies: &apList,
			Sku: &keyvault.Sku{
				Family: to.StringPtr("A"),
				Name:   keyvault.Standard,
			},
		},
		Location: to.StringPtr(location),
	}

	k.Log.Info(fmt.Sprintf("creating keyvault '%s' in resource group '%s' and location: %v with access policies granted to %v", vaultName, groupName, location, clientID))
	future, err := vaultsClient.CreateOrUpdate(ctx, groupName, vaultName, params)
	if err != nil {
		return keyvault.Vault{}, err
	}

	return future.Result(vaultsClient)
}

// DeleteVault removes the resource group named by env var
func (k *azureKeyVaultManager) DeleteVault(ctx context.Context, groupName string, vaultName string) (result autorest.Response, err error) {
	vaultsClient := getVaultsClient()
	return vaultsClient.Delete(ctx, groupName, vaultName)
}

// CheckExistence checks for the presence of a keyvault instance on Azure
func (k *azureKeyVaultManager) GetVault(ctx context.Context, groupName string, vaultName string) (result keyvault.Vault, err error) {
	vaultsClient := getVaultsClient()
	return vaultsClient.Get(ctx, groupName, vaultName)

}

func (k *azureKeyVaultManager) Ensure(ctx context.Context, obj runtime.Object) (bool, error) {
	instance, err := k.convert(obj)
	if err != nil {
		return true, err
	}

	// convert kube labels to expected tag format
	labels := map[string]*string{}
	for k, v := range instance.GetLabels() {
		labels[k] = &v
	}
	instance.Status.Provisioning = true

	// Check if this KeyVault already exists and its state if it does.

	keyvault, err := k.GetVault(ctx, instance.Spec.ResourceGroup, instance.Name)
	if err == nil {
		instance.Status.Message = "Successfully Provisioned"
		instance.Status.Provisioned = true
		instance.Status.Provisioning = false
		return true, nil
	}
	k.Log.Info("KeyVault not present, creating")

	keyvault, err = k.CreateVault(
		ctx,
		instance.Spec.ResourceGroup,
		instance.Name,
		instance.Spec.Location,
		labels,
	)

	if err != nil {
		// let the user know what happened
		instance.Status.Message = err.Error()
		instance.Status.Provisioning = false
		// errors we expect might happen that we are ok with waiting for
		catch := []string{
			errhelp.ResourceGroupNotFoundErrorCode,
			errhelp.ParentNotFoundErrorCode,
			errhelp.NotFoundErrorCode,
			errhelp.AsyncOpIncompleteError,
		}

		catchUnrecoverableErrors := []string{
			errhelp.AccountNameInvalid,
			errhelp.AlreadyExists,
		}

		azerr := errhelp.NewAzureErrorAzureError(err)
		if helpers.ContainsString(catch, azerr.Type) {
			// most of these error technically mean the resource is actually not provisioning
			switch azerr.Type {
			case errhelp.AsyncOpIncompleteError:
				instance.Status.Provisioning = true
			}
			// reconciliation is not done but error is acceptable
			return false, nil
		}
		if helpers.ContainsString(catchUnrecoverableErrors, azerr.Type) {
			// Unrecoverable error, so stop reconcilation
			instance.Status.Provisioning = false
			instance.Status.Message = "Reconcilation hit unrecoverable error"
			k.Log.Info("Reconcilation hit unrecoverable error", "error=", err.Error())
			return true, nil
		}
		// reconciliation not done and we don't know what happened
		return false, err

	}

	instance.Status.State = keyvault.Status

	if instance.Status.Provisioning {
		instance.Status.Provisioned = true
		instance.Status.Provisioning = false
		instance.Status.Message = "Suceeded"
	} else {
		instance.Status.Provisioned = false
		instance.Status.Provisioning = true
	}

	return true, nil
}

func (k *azureKeyVaultManager) Delete(ctx context.Context, obj runtime.Object) (bool, error) {
	instance, err := k.convert(obj)
	if err != nil {
		return true, err
	}

	_, err = k.GetVault(ctx, instance.Spec.ResourceGroup, instance.Name)
	if err == nil {
		k.Log.Info("KeyVault present, deleting")
		_, err := k.DeleteVault(ctx, instance.Spec.ResourceGroup, instance.Name)
		if err != nil {
			k.Log.Info("Delete:", "KeyVault Delete returned=", err.Error())
			if !errhelp.IsAsynchronousOperationNotComplete(err) {
				k.Log.Info("Error from delete call")
				return true, err
			}
		}
		return true, nil
	}

	instance.Status.State = "Deleted"

	return false, nil
}

func (k *azureKeyVaultManager) GetParents(obj runtime.Object) ([]resourcemanager.KubeParent, error) {

	instance, err := k.convert(obj)
	if err != nil {
		return nil, err
	}

	return []resourcemanager.KubeParent{
		{
			Key: types.NamespacedName{
				Namespace: instance.Namespace,
				Name:      instance.Spec.ResourceGroup,
			},
			Target: &azurev1alpha1.ResourceGroup{},
		},
	}, nil

}

func (k *azureKeyVaultManager) convert(obj runtime.Object) (*v1alpha1.KeyVault, error) {
	local, ok := obj.(*v1alpha1.KeyVault)
	if !ok {
		return nil, fmt.Errorf("failed type assertion on kind: %s", obj.GetObjectKind().GroupVersionKind().String())
	}
	return local, nil
}
