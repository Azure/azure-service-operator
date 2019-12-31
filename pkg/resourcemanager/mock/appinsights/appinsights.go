/*
Copyright 2019 Microsoft.

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
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/go-logr/logr"

	"k8s.io/apimachinery/pkg/runtime"
)

type appInsightsResource struct {
	resourceGroupName string
	location          string
	appInsightsName   string
}

type mockAppInsightsManager struct {
	appInsightsResource *appInsightsResource
	Log                 logr.Logger
	Scheme              *runtime.Scheme
}

// NewMockAppInsightsManager creates a new mock of an ApplicationInsightsManager
func NewMockAppInsightsManager(log logr.Logger, scheme *runtime.Scheme) *mockAppInsightsManager {
	return &mockAppInsightsManager{}
}

func (m *mockAppInsightsManager) CreateAppInsights(
	ctx context.Context,
	kind string,
	resourceGroupName string,
	location string,
	resourceName string,
	tags map[string]*string,
	properties insights.ApplicationInsightsComponentProperties) (insights.ApplicationInsightsComponent, error) {

	i := insights.ApplicationInsightsComponent{
		ApplicationInsightsComponentProperties: &properties,
		Kind:                                   to.StringPtr(kind),
		Location:                               to.StringPtr(location),
		Tags:                                   tags,
	}
	return i, nil
}
