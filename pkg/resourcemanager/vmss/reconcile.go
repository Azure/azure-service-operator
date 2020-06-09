// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package vmss

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

func (g *AzureVMScaleSetClient) Ensure(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {

	instance, err := g.convert(obj)
	if err != nil {
		return true, err
	}

	client := getVMScaleSetClient()

	location := instance.Spec.Location
	resourceGroup := instance.Spec.ResourceGroup
	resourceName := instance.Name
	vmSize := instance.Spec.VMSize
	capacity := int64(instance.Spec.Capacity)
	osType := instance.Spec.OSType
	adminUserName := instance.Spec.AdminUserName
	sshPublicKeyData := instance.Spec.SSHPublicKeyData
	imageURN := instance.Spec.PlatformImageURN
	vnetName := instance.Spec.VirtualNetworkName
	subnetName := instance.Spec.SubnetName
	loadBalancerName := instance.Spec.LoadBalancerName
	bePoolName := instance.Spec.BackendAddressPoolName
	natPoolName := instance.Spec.InboundNatPoolName

	// Check to see if secret exists and if yes retrieve the admin login and password
	secret, err := g.GetOrPrepareSecret(ctx, instance)
	if err != nil {
		return false, err
	}
	// Update secret
	err = g.AddVMScaleSetCredsToSecrets(ctx, instance.Name, secret, instance)
	if err != nil {
		return false, err
	}

	adminPassword := string(secret["password"])

	instance.Status.Provisioning = true
	// Check if this item already exists. This is required
	// to overcome the issue with the lack of idempotence of the Create call
	item, err := g.GetVMScaleSet(ctx, resourceGroup, resourceName)
	if err == nil {
		instance.Status.Provisioned = true
		instance.Status.Provisioning = false
		instance.Status.Message = resourcemanager.SuccessMsg
		instance.Status.ResourceId = *item.ID
		return true, nil
	}

	future, err := g.CreateVMScaleSet(
		ctx,
		location,
		resourceGroup,
		resourceName,
		vmSize,
		capacity,
		string(osType),
		adminUserName,
		adminPassword,
		sshPublicKeyData,
		imageURN,
		vnetName,
		subnetName,
		loadBalancerName,
		bePoolName,
		natPoolName,
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
			errhelp.ResourceNotFound,
			errhelp.InvalidResourceReference,
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
		// reconciliation not done and we don't know what happened
		return false, err
	}

	_, err = future.Result(client)
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
			errhelp.SubscriptionDoesNotHaveServer,
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
		// reconciliation not done and we don't know what happened
		return false, err
	}

	if instance.Status.Provisioning {
		instance.Status.Provisioned = true
		instance.Status.Provisioning = false
		instance.Status.Message = resourcemanager.SuccessMsg
	} else {
		instance.Status.Provisioned = false
		instance.Status.Provisioning = true
	}
	return true, nil
}

func (g *AzureVMScaleSetClient) Delete(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {

	instance, err := g.convert(obj)
	if err != nil {
		return true, err
	}

	resourceGroup := instance.Spec.ResourceGroup
	resourceName := instance.Name

	status, err := g.DeleteVMScaleSet(
		ctx,
		resourceName,
		resourceGroup,
	)
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

	if err == nil {
		if status != "InProgress" {
			return false, nil
		}
	}

	return true, nil
}
func (g *AzureVMScaleSetClient) GetParents(obj runtime.Object) ([]resourcemanager.KubeParent, error) {

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

func (g *AzureVMScaleSetClient) GetStatus(obj runtime.Object) (*azurev1alpha1.ASOStatus, error) {

	instance, err := g.convert(obj)
	if err != nil {
		return nil, err
	}
	return &instance.Status, nil
}

func (g *AzureVMScaleSetClient) convert(obj runtime.Object) (*azurev1alpha1.AzureVMScaleSet, error) {
	local, ok := obj.(*azurev1alpha1.AzureVMScaleSet)
	if !ok {
		return nil, fmt.Errorf("failed type assertion on kind: %s", obj.GetObjectKind().GroupVersionKind().String())
	}
	return local, nil
}
