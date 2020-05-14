// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package vmext

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

func (g *AzureVirtualMachineExtensionClient) Ensure(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {

	instance, err := g.convert(obj)
	if err != nil {
		return true, err
	}

	client := getVirtualMachineExtensionClient()

	location := instance.Spec.Location
	resourceGroup := instance.Spec.ResourceGroup
	extName := instance.Name
	vmName := instance.Spec.VMName
	autoUpgradeMinorVersion := instance.Spec.AutoUpgradeMinorVersion
	forceUpdateTag := instance.Spec.ForceUpdateTag
	publisher := instance.Spec.Publisher
	typeName := instance.Spec.TypeName
	typeHandlerVersion := instance.Spec.TypeHandlerVersion
	settings := instance.Spec.Settings
	protectedSettings := instance.Spec.ProtectedSettings

	// Check to see if secret exists and if yes retrieve the admin login and password
	secret, err := g.GetOrPrepareSecret(ctx, instance)
	if err != nil {
		return false, err
	}
	// Update secret
	err = g.AddVirtualMachineExtensionCredsToSecrets(ctx, instance.Name, secret, instance)
	if err != nil {
		return false, err
	}

	// Use the protected settings piped from the secret if the regular input one is null or empty
	if len(protectedSettings) <= 0 {
		secretProtectedSettings := string(secret["protecetdSettings"])
		protectedSettings = secretProtectedSettings
	}

	instance.Status.Provisioning = true
	// Check if this item already exists. This is required
	// to overcome the issue with the lack of idempotence of the Create call
	item, err := g.GetVirtualMachineExtension(ctx, resourceGroup, vmName, extName)
	if err == nil {
		instance.Status.Provisioned = true
		instance.Status.Provisioning = false
		instance.Status.Message = resourcemanager.SuccessMsg
		instance.Status.ResourceId = *item.ID
		return true, nil
	}
	future, err := g.CreateVirtualMachineExtension(
		ctx,
		location,
		resourceGroup,
		vmName,
		extName,
		autoUpgradeMinorVersion,
		forceUpdateTag,
		publisher,
		typeName,
		typeHandlerVersion,
		settings,
		protectedSettings,
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

func (g *AzureVirtualMachineExtensionClient) Delete(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {

	instance, err := g.convert(obj)
	if err != nil {
		return true, err
	}

	resourceGroup := instance.Spec.ResourceGroup
	vmName := instance.Spec.VMName
	extName := instance.Name

	status, err := g.DeleteVirtualMachineExtension(
		ctx,
		extName,
		vmName,
		resourceGroup,
	)
	if err != nil {
		if !errhelp.IsAsynchronousOperationNotComplete(err) {
			return true, err
		}
	}

	if err == nil {
		if status != "InProgress" {
			return false, nil
		}
	}

	return true, nil
}
func (g *AzureVirtualMachineExtensionClient) GetParents(obj runtime.Object) ([]resourcemanager.KubeParent, error) {

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

func (g *AzureVirtualMachineExtensionClient) GetStatus(obj runtime.Object) (*azurev1alpha1.ASOStatus, error) {

	instance, err := g.convert(obj)
	if err != nil {
		return nil, err
	}
	return &instance.Status, nil
}

func (g *AzureVirtualMachineExtensionClient) convert(obj runtime.Object) (*azurev1alpha1.AzureVirtualMachineExtension, error) {
	local, ok := obj.(*azurev1alpha1.AzureVirtualMachineExtension)
	if !ok {
		return nil, fmt.Errorf("failed type assertion on kind: %s", obj.GetObjectKind().GroupVersionKind().String())
	}
	return local, nil
}
