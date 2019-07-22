package rediscache

import (
	"context"

	uuid "github.com/satori/go.uuid"

	azureV1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/client/deployment"
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
	templateURI := "https://azureserviceoperator.blob.core.windows.net/templates/rediscache.json"
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

	err := deployment.CreateDeployment(ctx, resourceGroupName, deploymentName, templateURI, &params)
	return deploymentName, err
}
