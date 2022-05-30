/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	eventgrid "github.com/Azure/azure-service-operator/v2/api/eventgrid/v1beta20200601"
	storage "github.com/Azure/azure-service-operator/v2/api/storage/v1beta20210401"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
)

func Test_EventGrid_Domain(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	publicNetworkAccess := eventgrid.DomainProperties_PublicNetworkAccess_Enabled

	// Create a domain
	domain := &eventgrid.Domain{
		ObjectMeta: tc.MakeObjectMeta("domain"),
		Spec: eventgrid.Domain_Spec{
			Location:            tc.AzureRegion,
			Owner:               testcommon.AsOwner(rg),
			PublicNetworkAccess: &publicNetworkAccess,
		},
	}

	// Create a storage account to use as destination
	accessTier := storage.StorageAccountPropertiesCreateParameters_AccessTier_Hot
	kind := storage.StorageAccount_Spec_Kind_StorageV2
	sku := storage.SkuName_Standard_LRS
	acctName := tc.NoSpaceNamer.GenerateName("dest")
	acct := &storage.StorageAccount{
		ObjectMeta: tc.MakeObjectMetaWithName(acctName),
		Spec: storage.StorageAccount_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			Kind:     &kind,
			Sku:      &storage.Sku{Name: &sku},
			// TODO: They mark this property as optional but actually it is required
			AccessTier: &accessTier,
		},
	}

	tc.CreateResourcesAndWait(domain, acct)

	queueServices := &storage.StorageAccountsQueueService{
		ObjectMeta: tc.MakeObjectMeta("dest-queues"),
		Spec: storage.StorageAccountsQueueService_Spec{
			Owner: testcommon.AsOwner(acct),
		},
	}

	tc.CreateResourceAndWait(queueServices)

	queue := &storage.StorageAccountsQueueServicesQueue{
		ObjectMeta: tc.MakeObjectMeta("dest-queue"),
		Spec: storage.StorageAccountsQueueServicesQueue_Spec{
			Owner: testcommon.AsOwner(queueServices),
		},
	}

	tc.CreateResourceAndWait(queue)

	armId := *domain.Status.Id

	// TODO: disabled pending (evildiscriminator)
	/*
		acctReference := tc.MakeReferenceFromResource(acct)
		tc.RunParallelSubtests(
			testcommon.Subtest{
				Name: "CreateDomainTopicAndSubscription",
				Test: func(tc *testcommon.KubePerTestContext) {
					topic := &eventgrid.DomainsTopic{
						ObjectMeta: tc.MakeObjectMeta("topic"),
						Spec: eventgrid.DomainsTopic_Spec{
							Owner: testcommon.AsOwner(domain),
						},
					}

					tc.CreateResourceAndWait(topic)
					// don’t bother deleting; deleting domain will clean up

					endpointType := eventgrid.EventSubscriptionDestination_EndpointType_StorageQueue
					subscription := &eventgrid.EventSubscription{
						ObjectMeta: tc.MakeObjectMeta("sub"),
						Spec: eventgrid.EventSubscription_Spec{
							Owner: tc.AsExtensionOwner(topic),
							Destination: &eventgrid.EventSubscriptionDestination{
								StorageQueue: &eventgrid.StorageQueueEventSubscriptionDestination{
									EndpointType: &endpointType,
									Properties: &eventgrid.StorageQueueEventSubscriptionDestinationProperties{
										ResourceReference: acctReference,
										QueueName:         &queue.Name,
									},
								},
							},
						},
					}

					tc.CreateResourceAndWait(subscription)
					// don’t bother deleting
				},
			},
			testcommon.Subtest{
				Name: "CreateDomainSubscription",
				Test: func(tc *testcommon.KubePerTestContext) {
					endpointType := eventgrid.EventSubscriptionDestination_EndpointType_StorageQueue
					subscription := &eventgrid.EventSubscription{
						ObjectMeta: tc.MakeObjectMeta("sub"),
						Spec: eventgrid.EventSubscription_Spec{
							Owner: tc.AsExtensionOwner(domain),
							Destination: &eventgrid.EventSubscriptionDestination{
								StorageQueue: &eventgrid.StorageQueueEventSubscriptionDestination{
									EndpointType: &endpointType,
									Properties: &eventgrid.StorageQueueEventSubscriptionDestinationProperties{
										ResourceReference: acctReference,
										QueueName:         &queue.Name,
									},
								},
							},
						},
					}

					tc.CreateResourceAndWait(subscription)
					// don’t bother deleting
				},
			},
		)
	*/

	tc.DeleteResourceAndWait(domain)

	// Ensure that the resource group was really deleted in Azure
	exists, _, err := tc.AzureClient.HeadByID(
		tc.Ctx,
		armId,
		string(eventgrid.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeFalse())
}
