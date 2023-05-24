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
	"github.com/Azure/azure-service-operator/v2/internal/reconcilers"
	"github.com/Azure/azure-service-operator/v2/internal/resolver"
	"github.com/Azure/azure-service-operator/v2/internal/util/kubeclient"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/extensions"
)

// Attention: A lot of code in this file is very similar to the logic in load_balancer_extensions.go and route_table_extensions.go.
// The two should be kept in sync as much as possible.

var _ extensions.ARMResourceModifier = &VirtualNetworkExtension{}

func (extension *VirtualNetworkExtension) ModifyARMResource(
	ctx context.Context,
	armObj genruntime.ARMResource,
	obj genruntime.ARMMetaObject,
	kubeClient kubeclient.Client,
	resolver *resolver.Resolver,
	log logr.Logger,
) (genruntime.ARMResource, error) {
	typedObj, ok := obj.(*network.VirtualNetwork)
	if !ok {
		return nil, errors.Errorf("cannot run on unknown resource type %T, expected *network.VirtualNetwork", obj)
	}

	// Type assert that we are the hub type. This will fail to compile if
	// the hub type has been changed but this extension has not been updated to match
	var _ conversion.Hub = typedObj

	subnetGVK := getSubnetGVK(obj)

	subnets := &network.VirtualNetworksSubnetList{}
	matchingFields := client.MatchingFields{".metadata.ownerReferences[0]": string(obj.GetUID())}
	err := kubeClient.List(ctx, subnets, matchingFields)
	if err != nil {
		return nil, errors.Wrapf(err, "failed listing subnets owned by VNET %s/%s", obj.GetNamespace(), obj.GetName())
	}

	armSubnets := make([]genruntime.ARMResourceSpec, 0, len(subnets.Items))
	for _, subnet := range subnets.Items {
		subnet := subnet

		var transformed genruntime.ARMResourceSpec
		transformed, err = transformToARM(ctx, &subnet, subnetGVK, kubeClient, resolver)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to transform subnet %s/%s", subnet.GetNamespace(), subnet.GetName())
		}
		armSubnets = append(armSubnets, transformed)
	}

	log.V(Info).Info("Found subnets to include on VNET", "count", len(armSubnets), "names", genruntime.ARMSpecNames(armSubnets))

	err = fuzzySetSubnets(armObj.Spec(), armSubnets)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to set subnets")
	}

	return armObj, nil
}

func getSubnetGVK(vnet genruntime.ARMMetaObject) schema.GroupVersionKind {
	gvk := genruntime.GetOriginalGVK(vnet)
	gvk.Kind = reflect.TypeOf(network.VirtualNetworksSubnet{}).Name() // "VirtualNetworksSubnet"

	return gvk
}

func transformToARM(
	ctx context.Context,
	obj genruntime.ARMMetaObject,
	gvk schema.GroupVersionKind,
	kubeClient kubeclient.Client,
	resolver *resolver.Resolver,
) (genruntime.ARMResourceSpec, error) {
	spec, err := genruntime.GetVersionedSpecFromGVK(obj, kubeClient.Scheme(), gvk)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get spec from %s/%s", obj.GetNamespace(), obj.GetName())
	}

	armTransformer, ok := spec.(genruntime.ARMTransformer)
	if !ok {
		return nil, errors.Errorf("spec was of type %T which doesn't implement genruntime.ArmTransformer", spec)
	}

	_, resolvedDetails, err := resolver.ResolveAll(ctx, obj)
	if err != nil {
		return nil, reconcilers.ClassifyResolverError(err)
	}

	armSpec, err := armTransformer.ConvertToARM(resolvedDetails)
	if err != nil {
		return nil, errors.Wrapf(err, "transforming resource %s to ARM", obj.GetName())
	}

	typedARMSpec, ok := armSpec.(genruntime.ARMResourceSpec)
	if !ok {
		return nil, errors.Errorf("casting armSpec of type %T to genruntime.ARMResourceSpec", armSpec)
	}

	return typedARMSpec, nil
}

