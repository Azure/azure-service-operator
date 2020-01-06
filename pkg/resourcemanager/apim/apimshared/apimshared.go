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

package apimshared

import (
	"context"
	"strings"

	apim "github.com/Azure/azure-sdk-for-go/services/apimanagement/mgmt/2019-01-01/apimanagement"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/iam"
)

// GetAPIMgmtSvcClient returns a new instance of an API Svc client
func GetAPIMgmtSvcClient() (apim.ServiceClient, error) {
	client := apim.NewServiceClient(config.SubscriptionID())
	a, err := iam.GetResourceManagementAuthorizer()
	if err != nil {
		client = apim.ServiceClient{}
	} else {
		client.Authorizer = a
		client.AddToUserAgent(config.UserAgent())
	}
	return client, err
}

// GetAPIMgmtSvc returns an instance of an APIM service
func GetAPIMgmtSvc(ctx context.Context, resourceGroupName string, resourceName string) (apim.ServiceResource, error) {
	client, err := GetAPIMgmtSvcClient()
	if err != nil {
		return apim.ServiceResource{}, err
	}

	return client.Get(
		ctx,
		resourceGroupName,
		resourceName,
	)
}

// IsAPIMgmtSvcActivated check to see if the API Mgmt Svc has been activated, returns "true" if it has been
// activated
func IsAPIMgmtSvcActivated(ctx context.Context, resourceGroupName string, resourceName string) (result bool, err error) {
	resource, err := GetAPIMgmtSvc(
		ctx,
		resourceGroupName,
		resourceName,
	)
	if err != nil {
		return false, err
	}

	result = false
	if strings.Compare(*resource.ServiceProperties.ProvisioningState, "Activating") != 0 {
		result = true
	}

	return result, err
}
