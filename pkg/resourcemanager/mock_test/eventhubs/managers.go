package eventhubs

import . "github.com/Azure/azure-service-operator/pkg/resourcemanager/eventhubs"

var MockEventHubManagers = EventHubManagers{
	EventHubNamespace: &mockEventHubNamespaceManager{},
	EventHub:          &mockEventHubManager{},
	ConsumerGroup:     &mockConsumerGroupManager{},
}
