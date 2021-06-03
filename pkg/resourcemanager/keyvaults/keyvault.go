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
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/to"
	uuid "github.com/gofrs/uuid"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"

	"github.com/Azure/azure-service-operator/api/v1alpha1"
	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/iam"
)

type AzureKeyVaultManager struct {
	Creds  config.Credentials
	Scheme *runtime.Scheme
}

func NewAzureKeyVaultManager(creds config.Credentials, scheme *runtime.Scheme) *AzureKeyVaultManager {
	return &AzureKeyVaultManager{
		Creds:  creds,
		Scheme: scheme,
	}
}

func GetKeyVaultClient(creds config.Credentials) (keyvault.VaultsClient, error) {
	vaultsClient := keyvault.NewVaultsClientWithBaseURI(config.BaseURI(), creds.SubscriptionID())
	a, err := iam.GetResourceManagementAuthorizer(creds)
	if err != nil {
		return vaultsClient, err
	}
	vaultsClient.Authorizer = a
	vaultsClient.AddToUserAgent(config.UserAgent())
	return vaultsClient, nil
}

func GetObjectID(ctx context.Context, creds config.Credentials, tenantID string, clientID string) (*string, error) {
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
		objID, err := GetObjectID(ctx, creds, policy.TenantID, policy.ClientID)
		if err != nil {
			return keyvault.AccessPolicyEntry{}, err
		}
		newEntry.ObjectID = objID
	} else if policy.ObjectID != "" {
		newEntry.ObjectID = &policy.ObjectID
	}

	return newEntry, nil
}

// CreateVault creates a new key vault
func (m *AzureKeyVaultManager) CreateVault(ctx context.Context, instance *v1alpha1.KeyVault, sku azurev1alpha1.KeyVaultSku, tags map[string]*string) (keyvault.Vault, error) {
	vaultName := instance.Name
	location := instance.Spec.Location
	groupName := instance.Spec.ResourceGroup
	enableSoftDelete := instance.Spec.EnableSoftDelete

	vaultsClient, err := GetKeyVaultClient(m.Creds)
	if err != nil {
		return keyvault.Vault{}, errors.Wrapf(err, "couldn't get vaults client")
	}
	id, err := uuid.FromString(m.Creds.TenantID())
	if err != nil {
		return keyvault.Vault{}, errors.Wrapf(err, "couldn't convert tenantID to UUID")
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
		// Policies must not be nil (API doesn't allow it)
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

	params := keyvault.VaultCreateOrUpdateParameters{
		Properties: &keyvault.VaultProperties{
			TenantID:         &id,
			AccessPolicies:   &accessPolicies,
			Sku:              &keyVaultSku,
			NetworkAcls:      &networkAcls,
			EnableSoftDelete: &enableSoftDelete,
		},
		Location: to.StringPtr(location),
		Tags:     tags,
	}

	future, err := vaultsClient.CreateOrUpdate(ctx, groupName, vaultName, params)
	if err != nil {
		return keyvault.Vault{}, err
	}

	return future.Result(vaultsClient)
}

// DeleteVault removes the resource group named by env var
func (m *AzureKeyVaultManager) DeleteVault(ctx context.Context, groupName string, vaultName string) (result autorest.Response, err error) {
	vaultsClient, err := GetKeyVaultClient(m.Creds)
	if err != nil {
		return autorest.Response{}, err
	}
	return vaultsClient.Delete(ctx, groupName, vaultName)
}

// CheckExistence checks for the presence of a keyvault instance on Azure
func (m *AzureKeyVaultManager) GetVault(ctx context.Context, groupName string, vaultName string) (result keyvault.Vault, err error) {
	vaultsClient, err := GetKeyVaultClient(m.Creds)
	if err != nil {
		return keyvault.Vault{}, err
	}
	return vaultsClient.Get(ctx, groupName, vaultName)

}

func (m *AzureKeyVaultManager) Ensure(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {
	instance, err := m.convert(obj)
	if err != nil {
		return true, err
	}

	// hash the spec
	hash := helpers.Hash256(instance.Spec)

	// convert kube labels to expected tag format
	labels := helpers.LabelsToTags(instance.GetLabels())

	instance.Status.SetProvisioning("")
	exists := false
	// Check if this KeyVault already exists and its state if it does.
	keyvault, err := m.GetVault(ctx, instance.Spec.ResourceGroup, instance.Name)
	if err == nil {
		exists = true
		if instance.Status.SpecHash == hash {
			instance.Status.SetProvisioned(resourcemanager.SuccessMsg)
			instance.Status.ResourceId = *keyvault.ID
			return true, nil
		}

		instance.Status.SpecHash = hash
	}

	keyvault, err = m.CreateVault(
		ctx,
		instance,
		instance.Spec.Sku,
		labels,
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
	instance.Status.SetProvisioned(resourcemanager.SuccessMsg)

	return true, nil
}

func HandleCreationError(instance *v1alpha1.KeyVault, err error) (bool, error) {
	// let the user know what happened
	instance.Status.Message = errhelp.StripErrorTimes(errhelp.StripErrorIDs(err))
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
		errhelp.ValidationError,
	}

	azerr := errhelp.NewAzureError(err)
	if helpers.ContainsString(catch, azerr.Type) {
		// reconciliation is not done but error is acceptable
		return false, nil
	}

	if helpers.ContainsString(catchUnrecoverableErrors, azerr.Type) {
		// Unrecoverable error, so stop reconcilation
		switch azerr.Type {
		case errhelp.AlreadyExists:
			timeNow := metav1.NewTime(time.Now())
			if timeNow.Sub(instance.Status.RequestedAt.Time) < (30 * time.Second) {
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

func (m *AzureKeyVaultManager) Delete(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {
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

func (m *AzureKeyVaultManager) GetParents(obj runtime.Object) ([]resourcemanager.KubeParent, error) {

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

func (g *AzureKeyVaultManager) GetStatus(obj runtime.Object) (*v1alpha1.ASOStatus, error) {
	instance, err := g.convert(obj)
	if err != nil {
		return nil, err
	}
	return &instance.Status, nil
}

func (m *AzureKeyVaultManager) convert(obj runtime.Object) (*v1alpha1.KeyVault, error) {
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
