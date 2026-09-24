/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"

	v1 "k8s.io/api/core/v1"

	postgresql "github.com/Azure/azure-service-operator/v2/api/dbforpostgresql/v20250801"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

// A replica must survive being PUT again once it exists, which is what happens on every spec change
// and whenever the operator restarts. See https://github.com/Azure/azure-service-operator/issues/5086.
func Test_DBForPostgreSQL_FlexibleServer_Replica_20250801(t *testing.T) {
	t.Parallel()

	if *isLive {
		t.Skip("can't run in live mode, postresql flexible server takes too long to be provisioned and deleted")
	}

	g := NewGomegaWithT(t)
	ctx := context.Background()
	tc := globalTestContext.ForTest(t)

	// Capacity crunch in West US 2 makes this not work when live
	tc.AzureRegion = to.Ptr("northeurope")

	rg := tc.CreateTestResourceGroupAndWait()

	adminPasswordKey := "adminPassword"
	secret := &v1.Secret{
		ObjectMeta: tc.MakeObjectMeta("postgresqlsecret"),
		StringData: map[string]string{
			adminPasswordKey: tc.Namer.GeneratePassword(),
		},
	}

	tc.CreateResource(secret)

	primary := &postgresql.FlexibleServer{
		ObjectMeta: tc.MakeObjectMeta("postgresql"),
		Spec: postgresql.FlexibleServer_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			Version:  to.Ptr(postgresql.PostgresMajorVersion_17),
			Sku: &postgresql.Sku{
				Name: to.Ptr("Standard_D2ds_v5"),
				Tier: to.Ptr(postgresql.SkuTier_GeneralPurpose),
			},
			AdministratorLogin: to.Ptr("myadmin"),
			AdministratorLoginPassword: &genruntime.SecretReference{
				Name: secret.Name,
				Key:  adminPasswordKey,
			},
			Storage: &postgresql.Storage{
				StorageSizeGB: to.Ptr(32),
			},
		},
	}

	tc.CreateResourceAndWait(primary)

	replica := &postgresql.FlexibleServer{
		ObjectMeta: tc.MakeObjectMeta("replica"),
		Spec: postgresql.FlexibleServer_Spec{
			Location:                      tc.AzureRegion,
			Owner:                         testcommon.AsOwner(rg),
			CreateMode:                    to.Ptr(postgresql.CreateMode_Replica),
			SourceServerResourceReference: tc.MakeReferenceFromResource(primary),
		},
	}

	tc.CreateResourceAndWait(replica)

	g.Expect(replica.Status.Id).ToNot(BeNil())
	g.Expect(replica.Status.ReplicationRole).ToNot(BeNil())
	g.Expect(*replica.Status.ReplicationRole).To(BeElementOf(
		postgresql.ReplicationRole_STATUS_AsyncReplica,
		postgresql.ReplicationRole_STATUS_GeoAsyncReplica,
	))
	armId := *replica.Status.Id

	// Any change to the replica sends it to Azure again, still with createMode Replica in its spec
	tc.LogSectionf("Updating replica %s", replica.Name)
	old := replica.DeepCopy()
	replica.Spec.Tags = map[string]string{
		"updated": "true",
	}
	tc.PatchResourceAndWait(old, replica)
	g.Expect(replica.Status.Tags).To(HaveKeyWithValue("updated", "true"))

	tc.DeleteResourceAndWait(replica)

	// Ensure that the replica was really deleted in Azure
	exists, retryAfter, err := tc.AzureClient.CheckExistenceWithGetByID(ctx, armId, string(postgresql.APIVersion_Value))
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(retryAfter).To(BeZero())
	g.Expect(exists).To(BeFalse())
}
