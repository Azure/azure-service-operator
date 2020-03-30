// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package rediscaches

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/redis/mgmt/2018-03-01/redis"
	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
)

// RedisCacheManager for RedisCache
type RedisCacheManager interface {
	// CreateRedisCache creates a new RedisCache
	CreateRedisCache(ctx context.Context,
		instance azurev1alpha1.RedisCache) (*redis.ResourceType, error)

	// DeleteRedisCache removes the resource group named by env var
	DeleteRedisCache(ctx context.Context, groupName string, redisCacheName string) (result redis.DeleteFuture, err error)
	// also embed async client methods
	resourcemanager.ARMClient
}
