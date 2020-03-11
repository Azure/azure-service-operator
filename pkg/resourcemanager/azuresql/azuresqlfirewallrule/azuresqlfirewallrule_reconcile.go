// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package azuresqlfirewallrule

import (
	"context"
	"fmt"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
)

// Ensure creates a sqlfirewallrule
func (fw *AzureSqlFirewallRuleManager) Ensure(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {
	instance, err := fw.convert(obj)
	if err != nil {
		return false, err
	}

	groupName := instance.Spec.ResourceGroup
	server := instance.Spec.Server
	ruleName := instance.ObjectMeta.Name
	startIP := instance.Spec.StartIPAddress
	endIP := instance.Spec.EndIPAddress

	fwr, err := fw.GetSQLFirewallRule(ctx, groupName, server, ruleName)
	if err == nil {
		instance.Status.Provisioning = false
		instance.Status.Provisioned = true
		instance.Status.Message = resourcemanager.SuccessMsg
		instance.Status.ResourceId = *fwr.ID
		return true, nil
	}
	instance.Status.Message = fmt.Sprintf("AzureSqlFirewallRule Get error %s", err.Error())

	_, err = fw.CreateOrUpdateSQLFirewallRule(ctx, groupName, server, ruleName, startIP, endIP)
	if err != nil {
		instance.Status.Message = err.Error()
		catch := []string{
			errhelp.AsyncOpIncompleteError,
			errhelp.ResourceGroupNotFoundErrorCode,
			errhelp.ParentNotFoundErrorCode,
			errhelp.AlreadyExists,
			errhelp.ResourceNotFound,
		}
		azerr := errhelp.NewAzureErrorAzureError(err)
		if helpers.ContainsString(catch, azerr.Type) {
			return false, nil
		}
		return false, err
	}
	return false, nil
}

// Delete drops a sqlfirewallrule
func (fw *AzureSqlFirewallRuleManager) Delete(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {
	instance, err := fw.convert(obj)
	if err != nil {
		return false, err
	}

	groupName := instance.Spec.ResourceGroup
	server := instance.Spec.Server
	ruleName := instance.ObjectMeta.Name

	err = fw.DeleteSQLFirewallRule(ctx, groupName, server, ruleName)
	if err != nil {
		if errhelp.IsStatusCode204(err) {
			// firewall does not exist
			return true, nil
		}
		if errhelp.IsStatusCode404(err) {
			return false, nil
		}
		instance.Status.Message = fmt.Sprintf("AzureSqlFirewallRule Delete failed with %s", err.Error())
		return false, err
	}
	instance.Status.Message = fmt.Sprintf("Delete AzureSqlFirewallRule succeeded")
	return false, nil
}

// GetParents returns the parents of sqlfirewallrule
func (fw *AzureSqlFirewallRuleManager) GetParents(obj runtime.Object) ([]resourcemanager.KubeParent, error) {
	instance, err := fw.convert(obj)
	if err != nil {
		return nil, err
	}

	return []resourcemanager.KubeParent{
		{
			Key: types.NamespacedName{
				Namespace: instance.Namespace,
				Name:      instance.Spec.Server,
			},
			Target: &azurev1alpha1.AzureSqlServer{},
		},
		{
			Key: types.NamespacedName{
				Namespace: instance.Namespace,
				Name:      instance.Spec.ResourceGroup,
			},
			Target: &azurev1alpha1.ResourceGroup{},
		},
	}, nil
}

func (g *AzureSqlFirewallRuleManager) GetStatus(obj runtime.Object) (*azurev1alpha1.ASOStatus, error) {
	instance, err := g.convert(obj)
	if err != nil {
		return nil, err
	}
	return &instance.Status, nil
}

func (fw *AzureSqlFirewallRuleManager) convert(obj runtime.Object) (*azurev1alpha1.AzureSqlFirewallRule, error) {
	local, ok := obj.(*azurev1alpha1.AzureSqlFirewallRule)
	if !ok {
		return nil, fmt.Errorf("failed type assertion on kind: %s", obj.GetObjectKind().GroupVersionKind().String())
	}
	return local, nil
}
