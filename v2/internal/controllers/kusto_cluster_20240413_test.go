/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	kusto "github.com/Azure/azure-service-operator/v2/api/kusto/v1api20240413"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
)

func Test_Kusto_Cluster_20240413_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	cluster := &kusto.Cluster{
		ObjectMeta: tc.MakeObjectMeta("cluster"),
		Spec: kusto.Cluster_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			Sku: &kusto.AzureSku{
				Capacity: to.Ptr(2),
				Name:     to.Ptr(kusto.AzureSku_Name_Standard_D13_V2),
				Tier:     to.Ptr(kusto.AzureSku_Tier_Standard),
			},
			Identity: &kusto.Identity{
				Type: to.Ptr(kusto.Identity_Type("SystemAssigned")),
			},
		},
	}

	db := &kusto.Database{
		ObjectMeta: tc.MakeObjectMeta("db"),
		Spec: kusto.Database_Spec{
			Owner: testcommon.AsOwner(cluster),
			ReadWrite: &kusto.ReadWriteDatabase{
				Kind:             to.Ptr(kusto.ReadWriteDatabase_Kind_ReadWrite),
				Location:         tc.AzureRegion,
				SoftDeletePeriod: to.Ptr("P1D"),
			},
		},
	}

	tc.CreateResourcesAndWait(cluster, db)
	tc.Expect(cluster.Status.Identity).ToNot(BeNil())
	tc.Expect(cluster.Status.Identity.PrincipalId).ToNot(BeNil())

	principalId := *cluster.Status.Identity.PrincipalId

	var tenantId *string
	if cluster.Status.Identity.TenantId != nil {
		tenantId = to.Ptr(*cluster.Status.Identity.TenantId)
	}

	tc.RunParallelSubtests(
		testcommon.Subtest{
			Name: "Kusto Cluster Update",
			Test: KustoClusterUpdate(cluster),
		},
		testcommon.Subtest{
			Name: "Kusto Database Update",
			Test: KustoDatabaseUpdate(db),
		},
		testcommon.Subtest{
			Name: "Kusto PrincipalAssignment CRUD",
			Test: KustoPrincipalAssignmentCRUD(db, principalId, tenantId),
		},
	)

	tc.DeleteResourceAndWait(cluster)
	exists, _, err := tc.AzureClient.CheckExistenceWithGetByID(tc.Ctx, *cluster.Status.Id, string(kusto.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeFalse())
}

func KustoClusterUpdate(cluster *kusto.Cluster) func(tc *testcommon.KubePerTestContext) {
	return func(tc *testcommon.KubePerTestContext) {
		old := cluster.DeepCopy()
		cluster.Spec.EnablePurge = to.Ptr(true)
		tc.PatchResourceAndWait(old, cluster)
		tc.Expect(cluster.Status.EnablePurge).ToNot(BeNil())
		tc.Expect(*cluster.Status.EnablePurge).To(BeTrue())
	}
}

func KustoDatabaseUpdate(db *kusto.Database) func(tc *testcommon.KubePerTestContext) {
	return func(tc *testcommon.KubePerTestContext) {
		old := db.DeepCopy()
		db.Spec.ReadWrite.SoftDeletePeriod = to.Ptr("P3D")
		tc.PatchResourceAndWait(old, db)
		tc.Expect(db.Status.ReadWrite).ToNot(BeNil())
		tc.Expect(db.Status.ReadWrite.SoftDeletePeriod).To(Equal(to.Ptr("P3D")))
	}
}

func KustoPrincipalAssignmentCRUD(db *kusto.Database, principalID string, tenantID *string) func(tc *testcommon.KubePerTestContext) {
	return func(tc *testcommon.KubePerTestContext) {
		pa := &kusto.PrincipalAssignment{
			ObjectMeta: tc.MakeObjectMeta("priass"),
			Spec: kusto.PrincipalAssignment_Spec{
				Owner:         testcommon.AsOwner(db),
				PrincipalId:   to.Ptr(principalID),
				PrincipalType: to.Ptr(kusto.DatabasePrincipalProperties_PrincipalType("App")),
				Role:          to.Ptr(kusto.DatabasePrincipalProperties_Role("Admin")),
				TenantId:      tenantID,
			},
		}

		tc.CreateResourcesAndWait(pa)
		tc.Expect(pa.Status.Id).ToNot(BeNil())

		old := pa.DeepCopy()
		pa.Spec.Role = to.Ptr(kusto.DatabasePrincipalProperties_Role("Viewer"))
		tc.PatchResourceAndWait(old, pa)
		tc.Expect(pa.Status.Role).ToNot(BeNil())
		tc.Expect(string(*pa.Status.Role)).To(Equal("Viewer"))

		armID := *pa.Status.Id
		tc.DeleteResourceAndWait(pa)
		exists, retryAfter, err := tc.AzureClient.CheckExistenceWithGetByID(tc.Ctx, armID, string(kusto.APIVersion_Value))
		tc.Expect(err).ToNot(HaveOccurred())
		tc.Expect(retryAfter).To(BeZero())
		tc.Expect(exists).To(BeFalse())
	}
}
