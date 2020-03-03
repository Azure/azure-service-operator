// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package eventhubs

type EventHubManagers struct {
	EventHubNamespace EventHubNamespaceManager
	EventHub          EventHubManager
	ConsumerGroup     ConsumerGroupManager
}

var AzureEventHubManagers = EventHubManagers{
	EventHubNamespace: &azureEventHubNamespaceManager{},
	EventHub:          &azureEventHubManager{},
	ConsumerGroup:     &azureConsumerGroupManager{},
}
