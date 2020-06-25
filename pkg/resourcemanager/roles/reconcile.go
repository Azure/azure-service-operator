// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package roles

import (
	"context"
	"fmt"
	"net/http"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	"k8s.io/apimachinery/pkg/runtime"
)

// Ensure makes sure that an VNet instance exists
func (r *RoleAssignManager) Ensure(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {

	instance, err := r.convert(obj)
	if err != nil {
		return false, err
	}

	// resourceGroup := instance.Spec.ResourceGroup
	// resourceName := instance.Name

	_, err = r.AssignRole(ctx, instance)
	if err != nil {
		instance.Status.Message = err.Error()
		azerr := errhelp.NewAzureErrorAzureError(err)

		switch azerr.Code {
		case http.StatusBadRequest:
			instance.Status.FailedProvisioning = true
			instance.Status.Provisioning = false
			return true, nil
		case http.StatusNotFound:
			return false, nil
		}
		return true, err
	}

	instance.Status.Message = resourcemanager.SuccessMsg
	instance.Status.Provisioned = true
	instance.Status.Provisioning = false

	return true, nil
}

// Delete makes sure that the roel assignment has been deleted
func (r *RoleAssignManager) Delete(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {
	instance, err := r.convert(obj)
	if err != nil {
		return false, err
	}

	if instance.Status.ResourceId == "" {
		return false, nil
	}

	_, err = r.DeleteRoleAssignment(ctx, instance)
	if err != nil {
		instance.Status.Message = err.Error()
		azerr := errhelp.NewAzureErrorAzureError(err)

		switch azerr.Code {
		case http.StatusNoContent:
			return false, nil
		}
		return true, err
	}

	return false, nil
}

// GetParents lists the parents for a VNet
func (r *RoleAssignManager) GetParents(obj runtime.Object) ([]resourcemanager.KubeParent, error) {
	return []resourcemanager.KubeParent{}, nil
}

func (r *RoleAssignManager) GetStatus(obj runtime.Object) (*azurev1alpha1.ASOStatus, error) {
	instance, err := r.convert(obj)
	if err != nil {
		return nil, err
	}
	return &instance.Status, nil
}

func (r *RoleAssignManager) convert(obj runtime.Object) (*azurev1alpha1.RoleAssign, error) {
	local, ok := obj.(*azurev1alpha1.RoleAssign)
	if !ok {
		return nil, fmt.Errorf("failed type assertion on kind: %s", obj.GetObjectKind().GroupVersionKind().String())
	}
	return local, nil
}
