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
	"github.com/Azure/go-autorest/autorest"
	"github.com/go-logr/logr"

	"k8s.io/apimachinery/pkg/runtime"
)

// NewAppInsightsManager creates a new AppInsightsManager
func NewAppInsightsManager(log logr.Logger, scheme *runtime.Scheme) AppInsightsManager {
	return AppInsightsManager{
		Log:    log,
		Scheme: scheme,
	}
}

// ApplicationInsightsManager manages Azure Application Insights service components
type ApplicationInsightsManager interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, location string) (result insights.ApplicationInsightsComponent, err error)
	Delete(ctx context.Context, resourceGroupName string, resourceName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, resourceName string) (result insights.ApplicationInsightsComponent, err error)
}
