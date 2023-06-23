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

// Attention: A lot of code in this file is very similar to the logic in network_security_group_extension.go, route_table_extensions.go and virtual_network_extensions.go.
// The two should be kept in sync as much as possible.
// NOTE: This wouldn't work without adding indexes in 'getGeneratedStorageTypes' method in controller_resources.go

var _ extensions.ARMResourceModifier = &LoadBalancerExtension{}

func (extension *LoadBalancerExtension) ModifyARMResource(
	ctx context.Context,
	armObj genruntime.ARMResource,
	obj genruntime.ARMMetaObject,
	kubeClient kubeclient.Client,
	resolver *resolver.Resolver,
	log logr.Logger,
) (genruntime.ARMResource, error) {
	typedObj, ok := obj.(*network.LoadBalancer)
	if !ok {
		return nil, errors.Errorf("cannot run on unknown resource type %T, expected *network.LoadBalancer", obj)
	}

	// Type assert that we are the hub type. This will fail to compile if
	// the hub type has been changed but this extension has not been updated to match
	var _ conversion.Hub = typedObj

	inboundNatRuleGVK := getInboundNatRuleGVK(obj)

	// We use namespace + owner name here rather than something slightly more specific like actual owner UUID
	// in order to account for cases where the subresources + owner were just created and so the subsresources haven't
	// actually been assigned to the owner yet. See https://github.com/Azure/azure-service-operator/issues/3077 for more
	// details.
	matchingOwner := client.MatchingFields{".spec.owner.name": obj.GetName()}
	matchingNamespace := client.InNamespace(obj.GetNamespace())

	inboundNatRules := &network.LoadBalancersInboundNatRuleList{}
	err := kubeClient.List(ctx, inboundNatRules, matchingOwner, matchingNamespace)
	if err != nil {
		return nil, errors.Wrapf(err, "failed listing InboundNatRules owned by LoadBalancer %s/%s", obj.GetNamespace(), obj.GetName())
	}

	armInboundNatRules := make([]genruntime.ARMResourceSpec, 0, len(inboundNatRules.Items))
	for _, inboundNatRule := range inboundNatRules.Items {
		inboundNatRule := inboundNatRule

		var transformed genruntime.ARMResourceSpec
		transformed, err = transformToARM(ctx, &inboundNatRule, inboundNatRuleGVK, kubeClient, resolver)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to transform InboundNatRules %s/%s", inboundNatRule.GetNamespace(), inboundNatRule.GetName())
		}
		armInboundNatRules = append(armInboundNatRules, transformed)
	}

	log.V(Info).Info("Found InboundNatRules to include on LoadBalancer", "count", len(armInboundNatRules), "names", genruntime.ARMSpecNames(armInboundNatRules))

	err = fuzzySetResources(armObj.Spec(), armInboundNatRules, "InboundNatRules")
	if err != nil {
		return nil, errors.Wrapf(err, "failed to set InboundNatRules")
	}

	return armObj, nil
}

func getInboundNatRuleGVK(lb genruntime.ARMMetaObject) schema.GroupVersionKind {
	gvk := genruntime.GetOriginalGVK(lb)
	gvk.Kind = reflect.TypeOf(network.LoadBalancersInboundNatRule{}).Name() // "LoadBalancersInboundNatRule"

	return gvk
}
