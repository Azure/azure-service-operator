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

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/azure-service-operator/pkg/reconciler"
	"github.com/go-logr/logr"

	"github.com/Azure/azure-service-operator/pkg/resourcemanager/eventhubs"
)

type ResourceManager struct {
	Logger                   logr.Logger
	EventHubNamespaceManager eventhubs.EventHubNamespaceManager
}

func CreateResourceManagerClient(eventHubNamespaceManager eventhubs.EventHubNamespaceManager, logger logr.Logger) ResourceManager {
	return ResourceManager{
		Logger:                   logger,
		EventHubNamespaceManager: eventHubNamespaceManager,
	}
}

func (client *ResourceManager) Create(ctx context.Context, r reconciler.ResourceSpec) (reconciler.ApplyResponse, error) {
	ehnDef, err := convertInstance(r.Instance)
	if err != nil {
		return reconciler.ApplyError, err
	}
	client.Logger.Info("EventhubNamespace " + ehnDef.Name + " creating on Azure. Please be patient.")
	_, err = client.EventHubNamespaceManager.CreateNamespaceAndWait(ctx, ehnDef.Spec.ResourceGroup, ehnDef.Name, ehnDef.Spec.Location)
	client.Logger.Info("EventhubNamespace " + ehnDef.Name + " finished creating on Azure.")
	if err != nil {
		return reconciler.ApplyError, err
	}
	return reconciler.ApplyAwaitingVerification, nil
}

func (client *ResourceManager) Update(ctx context.Context, r reconciler.ResourceSpec) (reconciler.ApplyResponse, error) {
	return reconciler.ApplyError, fmt.Errorf("EventhubNamespace cannot be updated")
}

func (client *ResourceManager) Verify(ctx context.Context, r reconciler.ResourceSpec) (reconciler.VerifyResponse, error) {
	ehnDef, err := convertInstance(r.Instance)
	if err != nil {
		return reconciler.VerifyError, err
	}

	client.Logger.Info("Fetching EventhubNamespace " + ehnDef.Name + " from Azure.")
	ehn, err := client.EventHubNamespaceManager.GetNamespace(ctx, ehnDef.Spec.ResourceGroup, ehnDef.Name)
	if ehn == nil || ehn.Response.Response == nil {
		return reconciler.VerifyError, fmt.Errorf("eventhubnamespace verify was nil for %s", ehnDef.Name)
	} else if ehn.Response.StatusCode == http.StatusNotFound {
		return reconciler.VerifyMissing, nil
	} else if err != nil {
		return reconciler.VerifyError, err
	} else if ehn.Response.StatusCode == http.StatusOK {
		if ehn.ProvisioningState != nil && *ehn.ProvisioningState == "Succeeded" {
			// TODO: handle any cases that lead to VerifyUpdateRequired and VerifyRecreateRequired
			return reconciler.VerifyReady, nil
		} else {
			// TODO: handle cases that lead to VerifyDeleting (what are the undocumented values of *ehn.ProvisioningState?)
			return reconciler.VerifyInProgress, nil
		}
	}

	// we ideally shouldn't get to this point - all cases should be handled explicitly
	return reconciler.VerifyMissing, nil
}

func (client *ResourceManager) Delete(ctx context.Context, r reconciler.ResourceSpec) (reconciler.DeleteResult, error) {
	ehnDef, err := convertInstance(r.Instance)
	if err != nil {
		return reconciler.DeleteError, err
	}

	client.Logger.Info("EventhubNamespace " + ehnDef.Name + " deleting on Azure. Please be patient.")
	resp, err := client.EventHubNamespaceManager.DeleteNamespace(ctx, ehnDef.Spec.ResourceGroup, ehnDef.Name)
	if resp.Response == nil {
		return reconciler.DeleteError, err
	}
	if resp.StatusCode == http.StatusNotFound {
		return reconciler.DeleteAlreadyDeleted, nil
	} else if resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusAccepted {
		return reconciler.DeleteSucceeded, nil
	}

	// TODO: handle all other cases
	return reconciler.DeleteSucceeded, nil
}
