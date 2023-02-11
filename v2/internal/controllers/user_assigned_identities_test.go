/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	mysql "github.com/Azure/azure-service-operator/v2/api/dbformysql/v1api20210501"
	managedidentity2018 "github.com/Azure/azure-service-operator/v2/api/managedidentity/v1api20181130"
	storage "github.com/Azure/azure-service-operator/v2/api/storage/v1api20210401"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
)

func Test_StorageAccountWithKubernetesReferenceUserAssignedIdentity(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	// Create a resource group
	rg := tc.CreateTestResourceGroupAndWait()

	mi := &managedidentity2018.UserAssignedIdentity{
		ObjectMeta: tc.MakeObjectMeta("mi"),
		Spec: managedidentity2018.UserAssignedIdentity_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
		},
	}

	// Create a storage account
	acct := newStorageAccount(tc, rg)
	userAssigned := storage.Identity_Type_UserAssigned
	acct.Spec.Identity = &storage.Identity{
		Type: &userAssigned,
		UserAssignedIdentities: []storage.UserAssignedIdentityDetails{
			{
				Reference: *tc.MakeReferenceFromResource(mi),
			},
		},
	}

	tc.CreateResourcesAndWait(mi, acct)

	tc.Expect(mi.Status.TenantId).ToNot(BeNil())
	tc.Expect(mi.Status.PrincipalId).ToNot(BeNil())
	tc.Expect(mi.Status.Id).ToNot(BeNil())

	tc.Expect(acct.Status.Identity).ToNot(BeNil())
	tc.Expect(acct.Status.Identity.UserAssignedIdentities).ToNot(BeNil())
	tc.Expect(acct.Status.Identity.UserAssignedIdentities).To(HaveLen(1))
	// There should be a single key in the UserAssignedIdentities, and it should match the mi.Status.Id. Unfortunately mi.Status.Id
	// differs from the map entry by case...
	for armID, details := range acct.Status.Identity.UserAssignedIdentities {
		tc.Expect(armID).To(testcommon.EqualsIgnoreCase(*mi.Status.Id))
		tc.Expect(details.PrincipalId).To(Equal(mi.Status.PrincipalId))
		tc.Expect(details.ClientId).To(Equal(mi.Status.ClientId))
	}
}

func Test_MySQLServerWithKubernetesReferenceUserAssignedIdentity(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	// Create a resource group
	rg := tc.CreateTestResourceGroupAndWait()

	mi := &managedidentity2018.UserAssignedIdentity{
		ObjectMeta: tc.MakeObjectMeta("mi"),
		Spec: managedidentity2018.UserAssignedIdentity_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
		},
	}

	tc.CreateResourceAndWait(mi)

	secretName := "mysqlsecret"
	adminPasswordKey := "adminPassword"
	adminPasswordSecretRef := createPasswordSecret(secretName, adminPasswordKey, tc)

	flexibleServer, _ := newFlexibleServer(tc, rg, adminPasswordSecretRef)
	flexibleServer.Spec.Identity = &mysql.Identity{
		Type: to.Ptr(mysql.Identity_Type_UserAssigned),
		UserAssignedIdentities: []mysql.UserAssignedIdentityDetails{
			{
				Reference: *tc.MakeReferenceFromResource(mi),
			},
		},
	}

	tc.CreateResourcesAndWait(flexibleServer)

	tc.Expect(mi.Status.TenantId).ToNot(BeNil())
	tc.Expect(mi.Status.PrincipalId).ToNot(BeNil())
	tc.Expect(mi.Status.Id).ToNot(BeNil())

	tc.Expect(flexibleServer.Status.Identity).ToNot(BeNil())
	tc.Expect(flexibleServer.Status.Identity.UserAssignedIdentities).ToNot(BeNil())
	tc.Expect(flexibleServer.Status.Identity.UserAssignedIdentities).To(HaveLen(1))
	// There should be a single key in the UserAssignedIdentities, and it should match the mi.Status.Id. Unfortunately mi.Status.Id
	// differs from the map entry by case...
	for armID := range flexibleServer.Status.Identity.UserAssignedIdentities {
		tc.Expect(armID).To(testcommon.EqualsIgnoreCase(*mi.Status.Id))
	}
}
