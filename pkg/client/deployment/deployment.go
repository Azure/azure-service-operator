package deployment

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2019-03-01/resources"

	"github.com/Azure/azure-service-operator/pkg/config"
	"github.com/Azure/azure-service-operator/pkg/iam"
)

func getDeploymentsClient() resources.DeploymentsClient {
	deployClient := resources.NewDeploymentsClient(config.Instance.SubscriptionID)
	a, _ := iam.GetResourceManagementAuthorizer()
	deployClient.Authorizer = a
	return deployClient
}

// CreateDeployment creates a template deployment using the
// referenced JSON files for the template and its parameters
func CreateDeployment(ctx context.Context, resourceGroupName, deploymentName, templateUri string, params *map[string]interface{}) (de resources.DeploymentExtended, err error) {
	deployClient := getDeploymentsClient()
	templateLink := resources.TemplateLink{
		URI: &templateUri,
	}
	future, err := deployClient.CreateOrUpdate(
		ctx,
		resourceGroupName,
		deploymentName,
		resources.Deployment{
			Properties: &resources.DeploymentProperties{
				TemplateLink: &templateLink,
				Parameters:   params,
				Mode:         resources.Incremental,
			},
		},
	)
	if err != nil {
		return de, fmt.Errorf("cannot create deployment: %v", err)
	}

	err = future.WaitForCompletionRef(ctx, deployClient.Client)
	if err != nil {
		return de, fmt.Errorf("cannot get the create deployment future respone: %v", err)
	}

	return future.Result(deployClient)
}
