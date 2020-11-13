// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package keyvaults

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

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
	uuid "github.com/satori/go.uuid"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
)

type azureKeyVaultManager struct {
	Creds  config.Credentials
	Scheme *runtime.Scheme
}

func NewAzureKeyVaultManager(creds config.Credentials, scheme *runtime.Scheme) *azureKeyVaultManager {
	return &azureKeyVaultManager{
		Creds:  creds,
		Scheme: scheme,
	}
}

func getVaultsClient(creds config.Credentials) (keyvault.VaultsClient, error) {
	vaultsClient := keyvault.NewVaultsClientWithBaseURI(config.BaseURI(), creds.SubscriptionID())
	a, err := iam.GetResourceManagementAuthorizer(creds)
	if err != nil {
		return vaultsClient, err
	}
	vaultsClient.Authorizer = a
	vaultsClient.AddToUserAgent(config.UserAgent())
	return vaultsClient, nil
}

func getObjectID(ctx context.Context, creds config.Credentials, tenantID string, clientID string) (*string, error) {
	appclient := auth.NewApplicationsClient(tenantID)
	a, err := iam.GetGraphAuthorizer(creds)
	if err != nil {
		return nil, err
	}
	appclient.Authorizer = a
	appclient.AddToUserAgent(config.UserAgent())

	result, err := appclient.GetServicePrincipalsIDByAppID(ctx, clientID)
	if err != nil {
		return nil, err
	}
	return result.Value, nil
}

// ParseNetworkPolicy - helper function to parse network policies from Kubernetes spec
func ParseNetworkPolicy(ruleSet *v1alpha1.NetworkRuleSet) keyvault.NetworkRuleSet {
	var bypass keyvault.NetworkRuleBypassOptions
	switch ruleSet.Bypass {
	case "AzureServices":
		bypass = keyvault.AzureServices
	case "None":
		bypass = keyvault.None
	default:
		bypass = keyvault.AzureServices
	}

	var defaultAction keyvault.NetworkRuleAction
	switch ruleSet.DefaultAction {
	case "Allow":
		defaultAction = keyvault.Allow
	case "Deny":
		defaultAction = keyvault.Deny
	default:
		defaultAction = keyvault.Deny
	}

	var ipInstances []keyvault.IPRule
	if ruleSet.IPRules != nil {
		for _, i := range *ruleSet.IPRules {
			ip := i
			ipInstances = append(ipInstances, keyvault.IPRule{Value: &ip})
		}
	}

	var virtualNetworkRules []keyvault.VirtualNetworkRule
	if ruleSet.VirtualNetworkRules != nil {
		for _, i := range *ruleSet.VirtualNetworkRules {
			id := i
			virtualNetworkRules = append(virtualNetworkRules, keyvault.VirtualNetworkRule{ID: &id})
		}
	}

	networkAcls := keyvault.NetworkRuleSet{
		Bypass:              bypass,
		DefaultAction:       defaultAction,
		IPRules:             &ipInstances,
		VirtualNetworkRules: &virtualNetworkRules,
	}

	return networkAcls
}

// ParseAccessPolicy - helper function to parse access policies from Kubernetes spec
func ParseAccessPolicy(ctx context.Context, creds config.Credentials, policy *v1alpha1.AccessPolicyEntry) (keyvault.AccessPolicyEntry, error) {
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
		objID, err := getObjectID(ctx, creds, policy.TenantID, policy.ClientID)
		if err != nil {
			return keyvault.AccessPolicyEntry{}, err
		}
		newEntry.ObjectID = objID

	} else if policy.ObjectID != "" {
		newEntry.ObjectID = &policy.ObjectID
	}

	return newEntry, nil
}

