// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package keyvaults

import (
	"context"
	"fmt"
	"strconv"

	auth "github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	"github.com/Azure/azure-sdk-for-go/services/keyvault/mgmt/2018-02-14/keyvault"
	kvops "github.com/Azure/azure-sdk-for-go/services/keyvault/v7.0/keyvault"
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
	"github.com/mitchellh/hashstructure"
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

func getVaultsClient() (keyvault.VaultsClient, error) {
	vaultsClient := keyvault.NewVaultsClientWithBaseURI(config.BaseURI(), config.SubscriptionID())
	a, err := iam.GetResourceManagementAuthorizer()
	if err != nil {
		return vaultsClient, err
	}
	vaultsClient.Authorizer = a
	vaultsClient.AddToUserAgent(config.UserAgent())
	return vaultsClient, nil
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

func parseNetworkPolicy(instance *v1alpha1.KeyVault) keyvault.NetworkRuleSet {
	var bypass keyvault.NetworkRuleBypassOptions
	switch instance.Spec.NetworkPolicies.Bypass {
	case "AzureServices":
		bypass = keyvault.AzureServices
	case "None":
		bypass = keyvault.None
	default:
		bypass = keyvault.AzureServices
	}

	var defaultAction keyvault.NetworkRuleAction
	switch instance.Spec.NetworkPolicies.DefaultAction {
	case "Allow":
		defaultAction = keyvault.Allow
	case "Deny":
		defaultAction = keyvault.Deny
	default:
		defaultAction = keyvault.Deny
	}

	var ipInstances []keyvault.IPRule
	for _, ip := range *instance.Spec.NetworkPolicies.IPRules {
		ipInstances = append(ipInstances, keyvault.IPRule{Value: &ip})
	}

	var virtualNetworkRules []keyvault.VirtualNetworkRule
	for _, id := range *instance.Spec.NetworkPolicies.VirtualNetworkRules {
		virtualNetworkRules = append(virtualNetworkRules, keyvault.VirtualNetworkRule{ID: &id})
	}

	networkAcls := keyvault.NetworkRuleSet{
		Bypass:              bypass,
		DefaultAction:       defaultAction,
		IPRules:             &ipInstances,
		VirtualNetworkRules: &virtualNetworkRules,
	}

	return networkAcls
}

// GenerateSpecHash - helper function that generates a unique hash for a Kubernetes runtime.Object resource
func GenerateSpecHash(obj runtime.Object) (string, error) {
	unstructured, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return "", err
	}

	hash, err := hashstructure.Hash(unstructured["spec"].(map[string]interface{}), nil)
	if err != nil {
		return "", err
	}

	if hash == 0 {
		return "", fmt.Errorf("InvalidHash: Hashing function returned 0")
	}

	return strconv.FormatUint(hash, 10), nil
}

func ParseAccessPolicy(policy *v1alpha1.AccessPolicyEntry, ctx context.Context) (keyvault.AccessPolicyEntry, error) {
	tenantID, err := uuid.FromString(policy.TenantID)
	if err != nil {
		return keyvault.AccessPolicyEntry{}, err
	}

	newEntry := keyvault.AccessPolicyEntry{
		TenantID:    &tenantID,
		Permissions: &keyvault.Permissions{},
	}

	if policy.Permissions.Keys != nil {
		var keyPermissions []keyvault.KeyPermissions
		permissions := keyvault.PossibleKeyPermissionsValues()
		validKeyPermissions := []string{}
		for _, item := range permissions {
			validKeyPermissions = append(validKeyPermissions, string(item))
		}

		for _, key := range *policy.Permissions.Keys {
			if helpers.ContainsString(validKeyPermissions, key) {
				keyPermissions = append(keyPermissions, keyvault.KeyPermissions(key))
			} else {
				return keyvault.AccessPolicyEntry{}, fmt.Errorf("InvalidAccessPolicy: Invalid Key Permission")
			}
		}

		newEntry.Permissions.Keys = &keyPermissions
	}

	if policy.Permissions.Secrets != nil {
		var secretPermissions []keyvault.SecretPermissions
		permissions := keyvault.PossibleSecretPermissionsValues()
		validSecretPermissions := []string{}
		for _, item := range permissions {
			validSecretPermissions = append(validSecretPermissions, string(item))
		}

		for _, key := range *policy.Permissions.Secrets {
			if helpers.ContainsString(validSecretPermissions, key) {
				secretPermissions = append(secretPermissions, keyvault.SecretPermissions(key))
			} else {
				return keyvault.AccessPolicyEntry{}, fmt.Errorf("InvalidAccessPolicy: Invalid Secret Permission")
			}
		}

		newEntry.Permissions.Secrets = &secretPermissions
	}

	if policy.Permissions.Certificates != nil {
		var certificatePermissions []keyvault.CertificatePermissions
		permissions := keyvault.PossibleCertificatePermissionsValues()
		validCertificatePermissions := []string{}
		for _, item := range permissions {
			validCertificatePermissions = append(validCertificatePermissions, string(item))
		}

		for _, key := range *policy.Permissions.Certificates {
			if helpers.ContainsString(validCertificatePermissions, key) {
				certificatePermissions = append(certificatePermissions, keyvault.CertificatePermissions(key))
			} else {
				return keyvault.AccessPolicyEntry{}, fmt.Errorf("InvalidAccessPolicy: Invalid Certificate Permission")
			}
		}

		newEntry.Permissions.Certificates = &certificatePermissions
	}

	if policy.Permissions.Storage != nil {
		var storagePermissions []keyvault.StoragePermissions
		permissions := keyvault.PossibleStoragePermissionsValues()
		validStoragePermissions := []string{}
		for _, item := range permissions {
			validStoragePermissions = append(validStoragePermissions, string(item))
		}

		for _, key := range *policy.Permissions.Storage {
			if helpers.ContainsString(validStoragePermissions, key) {
				storagePermissions = append(storagePermissions, keyvault.StoragePermissions(key))
			} else {
				return keyvault.AccessPolicyEntry{}, fmt.Errorf("InvalidAccessPolicy: Invalid Storage Permission")
			}
		}

		newEntry.Permissions.Storage = &storagePermissions
	}

	if policy.ApplicationID != "" {
		appID, err := uuid.FromString(policy.ApplicationID)
		if err != nil {
			return keyvault.AccessPolicyEntry{}, err
		}

		newEntry.ApplicationID = &appID
	}

	if policy.ClientID != "" {
		if objID := getObjectID(ctx, policy.TenantID, policy.ClientID); objID != nil {
			newEntry.ObjectID = objID
		}
	}

	return newEntry, nil
}

// InstantiateVault will instantiate VaultsClient
func InstantiateVault(ctx context.Context, vaultName string, containsUpdate bool) (keyvault.VaultsClient, uuid.UUID, error) {
	vaultsClient, err := getVaultsClient()
	if err != nil {
		return keyvault.VaultsClient{}, uuid.UUID{}, err
	}
	id, err := uuid.FromString(config.TenantID())
	if err != nil {
		return keyvault.VaultsClient{}, uuid.UUID{}, err
	}

	// Check if keyvault name is valid
	if !containsUpdate {
		vaultNameCheck := keyvault.VaultCheckNameAvailabilityParameters{
			Name: to.StringPtr(vaultName),
			Type: to.StringPtr("Microsoft.KeyVault/vaults"),
		}
		result, err := vaultsClient.CheckNameAvailability(ctx, vaultNameCheck)
		if err != nil {
			return keyvault.VaultsClient{}, uuid.UUID{}, err
		}
		if result.Reason == keyvault.Reason("Invalid") || result.Reason == keyvault.AccountNameInvalid {
			return keyvault.VaultsClient{}, uuid.UUID{}, fmt.Errorf("AccountNameInvalid")
		} else if result.Reason == keyvault.AlreadyExists {
			return keyvault.VaultsClient{}, uuid.UUID{}, fmt.Errorf("AlreadyExists")
		}
	}

	return vaultsClient, id, nil
}

// CreateVault creates a new key vault
func (k *azureKeyVaultManager) CreateVault(ctx context.Context, instance *v1alpha1.KeyVault, tags map[string]*string) (keyvault.Vault, error) {
	vaultName := instance.Name
	location := instance.Spec.Location
	groupName := instance.Spec.ResourceGroup
	enableSoftDelete := instance.Spec.EnableSoftDelete

	vaultsClient, id, err := InstantiateVault(ctx, vaultName, instance.Status.ContainsUpdate)
	if err != nil {
		return keyvault.Vault{}, err
	}

	var accessPolicies []keyvault.AccessPolicyEntry
	if instance.Spec.AccessPolicies != nil {
		for _, policy := range *instance.Spec.AccessPolicies {
			newEntry, err := ParseAccessPolicy(&policy, ctx)
			if err != nil {
				return keyvault.Vault{}, err
			}
			accessPolicies = append(accessPolicies, newEntry)
		}
	} else {
		accessPolicies = []keyvault.AccessPolicyEntry{}
	}

	var networkAcls keyvault.NetworkRuleSet
	if instance.Spec.NetworkPolicies != nil {
		networkAcls = parseNetworkPolicy(instance)
	} else {
		networkAcls = keyvault.NetworkRuleSet{}
	}

	params := keyvault.VaultCreateOrUpdateParameters{
		Properties: &keyvault.VaultProperties{
			TenantID:       &id,
			AccessPolicies: &accessPolicies,
			Sku: &keyvault.Sku{
				Family: to.StringPtr("A"),
				Name:   keyvault.Standard,
			},
			NetworkAcls:      &networkAcls,
			EnableSoftDelete: &enableSoftDelete,
		},
		Location: to.StringPtr(location),
		Tags:     tags,
	}

	future, err := vaultsClient.CreateOrUpdate(ctx, groupName, vaultName, params)

	return future.Result(vaultsClient)
}

// CreateVaultWithAccessPolicies creates a new key vault and provides access policies to the specified user
func (k *azureKeyVaultManager) CreateVaultWithAccessPolicies(ctx context.Context, groupName string, vaultName string, location string, clientID string) (keyvault.Vault, error) {
	vaultsClient, id, err := InstantiateVault(ctx, vaultName, false)
	if err != nil {
		return keyvault.Vault{}, err
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

	future, err := vaultsClient.CreateOrUpdate(ctx, groupName, vaultName, params)
	if err != nil {
		return keyvault.Vault{}, err
	}

	return future.Result(vaultsClient)
}

// DeleteVault removes the resource group named by env var
func (k *azureKeyVaultManager) DeleteVault(ctx context.Context, groupName string, vaultName string) (result autorest.Response, err error) {
	vaultsClient, err := getVaultsClient()
	if err != nil {
		return autorest.Response{}, err
	}
	return vaultsClient.Delete(ctx, groupName, vaultName)
}

// CheckExistence checks for the presence of a keyvault instance on Azure
func (k *azureKeyVaultManager) GetVault(ctx context.Context, groupName string, vaultName string) (result keyvault.Vault, err error) {
	vaultsClient, err := getVaultsClient()
	if err != nil {
		return keyvault.Vault{}, err
	}
	return vaultsClient.Get(ctx, groupName, vaultName)

}

func (k *azureKeyVaultManager) Ensure(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {
	instance, err := k.convert(obj)
	if err != nil {
		return true, err
	}

	hash, err := GenerateSpecHash(obj)
	if err != nil {
		return true, err
	}

	if instance.Status.SpecHash == "" {
		instance.Status.SpecHash = hash
	}

	// convert kube labels to expected tag format
	labels := map[string]*string{}
	for k, v := range instance.GetLabels() {
		value := v
		labels[k] = &value
	}
	instance.Status.Provisioning = true

	// Check if this KeyVault already exists and its state if it does.

	keyvault, err := k.GetVault(ctx, instance.Spec.ResourceGroup, instance.Name)
	if err == nil {
		if instance.Status.SpecHash == hash {
			instance.Status.Message = resourcemanager.SuccessMsg
			instance.Status.Provisioned = true
			instance.Status.Provisioning = false
			instance.Status.ResourceId = *keyvault.ID
			return true, nil
		}

		instance.Status.SpecHash = hash
		instance.Status.ContainsUpdate = true
	}

	keyvault, err = k.CreateVault(
		ctx,
		instance,
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
			errhelp.InvalidAccessPolicy,
			errhelp.BadRequest,
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
			instance.Status.Message = "Reconcilation hit unrecoverable error " + err.Error()
			return true, nil
		}
		// reconciliation not done and we don't know what happened
		return false, err

	}
	instance.Status.ContainsUpdate = false
	instance.Status.State = keyvault.Status

	instance.Status.Provisioned = true
	instance.Status.Provisioning = false
	instance.Status.Message = resourcemanager.SuccessMsg
	instance.Status.ResourceId = *keyvault.ID

	return true, nil
}

func (k *azureKeyVaultManager) Delete(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {
	instance, err := k.convert(obj)
	if err != nil {
		return true, err
	}

	_, err = k.GetVault(ctx, instance.Spec.ResourceGroup, instance.Name)
	if err == nil {
		_, err := k.DeleteVault(ctx, instance.Spec.ResourceGroup, instance.Name)
		if err != nil {
			if !errhelp.IsAsynchronousOperationNotComplete(err) {
				return true, err
			}
		}
		return true, nil
	}

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

func (g *azureKeyVaultManager) GetStatus(obj runtime.Object) (*v1alpha1.ASOStatus, error) {
	instance, err := g.convert(obj)
	if err != nil {
		return nil, err
	}
	return &instance.Status, nil
}

func (k *azureKeyVaultManager) convert(obj runtime.Object) (*v1alpha1.KeyVault, error) {
	local, ok := obj.(*v1alpha1.KeyVault)
	if !ok {
		return nil, fmt.Errorf("failed type assertion on kind: %s", obj.GetObjectKind().GroupVersionKind().String())
	}
	return local, nil
}

func NewOpsClient(keyvaultName string) *kvops.BaseClient {
	keyvaultClient := kvops.New()
	a, _ := iam.GetKeyvaultAuthorizer()
	keyvaultClient.Authorizer = a
	keyvaultClient.AddToUserAgent(config.UserAgent())
	return &keyvaultClient
}
