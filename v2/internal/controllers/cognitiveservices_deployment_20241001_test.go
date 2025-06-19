/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	cognitiveservices "github.com/Azure/azure-service-operator/v2/api/cognitiveservices/v1api20241001"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func Test_CognitiveServices_Deployment_CRUD(t *testing.T) {
	t.Parallel()
	tc := globalTestContext.ForTest(t)
	rg := tc.CreateTestResourceGroupAndWait()

	account := &cognitiveservices.Account{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.NoSpaceNamer.GenerateName("cogsvcacc")),
		Spec: cognitiveservices.Account_Spec{
			Identity: &cognitiveservices.Identity{
				Type: to.Ptr(cognitiveservices.Identity_Type("SystemAssigned")),
			},
			Kind:     to.Ptr("OpenAI"),
			Location: to.Ptr("eastus"),
			OperatorSpec: &cognitiveservices.AccountOperatorSpec{
				Secrets: &cognitiveservices.AccountOperatorSecrets{
					Key1:     &genruntime.SecretDestination{Name: "cogsecrets", Key: "key1"},
					Key2:     &genruntime.SecretDestination{Name: "cogsecrets", Key: "key2"},
					Endpoint: &genruntime.SecretDestination{Name: "cogsecrets", Key: "endpoint"},
					Endpoints: &genruntime.SecretDestination{
						Name: "cogsecrets", Key: "endpoints",
					},
				},
			},
			Owner: testcommon.AsOwner(rg),
			Properties: &cognitiveservices.AccountProperties{
				PublicNetworkAccess: to.Ptr(cognitiveservices.AccountProperties_PublicNetworkAccess("Enabled")),
			},
			Sku: &cognitiveservices.Sku{
				Name: to.Ptr("S0"),
			},
		},
	}

	tc.CreateResourcesAndWait(account)
	tc.Expect(account.Status.Id).ToNot(BeNil())

	deployment := &cognitiveservices.Deployment{
		ObjectMeta: tc.MakeObjectMetaWithName("cogsvcdep"),
		Spec: cognitiveservices.Deployment_Spec{
			Owner: testcommon.AsOwner(account),
			Properties: &cognitiveservices.DeploymentProperties{
				Model: &cognitiveservices.DeploymentModel{
					Name:      to.Ptr("gpt-4o"),
					Format:    to.Ptr("OpenAI"),
					Publisher: to.Ptr("OpenAI"),
					Version:   to.Ptr("2024-08-06"),
				},
			},
			Sku: &cognitiveservices.Sku{
				Capacity: to.Ptr(1),
				Name:     to.Ptr("Standard"),
			},
		},
	}

	tc.CreateResourcesAndWait(deployment)
	tc.Expect(deployment.Status.Id).ToNot(BeNil())

	old := deployment.DeepCopy()
	deployment.Spec.Properties.Model.Version = to.Ptr("2024-11-20")
	tc.PatchResourceAndWait(old, deployment)
	tc.Expect(deployment.Status.Properties.Model.Version).To(Equal(to.Ptr("2024-11-20")))

	tc.DeleteResourceAndWait(deployment)
	exists, _, err := tc.AzureClient.CheckExistenceWithGetByID(tc.Ctx, *deployment.Status.Id, string(cognitiveservices.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeFalse())
	tc.Expect(account.Status.Id).ToNot(BeNil())

	tc.DeleteResourceAndWait(account)
	exists, _, err = tc.AzureClient.CheckExistenceWithGetByID(tc.Ctx, *account.Status.Id, string(cognitiveservices.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeFalse())
}
