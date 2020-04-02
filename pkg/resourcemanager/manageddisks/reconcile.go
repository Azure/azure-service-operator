// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package manageddisks

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

// Ensure makes sure that an ManagedDisk instance exists
func (g *AzureManagedDiskManager) Ensure(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {

	instance, err := g.convert(obj)
	if err != nil {
		return false, err
	}

	location := instance.Spec.Location
	resourceGroup := instance.Spec.ResourceGroup
	resourceName := instance.Name
	createOption := instance.Spec.CreateOption
	diskSizeGB := instance.Spec.DiskSizeGB

	instance.Status.Provisioning = true
	instance.Status.Provisioned = false

	// check first to see if the ManagedDisk exists, if it does, dont create it and
	// 	consider the reconcilliation successful
	if exists, _ := g.ManagedDiskExists(ctx, resourceGroup, resourceName); exists {
		instance.Status.Provisioning = false
		instance.Status.Provisioned = true
		instance.Status.Message = resourcemanager.SuccessMsg
		return true, nil
	}

	_, err = g.CreateManageDisk(
		ctx,
		location,
		resourceGroup,
		resourceName,
		createOption,
		diskSizeGB,
	)
	if err != nil {
		azerr := errhelp.NewAzureErrorAzureError(err)
		catch := []string{
			errhelp.ResourceGroupNotFoundErrorCode,
			errhelp.ParentNotFoundErrorCode,
			errhelp.NotFoundErrorCode,
			errhelp.AsyncOpIncompleteError,
		}
		catchUnrecoverableErrors := []string{
			errhelp.NetcfgInvalidIPAddressPrefix,
			errhelp.NetcfgInvalidSubnet,
			errhelp.NetcfgInvalidVirtualNetworkSite,
			errhelp.InvalidCIDRNotation,
			errhelp.InvalidRequestFormat,
		}
		if helpers.ContainsString(catch, azerr.Type) {
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
			g.Telemetry.LogError("Reconcilation hit unrecoverable error", err)
			return true, nil
		}
		instance.Status.Provisioning = false
		return false, fmt.Errorf("Error creating ManagedDisk: %s, %s - %v", resourceGroup, resourceName, err)
	}

	// success
	instance.Status.Message = resourcemanager.SuccessMsg
	instance.Status.Provisioning = false
	instance.Status.Provisioned = true
	instance.Status.Message = resourcemanager.SuccessMsg

	return true, nil
}

// Delete makes sure that the ManagedDisk has been deleted
func (g *AzureManagedDiskManager) Delete(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {

	instance, err := g.convert(obj)
	if err != nil {
		return false, err
	}

	resourceGroup := instance.Spec.ResourceGroup
	resourceName := instance.Name

	g.DeleteManagedDisk(
		ctx,
		resourceGroup,
		resourceName,
	)

	return false, nil
}

// GetParents lists the parents for a ManagedDisk
func (g *AzureManagedDiskManager) GetParents(obj runtime.Object) ([]resourcemanager.KubeParent, error) {

	instance, err := g.convert(obj)
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

func (g *AzureManagedDiskManager) GetStatus(obj runtime.Object) (*azurev1alpha1.ASOStatus, error) {
	instance, err := g.convert(obj)
	if err != nil {
		return nil, err
	}
	return &instance.Status, nil
}

func (g *AzureManagedDiskManager) convert(obj runtime.Object) (*azurev1alpha1.AzureManagedDisk, error) {
	local, ok := obj.(*azurev1alpha1.AzureManagedDisk)
	if !ok {
		return nil, fmt.Errorf("failed type assertion on kind: %s", obj.GetObjectKind().GroupVersionKind().String())
	}
	return local, nil
}
