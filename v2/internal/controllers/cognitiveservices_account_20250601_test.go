/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	cognitiveservices "github.com/Azure/azure-service-operator/v2/api/cognitiveservices/v1api20250601"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/core"
)

func Test_CognitiveServices_Account_20250601_CRUD(t *testing.T) {
	t.Parallel()
	tc := globalTestContext.ForTest(t)

	// Moving test to a different region:
	// AI Foundry not available in westus2: https://learn.microsoft.com/en-us/azure/ai-foundry/reference/region-support
	tc.AzureRegion = to.Ptr("westus3")

	rg := tc.CreateTestResourceGroupAndWait()

	account := &cognitiveservices.Account{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.NoSpaceNamer.GenerateName("cogsvcacc")),
		Spec: cognitiveservices.Account_Spec{
			Identity: &cognitiveservices.Identity{
				Type: to.Ptr(cognitiveservices.Identity_Type_SystemAssigned),
			},
			Kind:     to.Ptr("OpenAI"),
			Location: tc.AzureRegion,
			OperatorSpec: &cognitiveservices.AccountOperatorSpec{
				Secrets: &cognitiveservices.AccountOperatorSecrets{
					Key1: &genruntime.SecretDestination{Name: "cogsecrets", Key: "key1"},
					Key2: &genruntime.SecretDestination{Name: "cogsecrets", Key: "key2"},
				},
				SecretExpressions: []*core.DestinationExpression{
					{
						Name:  "cogsecrets",
						Key:   "endpoint",
						Value: "self.status.properties.endpoint",
					},
					{
						Name:  "cogsecrets",
						Value: `self.status.properties.endpoints.transformMapEntry(k, v, {k.replace(" ", "-").lowerAscii(): v})`,
					},
				},
			},
			Owner: testcommon.AsOwner(rg),
			Properties: &cognitiveservices.AccountProperties{
				PublicNetworkAccess: to.Ptr(cognitiveservices.AccountProperties_PublicNetworkAccess_Enabled),
			},
			Sku: &cognitiveservices.Sku{
				Name: to.Ptr("S0"),
			},
		},
	}

	tc.CreateResourcesAndWait(account)
	tc.Expect(account.Status.Id).ToNot(BeNil())
	tc.ExpectSecretHasKeys(
		"cogsecrets",
		"key1",
		"key2",
		"endpoint",
		"openai-dall-e-api",
		"openai-language-model-instance-api",
		"openai-model-scaleset-api",
		"openai-moderations-api",
		"openai-realtime-api",
		"openai-sora-api",
		"openai-whisper-api",
		"token-service-api",
	)

	oldAcc := account.DeepCopy()
	if oldAcc.Spec.Tags == nil {
		oldAcc.Spec.Tags = map[string]string{}
	}
	account.Spec.Tags = map[string]string{"env": "test"}
	tc.PatchResourceAndWait(oldAcc, account)
	tc.Expect(account.Status.Tags).To(Equal(map[string]string{"env": "test"}))

	tc.RunParallelSubtests(
		testcommon.Subtest{
			Name: "CognitiveServices Deployment CRUD",
			Test: cognitiveServicesDeploymentCRUD(account),
		},
	)

	tc.DeleteResourceAndWait(account)
	exists, _, err := tc.AzureClient.CheckExistenceWithGetByID(tc.Ctx, *account.Status.Id, string(cognitiveservices.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeFalse())
}

func cognitiveServicesDeploymentCRUD(account *cognitiveservices.Account) func(tc *testcommon.KubePerTestContext) {
	return func(tc *testcommon.KubePerTestContext) {
		deployment := &cognitiveservices.Deployment{
			ObjectMeta: tc.MakeObjectMetaWithName(tc.NoSpaceNamer.GenerateName("cogsvcdep")),
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

		oldDep := deployment.DeepCopy()
		deployment.Spec.Properties.Model.Version = to.Ptr("2024-11-20")
		tc.PatchResourceAndWait(oldDep, deployment)
		tc.Expect(deployment.Status.Properties.Model.Version).To(Equal(to.Ptr("2024-11-20")))

		tc.DeleteResourceAndWait(deployment)
		exists, _, err := tc.AzureClient.CheckExistenceWithGetByID(tc.Ctx, *deployment.Status.Id, string(cognitiveservices.APIVersion_Value))
		tc.Expect(err).ToNot(HaveOccurred())
		tc.Expect(exists).To(BeFalse())
	}
}
