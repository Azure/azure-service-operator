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

// Attention: A lot of code in this file is very similar to the logic in network_security_group_extension.go, load_balancer_extensions.go and virtual_network_extensions.go.
// The two should be kept in sync as much as possible.
// NOTE: This wouldn't work without adding indexes in 'getGeneratedStorageTypes' method in controller_resources.go

var _ extensions.ARMResourceModifier = &RouteTableExtension{}

func (extension *RouteTableExtension) ModifyARMResource(
	ctx context.Context,
	armObj genruntime.ARMResource,
	obj genruntime.ARMMetaObject,
	kubeClient kubeclient.Client,
	resolver *resolver.Resolver,
	log logr.Logger,
) (genruntime.ARMResource, error) {
	typedObj, ok := obj.(*network.RouteTable)
	if !ok {
		return nil, errors.Errorf("cannot run on unknown resource type %T", obj)
	}

	// Type assert that we are the hub type. This will fail to compile if
	// the hub type has been changed but this extension has not been updated to match
	var _ conversion.Hub = typedObj

	routeGVK := getRouteGVK(obj)

	// We use namespace + owner name here rather than something slightly more specific like actual owner UUID
	// in order to account for cases where the subresources + owner were just created and so the subsresources haven't
	// actually been assigned to the owner yet. See https://github.com/Azure/azure-service-operator/issues/3077 for more
	// details.
	matchingOwner := client.MatchingFields{".spec.owner.name": obj.GetName()}
	matchingNamespace := client.InNamespace(obj.GetNamespace())

	routes := &network.RouteTablesRouteList{}
	err := kubeClient.List(ctx, routes, matchingOwner, matchingNamespace)
	if err != nil {
		return nil, errors.Wrapf(err, "failed listing routes owned by RouteTable %s/%s", obj.GetNamespace(), obj.GetName())
	}

	armRoutes := make([]genruntime.ARMResourceSpec, 0, len(routes.Items))
	for _, route := range routes.Items {
		route := route

		var transformed genruntime.ARMResourceSpec
		transformed, err = transformToARM(ctx, &route, routeGVK, kubeClient, resolver)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to transform route %s/%s", route.GetNamespace(), route.GetName())
		}
		armRoutes = append(armRoutes, transformed)
	}

	log.V(Info).Info("Found routes to include on RouteTable", "count", len(armRoutes), "names", genruntime.ARMSpecNames(armRoutes))

	err = fuzzySetResources(armObj.Spec(), armRoutes, "Routes")
	if err != nil {
		return nil, errors.Wrapf(err, "failed to set routes")
	}

	return armObj, nil
}

func getRouteGVK(routeTable genruntime.ARMMetaObject) schema.GroupVersionKind {
	gvk := genruntime.GetOriginalGVK(routeTable)
	gvk.Kind = reflect.TypeOf(network.RouteTablesRoute{}).Name() // "RouteTableRoute"

	return gvk
}
