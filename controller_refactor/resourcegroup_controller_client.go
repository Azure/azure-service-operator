package controller_refactor

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/resourcegroups"

	"k8s.io/apimachinery/pkg/runtime"
)

type ResourceGroupClient struct {
	ResourceGroupManager resourcegroups.ResourceGroupManager
}

func (client *ResourceGroupClient) Create(ctx context.Context, r runtime.Object) (EnsureResult, error) {
	rg, err := client.convert(r)
	if err != nil {
		return EnsureFailed, err
	}
	_, err = client.ResourceGroupManager.CreateGroup(ctx, rg.Name, rg.Spec.Location)

	if err != nil {
		return EnsureFailed, err
	}
	return EnsureAwaitingVerification, nil
}

func (client *ResourceGroupClient) Update(ctx context.Context, r runtime.Object) (EnsureResult, error) {
	return EnsureFailed, fmt.Errorf("resource group cannon be updated")
}

func (client *ResourceGroupClient) Verify(ctx context.Context, r runtime.Object) (VerifyResult, error) {
	rg, err := client.convert(r)
	if err != nil {
		return VerifyError, err
	}
	resp, err := client.ResourceGroupManager.CheckExistence(ctx, rg.Name)
	if resp.Response != nil && resp.StatusCode == http.StatusNotFound {
		return VerifyMissing, nil
	}
	if err != nil {
		return VerifyError, err
	}
	if resp.Response != nil && (resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusNoContent) {
		return VerifyReady, nil
	}

	return VerifyMissing, nil
}

func (client *ResourceGroupClient) Delete(ctx context.Context, r runtime.Object) (DeleteResult, error) {
	rg, err := client.convert(r)
	if err != nil {
		return DeleteError, err
	}
	if _, err := client.ResourceGroupManager.DeleteGroupAsync(ctx, rg.Name); err == nil {
		return DeleteAwaitingVerification, nil
	}
	return DeleteSucceed, nil
}

func (_ *ResourceGroupClient) convert(obj runtime.Object) (*v1alpha1.ResourceGroup, error) {
	local, ok := obj.(*v1alpha1.ResourceGroup)
	if !ok {
		return nil, fmt.Errorf("failed type assertion on kind: %s", obj.GetObjectKind().GroupVersionKind().String())
	}
	return local, nil
}
