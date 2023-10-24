/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	cache "github.com/Azure/azure-service-operator/v2/api/cache/v1api20201201"
	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1api20200601"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
)

func Test_Cache_Redis_20201201_CRUD(t *testing.T) {
	t.Parallel()

	if *isLive {
		t.Skip("can't run in live mode, redis server takes too long to be provisioned and deletion")
	}

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()
	redis1 := newRedis20201201(tc, rg, "redis1")
	redis2 := newRedis20201201(tc, rg, "redis2")

	tc.CreateResourcesAndWait(redis1, redis2)

	// It should be created in Kubernetes
	tc.Expect(redis1.Status.Id).ToNot(BeNil())
	tc.Expect(redis2.Status.Id).ToNot(BeNil())
	armId := *redis1.Status.Id

	// Perform a simple patch
	old := redis1.DeepCopy()
	enabled := cache.RedisCreateProperties_PublicNetworkAccess_Enabled
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
				Redis_FirewallRule_20201201_CRUD(tc, redis1)
			},
		})

	tc.RunParallelSubtests(
		testcommon.Subtest{
			Name: "Redis linked server CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				Redis_LinkedServer_20201201_CRUD(tc, rg, redis1, redis2)
			},
		},
		testcommon.Subtest{
			Name: "Redis patch schedule CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				Redis_PatchSchedule_20201201_CRUD(tc, redis1)
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

func newRedis20201201(tc *testcommon.KubePerTestContext, rg *resources.ResourceGroup, prefix string) *cache.Redis {
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

func Redis_LinkedServer_20201201_CRUD(tc *testcommon.KubePerTestContext, _ *resources.ResourceGroup, redis1, redis2 *cache.Redis) {
	// Interesting - the link needs to have the same name as the
	// secondary server.
	serverRole := cache.RedisLinkedServerCreateProperties_ServerRole_Secondary
	linkedServer := cache.RedisLinkedServer{
		ObjectMeta: tc.MakeObjectMetaWithName(redis2.ObjectMeta.Name),
		Spec: cache.Redis_LinkedServer_Spec{
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

func Redis_PatchSchedule_20201201_CRUD(tc *testcommon.KubePerTestContext, redis *cache.Redis) {
	monday := cache.ScheduleEntry_DayOfWeek_Monday
	schedule := cache.RedisPatchSchedule{
		ObjectMeta: tc.MakeObjectMeta("patchsched"),
		Spec: cache.Redis_PatchSchedule_Spec{
			Owner: testcommon.AsOwner(redis),
			ScheduleEntries: []cache.ScheduleEntry{{
				DayOfWeek:         &monday,
				MaintenanceWindow: to.Ptr("PT6H"),
				StartHourUtc:      to.Ptr(6),
			}},
		},
	}
	tc.CreateResourceAndWait(&schedule)
	tc.Expect(schedule.Status.Id).ToNot(BeNil())

	wednesday := cache.ScheduleEntry_DayOfWeek_Wednesday
	old := schedule.DeepCopy()
	schedule.Spec.ScheduleEntries = append(schedule.Spec.ScheduleEntries, cache.ScheduleEntry{
		DayOfWeek:         &wednesday,
		MaintenanceWindow: to.Ptr("PT6H30S"),
		StartHourUtc:      to.Ptr(7),
	})
	tc.PatchResourceAndWait(old, &schedule)
	statusMonday := cache.ScheduleEntry_DayOfWeek_STATUS_Monday
	statusWednesday := cache.ScheduleEntry_DayOfWeek_STATUS_Wednesday
	tc.Expect(schedule.Status.ScheduleEntries).To(Equal([]cache.ScheduleEntry_STATUS{{
		DayOfWeek:         &statusMonday,
		MaintenanceWindow: to.Ptr("PT6H"),
		StartHourUtc:      to.Ptr(6),
	}, {
		DayOfWeek:         &statusWednesday,
		MaintenanceWindow: to.Ptr("PT6H30S"),
		StartHourUtc:      to.Ptr(7),
	}}))
	// The patch schedule is always named default in Azure whatever we
	// call it in k8s, and can't be deleted once it's created.
}

func Redis_FirewallRule_20201201_CRUD(tc *testcommon.KubePerTestContext, redis *cache.Redis) {
	// The RP doesn't like rules with hyphens in the name.
	rule := cache.RedisFirewallRule{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.NoSpaceNamer.GenerateName("fwrule")),
		Spec: cache.Redis_FirewallRule_Spec{
			Owner:   testcommon.AsOwner(redis),
			StartIP: to.Ptr("1.2.3.4"),
			EndIP:   to.Ptr("1.2.3.4"),
		},
	}

	tc.CreateResourceAndWait(&rule)
	defer tc.DeleteResourceAndWait(&rule)

	old := rule.DeepCopy()
	rule.Spec.EndIP = to.Ptr("1.2.3.5")
	tc.PatchResourceAndWait(old, &rule)
	tc.Expect(rule.Status.EndIP).ToNot(BeNil())
	tc.Expect(*rule.Status.EndIP).To(Equal("1.2.3.5"))
	tc.Expect(rule.Status.Id).ToNot(BeNil())
}
