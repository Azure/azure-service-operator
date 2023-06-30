// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package rediscachefirewallrules

import (
	"context"
	"fmt"

	"github.com/Azure/azure-service-operator/api/v1alpha1"
	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
)

// Ensure creates a rediscachefirewallrule
func (fw *AzureRedisCacheFirewallRuleManager) Ensure(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {
	instance, err := fw.convert(obj)
	if err != nil {
		return true, err
	}

	fwr, err := fw.Get(ctx, instance.Spec.ResourceGroup, instance.Spec.CacheName, instance.ObjectMeta.Name)
	if err == nil {
		instance.Status.Provisioned = true
		instance.Status.Provisioning = false
		instance.Status.Message = resourcemanager.SuccessMsg
		instance.Status.ResourceId = *fwr.ID
		instance.Status.State = fwr.Status
		return true, nil
	}

	instance.Status.Provisioning = true
	instance.Status.FailedProvisioning = false

	_, err = fw.CreateRedisCacheFirewallRule(ctx, *instance)
	if err != nil {
		instance.Status.Message = err.Error()
		instance.Status.Provisioning = false
		azerr := errhelp.NewAzureError(err)

		inProgress := []string{
			errhelp.RequestConflictError,
			errhelp.AsyncOpIncompleteError,
		}
		ignorableErr := []string{
			errhelp.ParentNotFoundErrorCode,
			errhelp.ResourceGroupNotFoundErrorCode,
			errhelp.ResourceNotFound,
		}
		fatalErr := []string{
			errhelp.BadRequest,
		}

		if helpers.ContainsString(inProgress, azerr.Type) {
			instance.Status.Provisioning = true
			instance.Status.Message = "RedisCacheFirewallRule exists but may not be ready"
			return false, nil
		}
		if helpers.ContainsString(ignorableErr, azerr.Type) {
			instance.Status.Provisioning = false
			return false, nil
		}
		if helpers.ContainsString(fatalErr, azerr.Type) {
			// serious error occured, end reconcilliation and mark it as failed
			instance.Status.Message = fmt.Sprintf("Error occurred creating the RedisCacheFirewallRule: %s", errhelp.StripErrorIDs(err))
			instance.Status.Provisioned = false
			instance.Status.Provisioning = false
			instance.Status.FailedProvisioning = true
			return true, nil
		}
		return false, nil
	}
	return false, nil
}

// Delete removes a rediscachefirewallrule
func (fw *AzureRedisCacheFirewallRuleManager) Delete(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {
	instance, err := fw.convert(obj)
	if err != nil {
		return true, err
	}

	_, err = fw.DeleteRedisCacheFirewallRule(ctx, instance.Spec.ResourceGroup, instance.Spec.CacheName, instance.ObjectMeta.Name)
	if err != nil {
		instance.Status.Message = err.Error()
		azerr := errhelp.NewAzureError(err)

		ignorableErr := []string{
			errhelp.AsyncOpIncompleteError,
		}

		finished := []string{
			errhelp.ResourceNotFound,
		}
		if helpers.ContainsString(ignorableErr, azerr.Type) {
			return true, nil
		}
		if helpers.ContainsString(finished, azerr.Type) {
			return false, nil
		}

		instance.Status.Message = fmt.Sprintf("AzureSqlFailoverGroup Delete failed with: %s", err.Error())
		return false, err
	}
	// successful return
	instance.Status.Message = "Delete AzureSqlFailoverGroup succeeded"
	return false, err
}

// GetParents returns the parents of rediscachefirewallrule
func (fw *AzureRedisCacheFirewallRuleManager) GetParents(obj runtime.Object) ([]resourcemanager.KubeParent, error) {
	instance, err := fw.convert(obj)
	if err != nil {
		return nil, err
	}

	return []resourcemanager.KubeParent{
		{
			Key: types.NamespacedName{
				Namespace: instance.Namespace,
				Name:      instance.Spec.CacheName,
			},
			Target: &v1alpha1.RedisCache{},
		},
		{
			Key: types.NamespacedName{
				Namespace: instance.Namespace,
				Name:      instance.Spec.ResourceGroup,
			},
			Target: &v1alpha1.ResourceGroup{},
		},
	}, nil
}

// GetStatus gets the ASOStatus
func (fw *AzureRedisCacheFirewallRuleManager) GetStatus(obj runtime.Object) (*azurev1alpha1.ASOStatus, error) {
	instance, err := fw.convert(obj)
	if err != nil {
		return nil, err
	}
	return &instance.Status, nil
}

func (fw *AzureRedisCacheFirewallRuleManager) convert(obj runtime.Object) (*v1alpha1.RedisCacheFirewallRule, error) {
	local, ok := obj.(*v1alpha1.RedisCacheFirewallRule)
	if !ok {
		return nil, fmt.Errorf("failed type assertion on kind: %s", obj.GetObjectKind().GroupVersionKind().String())
	}
	return local, nil
}
