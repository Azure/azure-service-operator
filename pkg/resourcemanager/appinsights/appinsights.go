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
	"fmt"
	"log"

	"github.com/Azure/azure-sdk-for-go/services/appinsights/mgmt/2015-05-01/insights"
	"github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/iam"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
)

// Manager manages Azure Application Insights services
type Manager struct {
	Log logr.Logger
}

// NewManager creates a new AppInsights Manager
func NewManager(log logr.Logger) *Manager {
	return &Manager{
		Log: log,
	}
}

func (m *Manager) convert(obj runtime.Object) (*v1alpha1.AppInsights, error) {
	i, ok := obj.(*v1alpha1.AppInsights)
	if !ok {
		return nil, fmt.Errorf("failed type assertion on kind: %s", obj.GetObjectKind().GroupVersionKind().String())
	}
	return i, nil
}

// GetParents fetches dependent ARM resources
func (m *Manager) GetParents(obj runtime.Object) ([]resourcemanager.KubeParent, error) {
	i, err := m.convert(obj)
	if err != nil {
		return nil, err
	}

	return []resourcemanager.KubeParent{
		{
			Key: types.NamespacedName{
				Namespace: i.Namespace,
				Name:      i.Spec.ResourceGroup,
			},
			Target: &v1alpha1.ResourceGroup{},
		},
		{
			Key: types.NamespacedName{
				Namespace: i.Namespace,
				Name:      i.Name,
			},
			Target: &v1alpha1.AppInsights{},
		},
	}, nil
}

// CreateAppInsights creates or updates an Application Insights service
func (m *Manager) CreateAppInsights(
	ctx context.Context,
	kind string,
	resourceGroupName string,
	location string,
	resourceName string) (insights.ApplicationInsightsComponent, error) {

	componentsClient := getComponentsClient()

	props := insights.ApplicationInsightsComponent{
		Kind:     to.StringPtr(kind),
		Location: to.StringPtr(location),
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

// Ensure checks the desired state of the operator
func (m *Manager) Ensure(ctx context.Context, obj runtime.Object) (bool, error) {
	instance, err := m.convert(obj)
	if err != nil {
		return false, err
	}

	// Set k8s status to provisioning at the beginning of this reconciliation
	instance.Status.Provisioning = true

	insights, err := m.GetAppInsights(ctx, instance.Spec.ResourceGroup, instance.Name)
	if err == nil {
		instance.Status.Provisioned = true
		instance.Status.Provisioning = false
		instance.Status.State = insights.Status
	}

	c, err := m.CreateAppInsights(ctx, instance.Spec.Kind, instance.Spec.ResourceGroup, instance.Spec.Location, instance.Name)
	if err != nil {
		instance.Status.Message = err.Error()
		instance.Status.Provisioning = false

		// errors we expect might happen that we are ok to wait for
		catch := []string{
			errhelp.ResourceGroupNotFoundErrorCode,
			errhelp.ParentNotFoundErrorCode,
			errhelp.NotFoundErrorCode,
			errhelp.AsyncOpIncompleteError,
		}

		azerr := errhelp.NewAzureErrorAzureError(err)

		if helpers.ContainsString(catch, azerr.Type) {
			// most of these errors mean the resource is actually not provisioning
			switch azerr.Type {
			case errhelp.AsyncOpIncompleteError:
				instance.Status.Provisioning = true
			}
			// reconciliation is not done but error is acceptable
			return false, nil
		}
		// reconciliation not done and we don't know what happened
		return false, err
	}

	instance.Status.State = autorest.String(c.ProvisioningState)

	return true, nil
}

// Delete removes an AppInsights resource
func (m *Manager) Delete(ctx context.Context, obj runtime.Object) (bool, error) {
	i, err := m.convert(obj)
	if err != nil {
		return false, err
	}

	response, err := m.DeleteAppInsights(ctx, i.Spec.ResourceGroup, i.Name)
	if err != nil {
		return false, err
	}
	i.Status.State = response.Status

	return true, err
}

// DeleteAppInsights removes an Application Insights service from a subscription
func (m *Manager) DeleteAppInsights(
	ctx context.Context,
	resourceGroupName string,
	resourceName string) (autorest.Response, error) {

	componentsClient := getComponentsClient()

	response, err := componentsClient.Get(ctx, resourceGroupName, resourceName)
	if err == nil {
		return componentsClient.Delete(ctx, resourceGroupName, resourceName)
	}
	return response.Response, nil
}

// GetAppInsights fetches an Application Insights service reference
func (m *Manager) GetAppInsights(
	ctx context.Context,
	resourceGroupName string,
	resourceName string) (insights.ApplicationInsightsComponent, error) {

	componentsClient := getComponentsClient()
	return componentsClient.Get(ctx, resourceGroupName, resourceName)
}

func getComponentsClient() insights.ComponentsClient {
	insightsClient := insights.NewComponentsClient(config.SubscriptionID())

	a, err := iam.GetResourceManagementAuthorizer()
	if err != nil {
		log.Fatalf("failed to initialize authorizer %v\n", err)
	}
	insightsClient.Authorizer = a
	insightsClient.AddToUserAgent(config.UserAgent())

	return insightsClient
}
