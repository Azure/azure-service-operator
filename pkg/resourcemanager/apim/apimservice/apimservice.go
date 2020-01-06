/*

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package apimservice

import (
	"context"

	apim "github.com/Azure/azure-sdk-for-go/services/apimanagement/mgmt/2019-01-01/apimanagement"
	apimshared "github.com/Azure/azure-service-operator/pkg/resourcemanager/apim/apimshared"
	"github.com/Azure/go-autorest/autorest/to"
)

type AzureAPIMgmtServiceManager struct{}

// CreateAPIMgmtSvc creates a new API Mgmt Svc
func (_ *AzureAPIMgmtServiceManager) CreateAPIMgmtSvc(ctx context.Context, location string, resourceGroupName string, resourceName string) (apim.ServiceResource, error) {
	client, err := apimshared.GetAPIMgmtSvcClient()
	if err != nil {
		return apim.ServiceResource{}, err
	}

	svcProp := apim.ServiceProperties{}
	sku := apim.ServiceSkuProperties{
		Name: apim.SkuTypeBasic,
	}
	future, err := client.CreateOrUpdate(
		ctx,
		resourceGroupName,
		resourceName,
		apim.ServiceResource{
			Location:          to.StringPtr(location),
			ServiceProperties: &svcProp,
			Sku:               &sku,
		},
	)

	return future.Result(client)
}

// GetAPIMgmtSvc gets an instance of an API Mgmt Svc
func (_ *AzureAPIMgmtServiceManager) GetAPIMgmtSvc(ctx context.Context, resourceGroupName string, resourceName string) (apim.ServiceResource, error) {
	return apimshared.GetAPIMgmtSvc(ctx, resourceGroupName, resourceName)
}

// DeleteAPIMgmtSvc an instance of an API Mgmt Svc
func (_ *AzureAPIMgmtServiceManager) DeleteAPIMgmtSvc(ctx context.Context, resourceGroupName string, resourceName string) (apim.ServiceResource, error) {
	client, err := apimshared.GetAPIMgmtSvcClient()
	if err != nil {
		return apim.ServiceResource{}, err
	}

	result, err := client.Delete(ctx, resourceGroupName, resourceName)
	if err != nil {
		return apim.ServiceResource{}, err
	}

	return result.Result(client)
}

// IsAPIMgmtSvcActivated checks to see if the API Mgmt Svc has been activated
func (_ *AzureAPIMgmtServiceManager) IsAPIMgmtSvcActivated(ctx context.Context, resourceGroupName string, resourceName string) (result bool, err error) {
	return apimshared.IsAPIMgmtSvcActivated(ctx, resourceGroupName, resourceName)
}
