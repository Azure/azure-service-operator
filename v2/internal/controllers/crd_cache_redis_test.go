/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	"github.com/Azure/go-autorest/autorest/to"
	. "github.com/onsi/gomega"
	v1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	cache "github.com/Azure/azure-service-operator/v2/api/cache/v1beta20201201"
	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1beta20200601"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func Test_Cache_Redis_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()
	redis1 := makeRedis(tc, rg, "redis1")
	redis2 := makeRedis(tc, rg, "redis2")

	tc.CreateResourcesAndWait(redis1, redis2)

	// It should be created in Kubernetes
	tc.Expect(redis1.Status.Id).ToNot(BeNil())
	tc.Expect(redis2.Status.Id).ToNot(BeNil())
	armId := *redis1.Status.Id

	// Perform a simple patch
	old := redis1.DeepCopy()
	enabled := cache.RedisCreatePropertiesPublicNetworkAccess_Enabled
	redis1.Spec.PublicNetworkAccess = &enabled
	tc.PatchResourceAndWait(old, redis1)
	tc.Expect(redis1.Status.PublicNetworkAccess).ToNot(BeNil())
	tc.Expect(string(*redis1.Status.PublicNetworkAccess)).To(Equal(string(enabled)))

	// Updating firewall rules and setting up linked servers at the
	// same time seems to wreak havoc (or at least make things take a
	// lot longer), so test firewall rules synchronously first -
	// they're quick when not clobbering links.
	tc.RunSubtests(
		testcommon.Subtest{
			Name: "Redis firewall rule CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				Redis_FirewallRule_CRUD(tc, redis1)
			},
		})

	tc.RunParallelSubtests(
		testcommon.Subtest{
			Name: "Redis linked server CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				Redis_LinkedServer_CRUD(tc, rg, redis1, redis2)
			},
		},
		testcommon.Subtest{
			Name: "Redis patch schedule CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				Redis_PatchSchedule_CRUD(tc, redis1)
			},
		},
	)

	tc.DeleteResourcesAndWait(redis1, redis2)

	// Ensure that the resource was really deleted in Azure
	exists, retryAfter, err := tc.AzureClient.HeadByID(tc.Ctx, armId, string(cache.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())
}

func makeRedis(tc *testcommon.KubePerTestContext, rg *resources.ResourceGroup, prefix string) *cache.Redis {
	tls12 := cache.RedisCreatePropertiesMinimumTlsVersion_12
	family := cache.SkuFamily_P
	sku := cache.SkuName_Premium
	return &cache.Redis{
		ObjectMeta: tc.MakeObjectMeta(prefix),
		Spec: cache.Redis_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			Sku: &cache.Sku{
				Family:   &family,
				Name:     &sku,
				Capacity: to.IntPtr(1),
			},
			EnableNonSslPort:  to.BoolPtr(false),
			MinimumTlsVersion: &tls12,
			RedisConfiguration: map[string]string{
				"maxmemory-delta":  "10",
				"maxmemory-policy": "allkeys-lru",
			},
			RedisVersion: to.StringPtr("6"),
		},
	}
}

func Redis_LinkedServer_CRUD(tc *testcommon.KubePerTestContext, rg *resources.ResourceGroup, redis1, redis2 *cache.Redis) {
	// Interesting - the link needs to have the same name as the
	// secondary server.
	serverRole := cache.RedisLinkedServerCreatePropertiesServerRole_Secondary
	linkedServer := cache.RedisLinkedServer{
		ObjectMeta: tc.MakeObjectMetaWithName(redis2.ObjectMeta.Name),
		Spec: cache.RedisLinkedServers_Spec{
			Owner:                     testcommon.AsOwner(redis1),
			LinkedRedisCacheLocation:  tc.AzureRegion,
			LinkedRedisCacheReference: tc.MakeReferenceFromResource(redis2),
			ServerRole:                &serverRole,
		},
	}

	tc.CreateResourceAndWait(&linkedServer)
	defer tc.DeleteResourceAndWait(&linkedServer)
	tc.Expect(linkedServer.Status.Id).ToNot(BeNil())

	// Linked servers can't be updated.
}

func Redis_PatchSchedule_CRUD(tc *testcommon.KubePerTestContext, redis *cache.Redis) {
	monday := cache.ScheduleEntryDayOfWeek_Monday
	schedule := cache.RedisPatchSchedule{
		ObjectMeta: tc.MakeObjectMeta("patchsched"),
		Spec: cache.RedisPatchSchedules_Spec{
			Owner: testcommon.AsOwner(redis),
			ScheduleEntries: []cache.ScheduleEntry{{
				DayOfWeek:         &monday,
				MaintenanceWindow: to.StringPtr("PT6H"),
				StartHourUtc:      to.IntPtr(6),
			}},
		},
	}
	tc.CreateResourceAndWait(&schedule)
	tc.Expect(schedule.Status.Id).ToNot(BeNil())

	wednesday := cache.ScheduleEntryDayOfWeek_Wednesday
	old := schedule.DeepCopy()
	schedule.Spec.ScheduleEntries = append(schedule.Spec.ScheduleEntries, cache.ScheduleEntry{
		DayOfWeek:         &wednesday,
		MaintenanceWindow: to.StringPtr("PT6H30S"),
		StartHourUtc:      to.IntPtr(7),
	})
	tc.PatchResourceAndWait(old, &schedule)
	statusMonday := cache.ScheduleEntrySTATUSDayOfWeek_Monday
	statusWednesday := cache.ScheduleEntrySTATUSDayOfWeek_Wednesday
	tc.Expect(schedule.Status.ScheduleEntries).To(Equal([]cache.ScheduleEntry_STATUS{{
		DayOfWeek:         &statusMonday,
		MaintenanceWindow: to.StringPtr("PT6H"),
		StartHourUtc:      to.IntPtr(6),
	}, {
		DayOfWeek:         &statusWednesday,
		MaintenanceWindow: to.StringPtr("PT6H30S"),
		StartHourUtc:      to.IntPtr(7),
	}}))
	// The patch schedule is always named default in Azure whatever we
	// call it in k8s, and can't be deleted once it's created.
}

