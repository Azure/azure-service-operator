// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package namespacenetworkrule

import (
	"context"
	"fmt"
	"strings"

	"github.com/Azure/azure-sdk-for-go/services/eventhub/mgmt/2017-04-01/eventhub"
	"github.com/Azure/azure-service-operator/api/v1alpha1"
	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
)

// Ensure creates a namespace network rule
func (vr *AzureNamespaceNetworkRuleManager) Ensure(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {
	instance, err := vr.convert(obj)
	if err != nil {
		return false, err
	}

	// set a spec hash if one hasn't been set
	hash := helpers.Hash256(instance.Spec)
	if instance.Status.SpecHash == hash && instance.Status.Provisioned {
		instance.Status.RequestedAt = nil
		return true, nil
	}

	if instance.Status.SpecHash == "" {
		instance.Status.SpecHash = hash
	}

	groupName := instance.Spec.ResourceGroup
	namespace := instance.Spec.Namespace

	networkRuleSet := ParseNetworkRules(instance.Spec.DefaultAction, instance.Spec.VirtualNetworkRules, instance.Spec.IPRules)

	/*
		nwruleset, err := vr.GetNetworkRuleSet(ctx, groupName, namespace)
		if err == nil {
			if instance.Status.SpecHash == hash { //Signifies Get is successful for the spec we requested
				if nwruleset.NetworkRuleSetProperties != nil {
					instance.Status.Provisioning = false
					instance.Status.Provisioned = true
					instance.Status.Message = resourcemanager.SuccessMsg
					instance.Status.ResourceId = *nwruleset.ID
					return true, nil

				}
				return false, nil
			}
		}*/

	instance.Status.Provisioning = true
	nwruleset, err := vr.CreateNetworkRuleSet(ctx, groupName, namespace, networkRuleSet)
	if err != nil {
		instance.Status.Message = err.Error()
		azerr := errhelp.NewAzureErrorAzureError(err)

		ignorableErrors := []string{
			errhelp.ResourceGroupNotFoundErrorCode,
			errhelp.ParentNotFoundErrorCode,
			errhelp.ResourceNotFound,
		}

		unrecoverableErrors := []string{
			errhelp.BadRequest, // error we get when namespace is "Basic" SKU
		}

		if helpers.ContainsString(ignorableErrors, azerr.Type) {
			instance.Status.Provisioning = false
			return false, nil //requeue
		}

		if helpers.ContainsString(unrecoverableErrors, azerr.Type) {
			instance.Status.Provisioning = false
			return true, nil // Stop reconciliation
		}

		return false, err
	}

	instance.Status.SpecHash = hash // Reset the Hash in status to the hash of the spec
	instance.Status.Provisioning = false
	instance.Status.Provisioned = true
	instance.Status.Message = resourcemanager.SuccessMsg
	instance.Status.ResourceId = *nwruleset.ID
	return true, nil
}

// Delete drops a namespace network rule
func (vr *AzureNamespaceNetworkRuleManager) Delete(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {
	instance, err := vr.convert(obj)
	if err != nil {
		return false, err
	}

	groupName := instance.Spec.ResourceGroup
	namespace := instance.Spec.Namespace

	_, err = vr.DeleteNetworkRuleSet(ctx, groupName, namespace)
	if err != nil {
		instance.Status.Message = err.Error()

		azerr := errhelp.NewAzureErrorAzureError(err)
		// these errors are expected
		ignore := []string{
			errhelp.AsyncOpIncompleteError,
		}

		// this means the resource doesn't exist
		finished := []string{
			errhelp.ResourceNotFound,
			errhelp.ParentNotFoundErrorCode,
			errhelp.ResourceGroupNotFoundErrorCode,
			errhelp.BadRequest, // get this if we had created this with namespace tier as "Basic"
		}

		if helpers.ContainsString(ignore, azerr.Type) {
			return true, nil //requeue
		}

		if helpers.ContainsString(finished, azerr.Type) {
			return false, nil //end reconcile
		}
		return false, err
	}

	return false, nil
}

// GetParents returns the parents of namespace network rule
func (vr *AzureNamespaceNetworkRuleManager) GetParents(obj runtime.Object) ([]resourcemanager.KubeParent, error) {
	instance, err := vr.convert(obj)
	if err != nil {
		return nil, err
	}

	return []resourcemanager.KubeParent{
		{
			Key: types.NamespacedName{
				Namespace: instance.Namespace,
				Name:      instance.Spec.Namespace,
			},
			Target: &azurev1alpha1.EventhubNamespace{},
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

func (vr *AzureNamespaceNetworkRuleManager) GetStatus(obj runtime.Object) (*azurev1alpha1.ASOStatus, error) {
	instance, err := vr.convert(obj)
	if err != nil {
		return nil, err
	}
	return &instance.Status, nil
}

func (vr *AzureNamespaceNetworkRuleManager) convert(obj runtime.Object) (*azurev1alpha1.EventhubNamespaceNetworkRule, error) {
	local, ok := obj.(*azurev1alpha1.EventhubNamespaceNetworkRule)
	if !ok {
		return nil, fmt.Errorf("failed type assertion on kind: %s", obj.GetObjectKind().GroupVersionKind().String())
	}
	return local, nil
}

// ParseNetworkRules parses the network rules specified in the yaml and converts to the SDK struct
func ParseNetworkRules(defAction v1alpha1.DefaultAction, vNetRules *[]v1alpha1.VirtualNetworkRules, ipRules *[]v1alpha1.IPRules) eventhub.NetworkRuleSet {

	defaultAction := eventhub.Deny

	switch strings.ToLower(defAction) {
	case "allow":
		defaultAction = eventhub.Allow
	case "deny":
		defaultAction = eventhub.Deny
	default:
		defaultAction = eventhub.Deny
	}

	var vNetRulesSet []eventhub.NWRuleSetVirtualNetworkRules
	if vNetRules != nil {
		for _, i := range *vNetRules {
			subnetID := i.SubnetID
			ignoreEndpoint := i.IgnoreMissingServiceEndpoint
			vNetRulesSet = append(vNetRulesSet, eventhub.NWRuleSetVirtualNetworkRules{
				Subnet: &eventhub.Subnet{
					ID: &subnetID,
				},
				IgnoreMissingVnetServiceEndpoint: &ignoreEndpoint,
			})
		}
	}

	var ipRulesSet []eventhub.NWRuleSetIPRules
	if ipRules != nil {
		for _, i := range *ipRules {
			ipmask := i.IPMask
			ipRulesSet = append(ipRulesSet, eventhub.NWRuleSetIPRules{
				IPMask: ipmask,
				Action: eventhub.NetworkRuleIPActionAllow, // only applicable value
			})
		}
	}

	return eventhub.NetworkRuleSet{
		NetworkRuleSetProperties: &eventhub.NetworkRuleSetProperties{
			DefaultAction:       defaultAction,
			VirtualNetworkRules: &vNetRulesSet,
			IPRules:             &ipRulesSet,
		},
	}

}
