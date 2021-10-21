/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	eventgrid "github.com/Azure/azure-service-operator/v2/api/microsoft.eventgrid/v1alpha1api20200601"
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

	tc.CreateResourceAndWait(domain)

	armId := *domain.Status.Id
	// objectKey := client.ObjectKeyFromObject(domain)

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

				subscription := &eventgrid.EventSubscription{
					ObjectMeta: tc.MakeObjectMeta("sub"),
					Spec: eventgrid.EventSubscriptions_Spec{
						Owner: tc.AsExtensionOwner(topic),
					},
				}

				tc.CreateResourceAndWait(subscription)
				// don’t bother deleting
			},
		},
		testcommon.Subtest{
			Name: "CreateDomainSubscription",
			Test: func(tc *testcommon.KubePerTestContext) {
				subscription := &eventgrid.EventSubscription{
					ObjectMeta: tc.MakeObjectMeta("sub"),
					Spec: eventgrid.EventSubscriptions_Spec{
						Owner: tc.AsExtensionOwner(domain),
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
