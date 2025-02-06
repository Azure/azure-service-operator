/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	v1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	signalrservice "github.com/Azure/azure-service-operator/v2/api/signalrservice/v1api20211001"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func Test_SignalRService_SignalR_CRUD(t *testing.T) {
	t.Parallel()
	tc := globalTestContext.ForTest(t)
	rg := tc.CreateTestResourceGroupAndWait()

	deny := signalrservice.ACLAction_Deny
	systemAssigned := signalrservice.ManagedIdentityType_SystemAssigned
	// Adapted from the quickstart example:
	// https://docs.microsoft.com/en-us/azure/azure-signalr/signalr-quickstart-azure-signalr-service-arm-template
	serviceModeFlag := signalrservice.FeatureFlags_ServiceMode
	connectivityLogsFlag := signalrservice.FeatureFlags_EnableConnectivityLogs
	enableMessagingLogsFlag := signalrservice.FeatureFlags_EnableMessagingLogs
	enableliveTraceFlag := signalrservice.FeatureFlags_EnableLiveTrace
	signalR := &signalrservice.SignalR{
		ObjectMeta: tc.MakeObjectMeta("signalr"),
		Spec: signalrservice.SignalR_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			Sku: &signalrservice.ResourceSku{
				Capacity: to.Ptr(1),
				Name:     to.Ptr("Standard_S1"),
			},
			Identity: &signalrservice.ManagedIdentity{
				Type: &systemAssigned,
			},
			Tls: &signalrservice.SignalRTlsSettings{
				ClientCertEnabled: to.Ptr(false),
			},
			Features: []signalrservice.SignalRFeature{{
				Flag:  &serviceModeFlag,
				Value: to.Ptr("Classic"),
			}, {
				Flag:  &connectivityLogsFlag,
				Value: to.Ptr("true"),
			}, {
				Flag:  &enableMessagingLogsFlag,
				Value: to.Ptr("true"),
			}, {
				Flag:  &enableliveTraceFlag,
				Value: to.Ptr("true"),
			}},
			Cors: &signalrservice.SignalRCorsSettings{
				AllowedOrigins: []string{"https://foo.com", "https://bar.com"},
			},
			NetworkACLs: &signalrservice.SignalRNetworkACLs{
				DefaultAction: &deny,
				PublicNetwork: &signalrservice.NetworkACL{
					Allow: []signalrservice.SignalRRequestType{
						signalrservice.SignalRRequestType_ClientConnection,
					},
				},
				PrivateEndpoints: []signalrservice.PrivateEndpointACL{{
					Name: to.Ptr("privateendpointname"),
					Allow: []signalrservice.SignalRRequestType{
						signalrservice.SignalRRequestType_ServerConnection,
					},
				}},
			},
			Upstream: &signalrservice.ServerlessUpstreamSettings{
				Templates: []signalrservice.UpstreamTemplate{{
					CategoryPattern: to.Ptr("*"),
					EventPattern:    to.Ptr("connect,disconnect"),
					HubPattern:      to.Ptr("*"),
					UrlTemplate:     to.Ptr("https://example.com/chat/api/connect"),
				}},
			},
		},
	}

	tc.CreateResourceAndWait(signalR)
	tc.Expect(signalR.Status.Id).ToNot(BeNil())
	armId := *signalR.Status.Id

	// Perform a patch to add another URL to the cors allow list.
	old := signalR.DeepCopy()
	signalR.Spec.Cors.AllowedOrigins = append(
		signalR.Spec.Cors.AllowedOrigins, "https://definitelymydomain.horse",
	)
	tc.PatchResourceAndWait(old, signalR)
	tc.Expect(signalR.Status.Cors).ToNot(BeNil())
	tc.Expect(signalR.Status.Cors.AllowedOrigins).To(ContainElement("https://definitelymydomain.horse"))

	// There should be no secrets at this point
	list := &v1.SecretList{}
	tc.ListResources(list, client.InNamespace(tc.Namespace))
	tc.Expect(list.Items).To(HaveLen(0))

	// Run sub-tests on the SignalR service in sequence
	tc.RunSubtests(
		testcommon.Subtest{
			Name: "SecretsWrittenToSameKubeSecret",
			Test: func(tc *testcommon.KubePerTestContext) {
				SignalR_SecretsWrittenToSameKubeSecret_20211001(tc, signalR)
			},
		},
		testcommon.Subtest{
			Name: "SecretsWrittenToDifferentKubeSecrets",
			Test: func(tc *testcommon.KubePerTestContext) {
				SignalR_SecretsWrittenToDifferentKubeSecrets_20211001(tc, signalR)
			},
		},
	)

	tc.DeleteResourcesAndWait(signalR)

	// Ensure that the resource was really deleted in Azure
	exists, retryAfter, err := tc.AzureClient.CheckExistenceWithGetByID(tc.Ctx, armId, string(signalrservice.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())
}

func SignalR_SecretsWrittenToSameKubeSecret_20211001(tc *testcommon.KubePerTestContext, signalR *signalrservice.SignalR) {
	old := signalR.DeepCopy()
	signalrSecret := "signalrsecret"
	signalR.Spec.OperatorSpec = &signalrservice.SignalROperatorSpec{
		Secrets: &signalrservice.SignalROperatorSecrets{
			PrimaryKey: &genruntime.SecretDestination{
				Name: signalrSecret,
				Key:  "primarykey",
			},
			PrimaryConnectionString: &genruntime.SecretDestination{
				Name: signalrSecret,
				Key:  "primaryconnectionstring",
			},
			SecondaryKey: &genruntime.SecretDestination{
				Name: signalrSecret,
				Key:  "secondarykey",
			},
			SecondaryConnectionString: &genruntime.SecretDestination{
				Name: signalrSecret,
				Key:  "secondaryconnectionstring",
			},
		},
	}
	tc.PatchResourceAndWait(old, signalR)

	tc.ExpectSecretHasKeys(signalrSecret, "primarykey", "primaryconnectionstring", "secondarykey", "secondaryconnectionstring")
}

func SignalR_SecretsWrittenToDifferentKubeSecrets_20211001(tc *testcommon.KubePerTestContext, signalR *signalrservice.SignalR) {
	old := signalR.DeepCopy()
	primaryKeySecret := "secret1"
	primaryConnectionString := "secret2"
	secondaryKeySecret := "secret3"
	secondaryConnectionString := "secret4"

	signalR.Spec.OperatorSpec = &signalrservice.SignalROperatorSpec{
		Secrets: &signalrservice.SignalROperatorSecrets{
			PrimaryKey: &genruntime.SecretDestination{
				Name: primaryKeySecret,
				Key:  "primarykey",
			},
			PrimaryConnectionString: &genruntime.SecretDestination{
				Name: primaryConnectionString,
				Key:  "primaryconnectionstring",
			},
			SecondaryKey: &genruntime.SecretDestination{
				Name: secondaryKeySecret,
				Key:  "secondarykey",
			},
			SecondaryConnectionString: &genruntime.SecretDestination{
				Name: secondaryConnectionString,
				Key:  "secondaryconnectionstring",
			},
		},
	}
	tc.PatchResourceAndWait(old, signalR)

	tc.ExpectSecretHasKeys(primaryKeySecret, "primarykey")
	tc.ExpectSecretHasKeys(primaryConnectionString, "primaryconnectionstring")
	tc.ExpectSecretHasKeys(secondaryKeySecret, "secondarykey")
	tc.ExpectSecretHasKeys(secondaryConnectionString, "secondaryconnectionstring")
}
