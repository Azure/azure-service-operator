/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	eventgrid "github.com/Azure/azure-service-operator/v2/api/eventgrid/v1alpha1api20200601"
	storage "github.com/Azure/azure-service-operator/v2/api/storage/v1alpha1api20210401"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
)

func Test_EventGrid_Domain(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	publicNetworkAccess := eventgrid.DomainPropertiesPublicNetworkAccessEnabled

	// Create a domain
	domain := &eventgrid.Domain{
		ObjectMeta: tc.MakeObjectMeta("domain"),
		Spec: eventgrid.Domains_Spec{
			Location:            tc.AzureRegion,
			Owner:               testcommon.AsOwner(rg),
			PublicNetworkAccess: &publicNetworkAccess,
		},
	}

	// Create a storage account to use as destination
	accessTier := storage.StorageAccountPropertiesCreateParametersAccessTierHot
	kind := storage.StorageAccountsSpecKindStorageV2
	sku := storage.SkuNameStandardLRS
	acctName := tc.NoSpaceNamer.GenerateName("dest")
	acct := &storage.StorageAccount{
		ObjectMeta: tc.MakeObjectMetaWithName(acctName),
		Spec: storage.StorageAccounts_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			Kind:     &kind,
			Sku:      &storage.Sku{Name: &sku},
			// TODO: They mark this property as optional but actually it is required
			AccessTier: &accessTier,
		},
	}

	acctReference := tc.MakeReferenceFromResource(acct)

	tc.CreateResourcesAndWait(domain, acct)

	queueServices := &storage.StorageAccountsQueueService{
		ObjectMeta: tc.MakeObjectMeta("dest-queues"),
		Spec: storage.StorageAccountsQueueServices_Spec{
			Owner: testcommon.AsOwner(acct),
		},
	}

	tc.CreateResourceAndWait(queueServices)

	queue := &storage.StorageAccountsQueueServicesQueue{
		ObjectMeta: tc.MakeObjectMeta("dest-queue"),
		Spec: storage.StorageAccountsQueueServicesQueues_Spec{
			Owner: testcommon.AsOwner(queueServices),
		},
	}

	tc.CreateResourceAndWait(queue)

	armId := *domain.Status.Id

	tc.RunParallelSubtests(
		testcommon.Subtest{
			Name: "CreateDomainTopicAndSubscription",
			Test: func(tc *testcommon.KubePerTestContext) {
				topic := &eventgrid.DomainsTopic{
					ObjectMeta: tc.MakeObjectMeta("topic"),
					Spec: eventgrid.DomainsTopics_Spec{
						Owner: testcommon.AsOwner(domain),
					},
				}

				tc.CreateResourceAndWait(topic)
				// don’t bother deleting; deleting domain will clean up

				endpointType := eventgrid.StorageQueueEventSubscriptionDestinationEndpointTypeStorageQueue
				subscription := &eventgrid.EventSubscription{
					ObjectMeta: tc.MakeObjectMeta("sub"),
					Spec: eventgrid.EventSubscriptions_Spec{
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
				endpointType := eventgrid.StorageQueueEventSubscriptionDestinationEndpointTypeStorageQueue
				subscription := &eventgrid.EventSubscription{
					ObjectMeta: tc.MakeObjectMeta("sub"),
					Spec: eventgrid.EventSubscriptions_Spec{
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

	tc.DeleteResourceAndWait(domain)

	// Ensure that the resource group was really deleted in Azure
	exists, _, err := tc.AzureClient.HeadByID(
		tc.Ctx,
		armId,
		string(eventgrid.TopicsSpecAPIVersion20200601))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeFalse())
}
