/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	cognitiveservices "github.com/Azure/azure-service-operator/v2/api/cognitiveservices/v20250601"
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
			Kind:     to.Ptr("AIServices"),
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
						Key:   "openai-endpoint",
						Value: `self.status.properties.endpoints["OpenAI Language Model Instance API"]`,
					},
				},
			},
			Owner: testcommon.AsOwner(rg),
			Properties: &cognitiveservices.AccountProperties{
				AllowProjectManagement: to.Ptr(true),
				CustomSubDomainName:    to.Ptr(tc.NoSpaceNamer.GenerateName("cogsvcacc")),
				PublicNetworkAccess:    to.Ptr(cognitiveservices.AccountProperties_PublicNetworkAccess_Enabled),
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
		"openai-endpoint",
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
		testcommon.Subtest{
			Name: "CognitiveServices Project CRUD",
			Test: cognitiveServicesProjectCRUD(account),
		},
	)

	tc.DeleteResourceAndWait(account)
	exists, _, err := tc.AzureClient.CheckExistenceWithGetByID(tc.Ctx, *account.Status.Id, string(cognitiveservices.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeFalse())
}

func cognitiveServicesProjectCRUD(account *cognitiveservices.Account) func(tc *testcommon.KubePerTestContext) {
	return func(tc *testcommon.KubePerTestContext) {
		project := &cognitiveservices.Project{
			ObjectMeta: tc.MakeObjectMetaWithName(tc.NoSpaceNamer.GenerateName("cogsvcprj")),
			Spec: cognitiveservices.Project_Spec{
				Identity: &cognitiveservices.Identity{
					Type: to.Ptr(cognitiveservices.Identity_Type_SystemAssigned),
				},
				Location: tc.AzureRegion,
				Owner:    testcommon.AsOwner(account),
				Properties: &cognitiveservices.ProjectProperties{
					Description: to.Ptr("Test project for ASO integration tests"),
					DisplayName: to.Ptr("ASO Test Project"),
				},
			},
		}

		tc.CreateResourcesAndWait(project)
		tc.Expect(project.Status.Id).ToNot(BeNil())
		tc.Expect(project.Status.Properties).ToNot(BeNil())
		tc.Expect(project.Status.Properties.DisplayName).To(Equal(to.Ptr("ASO Test Project")))

		// Update the project description
		oldProject := project.DeepCopy()
		project.Spec.Properties.Description = to.Ptr("Updated description")
		tc.PatchResourceAndWait(oldProject, project)
		tc.Expect(project.Status.Properties.Description).To(Equal(to.Ptr("Updated description")))

		tc.DeleteResourceAndWait(project)
		exists, _, err := tc.AzureClient.CheckExistenceWithGetByID(tc.Ctx, *project.Status.Id, string(cognitiveservices.APIVersion_Value))
		tc.Expect(err).ToNot(HaveOccurred())
		tc.Expect(exists).To(BeFalse())
	}
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
						Version:   to.Ptr("2024-11-20"),
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
		deployment.Spec.Sku.Capacity = to.Ptr(2)
		tc.PatchResourceAndWait(oldDep, deployment)
		tc.Expect(deployment.Status.Sku.Capacity).To(Equal(to.Ptr(2)))
		tc.DeleteResourceAndWait(deployment)
		exists, _, err := tc.AzureClient.CheckExistenceWithGetByID(tc.Ctx, *deployment.Status.Id, string(cognitiveservices.APIVersion_Value))
		tc.Expect(err).ToNot(HaveOccurred())
		tc.Expect(exists).To(BeFalse())
	}
}
