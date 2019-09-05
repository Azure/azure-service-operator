package rediscaches

import (
	"context"
	"errors"
	"log"

	"github.com/Azure/azure-sdk-for-go/services/redis/mgmt/2018-03-01/redis"
	azurev1 "github.com/Azure/azure-service-operator/api/v1"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/iam"
	"github.com/Azure/go-autorest/autorest/to"
)

func getRedisCacheClient() redis.Client {
	redisClient := redis.NewClient(config.SubscriptionID())
	a, err := iam.GetResourceManagementAuthorizer()
	if err != nil {
		log.Fatalf("failed to initialize authorizer: %v\n", err)
	}
	redisClient.Authorizer = a
	redisClient.AddToUserAgent(config.UserAgent())
	return redisClient
}

// CreateRedisCache creates a new RedisCache
func CreateRedisCache(ctx context.Context,
	groupName string,
	redisCacheName string,
	location string,
	sku azurev1.RedisCacheSku,
	enableNonSSLPort bool,
	tags map[string]*string) (*redis.ResourceType, error) {
	redisClient := getRedisCacheClient()

	//log.Println("RedisCache:CacheName" + redisCacheName)

	//Check if name is available
	redisType := "Microsoft.Cache/redis"
	checkNameParams := redis.CheckNameAvailabilityParameters{
		Name: &redisCacheName,
		Type: &redisType,
	}
	checkNameResult, err := redisClient.CheckNameAvailability(ctx, checkNameParams)
	if err != nil {
		return nil, err
	}

	if checkNameResult.StatusCode != 200 {
		log.Fatalf("redis cache name (%s) not available: %v\n", redisCacheName, checkNameResult.Status)
		return nil, errors.New("redis cache name not available")
	}

	redisSku := redis.Sku{
		Name:     redis.SkuName(sku.Name),
		Family:   redis.SkuFamily(sku.Family),
		Capacity: to.Int32Ptr(sku.Capacity),
	}

	createParams := redis.CreateParameters{
		Location: to.StringPtr(location),
		Tags:     tags,
		CreateProperties: &redis.CreateProperties{
			EnableNonSslPort: &enableNonSSLPort,
			Sku:              &redisSku,
		},
	}

	future, err := redisClient.Create(
		ctx, groupName, redisCacheName, createParams)
	if err != nil {
		return nil, err
	}

	err = future.WaitForCompletionRef(ctx, redisClient.Client)
	if err != nil {
		return nil, err
	}
	result, err := future.Result(redisClient)
	return &result, err
}

// DeleteRedisCache removes the resource group named by env var
func DeleteRedisCache(ctx context.Context, groupName string, redisCacheName string) (result redis.DeleteFuture, err error) {
	redisClient := getRedisCacheClient()
	return redisClient.Delete(ctx, groupName, redisCacheName)
}
