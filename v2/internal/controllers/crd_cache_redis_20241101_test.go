/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	cache "github.com/Azure/azure-service-operator/v2/api/cache/v1api20241101"
	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1api20200601"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
)

func Test_Cache_Redis_20241101_CRUD(t *testing.T) {
	t.Parallel()

	if *isLive {
		t.Skip("can't run in live mode, redis server takes too long to be provisioned and deletion")
	}

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()
	redis := newRedis20241101(tc, rg, "redis")

	tc.CreateResourcesAndWait(redis)

	// It should be created in Kubernetes
	tc.Expect(redis.Status.Id).ToNot(BeNil())
	armId := *redis.Status.Id

	// Perform a simple patch
	old := redis.DeepCopy()
	enabled := cache.RedisCreateProperties_PublicNetworkAccess_Enabled
	redis.Spec.PublicNetworkAccess = &enabled
	tc.PatchResourceAndWait(old, redis)
	tc.Expect(redis.Status.PublicNetworkAccess).ToNot(BeNil())
	tc.Expect(string(*redis.Status.PublicNetworkAccess)).To(Equal(string(enabled)))

	// Test the new AccessPolicy and AccessPolicyAssignment resources
	tc.RunParallelSubtests(
		testcommon.Subtest{
			Name: "Redis access policy CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				Redis_AccessPolicy_20241101_CRUD(tc, redis)
			},
		},
		testcommon.Subtest{
			Name: "Redis access policy assignment CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				Redis_AccessPolicyAssignment_20241101_CRUD(tc, redis)
			},
		},
	)

	tc.DeleteResourcesAndWait(redis)

	// Ensure that the resource was really deleted in Azure
	exists, retryAfter, err := tc.AzureClient.CheckExistenceWithGetByID(tc.Ctx, armId, string(cache.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())
}

func newRedis20241101(tc *testcommon.KubePerTestContext, rg *resources.ResourceGroup, prefix string) *cache.Redis {
	tls12 := cache.RedisCreateProperties_MinimumTlsVersion_12
	family := cache.Sku_Family_P
	sku := cache.Sku_Name_Premium
	return &cache.Redis{
		ObjectMeta: tc.MakeObjectMeta(prefix),
		Spec: cache.Redis_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			Sku: &cache.Sku{
				Family:   &family,
				Name:     &sku,
				Capacity: to.Ptr(1),
			},
			EnableNonSslPort:  to.Ptr(false),
			MinimumTlsVersion: &tls12,
			RedisConfiguration: &cache.RedisCreateProperties_RedisConfiguration{
				MaxmemoryDelta:  to.Ptr("10"),
				MaxmemoryPolicy: to.Ptr("allkeys-lru"),
			},
			RedisVersion: to.Ptr("6"),
		},
	}
}

func Redis_AccessPolicy_20241101_CRUD(tc *testcommon.KubePerTestContext, redis *cache.Redis) {
	// Create an access policy
	accessPolicy := cache.RedisAccessPolicy{
		ObjectMeta: tc.MakeObjectMeta("testpolicy"),
		Spec: cache.RedisAccessPolicy_Spec{
			Owner:       testcommon.AsOwner(redis),
			Type:        to.Ptr("custom"),
			Permissions: to.Ptr("+get +set"),
		},
	}

	tc.CreateResourceAndWait(&accessPolicy)
	defer tc.DeleteResourceAndWait(&accessPolicy)

	tc.Expect(accessPolicy.Status.Id).ToNot(BeNil())
	tc.Expect(accessPolicy.Status.Type).ToNot(BeNil())
	tc.Expect(*accessPolicy.Status.Type).To(Equal("custom"))
	tc.Expect(accessPolicy.Status.Permissions).ToNot(BeNil())
	tc.Expect(*accessPolicy.Status.Permissions).To(Equal("+get +set"))

	// Update the policy permissions
	old := accessPolicy.DeepCopy()
	accessPolicy.Spec.Permissions = to.Ptr("+get +set +del")
	tc.PatchResourceAndWait(old, &accessPolicy)
	tc.Expect(accessPolicy.Status.Permissions).ToNot(BeNil())
	tc.Expect(*accessPolicy.Status.Permissions).To(Equal("+get +set +del"))
}

func Redis_AccessPolicyAssignment_20241101_CRUD(tc *testcommon.KubePerTestContext, redis *cache.Redis) {
	// First create an access policy that we can assign
	accessPolicy := cache.RedisAccessPolicy{
		ObjectMeta: tc.MakeObjectMeta("assignmentpolicy"),
		Spec: cache.RedisAccessPolicy_Spec{
			Owner:       testcommon.AsOwner(redis),
			Type:        to.Ptr("custom"),
			Permissions: to.Ptr("+get +set"),
		},
	}

	tc.CreateResourceAndWait(&accessPolicy)
	defer tc.DeleteResourceAndWait(&accessPolicy)

	// Create an access policy assignment
	assignment := cache.RedisAccessPolicyAssignment{
		ObjectMeta: tc.MakeObjectMeta("testassignment"),
		Spec: cache.RedisAccessPolicyAssignment_Spec{
			Owner:            testcommon.AsOwner(redis),
			AccessPolicyName: to.Ptr(accessPolicy.Name),
			ObjectId:         to.Ptr("00000000-0000-0000-0000-000000000000"), // Dummy object ID
			ObjectIdAlias:    to.Ptr("test-user"),
		},
	}

	tc.CreateResourceAndWait(&assignment)
	defer tc.DeleteResourceAndWait(&assignment)

	tc.Expect(assignment.Status.Id).ToNot(BeNil())
	tc.Expect(assignment.Status.AccessPolicyName).ToNot(BeNil())
	tc.Expect(*assignment.Status.AccessPolicyName).To(Equal(accessPolicy.Name))
	tc.Expect(assignment.Status.ObjectId).ToNot(BeNil())
	tc.Expect(*assignment.Status.ObjectId).To(Equal("00000000-0000-0000-0000-000000000000"))
	tc.Expect(assignment.Status.ObjectIdAlias).ToNot(BeNil())
	tc.Expect(*assignment.Status.ObjectIdAlias).To(Equal("test-user"))

	// Update the object alias
	old := assignment.DeepCopy()
	assignment.Spec.ObjectIdAlias = to.Ptr("updated-user")
	tc.PatchResourceAndWait(old, &assignment)
	tc.Expect(assignment.Status.ObjectIdAlias).ToNot(BeNil())
	tc.Expect(*assignment.Status.ObjectIdAlias).To(Equal("updated-user"))
}
