package eventhubs

import (
	"context"

	"Telstra.Dx.AzureOperator/resourcemanager/config"
	"Telstra.Dx.AzureOperator/resourcemanager/iam"

	"github.com/Azure/azure-sdk-for-go/services/eventhub/mgmt/2017-04-01/eventhub"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/to"
)

func getHubsClient() eventhub.EventHubsClient {
	hubClient := eventhub.NewEventHubsClient(config.SubscriptionID())
	auth, _ := iam.GetResourceManagementAuthorizer()
	hubClient.Authorizer = auth
	hubClient.AddToUserAgent(config.UserAgent())
	return hubClient
}

// DeleteHub deletes an Event Hub from the specified Namespace and resource group.
// Parameters:
// resourceGroupName - name of the resource group within the azure subscription.
// namespaceName - the Namespace name
// eventHubName - the Event Hub name
func DeleteHub(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string) (result autorest.Response, err error) {
	hubClient := getHubsClient()
	return hubClient.Delete(ctx,
		resourceGroupName,
		namespaceName,
		eventHubName)

}

// CreateHub creates an Event Hubs hub in a namespace
// Parameters:
// resourceGroupName - name of the resource group within the azure subscription.
// namespaceName - the Namespace name
// eventHubName - the Event Hub name
func CreateHub(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string) (eventhub.Model, error) {
	hubClient := getHubsClient()
	return hubClient.CreateOrUpdate(
		ctx,
		resourceGroupName,
		namespaceName,
		eventHubName,
		eventhub.Model{
			Properties: &eventhub.Properties{
				PartitionCount: to.Int64Ptr(4),
			},
		},
	)
}
