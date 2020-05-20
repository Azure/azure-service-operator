// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package vnet

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/network/mgmt/network"
	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
)

// Ensure makes sure that an VNet instance exists
func (g *AzureVNetManager) Ensure(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {

	instance, err := g.convert(obj)
	if err != nil {
		return false, err
	}

	location := instance.Spec.Location
	resourceGroup := instance.Spec.ResourceGroup
	resourceName := instance.Name
	addressSpace := instance.Spec.AddressSpace
	subnets := instance.Spec.Subnets

	// check first to see if the VNet exists, if it does, dont create it and
	// 	consider the reconcilliation successful
	vNet, err := g.GetVNet(ctx, resourceGroup, resourceName)
	if err == nil {
		if vNet.ProvisioningState == network.Succeeded {
			// succeeded! end reconcilliation successfully
			instance.Status.Provisioning = false
			instance.Status.Provisioned = true
			instance.Status.FailedProvisioning = false
			instance.Status.Message = resourcemanager.SuccessMsg
			instance.Status.ResourceId = *vNet.ID
			return true, nil
		}
	}

	instance.Status.Provisioning = true
	result, err := g.CreateVNet(
		ctx,
		location,
		resourceGroup,
		resourceName,
		addressSpace,
		subnets,
	)
	if err != nil {
		azerr := errhelp.NewAzureErrorAzureError(err)
		instance.Status.Message = err.Error()

		if result.Response.Response != nil && result.Response.Response.StatusCode == http.StatusBadRequest {
			instance.Status.Provisioning = false
			return true, nil
		}

		catch := []string{
			errhelp.ResourceGroupNotFoundErrorCode,
			errhelp.ParentNotFoundErrorCode,
			errhelp.NotFoundErrorCode,
		}
		catchUnrecoverableErrors := []string{
			errhelp.NetcfgInvalidIPAddressPrefix,
			errhelp.NetcfgInvalidSubnet,
			errhelp.NetcfgInvalidVirtualNetworkSite,
			errhelp.InvalidCIDRNotation,
			errhelp.InvalidRequestFormat,
			errhelp.LocationNotAvailableForResourceType,
			errhelp.InvalidAddressPrefixFormat,
		}

		// everything ok - just requeue
		if strings.Contains(azerr.Type, errhelp.AsyncOpIncompleteError) {
			return false, nil
		}

		// reconciliation is not done but error is acceptable
		if helpers.ContainsString(catch, azerr.Type) {
			instance.Status.Provisioning = false
			return false, nil

		}

		instance.Status.Provisioning = false

		// Unrecoverable error, so stop reconcilation
		if helpers.ContainsString(catchUnrecoverableErrors, azerr.Type) {
			instance.Status.FailedProvisioning = true
			instance.Status.Message = fmt.Sprintf("Reconcilation hit unrecoverable error: %s", errhelp.StripErrorIDs(err))
			return true, nil
		}

		return false, fmt.Errorf("Error creating VNet: %s, %s - %v", resourceGroup, resourceName, err)
	}

	return false, nil
}

// Delete makes sure that the VNet has been deleted
func (g *AzureVNetManager) Delete(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {

	instance, err := g.convert(obj)
	if err != nil {
		return false, err
	}

	resourceGroup := instance.Spec.ResourceGroup
	resourceName := instance.Name

	g.DeleteVNet(
		ctx,
		resourceGroup,
		resourceName,
	)

	return false, nil
}

// GetParents lists the parents for a VNet
func (g *AzureVNetManager) GetParents(obj runtime.Object) ([]resourcemanager.KubeParent, error) {

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

func (g *AzureVNetManager) GetStatus(obj runtime.Object) (*azurev1alpha1.ASOStatus, error) {
	instance, err := g.convert(obj)
	if err != nil {
		return nil, err
	}
	return &instance.Status, nil
}

func (g *AzureVNetManager) convert(obj runtime.Object) (*azurev1alpha1.VirtualNetwork, error) {
	local, ok := obj.(*azurev1alpha1.VirtualNetwork)
	if !ok {
		return nil, fmt.Errorf("failed type assertion on kind: %s", obj.GetObjectKind().GroupVersionKind().String())
	}
	return local, nil
}
