// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package eventhubs

import . "github.com/Azure/azure-service-operator/pkg/resourcemanager/eventhubs"

var MockEventHubManagers = EventHubManagers{
	EventHubNamespace: &mockEventHubNamespaceManager{},
	EventHub:          &mockEventHubManager{},
	ConsumerGroup:     &mockConsumerGroupManager{},
}
