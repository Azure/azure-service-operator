// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package eventhubs

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/eventhub/mgmt/2017-04-01/eventhub"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	"github.com/Azure/go-autorest/autorest"
)

type ConsumerGroupManager interface {
	// CreateConsumerGroup creates an Event Hub Consumer Group
	// Parameters:
	// resourceGroupName - name of the resource group within the azure subscription.
	// namespaceName - the Namespace name
	// eventHubName - the Event Hub name
	// consumerGroupName - the consumer group name
	// parameters - parameters supplied to create or update a consumer group resource.
	CreateConsumerGroup(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, consumerGroupName string) (eventhub.ConsumerGroup, error)

	//GetConsumerGroup gets consumer group description for the specified Consumer Group.
	// Parameters:
	// resourceGroupName - name of the resource group within the azure subscription.
	// namespaceName - the Namespace name
	// eventHubName - the Event Hub name
	// consumerGroupName - the consumer group name
	GetConsumerGroup(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, consumerGroupName string) (eventhub.ConsumerGroup, error)

	// DeleteConsumerGroup deletes an Event Hub Consumer Group
	// Parameters:
	// resourceGroupName - name of the resource group within the azure subscription.
	// namespaceName - the Namespace name
	// eventHubName - the Event Hub name
	// consumerGroupName - the consumer group name
	DeleteConsumerGroup(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, consumerGroupName string) (result autorest.Response, err error)

	// also embed async client methods
	resourcemanager.ARMClient
}
