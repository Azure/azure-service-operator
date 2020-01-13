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

package apimservice

import (
	"context"
	"fmt"
	"strings"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
)

// Ensure makes sure that an API Mgmt Svc instance exists
func (g *AzureAPIMgmtServiceManager) Ensure(ctx context.Context, obj runtime.Object) (bool, error) {

	instance, err := g.convert(obj)
	if err != nil {
		return false, err
	}
	location := instance.Spec.Location
	resourceGroupName := instance.Spec.ResourceGroup
	resourceName := instance.ObjectMeta.Name

	exists, activated, _ := g.MgmtSvcStatus(ctx, resourceGroupName, resourceName)
	if !exists {

		// STEP 1:
		// 	need to provision
		// check name
		available, err := g.CheckAPIMgmtSvcName(ctx, resourceName)
		if err != nil {
			return false, err
		}
		if available {

			// create service
			instance.Status.Provisioned = false
			instance.Status.Provisioning = false
			_, err := g.CreateAPIMgmtSvc(ctx, location, resourceGroupName, resourceName)
			if err != nil {
				return false, fmt.Errorf("API Mgmt Svc create error %v", err)
			}
			instance.Status.Provisioning = true
			return false, nil
		} else {
			g.Telemetry.LogError("could not create API Mgmt Service due to bad resource name", fmt.Errof())
			return true, nil
		}
	} else {
		if !activated {

			// STEP 2:
			// 	still in the proccess of provisioning
			instance.Status.Provisioned = false
			instance.Status.Provisioning = true
			return false, nil
		} else {

			// STEP 3:
			// 	provisioned, now need to update with a vnet?
			vnetType := instance.Spec.VnetType
			if vnetType != "" && !strings.EqualFold(vnetType, "none") {
				vnetResourceGroup := instance.Spec.VnetType
				vnetName := instance.Spec.VnetType
				subnetName := instance.Spec.VnetSubnetName
				err = g.SetVNetForAPIMgmtSvc(
					ctx,
					resourceGroupName,
					resourceName,
					vnetType,
					vnetResourceGroup,
					vnetName,
					subnetName,
				)
				if err != nil {
					instance.Status.Provisioned = false
					instance.Status.Provisioning = true
					return false, fmt.Errorf("API Mgmt Svc could not update VNet %s, %s - %v", vnetResourceGroup, vnetName, err)
				}
			}

			// STEP 4:
			// 	everything is now completed!
			instance.Status.Provisioned = true
			instance.Status.Provisioning = false
			return true, nil
		}
	}
}

// Delete makes sure that an API Mgmt Svc has been deleted
func (g *AzureAPIMgmtServiceManager) Delete(ctx context.Context, obj runtime.Object) (bool, error) {

	instance, err := g.convert(obj)
	if err != nil {
		return false, err
	}

	resourceGroupName := instance.Spec.ResourceGroup
	resourceName := instance.ObjectMeta.Name

	_, err = g.DeleteAPIMgmtSvc(ctx, resourceGroupName, resourceName)
	if err != nil {
		return true, fmt.Errorf("API Mgmt Svc delete error %v", err)
	}

	return false, nil
}

// GetParents lists the parents for an API Mgmt Svc
func (g *AzureAPIMgmtServiceManager) GetParents(obj runtime.Object) ([]resourcemanager.KubeParent, error) {

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

func (g *AzureAPIMgmtServiceManager) convert(obj runtime.Object) (*azurev1alpha1.ApimService, error) {
	local, ok := obj.(*azurev1alpha1.ApimService)
	if !ok {
		return nil, fmt.Errorf("failed type assertion on kind: %s", obj.GetObjectKind().GroupVersionKind().String())
	}
	return local, nil
}
