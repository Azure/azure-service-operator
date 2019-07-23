package cosmosdb

import (
	"context"

	uuid "github.com/satori/go.uuid"

	azureV1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/client/deployment"
)

// New generates a new object
func New(cosmosdb *azureV1alpha1.CosmosDB) *Template {
	return &Template{
		CosmosDB: cosmosdb,
	}
}

// Template defines the dynamodb cfts
type Template struct {
	CosmosDB *azureV1alpha1.CosmosDB
}

func (t *Template) CreateDeployment(ctx context.Context, resourceGroupName string) (string, error) {
	deploymentName := uuid.NewV4().String()
	templateURI := "https://azureserviceoperator.blob.core.windows.net/templates/cosmosdb.json"
	params := map[string]interface{}{
		"location": map[string]interface{}{
			"value": t.CosmosDB.Spec.Location,
		},
		"kind": map[string]interface{}{
			"value": t.CosmosDB.Spec.Kind,
		},
		"properties": map[string]interface{}{
			"value": t.CosmosDB.Spec.Properties,
		},
	}

	err := deployment.CreateDeployment(ctx, resourceGroupName, deploymentName, templateURI, &params)
	return deploymentName, err
}
