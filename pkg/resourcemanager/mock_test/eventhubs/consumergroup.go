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

type consumerGroupResource struct {
	resourceGroupName string
	namespaceName     string
	eventHubName      string
	consumerGroupName string
	ConsumerGroup     eventhub.ConsumerGroup
}

type mockConsumerGroupManager struct {
	consumerGroupResources []consumerGroupResource
}

func findConsumerGroup(res []consumerGroupResource, predicate func(consumerGroupResource) bool) (int, consumerGroupResource) {
	for index, r := range res {
		if predicate(r) {
			return index, r
		}
	}
	return -1, consumerGroupResource{}
}

func find(res []interface{}, predicate func(interface{}) bool) (int, interface{}) {
	for index, r := range res {
		if predicate(r) {
			return index, r
		}
	}
	return -1, consumerGroupResource{}
}

func (manager *mockConsumerGroupManager) CreateConsumerGroup(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, consumerGroupName string) (eventhub.ConsumerGroup, error) {
	var consumerGroup = eventhub.ConsumerGroup{
		Response: helpers.GetRestResponse(201),
		Name:     to.StringPtr(consumerGroupName),
	}
	manager.consumerGroupResources = append(manager.consumerGroupResources, consumerGroupResource{
		resourceGroupName: resourceGroupName,
		namespaceName:     namespaceName,
		eventHubName:      eventHubName,
		consumerGroupName: consumerGroupName,
		ConsumerGroup:     consumerGroup,
	})
	return consumerGroup, nil
}

func (manager *mockConsumerGroupManager) DeleteConsumerGroup(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, consumerGroupName string) (autorest.Response, error) {
	groups := manager.consumerGroupResources

	index, _ := findConsumerGroup(groups, func(g consumerGroupResource) bool {
		return g.resourceGroupName == resourceGroupName &&
			g.namespaceName == namespaceName &&
			g.eventHubName == eventHubName &&
			g.consumerGroupName == consumerGroupName
	})

	if index == -1 {
		return helpers.GetRestResponse(404), errors.New("consumer group not found")
	}

	manager.consumerGroupResources = append(groups[:index], groups[index+1:]...)

	return helpers.GetRestResponse(200), nil
}

func (manager *mockConsumerGroupManager) GetConsumerGroup(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, consumerGroupName string) (eventhub.ConsumerGroup, error) {
	groups := manager.consumerGroupResources

	index, group := findConsumerGroup(groups, func(g consumerGroupResource) bool {
		return g.resourceGroupName == resourceGroupName &&
			g.namespaceName == namespaceName &&
			g.eventHubName == eventHubName &&
			g.consumerGroupName == consumerGroupName
	})

	if index == -1 {
		return eventhub.ConsumerGroup{}, errors.New("consumer group not found")
	}

	return group.ConsumerGroup, nil
}
