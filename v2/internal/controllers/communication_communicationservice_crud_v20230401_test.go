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

	communication "github.com/Azure/azure-service-operator/v2/api/communication/v20230401"
	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1api20200601"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func Test_Communication_CommunicationService_20230401_CRUD(t *testing.T) {
	t.Parallel()
	tc := globalTestContext.ForTest(t)
	rg := tc.CreateTestResourceGroupAndWait()

	// CommunicationService requires location "global" and a dataLocation for data residency.
	svc := &communication.CommunicationService{
		ObjectMeta: tc.MakeObjectMeta("commssvc"),
		Spec: communication.CommunicationService_Spec{
			Location:     to.Ptr("global"),
			Owner:        testcommon.AsOwner(rg),
			DataLocation: to.Ptr("UnitedStates"),
		},
	}

	tc.CreateResourceAndWait(svc)
	tc.Expect(svc.Status.Id).ToNot(BeNil())
	armId := *svc.Status.Id

	// Verify Azure-populated fields are present in status
	tc.Expect(svc.Status.HostName).ToNot(BeNil())

	// There should be no secrets at this point
	list := &v1.SecretList{}
	tc.ListResources(list, client.InNamespace(tc.Namespace))
	tc.Expect(list.Items).To(HaveLen(0))

	// Run sequential subtests for secrets
	tc.RunSubtests(
		testcommon.Subtest{
			Name: "SecretsWrittenToSameKubeSecret",
			Test: func(tc *testcommon.KubePerTestContext) {
				CommunicationService_SecretsWrittenToSameKubeSecret_20230401(tc, svc)
			},
		},
		testcommon.Subtest{
			Name: "SecretsWrittenToDifferentKubeSecrets",
			Test: func(tc *testcommon.KubePerTestContext) {
				CommunicationService_SecretsWrittenToDifferentKubeSecrets_20230401(tc, svc)
			},
		},
	)

	// Run parallel subtests for child resource hierarchies
	tc.RunParallelSubtests(
		testcommon.Subtest{
			Name: "EmailService CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				Communication_EmailService_20230401(tc, rg)
			},
		},
	)

	tc.DeleteResourceAndWait(svc)

	// Verify the resource was deleted in Azure
	exists, retryAfter, err := tc.AzureClient.CheckExistenceWithGetByID(tc.Ctx, armId, string(communication.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())
}

func CommunicationService_SecretsWrittenToSameKubeSecret_20230401(tc *testcommon.KubePerTestContext, svc *communication.CommunicationService) {
	old := svc.DeepCopy()
	secretName := "commssecret"
	svc.Spec.OperatorSpec = &communication.CommunicationServiceOperatorSpec{
		Secrets: &communication.CommunicationServiceOperatorSecrets{
			PrimaryKey: &genruntime.SecretDestination{
				Name: secretName,
				Key:  "primaryKey",
			},
			PrimaryConnectionString: &genruntime.SecretDestination{
				Name: secretName,
				Key:  "primaryConnectionString",
			},
			SecondaryKey: &genruntime.SecretDestination{
				Name: secretName,
				Key:  "secondaryKey",
			},
			SecondaryConnectionString: &genruntime.SecretDestination{
				Name: secretName,
				Key:  "secondaryConnectionString",
			},
		},
	}
	tc.PatchResourceAndWait(old, svc)
	tc.ExpectSecretHasKeys(secretName, "primaryKey", "primaryConnectionString", "secondaryKey", "secondaryConnectionString")
}

func CommunicationService_SecretsWrittenToDifferentKubeSecrets_20230401(tc *testcommon.KubePerTestContext, svc *communication.CommunicationService) {
	old := svc.DeepCopy()
	svc.Spec.OperatorSpec = &communication.CommunicationServiceOperatorSpec{
		Secrets: &communication.CommunicationServiceOperatorSecrets{
			PrimaryKey: &genruntime.SecretDestination{
				Name: "commssecret1",
				Key:  "primaryKey",
			},
			PrimaryConnectionString: &genruntime.SecretDestination{
				Name: "commssecret2",
				Key:  "primaryConnectionString",
			},
			SecondaryKey: &genruntime.SecretDestination{
				Name: "commssecret3",
				Key:  "secondaryKey",
			},
			SecondaryConnectionString: &genruntime.SecretDestination{
				Name: "commssecret4",
				Key:  "secondaryConnectionString",
			},
		},
	}
	tc.PatchResourceAndWait(old, svc)
	tc.ExpectSecretHasKeys("commssecret1", "primaryKey")
	tc.ExpectSecretHasKeys("commssecret2", "primaryConnectionString")
	tc.ExpectSecretHasKeys("commssecret3", "secondaryKey")
	tc.ExpectSecretHasKeys("commssecret4", "secondaryConnectionString")
}

