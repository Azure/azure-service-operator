/*
Copyright 2019 microsoft.
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
func (fw *AzureSqlFirewallRuleManager) Ensure(ctx context.Context, obj runtime.Object) (bool, error) {
	instance, err := fw.convert(obj)
	if err != nil {
		return false, err
	}

	groupName := instance.Spec.ResourceGroup
	server := instance.Spec.Server
	ruleName := instance.ObjectMeta.Name
	startIP := instance.Spec.StartIPAddress
	endIP := instance.Spec.EndIPAddress

	_, err = fw.CreateOrUpdateSQLFirewallRule(ctx, groupName, server, ruleName, startIP, endIP)
	if err != nil {
		catch := []string{
			errhelp.AsyncOpIncompleteError,
			errhelp.ResourceGroupNotFoundErrorCode,
		}
		azerr := errhelp.NewAzureErrorAzureError(err)
		if helpers.ContainsString(catch, azerr.Type) {
			return false, nil
		}
		return true, fmt.Errorf("AzureSqlFirewallRule CreateOrUpdate error %v", err)
	}

	resp, err := fw.GetSQLFirewallRule(ctx, groupName, server, ruleName)
	if err != nil {
		return true, fmt.Errorf("AzureSqlFirewallRule GetSQLFirewallRule error %v", err)
	}

	instance.Status.Provisioning = false
	instance.Status.Provisioned = true
	instance.Status.State = string(resp.Status)
	instance.Status.Message = "Success"

	return true, nil
}

// Delete drops a sqlfirewallrule
func (fw *AzureSqlFirewallRuleManager) Delete(ctx context.Context, obj runtime.Object) (bool, error) {
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
			return false, nil
		}
		return true, fmt.Errorf("Azure SqlFirewallRule Delete error %v", err)
	}

	return true, nil
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

func (fw *AzureSqlFirewallRuleManager) convert(obj runtime.Object) (*azurev1alpha1.AzureSqlFirewallRule, error) {
	local, ok := obj.(*azurev1alpha1.AzureSqlFirewallRule)
	if !ok {
		return nil, fmt.Errorf("failed type assertion on kind: %s", obj.GetObjectKind().GroupVersionKind().String())
	}
	return local, nil
}
