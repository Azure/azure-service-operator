/*
Copyright 2019 microsoft.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package vnet

import (
	"context"
	"fmt"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
)

// Ensure makes sure that an API Mgmt Svc instance exists
func (g *AzureVNetManager) Ensure(ctx context.Context, obj runtime.Object) (bool, error) {

	instance, err := g.convert(obj)
	if err != nil {
		return false, err
	}

	location := instance.Spec.Location
	resourceGroup := instance.Spec.ResourceGroup
	resourceName := instance.Name
	addressSpace := instance.Spec.AddressSpace
	subnets := instance.Spec.Subnets

	instance.Status.Provisioning = true
	instance.Status.Provisioned = false
	_, err = g.CreateVNet(
		ctx,
		location,
		resourceGroup,
		resourceName,
		addressSpace,
		subnets,
	)
	if err != nil {
		return false, fmt.Errorf("Error creating VNet: %s, %s - %v", resourceGroup, resourceName, err)
	}
	instance.Status.Provisioning = false
	instance.Status.Provisioned = true

	return true, nil
}

// Delete makes sure that an API Mgmt Svc has been deleted
func (g *AzureVNetManager) Delete(ctx context.Context, obj runtime.Object) (bool, error) {

	instance, err := g.convert(obj)
	if err != nil {
		return false, err
	}

	resourceGroup := instance.Spec.ResourceGroup
	resourceName := instance.Name

	_, err = g.DeleteVNet(
		ctx,
		resourceGroup,
		resourceName,
	)
	if err != nil {
		return true, fmt.Errorf("Error deleting VNet: %s, %s - %v", resourceGroup, resourceName, err)
	}

	return true, nil
}

// GetParents lists the parents for an API Mgmt Svc
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

func (g *AzureVNetManager) convert(obj runtime.Object) (*azurev1alpha1.VirtualNetwork, error) {
	local, ok := obj.(*azurev1alpha1.VirtualNetwork)
	if !ok {
		return nil, fmt.Errorf("failed type assertion on kind: %s", obj.GetObjectKind().GroupVersionKind().String())
	}
	return local, nil
}
