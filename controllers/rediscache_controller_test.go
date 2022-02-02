// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

//go:build rediscache
// +build rediscache

package controllers

import (
	"context"
	"testing"
	"time"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const longRunningTimeout = 25 * time.Minute

func TestRedisCacheControllerHappyPath(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()

	var rgLocation string
	var rgName string
	var redisCacheName string

	rgName = tc.resourceGroupName
	rgLocation = tc.resourceGroupLocation
	redisCacheName = GenerateTestResourceNameWithRandom("rediscache", 10)

	// Create the RedisCache object and expect the Reconcile to be created
	redisCacheInstance := &azurev1alpha1.RedisCache{
		ObjectMeta: metav1.ObjectMeta{
			Name:      redisCacheName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.RedisCacheSpec{
			Location:          rgLocation,
			ResourceGroupName: rgName,
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
	EnsureSecrets(ctx, t, tc, redisCacheInstance, tc.secretClient, redisCacheInstance.Name, redisCacheInstance.Namespace)

	// delete rc
	EnsureDelete(ctx, t, tc, redisCacheInstance)
}
func TestRedisCacheControllerNoResourceGroup(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()

	var rgLocation string
	var rgName string
	var redisCacheName string

	rgName = GenerateTestResourceNameWithRandom("rcfwr-rg", 10)
	rgLocation = tc.resourceGroupLocation
	redisCacheName = GenerateTestResourceNameWithRandom("rediscache", 10)

	// Create the RedisCache object and expect the Reconcile to be created
	redisCacheInstance := &azurev1alpha1.RedisCache{
		ObjectMeta: metav1.ObjectMeta{
			Name:      redisCacheName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.RedisCacheSpec{
			Location:          rgLocation,
			ResourceGroupName: rgName,
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
	EnsureInstanceWithResult(ctx, t, tc, redisCacheInstance, errhelp.ResourceGroupNotFoundErrorCode, false)

	// delete rc
	EnsureDelete(ctx, t, tc, redisCacheInstance)
}
