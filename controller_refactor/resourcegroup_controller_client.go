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

func (client *ResourceGroupClient) Create(ctx context.Context, r runtime.Object) error {
	rg, err := client.convert(r)
	if err != nil {
		return err
	}
	_, err = client.ResourceGroupManager.CreateGroup(ctx, rg.Name, rg.Spec.Location)
	return err
}

func (client *ResourceGroupClient) Validate(ctx context.Context, r runtime.Object) (bool, error) {
	rg, err := client.convert(r)
	if err != nil {
		return false, err
	}
	resp, err := client.ResourceGroupManager.CheckExistence(ctx, rg.Name)
	if err != nil {
		return false, err
	}
	return resp.Response != nil && (resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusNoContent), nil
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

