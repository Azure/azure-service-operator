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
	"errors"
	"github.com/Azure/azure-sdk-for-go/services/eventhub/mgmt/2017-04-01/eventhub"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/mock_test/helpers"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/to"
)

type eventHubNamespaceResource struct {
	resourceGroupName string
	namespaceName     string
	eHNamespace       eventhub.EHNamespace
}

type mockEventHubNamespaceManager struct {
	eventHubNamespaceResources []eventHubNamespaceResource
}

func findEventHubNamespace(res []eventHubNamespaceResource, predicate func(eventHubNamespaceResource) bool) (int, eventHubNamespaceResource) {
	for index, r := range res {
		if predicate(r) {
			return index, r
		}
	}
	return -1, eventHubNamespaceResource{}
}

func (manager *mockEventHubNamespaceManager) CreateNamespaceAndWait(ctx context.Context, resourceGroupName string, namespaceName string, location string) (*eventhub.EHNamespace, error) {
	var eventHubNamespace = eventhub.EHNamespace{
		Response: helpers.GetRestResponse(201),
		EHNamespaceProperties: &eventhub.EHNamespaceProperties{
			ProvisioningState: to.StringPtr("Succeeded"),
		},
		Location: &location,
		Name:     to.StringPtr(namespaceName),
	}
	manager.eventHubNamespaceResources = append(manager.eventHubNamespaceResources, eventHubNamespaceResource{
		resourceGroupName: resourceGroupName,
		namespaceName:     namespaceName,
		eHNamespace:       eventHubNamespace,
	})
	return &eventHubNamespace, nil
}

func (manager *mockEventHubNamespaceManager) DeleteNamespace(ctx context.Context, resourceGroupName string, namespaceName string) (autorest.Response, error) {
	namespaces := manager.eventHubNamespaceResources

	index, _ := findEventHubNamespace(namespaces, func(g eventHubNamespaceResource) bool {
		return g.resourceGroupName == resourceGroupName &&
			g.namespaceName == namespaceName
	})

	if index == -1 {
		return helpers.GetRestResponse(404), errors.New("eventhub namespace not found")
	}

	manager.eventHubNamespaceResources = append(namespaces[:index], namespaces[index+1:]...)

	return helpers.GetRestResponse(200), nil
}

func (manager *mockEventHubNamespaceManager) GetNamespace(ctx context.Context, resourceGroupName string, namespaceName string) (*eventhub.EHNamespace, error) {
	groups := manager.eventHubNamespaceResources

	index, group := findEventHubNamespace(groups, func(g eventHubNamespaceResource) bool {
		return g.resourceGroupName == resourceGroupName &&
			g.namespaceName == namespaceName
	})

	if index == -1 {
		return &eventhub.EHNamespace{}, errors.New("eventhub namespace not found")
	}

	return &group.eHNamespace, nil
}
