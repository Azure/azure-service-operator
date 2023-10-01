/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package customizations

import (
	"context"

	keyvaultapi "github.com/Azure/azure-service-operator/v2/api/keyvault/v1api20210401preview"
	keyvault "github.com/Azure/azure-service-operator/v2/api/keyvault/v1api20210401previewstorage"

	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1api20200601storage"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	"sigs.k8s.io/controller-runtime/pkg/conversion"

	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/internal/resolver"
	"github.com/Azure/azure-service-operator/v2/internal/util/kubeclient"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/extensions"
)

var _ extensions.ARMResourceModifier = &VaultExtension{}

// ModifyARMResource implements extensions.ARMResourceModifier.
func (ex *VaultExtension) ModifyARMResource(
	ctx context.Context,
	armClient *genericarmclient.GenericClient,
	armObj genruntime.ARMResource,
	obj genruntime.ARMMetaObject,
	kubeClient kubeclient.Client,
	resolver *resolver.Resolver,
	log logr.Logger,
) (genruntime.ARMResource, error) {

	kv, ok := obj.(*keyvault.Vault)
	if !ok {
		return nil, errors.Errorf(
			"Cannot run VaultExtension.ModifyARMResource() with unexpected resource type %T",
			obj)
	}

	// Type assert that we are the hub type. This will fail to compile if
	// the hub type has been changed but this extension has not been updated to match
	var _ conversion.Hub = kv

	// If createMode is nil (!!), Default or Recover, nothing for us to do
	if kv.Spec.Properties == nil ||
		kv.Spec.Properties.CreateMode == nil ||
		*kv.Spec.Properties.CreateMode == string(keyvaultapi.VaultProperties_CreateMode_Default) ||
		*kv.Spec.Properties.CreateMode == string(keyvaultapi.VaultProperties_CreateMode_Recover) {
		return armObj, nil
	}

	// Find out whether a soft-deleted KeyVault with the same name exists
	exists, err := ex.checkForExistenceOfDeletedKeyVault(ctx, kv, armClient, kubeClient, resolver, log)
	if err != nil {
		// Couldn't determine whether a soft-deleted keyvault exists, assume it doesn't
		log.Error(err, "error checking for existence of soft-deleted KeyVault")
		return armObj, nil
	}

	if *kv.Spec.Properties.CreateMode == string(keyvaultapi.VaultProperties_CreateMode_CreateOrRecover) {
		// If a soft-deleted KeyVault exists, we need to recover it, otherwise we create a new one
		createMode := keyvaultapi.VaultProperties_CreateMode_Default
		if exists {
			createMode = keyvaultapi.VaultProperties_CreateMode_Recover
		}

		spec := armObj.Spec()
		err = reflecthelpers.SetProperty(spec, "Properties.CreateMode", &createMode)
		if err != nil {
			return nil, errors.Wrapf(err, "error setting CreateMode to %s", createMode)
		}

		return armObj, nil
	}

	return armObj, nil
}

// checkForExistenceOfDeletedKeyVault checks to see whether there's a soft deleted KeyVault with the same name.
// This might be true if another party has deleted the KeyVault, even if we previously created it
func (ex *VaultExtension) checkForExistenceOfDeletedKeyVault(
	ctx context.Context,
	kv *keyvault.Vault,
	armClient *genericarmclient.GenericClient,
	kubeClient kubeclient.Client,
	resolver *resolver.Resolver,
	log logr.Logger,
) (bool, error) {
	// Get the owner of the KeyVault, we need this resource group to determine the subscription
	owner, ownerErr := ex.getOwner(ctx, kv, resolver, log)
	if ownerErr != nil {
		return false, errors.Wrapf(ownerErr, "unable to find owner of KeyVault %s", kv.Name)
	}

	// Get the subscription of the KeyVault
	subscription, subscriptionErr := ex.getSubscription(owner)
	if subscriptionErr != nil {
		return false, errors.Wrapf(subscriptionErr, "unable to determine subscription for KeyVault %s", kv.Name)
	}

	// Get the location of the KeyVault
	location, locationOk := ex.getLocation(kv, owner)
	if !locationOk {
		return false, errors.Errorf("unable to determine location of KeyVault %s", kv.Name)
	}

	// Create the ARM ID of the soft-deleted KeyVault, if any
	armIdOfDeletedKeyVault, err := ex.createDeletedKeyVaultARMID(
		subscription,
		location,
		kv.Name)
	if err != nil {
		return false, errors.Wrapf(err, "failed to create ARM ID for potential soft-deleted KeyVault %s", kv.Name)
	}

	// Get the API version of the KeyVault API to use
	apiVersion, err := genruntime.GetAPIVersion(kv, kubeClient.Scheme())
	if err != nil {
		return false, errors.Wrapf(err, "unable to get api version for KeyVault")
	}

	// Check to see if the deleted keyvault exists
	exists, _, err := armClient.HeadByID(ctx, armIdOfDeletedKeyVault, apiVersion)
	if err != nil {
		return false, errors.Wrapf(err, "error checking for existence of deleted KeyVault %s", kv.Name)
	}

	log.Info(
		"Checking for existence of soft-deleted KeyVault",
		"KeyVault", kv.Name,
		"Subscription", subscription,
		"Location", location,
		"Exists", exists,
	)

	return exists, nil
}

