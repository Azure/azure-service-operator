package eventhubs

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/services/eventhub/mgmt/2017-04-01/eventhub"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/mock_test/helpers"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/to"
)

type EventHubNamespaceResource struct {
	ResourceGroupName string
	NamespaceName     string
	EHNamespace       eventhub.EHNamespace
}

type mockEventHubNamespaceManager struct {
	eventHubNamespaceResources []EventHubNamespaceResource
}

func findEventHubNamespace(res []EventHubNamespaceResource, predicate func(EventHubNamespaceResource) bool) (int, EventHubNamespaceResource) {
	for index, r := range res {
		if predicate(r) {
			return index, r
		}
	}
	return -1, EventHubNamespaceResource{}
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
	manager.eventHubNamespaceResources = append(manager.eventHubNamespaceResources, EventHubNamespaceResource{
		ResourceGroupName: resourceGroupName,
		NamespaceName:     namespaceName,
		EHNamespace:       eventHubNamespace,
	})
	return &eventHubNamespace, nil
}

func (manager *mockEventHubNamespaceManager) DeleteNamespace(ctx context.Context, resourceGroupName string, namespaceName string) (result autorest.Response, err error) {
	namespaces := manager.eventHubNamespaceResources

	index, _ := findEventHubNamespace(namespaces, func(g EventHubNamespaceResource) bool {
		return g.ResourceGroupName == resourceGroupName &&
			g.NamespaceName == namespaceName
	})

	if index == -1 {
		return helpers.GetRestResponse(404), errors.New("eventhub namespace not found")
	}

	manager.eventHubNamespaceResources = append(namespaces[:index], namespaces[index+1:]...)

	return helpers.GetRestResponse(200), nil
}

func (manager *mockEventHubNamespaceManager) GetNamespace(ctx context.Context, resourceGroupName string, namespaceName string) (*eventhub.EHNamespace, error) {
	groups := manager.eventHubNamespaceResources

	index, group := findEventHubNamespace(groups, func(g EventHubNamespaceResource) bool {
		return g.ResourceGroupName == resourceGroupName &&
			g.NamespaceName == namespaceName
	})

	if index == -1 {
		return &eventhub.EHNamespace{}, errors.New("eventhub namespace not found")
	}

	return &group.EHNamespace, nil
}
