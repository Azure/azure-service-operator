// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package eventhubs

import (
	"context"

	"azure-operator/aztestcreator/config"
	"azure-operator/aztestcreator/iam"

	"github.com/Azure/azure-sdk-for-go/services/eventhub/mgmt/2017-04-01/eventhub"
	"github.com/Azure/go-autorest/autorest/to"
)

func getHubsClient() eventhub.EventHubsClient {
	hubClient := eventhub.NewEventHubsClient(config.SubscriptionID())
	auth, _ := iam.GetResourceManagementAuthorizer()
	hubClient.Authorizer = auth
	hubClient.AddToUserAgent(config.UserAgent())
	return hubClient
}

// CreateHub creates an Event Hubs hub in a namespace
func CreateHub(ctx context.Context, groupname string, nsName string, hubName string) (eventhub.Model, error) {
	hubClient := getHubsClient()
	return hubClient.CreateOrUpdate(
		ctx,
		groupname,
		nsName,
		hubName,
		eventhub.Model{
			Properties: &eventhub.Properties{
				PartitionCount: to.Int64Ptr(4),
			},
		},
	)
}
