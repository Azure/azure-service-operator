// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

//go:build all || psqlserver
// +build all psqlserver

package controllers

import (
	"context"
	"testing"

	"github.com/Azure/azure-service-operator/api/v1alpha2"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

//test Postgre SQL server unhappy path
func TestPSQLServerControllerNoResourceGroup(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()

	// Add any setup steps that needs to be executed before each test
	rgName := GenerateTestResourceNameWithRandom("psqlsrv-rg", 10)
	rgLocation := tc.resourceGroupLocation

	postgreSQLServerName := GenerateTestResourceNameWithRandom("psql-srv", 10)

	// Create the PostgreSQLServer object and expect the Reconcile to be created
	postgreSQLServerInstance := &v1alpha2.PostgreSQLServer{
		ObjectMeta: metav1.ObjectMeta{
			Name:      postgreSQLServerName,
			Namespace: "default",
		},
		Spec: v1alpha2.PostgreSQLServerSpec{
			Location:      rgLocation,
			ResourceGroup: rgName,
			Sku: v1alpha2.AzureDBsSQLSku{
				Name:     "B_Gen5_2",
				Tier:     v1alpha2.SkuTier("Basic"),
				Family:   "Gen5",
				Size:     "51200",
				Capacity: 2,
			},
			ServerVersion:  v1alpha2.ServerVersion("10"),
			SSLEnforcement: v1alpha2.SslEnforcementEnumEnabled,
		},
	}
	EnsureInstanceWithResult(ctx, t, tc, postgreSQLServerInstance, errhelp.ResourceGroupNotFoundErrorCode, false)
	EnsureDelete(ctx, t, tc, postgreSQLServerInstance)

}

func TestPSQLServerControllerReplicaNoSourceServer(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()

	// Add any setup steps that needs to be executed before each test
	rgName := tc.resourceGroupName
	rgLocation := tc.resourceGroupLocation

	postgreSQLServerReplicaName := GenerateTestResourceNameWithRandom("psql-rep", 10)

	// Create the PostgreSQL Replica Server object and expect the Reconcile to be created

	postgreSQLServerReplicaInstance := &v1alpha2.PostgreSQLServer{
		ObjectMeta: metav1.ObjectMeta{
			Name:      postgreSQLServerReplicaName,
			Namespace: "default",
		},
		Spec: v1alpha2.PostgreSQLServerSpec{
			Location:      rgLocation,
			ResourceGroup: rgName,
			CreateMode:    "Replica",
		},
	}

	errMessage := "Replica requested but source server unspecified"
	EnsureInstanceWithResult(ctx, t, tc, postgreSQLServerReplicaInstance, errMessage, false)
	EnsureDelete(ctx, t, tc, postgreSQLServerReplicaInstance)

}
