// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

// +build rediscache

package controllers

import (
	"context"
	"testing"
	"time"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/stretchr/testify/assert"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const longRunningTimeout = 25 * time.Minute

func TestRedisCacheControllerHappyPath(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()
	assert := assert.New(t)

	var rgLocation string
	var rgName string
	var redisCacheName string
	var err error

	rgName = tc.resourceGroup
	rgLocation = tc.resourceGroupLocation
	redisCacheName = GenerateTestResourceNameWithRandom("rediscache", 10)

	// Create the RedisCache object and expect the Reconcile to be created
	redisCacheInstance := &azurev1alpha1.RedisCache{
		ObjectMeta: metav1.ObjectMeta{
			Name:      redisCacheName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.RedisCacheSpec{
			Location:      rgLocation,
			ResourceGroup: rgName,
			Properties: azurev1alpha1.RedisCacheProperties{
				Sku: azurev1alpha1.RedisCacheSku{
					Name:     "Basic",
					Family:   "C",
					Capacity: 0,
				},
				EnableNonSslPort: true,
			},
		},
	}

	// create rc
	EnsureInstance(ctx, t, tc, redisCacheInstance)

	// verify secret exists in secretclient
	EnsureSecrets(ctx, t, tc, redisCacheInstance, tc.SecretClient, redisCacheInstance.Name, redisCacheInstance.Namespace)

	// delete rc
	EnsureDelete(ctx, t, tc, redisCacheInstance)
}
