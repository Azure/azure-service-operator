// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package vnet

import (
	"context"
	"fmt"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
)

// Ensure makes sure that an VNet instance exists
func (r *RoleAssignManager) Ensure(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {

	instance, err := r.convert(obj)
	if err != nil {
		return false, err
	}

	// resourceGroup := instance.Spec.ResourceGroup
	// resourceName := instance.Name

	_, err = r.AssignRole(ctx, instance.Spec.PrincipalID, "Contributor")
	if err != nil {
		return true, err
	}

	return false, nil
}

// Delete makes sure that the VNet has been deleted
func (r *RoleAssignManager) Delete(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {

	return false, nil
}

// GetParents lists the parents for a VNet
func (r *RoleAssignManager) GetParents(obj runtime.Object) ([]resourcemanager.KubeParent, error) {

	instance, err := r.convert(obj)
	if err != nil {
		return nil, err
	}

	return []resourcemanager.KubeParent{
		{
			Key: types.NamespacedName{
				Namespace: instance.Namespace,
				Name:      instance.Spec.ResourceGroup,
			},
			Target: &azurev1alpha1.ResourceGroup{},
		},
	}, nil
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
