// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package actions

import (
	"context"
	"fmt"

	"github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/rediscaches"
	"github.com/Azure/azure-service-operator/pkg/secrets"

	model "github.com/Azure/azure-sdk-for-go/services/redis/mgmt/2018-03-01/redis"

	"k8s.io/apimachinery/pkg/runtime"
)

// AzureRedisCacheActionManager creates a new RedisCacheManager
type AzureRedisCacheActionManager struct {
	rediscaches.AzureRedisManager
}

// NewAzureRedisCacheActionManager creates a new RedisCacheManager
func NewAzureRedisCacheActionManager(creds config.Credentials, secretClient secrets.SecretClient, scheme *runtime.Scheme) *AzureRedisCacheActionManager {
	return &AzureRedisCacheActionManager{
		rediscaches.AzureRedisManager{
			Creds:        creds,
			SecretClient: secretClient,
			Scheme:       scheme,
		},
	}
}

// RegeneratePrimaryAccessKey regenerates either the primary or secondary access keys
func (r *AzureRedisCacheActionManager) RegeneratePrimaryAccessKey(ctx context.Context, resourceGroup string, cacheName string) error {
	client, err := r.GetRedisCacheClient()
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
	client, err := r.GetRedisCacheClient()
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

func (r *AzureRedisCacheActionManager) ForceReboot(ctx context.Context, resourceGroup string, cacheName string, actionName v1alpha1.RedisCacheActionName, shardID *int32) error {
	client, err := r.GetRedisCacheClient()
	if err != nil {
		return err
	}

	var rebootType model.RebootType
	switch actionName {
	case v1alpha1.RedisCacheActionNameRebootAllNodes:
		rebootType = model.AllNodes
	case v1alpha1.RedisCacheActionNameRebootPrimaryNode:
		rebootType = model.PrimaryNode
	case v1alpha1.RedisCacheActionNameRebootSecondaryNode:
		rebootType = model.SecondaryNode
	default:
		return fmt.Errorf("%s is not a valid reboot action", actionName)
	}

	_, err = client.ForceReboot(ctx, resourceGroup, cacheName, model.RebootParameters{
		RebootType: rebootType,
		ShardID:    shardID,
	})
	if err != nil {
		return err
	}
	return err
}
