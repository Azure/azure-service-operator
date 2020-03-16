// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

// +build all eventhubnamespace

package controllers

import (
	"context"
	"strings"
	"testing"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/stretchr/testify/assert"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

func TestEventHubNamespaceControllerNoResourceGroup(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()
	assert := assert.New(t)

	var rgLocation string
	rgLocation = tc.resourceGroupLocation

	// Add Tests for OpenAPI validation (or additonal CRD features) specified in
	// your API definition.
	// Avoid adding tests for vanilla CRUD operations because they would
	// test Kubernetes API server, which isn't the goal here.

	// setting this rg name tells the mocks to set a proper error
	resourceGroupName := "gone"
	eventhubNamespaceName := GenerateTestResourceNameWithRandom("ns-dev-eh", 10)

	// Create the EventHubNamespace object and expect the Reconcile to be created
	eventhubNamespaceInstance := &azurev1alpha1.EventhubNamespace{
		ObjectMeta: metav1.ObjectMeta{
			Name:      eventhubNamespaceName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.EventhubNamespaceSpec{
			Location:      rgLocation,
			ResourceGroup: resourceGroupName,
		},
	}

	err := tc.k8sClient.Create(ctx, eventhubNamespaceInstance)
	assert.Equal(nil, err, "create eventhubns in k8s")

	eventhubNamespacedName := types.NamespacedName{Name: eventhubNamespaceName, Namespace: "default"}

	assert.Eventually(func() bool {
		_ = tc.k8sClient.Get(ctx, eventhubNamespacedName, eventhubNamespaceInstance)
		return strings.Contains(eventhubNamespaceInstance.Status.Message, errhelp.ResourceGroupNotFoundErrorCode)
	}, tc.timeout, tc.retry, "wait for eventhubns to have no rg error")

	err = tc.k8sClient.Delete(ctx, eventhubNamespaceInstance)
	assert.Equal(nil, err, "delete eventhubns in k8s")

	assert.Eventually(func() bool {
		err = tc.k8sClient.Get(ctx, eventhubNamespacedName, eventhubNamespaceInstance)
		return apierrors.IsNotFound(err)
	}, tc.timeout, tc.retry, "wait for eventHubnamespaceInstance to be gone from k8s")

}

func TestEventHubNamespaceControllerHappy(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()
	assert := assert.New(t)
	var err error

	var rgName string = tc.resourceGroupName
	var rgLocation string = tc.resourceGroupLocation
	eventhubNamespaceName := GenerateTestResourceNameWithRandom("ns-dev-eh", 10)

	// Create the Eventhub namespace object and expect the Reconcile to be created
	eventhubNamespaceInstance := &azurev1alpha1.EventhubNamespace{
		ObjectMeta: metav1.ObjectMeta{
			Name:      eventhubNamespaceName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.EventhubNamespaceSpec{
			Location:      rgLocation,
			ResourceGroup: rgName,
		},
	}

	EnsureInstance(ctx, t, tc, eventhubNamespaceInstance)

	eventhubNamespacedName := types.NamespacedName{Name: eventhubNamespaceName, Namespace: "default"}

	err = tc.k8sClient.Delete(ctx, eventhubNamespaceInstance)
	assert.Equal(nil, err, "delete eventhubns in k8s")

	assert.Eventually(func() bool {
		err = tc.k8sClient.Get(ctx, eventhubNamespacedName, eventhubNamespaceInstance)
		return apierrors.IsNotFound(err)
	}, tc.timeout, tc.retry, "wait for eventHubnamespaceInstance to be gone from k8s")
}
