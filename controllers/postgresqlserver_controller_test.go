// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

// +build all psqlserver

package controllers

import (
	"context"
	"testing"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
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
	postgreSQLServerInstance := &azurev1alpha1.PostgreSQLServer{
		ObjectMeta: metav1.ObjectMeta{
			Name:      postgreSQLServerName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.PostgreSQLServerSpec{
			Location:      rgLocation,
			ResourceGroup: rgName,
			Sku: azurev1alpha1.AzureDBsSQLSku{
				Name:     "B_Gen5_2",
				Tier:     azurev1alpha1.SkuTier("Basic"),
				Family:   "Gen5",
				Size:     "51200",
				Capacity: 2,
			},
			ServerVersion:  azurev1alpha1.ServerVersion("10"),
			SSLEnforcement: azurev1alpha1.SslEnforcementEnumEnabled,
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

	postgreSQLServerReplicaInstance := &azurev1alpha1.PostgreSQLServer{
		ObjectMeta: metav1.ObjectMeta{
			Name:      postgreSQLServerReplicaName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.PostgreSQLServerSpec{
			Location:      rgLocation,
			ResourceGroup: rgName,
			CreateMode:    "Replica",
		},
	}

	errMessage := "Replica requested but source server unspecified"
	EnsureInstanceWithResult(ctx, t, tc, postgreSQLServerReplicaInstance, errMessage, false)
	EnsureDelete(ctx, t, tc, postgreSQLServerReplicaInstance)

}
