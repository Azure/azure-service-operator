package resourcemanager

import "github.com/Azure/azure-service-operator/pkg/resourcemanager/eventhubs"

type ResourceManagers struct {
	EventHubManagers eventhubs.EventHubManagers
}

var AzureResourceManagers = ResourceManagers{
	EventHubManagers: eventhubs.AzureEventHubManagers,
}
