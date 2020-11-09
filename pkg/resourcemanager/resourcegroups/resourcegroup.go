// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package resourcegroups

import (
	"context"
	"net/http"
	"strings"
	"sync"

	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2017-05-10/resources"

	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/iam"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	"github.com/Azure/go-autorest/autorest/to"
)

// AzureResourceGroupManager is the struct which contains helper functions for resource groups
type AzureResourceGroupManager struct {
	creds config.Credentials
}

func getGroupsClient(creds config.Credentials) (resources.GroupsClient, error) {
	groupsClient := resources.NewGroupsClientWithBaseURI(config.BaseURI(), creds.SubscriptionID())
	a, err := iam.GetResourceManagementAuthorizer(creds)
	if err != nil {
		return resources.GroupsClient{}, err
	}
	groupsClient.Authorizer = a
	groupsClient.AddToUserAgent(config.UserAgent())
	return groupsClient, nil
}

func getGroupsClientWithAuthFile(creds config.Credentials) (resources.GroupsClient, error) {
	groupsClient := resources.NewGroupsClientWithBaseURI(config.BaseURI(), creds.SubscriptionID())
	// requires env var AZURE_AUTH_LOCATION set to output of
	// `az ad sp create-for-rbac --sdk-auth`
	a, err := auth.NewAuthorizerFromFile(config.BaseURI())
	if err != nil {
		return resources.GroupsClient{}, err
	}
	groupsClient.Authorizer = a
	groupsClient.AddToUserAgent(config.UserAgent())
	return groupsClient, nil
}

// CreateGroup creates a new resource group named by env var
func (m *AzureResourceGroupManager) CreateGroup(ctx context.Context, groupName string, location string) (resources.Group, error) {
	groupsClient, err := getGroupsClient(m.creds)
	if err != nil {
		return resources.Group{}, err
	}

	return groupsClient.CreateOrUpdate(
		ctx,
		groupName,
		resources.Group{
			Location: to.StringPtr(location),
		})
}

// CreateGroupWithAuthFile creates a new resource group. The client authorizer
// is set up based on an auth file created using the Azure CLI.
func (m *AzureResourceGroupManager) CreateGroupWithAuthFile(ctx context.Context, groupName, location string) (resources.Group, error) {
	groupsClient, err := getGroupsClientWithAuthFile(m.creds)
	if err != nil {
		return resources.Group{}, err
	}

	return groupsClient.CreateOrUpdate(
		ctx,
		groupName,
		resources.Group{
			Location: to.StringPtr(location),
		})
}

// DeleteGroup removes the resource group named by env var
func (m *AzureResourceGroupManager) DeleteGroup(ctx context.Context, groupName string) (result autorest.Response, err error) {
	client, err := getGroupsClient(m.creds)
	if err != nil {
		return autorest.Response{
			Response: &http.Response{
				StatusCode: 500,
			},
		}, err
	}

	future, err := client.Delete(ctx, groupName)
	if err != nil {
		return autorest.Response{}, err
	}

	return future.Result(client)
}

func (m *AzureResourceGroupManager) DeleteGroupAsync(ctx context.Context, groupName string) (result resources.GroupsDeleteFuture, err error) {
	return m.deleteGroupAsync(ctx, groupName)
}

func (m *AzureResourceGroupManager) deleteGroupAsync(ctx context.Context, groupName string) (result resources.GroupsDeleteFuture, err error) {
	groupsClient, err := getGroupsClient(m.creds)
	if err != nil {
		return resources.GroupsDeleteFuture{}, err
	}

	return groupsClient.Delete(ctx, groupName)
}

// ListGroups gets an interator that gets all resource groups in the subscription
func (m *AzureResourceGroupManager) ListGroups(ctx context.Context) (resources.GroupListResultIterator, error) {
	groupsClient, err := getGroupsClient(m.creds)
	if err != nil {
		return resources.GroupListResultIterator{}, err
	}

	return groupsClient.ListComplete(ctx, "", nil)
}

// GetGroup gets info on the resource group in use
func (m *AzureResourceGroupManager) GetGroup(ctx context.Context, groupName string) (resources.Group, error) {
	groupsClient, err := getGroupsClient(m.creds)
	if err != nil {
		return resources.Group{}, err
	}

	return groupsClient.Get(ctx, groupName)
}

// DeleteAllGroupsWithPrefix deletes all rescource groups that start with a certain prefix
func (m *AzureResourceGroupManager) DeleteAllGroupsWithPrefix(ctx context.Context, prefix string) (futures []resources.GroupsDeleteFuture, groups []string) {

	for list, err := m.ListGroups(ctx); list.NotDone(); err = list.Next() {
		if err != nil {
			return
		}
		rgName := *list.Value().Name
		if strings.HasPrefix(rgName, prefix) {
			future, err := m.deleteGroupAsync(ctx, rgName)
			if err != nil {
				return
			}
			futures = append(futures, future)
			groups = append(groups, rgName)
		}
	}
	return
}

// WaitForDeleteCompletion concurrently waits for delete group operations to finish
func (m *AzureResourceGroupManager) WaitForDeleteCompletion(ctx context.Context, wg *sync.WaitGroup, futures []resources.GroupsDeleteFuture, groups []string) {
	for i, f := range futures {
		wg.Add(1)
		go func(ctx context.Context, future resources.GroupsDeleteFuture, rg string) {
			client, err := getGroupsClient(m.creds)
			if err != nil {
				return
			}

			err = future.WaitForCompletionRef(ctx, client.Client)
			if err != nil {
				return
			}
			wg.Done()
		}(ctx, f, groups[i])
	}
}

// CheckExistence checks whether a resource exists
func (m *AzureResourceGroupManager) CheckExistence(ctx context.Context, resourceGroupName string) (result autorest.Response, err error) {
	groupsClient, err := getGroupsClient(m.creds)
	if err != nil {
		return autorest.Response{
			Response: &http.Response{
				StatusCode: 500,
			},
		}, err
	}

	result, _ = groupsClient.CheckExistence(ctx, resourceGroupName)
	return
}
