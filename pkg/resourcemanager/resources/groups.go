// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package resources

import (
	"context"
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2018-02-01/resources"

	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	"github.com/Azure/go-autorest/autorest/to"

	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/iam"
)

func getGroupsClient() (client resources.GroupsClient, err error) {
	client = resources.NewGroupsClient(config.SubscriptionID())
	a, err := iam.GetResourceManagementAuthorizer()
	if err != nil {
		client = resources.GroupsClient{}
	} else {
		client.Authorizer = a
		client.AddToUserAgent(config.UserAgent())
	}
	return client, err
}

func getGroupsClientWithAuthFile() (client resources.GroupsClient, err error) {
	client = resources.NewGroupsClient(config.SubscriptionID())
	// requires env var AZURE_AUTH_LOCATION set to output of
	// `az ad sp create-for-rbac --sdk-auth`
	a, err := auth.NewAuthorizerFromFile(azure.PublicCloud.ResourceManagerEndpoint)
	if err != nil {
		client = resources.GroupsClient{}
	} else {
		client.Authorizer = a
		client.AddToUserAgent(config.UserAgent())
	}
	return client, err
}

// CreateGroup creates a new resource group named by env var
func CreateGroup(ctx context.Context, groupName string) (group resources.Group, err error) {
	groupsClient, err := getGroupsClient()
	if err != nil {
		return resources.Group{}, err
	}
	log.Println(fmt.Sprintf("creating resource group '%s' on location: %v", groupName, config.Location()))
	return groupsClient.CreateOrUpdate(
		ctx,
		groupName,
		resources.Group{
			Location: to.StringPtr(config.Location()),
		})
}

// CreateGroupWithAuthFile creates a new resource group. The client authorizer
// is set up based on an auth file created using the Azure CLI.
func CreateGroupWithAuthFile(ctx context.Context, groupName string) (resources.Group, error) {
	groupsClient := getGroupsClientWithAuthFile()
	log.Println(fmt.Sprintf("creating resource group '%s' on location: %v", groupName, config.Location()))
	return groupsClient.CreateOrUpdate(
		ctx,
		groupName,
		resources.Group{
			Location: to.StringPtr(config.Location()),
		})
}

// DeleteGroup removes the resource group named by env var
func DeleteGroup(ctx context.Context, groupName string) (result resources.GroupsDeleteFuture, err error) {
	groupsClient := getGroupsClient()
	return groupsClient.Delete(ctx, groupName)
}

// ListGroups gets an interator that gets all resource groups in the subscription
func ListGroups(ctx context.Context) (resources.GroupListResultIterator, error) {
	groupsClient := getGroupsClient()
	return groupsClient.ListComplete(ctx, "", nil)
}

// GetGroup gets info on the resource group in use
func GetGroup(ctx context.Context) (resources.Group, error) {
	groupsClient := getGroupsClient()
	return groupsClient.Get(ctx, config.GroupName())
}

// DeleteAllGroupsWithPrefix deletes all rescource groups that start with a certain prefix
func DeleteAllGroupsWithPrefix(ctx context.Context, prefix string) (futures []resources.GroupsDeleteFuture, groups []string) {
	if config.KeepResources() {
		log.Println("keeping resource groups")
		return
	}
	for list, _ := ListGroups(ctx); list.NotDone(); {
		rgName := *list.Value().Name
		if strings.HasPrefix(rgName, prefix) {
			fmt.Printf("deleting group '%s'\n", rgName)
			future, _ := DeleteGroup(ctx, rgName)
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
			if err == nil {
				fmt.Printf("finished deleting group '%s'\n", rg)
			}
			wg.Done()
		}(ctx, f, groups[i])
	}
}
