// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package rediscaches

import (
	"context"

	"github.com/Azure/azure-service-operator/pkg/secrets"

	model "github.com/Azure/azure-sdk-for-go/services/redis/mgmt/2018-03-01/redis"

	"k8s.io/apimachinery/pkg/runtime"
)

// AzureRedisCacheActionManager creates a new RedisCacheManager
type AzureRedisCacheActionManager struct {
	AzureRedisManager
}

// NewAzureRedisCacheActionManager creates a new RedisCacheManager
func NewAzureRedisCacheActionManager(secretClient secrets.SecretClient, scheme *runtime.Scheme) *AzureRedisCacheActionManager {
	return &AzureRedisCacheActionManager{
		AzureRedisManager{
			SecretClient: secretClient,
			Scheme:       scheme,
		},
	}
}

// RegeneratePrimaryAccessKey regenerates either the primary or secondary access keys
func (r *AzureRedisCacheActionManager) RegeneratePrimaryAccessKey(ctx context.Context, resourceGroup string, cacheName string) error {
	client, err := getRedisCacheClient()
	if err != nil {
		return err
	}

	_, err = client.RegenerateKey(ctx, resourceGroup, cacheName, model.RegenerateKeyParameters{
		KeyType: model.Primary,
	})
	if err != nil {
		return err
	}

	return nil
}

// RegenerateSecondaryAccessKey regenerates either the primary or secondary access keys
func (r *AzureRedisCacheActionManager) RegenerateSecondaryAccessKey(ctx context.Context, resourceGroup string, cacheName string) error {
	client, err := getRedisCacheClient()
	if err != nil {
		return err
	}

	_, err = client.RegenerateKey(ctx, resourceGroup, cacheName, model.RegenerateKeyParameters{
		KeyType: model.Secondary,
	})
	if err != nil {
		return err
	}

	return nil
}
