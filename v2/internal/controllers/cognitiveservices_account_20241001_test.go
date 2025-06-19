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

func Test_CognitiveServices_Account_CRUD(t *testing.T) {
	t.Parallel()
	tc := globalTestContext.ForTest(t)
	rg := tc.CreateTestResourceGroupAndWait()

	account := &cognitiveservices.Account{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.NoSpaceNamer.GenerateName("cogsvc")),
		Spec: cognitiveservices.Account_Spec{
			Identity: &cognitiveservices.Identity{
				Type: to.Ptr(cognitiveservices.Identity_Type("SystemAssigned")),
			},
			Kind:     to.Ptr("AIServices"),
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

	old := account.DeepCopy()
	if old.Spec.Tags == nil {
		old.Spec.Tags = map[string]string{}
	}
	account.Spec.Tags = map[string]string{"env": "test"}
	tc.PatchResourceAndWait(old, account)
	tc.Expect(account.Status.Tags).To(Equal(map[string]string{"env": "test"}))

	tc.DeleteResourceAndWait(account)
	exists, _, err := tc.AzureClient.CheckExistenceWithGetByID(tc.Ctx, *account.Status.Id, string(cognitiveservices.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeFalse())
}
