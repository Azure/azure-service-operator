/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"
	"sigs.k8s.io/controller-runtime/pkg/client"

	eventgrid "github.com/Azure/azure-service-operator/v2/api/eventgrid/v1alpha1api20200601"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
)

func Test_EventGrid_Topic(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	// Create a topic
	topic := &eventgrid.Topic{
		ObjectMeta: tc.MakeObjectMeta("topic"),
		Spec: eventgrid.Topics_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			Tags:     map[string]string{"cheese": "blue"},
		},
	}

	tc.CreateResourceAndWait(topic)

	// Perform a simple patch.
	old := topic.DeepCopy()
	topic.Spec.Tags["cheese"] = "époisses"
	tc.Patch(old, topic)

	armId := *topic.Status.Id
	objectKey := client.ObjectKeyFromObject(topic)

	// Ensure state eventually gets updated in k8s from change in Azure.
	tc.Eventually(func() map[string]string {
		var updatedTopic eventgrid.Topic
		tc.GetResource(objectKey, &updatedTopic)
		return updatedTopic.Status.Tags
	}).Should(Equal(map[string]string{"cheese": "époisses"}))

	tc.RunParallelSubtests(
		testcommon.Subtest{
			Name: "CreateTopicSubscription",
			Test: func(tc *testcommon.KubePerTestContext) {
				subscription := eventgrid.EventSubscription{
					ObjectMeta: tc.MakeObjectMeta("sub"),
					Spec: eventgrid.EventSubscriptions_Spec{
						Owner: tc.AsExtensionOwner(topic),
						/*
							Filter: &eventgrid.EventSubscriptionFilter{
								AdvancedFilters: []eventgrid.AdvancedFilter{
									{
										NumberGreaterThanOrEquals: &eventgrid.AdvancedFilter_NumberGreaterThanOrEquals{
											// TODO: this should be auto-populated
											OperatorType: eventgrid.AdvancedFilterNumberGreaterThanOrEqualsOperatorTypeNumberGreaterThanOrEquals,
											Key:          to.StringPtr("key"),
											Value:        to.Float64Ptr(123),
										},
									},
								},
							},
						*/
					},
				}

				tc.CreateResourceAndWait(&subscription)
				// don’t delete; deleting topic will clean up
			},
		},
	)

	tc.DeleteResourceAndWait(topic)

	// Ensure that the resource group was really deleted in Azure
	exists, _, err := tc.AzureClient.HeadByID(
		tc.Ctx,
		armId,
		string(eventgrid.TopicsSpecAPIVersion20200601))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeFalse())
}
