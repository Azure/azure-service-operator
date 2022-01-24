// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

//go:build all || rediscache
// +build all rediscache

package controllers

import (
	"context"
	"testing"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestRedisCacheActionControllerNoResourceGroup(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()

	var rgName string
	var redisCache string
	var redisCacheAction string

	rgName = GenerateTestResourceNameWithRandom("rcfwr-rg", 10)
	redisCache = GenerateTestResourceNameWithRandom("rediscache", 10)
	redisCacheAction = GenerateTestResourceNameWithRandom("rediscacheaction", 10)

	redisCacheActionInstance := &azurev1alpha1.RedisCacheAction{
		ObjectMeta: metav1.ObjectMeta{
			Name:      redisCacheAction,
			Namespace: "default",
		},
		Spec: azurev1alpha1.RedisCacheActionSpec{
			ResourceGroup: rgName,
			CacheName:     redisCache,
			ActionName:    azurev1alpha1.RedisCacheActionNameRollAllKeys,
		},
	}

	EnsureInstanceWithResult(ctx, t, tc, redisCacheActionInstance, errhelp.ResourceGroupNotFoundErrorCode, false)
	EnsureDelete(ctx, t, tc, redisCacheActionInstance)
}

func TestRedisCacheActionNoRedisCache(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()

	var rgName string
	var redisCache string
	var redisCacheAction string

	rgName = tc.resourceGroupName
	redisCache = GenerateTestResourceNameWithRandom("rediscache", 10)
	redisCacheAction = GenerateTestResourceNameWithRandom("rediscacheaction", 10)

	redisCacheActionInstance := &azurev1alpha1.RedisCacheAction{
		ObjectMeta: metav1.ObjectMeta{
			Name:      redisCacheAction,
			Namespace: "default",
		},
		Spec: azurev1alpha1.RedisCacheActionSpec{
			ResourceGroup: rgName,
			CacheName:     redisCache,
			ActionName:    azurev1alpha1.RedisCacheActionNameRollAllKeys,
		},
	}
	EnsureInstanceWithResult(ctx, t, tc, redisCacheActionInstance, errhelp.ResourceNotFound, false)
	EnsureDelete(ctx, t, tc, redisCacheActionInstance)
}
