// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package eventhubs

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/eventhub/mgmt/2017-04-01/eventhub"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	"github.com/Azure/go-autorest/autorest"
)

type EventHubManager interface {
	// DeleteHub deletes an Event Hub from the specified Namespace and resource group.
	// Parameters:
	// resourceGroupName - name of the resource group within the azure subscription.
	// namespaceName - the Namespace name
	// eventHubName - the Event Hub name
	DeleteHub(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string) (result autorest.Response, err error)

	// CreateHub creates an Event Hubs hub in a namespace
	// Parameters:
	// resourceGroupName - name of the resource group within the azure subscription.
	// namespaceName - the Namespace name
	// eventHubName - the Event Hub name
	CreateHub(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, MessageRetentionInDays int32, PartitionCount int32, captureDescription *eventhub.CaptureDescription) (eventhub.Model, error)

	//GetHub gets an Event Hubs description for the specified Event Hub.
	// Parameters:
	// resourceGroupName - name of the resource group within the azure subscription.
	// namespaceName - the Namespace name
	// eventHubName - the Event Hub name
	GetHub(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string) (eventhub.Model, error)

	// CreateOrUpdateAuthorizationRule creates or updates an AuthorizationRule for the specified Event Hub.
	// Parameters:
	// resourceGroupName - name of the resource group within the azure subscription.
	// namespaceName - the Namespace name
	// eventHubName - the Event Hub name
	// authorizationRuleName - the authorization rule name.
	// parameters - the shared access AuthorizationRule.
	CreateOrUpdateAuthorizationRule(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, authorizationRuleName string, parameters eventhub.AuthorizationRule) (result eventhub.AuthorizationRule, err error)

	// ListKeys gets the ACS and SAS connection strings for the Event Hub.
	// Parameters:
	// resourceGroupName - name of the resource group within the azure subscription.
	// namespaceName - the Namespace name
	// eventHubName - the Event Hub name
	// authorizationRuleName - the authorization rule name.
	ListKeys(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, authorizationRuleName string) (result eventhub.AccessKeys, err error)

	// also embed arm client methods
	resourcemanager.ARMClient
}
