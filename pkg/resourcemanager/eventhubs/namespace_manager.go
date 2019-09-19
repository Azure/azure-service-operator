package eventhubs

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/eventhub/mgmt/2017-04-01/eventhub"
	"github.com/Azure/go-autorest/autorest"
)

type EventHubNamespaceManager interface {
	// DeleteNamespace deletes an existing namespace. This operation also removes all associated resources under the namespace.
	// Parameters:
	// resourceGroupName - name of the resource group within the azure subscription.
	// namespaceName - the Namespace name

	DeleteNamespace(ctx context.Context, resourceGroupName string, namespaceName string) (result autorest.Response, err error)

	// Get gets the description of the specified namespace.
	// Parameters:
	// resourceGroupName - name of the resource group within the azure subscription.
	// namespaceName - the Namespace name
	GetNamespace(ctx context.Context, resourceGroupName string, namespaceName string) (result *eventhub.EHNamespace, err error)

	// CreateNamespaceAndWait creates an Event Hubs namespace
	// Parameters:
	// resourceGroupName - name of the resource group within the azure subscription.
	// namespaceName - the Namespace name
	// location - azure region
	CreateNamespaceAndWait(ctx context.Context, resourceGroupName string, namespaceName string, location string) (*eventhub.EHNamespace, error)
}
