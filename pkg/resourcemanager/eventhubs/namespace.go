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

package eventhubs

import (
	"context"

	"github.com/Azure/go-autorest/autorest"

	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/iam"

	"github.com/Azure/azure-sdk-for-go/services/eventhub/mgmt/2017-04-01/eventhub"
	"github.com/Azure/go-autorest/autorest/to"
)

type azureEventHubNamespaceManager struct{}

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
func (_ *azureEventHubNamespaceManager) DeleteNamespace(ctx context.Context, resourceGroupName string, namespaceName string) (autorest.Response, error) {

	nsClient := getNamespacesClient()
	future, err := nsClient.Delete(ctx,
		resourceGroupName,
		namespaceName)

	return autorest.Response{Response: future.Response()}, err
}

// Get gets the description of the specified namespace.
// Parameters:
// resourceGroupName - name of the resource group within the azure subscription.
// namespaceName - the Namespace name
func (_ *azureEventHubNamespaceManager) GetNamespace(ctx context.Context, resourceGroupName string, namespaceName string) (*eventhub.EHNamespace, error) {
	nsClient := getNamespacesClient()
	x, err := nsClient.Get(ctx, resourceGroupName, namespaceName)

	if err != nil {
		return &eventhub.EHNamespace{
			Response: x.Response,
		}, err
	}

	return &x, err
}

// CreateNamespaceAndWait creates an Event Hubs namespace
// Parameters:
// resourceGroupName - name of the resource group within the azure subscription.
// namespaceName - the Namespace name
// location - azure region
func (_ *azureEventHubNamespaceManager) CreateNamespaceAndWait(ctx context.Context, resourceGroupName string, namespaceName string, location string) (*eventhub.EHNamespace, error) {
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