func Communication_EmailService_20230401(tc *testcommon.KubePerTestContext, rg *resources.ResourceGroup) {
	emailSvc := &communication.EmailService{
		ObjectMeta: tc.MakeObjectMeta("emailsvc"),
		Spec: communication.EmailService_Spec{
			Location:     to.Ptr("global"),
			Owner:        testcommon.AsOwner(rg),
			DataLocation: to.Ptr("UnitedStates"),
		},
	}

	tc.CreateResourceAndWait(emailSvc)
	tc.Expect(emailSvc.Status.Id).ToNot(BeNil())
	armId := *emailSvc.Status.Id

	// Run Domain CRUD as parallel subtest
	tc.RunParallelSubtests(
		testcommon.Subtest{
			Name: "Domain CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				Communication_Domain_20230401(tc, emailSvc)
			},
		},
	)

	tc.DeleteResourceAndWait(emailSvc)

	// Verify deleted in Azure
	exists, retryAfter, err := tc.AzureClient.CheckExistenceWithGetByID(tc.Ctx, armId, string(communication.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())
}

func Communication_Domain_20230401(tc *testcommon.KubePerTestContext, emailSvc *communication.EmailService) {
	azureManaged := communication.DomainManagement_AzureManaged
	disabled := communication.UserEngagementTracking_Disabled

	// AzureManaged domains require the Azure name to be exactly "AzureManagedDomain" per the Azure API.
	// The K8s metadata name must be lowercase, so we set AzureName explicitly.
	domain := &communication.Domain{
		ObjectMeta: tc.MakeObjectMetaWithName("azuremanageddomain"),
		Spec: communication.Domain_Spec{
			AzureName:              "AzureManagedDomain",
			Location:               to.Ptr("global"),
			Owner:                  testcommon.AsOwner(emailSvc),
			DomainManagement:       &azureManaged,
			UserEngagementTracking: &disabled,
		},
	}

	tc.CreateResourceAndWait(domain)
	tc.Expect(domain.Status.Id).ToNot(BeNil())
	armId := *domain.Status.Id

	// Verify Azure-managed domain fields are populated
	tc.Expect(domain.Status.VerificationStates).ToNot(BeNil())
	tc.Expect(domain.Status.MailFromSenderDomain).ToNot(BeNil())

	// Run SenderUsername CRUD as parallel subtest
	tc.RunParallelSubtests(
		testcommon.Subtest{
			Name: "SenderUsername CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				Communication_SenderUsername_20230401(tc, domain)
			},
		},
	)

	tc.DeleteResourceAndWait(domain)

	// Verify deleted in Azure
	exists, retryAfter, err := tc.AzureClient.CheckExistenceWithGetByID(tc.Ctx, armId, string(communication.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())
}

func Communication_SenderUsername_20230401(tc *testcommon.KubePerTestContext, domain *communication.Domain) {
	sender := &communication.SenderUsername{
		ObjectMeta: tc.MakeObjectMeta("sender"),
		Spec: communication.SenderUsername_Spec{
			// AzureName must match Username - Azure requires the resource name in the URL
			// to match the username field in the request body.
			AzureName:   "testuser",
			Owner:       testcommon.AsOwner(domain),
			Username:    to.Ptr("testuser"),
			DisplayName: to.Ptr("ASO Test User"),
		},
	}

	tc.CreateResourceAndWait(sender)
	tc.Expect(sender.Status.Id).ToNot(BeNil())
	armId := *sender.Status.Id

	// Verify fields from status
	tc.Expect(sender.Status.Username).ToNot(BeNil())
	tc.Expect(*sender.Status.Username).To(Equal("testuser"))

	tc.DeleteResourceAndWait(sender)

	// Verify deleted in Azure
	exists, retryAfter, err := tc.AzureClient.CheckExistenceWithGetByID(tc.Ctx, armId, string(communication.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())
}
