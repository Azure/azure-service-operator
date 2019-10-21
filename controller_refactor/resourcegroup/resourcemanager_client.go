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
	"github.com/Azure/azure-service-operator/controller_refactor"
	"github.com/go-logr/logr"
	"net/http"

	"github.com/Azure/azure-service-operator/pkg/resourcemanager/resourcegroups"

	"k8s.io/apimachinery/pkg/runtime"
)

type ResourceManagerClient struct {
	logger               logr.Logger
	resourceGroupManager resourcegroups.ResourceGroupManager
}

func (client *ResourceManagerClient) Create(ctx context.Context, r runtime.Object) (controller_refactor.EnsureResult, error) {
	rg, err := convertInstance(r)
	if err != nil {
		return controller_refactor.EnsureFailed, err
	}
	_, err = client.resourceGroupManager.CreateGroup(ctx, rg.Name, rg.Spec.Location)

	if err != nil {
		return controller_refactor.EnsureFailed, err
	}
	return controller_refactor.EnsureAwaitingVerification, nil
}

func (client *ResourceManagerClient) Update(ctx context.Context, r runtime.Object) (controller_refactor.EnsureResult, error) {
	return controller_refactor.EnsureFailed, fmt.Errorf("resource group cannot be updated")
}

func (client *ResourceManagerClient) Verify(ctx context.Context, r runtime.Object) (controller_refactor.VerifyResult, error) {
	rg, err := convertInstance(r)
	if err != nil {
		return controller_refactor.VerifyError, err
	}

	// TODO: need to get the object itself to check if it's creating or deleting
	resp, err := client.resourceGroupManager.CheckExistence(ctx, rg.Name)
	if resp.Response != nil && resp.StatusCode == http.StatusNotFound {
		return controller_refactor.VerifyMissing, nil
	}
	if err != nil {
		return controller_refactor.VerifyError, err
	}
	if resp.Response != nil && (resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusNoContent) {
		return controller_refactor.VerifyReady, nil
	}

	return controller_refactor.VerifyMissing, nil
}

func (client *ResourceManagerClient) Delete(ctx context.Context, r runtime.Object) (controller_refactor.DeleteResult, error) {
	rg, err := convertInstance(r)
	if err != nil {
		return controller_refactor.DeleteError, err
	}
	if _, err := client.resourceGroupManager.DeleteGroupAsync(ctx, rg.Name); err == nil {
		return controller_refactor.DeleteAwaitingVerification, nil
	}
	return controller_refactor.DeleteSucceed, nil
}