func Redis_FirewallRule_CRUD(tc *testcommon.KubePerTestContext, redis *cache.Redis) {
	// The RP doesn't like rules with hyphens in the name.
	rule := cache.RedisFirewallRule{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.NoSpaceNamer.GenerateName("fwrule")),
		Spec: cache.RedisFirewallRules_Spec{
			Owner:   testcommon.AsOwner(redis),
			StartIP: to.StringPtr("1.2.3.4"),
			EndIP:   to.StringPtr("1.2.3.4"),
		},
	}

	tc.CreateResourceAndWait(&rule)
	defer tc.DeleteResourceAndWait(&rule)

	old := rule.DeepCopy()
	rule.Spec.EndIP = to.StringPtr("1.2.3.5")
	tc.PatchResourceAndWait(old, &rule)
	tc.Expect(rule.Status.EndIP).ToNot(BeNil())
	tc.Expect(*rule.Status.EndIP).To(Equal("1.2.3.5"))
	tc.Expect(rule.Status.Id).ToNot(BeNil())
}

func Test_Cache_Redis_SecretsFromAzure(t *testing.T) {
	t.Parallel()
	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()
	redis := makeRedis(tc, rg, "redis")

	tc.CreateResourceAndWait(redis)

	// There should be no secrets at this point
	list := &v1.SecretList{}
	tc.ListResources(list, client.InNamespace(tc.Namespace))
	tc.Expect(list.Items).To(HaveLen(0))

	// Run sub-tests on the redis
	tc.RunSubtests(
		testcommon.Subtest{
			Name: "SecretsWrittenToSameKubeSecret",
			Test: func(tc *testcommon.KubePerTestContext) {
				Redis_SecretsWrittenToSameKubeSecret(tc, redis)
			},
		},
		testcommon.Subtest{
			Name: "SecretsWrittenToDifferentKubeSecrets",
			Test: func(tc *testcommon.KubePerTestContext) {
				Redis_SecretsWrittenToDifferentKubeSecrets(tc, redis)
			},
		},
	)
}

func Redis_SecretsWrittenToSameKubeSecret(tc *testcommon.KubePerTestContext, redis *cache.Redis) {
	old := redis.DeepCopy()
	redisSecret := "storagekeys"
	redis.Spec.OperatorSpec = &cache.RedisOperatorSpec{
		Secrets: &cache.RedisOperatorSecrets{
			PrimaryKey: &genruntime.SecretDestination{
				Name: redisSecret,
				Key:  "primarykey",
			},
			HostName: &genruntime.SecretDestination{
				Name: redisSecret,
				Key:  "hostname",
			},
			SSLPort: &genruntime.SecretDestination{
				Name: redisSecret,
				Key:  "sslport",
			},
		},
	}
	tc.PatchResourceAndWait(old, redis)

	tc.ExpectSecretHasKeys(redisSecret, "primarykey", "hostname", "sslport")
}

func Redis_SecretsWrittenToDifferentKubeSecrets(tc *testcommon.KubePerTestContext, redis *cache.Redis) {
	old := redis.DeepCopy()
	primaryKeySecret := "secret1"
	secondaryKeySecret := "secret2"
	hostnameSecret := "secret3"
	sslPortSecret := "secret4"

	// Not testing port as it's not returned by default so won't be written anyway

	redis.Spec.OperatorSpec = &cache.RedisOperatorSpec{
		Secrets: &cache.RedisOperatorSecrets{
			PrimaryKey: &genruntime.SecretDestination{
				Name: primaryKeySecret,
				Key:  "primarykey",
			},
			SecondaryKey: &genruntime.SecretDestination{
				Name: secondaryKeySecret,
				Key:  "secondarykey",
			},
			HostName: &genruntime.SecretDestination{
				Name: hostnameSecret,
				Key:  "hostname",
			},
			SSLPort: &genruntime.SecretDestination{
				Name: sslPortSecret,
				Key:  "sslport",
			},
		},
	}
	tc.PatchResourceAndWait(old, redis)

	tc.ExpectSecretHasKeys(primaryKeySecret, "primarykey")
	tc.ExpectSecretHasKeys(secondaryKeySecret, "secondarykey")
	tc.ExpectSecretHasKeys(hostnameSecret, "hostname")
	tc.ExpectSecretHasKeys(sslPortSecret, "sslport")
}
