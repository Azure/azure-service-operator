// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

//go:build all || consumergroup || eventhub || eventhubnamespace
// +build all consumergroup eventhub eventhubnamespace

package controllers

import (
	"context"
	"net/http"
	"testing"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/stretchr/testify/assert"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestConsumerGroupEventHubAndNamespaceControllerHappy(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()
	assert := assert.New(t)

	var rgName string = tc.resourceGroupName
	var rgLocation string = tc.resourceGroupLocation
	eventhubNamespaceName := GenerateTestResourceNameWithRandom("ns-dev-eh", 10)

	consumerGroupName := GenerateTestResourceNameWithRandom("cg", 10)
	azureConsumerGroupName := consumerGroupName + "-azure"

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

	eventhubName := GenerateTestResourceNameWithRandom("eh-cd", 10)

	// Create the EventHub object and expect the Reconcile to be created
	eventhubInstance := &azurev1alpha1.Eventhub{
		ObjectMeta: metav1.ObjectMeta{
			Name:      eventhubName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.EventhubSpec{
			Location:      rgLocation,
			Namespace:     eventhubNamespaceName,
			ResourceGroup: rgName,
			Properties: azurev1alpha1.EventhubProperties{
				MessageRetentionInDays: 7,
				PartitionCount:         2,
			},
			AuthorizationRule: azurev1alpha1.EventhubAuthorizationRule{
				Name:   "RootManageSharedAccessKey",
				Rights: []string{"Listen"},
			},
		},
	}

	// verify eventhub is created successfully
	EnsureInstance(ctx, t, tc, eventhubInstance)

	// Create a consumer group instance and verify
	consumerGroupInstance := &azurev1alpha1.ConsumerGroup{
		ObjectMeta: metav1.ObjectMeta{
			Name:      consumerGroupName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.ConsumerGroupSpec{
			Namespace:         eventhubNamespaceName,
			ResourceGroup:     rgName,
			Eventhub:          eventhubName,
			ConsumerGroupName: azureConsumerGroupName,
		},
	}
	EnsureInstance(ctx, t, tc, consumerGroupInstance)

	assert.Eventually(func() bool {
		cg, _ := tc.consumerGroupClient.GetConsumerGroup(ctx, rgName, eventhubNamespaceName, eventhubName, azureConsumerGroupName)
		return cg.Name != nil && *cg.Name == azureConsumerGroupName && cg.Response.StatusCode == http.StatusOK
	}, tc.timeout, tc.retry, "wait for consumergroup to exist in Azure")

	EnsureDelete(ctx, t, tc, consumerGroupInstance)

	// verify eventhub is deleted
	EnsureDelete(ctx, t, tc, eventhubInstance)

	// verify eventhubnamespace gets deleted
	EnsureDelete(ctx, t, tc, eventhubNamespaceInstance)
}
