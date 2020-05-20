// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package storageaccount

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2019-04-01/storage"
	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/pollclient"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
)

// Ensure creates a storage account
func (sa *azureStorageManager) Ensure(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {

	instance, err := sa.convert(obj)
	if err != nil {
		return false, err
	}

	location := instance.Spec.Location
	name := instance.ObjectMeta.Name
	groupName := instance.Spec.ResourceGroup
	sku := instance.Spec.Sku
	kind := instance.Spec.Kind
	accessTier := instance.Spec.AccessTier
	enableHTTPSTrafficOnly := instance.Spec.EnableHTTPSTrafficOnly
	dataLakeEnabled := instance.Spec.DataLakeEnabled
	pollURL := instance.Status.PollingURL

	networkAcls := storage.NetworkRuleSet{}

	if instance.Spec.NetworkRule != nil {
		networkAcls = ParseNetworkPolicy(instance.Spec.NetworkRule)
	}
	// convert kube labels to expected tag format
	labels := helpers.LabelsToTags(instance.GetLabels())

	hash := ""
	stor, err := sa.GetStorage(ctx, groupName, name)
	if err != nil {
		instance.Status.Message = err.Error()
		instance.Status.State = "NotReady"

		// handle failures in the async operation
		if pollURL != "" {
			pClient := pollclient.NewPollClient()
			res, err := pClient.Get(ctx, pollURL)
			azerr := errhelp.NewAzureErrorAzureError(err)
			if err != nil {
				if azerr.Type == errhelp.NetworkAclsValidationFailure {
					instance.Status.Message = "Unable to provision Azure Storage Account due to error: " + errhelp.StripErrorIDs(err)
					instance.Status.Provisioning = false
					instance.Status.Provisioned = false
					return true, nil
				}
				return false, err
			}

			if res.Status == "Failed" {
				instance.Status.Message = res.Error.Error()
				instance.Status.Provisioning = false
				return true, nil
			}
		}
	} else {
		instance.Status.State = string(stor.ProvisioningState)

		hash = helpers.Hash256(instance.Spec)
		if instance.Status.SpecHash == hash && (instance.Status.Provisioned || instance.Status.FailedProvisioning) {
			instance.Status.RequestedAt = nil
			return true, nil
		}
	}

	if instance.Status.State == "Succeeded" {

		// upsert
		err = sa.StoreSecrets(ctx, groupName, name, instance)
		if err != nil {
			return false, err
		}

		// everything finished successfully!
		instance.Status.Message = resourcemanager.SuccessMsg
		instance.Status.Provisioned = true
		instance.Status.Provisioning = false
		instance.Status.SpecHash = hash
		instance.Status.ResourceId = *stor.ID
		instance.Status.PollingURL = ""
		return true, nil
	}

	instance.Status.Provisioning = true
	instance.Status.Provisioned = false

	pollURL, _, err = sa.CreateStorage(ctx, groupName, name, location, sku, kind, labels, accessTier, enableHTTPSTrafficOnly, dataLakeEnabled, &networkAcls)
	if err != nil {
		instance.Status.Message = err.Error()
		azerr := errhelp.NewAzureErrorAzureError(err)
		instance.Status.Provisioning = false

		ignore := []string{
			errhelp.ParentNotFoundErrorCode,
			errhelp.ResourceGroupNotFoundErrorCode,
		}
		if helpers.ContainsString(ignore, azerr.Type) {
			instance.Status.Provisioning = false
			return false, nil
		}

		wait := []string{
			errhelp.AsyncOpIncompleteError,
			errhelp.AlreadyExists,
		}
		if helpers.ContainsString(wait, azerr.Type) {

			if azerr.Type == errhelp.AlreadyExists {
				// This error could happen in two cases - when the storage account
				// exists in some other resource group or when this is a repeat
				// call to the reconcile loop for an update of this exact resource. So
				// we call a Get to check if this is the current resource and if
				// yes, we let the call go through instead of ending the reconcile loop
				_, err := sa.GetStorage(ctx, instance.Spec.ResourceGroup, instance.ObjectMeta.Name)
				if err != nil {
					// This means that the Server exists elsewhere and we should
					// terminate the reconcile loop
					instance.Status.Message = "Storage Account Already exists somewhere else"
					return true, nil
				}

				instance.Status.Message = "Storage Account already exists and should be available shortly"
				instance.Status.Provisioning = true
			}

			if azerr.Type == errhelp.AsyncOpIncompleteError {
				instance.Status.Provisioning = true
				instance.Status.PollingURL = pollURL
			}
			return false, nil
		}

		stop := []string{
			errhelp.AccountNameInvalid,
			errhelp.NetworkAclsValidationFailure,
		}
		if helpers.ContainsString(stop, azerr.Type) {
			instance.Status.Message = "Unable to provision Azure Storage Account due to error: " + errhelp.StripErrorIDs(err)
			instance.Status.Provisioning = false
			instance.Status.Provisioned = false
			return true, nil
		}

		return false, err
	}

	return false, nil
}

// Delete drops a storage account
func (sa *azureStorageManager) Delete(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {
	instance, err := sa.convert(obj)
	if err != nil {
		return false, err
	}

	name := instance.ObjectMeta.Name
	groupName := instance.Spec.ResourceGroup
	key := types.NamespacedName{
		Name: fmt.Sprintf("storageaccount-%s-%s",
			instance.Spec.ResourceGroup,
			instance.Name),
		Namespace: instance.Namespace,
	}

	_, err = sa.DeleteStorage(ctx, groupName, name)
	if err != nil {
		catch := []string{
			errhelp.ValidationError,
			errhelp.ResourceGroupNotFoundErrorCode,
		}

		err = errhelp.NewAzureError(err)
		if azerr, ok := err.(*errhelp.AzureError); ok {
			if helpers.ContainsString(catch, azerr.Type) {
				// Best case deletion of secrets
				sa.SecretClient.Delete(ctx, key)
				return false, nil
			}
		}
		return true, err
	}

	_, err = sa.GetStorage(ctx, groupName, name)
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
		azerr := errhelp.NewAzureErrorAzureError(err)
		if helpers.ContainsString(catch, azerr.Type) {
			return true, nil
		} else if helpers.ContainsString(gone, azerr.Type) {
			return false, nil
		}
		return true, err
	}
	return true, nil
}

// GetParents returns the parents of a storage account
func (sa *azureStorageManager) GetParents(obj runtime.Object) ([]resourcemanager.KubeParent, error) {
	instance, err := sa.convert(obj)
	if err != nil {
		return nil, err
	}

	return []resourcemanager.KubeParent{
		{
			Key: types.NamespacedName{
				Name:      instance.Spec.ResourceGroup,
				Namespace: instance.Namespace,
			},
			Target: &azurev1alpha1.ResourceGroup{},
		},
	}, nil
}

func (g *azureStorageManager) GetStatus(obj runtime.Object) (*azurev1alpha1.ASOStatus, error) {
	instance, err := g.convert(obj)
	if err != nil {
		return nil, err
	}
	return &instance.Status, nil
}

func (sa *azureStorageManager) convert(obj runtime.Object) (*azurev1alpha1.StorageAccount, error) {
	local, ok := obj.(*azurev1alpha1.StorageAccount)
	if !ok {
		return nil, fmt.Errorf("failed type assertion on kind: %s", obj.GetObjectKind().GroupVersionKind().String())
	}
	return local, nil
}
