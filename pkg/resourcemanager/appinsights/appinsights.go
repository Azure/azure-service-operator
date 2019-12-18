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
	"github.com/go-logr/logr"
)

// Manager manages Azure Application Insights services
type Manager struct {
	Log logr.Logger
}

// CreateOrUpdate creates or updates an Application Insights service
func (m *Manager) CreateOrUpdate(
	ctx context.Context,
	resourceGroupName string,
	location string,
	resourceName string,
	properties insights.ApplicationInsightsComponent) (result insights.ApplicationInsightsComponent, err error) {

	componentsClient := getComponentsClient()

	props := insights.ApplicationInsightsComponent{
		ApplicationInsightsComponentProperties: properties.ApplicationInsightsComponentProperties,
		Kind:                                   properties.Kind,
		Location:                               properties.Location,
		Type:                                   properties.Type,
		Tags:                                   properties.Tags,
	}

	// submit the ARM request
	component, err := componentsClient.CreateOrUpdate(
		ctx,
		resourceGroupName,
		resourceName,
		props,
	)

	return component, err
}

// Delete removes an Application Insights service from a subscription
func (m *Manager) Delete(
	ctx context.Context,
	resourceGroupName string,
	resourceName string) (result autorest.Response, err error) {

	componentsClient := getComponentsClient()
	return componentsClient.Delete(ctx, resourceGroupName, resourceName)
}

// Get fetches an Application Insights service reference
func (m *Manager) Get(
	ctx context.Context,
	resourceGroupName string,
	resourceName string) (result insights.ApplicationInsightsComponent, err error) {

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
