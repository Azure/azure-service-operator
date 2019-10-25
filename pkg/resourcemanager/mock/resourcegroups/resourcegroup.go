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
	"errors"
	"fmt"
	"net/http"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"

	"github.com/Azure/azure-service-operator/pkg/resourcemanager/mock/helpers"
	"k8s.io/apimachinery/pkg/runtime"

	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2017-05-10/resources"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/to"
)

type MockResourceGroupManager struct {
	resourceGroups []resources.Group
}

func findResourceGroup(res []resources.Group, predicate func(resources.Group) bool) (int, resources.Group) {
	for index, r := range res {
		if predicate(r) {
			return index, r
		}
	}
	return -1, resources.Group{}
}

// CreateGroup creates a new resource group
func (manager *MockResourceGroupManager) CreateGroup(ctx context.Context, groupName string, location string) (resources.Group, error) {
	index, _ := findResourceGroup(manager.resourceGroups, func(g resources.Group) bool {
		return *g.Name == groupName
	})

	r := resources.Group{
		Response: helpers.GetRestResponse(201),
		Location: to.StringPtr(location),
		Name:     to.StringPtr(groupName),
	}

	if index == -1 {
		manager.resourceGroups = append(manager.resourceGroups, r)
	}

	return r, nil
}

// DeleteGroup removes the resource group
func (manager *MockResourceGroupManager) DeleteGroup(ctx context.Context, groupName string) (autorest.Response, error) {
	groups := manager.resourceGroups

	index, _ := findResourceGroup(groups, func(g resources.Group) bool {
		return *g.Name == groupName
	})

	if index == -1 {
		return helpers.GetRestResponse(http.StatusNotFound), errors.New("resource group not found")
	}

	manager.resourceGroups = append(groups[:index], groups[index+1:]...)

	return helpers.GetRestResponse(http.StatusOK), nil
}

func (manager *MockResourceGroupManager) DeleteGroupAsync(ctx context.Context, groupName string) (resources.GroupsDeleteFuture, error) {
	_, err := manager.DeleteGroup(ctx, groupName)

	return resources.GroupsDeleteFuture{}, err
}

func (manager *MockResourceGroupManager) CheckExistence(ctx context.Context, groupName string) (autorest.Response, error) {
	groups := manager.resourceGroups
	index, _ := findResourceGroup(groups, func(g resources.Group) bool {
		return *g.Name == groupName
	})

	if index == -1 {
		return helpers.GetRestResponse(http.StatusNotFound), errors.New("resource group not found")
	}

	return helpers.GetRestResponse(http.StatusNoContent), nil
}

func (g *MockResourceGroupManager) Ensure(ctx context.Context, obj runtime.Object) (bool, error) {
	instance, err := g.convert(obj)
	if err != nil {
		return false, err
	}

	_, err = g.CreateGroup(ctx, instance.ObjectMeta.Name, instance.Spec.Location)
	if err != nil {
		return false, err
	}

	instance.Status.Provisioning = true
	instance.Status.Provisioned = true

	return true, nil
}

func (g *MockResourceGroupManager) Delete(ctx context.Context, obj runtime.Object) (bool, error) {
	instance, err := g.convert(obj)
	if err != nil {
		return false, err
	}

	_, err = g.DeleteGroup(ctx, instance.ObjectMeta.Name)
	if err != nil {
		return false, err
	}

	return false, nil
}

func (g *MockResourceGroupManager) convert(obj runtime.Object) (*azurev1alpha1.ResourceGroup, error) {
	local, ok := obj.(*azurev1alpha1.ResourceGroup)
	if !ok {
		return nil, fmt.Errorf("failed type assertion on kind: %s", obj.GetObjectKind().GroupVersionKind().String())
	}
	return local, nil
}
