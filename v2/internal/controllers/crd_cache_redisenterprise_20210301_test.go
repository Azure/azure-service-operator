/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"
	"sigs.k8s.io/controller-runtime/pkg/client"

	cache "github.com/Azure/azure-service-operator/v2/api/cache/v1api20210301"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
)

func Test_Cache_RedisEnterprise_20210301_CRUD(t *testing.T) {
	t.Parallel()
	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()
	tls12 := cache.ClusterProperties_MinimumTlsVersion_12
	sku := cache.Sku_Name_Enterprise_E10
	redis := cache.RedisEnterprise{
		ObjectMeta: tc.MakeObjectMeta("redisent"),
		Spec: cache.RedisEnterprise_Spec{
			Location:          tc.AzureRegion,
			Owner:             testcommon.AsOwner(rg),
			MinimumTlsVersion: &tls12,
			Sku: &cache.Sku{
				Capacity: to.Ptr(2),
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

	old := redis.DeepCopy()
	redis.Spec.Tags["nomai"] = "vessel"
	tc.Patch(old, &redis)

	objectKey := client.ObjectKeyFromObject(&redis)

	// Ensure state got updated in Azure.
	tc.Eventually(func() map[string]string {
		var updated cache.RedisEnterprise
		tc.GetResource(objectKey, &updated)
		return updated.Status.Tags
	}).Should(Equal(map[string]string{
		"elks":  "stranger",
		"nomai": "vessel",
	}))

	tc.RunParallelSubtests(testcommon.Subtest{
		Name: "RedisEnterprise database CRUD",
		Test: func(tc *testcommon.KubePerTestContext) {
			RedisEnterprise_Database_20210301_CRUD(tc, &redis)
		},
	})

	tc.DeleteResourceAndWait(&redis)

	// Ensure that the resource was really deleted in Azure
	exists, retryAfter, err := tc.AzureClient.HeadByID(tc.Ctx, armId, string(cache.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())
}

func RedisEnterprise_Database_20210301_CRUD(tc *testcommon.KubePerTestContext, redis *cache.RedisEnterprise) {
	encrypted := cache.DatabaseProperties_ClientProtocol_Encrypted
	enterpriseCluster := cache.DatabaseProperties_ClusteringPolicy_EnterpriseCluster
	allKeysLRU := cache.DatabaseProperties_EvictionPolicy_AllKeysLRU
	always := cache.Persistence_AofFrequency_Always

	database := cache.RedisEnterpriseDatabase{
		// The RP currently only allows one database, which must be
		// named "default", in a cluster.
		ObjectMeta: tc.MakeObjectMetaWithName("default"),
		Spec: cache.RedisEnterprise_Database_Spec{
			Owner:            testcommon.AsOwner(redis),
			ClientProtocol:   &encrypted,
			ClusteringPolicy: &enterpriseCluster,
			EvictionPolicy:   &allKeysLRU,
			Modules: []cache.Module{{
				Name: to.Ptr("RedisBloom"),
				Args: to.Ptr("ERROR_RATE 0.1 INITIAL_SIZE 400"),
			}},
			Persistence: &cache.Persistence{
				AofEnabled:   to.Ptr(true),
				AofFrequency: &always,
				RdbEnabled:   to.Ptr(false),
			},
			// Port is required to be 10000 at the moment.
			Port: to.Ptr(10000),
		},
	}

	tc.CreateResourceAndWait(&database)
	defer tc.DeleteResourceAndWait(&database)
	tc.Expect(database.Status.Id).ToNot(BeNil())

	old := database.DeepCopy()
	oneSecond := cache.Persistence_AofFrequency_1S
	database.Spec.Persistence.AofFrequency = &oneSecond
	tc.PatchResourceAndWait(old, &database)

	oneSecondStatus := cache.Persistence_AofFrequency_STATUS_1S
	expectedPersistenceStatus := &cache.Persistence_STATUS{
		AofEnabled:   to.Ptr(true),
		AofFrequency: &oneSecondStatus,
		RdbEnabled:   to.Ptr(false),
	}
	tc.Expect(database.Status.Persistence).To(Equal(expectedPersistenceStatus))
}
