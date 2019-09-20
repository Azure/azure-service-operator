package eventhubs

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/services/eventhub/mgmt/2017-04-01/eventhub"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/mock_test/helpers"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/to"
)

type ConsumerGroupResource struct {
	ResourceGroupName string
	NamespaceName     string
	EventHubName      string
	ConsumerGroupName string
	ConsumerGroup     eventhub.ConsumerGroup
}

type mockConsumerGroupManager struct {
	consumerGroupResources []ConsumerGroupResource
}

func findConsumerGroup(res []ConsumerGroupResource, predicate func(ConsumerGroupResource) bool) (int, ConsumerGroupResource) {
	for index, r := range res {
		if predicate(r) {
			return index, r
		}
	}
	return -1, ConsumerGroupResource{}
}

func find(res []interface{}, predicate func(interface{}) bool) (int, interface{}) {
	for index, r := range res {
		if predicate(r) {
			return index, r
		}
	}
	return -1, ConsumerGroupResource{}
}

func (manager *mockConsumerGroupManager) CreateConsumerGroup(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, consumerGroupName string) (eventhub.ConsumerGroup, error) {
	var consumerGroup = eventhub.ConsumerGroup{
		Response: helpers.GetRestResponse(201),
		Name:     to.StringPtr(consumerGroupName),
	}
	manager.consumerGroupResources = append(manager.consumerGroupResources, ConsumerGroupResource{
		ResourceGroupName: resourceGroupName,
		NamespaceName:     namespaceName,
		EventHubName:      eventHubName,
		ConsumerGroupName: consumerGroupName,
		ConsumerGroup:     consumerGroup,
	})
	return consumerGroup, nil
}

func (manager *mockConsumerGroupManager) DeleteConsumerGroup(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, consumerGroupName string) (autorest.Response, error) {
	groups := manager.consumerGroupResources

	index, _ := findConsumerGroup(groups, func(g ConsumerGroupResource) bool {
		return g.ResourceGroupName == resourceGroupName &&
			g.NamespaceName == namespaceName &&
			g.EventHubName == eventHubName &&
			g.ConsumerGroupName == consumerGroupName
	})

	if index == -1 {
		return helpers.GetRestResponse(404), errors.New("consumer group not found")
	}

	manager.consumerGroupResources = append(groups[:index], groups[index+1:]...)

	return helpers.GetRestResponse(200), nil
}

func (manager *mockConsumerGroupManager) GetConsumerGroup(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, consumerGroupName string) (eventhub.ConsumerGroup, error) {
	groups := manager.consumerGroupResources

	index, group := findConsumerGroup(groups, func(g ConsumerGroupResource) bool {
		return g.ResourceGroupName == resourceGroupName &&
			g.NamespaceName == namespaceName &&
			g.EventHubName == eventHubName &&
			g.ConsumerGroupName == consumerGroupName
	})

	if index == -1 {
		return eventhub.ConsumerGroup{}, errors.New("consumer group not found")
	}

	return group.ConsumerGroup, nil
}
