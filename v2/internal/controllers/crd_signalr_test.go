/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	"github.com/Azure/go-autorest/autorest/to"
	. "github.com/onsi/gomega"

	signalrservice "github.com/Azure/azure-service-operator/v2/api/signalrservice/v1alpha1api20211001"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
)

func Test_SignalRService_SignalR_CRUD(t *testing.T) {
	t.Parallel()
	tc := globalTestContext.ForTest(t)
	rg := tc.CreateTestResourceGroupAndWait()

	deny := signalrservice.SignalRNetworkACLsDefaultActionDeny
	systemAssigned := signalrservice.ManagedIdentityTypeSystemAssigned
	// Adapted from the quickstart example:
	// https://docs.microsoft.com/en-us/azure/azure-signalr/signalr-quickstart-azure-signalr-service-arm-template
	signalR := signalrservice.SignalR{
		ObjectMeta: tc.MakeObjectMeta("signalr"),
		Spec: signalrservice.SignalR_Spec{
			Location: &tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			Sku: &signalrservice.ResourceSku{
				Capacity: to.IntPtr(1),
				Name:     "Standard_S1",
			},
			Identity: &signalrservice.ManagedIdentity{
				Type: &systemAssigned,
			},
			Tls: &signalrservice.SignalRTlsSettings{
				ClientCertEnabled: to.BoolPtr(false),
			},
			Features: []signalrservice.SignalRFeature{{
				Flag:  signalrservice.SignalRFeatureFlagServiceMode,
				Value: "Classic",
			}, {
				Flag:  signalrservice.SignalRFeatureFlagEnableConnectivityLogs,
				Value: "true",
			}, {
				Flag:  signalrservice.SignalRFeatureFlagEnableMessagingLogs,
				Value: "true",
			}, {
				Flag:  signalrservice.SignalRFeatureFlagEnableLiveTrace,
				Value: "true",
			}},
			Cors: &signalrservice.SignalRCorsSettings{
				AllowedOrigins: []string{"https://foo.com", "https://bar.com"},
			},
			NetworkACLs: &signalrservice.SignalRNetworkACLs{
				DefaultAction: &deny,
				PublicNetwork: &signalrservice.NetworkACL{
					Allow: []signalrservice.NetworkACLAllow{
						signalrservice.NetworkACLAllowClientConnection,
					},
				},
				PrivateEndpoints: []signalrservice.PrivateEndpointACL{{
					Name: "privateendpointname",
					Allow: []signalrservice.PrivateEndpointACLAllow{
						signalrservice.PrivateEndpointACLAllowServerConnection,
					},
				}},
			},
			Upstream: &signalrservice.ServerlessUpstreamSettings{
				Templates: []signalrservice.UpstreamTemplate{{
					CategoryPattern: to.StringPtr("*"),
					EventPattern:    to.StringPtr("connect,disconnect"),
					HubPattern:      to.StringPtr("*"),
					UrlTemplate:     "https://example.com/chat/api/connect",
				}},
			},
		},
	}

	tc.CreateResourceAndWait(&signalR)
	tc.Expect(signalR.Status.Id).ToNot(BeNil())
	armId := *signalR.Status.Id

	// Perform a patch to add another URL to the cors allow list.
	old := signalR.DeepCopy()
	signalR.Spec.Cors.AllowedOrigins = append(
		signalR.Spec.Cors.AllowedOrigins, "https://definitelymydomain.horse",
	)
	tc.PatchResourceAndWait(old, &signalR)
	tc.Expect(signalR.Status.Cors).ToNot(BeNil())
	tc.Expect(signalR.Status.Cors.AllowedOrigins).To(ContainElement("https://definitelymydomain.horse"))

	tc.DeleteResourcesAndWait(&signalR, rg)

	// Ensure that the resource was really deleted in Azure
	exists, retryAfter, err := tc.AzureClient.HeadByID(tc.Ctx, armId, string(signalrservice.SignalRSpecAPIVersion20211001))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())
}
