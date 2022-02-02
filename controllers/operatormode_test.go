// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

//go:build all
// +build all

package controllers

import (
	"context"
	"testing"
	"time"

	"github.com/gobuffalo/envy"
	"github.com/stretchr/testify/require"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"

	"github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/go-autorest/autorest/to"
)

func TestOperatorModeWebhooks(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()
	require := require.New(t)

	operatorMode, err := config.ParseOperatorMode(
		envy.Get("AZURE_OPERATOR_MODE", config.OperatorModeBoth.String()))
	require.Equal(nil, err)

	rgName := tc.resourceGroupName
	sqlServerName := GenerateTestResourceNameWithRandom("sqlserver", 10)
	userName := GenerateTestResourceNameWithRandom("sqluser", 5)

	database := v1alpha1.AzureSQLUser{
		ObjectMeta: metav1.ObjectMeta{
			Name:      userName,
			Namespace: "default",
		},
		Spec: v1alpha1.AzureSQLUserSpec{
			Server:        sqlServerName,
			DbName:        "master",
			ResourceGroup: rgName,
			Roles:         []string{"db_datareader", "db_datawriter"},
		},
	}

	err = tc.k8sClient.Create(ctx, &database)
	require.NotNil(err)
	if operatorMode.IncludesWebhooks() {
		// If we're running in a mode that enables webhooks, a user
		// being created in the master database should be forbidden.
		require.Contains(err.Error(), "'master' is a reserved database name and cannot be used")
	} else {
		// Otherwise we should fail because the webhook isn't
		// registered (in a real multi-operator deployment it would be
		// routed to a different operator running in webhook-only
		// mode).
		require.Contains(err.Error(), `failed calling webhook "vazuresqluser.kb.io"`)
	}
}

func TestOperatorModeWatchers(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()
	require := require.New(t)

	rgName := tc.resourceGroupName
	rgLocation := tc.resourceGroupLocation
	acctName := "storageacct" + helpers.RandomString(6)
	instance := v1alpha1.StorageAccount{
		ObjectMeta: metav1.ObjectMeta{
			Name:      acctName,
			Namespace: "default",
		},
		Spec: v1alpha1.StorageAccountSpec{
			Kind:          "BlobStorage",
			Location:      rgLocation,
			ResourceGroup: rgName,
			Sku: v1alpha1.StorageAccountSku{
				Name: "Standard_LRS",
			},
			AccessTier:             "Hot",
			EnableHTTPSTrafficOnly: to.BoolPtr(true),
		},
	}

	err := tc.k8sClient.Create(ctx, &instance)
	require.Equal(nil, err)
	defer EnsureDelete(ctx, t, tc, &instance)

	operatorMode, err := config.ParseOperatorMode(
		envy.Get("AZURE_OPERATOR_MODE", config.OperatorModeBoth.String()))
	require.Equal(nil, err)

	names := types.NamespacedName{
		Namespace: "default",
		Name:      acctName,
	}
	gotFinalizer := func() bool {
		var instance v1alpha1.StorageAccount
		err := tc.k8sClient.Get(ctx, names, &instance)
		require.Equal(nil, err)
		res, err := meta.Accessor(&instance)
		require.Equal(nil, err)
		return HasFinalizer(res, finalizerName)
	}

	if operatorMode.IncludesWatchers() {
		// The operator should see this account and start reconciling it.
		require.Eventually(
			gotFinalizer,
			tc.timeoutFast,
			tc.retry,
			"instance never got finalizer even though operator mode is %q",
			operatorMode,
		)
	} else {
		// The operator shouldn't have registered a watcher, so there
		// shouldn't be a finalizer.
		require.Never(
			gotFinalizer,
			20*time.Second,
			time.Second,
			"instance got finalizer when operator mode is %q",
			operatorMode,
		)
	}
}
