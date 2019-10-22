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

package resourcegroups

import (
	"context"
	"fmt"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	helpers "github.com/Azure/azure-service-operator/pkg/helpers"
	"k8s.io/apimachinery/pkg/runtime"
)

func (g *AzureResourceGroupManager) Ensure(ctx context.Context, obj runtime.Object) (bool, error) {

	instance, err := g.convert(obj)
	if err != nil {
		return false, err
	}
	resourcegroupLocation := instance.Spec.Location
	resourcegroupName := instance.ObjectMeta.Name
	instance.Status.Provisioning = true

	g.Log.Info("creating resource group", "name", resourcegroupName, "location", resourcegroupLocation)
	_, err = g.CreateGroup(ctx, resourcegroupName, resourcegroupLocation)
	if err != nil {

		// catch := []string{}
		// err = errhelp.NewAzureError(err)
		// if azerr, ok := err.(*errhelp.AzureError); ok {
		// 	if helpers.ContainsString(catch, azerr.Type) {
		// 		return true, nil
		// 	}
		// }
		instance.Status.Provisioned = false
		return false, fmt.Errorf("ResourceGroup create error %v", err)

	}
	if instance.Status.Provisioning {
		instance.Status.Provisioned = true
		instance.Status.Provisioning = false
	} else {
		instance.Status.Provisioned = false
		instance.Status.Provisioning = true
	}

	return true, nil
}

func (g *AzureResourceGroupManager) Delete(ctx context.Context, obj runtime.Object) (bool, error) {
	instance, err := g.convert(obj)
	if err != nil {
		return false, err
	}

	resourcegroup := instance.ObjectMeta.Name
	g.Log.Info("Deleting resource group", "name", resourcegroup)

	_, err = g.DeleteGroup(ctx, resourcegroup)
	if err != nil {

		catch := []string{
			errhelp.ResourceGroupNotFoundErrorCode,
			errhelp.AsyncOpIncompleteError,
		}
		err = errhelp.NewAzureError(err)
		if azerr, ok := err.(*errhelp.AzureError); ok {
			if helpers.ContainsString(catch, azerr.Type) {
				return false, nil
			}
		}

		return true, fmt.Errorf("ResourceGroup delete error %v", err)

	}

	return true, nil
}

func (g *AzureResourceGroupManager) ForSubscription(context.Context, runtime.Object) error {
	return nil
}

func (g *AzureResourceGroupManager) convert(obj runtime.Object) (*azurev1alpha1.ResourceGroup, error) {
	local, ok := obj.(*azurev1alpha1.ResourceGroup)
	if !ok {
		return nil, fmt.Errorf("failed type assertion on kind: %s", obj.GetObjectKind().GroupVersionKind().String())
	}
	return local, nil
}
