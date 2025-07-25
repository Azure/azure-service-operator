/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	"sigs.k8s.io/controller-runtime/pkg/client"

	cache "github.com/Azure/azure-service-operator/v2/api/cache/v1api20250401"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
)

func Test_Cache_RedisEnterprise_20250401_CRUD(t *testing.T) {
	t.Parallel()
	tc := globalTestContext.ForTest(t)

	// Create a resource group for the test
	rg := tc.CreateTestResourceGroupAndWait()
	tls12 := cache.ClusterProperties_MinimumTlsVersion_12
	highAvailability := cache.ClusterProperties_HighAvailability_Enabled
	redis := cache.RedisEnterprise{
		ObjectMeta: tc.MakeObjectMeta("redisent"),
		Spec: cache.RedisEnterprise_Spec{
			Location:          tc.AzureRegion,
			Owner:             testcommon.AsOwner(rg),
			MinimumTlsVersion: &tls12,
			HighAvailability:  &highAvailability,
			Sku: &cache.Sku{
				Name: to.Ptr(cache.Sku_Name_Balanced_B0),
			},
			Tags: map[string]string{
				"environment": "test",
				"created-by":  "aso-test",
			},
		},
	}

	tc.CreateResourceAndWait(&redis)
	tc.Expect(redis.Status.Id).ToNot(BeNil())
	armId := *redis.Status.Id

	// Test resource update by modifying tags
	old := redis.DeepCopy()
	redis.Spec.Tags["updated"] = "true"
	tc.Patch(old, &redis)

	objectKey := client.ObjectKeyFromObject(&redis)

	// Ensure state got updated in Azure
	tc.Eventually(func() map[string]string {
		var updated cache.RedisEnterprise
		tc.GetResource(objectKey, &updated)
		return updated.Status.Tags
	}).Should(Equal(map[string]string{
		"environment": "test",
		"created-by":  "aso-test",
		"updated":     "true",
	}))

	// Verify status fields are populated correctly
	tc.Expect(redis.Status.Location).ToNot(BeNil())
	tc.Expect(redis.Status.MinimumTlsVersion).ToNot(BeNil())
	tc.Expect(*redis.Status.MinimumTlsVersion).To(Equal(cache.ClusterProperties_MinimumTlsVersion_STATUS_12))
	tc.Expect(redis.Status.HighAvailability).ToNot(BeNil())
	tc.Expect(*redis.Status.HighAvailability).To(Equal(cache.ClusterProperties_HighAvailability_STATUS_Enabled))
	tc.Expect(redis.Status.Sku).ToNot(BeNil())
	tc.Expect(redis.Status.Sku.Name).ToNot(BeNil())
	tc.Expect(*redis.Status.Sku.Name).To(Equal(cache.Sku_Name_STATUS_Balanced_B0))

	// Run parallel subtests to exercise different aspects of the cluster
	tc.RunParallelSubtests(
		testcommon.Subtest{
			Name: "Database creation",
			Test: func(tc *testcommon.KubePerTestContext) {
				RedisEnterprise_Database_20250401(tc, &redis)
			},
		},
	)

	tc.DeleteResourceAndWait(&redis)

	// Ensure that the resource was really deleted in Azure
	exists, retryAfter, err := tc.AzureClient.CheckExistenceWithGetByID(tc.Ctx, armId, string(cache.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())
}

func RedisEnterprise_Database_20250401(tc *testcommon.KubePerTestContext, redis *cache.RedisEnterprise) {
	// Create a redis database on the newly-created cluster
	objectKey := client.ObjectKeyFromObject(redis)

	var redisCurrent cache.RedisEnterprise
	tc.GetResource(objectKey, &redisCurrent)

	db := cache.RedisEnterpriseDatabase{
		ObjectMeta: tc.MakeObjectMeta("redisdb"),
		Spec: cache.RedisEnterpriseDatabase_Spec{
			Owner:            testcommon.AsOwner(redis),
			AzureName:        "default",
			ClusteringPolicy: to.Ptr(cache.DatabaseProperties_ClusteringPolicy_OSSCluster),
		},
	}

	tc.CreateResourceAndWait(&db)
	tc.Expect(db.Status.Id).ToNot(BeNil())
	tc.Expect(to.Value(db.Status.Name)).To(Equal("default"))
	tc.Expect(to.Value(db.Status.ClusteringPolicy)).To(Equal(cache.DatabaseProperties_ClusteringPolicy_STATUS_OSSCluster))
}
