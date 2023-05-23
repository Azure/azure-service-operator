// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package customizations

import (
	"context"
	"encoding/json"
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

// Attention: A lot of code in this file is very similar to the logic in load_balancer_extensions.go and virtual_network_extensions.go.
// The two should be kept in sync as much as possible.

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

	routes := &network.RouteTablesRouteList{}
	matchingFields := client.MatchingFields{".metadata.ownerReferences[0]": string(obj.GetUID())}
	err := kubeClient.List(ctx, routes, matchingFields)
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

	err = fuzzySetRoutes(armObj.Spec(), armRoutes)
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

// TODO: When we move to Swagger as the source of truth, the type for routetable.properties.routes and routtableroutes.properties
// TODO: may be the same, so we can do away with the JSON serialization part of this assignment.
// fuzzySetRoutes assigns a collection of routes to the routes property of the route table. Since there are
// many possible ARM API versions and we don't know which one we're using, we cannot do this statically.
// To make matters even more horrible, the type used in the routetable.properties.routes is not the same
// type as used for routes.properties (although structurally they are basically the same). To overcome this
// we JSON serialize the route and deserialize it into the routetableroutes.properties.routes field.
func fuzzySetRoutes(routeTable genruntime.ARMResourceSpec, routes []genruntime.ARMResourceSpec) (err error) {
	if len(routes) == 0 {
		// Nothing to do
		return nil
	}

	defer func() {
		if x := recover(); x != nil {
			err = errors.Errorf("caught panic: %s", x)
		}
	}()

	// Here be dragons
	routeTableValue := reflect.ValueOf(routeTable)
	routeTableValue = reflect.Indirect(routeTableValue)

	if (routeTableValue == reflect.Value{}) {
		return errors.Errorf("cannot assign to nil route table")
	}

	propertiesField := routeTableValue.FieldByName("Properties")
	if (propertiesField == reflect.Value{}) {
		return errors.Errorf("couldn't find properties field on route table")
	}

	propertiesValue := reflect.Indirect(propertiesField)
	if (propertiesValue == reflect.Value{}) {
		// If the properties field is nil, we must construct an entirely new properties and assign it here
		temp := reflect.New(propertiesField.Type().Elem())
		propertiesField.Set(temp)
		propertiesValue = reflect.Indirect(temp)
	}

	routesField := propertiesValue.FieldByName("Routes")
	if (routesField == reflect.Value{}) {
		return errors.Errorf("couldn't find routes field on route table")
	}

	if routesField.Type().Kind() != reflect.Slice {
		return errors.Errorf("routes field on route table was not of kind Slice")
	}

	elemType := routesField.Type().Elem()
	routesSlice := reflect.MakeSlice(routesField.Type(), 0, 0)

	for _, route := range routes {
		embeddedRoute := reflect.New(elemType)
		err = fuzzySetRoute(route, embeddedRoute)
		if err != nil {
			return err
		}

		routesSlice = reflect.Append(routesSlice, reflect.Indirect(embeddedRoute))
	}

	// Now do the assignment
	routesField.Set(routesSlice)

	return err
}

func fuzzySetRoute(route genruntime.ARMResourceSpec, embeddedRoute reflect.Value) error {
	routeJSON, err := json.Marshal(route)
	if err != nil {
		return errors.Wrapf(err, "failed to marshal route json")
	}

	err = json.Unmarshal(routeJSON, embeddedRoute.Interface())
	if err != nil {
		return errors.Wrapf(err, "failed to unmarshal route JSON")
	}

	// Safety check that these are actually the same:
	// We can't use reflect.DeepEqual because the types are not the same.
	embeddedRouteJSON, err := json.Marshal(embeddedRoute.Interface())
	if err != nil {
		return errors.Wrap(err, "unable to check that embedded route is the same as route")
	}
	if string(embeddedRouteJSON) != string(routeJSON) {
		return errors.Errorf("embeddedRouteJSON (%s) != routeJSON (%s)", string(embeddedRouteJSON), string(routeJSON))
	}

	return nil
}
