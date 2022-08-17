/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	"github.com/Azure/go-autorest/autorest/to"
	. "github.com/onsi/gomega"

	cache "github.com/Azure/azure-service-operator/v2/api/cache/v1beta20210301"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
)

func Test_Cache_RedisEnterprise_CRUD(t *testing.T) {
	t.Parallel()
	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()
	tls12 := cache.ClusterPropertiesMinimumTlsVersion_12
	sku := cache.SkuName_EnterpriseE10
	redis := cache.RedisEnterprise{
		ObjectMeta: tc.MakeObjectMeta("redisent"),
		Spec: cache.RedisEnterprise_Spec{
			Location:          tc.AzureRegion,
			Owner:             testcommon.AsOwner(rg),
			MinimumTlsVersion: &tls12,
			Sku: &cache.Sku{
				Capacity: to.IntPtr(2),
				Name:     &sku,
			},
			Tags: map[string]string{
				"elks": "stranger",
			},
		},
	}

	tc.CreateResourceAndWait(&redis)
	tc.Expect(redis.Status.Id).ToNot(BeNil())
	armId := *redis.Status.Id

	// TODO(babbageclunk): It seems like this isn't working because
	// the RP expects updates to be done using PATCH but the operator
	// issues PUTs. I've reached out to the Azure Redis team.

	// old := redis.DeepCopy()
	// redis.Spec.Tags["nomai"] = "vessel"
	// tc.Patch(old, &redis)

	// objectKey := client.ObjectKeyFromObject(&redis)

	// // Ensure state got updated in Azure.
	// tc.Eventually(func() map[string]string {
	// 	var updated cache.RedisEnterprise
	// 	tc.GetResource(objectKey, &updated)
	// 	tc.T.Log(pretty.Sprint(updated.Status.Tags))
	// 	return updated.Status.Tags
	// }).Should(Equal(map[string]string{
	// 	"elks":  "stranger",
	// 	"nomai": "vessel",
	// }))

	tc.RunParallelSubtests(testcommon.Subtest{
		Name: "RedisEnterprise database CRUD",
		Test: func(tc *testcommon.KubePerTestContext) {
			RedisEnterprise_Database_CRUD(tc, &redis)
		},
	})

	tc.DeleteResourceAndWait(&redis)

	// Ensure that the resource was really deleted in Azure
	exists, retryAfter, err := tc.AzureClient.HeadByID(tc.Ctx, armId, string(cache.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())
}

func RedisEnterprise_Database_CRUD(tc *testcommon.KubePerTestContext, redis *cache.RedisEnterprise) {
	encrypted := cache.DatabasePropertiesClientProtocol_Encrypted
	enterpriseCluster := cache.DatabasePropertiesClusteringPolicy_EnterpriseCluster
	allKeysLRU := cache.DatabasePropertiesEvictionPolicy_AllKeysLRU
	always := cache.PersistenceAofFrequency_Always

	database := cache.RedisEnterpriseDatabase{
		// The RP currently only allows one database, which must be
		// named "default", in a cluster.
		ObjectMeta: tc.MakeObjectMetaWithName("default"),
		Spec: cache.RedisEnterpriseDatabases_Spec{
			Owner:            testcommon.AsOwner(redis),
			ClientProtocol:   &encrypted,
			ClusteringPolicy: &enterpriseCluster,
			EvictionPolicy:   &allKeysLRU,
			Modules: []cache.Module{{
				Name: to.StringPtr("RedisBloom"),
				Args: to.StringPtr("ERROR_RATE 0.1 INITIAL_SIZE 400"),
			}},
			Persistence: &cache.Persistence{
				AofEnabled:   to.BoolPtr(true),
				AofFrequency: &always,
				RdbEnabled:   to.BoolPtr(false),
			},
			// Port is required to be 10000 at the moment.
			Port: to.IntPtr(10000),
		},
	}

	tc.CreateResourceAndWait(&database)
	defer tc.DeleteResourceAndWait(&database)
	tc.Expect(database.Status.Id).ToNot(BeNil())

	old := database.DeepCopy()
	oneSecond := cache.PersistenceAofFrequency_1S
	database.Spec.Persistence.AofFrequency = &oneSecond
	tc.PatchResourceAndWait(old, &database)

	oneSecondStatus := cache.PersistenceStatusAofFrequency_1S
	expectedPersistenceStatus := &cache.Persistence_Status{
		AofEnabled:   to.BoolPtr(true),
		AofFrequency: &oneSecondStatus,
		RdbEnabled:   to.BoolPtr(false),
	}
	tc.Expect(database.Status.Persistence).To(Equal(expectedPersistenceStatus))
}
