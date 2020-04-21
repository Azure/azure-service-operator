// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

// +build all azuresqlserver

package controllers

import (
	"context"
	"testing"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestAzureSqlServerControllerNoResourceGroup(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()

	sqlServerName := GenerateTestResourceNameWithRandom("sqlserver-dev", 10)

	// Create the SqlServer object and expect the Reconcile to be created
	sqlServerInstance := &azurev1alpha1.AzureSqlServer{
		ObjectMeta: metav1.ObjectMeta{
			Name:      sqlServerName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.AzureSqlServerSpec{
			Location:      tc.resourceGroupLocation,
			ResourceGroup: GenerateTestResourceNameWithRandom("rg-fake-dev", 10),
		},
	}

	EnsureInstanceWithResult(ctx, t, tc, sqlServerInstance, errhelp.ResourceGroupNotFoundErrorCode, false)

	EnsureDelete(ctx, t, tc, sqlServerInstance)
}