// InstantiateVault will instantiate VaultsClient
func InstantiateVault(ctx context.Context, creds config.Credentials, vaultName string, containsUpdate bool) (keyvault.VaultsClient, uuid.UUID, error) {
	vaultsClient, err := getVaultsClient(creds)
	if err != nil {
		return keyvault.VaultsClient{}, uuid.UUID{}, err
	}
	id, err := uuid.FromString(creds.TenantID())
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
func (m *azureKeyVaultManager) CreateVault(ctx context.Context, instance *v1alpha1.KeyVault, sku azurev1alpha1.KeyVaultSku, tags map[string]*string, vaultExists bool) (keyvault.Vault, error) {
	vaultName := instance.Name
	location := instance.Spec.Location
	groupName := instance.Spec.ResourceGroup
	enableSoftDelete := instance.Spec.EnableSoftDelete

	vaultsClient, id, err := InstantiateVault(ctx, m.Creds, vaultName, instance.Status.ContainsUpdate)
	if err != nil {
		return keyvault.Vault{}, err
	}

	var accessPolicies []keyvault.AccessPolicyEntry
	if instance.Spec.AccessPolicies != nil {
		for _, policy := range *instance.Spec.AccessPolicies {
			policy := policy // Make a copy of the variable and redeclare it
			newEntry, err := ParseAccessPolicy(ctx, m.Creds, &policy)
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
		networkAcls = ParseNetworkPolicy(instance.Spec.NetworkPolicies)
	} else {
		networkAcls = keyvault.NetworkRuleSet{}
	}

	keyVaultSku := keyvault.Sku{
		Family: to.StringPtr("A"),
		Name:   keyvault.Standard,
	}

	if strings.ToLower(sku.Name) == "premium" {
		keyVaultSku.Name = keyvault.Premium
	}

	pols := []keyvault.AccessPolicyEntry{}
	params := keyvault.VaultCreateOrUpdateParameters{
		Properties: &keyvault.VaultProperties{
			TenantID:         &id,
			AccessPolicies:   &pols,
			Sku:              &keyVaultSku,
			NetworkAcls:      &networkAcls,
			EnableSoftDelete: &enableSoftDelete,
		},
		Location: to.StringPtr(location),
		Tags:     tags,
	}

	if vaultExists {
		params.Properties.AccessPolicies = &accessPolicies
	}

	future, err := vaultsClient.CreateOrUpdate(ctx, groupName, vaultName, params)
	if err != nil {
		return keyvault.Vault{}, err
	}

	return future.Result(vaultsClient)
}

//CreateVaultWithAccessPolicies creates a new key vault and provides access policies to the specified user
func (m *azureKeyVaultManager) CreateVaultWithAccessPolicies(ctx context.Context, groupName string, vaultName string, location string, clientID string) (keyvault.Vault, error) {
	vaultsClient, id, err := InstantiateVault(ctx, m.Creds, vaultName, false)
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
		objID, err := getObjectID(ctx, m.Creds, m.Creds.TenantID(), clientID)
		if err != nil {
			return keyvault.Vault{}, err
		}
		if objID != nil {
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
func (m *azureKeyVaultManager) DeleteVault(ctx context.Context, groupName string, vaultName string) (result autorest.Response, err error) {
	vaultsClient, err := getVaultsClient(m.Creds)
	if err != nil {
		return autorest.Response{}, err
	}
	return vaultsClient.Delete(ctx, groupName, vaultName)
}

// CheckExistence checks for the presence of a keyvault instance on Azure
func (m *azureKeyVaultManager) GetVault(ctx context.Context, groupName string, vaultName string) (result keyvault.Vault, err error) {
	vaultsClient, err := getVaultsClient(m.Creds)
	if err != nil {
		return keyvault.Vault{}, err
	}
	return vaultsClient.Get(ctx, groupName, vaultName)

}

func (m *azureKeyVaultManager) Ensure(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {
	instance, err := m.convert(obj)
	if err != nil {
		return true, err
	}

	// hash the spec
	hash := helpers.Hash256(instance.Spec)

	// convert kube labels to expected tag format
	labels := helpers.LabelsToTags(instance.GetLabels())

	instance.Status.Provisioning = true
	instance.Status.FailedProvisioning = false
	exists := false
	// Check if this KeyVault already exists and its state if it does.
	keyvault, err := m.GetVault(ctx, instance.Spec.ResourceGroup, instance.Name)
	if err == nil {
		exists = true
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

	keyvault, err = m.CreateVault(
		ctx,
		instance,
		instance.Spec.Sku,
		labels,
		exists,
	)
	if err != nil {
		done, err := HandleCreationError(instance, err)
		if done && exists {
			instance.Status.Message = "key vault created but access policies failed: " + instance.Status.Message
			return false, nil
		}

		return done, err
	}

	instance.Status.State = keyvault.Status
	if keyvault.ID != nil {
		instance.Status.ResourceId = *keyvault.ID
	}
	instance.Status.ContainsUpdate = false
	instance.Status.Provisioned = true
	instance.Status.Provisioning = false
	instance.Status.Message = resourcemanager.SuccessMsg

	return true, nil
}

func HandleCreationError(instance *v1alpha1.KeyVault, err error) (bool, error) {
	// let the user know what happened
	instance.Status.Message = errhelp.StripErrorTimes(errhelp.StripErrorIDs(err))
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
		errhelp.LocationNotAvailableForResourceType,
	}

	azerr := errhelp.NewAzureError(err)
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
		switch azerr.Type {
		case errhelp.AlreadyExists:
			timeNow := metav1.NewTime(time.Now())
			if timeNow.Sub(instance.Status.RequestedAt.Time) < (30 * time.Second) {
				instance.Status.Provisioning = true
				return false, nil
			}

		}
		instance.Status.Message = "Reconcilation hit unrecoverable error " + err.Error()
		return true, nil
	}

	if azerr.Code == http.StatusForbidden {
		// permission errors when applying access policies are generally worth waiting on
		return false, nil
	}

	// reconciliation not done and we don't know what happened
	return false, err
}

func (m *azureKeyVaultManager) Delete(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {
	instance, err := m.convert(obj)
	if err != nil {
		return true, err
	}

	_, err = m.GetVault(ctx, instance.Spec.ResourceGroup, instance.Name)
	if err == nil {
		_, err := m.DeleteVault(ctx, instance.Spec.ResourceGroup, instance.Name)
		if err != nil {
			catch := []string{
				errhelp.AsyncOpIncompleteError,
			}
			gone := []string{
				errhelp.ResourceGroupNotFoundErrorCode,
				errhelp.ParentNotFoundErrorCode,
				errhelp.NotFoundErrorCode,
				errhelp.ResourceNotFound,
			}
			azerr := errhelp.NewAzureError(err)
			if helpers.ContainsString(catch, azerr.Type) {
				return true, nil
			} else if helpers.ContainsString(gone, azerr.Type) {
				return false, nil
			}
			return true, err
		}
		return true, nil
	}

	return false, nil
}

func (m *azureKeyVaultManager) GetParents(obj runtime.Object) ([]resourcemanager.KubeParent, error) {

	instance, err := m.convert(obj)
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

func (m *azureKeyVaultManager) convert(obj runtime.Object) (*v1alpha1.KeyVault, error) {
	local, ok := obj.(*v1alpha1.KeyVault)
	if !ok {
		return nil, fmt.Errorf("failed type assertion on kind: %s", obj.GetObjectKind().GroupVersionKind().String())
	}
	return local, nil
}

func NewOpsClient(creds config.Credentials, keyvaultName string) *kvops.BaseClient {
	keyvaultClient := kvops.New()
	a, _ := iam.GetKeyvaultAuthorizer(creds)
	keyvaultClient.Authorizer = a
	keyvaultClient.AddToUserAgent(config.UserAgent())
	return &keyvaultClient
}
