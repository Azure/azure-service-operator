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

package eventhubnamespace

// TODO

import (
	"context"
	"fmt"
	"github.com/Azure/azure-service-operator/controller_refactor"
	"net/http"

	"github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/eventhubs"

	"k8s.io/apimachinery/pkg/runtime"
)

type ResourceManagerClient struct {
	EventHubNamespaceManager eventhubs.EventHubNamespaceManager
}

func (client *ResourceManagerClient) Create(ctx context.Context, r runtime.Object) (controller_refactor.EnsureResult, error) {
	ehnDef, err := client.convert(r)
	if err != nil {
		return controller_refactor.EnsureFailed, err
	}
	_, err = client.EventHubNamespaceManager.CreateNamespaceAndWait(ctx, ehnDef.Spec.ResourceGroup, ehnDef.Name, ehnDef.Spec.Location)
	if err != nil {
		return controller_refactor.EnsureFailed, err
	}
	return controller_refactor.EnsureAwaitingVerification, nil
}

func (client *ResourceManagerClient) Update(ctx context.Context, r runtime.Object) (controller_refactor.EnsureResult, error) {
	return controller_refactor.EnsureFailed, fmt.Errorf("eventhubnamespace cannot be updated")
}

func (client *ResourceManagerClient) Verify(ctx context.Context, r runtime.Object) (controller_refactor.VerifyResult, error) {
	ehnDef, err := client.convert(r)
	if err != nil {
		return controller_refactor.VerifyError, err
	}

	ehn, err := client.EventHubNamespaceManager.GetNamespace(ctx, ehnDef.Spec.ResourceGroup, ehnDef.Name)
	if err == nil {
		return controller_refactor.VerifyError, err
	} else if ehn == nil {
		return controller_refactor.VerifyError, fmt.Errorf("eventhubnamespace verify was nil for %s", ehnDef.Name)
	} else if ehn.Response.StatusCode == http.StatusNotFound {
		return controller_refactor.VerifyMissing, nil
	} else if ehn.Response.StatusCode == http.StatusOK {
		if ehn.ProvisioningState != nil && *ehn.ProvisioningState == "Succeeded" {
			// TODO: handle cases that lead to VerifyUpdateRequired and VerifyRecreateRequired
			return controller_refactor.VerifyReady, nil
		} else {
			// TODO: handle cases that lead to VerifyDeleting (what are the undocumented values of *ehn.ProvisioningState?)
			return controller_refactor.VerifyProvisioning, nil
		}
	}

	// we ideally shouldn't get to this point - all cases should be handled explicitly
	return controller_refactor.VerifyMissing, nil
}

func (client *ResourceManagerClient) Delete(ctx context.Context, r runtime.Object) (controller_refactor.DeleteResult, error) {
	ehnDef, err := client.convert(r)
	if err != nil {
		return controller_refactor.DeleteError, err
	}

	resp, err := client.EventHubNamespaceManager.DeleteNamespace(ctx, ehnDef.Spec.ResourceGroup, ehnDef.Name)
	if resp.Response != nil {
		return controller_refactor.DeleteError, err
	}
	if resp.StatusCode == http.StatusNotFound {
		return controller_refactor.DeleteAlreadyDeleted, nil
	} else if resp.StatusCode == http.StatusOK {
		return controller_refactor.DeleteSucceed, nil
	}

	// TODO: handle all other cases
	return controller_refactor.DeleteSucceed, nil
}

func (_ *ResourceManagerClient) convert(obj runtime.Object) (*v1alpha1.EventhubNamespace, error) {
	local, ok := obj.(*v1alpha1.EventhubNamespace)
	if !ok {
		return nil, fmt.Errorf("failed type assertion on kind: %s", obj.GetObjectKind().GroupVersionKind().String())
	}
	return local, nil
}
