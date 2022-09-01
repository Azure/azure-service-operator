/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"fmt"
	"testing"

	"github.com/Azure/go-autorest/autorest/to"
	. "github.com/onsi/gomega"
	"sigs.k8s.io/controller-runtime/pkg/client"

	managedidentity2018 "github.com/Azure/azure-service-operator/v2/api/managedidentity/v1beta20181130"
	managedidentity2022 "github.com/Azure/azure-service-operator/v2/api/managedidentity/v1beta20220131preview"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
)

func Test_ManagedIdentity_UserAssignedIdentity_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	mi := &managedidentity2018.UserAssignedIdentity{
		ObjectMeta: tc.MakeObjectMeta("mi"),
		Spec: managedidentity2018.UserAssignedIdentities_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
		},
	}

	tc.CreateResourceAndWait(mi)

	tc.Expect(mi.Status.TenantId).ToNot(BeNil())
	tc.Expect(mi.Status.PrincipalId).ToNot(BeNil())
	tc.Expect(mi.Status.Id).ToNot(BeNil())
	armId := *mi.Status.Id

	tc.RunParallelSubtests(
		testcommon.Subtest{
			Name: "Federated Identity Credentials CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				FederatedIdentityCredentials_CRUD(tc, mi)
			},
		},
	)

	// Perform a simple patch
	old := mi.DeepCopy()
	mi.Spec.Tags = map[string]string{
		"foo": "bar",
	}
	tc.PatchResourceAndWait(old, mi)
	tc.Expect(mi.Status.Tags).To(HaveKey("foo"))

	tc.DeleteResourceAndWait(mi)

	// Ensure that the resource group was really deleted in Azure
	exists, retryAfter, err := tc.AzureClient.HeadByID(tc.Ctx, armId, string(managedidentity2018.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())
}

func FederatedIdentityCredentials_CRUD(tc *testcommon.KubePerTestContext, umi *managedidentity2018.UserAssignedIdentity) {

	fic := &managedidentity2022.FederatedIdentityCredential{
		ObjectMeta: tc.MakeObjectMeta("fic"),
		Spec: managedidentity2022.UserAssignedIdentities_FederatedIdentityCredentials_Spec{
			Owner: testcommon.AsOwner(umi),
			// For Workload Identity, Audiences should always be "api://AzureADTokenExchange"
			Audiences: []string{
				"api://AzureADTokenExchange",
			},
			// For Workload Identity, Issuer should be the OIDC endpoint of the cluster. For AKS this will look like
			// https://oidc.prod-aks.azure.com/00000000-0000-0000-0000-00000000000/
			Issuer: to.StringPtr("https://oidc.prod-aks.azure.com/00000000-0000-0000-0000-00000000000/"),
			// For Workload Identity, Subject should always be system:serviceaccount:<namespace>:<serviceaccount>
			Subject: to.StringPtr(fmt.Sprintf("system:serviceaccount:%s:%s", tc.Namespace, "default")),
		},
	}

	tc.CreateResourceAndWait(fic)

	tc.Expect(fic.Status.Id).ToNot(BeNil())
	armId := *fic.Status.Id

	// Update the FIC
	old := fic.DeepCopy()
	fic.Spec.Issuer = to.StringPtr("https://oidc.prod-aks.azure.com/1234/")
	tc.Patch(old, fic)

	objectKey := client.ObjectKeyFromObject(fic)

	// ensure state got updated in Azure
	tc.Eventually(func() *string {
		updated := &managedidentity2022.FederatedIdentityCredential{}
		tc.GetResource(objectKey, updated)
		return updated.Status.Issuer
	}).Should(Equal(to.StringPtr("https://oidc.prod-aks.azure.com/1234/")))

	tc.DeleteResourceAndWait(fic)

	// Ensure that the resource was really deleted in Azure
	exists, retryAfter, err := tc.AzureClient.HeadByID(tc.Ctx, armId, string(managedidentity2022.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())
}
