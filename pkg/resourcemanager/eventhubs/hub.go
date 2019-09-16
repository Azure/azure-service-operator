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

type AzureEventHubManager struct {}

func getHubsClient() eventhub.EventHubsClient {
	hubClient := eventhub.NewEventHubsClient(config.SubscriptionID())
	auth, _ := iam.GetResourceManagementAuthorizer()
	hubClient.Authorizer = auth
	hubClient.AddToUserAgent(config.UserAgent())
	return hubClient
}

func (_ AzureEventHubManager) DeleteHub(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string) (result autorest.Response, err error) {
	hubClient := getHubsClient()
	return hubClient.Delete(ctx,
		resourceGroupName,
		namespaceName,
		eventHubName)

}

func (_ AzureEventHubManager) CreateHub(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, MessageRetentionInDays int32, PartitionCount int32, captureDescription *eventhub.CaptureDescription) (eventhub.Model, error) {
	hubClient := getHubsClient()

	// MessageRetentionInDays - Number of days to retain the events for this Event Hub, value should be 1 to 7 days
	if MessageRetentionInDays < 1 || MessageRetentionInDays > 7 {
		return eventhub.Model{}, fmt.Errorf("MessageRetentionInDays is invalid")
	}

	// PartitionCount - Number of partitions created for the Event Hub, allowed values are from 2 to 32 partitions.
	if PartitionCount < 2 || PartitionCount > 32 {
		return eventhub.Model{}, fmt.Errorf("PartitionCount is invalid")
	}

	properties := eventhub.Properties{
		PartitionCount:         to.Int64Ptr(int64(PartitionCount)),
		MessageRetentionInDays: to.Int64Ptr(int64(MessageRetentionInDays)),
		CaptureDescription:     captureDescription,
	}

	return hubClient.CreateOrUpdate(
		ctx,
		resourceGroupName,
		namespaceName,
		eventHubName,
		eventhub.Model{
			Properties: &properties,
		},
	)
}

func (_ AzureEventHubManager) GetHub(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string) (eventhub.Model, error) {
	hubClient := getHubsClient()
	return hubClient.Get(ctx, resourceGroupName, namespaceName, eventHubName)
}

func (_ AzureEventHubManager) CreateOrUpdateAuthorizationRule(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, authorizationRuleName string, parameters eventhub.AuthorizationRule) (result eventhub.AuthorizationRule, err error) {
	hubClient := getHubsClient()
	return hubClient.CreateOrUpdateAuthorizationRule(ctx, resourceGroupName, namespaceName, eventHubName, authorizationRuleName, parameters)
}

func (_ AzureEventHubManager) ListKeys(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, authorizationRuleName string) (result eventhub.AccessKeys, err error) {
	hubClient := getHubsClient()
	return hubClient.ListKeys(ctx, resourceGroupName, namespaceName, eventHubName, authorizationRuleName)
}
