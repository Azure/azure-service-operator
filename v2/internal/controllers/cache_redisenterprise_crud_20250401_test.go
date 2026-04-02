/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"os"
	"testing"

	"github.com/google/uuid"
	. "github.com/onsi/gomega"

	"sigs.k8s.io/controller-runtime/pkg/client"

	cache "github.com/Azure/azure-service-operator/v2/api/cache/v1api20250401"
	cache20250401 "github.com/Azure/azure-service-operator/v2/api/cache/v20250401"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/core"
)

const (
	redisEnterpriseAssignmentObjectIDEnvVar   = "TEST_REDIS_ENTERPRISE_ASSIGNMENT_OBJECT_ID"
	redisEnterpriseAssignmentObjectIDSentinel = "00000000-0000-0000-0000-000000000000"
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

	// Run subtests sequentially because the assignment scenario also uses a database named "default"
	// beneath the same Redis Enterprise owner.
	tc.RunSubtests(
		testcommon.Subtest{
			Name: "RedisEnterprise database CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				RedisEnterprise_Database_20250401_CRUD(tc, &redis)
			},
		},
		testcommon.Subtest{
			Name: "RedisEnterprise database access policy assignment CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				RedisEnterprise_Database_AccessPolicyAssignment_20250401_CRUD(tc, &redis)
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

func RedisEnterprise_Database_20250401_CRUD(tc *testcommon.KubePerTestContext, redis *cache.RedisEnterprise) {
	// Create a redis database on the newly-created cluster
	secretName := "redissecret"

	db := cache.RedisEnterpriseDatabase{
		ObjectMeta: tc.MakeObjectMeta("redisdb"),
		Spec: cache.RedisEnterpriseDatabase_Spec{
			Owner:            testcommon.AsOwner(redis),
			AzureName:        "default",
			ClusteringPolicy: to.Ptr(cache.DatabaseProperties_ClusteringPolicy_OSSCluster),
			OperatorSpec: &cache.RedisEnterpriseDatabaseOperatorSpec{
				SecretExpressions: []*core.DestinationExpression{
					{
						Name:  secretName,
						Key:   "primaryKey",
						Value: "secret.primaryKey",
					},
				},
				Secrets: &cache.RedisEnterpriseDatabaseOperatorSecrets{
					SecondaryKey: &genruntime.SecretDestination{
						Name: secretName,
						Key:  "secondaryKey",
					},
				},
			},
		},
	}

	tc.CreateResourceAndWait(&db)
	tc.Expect(db.Status.Id).ToNot(BeNil())
	tc.Expect(to.Value(db.Status.Name)).To(Equal("default"))
	tc.Expect(to.Value(db.Status.ClusteringPolicy)).To(Equal(cache.DatabaseProperties_ClusteringPolicy_STATUS_OSSCluster))
	armId := *db.Status.Id

	tc.ExpectSecretHasKeys(secretName, "primaryKey", "secondaryKey")

	tc.DeleteResourceAndWait(&db)
	exists, retryAfter, err := tc.AzureClient.CheckExistenceWithGetByID(tc.Ctx, armId, string(cache.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())
}

func RedisEnterprise_Database_AccessPolicyAssignment_20250401_CRUD(tc *testcommon.KubePerTestContext, redis *cache.RedisEnterprise) {
	assignmentObjectID := getRedisEnterpriseAssignmentObjectID(tc)
	assignmentName := tc.NoSpaceNamer.GenerateName("assign")
	accessKeysAuthenticationDisabled := cache.DatabaseProperties_AccessKeysAuthentication_Disabled

	db := cache.RedisEnterpriseDatabase{
		ObjectMeta: tc.MakeObjectMeta("redisdbassign"),
		Spec: cache.RedisEnterpriseDatabase_Spec{
			Owner:                    testcommon.AsOwner(redis),
			AzureName:                "default",
			ClusteringPolicy:         to.Ptr(cache.DatabaseProperties_ClusteringPolicy_OSSCluster),
			AccessKeysAuthentication: &accessKeysAuthenticationDisabled,
		},
	}

	assignment := cache20250401.RedisEnterpriseDatabaseAccessPolicyAssignment{
		ObjectMeta: tc.MakeObjectMetaWithName(assignmentName),
		Spec: cache20250401.RedisEnterpriseDatabaseAccessPolicyAssignment_Spec{
			Owner:            testcommon.AsOwner(&db),
			AccessPolicyName: to.Ptr("default"),
			User: &cache20250401.AccessPolicyAssignmentProperties_User{
				ObjectId: to.Ptr(assignmentObjectID),
			},
		},
	}

	tc.CreateResourcesAndWait(&db, &assignment)
	defer tc.DeleteResourceAndWait(&db)

	tc.Expect(db.Status.Id).ToNot(BeNil())
	tc.Expect(db.Status.AccessKeysAuthentication).ToNot(BeNil())
	tc.Expect(*db.Status.AccessKeysAuthentication).To(Equal(cache.DatabaseProperties_AccessKeysAuthentication_STATUS_Disabled))

	tc.Expect(assignment.Status.Id).ToNot(BeNil())
	tc.Expect(assignment.Status.AccessPolicyName).ToNot(BeNil())
	tc.Expect(*assignment.Status.AccessPolicyName).To(Equal("default"))
	tc.Expect(assignment.Status.User).ToNot(BeNil())
	tc.Expect(assignment.Status.User.ObjectId).ToNot(BeNil())
	if tc.AzureClientRecorder.IsReplaying() {
		tc.Expect(*assignment.Status.User.ObjectId).To(Equal(redisEnterpriseAssignmentObjectIDSentinel))
	} else {
		tc.Expect(*assignment.Status.User.ObjectId).To(Equal(assignmentObjectID))
	}
	assignmentARMID := *assignment.Status.Id

	tc.DeleteResourceAndWait(&assignment)
	exists, retryAfter, err := tc.AzureClient.CheckExistenceWithGetByID(tc.Ctx, assignmentARMID, string(cache20250401.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())
}

func getRedisEnterpriseAssignmentObjectID(tc *testcommon.KubePerTestContext) string {
	if tc.AzureClientRecorder.IsReplaying() {
		return redisEnterpriseAssignmentObjectIDSentinel
	}

	assignmentObjectID := os.Getenv(redisEnterpriseAssignmentObjectIDEnvVar)
	if assignmentObjectID == "" {
		tc.T.Skipf("%s must be set to a real Microsoft Entra object ID when recording this test", redisEnterpriseAssignmentObjectIDEnvVar)
	}

	parsed, err := uuid.Parse(assignmentObjectID)
	if err != nil {
		tc.T.Fatalf("%s must be a valid GUID, got %q: %s", redisEnterpriseAssignmentObjectIDEnvVar, assignmentObjectID, err)
	}

	if parsed == uuid.Nil {
		tc.T.Fatalf("%s must not use the placeholder object ID %s", redisEnterpriseAssignmentObjectIDEnvVar, redisEnterpriseAssignmentObjectIDSentinel)
	}

	tc.WithLiteralRedaction(assignmentObjectID, redisEnterpriseAssignmentObjectIDSentinel)

	return assignmentObjectID
}
