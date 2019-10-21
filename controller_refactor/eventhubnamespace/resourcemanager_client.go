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

//import (
//	"context"
//	"fmt"
//	"github.com/Azure/azure-service-operator/controller_refactor"
//	"net/http"
//
//	"github.com/Azure/azure-service-operator/api/v1alpha1"
//	"github.com/Azure/azure-service-operator/pkg/resourcemanager/resourcegroups"
//
//	"k8s.io/apimachinery/pkg/runtime"
//)
//
//type ResourceManagerClient struct {
//	ResourceGroupManager resourcegroups.ResourceGroupManager
//}
//
//func (client *ResourceManagerClient) Create(ctx context.Context, r runtime.Object) (controller_refactor.EnsureResult, error) {
//	rg, err := client.convert(r)
//	if err != nil {
//		return controller_refactor.EnsureFailed, err
//	}
//	_, err = client.ResourceGroupManager.CreateGroup(ctx, rg.Name, rg.Spec.Location)
//
//	if err != nil {
//		return controller_refactor.EnsureFailed, err
//	}
//	return controller_refactor.EnsureAwaitingVerification, nil
//}
//
//func (client *ResourceManagerClient) Update(ctx context.Context, r runtime.Object) (controller_refactor.EnsureResult, error) {
//	return controller_refactor.EnsureFailed, fmt.Errorf("resource group cannon be updated")
//}
//
//func (client *ResourceManagerClient) Verify(ctx context.Context, r runtime.Object) (controller_refactor.VerifyResult, error) {
//	rg, err := client.convert(r)
//	if err != nil {
//		return controller_refactor.VerifyError, err
//	}
//	resp, err := client.ResourceGroupManager.CheckExistence(ctx, rg.Name)
//	if resp.Response != nil && resp.StatusCode == http.StatusNotFound {
//		return controller_refactor.VerifyMissing, nil
//	}
//	if err != nil {
//		return controller_refactor.VerifyError, err
//	}
//	if resp.Response != nil && (resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusNoContent) {
//		return controller_refactor.VerifyReady, nil
//	}
//
//	return controller_refactor.VerifyMissing, nil
//}
//
//func (client *ResourceManagerClient) Delete(ctx context.Context, r runtime.Object) (controller_refactor.DeleteResult, error) {
//	rg, err := client.convert(r)
//	if err != nil {
//		return controller_refactor.DeleteError, err
//	}
//	if _, err := client.ResourceGroupManager.DeleteGroupAsync(ctx, rg.Name); err == nil {
//		return controller_refactor.DeleteAwaitingVerification, nil
//	}
//	return controller_refactor.DeleteSucceed, nil
//}
//
//func (_ *ResourceManagerClient) convert(obj runtime.Object) (*v1alpha1.ResourceGroup, error) {
//	local, ok := obj.(*v1alpha1.ResourceGroup)
//	if !ok {
//		return nil, fmt.Errorf("failed type assertion on kind: %s", obj.GetObjectKind().GroupVersionKind().String())
//	}
//	return local, nil
//}
