// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package rediscaches

import (
	"context"
	"log"

	"github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/iam"
	"github.com/Azure/azure-service-operator/pkg/secrets"

	"github.com/Azure/azure-sdk-for-go/services/redis/mgmt/2018-03-01/redis"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
)

// AzureRedisManager
type AzureRedisManager struct {
	SecretClient secrets.SecretClient
	Scheme       *runtime.Scheme
}

func getRedisCacheClient() (redis.Client, error) {
	redisClient := redis.NewClientWithBaseURI(config.BaseURI(), config.SubscriptionID())
	a, err := iam.GetResourceManagementAuthorizer()
	if err != nil {
		log.Println("failed to initialize authorizer: " + err.Error())
		return redisClient, err
	}
	redisClient.Authorizer = a
	redisClient.AddToUserAgent(config.UserAgent())
	return redisClient, nil
}

//ListKeys lists the keys for redis cache
func (r *AzureRedisManager) ListKeys(ctx context.Context, resourceGroupName string, redisCacheName string) (result redis.AccessKeys, err error) {
	redisClient, err := getRedisCacheClient()
	if err != nil {
		return result, err
	}
	return redisClient.ListKeys(ctx, resourceGroupName, redisCacheName)
}

// CreateSecrets creates a secret for a redis cache
func (r *AzureRedisManager) CreateSecrets(ctx context.Context, instance *v1alpha1.RedisCache, data map[string][]byte) error {
	secretName := instance.Spec.SecretName
	if secretName == "" {
		secretName = instance.Name
	}

	key := types.NamespacedName{Name: secretName, Namespace: instance.Namespace}

	err := r.SecretClient.Upsert(
		ctx,
		key,
		data,
		secrets.WithOwner(instance),
		secrets.WithScheme(r.Scheme),
	)
	if err != nil {
		return err
	}

	return nil
}

// ListKeysAndCreateSecrets lists keys and creates secrets
func (r *AzureRedisManager) ListKeysAndCreateSecrets(ctx context.Context, instance *v1alpha1.RedisCache) error {
	var err error
	var result redis.AccessKeys

	result, err = r.ListKeys(ctx, instance.Spec.ResourceGroupName, instance.Name)
	if err != nil {
		return err
	}
	data := map[string][]byte{
		"primaryKey":   []byte(*result.PrimaryKey),
		"secondaryKey": []byte(*result.SecondaryKey),
	}

	err = r.CreateSecrets(
		ctx,
		instance,
		data,
	)
	if err != nil {
		return err
	}

	return nil
}