func (*VaultExtension) getOwner(
	ctx context.Context,
	kv *keyvault.Vault,
	resolver *resolver.Resolver,
	log logr.Logger,
) (*resources.ResourceGroup, error) {
	owner, err := resolver.ResolveOwner(ctx, kv)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to resolve owner of KeyVault %s", kv.Name)
	}

	// No need to wait for resources that don't have an owner
	if !owner.FoundKubernetesOwner() {
		log.Info(
			"KeyVault owner is not within the cluster, cannot determine subscription",
			"keyVault", kv.Name)
		return nil, errors.Errorf("owner of KeyVault %s is not within the cluster", kv.Name)
	}

	rg, ok := owner.Owner.(*resources.ResourceGroup)
	if !ok {
		return nil, errors.Errorf("expected owner of KeyVault %s to be a ResourceGroup", kv.Name)
	}

	// Type assert that the ResourceGroup is the hub type. This will fail to compile if
	// the hub type has been changed but this extension has not been updated to match
	var _ conversion.Hub = rg

	return rg, nil
}

func (*VaultExtension) getSubscription(
	rg *resources.ResourceGroup,
) (string, error) {
	if rg.Status.Id == nil {
		return "", errors.Errorf(
			"unable to determine subscription as resource group %s has no id",
			rg.Name)
	}

	// Parse id into a more useful form
	armID, err := arm.ParseResourceID(*rg.Status.Id)
	if err != nil {
		return "", errors.Wrapf(
			err,
			"unable to parse ARMID of ResourceGroup %s",
			rg.AzureName())
	}

	return armID.SubscriptionID, nil
}

// findKeyVaultLocation determines which location we're trying to create KeyVault within
func (*VaultExtension) getLocation(
	kv *keyvault.Vault,
	rg *resources.ResourceGroup,
) (string, bool) {
	// Prefer location on the KeyVault
	if kv.Spec.Location != nil && *kv.Spec.Location != "" {
		return *kv.Spec.Location, true
	}

	// Fallback to location on ResourceGroup
	if rg.Spec.Location != nil && *rg.Spec.Location != "" {
		return *rg.Spec.Location, true
	}

	return "", false
}

func (*VaultExtension) getAPIVersion(
	kv *keyvault.Vault,
	kubeClient kubeclient.Client,
) (string, error) {
	api, err := genruntime.GetAPIVersion(kv, kubeClient.Scheme())
	if err != nil {
		return "", errors.Wrapf(err, "error getting api version for resource %s", kv.GetName())
	}

	return api, nil
}

func (*VaultExtension) createDeletedKeyVaultARMID(
	subscription string,
	location string,
	name string,
) (string, error) {
	// For a deleted KeyVault, the ARM ID has this format:
	// /subscriptions/{subscriptionId}/providers/Microsoft.KeyVault/locations/{location}/deletedVaults/{vaultName}?api-version=2022-07-01

	return genericarmclient.MakeSubscriptionScopeARMID(
		subscription,
		"Microsoft.KeyVault",
		"locations", location,
		"deletedVaults", name)
}

/*

  logr.go:278: I2023-09-27T19:52:21Z] VaultController "msg"="updated resource in etcd" name="aso-kv-gs" namespace="aso-test-keyvault-whenrecoverspecified-recoverssuccessfully" kind="&TypeMeta{Kind:Vault,APIVersion:keyvault.azure.com/v1api20210401previewstorage,}" resourceVersion="983" generation=1 uid="3d532c3b-e7b7-484c-9432-485c9bf4539e" ownerReferences=[]v1.OwnerReference{v1.OwnerReference{APIVersion:"resources.azure.com/v1api20200601storage", Kind:"ResourceGroup", Name:"asotest-rg-bmhvxo", UID:"77167b62-35a0-408e-b62f-a067ec8f0851", Controller:(*bool)(nil), BlockOwnerDeletion:(*bool)(nil)}} creationTimestamp="2023-09-27 19:43:35 +0000 UTC" deletionTimestamp=<nil> finalizers=[]string{"serviceoperator.azure.com/finalizer"} annotations=map[string]string{"serviceoperator.azure.com/resource-id":"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/asotest-rg-bmhvxo/providers/Microsoft.KeyVault/vaults/aso-kv-gs"} conditions="[Condition [Ready], Status = \"False\", ObservedGeneration = 1, Severity = \"Warning\", Reason = \"ConflictError\", Message = \"A vault with the same name already exists in deleted state. You need to either recover or purge existing key vault. Follow this link https://go.microsoft.com/fwlink/?linkid=2149745 for more information on soft delete.: PUT https://management.azure.com/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/asotest-rg-bmhvxo/providers/Microsoft.KeyVault/vaults/aso-kv-gs\\n--------------------------------------------------------------------------------\\nRESPONSE 409: 409 Conflict\\nERROR CODE: ConflictError\\n--------------------------------------------------------------------------------\\n{\\n  \\\"error\\\": {\\n    \\\"code\\\": \\\"ConflictError\\\",\\n    \\\"message\\\": \\\"A vault with the same name already exists in deleted state. You need to either recover or purge existing key vault. Follow this link https://go.microsoft.com/fwlink/?linkid=2149745 for more information on soft delete.\\\"\\n  }\\n}\\n--------------------------------------------------------------------------------\\n\", LastTransitionTime = \"2023-09-27 19:52:10 +0000 UTC\"]" owner="asotest-rg-bmhvxo, Group/Kind: resources.azure.com/ResourceGroup"

*/
