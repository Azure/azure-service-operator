package storage

import (
	"context"
	"encoding/json"

	uuid "github.com/satori/go.uuid"

	azureV1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/client/deployment"
	"github.com/Azure/azure-service-operator/pkg/template"
)

// New generates a new object
func New(storage *azureV1alpha1.Storage) *Template {
	return &Template{
		Storage: storage,
	}
}

// Template defines the dynamodb cfts
type Template struct {
	Storage *azureV1alpha1.Storage
}

func (t *Template) CreateDeployment(ctx context.Context, resourceGroupName string) (string, error) {
	deploymentName := uuid.NewV4().String()
	asset, err := template.Asset("storage.json")
	templateContents := make(map[string]interface{})
	json.Unmarshal(asset, &templateContents)
	params := map[string]interface{}{
		"location": map[string]interface{}{
			"value": t.Storage.Spec.Location,
		},
		"accountType": map[string]interface{}{
			"value": t.Storage.Spec.Sku.Name,
		},
		"kind": map[string]interface{}{
			"value": t.Storage.Spec.Kind,
		},
		"accessTier": map[string]interface{}{
			"value": t.Storage.Spec.AccessTier,
		},
		"supportsHttpsTrafficOnly": map[string]interface{}{
			"value": *t.Storage.Spec.EnableHTTPSTrafficOnly,
		},
	}

	err = deployment.CreateDeployment(ctx, resourceGroupName, deploymentName, &templateContents, &params)
	return deploymentName, err
}
