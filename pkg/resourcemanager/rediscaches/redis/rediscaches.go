// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package redis

import (
	"context"
	"errors"
	"log"

	"github.com/Azure/azure-sdk-for-go/services/redis/mgmt/2018-03-01/redis"
	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/rediscaches"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/vnet"
	"github.com/Azure/azure-service-operator/pkg/secrets"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/to"
	"k8s.io/apimachinery/pkg/runtime"
)

// AzureRedisCacheManager creates a new RedisCacheManager
type AzureRedisCacheManager struct {
	rediscaches.AzureRedisManager
}

// NewAzureRedisCacheManager creates a new RedisCacheManager
func NewAzureRedisCacheManager(creds config.Credentials, secretClient secrets.SecretClient, scheme *runtime.Scheme) *AzureRedisCacheManager {
	return &AzureRedisCacheManager{
		rediscaches.AzureRedisManager{
			Creds:        creds,
			SecretClient: secretClient,
			Scheme:       scheme,
		},
	}
}

// CreateRedisCache creates a new RedisCache
func (r *AzureRedisCacheManager) CreateRedisCache(
	ctx context.Context,
	instance azurev1alpha1.RedisCache) (*redis.ResourceType, error) {

	props := instance.Spec.Properties

	// convert kube labels to expected tag format
	tags := helpers.LabelsToTags(instance.GetLabels())

	redisClient, err := r.GetRedisCacheClient()
	if err != nil {
		return nil, err
	}

	//Check if name is available
	redisType := "Microsoft.Cache/redis"
	checkNameParams := redis.CheckNameAvailabilityParameters{
		Name: &instance.Name,
		Type: &redisType,
	}
	checkNameResult, err := redisClient.CheckNameAvailability(ctx, checkNameParams)
	if err != nil {
		return nil, err
	}

	if checkNameResult.StatusCode != 200 {
		log.Println("redis cache name (%s) not available: " + instance.Name + checkNameResult.Status)
		return nil, errors.New("redis cache name not available")
	}

	redisSku := &redis.Sku{
		Name:     redis.SkuName(props.Sku.Name),
		Family:   redis.SkuFamily(props.Sku.Family),
		Capacity: to.Int32Ptr(props.Sku.Capacity),
	}

	createParams := redis.CreateParameters{
		Location: to.StringPtr(instance.Spec.Location),
		Tags:     tags,
		CreateProperties: &redis.CreateProperties{
			EnableNonSslPort: &props.EnableNonSslPort,
			Sku:              redisSku,
		},
	}

	// handle vnet settings
	if len(props.SubnetID) > 0 {
		ip := props.StaticIP
		if len(props.StaticIP) == 0 {
			vnetManager := vnet.NewAzureVNetManager(r.Creds)
			sid := vnet.ParseSubnetID(props.SubnetID)

			ip, err = vnetManager.GetAvailableIP(ctx, sid.ResourceGroup, sid.VNet, sid.Subnet)
			if err != nil {
				return nil, err
			}
		}

		createParams.CreateProperties.SubnetID = &props.SubnetID
		createParams.CreateProperties.StaticIP = &ip
	}

	if redisSku.Name == redis.Premium && props.ShardCount != nil {
		createParams.CreateProperties.ShardCount = props.ShardCount
	}

	// set redis config if one was provided
	if len(props.Configuration) > 0 {
		config := map[string]*string{}
		for k, v := range props.Configuration {
			value := v
			config[k] = &value
		}
		createParams.CreateProperties.RedisConfiguration = config
	}

	future, err := redisClient.Create(
		ctx, instance.Spec.ResourceGroupName, instance.Name, createParams,
	)
	if err != nil {
		return nil, err
	}

	result, err := future.Result(redisClient)
	return &result, err
}

// GetRedisCache returns a redis cache object if it exists
func (r *AzureRedisCacheManager) GetRedisCache(ctx context.Context, groupName string, redisCacheName string) (result redis.ResourceType, err error) {
	redisClient, err := r.GetRedisCacheClient()
	if err != nil {
		return result, err
	}
	return redisClient.Get(ctx, groupName, redisCacheName)
}

// DeleteRedisCache removes the resource group named by env var
func (r *AzureRedisCacheManager) DeleteRedisCache(ctx context.Context, groupName string, redisCacheName string) (result autorest.Response, err error) {
	redisClient, err := r.GetRedisCacheClient()
	if err != nil {
		return result, err
	}
	future, err := redisClient.Delete(ctx, groupName, redisCacheName)
	if err != nil {
		return result, err
	}

	return future.Result(redisClient)

}
