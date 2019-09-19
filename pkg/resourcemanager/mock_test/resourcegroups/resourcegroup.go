package resourcegroups

import (
	"context"
	"errors"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/mock_test/helpers"

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
	r := resources.Group{
		Response: helpers.GetRestResponse(201),
		Location: to.StringPtr(location),
		Name:     to.StringPtr(groupName),
	}
	manager.resourceGroups = append(manager.resourceGroups, r)

	return r, nil
}

// DeleteGroup removes the resource group
func (manager *MockResourceGroupManager) DeleteGroup(ctx context.Context, groupName string) (result autorest.Response, err error) {
	groups := manager.resourceGroups
	index, _ := findResourceGroup(groups, func(g resources.Group) bool {
		return *g.Name == groupName
	})

	if index == -1 {
		return helpers.GetRestResponse(404), errors.New("resource group not found")
	}

	manager.resourceGroups = append(groups[:index], groups[index+1:]...)

	return helpers.GetRestResponse(200), nil
}

func (manager *MockResourceGroupManager) CheckExistence(ctx context.Context, groupName string) (result autorest.Response, err error) {
	groups := manager.resourceGroups
	index, _ := findResourceGroup(groups, func(g resources.Group) bool {
		return *g.Name == groupName
	})

	if index == -1 {
		return helpers.GetRestResponse(404), errors.New("resource group not found")
	}

	return helpers.GetRestResponse(204), nil
}
