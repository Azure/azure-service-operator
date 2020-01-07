// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package rediscaches

import (
	"context"
	"errors"
	"net/http"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/mock/helpers"
	"github.com/Azure/go-autorest/autorest/to"

	"github.com/Azure/azure-sdk-for-go/services/redis/mgmt/2018-03-01/redis"
)

type MockRedisCacheManager struct {
	redisCaches []MockRedisCacheResource
}

type MockRedisCacheResource struct {
	resourceGroupName string
	redis             redis.ResourceType
}

func NewMockRedisCacheManager() *MockRedisCacheManager {
	return &MockRedisCacheManager{}
}

func findRedisCache(res []MockRedisCacheResource, predicate func(MockRedisCacheResource) bool) (int, MockRedisCacheResource) {
	for index, r := range res {
		if predicate(r) {
			return index, r
		}
	}
	return -1, MockRedisCacheResource{}
}

func (manager *MockRedisCacheManager) CreateRedisCache(ctx context.Context,
	groupName string,
	redisCacheName string,
	location string,
	sku azurev1alpha1.RedisCacheSku,
	enableNonSSLPort bool,
	tags map[string]*string) (*redis.ResourceType, error) {
	index, _ := findRedisCache(manager.redisCaches, func(s MockRedisCacheResource) bool {
		return s.resourceGroupName == groupName && *s.redis.Name == redisCacheName
	})

	rc := redis.ResourceType{
		Response: helpers.GetRestResponse(http.StatusCreated),
		Location: to.StringPtr(location),
		Name:     to.StringPtr(redisCacheName),
	}

	r := MockRedisCacheResource{
		resourceGroupName: groupName,
		redis:             rc,
	}

	if index == -1 {
		manager.redisCaches = append(manager.redisCaches, r)
	}

	return &r.redis, nil
}

func (manager *MockRedisCacheManager) DeleteRedisCache(ctx context.Context, groupName string, redisCacheName string) (result redis.DeleteFuture, err error) {
	redisCaches := manager.redisCaches

	index, _ := findRedisCache(redisCaches, func(s MockRedisCacheResource) bool {
		return s.resourceGroupName == groupName &&
			*s.redis.Name == redisCacheName
	})

	if index == -1 {
		return redis.DeleteFuture{}, errors.New("Redis Cache Not Found")
	}

	manager.redisCaches = append(redisCaches[:index], redisCaches[index+1:]...)

	return redis.DeleteFuture{}, nil
}
