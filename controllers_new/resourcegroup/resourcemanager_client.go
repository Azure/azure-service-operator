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

package resourcegroup

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/azure-service-operator/pkg/reconciler"
	"github.com/go-logr/logr"

	"github.com/Azure/azure-service-operator/pkg/resourcemanager/resourcegroups"

	"k8s.io/apimachinery/pkg/runtime"
)

type ResourceManagerClient struct {
	Logger               logr.Logger
	ResourceGroupManager resourcegroups.ResourceGroupManager
}

func CreateResourceManagerClient(resourceGroupManager resourcegroups.ResourceGroupManager, logger logr.Logger) ResourceManagerClient {
	return ResourceManagerClient{
		Logger:               logger,
		ResourceGroupManager: resourceGroupManager,
	}
}

func (client *ResourceManagerClient) Create(ctx context.Context, r runtime.Object) (reconciler.EnsureResult, error) {
	rg, err := convertInstance(r)
	if err != nil {
		return reconciler.EnsureError, err
	}
	_, err = client.ResourceGroupManager.CreateGroup(ctx, rg.Name, rg.Spec.Location)

	if err != nil {
		return reconciler.EnsureError, err
	}
	return reconciler.EnsureAwaitingVerification, nil
}

func (client *ResourceManagerClient) Update(ctx context.Context, r runtime.Object) (reconciler.EnsureResult, error) {
	return reconciler.EnsureError, fmt.Errorf("resource group cannot be updated")
}

func (client *ResourceManagerClient) Verify(ctx context.Context, r runtime.Object) (reconciler.VerifyResult, error) {
	rg, err := convertInstance(r)
	if err != nil {
		return reconciler.VerifyError, err
	}

	resp, err := client.ResourceGroupManager.GetGroup(ctx, rg.Name)
	if resp.Response.Response == nil || resp.StatusCode == http.StatusNotFound || resp.Properties.ProvisioningState == nil {
		return reconciler.VerifyMissing, nil
	} else if err != nil {
		return reconciler.VerifyError, err
	}
	if resp.StatusCode == http.StatusOK {
		switch *resp.Properties.ProvisioningState {
		// TODO: capture other ProvisioningStates
		case "Deleting":
			return reconciler.VerifyDeleting, nil
		default:
			return reconciler.VerifyReady, nil
		}
	}

	// we probably shouldn't get to this point
	return reconciler.VerifyMissing, nil
}

func (client *ResourceManagerClient) Delete(ctx context.Context, r runtime.Object) (reconciler.DeleteResult, error) {
	rg, err := convertInstance(r)
	if err != nil {
		return reconciler.DeleteError, err
	}
	if _, err := client.ResourceGroupManager.DeleteGroupAsync(ctx, rg.Name); err == nil {
		return reconciler.DeleteAwaitingVerification, nil
	}
	return reconciler.DeleteSucceed, nil
}
