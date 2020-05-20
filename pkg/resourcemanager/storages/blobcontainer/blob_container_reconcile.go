// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package blobcontainer

import (
	"context"
	"fmt"
	"strings"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/api/v1alpha2"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
)

// Ensure creates a blob container
func (bc *AzureBlobContainerManager) Ensure(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {

	instance, err := bc.convert(obj)
	if err != nil {
		return false, err
	}

	instance.Status.Provisioning = true
	instance.Status.Provisioned = false

	groupName := instance.Spec.ResourceGroup
	accountName := instance.Spec.AccountName
	containerName := instance.ObjectMeta.Name
	accessLevel := instance.Spec.AccessLevel

	// Create the blob container
	_, err = bc.CreateBlobContainer(ctx, groupName, accountName, containerName, accessLevel)
	if err != nil {
		instance.Status.Message = err.Error()

		// WIP: Validation error handling investigation
		// Log TypeOf error to identify type of MinLength/MaxLength validation error
		if strings.Contains(err.Error(), "MinLength") || strings.Contains(err.Error(), "MaxLength") {
			return true, nil
		}
		// END WIP: Validation error handling investigation

		azerr := errhelp.NewAzureErrorAzureError(err)
		catchIgnorable := []string{
			errhelp.ParentNotFoundErrorCode,
			errhelp.ResourceGroupNotFoundErrorCode,
			errhelp.NotFoundErrorCode,
			errhelp.StorageAccountIsNotProvisioned,
		}
		if helpers.ContainsString(catchIgnorable, azerr.Type) {
			instance.Status.Provisioning = false
			return false, nil
		}

		catchWaiting := []string{
			errhelp.AsyncOpIncompleteError,
		}
		if helpers.ContainsString(catchWaiting, azerr.Type) {
			return false, nil
		}

		catchNotIgnorable := []string{
			errhelp.ContainerOperationFailure, // Container name was invalid
			errhelp.ValidationError,           // Some validation (such as min/max length for ContainerName) failed
		}
		if helpers.ContainsString(catchNotIgnorable, azerr.Type) {
			instance.Status.Provisioning = false
			return true, nil
		}

		return false, err
	}

	instance.Status.Provisioning = false
	instance.Status.Provisioned = true
	instance.Status.Message = resourcemanager.SuccessMsg

	return true, nil
}

// Delete drops a blob container
func (bc *AzureBlobContainerManager) Delete(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {
	instance, err := bc.convert(obj)
	if err != nil {
		return false, err
	}

	groupName := instance.Spec.ResourceGroup
	accountName := instance.Spec.AccountName
	containerName := instance.ObjectMeta.Name

	_, err = bc.DeleteBlobContainer(ctx, groupName, accountName, containerName)
	if err != nil {
		instance.Status.Message = err.Error()

		azerr := errhelp.NewAzureErrorAzureError(err)
		catch := []string{
			errhelp.AsyncOpIncompleteError,
		}
		if helpers.ContainsString(catch, azerr.Type) {
			return true, nil
		}
		catch = []string{
			errhelp.ParentNotFoundErrorCode,
			errhelp.ResourceGroupNotFoundErrorCode,
			errhelp.NotFoundErrorCode,
		}
		if helpers.ContainsString(catch, azerr.Type) {
			return false, nil
		}

		return true, err
	}

	return false, nil
}

// GetParents returns the parents of a blob container
func (bc *AzureBlobContainerManager) GetParents(obj runtime.Object) ([]resourcemanager.KubeParent, error) {
	instance, err := bc.convert(obj)
	if err != nil {
		return nil, err
	}

	return []resourcemanager.KubeParent{
		{
			Key: types.NamespacedName{
				Name:      instance.Spec.AccountName,
				Namespace: instance.Namespace,
			},
			Target: &azurev1alpha1.StorageAccount{},
		},
		{
			Key: types.NamespacedName{
				Name:      instance.Spec.ResourceGroup,
				Namespace: instance.Namespace,
			},
			Target: &azurev1alpha1.ResourceGroup{},
		},
	}, nil

}

func (bc *AzureBlobContainerManager) GetStatus(obj runtime.Object) (*azurev1alpha1.ASOStatus, error) {
	instance, err := bc.convert(obj)
	if err != nil {
		return nil, err
	}
	st := azurev1alpha1.ASOStatus(instance.Status)
	return &st, nil
}

func (bc *AzureBlobContainerManager) convert(obj runtime.Object) (*v1alpha2.BlobContainer, error) {
	local, ok := obj.(*v1alpha2.BlobContainer)
	if !ok {
		return nil, fmt.Errorf("failed type assertion on kind: %s", obj.GetObjectKind().GroupVersionKind().String())
	}
	return local, nil
}
