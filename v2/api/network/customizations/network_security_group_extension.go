// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package customizations

import (
	"context"
	"reflect"

	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/conversion"

	network "github.com/Azure/azure-service-operator/v2/api/network/v1api20201101storage"
	. "github.com/Azure/azure-service-operator/v2/internal/logging"
	"github.com/Azure/azure-service-operator/v2/internal/resolver"
	"github.com/Azure/azure-service-operator/v2/internal/util/kubeclient"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/extensions"
)

// Attention: A lot of code in this file is very similar to the logic in load_balancer_extension.go, route_table_extensions.go and virtual_network_extensions.go.
// The two should be kept in sync as much as possible.
// NOTE: This wouldn't work without adding indexes in 'getGeneratedStorageTypes' method in controller_resources.go

var _ extensions.ARMResourceModifier = &NetworkSecurityGroupExtension{}

func (extension *NetworkSecurityGroupExtension) ModifyARMResource(
	ctx context.Context,
	armObj genruntime.ARMResource,
	obj genruntime.ARMMetaObject,
	kubeClient kubeclient.Client,
	resolver *resolver.Resolver,
	log logr.Logger,
) (genruntime.ARMResource, error) {
	typedObj, ok := obj.(*network.NetworkSecurityGroup)
	if !ok {
		return nil, errors.Errorf("cannot run on unknown resource type %T, expected *network.NetworkSecurityGroup", obj)
	}

	// Type assert that we are the hub type. This will fail to compile if
	// the hub type has been changed but this extension has not been updated to match
	var _ conversion.Hub = typedObj

	networkSecurityGroupsSecurityRuleGVK := getNetworkSecurityGroupsSecurityRuleGVK(obj)

	networkSecurityGroupsSecurityRules := &network.NetworkSecurityGroupsSecurityRuleList{}
	matchingFields := client.MatchingFields{".metadata.ownerReferences[0]": string(obj.GetUID())}
	err := kubeClient.List(ctx, networkSecurityGroupsSecurityRules, matchingFields)
	if err != nil {
		return nil, errors.Wrapf(err, "failed listing NetworkSecurityGroupsSecurityRule owned by NetworkSecurityGroup %s/%s", obj.GetNamespace(), obj.GetName())
	}

	armNetworkSecurityGroupsSecurityRule := make([]genruntime.ARMResourceSpec, 0, len(networkSecurityGroupsSecurityRules.Items))
	for _, networkSecurityGroupsSecurityRule := range networkSecurityGroupsSecurityRules.Items {
		networkSecurityGroupsSecurityRule := networkSecurityGroupsSecurityRule

		var transformed genruntime.ARMResourceSpec
		transformed, err = transformToARM(ctx, &networkSecurityGroupsSecurityRule, networkSecurityGroupsSecurityRuleGVK, kubeClient, resolver)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to transform NetworkSecurityGroupsSecurityRules %s/%s", networkSecurityGroupsSecurityRule.GetNamespace(), networkSecurityGroupsSecurityRule.GetName())
		}
		armNetworkSecurityGroupsSecurityRule = append(armNetworkSecurityGroupsSecurityRule, transformed)
	}

	log.V(Info).Info("Found NetworkSecurityGroupsSecurityRule to include on NetworkSecurityGroup", "count", len(armNetworkSecurityGroupsSecurityRule), "names", genruntime.ARMSpecNames(armNetworkSecurityGroupsSecurityRule))

	err = fuzzySetResources(armObj.Spec(), armNetworkSecurityGroupsSecurityRule, "SecurityRules")
	if err != nil {
		return nil, errors.Wrapf(err, "failed to set networkSecurityGroupsSecurityRules")
	}

	return armObj, nil
}

func getNetworkSecurityGroupsSecurityRuleGVK(nsg genruntime.ARMMetaObject) schema.GroupVersionKind {
	gvk := genruntime.GetOriginalGVK(nsg)
	gvk.Kind = reflect.TypeOf(network.NetworkSecurityGroupsSecurityRule{}).Name() // "NetworkSecurityGroupsSecurityRule"

	return gvk
}
