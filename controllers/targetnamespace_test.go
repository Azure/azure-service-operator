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
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

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

	require := require.New(t)

	// Check that the instance is annotated with the operator namespace.
	checkNamespaceAnnotation(require, &instanceDefault, "azureoperator-system")

	// The watched namespace is also reconciled.
	instanceWatched := instanceDefault
	instanceWatched.ObjectMeta = metav1.ObjectMeta{
		Name:      newName(),
		Namespace: "watched",
	}
	EnsureInstance(ctx, t, tc, &instanceWatched)

	checkNamespaceAnnotation(require, &instanceWatched, "azureoperator-system")

	// But the unwatched namespace isn't...
	unwatchedName := newName()
	instanceUnwatched := instanceDefault
	instanceUnwatched.ObjectMeta = metav1.ObjectMeta{
		Name:      unwatchedName,
		Namespace: "unwatched",
	}
	err := tc.k8sClient.Create(ctx, &instanceUnwatched)
	require.Equal(nil, err)

	names := types.NamespacedName{Name: unwatchedName, Namespace: "unwatched"}

	gotFinalizer := func() bool {
		var instance v1alpha1.StorageAccount
		err := tc.k8sClient.Get(ctx, names, &instance)
		require.Equal(nil, err)
		res, err := meta.Accessor(&instance)
		require.Equal(nil, err)
		return HasFinalizer(res, finalizerName)
	}

	gotNamespaceAnnotation := func() bool {
		var instance v1alpha1.StorageAccount
		err := tc.k8sClient.Get(ctx, names, &instance)
		require.Equal(nil, err)
		res, err := meta.Accessor(&instance)
		require.Equal(nil, err)
		return res.GetAnnotations()[namespaceAnnotation] == "azureoperator-system"
	}

	if configuredNamespaces == "" {
		// The operator should be watching all namespaces.
		require.Eventually(
			gotFinalizer,
			tc.timeoutFast,
			tc.retry,
			"instance in some namespace never got a finalizer",
		)
		// And there should also be a namespace annotation.
		require.Eventually(
			gotNamespaceAnnotation,
			tc.timeoutFast,
			tc.retry,
			"instance in some namespace never got an operator namespace annotation",
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
		// There also shouldn't be a namespace annotation.
		checkNoNamespaceAnnotation(require, &instanceUnwatched)
	}

	EnsureDelete(ctx, t, tc, &instanceDefault)
	EnsureDelete(ctx, t, tc, &instanceWatched)
	EnsureDelete(ctx, t, tc, &instanceUnwatched)
}

func TestOperatorNamespacePreventsReconciling(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()

	// If a resource has a different operator's namespace it won't be
	// reconciled.
	notMine := v1alpha1.StorageAccount{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "storageacct" + helpers.RandomString(6),
			Namespace: "default",
			Annotations: map[string]string{
				namespaceAnnotation: "hard-times",
			},
		},
		Spec: v1alpha1.StorageAccountSpec{
			Kind:          "BlobStorage",
			Location:      tc.resourceGroupLocation,
			ResourceGroup: tc.resourceGroupName,
			Sku: v1alpha1.StorageAccountSku{
				Name: "Standard_LRS",
			},
			AccessTier:             "Hot",
			EnableHTTPSTrafficOnly: to.BoolPtr(true),
		},
	}

	require := require.New(t)
	err := tc.k8sClient.Create(ctx, &notMine)
	require.Equal(nil, err)
	defer EnsureDelete(ctx, t, tc, &notMine)

	names := types.NamespacedName{
		Name:      notMine.ObjectMeta.Name,
		Namespace: "default",
	}

	gotFinalizer := func() bool {
		var instance v1alpha1.StorageAccount
		err := tc.k8sClient.Get(ctx, names, &instance)
		require.Equal(nil, err)
		res, err := meta.Accessor(&instance)
		require.Equal(nil, err)
		return HasFinalizer(res, finalizerName)
	}

	require.Never(
		gotFinalizer,
		20*time.Second,
		time.Second,
		"instance claimed by some other operator got finalizer",
	)

	var events corev1.EventList
	err = tc.k8sClient.List(ctx, &events, &client.ListOptions{
		FieldSelector: fields.ParseSelectorOrDie("involvedObject.name=" + notMine.ObjectMeta.Name),
		Namespace:     "default",
	})
	require.Equal(nil, err)
	require.Len(events.Items, 1)
	event := events.Items[0]
	require.Equal(event.Type, "Warning")
	require.Equal(event.Reason, "Overlap")
	require.Equal(event.Message, `Operators in "azureoperator-system" and "hard-times" are both configured to manage this resource`)

	// But an instance that I've claimed gets reconciled fine.
	mine := notMine
	mine.ObjectMeta = metav1.ObjectMeta{
		Name:      "storaceacct" + helpers.RandomString(6),
		Namespace: "default",
		Annotations: map[string]string{
			namespaceAnnotation: "azureoperator-system",
		},
	}
	EnsureInstance(ctx, t, tc, &mine)
	EnsureDelete(ctx, t, tc, &mine)
}

func checkNoNamespaceAnnotation(r *require.Assertions, instance metav1.Object) {
	res, err := meta.Accessor(instance)
	r.Equal(nil, err)
	_, found := res.GetAnnotations()[namespaceAnnotation]
	r.Equal(false, found)
}

func checkNamespaceAnnotation(r *require.Assertions, instance metav1.Object, expected string) {
	res, err := meta.Accessor(instance)
	r.Equal(nil, err)
	actual, found := res.GetAnnotations()[namespaceAnnotation]
	r.Equal(true, found)
	r.Equal(expected, actual)
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
