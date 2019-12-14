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
	"strings"
	"sync"

	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2017-05-10/resources"

	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/iam"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	"github.com/Azure/go-autorest/autorest/to"
)

// AzureResourceGroupManager is the struct which contains helper functions for resource groups
type AzureResourceGroupManager struct{}

func getGroupsClient() resources.GroupsClient {
	groupsClient := resources.NewGroupsClient(config.SubscriptionID())
	a, err := iam.GetResourceManagementAuthorizer()
	if err != nil {
		return resources.GroupsClient{}
	}
	groupsClient.Authorizer = a
	groupsClient.AddToUserAgent(config.UserAgent())
	return groupsClient
}

func getGroupsClientWithAuthFile() resources.GroupsClient {
	groupsClient := resources.NewGroupsClient(config.SubscriptionID())
	// requires env var AZURE_AUTH_LOCATION set to output of
	// `az ad sp create-for-rbac --sdk-auth`
	a, err := auth.NewAuthorizerFromFile(azure.PublicCloud.ResourceManagerEndpoint)
	if err != nil {
		return resources.GroupsClient{}
	}
	groupsClient.Authorizer = a
	groupsClient.AddToUserAgent(config.UserAgent())
	return groupsClient
}

// CreateGroup creates a new resource group named by env var
func (_ *AzureResourceGroupManager) CreateGroup(ctx context.Context, groupName string, location string) (resources.Group, error) {
	groupsClient := getGroupsClient()

	return groupsClient.CreateOrUpdate(
		ctx,
		groupName,
		resources.Group{
			Location: to.StringPtr(location),
		})
}

// CreateGroupWithAuthFile creates a new resource group. The client authorizer
// is set up based on an auth file created using the Azure CLI.
func CreateGroupWithAuthFile(ctx context.Context, groupName string, location string) (resources.Group, error) {
	groupsClient := getGroupsClientWithAuthFile()
	return groupsClient.CreateOrUpdate(
		ctx,
		groupName,
		resources.Group{
			Location: to.StringPtr(location),
		})
}

// DeleteGroup removes the resource group named by env var
func (_ *AzureResourceGroupManager) DeleteGroup(ctx context.Context, groupName string) (result autorest.Response, err error) {
	var client = getGroupsClient()

	future, err := client.Delete(ctx, groupName)
	if err != nil {
		return autorest.Response{}, err
	}

	future.WaitForCompletionRef(ctx, client.Client)

	return future.Result(client)
}

func (_ *AzureResourceGroupManager) DeleteGroupAsync(ctx context.Context, groupName string) (result resources.GroupsDeleteFuture, err error) {
	return deleteGroupAsync(ctx, groupName)
}

func deleteGroupAsync(ctx context.Context, groupName string) (result resources.GroupsDeleteFuture, err error) {
	groupsClient := getGroupsClient()
	return groupsClient.Delete(ctx, groupName)
}

// ListGroups gets an interator that gets all resource groups in the subscription
func ListGroups(ctx context.Context) (resources.GroupListResultIterator, error) {
	groupsClient := getGroupsClient()
	return groupsClient.ListComplete(ctx, "", nil)
}

// GetGroup gets info on the resource group in use
func GetGroup(ctx context.Context, groupName string) (resources.Group, error) {
	groupsClient := getGroupsClient()
	return groupsClient.Get(ctx, groupName)
}

// DeleteAllGroupsWithPrefix deletes all rescource groups that start with a certain prefix
func DeleteAllGroupsWithPrefix(ctx context.Context, prefix string) (futures []resources.GroupsDeleteFuture, groups []string) {

	for list, err := ListGroups(ctx); list.NotDone(); err = list.Next() {
		if err != nil {
			return
		}
		rgName := *list.Value().Name
		if strings.HasPrefix(rgName, prefix) {
			future, err := deleteGroupAsync(ctx, rgName)
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
func WaitForDeleteCompletion(ctx context.Context, wg *sync.WaitGroup, futures []resources.GroupsDeleteFuture, groups []string) {
	for i, f := range futures {
		wg.Add(1)
		go func(ctx context.Context, future resources.GroupsDeleteFuture, rg string) {
			err := future.WaitForCompletionRef(ctx, getGroupsClient().Client)
			if err != nil {
				return
			}
			wg.Done()
		}(ctx, f, groups[i])
	}
}

// CheckExistence checks whether a resource exists
func (_ *AzureResourceGroupManager) CheckExistence(ctx context.Context, resourceGroupName string) (result autorest.Response, err error) {
	groupsClient := getGroupsClient()
	result, _ = groupsClient.CheckExistence(ctx, resourceGroupName)
	return
}
