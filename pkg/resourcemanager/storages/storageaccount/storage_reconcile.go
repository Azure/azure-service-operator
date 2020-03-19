// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package storageaccount

import (
	"context"
	"fmt"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
)

// Ensure creates an AzureSqlDb
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

	// convert kube labels to expected tag format
	labels := map[string]*string{}
	for k, v := range instance.GetLabels() {
		value := v
		labels[k] = &value
	}

	stor, err := sa.GetStorage(ctx, groupName, name)
	if err != nil {
		azerr := errhelp.NewAzureErrorAzureError(err)
		if azerr.Type != errhelp.ResourceNotFound && instance.Status.Provisioning == true {
			instance.Status.Provisioning = false
			return false, nil
		}

		instance.Status.State = "NotReady"
	} else {
		instance.Status.State = string(stor.ProvisioningState)
	}

	if instance.Status.State == "Succeeded" {
		instance.Status.Message = resourcemanager.SuccessMsg
		instance.Status.Provisioned = true
		instance.Status.Provisioning = false
		return true, nil
	}

	instance.Status.Provisioning = true
	instance.Status.Provisioned = false

	_, err = sa.CreateStorage(ctx, groupName, name, location, sku, kind, labels, accessTier, enableHTTPSTrafficOnly, dataLakeEnabled)
	if err != nil {
		instance.Status.Message = err.Error()

		catch := []string{
			errhelp.ParentNotFoundErrorCode,
			errhelp.ResourceGroupNotFoundErrorCode,
			errhelp.AccountNameInvalid,
			errhelp.AlreadyExists,
			errhelp.AsyncOpIncompleteError,
		}
		azerr := errhelp.NewAzureErrorAzureError(err)
		if helpers.ContainsString(catch, azerr.Type) {
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
					instance.Status.Provisioning = false
					return true, nil
				}
			}

			if azerr.Type == errhelp.AccountNameInvalid {
				instance.Status.Message = "Invalid Storage Account Name"
				return true, nil
			}

			return false, nil
		}
		return false, err
	}

	return false, nil
}

// Delete drops a AzureSqlDb
func (sa *azureStorageManager) Delete(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {
	instance, err := sa.convert(obj)
	if err != nil {
		return false, err
	}

	name := instance.ObjectMeta.Name
	groupName := instance.Spec.ResourceGroup
	_, err = sa.DeleteStorage(ctx, groupName, name)
	if err != nil {
		if errhelp.IsStatusCode204(err) {
			return true, nil
		}
		if errhelp.IsStatusCode404(err) {
			return false, nil
		}

		return true, err
	}

	return false, nil
}

// GetParents returns the parents of AzureSqlDatabase
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
			Target: &azurev1alpha1.AzureSqlServer{},
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

func (sa *azureStorageManager) convert(obj runtime.Object) (*azurev1alpha1.Storage, error) {
	local, ok := obj.(*azurev1alpha1.Storage)
	if !ok {
		return nil, fmt.Errorf("failed type assertion on kind: %s", obj.GetObjectKind().GroupVersionKind().String())
	}
	return local, nil
}
