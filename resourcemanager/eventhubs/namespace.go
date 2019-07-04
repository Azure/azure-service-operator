// Copyright (c) Telstra and contributors.  All rights reserved.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package eventhubs

import (
	"context"

	"Telstra.Dx.AzureOperator/resourcemanager/config"
	"Telstra.Dx.AzureOperator/resourcemanager/iam"

	"github.com/Azure/azure-sdk-for-go/services/eventhub/mgmt/2017-04-01/eventhub"
	"github.com/Azure/go-autorest/autorest/to"
)

func getNamespacesClient() eventhub.NamespacesClient {
	nsClient := eventhub.NewNamespacesClient(config.SubscriptionID())
	auth, _ := iam.GetResourceManagementAuthorizer()
	nsClient.Authorizer = auth
	nsClient.AddToUserAgent(config.UserAgent())
	return nsClient
}

// DeleteNamespace deletes an existing namespace. This operation also removes all associated resources under the namespace.
// Parameters:
// resourceGroupName - name of the resource group within the azure subscription.
// namespaceName - the Namespace name
func DeleteNamespace(ctx context.Context, resourceGroupName string, namespaceName string) (result eventhub.NamespacesDeleteFuture, err error) {

	nsClient := getNamespacesClient()
	return nsClient.Delete(ctx,
		resourceGroupName,
		namespaceName)

}

// CreateNamespace creates an Event Hubs namespace
// Parameters:
// resourceGroupName - name of the resource group within the azure subscription.
// namespaceName - the Namespace name
// location - azure region
func CreateNamespace(ctx context.Context, resourceGroupName string, namespaceName string, location string) (*eventhub.EHNamespace, error) {
	nsClient := getNamespacesClient()
	future, err := nsClient.CreateOrUpdate(
		ctx,
		resourceGroupName,
		namespaceName,
		eventhub.EHNamespace{
			Location: to.StringPtr(location),
		},
	)
	if err != nil {
		return nil, err
	}

	err = future.WaitForCompletionRef(ctx, nsClient.Client)
	if err != nil {
		return nil, err
	}

	result, err := future.Result(nsClient)
	return &result, err
}
