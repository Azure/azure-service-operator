package eventhubs

import (
	"context"
	"fmt"

	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/iam"

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
func CreateHub(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, MessageRetentionInDays int32, PartitionCount int32) (eventhub.Model, error) {
	hubClient := getHubsClient()

	// MessageRetentionInDays - Number of days to retain the events for this Event Hub, value should be 1 to 7 days
	if MessageRetentionInDays < 1 || MessageRetentionInDays > 7 {
		return eventhub.Model{}, fmt.Errorf("MessageRetentionInDays is invalid")
	}

	// PartitionCount - Number of partitions created for the Event Hub, allowed values are from 1 to 32 partitions.
	if PartitionCount < 1 || PartitionCount > 32 {
		return eventhub.Model{}, fmt.Errorf("PartitionCount is invalid")
	}

	return hubClient.CreateOrUpdate(
		ctx,
		resourceGroupName,
		namespaceName,
		eventHubName,
		eventhub.Model{
			Properties: &eventhub.Properties{
				PartitionCount:         to.Int64Ptr(int64(PartitionCount)),
				MessageRetentionInDays: to.Int64Ptr(int64(MessageRetentionInDays)),
			},
		},
	)
}

//GetHub gets an Event Hubs description for the specified Event Hub.
func GetHub(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string) (eventhub.Model, error) {
	hubClient := getHubsClient()
	return hubClient.Get(ctx, resourceGroupName, namespaceName, eventHubName)
}

// CreateOrUpdateAuthorizationRule creates or updates an AuthorizationRule for the specified Event Hub.
// Parameters:
// resourceGroupName - name of the resource group within the azure subscription.
// namespaceName - the Namespace name
// eventHubName - the Event Hub name
// authorizationRuleName - the authorization rule name.
// parameters - the shared access AuthorizationRule.
func CreateOrUpdateAuthorizationRule(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, authorizationRuleName string, parameters eventhub.AuthorizationRule) (result eventhub.AuthorizationRule, err error) {
	hubClient := getHubsClient()
	return hubClient.CreateOrUpdateAuthorizationRule(ctx, resourceGroupName, namespaceName, eventHubName, authorizationRuleName, parameters)
}

// ListKeys gets the ACS and SAS connection strings for the Event Hub.
// Parameters:
// resourceGroupName - name of the resource group within the azure subscription.
// namespaceName - the Namespace name
// eventHubName - the Event Hub name
// authorizationRuleName - the authorization rule name.
func ListKeys(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, authorizationRuleName string) (result eventhub.AccessKeys, err error) {
	hubClient := getHubsClient()
	return hubClient.ListKeys(ctx, resourceGroupName, namespaceName, eventHubName, authorizationRuleName)
}
