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
)

// AzureRedisManager
type AzureRedisManager struct {
	Creds        config.Credentials
	SecretClient secrets.SecretClient
	Scheme       *runtime.Scheme
}

func (m *AzureRedisManager) GetRedisCacheClient() (redis.Client, error) {
	redisClient := redis.NewClientWithBaseURI(config.BaseURI(), m.Creds.SubscriptionID())
	a, err := iam.GetResourceManagementAuthorizer(m.Creds)
	if err != nil {
		log.Println("failed to initialize authorizer: " + err.Error())
		return redisClient, err
	}
	redisClient.Authorizer = a
	redisClient.AddToUserAgent(config.UserAgent())
	return redisClient, nil
}

// ListKeys lists the keys for redis cache
func (m *AzureRedisManager) ListKeys(ctx context.Context, resourceGroupName string, redisCacheName string) (result redis.AccessKeys, err error) {
	redisClient, err := m.GetRedisCacheClient()
	if err != nil {
		return result, err
	}
	return redisClient.ListKeys(ctx, resourceGroupName, redisCacheName)
}

// CreateSecrets creates a secret for a redis cache
func (m *AzureRedisManager) CreateSecrets(ctx context.Context, secretClient secrets.SecretClient, instance *v1alpha1.RedisCache, data map[string][]byte) error {
	secretName := instance.Spec.SecretName
	if secretName == "" {
		secretName = instance.Name
	}

	secretKey := secrets.SecretKey{Name: secretName, Namespace: instance.Namespace, Kind: instance.TypeMeta.Kind}
	err := secretClient.Upsert(
		ctx,
		secretKey,
		data,
		secrets.WithOwner(instance),
		secrets.WithScheme(m.Scheme),
	)
	if err != nil {
		return err
	}

	return nil
}

// ListKeysAndCreateSecrets lists keys and creates secrets
func (m *AzureRedisManager) ListKeysAndCreateSecrets(ctx context.Context, secretClient secrets.SecretClient, instance *v1alpha1.RedisCache) error {
	var err error
	var result redis.AccessKeys

	result, err = m.ListKeys(ctx, instance.Spec.ResourceGroupName, instance.Name)
	if err != nil {
		return err
	}
	data := map[string][]byte{
		"primaryKey":   []byte(*result.PrimaryKey),
		"secondaryKey": []byte(*result.SecondaryKey),
	}

	err = m.CreateSecrets(
		ctx,
		secretClient,
		instance,
		data,
	)
	if err != nil {
		return err
	}

	return nil
}
