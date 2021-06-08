// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

// +build all

package controllers

import (
	"context"
	"testing"
	"time"

	"github.com/Azure/go-autorest/autorest/to"
	"github.com/gobuffalo/envy"
	"github.com/stretchr/testify/require"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"

	v1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/helpers"
)

func TestTargetNamespaces(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()

	// Check that resources in default and watched get reconciled
	// successfully, but ones created in other ones don't.
	rgName := tc.resourceGroupName
	rgLocation := tc.resourceGroupLocation

	newName := func() string {
		return "storageacct" + helpers.RandomString(6)
	}

	createNamespaces(ctx, t, "watched", "unwatched")

	configuredNamespaces := envy.Get("AZURE_TARGET_NAMESPACES", "")

	instance := v1alpha1.StorageAccount{
		ObjectMeta: metav1.ObjectMeta{
			Name:      newName(),
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

	EnsureInstance(ctx, t, tc, &instance)

	// The watched namespace is also reconciled.
	instance2 := instance
	instance2.ObjectMeta = metav1.ObjectMeta{
		Name:      newName(),
		Namespace: "watched",
	}
	EnsureInstance(ctx, t, tc, &instance2)

	// But the unwatched namespace isn't...
	instance3 := instance
	instance3.ObjectMeta = metav1.ObjectMeta{
		Name:      newName(),
		Namespace: "unwatched",
	}
	require := require.New(t)
	err := tc.k8sClient.Create(ctx, &instance3)
	require.Equal(nil, err)

	res, err := meta.Accessor(&instance3)
	require.Equal(nil, err)
	names := types.NamespacedName{Name: res.GetName(), Namespace: res.GetNamespace()}

	gotFinalizer := func() bool {
		err := tc.k8sClient.Get(ctx, names, &instance3)
		require.Equal(nil, err)
		return HasFinalizer(res, finalizerName)
	}

	if configuredNamespaces == "" {
		// The operator should be watching all namespaces.
		require.Eventually(
			gotFinalizer,
			tc.timeoutFast,
			tc.retry,
			"instance in some namespace never got a finalizer",
		)
	} else {
		// We can tell that the resource isn't being reconciled if it
		// never gets a finalizer.
		require.Never(
			gotFinalizer,
			20*time.Second,
			time.Second,
			"instance in unwatched namespace got finalizer",
		)
	}

	EnsureDelete(ctx, t, tc, &instance)
	EnsureDelete(ctx, t, tc, &instance2)
	EnsureDelete(ctx, t, tc, &instance3)
}

func createNamespaces(ctx context.Context, t *testing.T, names ...string) {
	for _, name := range names {
		err := tc.k8sClient.Create(ctx, &corev1.Namespace{
			ObjectMeta: metav1.ObjectMeta{
				Name: name,
			},
		})
		if err != nil {
			t.Fatal(err)
		}
	}
}
