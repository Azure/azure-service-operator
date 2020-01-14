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

package appinsights

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/appinsights/mgmt/2015-05-01/insights"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	"github.com/Azure/go-autorest/autorest"
)

// ApplicationInsightsManager manages Azure Application Insights service components
type ApplicationInsightsManager interface {
	CreateAppInsights(
		ctx context.Context,
		resourceGroupName string,
		kind string,
		applicationType string,
		location string,
		resourceName string) (insights.ApplicationInsightsComponent, error)
	DeleteAppInsights(ctx context.Context, resourceGroupName string, resourceName string) (autorest.Response, error)
	GetAppInsights(ctx context.Context, resourceGroupName string, resourceName string) (insights.ApplicationInsightsComponent, error)

	// ARM Client
	resourcemanager.ARMClient
}
