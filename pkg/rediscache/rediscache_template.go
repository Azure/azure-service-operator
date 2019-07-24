package rediscache

import (
	"context"
	"encoding/json"

	uuid "github.com/satori/go.uuid"

	azureV1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/client/deployment"
	"github.com/Azure/azure-service-operator/pkg/template"
)

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
