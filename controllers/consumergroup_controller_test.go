// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

// +build all consumergroup

package controllers

import (
	"context"
	"net/http"
	"testing"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/stretchr/testify/assert"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
)

func TestConsumerGroup(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	var rgName string = tc.resourceGroupName
	var ehnName string = tc.eventhubNamespaceName
	var ehName string = tc.eventhubName
	var ctx = context.Background()
	defer PanicRecover(t)

	consumerGroupName := GenerateTestResourceNameWithRandom("cg", 10)
	azureConsumerGroupName := consumerGroupName + "-azure"

	var err error

	// Create the consumer group object and expect the Reconcile to be created
	consumerGroupInstance := &azurev1alpha1.ConsumerGroup{
		ObjectMeta: metav1.ObjectMeta{
			Name:      consumerGroupName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.ConsumerGroupSpec{
			Namespace:         ehnName,
			ResourceGroup:     rgName,
			Eventhub:          ehName,
			ConsumerGroupName: azureConsumerGroupName,
		},
	}

	err = tc.k8sClient.Create(ctx, consumerGroupInstance)
	assert.Equal(nil, err, "create consumergroup in k8s")

	consumerGroupNamespacedName := types.NamespacedName{Name: consumerGroupName, Namespace: "default"}

	assert.Eventually(func() bool {
		_ = tc.k8sClient.Get(ctx, consumerGroupNamespacedName, consumerGroupInstance)
		return HasFinalizer(consumerGroupInstance, finalizerName)
	}, tc.timeout, tc.retry, "wait for finalizer")

	assert.Eventually(func() bool {
		_ = tc.k8sClient.Get(ctx, consumerGroupNamespacedName, consumerGroupInstance)
		return consumerGroupInstance.Status.Provisioned
	}, tc.timeout, tc.retry, "wait for provision")

	assert.Eventually(func() bool {
		cg, _ := tc.consumerGroupClient.GetConsumerGroup(ctx, rgName, ehnName, ehName, azureConsumerGroupName)
		return cg.Name != nil && *cg.Name == azureConsumerGroupName && cg.Response.StatusCode == http.StatusOK
	}, tc.timeout, tc.retry, "wait for consumergroup to exist in Azure")

	err = tc.k8sClient.Delete(ctx, consumerGroupInstance)
	assert.Equal(nil, err, "delete consumergroup in k8s")

	assert.Eventually(func() bool {
		err = tc.k8sClient.Get(ctx, consumerGroupNamespacedName, consumerGroupInstance)
		return apierrors.IsNotFound(err)
	}, tc.timeout, tc.retry, "wait for consumergroup to be gone from k8s")

	assert.Eventually(func() bool {
		cg, _ := tc.consumerGroupClient.GetConsumerGroup(ctx, rgName, ehnName, ehName, azureConsumerGroupName)
		return cg.Response.StatusCode != http.StatusOK
	}, tc.timeout, tc.retry, "wait for consumergroup to be gone from azure")

}
