/*
Copyright 2019 microsoft.

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
	"errors"
	"fmt"
	"net/http"

	"github.com/Azure/azure-sdk-for-go/services/appinsights/mgmt/2015-05-01/insights"
	"github.com/Azure/azure-service-operator/api/v1alpha1"
	resourcemanager "github.com/Azure/azure-service-operator/pkg/resourcemanager"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/mock/helpers"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/to"
	"k8s.io/apimachinery/pkg/runtime"
)

// MockAppInsightsManager mocks the implementation of an AppInsights Manager type
type MockAppInsightsManager struct {
	Scheme            *runtime.Scheme
	resourceGroupName string
	appinsights       []insights.ApplicationInsightsComponent
}

// NewMockAppInsightsManager creates a new MockAppInsightsManager
func NewMockAppInsightsManager(scheme *runtime.Scheme) *MockAppInsightsManager {
	return &MockAppInsightsManager{
		Scheme:      scheme,
		appinsights: []insights.ApplicationInsightsComponent{},
	}
}

// GetParents fetches the ARM hierarchy of resources for this operator
func (m *MockAppInsightsManager) GetParents(obj runtime.Object) ([]resourcemanager.KubeParent, error) {
	return []resourcemanager.KubeParent{}, nil
}

// CreateAppInsights creates or updates a mock Application Insights service
func (m *MockAppInsightsManager) CreateAppInsights(
	ctx context.Context,
	resourceGroupName string,
	kind string,
	applicationType string,
	location string,
	resourceName string) (insights.ApplicationInsightsComponent, error) {

	index, _ := findAppInsight(m.appinsights, func(i insights.ApplicationInsightsComponent) bool {
		return *i.Name == resourceName
	})

	insights := insights.ApplicationInsightsComponent{
		Kind:     to.StringPtr(kind),
		Location: to.StringPtr(location),
		Name:     to.StringPtr(resourceName),
	}

	if index == -1 {
		m.appinsights = append(m.appinsights, insights)
	}

	return insights, nil
}

// Delete removes the operator from desired state
func (m *MockAppInsightsManager) Delete(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {

	instance, err := m.convert(obj)
	if err != nil {
		return false, err
	}

	_, err = m.DeleteAppInsights(ctx, instance.Spec.ResourceGroup, instance.ObjectMeta.Name)
	if err != nil {
		return false, err
	}

	return false, nil
}

// DeleteAppInsights removes an Application Insights service from a subscription
func (m *MockAppInsightsManager) DeleteAppInsights(ctx context.Context, resourceGroupName string, resourceName string) (autorest.Response, error) {
	appinsights := m.appinsights

	index, _ := findAppInsight(appinsights, func(i insights.ApplicationInsightsComponent) bool {
		return *i.Name == resourceName
	})

	if index == -1 {
		return helpers.GetRestResponse(http.StatusNotFound), errors.New("appinsights instance not found")
	}

	m.appinsights = append(appinsights[:index], appinsights[index+1:]...)

	return helpers.GetRestResponse(http.StatusOK), nil
}

// GetAppInsights fetches an Application Insights service reference
func (m *MockAppInsightsManager) GetAppInsights(ctx context.Context, resourceGroupName string, resourceName string) (insights.ApplicationInsightsComponent, error) {
	appinsights := m.appinsights

	index, appinsight := findAppInsight(appinsights, func(i insights.ApplicationInsightsComponent) bool {
		return *i.Name == resourceName
	})

	if index == -1 {
		return insights.ApplicationInsightsComponent{
			Response: helpers.GetRestResponse(http.StatusNotFound),
		}, errors.New("appinsights instance not found")
	}

	appinsight.Response = helpers.GetRestResponse(http.StatusOK)
	return appinsight, nil
}

// Ensure checks the desired state of the operator
func (m *MockAppInsightsManager) Ensure(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {

	i, err := m.convert(obj)
	if err != nil {
		return true, err
	}

	_, _ = m.CreateAppInsights(ctx, i.Spec.ResourceGroup, "web", "other", i.Spec.Location, i.Name)

	i.Status.Provisioned = true

	return true, nil
}

func findAppInsight(res []insights.ApplicationInsightsComponent, predicate func(insights.ApplicationInsightsComponent) bool) (int, insights.ApplicationInsightsComponent) {
	for index, r := range res {
		if predicate(r) {
			return index, r
		}
	}
	return -1, insights.ApplicationInsightsComponent{}
}

func (m *MockAppInsightsManager) convert(obj runtime.Object) (*v1alpha1.AppInsights, error) {
	i, ok := obj.(*v1alpha1.AppInsights)
	if !ok {
		return nil, fmt.Errorf("failed type assertion on kind: %s", obj.GetObjectKind().GroupVersionKind().String())
	}
	return i, nil
}
