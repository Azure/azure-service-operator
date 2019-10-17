package controller_refactor

import (
	"context"
	"fmt"
	"github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/resourcegroups"
	"k8s.io/apimachinery/pkg/runtime"
	"net/http"
)

type ResourceGroupClient struct {
	ResourceGroupManager resourcegroups.ResourceGroupManager
}

func (client *ResourceGroupClient) Ensure(ctx context.Context, r runtime.Object) error {
	rg, err := client.convert(r)
	if err != nil {
		return err
	}
	_, err = client.ResourceGroupManager.CreateGroup(ctx, rg.Name, rg.Spec.Location)
	return err
}

func (client *ResourceGroupClient) Verify(ctx context.Context, r runtime.Object) (VerifyResult, error) {
	rg, err := client.convert(r)
	if err != nil {
		return Invalid, err
	}
	resp, err := client.ResourceGroupManager.CheckExistence(ctx, rg.Name)
	if err != nil {
		return Invalid, err
	}
	if resp.Response != nil && (resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusNoContent) {
		return Ready, nil
	}

	return Missing, nil
}

func (client *ResourceGroupClient) Delete(ctx context.Context, r runtime.Object) error {
	rg, err := client.convert(r)
	if err != nil {
		return err
	}
	if resp, err := client.ResourceGroupManager.DeleteGroup(ctx, rg.Name); err == nil || resp.IsHTTPStatus(http.StatusNotFound) {
		return nil
	}
	return nil
}

func (_ *ResourceGroupClient) convert(obj runtime.Object) (*v1alpha1.ResourceGroup, error) {
	local, ok := obj.(*v1alpha1.ResourceGroup)
	if !ok {
		return nil, fmt.Errorf("failed type assertion on kind: %s", obj.GetObjectKind().GroupVersionKind().String())
	}
	return local, nil
}
