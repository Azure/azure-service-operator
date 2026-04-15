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
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func Test_Communication_CommunicationService_20230401_CRUD(t *testing.T) {
	t.Parallel()
	tc := globalTestContext.ForTest(t)
	rg := tc.CreateTestResourceGroupAndWait()

	svc := newCommunicationService(tc, rg)
	emailSvc := newEmailService(tc, rg)
	domain := newAzureManagedDomain(tc, emailSvc)
	sender := newSenderUsername(tc, domain)

	// Create all resources in one call — ASO uses ownership to orchestrate the correct ordering.
	tc.CreateResourcesAndWait(svc, emailSvc, domain, sender)

	// Verify CommunicationService
	tc.Expect(svc.Status.Id).ToNot(BeNil())
	tc.Expect(svc.Status.HostName).ToNot(BeNil())

	// Verify EmailService
	tc.Expect(emailSvc.Status.Id).ToNot(BeNil())

	// Verify Domain
	tc.Expect(domain.Status.Id).ToNot(BeNil())
	tc.Expect(domain.Status.VerificationStates).ToNot(BeNil())
	tc.Expect(domain.Status.MailFromSenderDomain).ToNot(BeNil())

	// Verify SenderUsername
	tc.Expect(sender.Status.Id).ToNot(BeNil())
	tc.Expect(sender.Status.Username).ToNot(BeNil())
	tc.Expect(*sender.Status.Username).To(Equal("testuser"))

	// There should be no secrets at this point
	list := &v1.SecretList{}
	tc.ListResources(list, client.InNamespace(tc.Namespace))
	tc.Expect(list.Items).To(HaveLen(0))

	// Run subtests for secrets and configmap exports (these need the resource already created)
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
		testcommon.Subtest{
			Name: "ConfigMapWritten",
			Test: func(tc *testcommon.KubePerTestContext) {
				CommunicationService_ConfigMapWritten_20230401(tc, svc)
			},
		},
	)

	// Delete leaf resources first, then parents
	svcArmId := *svc.Status.Id
	senderArmId := *sender.Status.Id
	domainArmId := *domain.Status.Id
	emailSvcArmId := *emailSvc.Status.Id

	tc.DeleteResourceAndWait(sender)
	tc.DeleteResourceAndWait(domain)
	tc.DeleteResourceAndWait(emailSvc)
	tc.DeleteResourceAndWait(svc)

	// Verify all resources were deleted in Azure
	for _, armId := range []string{senderArmId, domainArmId, emailSvcArmId, svcArmId} {
		exists, retryAfter, err := tc.AzureClient.CheckExistenceWithGetByID(tc.Ctx, armId, string(communication.APIVersion_Value))
		tc.Expect(err).ToNot(HaveOccurred())
		tc.Expect(retryAfter).To(BeZero())
		tc.Expect(exists).To(BeFalse())
	}
}

func newCommunicationService(tc *testcommon.KubePerTestContext, rg client.Object) *communication.CommunicationService {
	return &communication.CommunicationService{
		ObjectMeta: tc.MakeObjectMeta("commssvc"),
		Spec: communication.CommunicationService_Spec{
			Location:     to.Ptr("global"),
			Owner:        testcommon.AsOwner(rg),
			DataLocation: to.Ptr("UnitedStates"),
		},
	}
}

func newEmailService(tc *testcommon.KubePerTestContext, rg client.Object) *communication.EmailService {
	return &communication.EmailService{
		ObjectMeta: tc.MakeObjectMeta("emailsvc"),
		Spec: communication.EmailService_Spec{
			Location:     to.Ptr("global"),
			Owner:        testcommon.AsOwner(rg),
			DataLocation: to.Ptr("UnitedStates"),
		},
	}
}

func newAzureManagedDomain(tc *testcommon.KubePerTestContext, emailSvc *communication.EmailService) *communication.Domain {
	azureManaged := communication.DomainManagement_AzureManaged
	disabled := communication.UserEngagementTracking_Disabled

	return &communication.Domain{
		ObjectMeta: tc.MakeObjectMetaWithName("azuremanageddomain"),
		Spec: communication.Domain_Spec{
			AzureName:              "AzureManagedDomain",
			Location:               to.Ptr("global"),
			Owner:                  testcommon.AsOwner(emailSvc),
			DomainManagement:       &azureManaged,
			UserEngagementTracking: &disabled,
		},
	}
}

func newSenderUsername(tc *testcommon.KubePerTestContext, domain *communication.Domain) *communication.SenderUsername {
	return &communication.SenderUsername{
		ObjectMeta: tc.MakeObjectMeta("sender"),
		Spec: communication.SenderUsername_Spec{
			AzureName:   "testuser",
			Owner:       testcommon.AsOwner(domain),
			Username:    to.Ptr("testuser"),
			DisplayName: to.Ptr("ASO Test User"),
		},
	}
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

func CommunicationService_ConfigMapWritten_20230401(tc *testcommon.KubePerTestContext, svc *communication.CommunicationService) {
	old := svc.DeepCopy()
	configMapName := "commsconfig"
	svc.Spec.OperatorSpec = &communication.CommunicationServiceOperatorSpec{
		ConfigMaps: &communication.CommunicationServiceOperatorConfigMaps{
			HostName: &genruntime.ConfigMapDestination{
				Name: configMapName,
				Key:  "hostName",
			},
		},
	}
	tc.PatchResourceAndWait(old, svc)
	tc.ExpectConfigMapHasKeysAndValues(configMapName, "hostName", *svc.Status.HostName)
}
