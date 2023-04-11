/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	sql "github.com/Azure/azure-service-operator/v2/api/sql/v1api20211101"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

// Note: We run this as a separate test from the other Azure SQL tests because it requires 2 servers
// See https://github.com/Azure/azure-quickstart-templates/blob/master/quickstarts/microsoft.sql/sql-with-failover-group/azuredeploy.json
func Test_SQL_Server_FailoverGroup_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)
	// Use a different region where we have quota
	tc.AzureRegion = to.Ptr("eastus")

	secondaryRegion := to.Ptr("eastus2")

	secretName := "sqlsecret"
	adminPasswordKey := "adminPassword"
	adminPasswordSecretRef := createPasswordSecret(secretName, adminPasswordKey, tc)

	rg := tc.CreateTestResourceGroupAndWait()

	// Make two servers, one primary and one secondary
	serverPrimary := &sql.Server{
		ObjectMeta: tc.MakeObjectMeta("sqlserverprimary"),
		Spec: sql.Server_Spec{
			Location:                   tc.AzureRegion,
			Owner:                      testcommon.AsOwner(rg),
			AdministratorLogin:         to.Ptr("myadmin"),
			AdministratorLoginPassword: &adminPasswordSecretRef,
			Version:                    to.Ptr("12.0"),
		},
	}

	serverSecondary := &sql.Server{
		ObjectMeta: tc.MakeObjectMeta("sqlserversecondary"),
		Spec: sql.Server_Spec{
			Location:                   secondaryRegion, // Must not be in the same region as the primary server
			Owner:                      testcommon.AsOwner(rg),
			AdministratorLogin:         to.Ptr("myadmin"),
			AdministratorLoginPassword: &adminPasswordSecretRef,
			Version:                    to.Ptr("12.0"),
		},
	}

	tc.CreateResourcesAndWait(serverPrimary, serverSecondary)

	// Make a database to mirror
	db := &sql.ServersDatabase{
		ObjectMeta: tc.MakeObjectMeta("db"),
		Spec: sql.Servers_Database_Spec{
			Owner:     testcommon.AsOwner(serverPrimary),
			Location:  tc.AzureRegion,
			Collation: to.Ptr("SQL_Latin1_General_CP1_CI_AS"),
		},
	}

	tc.CreateResourceAndWait(db)

	automatic := sql.FailoverGroupReadWriteEndpoint_FailoverPolicy_Automatic
	failoverGroup := &sql.ServersFailoverGroup{
		ObjectMeta: tc.MakeObjectMeta("failovergroup"),
		Spec: sql.Servers_FailoverGroup_Spec{
			Owner: testcommon.AsOwner(serverPrimary),
			PartnerServers: []sql.PartnerInfo{
				{
					Reference: tc.MakeReferenceFromResource(serverSecondary),
				},
			},
			DatabasesReferences: []genruntime.ResourceReference{
				*tc.MakeReferenceFromResource(db),
			},
			ReadWriteEndpoint: &sql.FailoverGroupReadWriteEndpoint{
				FailoverPolicy:                         &automatic,
				FailoverWithDataLossGracePeriodMinutes: to.Ptr(60),
			},
		},
	}

	tc.CreateResourceAndWait(failoverGroup)

	tc.Expect(failoverGroup.Status.Id).ToNot(BeNil())
	armId := *failoverGroup.Status.Id

	tc.DeleteResourceAndWait(failoverGroup)
	exists, _, err := tc.AzureClient.HeadByID(
		tc.Ctx,
		armId,
		string(sql.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeFalse())
}
