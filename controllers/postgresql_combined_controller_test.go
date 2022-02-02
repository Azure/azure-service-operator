// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

//go:build all || psql
// +build all psql

package controllers

import (
	"context"
	"testing"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/api/v1alpha2"
	"github.com/Azure/go-autorest/autorest/to"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestPSQLDatabaseController(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()

	var rgName string
	var rgLocation string
	var postgreSQLServerName string
	var postgreSQLServerInstance *v1alpha2.PostgreSQLServer

	// Add any setup steps that needs to be executed before each test
	rgName = tc.resourceGroupName
	rgLocation = tc.resourceGroupLocation

	postgreSQLServerName = GenerateTestResourceNameWithRandom("psql-srv", 10)

	// Create the PostgreSQLServer object and expect the Reconcile to be created
	postgreSQLServerInstance = &v1alpha2.PostgreSQLServer{
		ObjectMeta: metav1.ObjectMeta{
			Name:      postgreSQLServerName,
			Namespace: "default",
		},
		Spec: v1alpha2.PostgreSQLServerSpec{
			Location:      rgLocation,
			ResourceGroup: rgName,
			CreateMode:    "Default",
			Sku: v1alpha2.AzureDBsSQLSku{
				Name:     "GP_Gen5_4",
				Tier:     v1alpha2.SkuTier("GeneralPurpose"),
				Family:   "Gen5",
				Size:     "51200",
				Capacity: 4,
			},
			ServerVersion:  v1alpha2.ServerVersion("10"),
			SSLEnforcement: v1alpha2.SslEnforcementEnumEnabled,
			StorageProfile: &v1alpha2.PSQLStorageProfile{
				BackupRetentionDays: to.Int32Ptr(10),
				GeoRedundantBackup:  "Disabled",
				StorageMB:           to.Int32Ptr(5120),
				StorageAutogrow:     "Disabled",
			},
		},
	}

	EnsureInstance(ctx, t, tc, postgreSQLServerInstance)

	postgreSQLFirewallRuleName := GenerateTestResourceNameWithRandom("psql-fwrule", 10)

	// Create the PostgreSQLFirewallRule object and expect the Reconcile to be created
	postgreSQLFirewallRuleInstance := &azurev1alpha1.PostgreSQLFirewallRule{
		ObjectMeta: metav1.ObjectMeta{
			Name:      postgreSQLFirewallRuleName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.PostgreSQLFirewallRuleSpec{
			ResourceGroup:  rgName,
			Server:         postgreSQLServerName,
			StartIPAddress: "0.0.0.0",
			EndIPAddress:   "255.255.255.255",
		},
	}

	EnsureInstance(ctx, t, tc, postgreSQLFirewallRuleInstance)

	postgreSQLDatabaseName := GenerateTestResourceNameWithRandom("psql-db", 10)

	// Create the PostgreSQLDatabase object and expect the Reconcile to be created
	postgreSQLDatabaseInstance := &azurev1alpha1.PostgreSQLDatabase{
		ObjectMeta: metav1.ObjectMeta{
			Name:      postgreSQLDatabaseName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.PostgreSQLDatabaseSpec{
			ResourceGroup: rgName,
			Server:        postgreSQLServerName,
		},
	}

	EnsureInstance(ctx, t, tc, postgreSQLDatabaseInstance)

	EnsureDelete(ctx, t, tc, postgreSQLDatabaseInstance)

	EnsureDelete(ctx, t, tc, postgreSQLFirewallRuleInstance)

	EnsureDelete(ctx, t, tc, postgreSQLServerInstance)

}
