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
	tc.AzureRegion = new("westus3")

	rg := tc.CreateTestResourceGroupAndWait()

	account := &cognitiveservices.Account{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.NoSpaceNamer.GenerateName("cogsvcacc")),
		Spec: cognitiveservices.Account_Spec{
			Identity: &cognitiveservices.Identity{
				Type: to.Ptr(cognitiveservices.Identity_Type_SystemAssigned),
			},
			Kind:     new("AIServices"),
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
				AllowProjectManagement: new(true),
				CustomSubDomainName:    new(tc.NoSpaceNamer.GenerateName("cogsvcacc")),
				PublicNetworkAccess:    to.Ptr(cognitiveservices.AccountProperties_PublicNetworkAccess_Enabled),
			},
			Sku: &cognitiveservices.Sku{
				Name: new("S0"),
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
					Description: new("Test project for ASO integration tests"),
					DisplayName: new("ASO Test Project"),
				},
			},
		}

		tc.CreateResourcesAndWait(project)
		tc.Expect(project.Status.Id).ToNot(BeNil())
		tc.Expect(project.Status.Properties).ToNot(BeNil())
		tc.Expect(project.Status.Properties.DisplayName).To(Equal(new("ASO Test Project")))

		// Update the project description
		oldProject := project.DeepCopy()
		project.Spec.Properties.Description = new("Updated description")
		tc.PatchResourceAndWait(oldProject, project)
		tc.Expect(project.Status.Properties.Description).To(Equal(new("Updated description")))

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
						Name:      new("gpt-4o"),
						Format:    new("OpenAI"),
						Publisher: new("OpenAI"),
						Version:   new("2024-11-20"),
					},
				},
				Sku: &cognitiveservices.Sku{
					Capacity: new(1),
					Name:     new("Standard"),
				},
			},
		}

		tc.CreateResourcesAndWait(deployment)
		tc.Expect(deployment.Status.Id).ToNot(BeNil())

		oldDep := deployment.DeepCopy()
		deployment.Spec.Sku.Capacity = new(2)
		tc.PatchResourceAndWait(oldDep, deployment)
		tc.Expect(deployment.Status.Sku.Capacity).To(Equal(new(2)))
		tc.DeleteResourceAndWait(deployment)
		exists, _, err := tc.AzureClient.CheckExistenceWithGetByID(tc.Ctx, *deployment.Status.Id, string(cognitiveservices.APIVersion_Value))
		tc.Expect(err).ToNot(HaveOccurred())
		tc.Expect(exists).To(BeFalse())
	}
}
