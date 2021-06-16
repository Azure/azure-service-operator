// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

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

	instanceDefault := v1alpha1.StorageAccount{
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

	EnsureInstance(ctx, t, tc, &instanceDefault)

	// The watched namespace is also reconciled.
	instanceWatched := instanceDefault
	instanceWatched.ObjectMeta = metav1.ObjectMeta{
		Name:      newName(),
		Namespace: "watched",
	}
	EnsureInstance(ctx, t, tc, &instanceWatched)

	// But the unwatched namespace isn't...
	instanceUnwatched := instanceDefault
	instanceUnwatched.ObjectMeta = metav1.ObjectMeta{
		Name:      newName(),
		Namespace: "unwatched",
	}
	require := require.New(t)
	err := tc.k8sClient.Create(ctx, &instanceUnwatched)
	require.Equal(nil, err)

	res, err := meta.Accessor(&instanceUnwatched)
	require.Equal(nil, err)
	names := types.NamespacedName{Name: res.GetName(), Namespace: res.GetNamespace()}

	gotFinalizer := func() bool {
		err := tc.k8sClient.Get(ctx, names, &instanceUnwatched)
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

	EnsureDelete(ctx, t, tc, &instanceDefault)
	EnsureDelete(ctx, t, tc, &instanceWatched)
	EnsureDelete(ctx, t, tc, &instanceUnwatched)
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
