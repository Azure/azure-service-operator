package eventhubs

type EventHubManagers struct {
	EventHubNamespace  		EventHubNamespaceManager
	EventHub 				EventHubManager
	ConsumerGroup  			ConsumerGroupManager
}

var AzureEventHubManagers = EventHubManagers{
	EventHubNamespace: AzureEventHubNamespaceManager{},
	EventHub: AzureEventHubManager{},
	ConsumerGroup: AzureConsumerGroupManager{},
}
