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
	"github.com/go-logr/logr"
	"net/http"

	"github.com/Azure/azure-service-operator/pkg/resourcemanager/eventhubs"

	"k8s.io/apimachinery/pkg/runtime"
)

type ResourceManagerClient struct {
	logger                   logr.Logger
	eventHubNamespaceManager eventhubs.EventHubNamespaceManager
}

func (client *ResourceManagerClient) Create(ctx context.Context, r runtime.Object) (controller_refactor.EnsureResult, error) {
	ehnDef, err := convertInstance(r)
	if err != nil {
		return controller_refactor.EnsureFailed, err
	}
	client.logger.Info("EventhubNamespace " + ehnDef.Name + " creating on Azure. Please be patient.")
	_, err = client.eventHubNamespaceManager.CreateNamespaceAndWait(ctx, ehnDef.Spec.ResourceGroup, ehnDef.Name, ehnDef.Spec.Location)
	client.logger.Info("EventhubNamespace " + ehnDef.Name + " finished creating on Azure.")
	if err != nil {
		return controller_refactor.EnsureFailed, err
	}
	return controller_refactor.EnsureAwaitingVerification, nil
}

func (client *ResourceManagerClient) Update(ctx context.Context, r runtime.Object) (controller_refactor.EnsureResult, error) {
	return controller_refactor.EnsureFailed, fmt.Errorf("eventhubnamespace cannot be updated")
}

func (client *ResourceManagerClient) Verify(ctx context.Context, r runtime.Object) (controller_refactor.VerifyResult, error) {
	ehnDef, err := convertInstance(r)
	if err != nil {
		return controller_refactor.VerifyError, err
	}

	client.logger.Info("Fetching EventhubNamespace " + ehnDef.Name + " from Azure.")
	ehn, err := client.eventHubNamespaceManager.GetNamespace(ctx, ehnDef.Spec.ResourceGroup, ehnDef.Name)
	if ehn == nil || ehn.Response.Response == nil {
		return controller_refactor.VerifyError, fmt.Errorf("eventhubnamespace verify was nil for %s", ehnDef.Name)
	} else if ehn.Response.StatusCode == http.StatusNotFound {
		return controller_refactor.VerifyMissing, nil
	} else if err != nil {
		return controller_refactor.VerifyError, err
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
	ehnDef, err := convertInstance(r)
	if err != nil {
		return controller_refactor.DeleteError, err
	}

	client.logger.Info("EventhubNamespace " + ehnDef.Name + " deleting on Azure. Please be patient.")
	resp, err := client.eventHubNamespaceManager.DeleteNamespace(ctx, ehnDef.Spec.ResourceGroup, ehnDef.Name)
	if resp.Response == nil {
		return controller_refactor.DeleteError, err
	}
	if resp.StatusCode == http.StatusNotFound {
		return controller_refactor.DeleteAlreadyDeleted, nil
	} else if resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusAccepted {
		return controller_refactor.DeleteSucceed, nil
	}

	// TODO: handle all other cases
	return controller_refactor.DeleteSucceed, nil
}
