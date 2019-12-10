// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package rediscaches

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/redis/mgmt/2018-03-01/redis"
	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
)

func NewAzureRedisCacheManager() *AzureRedisCacheManager {
	return &AzureRedisCacheManager{}
}

type RedisCacheManager interface {
	// CreateRedisCache creates a new RedisCache
	CreateRedisCache(ctx context.Context,
		groupName string,
		redisCacheName string,
		location string,
		sku azurev1alpha1.RedisCacheSku,
		enableNonSSLPort bool,
		tags map[string]*string) (*redis.ResourceType, error)

	// DeleteRedisCache removes the resource group named by env var
	DeleteRedisCache(ctx context.Context, groupName string, redisCacheName string) (result redis.DeleteFuture, err error)
}
