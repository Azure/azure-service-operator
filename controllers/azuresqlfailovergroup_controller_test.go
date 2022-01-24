// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

//go:build all || azuresqlserver || fog
// +build all azuresqlserver fog

package controllers

import (
	"context"
	"testing"

	"github.com/Azure/azure-service-operator/api/v1beta1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestAzureSqlFailoverGroupControllerNoResourceGroup(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()

	var rgName string
	var rgLocation1 string
	var sqlServerOneName string
	var sqlServerTwoName string
	var sqlDatabaseName string

	// Add any setup steps that needs to be executed before each test
	rgName = tc.resourceGroupName
	rgLocation1 = "westus2"
	sqlServerOneName = GenerateTestResourceNameWithRandom("sqlfog-srvone", 10)
	sqlServerTwoName = GenerateTestResourceNameWithRandom("sqlfog-srvtwo", 10)
	sqlDatabaseName = GenerateTestResourceNameWithRandom("sqldb", 10)
	sqlFailoverGroupName := GenerateTestResourceNameWithRandom("sqlfog-dev", 10)

	// Create the SqlFailoverGroup object and expect the Reconcile to be created
	sqlFailoverGroupInstance := &v1beta1.AzureSqlFailoverGroup{
		ObjectMeta: metav1.ObjectMeta{
			Name:      sqlFailoverGroupName,
			Namespace: "default",
		},
		Spec: v1beta1.AzureSqlFailoverGroupSpec{
			Location:                     rgLocation1,
			ResourceGroup:                GenerateTestResourceNameWithRandom("rg-fake", 10),
			Server:                       sqlServerOneName,
			FailoverPolicy:               v1beta1.FailoverPolicyAutomatic,
			FailoverGracePeriod:          30,
			SecondaryServer:              sqlServerTwoName,
			SecondaryServerResourceGroup: rgName,
			DatabaseList:                 []string{sqlDatabaseName},
		},
	}

	EnsureInstanceWithResult(ctx, t, tc, sqlFailoverGroupInstance, errhelp.ResourceGroupNotFoundErrorCode, false)

	EnsureDelete(ctx, t, tc, sqlFailoverGroupInstance)
}