// TODO: When we move to Swagger as the source of truth, the type for vnet.properties.subnets and subnet.properties
// TODO: may be the same, so we can do away with the JSON serialization part of this assignment.
// fuzzySetSubnets assigns a collection of subnets to the subnets property of the vnet. Since there are
// many possible ARM API versions and we don't know which one we're using, we cannot do this statically.
// To make matters even more horrible, the type used in the vnet.properties.subnets property is not the same
// type as used for subnet.properties (although structurally they are basically the same). To overcome this
// we JSON serialize the subnet and deserialize it into the vnet.properties.subnets field.
func fuzzySetSubnets(vnet genruntime.ARMResourceSpec, subnets []genruntime.ARMResourceSpec) (err error) {
	if len(subnets) == 0 {
		// Nothing to do
		return nil
	}

	defer func() {
		if x := recover(); x != nil {
			err = errors.Errorf("caught panic: %s", x)
		}
	}()

	// Here be dragons
	vnetValue := reflect.ValueOf(vnet)
	vnetValue = reflect.Indirect(vnetValue)
	if (vnetValue == reflect.Value{}) {
		return errors.Errorf("cannot assign to nil vnet")
	}

	propertiesField := vnetValue.FieldByName("Properties")
	if (propertiesField == reflect.Value{}) {
		return errors.Errorf("couldn't find properties field on vnet")
	}

	propertiesValue := reflect.Indirect(propertiesField)
	if (propertiesValue == reflect.Value{}) {
		// If the properties field is nil, we must construct an entirely new properties and assign it here
		temp := reflect.New(propertiesField.Type().Elem())
		propertiesField.Set(temp)
		propertiesValue = reflect.Indirect(temp)
	}

	subnetsField := propertiesValue.FieldByName("Subnets")
	if (subnetsField == reflect.Value{}) {
		return errors.Errorf("couldn't find subnets field on vnet")
	}

	if subnetsField.Type().Kind() != reflect.Slice {
		return errors.Errorf("subnets field on vnet was not of kind Slice")
	}

	elemType := subnetsField.Type().Elem()
	subnetSlice := reflect.MakeSlice(subnetsField.Type(), 0, 0)

	for _, subnet := range subnets {
		embeddedSubnet := reflect.New(elemType)
		err := fuzzySetSubnet(subnet, embeddedSubnet)
		if err != nil {
			return err
		}

		subnetSlice = reflect.Append(subnetSlice, reflect.Indirect(embeddedSubnet))
	}

	// Now do the assignment
	subnetsField.Set(subnetSlice)

	return nil
}

func fuzzySetSubnet(subnet genruntime.ARMResourceSpec, embeddedSubnet reflect.Value) error {
	subnetJSON, err := json.Marshal(subnet)
	if err != nil {
		return errors.Wrapf(err, "failed to marshal subnet json")
	}

	err = json.Unmarshal(subnetJSON, embeddedSubnet.Interface())
	if err != nil {
		return errors.Wrapf(err, "failed to unmarshal subnet JSON")
	}

	// Safety check that these are actually the same:
	// We can't use reflect.DeepEqual because the types are not the same.
	embeddedSubnetJSON, err := json.Marshal(embeddedSubnet.Interface())
	if err != nil {
		return errors.Wrap(err, "unable to check that embedded subnet is the same as subnet")
	}

	var embeddedSubnetJSONMap map[string]interface{}
	err = json.Unmarshal(embeddedSubnetJSON, &embeddedSubnetJSONMap)
	if err != nil {
		return errors.Wrap(err, "unable to check that embedded subnet is the same as subnet")
	}
	var subnetJSONMap map[string]interface{}
	err = json.Unmarshal(subnetJSON, &subnetJSONMap)
	if err != nil {
		return errors.Wrap(err, "unable to check that embedded subnet is the same as subnet")
	}

	if !reflect.DeepEqual(embeddedSubnetJSONMap, subnetJSONMap) {
		return errors.Errorf("embeddedSubnetJSON (%s) != subnetJSON (%s)", string(embeddedSubnetJSON), string(subnetJSON))
	}

	return nil
}
