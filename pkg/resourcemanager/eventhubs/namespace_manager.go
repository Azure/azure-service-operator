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

package eventhubs

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/eventhub/mgmt/2017-04-01/eventhub"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	"github.com/Azure/go-autorest/autorest"
)

type EventHubNamespaceManager interface {
	// DeleteNamespace deletes an existing namespace. This operation also removes all associated resources under the namespace.
	// Parameters:
	// resourceGroupName - name of the resource group within the azure subscription.
	// namespaceName - the Namespace name

	DeleteNamespace(ctx context.Context, resourceGroupName string, namespaceName string) (result autorest.Response, err error)

	// Get gets the description of the specified namespace.
	// Parameters:
	// resourceGroupName - name of the resource group within the azure subscription.
	// namespaceName - the Namespace name
	GetNamespace(ctx context.Context, resourceGroupName string, namespaceName string) (result *eventhub.EHNamespace, err error)

	// CreateNamespaceAndWait creates an Event Hubs namespace
	// Parameters:
	// resourceGroupName - name of the resource group within the azure subscription.
	// namespaceName - the Namespace name
	// location - azure region
	CreateNamespaceAndWait(ctx context.Context, resourceGroupName string, namespaceName string, location string) (*eventhub.EHNamespace, error)

	// also embed async client methods
	resourcemanager.AsyncClient
}
