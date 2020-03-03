// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package rediscaches

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	resourcemanager "github.com/Azure/azure-service-operator/pkg/resourcemanager"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/mock/helpers"
	"github.com/Azure/go-autorest/autorest/to"

	"github.com/Azure/azure-sdk-for-go/services/redis/mgmt/2018-03-01/redis"
	"k8s.io/apimachinery/pkg/runtime"
)

type redisCacheResource struct {
	resourceGroupName string
	redis             redis.ResourceType
}

type mockRedisCacheManager struct {
	resourceGroupName string
	redisCaches       []redisCacheResource
}

func NewMockRedisCacheManager() *mockRedisCacheManager {
	return &mockRedisCacheManager{}
}

func findRedisCache(res []redisCacheResource, predicate func(redisCacheResource) bool) (int, redisCacheResource) {
	for index, r := range res {
		if predicate(r) {
			return index, r
		}
	}
	return -1, redisCacheResource{}
}

func (manager *mockRedisCacheManager) CreateRedisCache(ctx context.Context,
	groupName string,
	redisCacheName string,
	location string,
	sku azurev1alpha1.RedisCacheSku,
	enableNonSSLPort bool,
	tags map[string]*string) (*redis.ResourceType, error) {
	index, _ := findRedisCache(manager.redisCaches, func(s redisCacheResource) bool {
		return s.resourceGroupName == groupName && *s.redis.Name == redisCacheName
	})

	rc := redis.ResourceType{
		Response: helpers.GetRestResponse(http.StatusCreated),
		Location: to.StringPtr(location),
		Name:     to.StringPtr(redisCacheName),
	}

	r := redisCacheResource{
		resourceGroupName: groupName,
		redis:             rc,
	}

	if index == -1 {
		manager.redisCaches = append(manager.redisCaches, r)
	}

	return &r.redis, nil
}

func (manager *mockRedisCacheManager) DeleteRedisCache(ctx context.Context, groupName string, redisCacheName string) (result redis.DeleteFuture, err error) {
	redisCaches := manager.redisCaches

	index, _ := findRedisCache(redisCaches, func(s redisCacheResource) bool {
		return s.resourceGroupName == groupName &&
			*s.redis.Name == redisCacheName
	})

	if index == -1 {
		return redis.DeleteFuture{}, errors.New("Redis Cache Not Found")
	}

	manager.redisCaches = append(redisCaches[:index], redisCaches[index+1:]...)

	return redis.DeleteFuture{}, nil
}

func (manager *mockRedisCacheManager) convert(obj runtime.Object) (*azurev1alpha1.RedisCache, error) {
	local, ok := obj.(*azurev1alpha1.RedisCache)
	if !ok {
		return nil, fmt.Errorf("failed type assertion on kind: %s", obj.GetObjectKind().GroupVersionKind().String())
	}
	return local, nil
}

func (manager *mockRedisCacheManager) Ensure(ctx context.Context, obj runtime.Object) (bool, error) {
	instance, err := manager.convert(obj)
	if err != nil {
		return false, err
	}
	tags := map[string]*string{}
	_, _ = manager.CreateRedisCache(ctx, instance.Spec.ResourceGroupName, instance.Name, instance.Spec.Location, instance.Spec.Properties.Sku, instance.Spec.Properties.EnableNonSslPort, tags)

	instance.Status.Provisioned = true
	return true, nil

}

func (manager *mockRedisCacheManager) Delete(ctx context.Context, obj runtime.Object) (bool, error) {
	instance, err := manager.convert(obj)
	if err != nil {
		return false, err
	}

	_, _ = manager.DeleteRedisCache(ctx, instance.Spec.ResourceGroupName, instance.Name)

	return false, nil
}
func (manager *mockRedisCacheManager) GetParents(obj runtime.Object) ([]resourcemanager.KubeParent, error) {
	return []resourcemanager.KubeParent{}, nil
}
