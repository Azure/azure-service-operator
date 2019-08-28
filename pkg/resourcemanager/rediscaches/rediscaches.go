package rediscaches

import (
	"context"
	//"encoding/json"
	"errors"
	"fmt"
	"log"

	//uuid "github.com/satori/go.uuid"

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
	// kind azurev1.CosmosDBKind,
	// dbType azurev1.CosmosDBDatabaseAccountOfferType,
	tags map[string]*string) (redis.ResourceType, error) {
	redisClient := getRedisCacheClient()

	log.Println("RedisCache:CacheName" + redisCacheName)

	//Check if name is available
	redisType := "Microsoft.Cache/redis"
	checkNameParams := redis.CheckNameAvailabilityParameters{
		Name: &redisCacheName,
		Type: &redisType,
	}
	result, err := redisClient.CheckNameAvailability(ctx, checkNameParams)
	if err != nil {
		return redis.ResourceType{}, err
	}

	if result.StatusCode != 200 {
		log.Fatalf("redis cache name (%s) not available: %v\n", redisCacheName, result.Status)
		return redis.ResourceType{}, errors.New("redis cache name not available")
	}

	log.Println(fmt.Sprintf("creating rediscache '%s' in resource group '%s' and location: %v", redisCacheName, groupName, location))

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
		log.Println(fmt.Sprintf("ERROR creating redisCache '%s' in resource group '%s' and location: %v", redisCacheName, groupName, location))
		log.Println(fmt.Printf("failed to initialize redis Cache: %v\n", err))
	}
	return future.Result(redisClient)
}

// DeleteRedisCache removes the resource group named by env var
func DeleteRedisCache(ctx context.Context, groupName string, redisCacheName string) (result redis.DeleteFuture, err error) {
	redisClient := getRedisCacheClient()
	return redisClient.Delete(ctx, groupName, redisCacheName)
}

/*  Before Refactor
// New generates a new object
func New(redisCache *azureV1alpha1.RedisCache) *Template {
	return &Template{
		RedisCache: redisCache,
	}
}

// Template defines the dynamodb cfts
type Template struct {
	RedisCache *azureV1alpha1.RedisCache
}

func (t *Template) CreateDeployment(ctx context.Context, resourceGroupName string) (string, error) {
	deploymentName := uuid.NewV4().String()
	asset, err := template.Asset("rediscache.json")
	templateContents := make(map[string]interface{})
	json.Unmarshal(asset, &templateContents)
	params := map[string]interface{}{
		"location": map[string]interface{}{
			"value": t.RedisCache.Spec.Location,
		},
		"properties.sku.name": map[string]interface{}{
			"value": t.RedisCache.Spec.Properties.Sku.Name,
		},
		"properties.sku.family": map[string]interface{}{
			"value": t.RedisCache.Spec.Properties.Sku.Family,
		},
		"properties.sku.capacity": map[string]interface{}{
			"value": t.RedisCache.Spec.Properties.Sku.Capacity,
		},
		"properties.enableNonSslPort": map[string]interface{}{
			"value": t.RedisCache.Spec.Properties.EnableNonSslPort,
		},
	}

	err = deployment.CreateDeployment(ctx, resourceGroupName, deploymentName, &templateContents, &params)
	return deploymentName, err
}

*/
