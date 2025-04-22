/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	kusto "github.com/Azure/azure-service-operator/v2/api/kusto/v1api20230815"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
)

func Test_Kusto_Cluster_20230815_CRUD(t *testing.T) {
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
		},
	}

	rwDatabase := &kusto.Database{
		ObjectMeta: tc.MakeObjectMeta("rwdatabase"),
		Spec: kusto.Database_Spec{
			Owner: testcommon.AsOwner(cluster),
			ReadWrite: &kusto.ReadWriteDatabase{
				Kind:     to.Ptr(kusto.ReadWriteDatabase_Kind_ReadWrite),
				Location: tc.AzureRegion,
			},
		},
	}

	tc.CreateResourcesAndWait(cluster, rwDatabase)
	tc.Expect(cluster.Status.Id).ToNot(BeNil())
	armId := *cluster.Status.Id

	old := cluster.DeepCopy()
	cluster.Spec.EnablePurge = to.Ptr(true)
	tc.PatchResourceAndWait(old, cluster)

	tc.DeleteResourceAndWait(cluster)

	// Ensure delete
	exists, retryAfter, err := tc.AzureClient.CheckExistenceWithGetByID(tc.Ctx, armId, string(kusto.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())
}
