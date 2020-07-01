// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package resourcegroups

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/Azure/azure-service-operator/api/v1alpha1"
	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	helpers "github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	"k8s.io/apimachinery/pkg/runtime"
)

func (g *AzureResourceGroupManager) Ensure(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {
	options := &resourcemanager.Options{}
	for _, opt := range opts {
		opt(options)
	}

	instance, err := g.convert(obj)
	if err != nil {
		return false, err
	}

	instance.Status.Provisioning = true

	group, err := g.CreateGroupWithCreds(ctx, instance, options.Credential)
	if err != nil {
		instance.Status.Message = err.Error()
		azerr := errhelp.NewAzureErrorAzureError(err)

		// this happens when op isnt complete, just requeue
		if strings.Contains(azerr.Type, errhelp.AsyncOpIncompleteError) {
			return false, nil
		}

		instance.Status.Provisioning = false

		// handle special cases that won't work without a change to spec
		if group.StatusCode == http.StatusBadRequest {
			instance.Status.FailedProvisioning = true
			return true, nil
		}

		return false, fmt.Errorf("ResourceGroup create error %v", err)
	}

	instance.Status.Provisioned = true
	instance.Status.Provisioning = false
	instance.Status.FailedProvisioning = false
	instance.Status.Message = resourcemanager.SuccessMsg
	instance.Status.ResourceId = *group.ID

	return true, nil
}

func (g *AzureResourceGroupManager) Delete(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {
	options := &resourcemanager.Options{}
	for _, opt := range opts {
		opt(options)
	}

	instance, err := g.convert(obj)
	if err != nil {
		return false, err
	}

	resourcegroup := instance.ObjectMeta.Name

	_, err = g.DeleteGroupWithCreds(ctx, resourcegroup, options.Credential)
	if err != nil {

		catch := []string{
			errhelp.ResourceGroupNotFoundErrorCode,
			errhelp.AsyncOpIncompleteError,
		}
		err = errhelp.NewAzureError(err)
		if azerr, ok := err.(*errhelp.AzureError); ok {
			if helpers.ContainsString(catch, azerr.Type) {
				return false, nil
			}
		}

		return true, fmt.Errorf("ResourceGroup delete error %v", err)

	}

	return true, nil
}

func (g *AzureResourceGroupManager) GetParents(obj runtime.Object) ([]resourcemanager.KubeParent, error) {
	return nil, nil
}

func (g *AzureResourceGroupManager) GetStatus(obj runtime.Object) (*v1alpha1.ASOStatus, error) {
	instance, err := g.convert(obj)
	if err != nil {
		return nil, err
	}
	return &instance.Status, nil
}

func (g *AzureResourceGroupManager) convert(obj runtime.Object) (*azurev1alpha1.ResourceGroup, error) {
	local, ok := obj.(*azurev1alpha1.ResourceGroup)
	if !ok {
		return nil, fmt.Errorf("failed type assertion on kind: %s", obj.GetObjectKind().GroupVersionKind().String())
	}
	return local, nil
}
