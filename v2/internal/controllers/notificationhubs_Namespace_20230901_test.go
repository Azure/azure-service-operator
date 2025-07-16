/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	notificationhubs "github.com/Azure/azure-service-operator/v2/api/notificationhubs/v1api20230901"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func Test_Notificationhubs_Namespace_20230901_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	namespace := &notificationhubs.Namespace{
		ObjectMeta: tc.MakeObjectMeta("ns"),
		Spec: notificationhubs.Namespace_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			Properties: &notificationhubs.NamespaceProperties{
				NamespaceType: to.Ptr(notificationhubs.NamespaceType_NotificationHub),
			},
			Sku: &notificationhubs.Sku{
				Capacity: to.Ptr(1),
				Name:     to.Ptr(notificationhubs.SkuName_Free),
			},
		},
	}

	tc.CreateResourceAndWait(namespace)

	tc.Expect(namespace.Status.Id).ToNot(BeNil())
	armId := *namespace.Status.Id

	tc.RunSubtests(
		testcommon.Subtest{
			Name: "WriteNamespacesSecrets",
			Test: func(tc *testcommon.KubePerTestContext) {
				Namespace_WriteSecrets(tc, namespace)
			},
		})

	tc.RunParallelSubtests(
		testcommon.Subtest{
			Name: "Test_NotificationHub_CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				NotificationHub_CRUD(tc, testcommon.AsOwner(namespace))
			},
		},
		testcommon.Subtest{
			Name: "Test_NamespacesAuthorizationRule_CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				NamespacesAuthorizationRule_CRUD(tc, testcommon.AsOwner(namespace))
			},
		},
	)

	tc.DeleteResourceAndWait(namespace)

	// Ensure that the resource was really deleted in Azure
	exists, retryAfter, err := tc.AzureClient.CheckExistenceWithGetByID(tc.Ctx, armId, string(notificationhubs.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())
}

func NamespacesAuthorizationRule_CRUD(tc *testcommon.KubePerTestContext, owner *genruntime.KnownResourceReference) {
	authrule := &notificationhubs.NamespacesAuthorizationRule{
		ObjectMeta: tc.MakeObjectMeta("policy"),
		Spec: notificationhubs.NamespacesAuthorizationRule_Spec{
			Location: tc.AzureRegion,
			Owner:    owner,
			Properties: &notificationhubs.SharedAccessAuthorizationRuleProperties{
				Rights: []notificationhubs.AccessRights{
					notificationhubs.AccessRights_Listen,
				},
			},
		},
	}

	tc.CreateResourceAndWait(authrule)
}

func NotificationHub_CRUD(tc *testcommon.KubePerTestContext, owner *genruntime.KnownResourceReference) {
	notificationHub := &notificationhubs.NotificationHub{
		ObjectMeta: tc.MakeObjectMeta("hub"),
		Spec: notificationhubs.NotificationHub_Spec{
			Location: tc.AzureRegion,
			Owner:    owner,
			Sku: &notificationhubs.Sku{
				Capacity: to.Ptr(1),
				Name:     to.Ptr(notificationhubs.SkuName_Free),
			},
		},
	}

	tc.CreateResourceAndWait(notificationHub)

	tc.Expect(notificationHub.Status.Id).ToNot(BeNil())
	armId := *notificationHub.Status.Id

	tc.RunSubtests(
		testcommon.Subtest{
			Name: "WriteNotificationHubSecrets",
			Test: func(tc *testcommon.KubePerTestContext) {
				NotificationHub_WriteSecrets(tc, notificationHub)
			},
		})

	tc.RunParallelSubtests(
		testcommon.Subtest{
			Name: "Test_NotificationHubAuthorizationRule_CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				NotificationHubAutorizationRule_CRUD(tc, testcommon.AsOwner(notificationHub))
			},
		},
	)

	tc.DeleteResourceAndWait(notificationHub)

	// Ensure that the resource was really deleted in Azure
	exists, retryAfter, err := tc.AzureClient.CheckExistenceWithGetByID(tc.Ctx, armId, string(notificationhubs.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())
}

func NotificationHubAutorizationRule_CRUD(tc *testcommon.KubePerTestContext, owner *genruntime.KnownResourceReference) {
	authrule := &notificationhubs.NotificationHubsAuthorizationRule{
		ObjectMeta: tc.MakeObjectMeta("policy"),
		Spec: notificationhubs.NotificationHubsAuthorizationRule_Spec{
			Location: tc.AzureRegion,
			Owner:    owner,
			Properties: &notificationhubs.SharedAccessAuthorizationRuleProperties{
				Rights: []notificationhubs.AccessRights{
					notificationhubs.AccessRights_Listen,
				},
			},
		},
	}

	tc.CreateResourceAndWait(authrule)
}

func Namespace_WriteSecrets(tc *testcommon.KubePerTestContext, namespace *notificationhubs.Namespace) {
	old := namespace.DeepCopy()
	namespaceKeysSecret := "namespacekeyssecret"
	namespace.Spec.OperatorSpec = &notificationhubs.NamespaceOperatorSpec{
		Secrets: &notificationhubs.NamespaceOperatorSecrets{
			PrimaryConnectionString:   &genruntime.SecretDestination{Name: namespaceKeysSecret, Key: "primaryConnectionString"},
			SecondaryConnectionString: &genruntime.SecretDestination{Name: namespaceKeysSecret, Key: "secondaryConnectionString"},
			PrimaryKey:                &genruntime.SecretDestination{Name: namespaceKeysSecret, Key: "primaryKey"},
			SecondaryKey:              &genruntime.SecretDestination{Name: namespaceKeysSecret, Key: "secondaryKey"},
		},
	}
	tc.PatchResourceAndWait(old, namespace)

	tc.ExpectSecretHasKeys(
		namespaceKeysSecret,
		"primaryConnectionString",
		"secondaryConnectionString",
		"primaryKey",
		"secondaryKey",
	)
}

func NotificationHub_WriteSecrets(tc *testcommon.KubePerTestContext, hub *notificationhubs.NotificationHub) {
	old := hub.DeepCopy()
	namespaceKeysSecret := "notificationhubkeyssecret"
	hub.Spec.OperatorSpec = &notificationhubs.NotificationHubOperatorSpec{
		Secrets: &notificationhubs.NotificationHubOperatorSecrets{
			PrimaryConnectionString:   &genruntime.SecretDestination{Name: namespaceKeysSecret, Key: "primaryConnectionString"},
			SecondaryConnectionString: &genruntime.SecretDestination{Name: namespaceKeysSecret, Key: "secondaryConnectionString"},
			PrimaryKey:                &genruntime.SecretDestination{Name: namespaceKeysSecret, Key: "primaryKey"},
			SecondaryKey:              &genruntime.SecretDestination{Name: namespaceKeysSecret, Key: "secondaryKey"},
		},
	}
	tc.PatchResourceAndWait(old, hub)

	tc.ExpectSecretHasKeys(
		namespaceKeysSecret,
		"primaryConnectionString",
		"secondaryConnectionString",
		"primaryKey",
		"secondaryKey",
	)
}
