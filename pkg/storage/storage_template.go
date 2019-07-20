package storage

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2019-03-01/resources"
	uuid "github.com/satori/go.uuid"

	azureV1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/client/deployment"
	"github.com/Azure/azure-service-operator/pkg/helpers"
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

func (t *Template) CreateDeployment(ctx context.Context, resourceGroupName string) (resources.DeploymentExtended, error) {
	deploymentName := uuid.NewV4().String()
	templateURI := "https://azureserviceoperator.blob.core.windows.net/templates/storage.json"
	params := map[string]interface{}{
		"location": map[string]interface{}{
			"value": t.Storage.Spec.Location,
		},
		"storageAccountName": map[string]interface{}{
			"value": helpers.AzureResourceName(t.Storage.Kind),
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

	return deployment.CreateDeployment(ctx, resourceGroupName, deploymentName, templateURI, &params)
}
