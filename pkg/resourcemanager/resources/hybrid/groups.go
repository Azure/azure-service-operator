// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package resources

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/resources/mgmt/resources"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/to"

	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/iam"
)

const (
	errorPrefix = "Cannot create resource group, reason: %v"
)

func getGroupsClient(tokenAudience string) resources.GroupsClient {
	var tokenAuthorizer, err := iam.getAuthorizerForResource(tokenAudience)
	if err != nil {
		log.Fatalf("failed to get token: %v\n", err)
	}

	groupsClient := resources.NewGroupsClientWithBaseURI(
		config.Environment().ResourceManagerEndpoint,
		config.SubscriptionID())
	groupsClient.Authorizer = tokenAuthorizer
	groupsClient.AddToUserAgent(config.UserAgent())

	return groupsClient
}

// CreateGroup creates a new resource group named by env var
func CreateGroup(ctx context.Context) (resources.Group, error) {
	groupClient := getGroupsClient(config.Environment().TokenAudience)

	return groupClient.CreateOrUpdate(ctx,
		config.GroupName(),
		resources.Group{
			Location: to.StringPtr(config.Location()),
		},
	)
}

// DeleteGroup removes the resource group named by env var
func DeleteGroup(ctx context.Context) (result resources.GroupsDeleteFuture, err error) {
	groupsClient := getGroupsClient(config.Environment().TokenAudience)

	return groupsClient.Delete(ctx, config.GroupName())
}
