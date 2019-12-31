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
	"log"

	"github.com/Azure/azure-sdk-for-go/services/appinsights/mgmt/2015-05-01/insights"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/iam"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/go-logr/logr"

	"k8s.io/apimachinery/pkg/runtime"
)

// Manager manages Azure Application Insights services
type AppInsightsManager struct {
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// CreateOrUpdate creates or updates an Application Insights service
func (m *AppInsightsManager) CreateOrUpdate(
	ctx context.Context,
	kind string,
	resourceGroupName string,
	location string,
	resourceName string,
	tags map[string]*string,
	properties insights.ApplicationInsightsComponentProperties) (insights.ApplicationInsightsComponent, error) {

	componentsClient := getComponentsClient()

	props := insights.ApplicationInsightsComponent{
		ApplicationInsightsComponentProperties: &properties,
		Kind:                                   to.StringPtr(kind),
		Location:                               to.StringPtr(location),
		Tags:                                   tags,
	}

	// submit the ARM request
	result, err := componentsClient.CreateOrUpdate(
		ctx,
		resourceGroupName,
		resourceName,
		props,
	)
	return result, err
}

// Delete removes an Application Insights service from a subscription
func (m *AppInsightsManager) Delete(
	ctx context.Context,
	resourceGroupName string,
	resourceName string) (autorest.Response, error) {

	componentsClient := getComponentsClient()
	return componentsClient.Delete(ctx, resourceGroupName, resourceName)
}

// Get fetches an Application Insights service reference
func (m *AppInsightsManager) Get(
	ctx context.Context,
	resourceGroupName string,
	resourceName string) (insights.ApplicationInsightsComponent, error) {

	componentsClient := getComponentsClient()
	return componentsClient.Get(ctx, resourceGroupName, resourceName)
}

func getComponentsClient() insights.ComponentsClient {
	componentsClient := insights.NewComponentsClient(config.SubscriptionID())

	a, err := iam.GetResourceManagementAuthorizer()
	if err != nil {
		log.Fatalf("failed to initialize authorizer %v\n", err)
	}
	componentsClient.Authorizer = a
	componentsClient.AddToUserAgent(config.UserAgent())

	return componentsClient
}
