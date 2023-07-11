// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package customizations

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

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

// Attention: A lot of code in this file is very similar to the logic in network_security_group_extension.go, load_balancer_extensions.go and route_table_extensions.go.
// The two should be kept in sync as much as possible.
// NOTE: This wouldn't work without adding indexes in 'getGeneratedStorageTypes' method in controller_resources.goould be kept in sync as much as possible.

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

	err = fuzzySetResources(armObj.Spec(), armSubnets, "Subnets")
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

// TODO: When we move to Swagger as the source of truth, the type for ownerResource.properties.propertyField and resource.properties
// TODO: may be the same, so we can do away with the JSON serialization part of this assignment.
// fuzzySetResources assigns a collection of subnets to the resources property of the ownerResource. Since there are
// many possible ARM API versions and we don't know which one we're using, we cannot do this statically.
// To make matters even more horrible, the type used in the ownerResource.properties.propertyField property is not the same
// type as used for resource.properties (although structurally they are basically the same). To overcome this
// we JSON serialize the resource and deserialize it into the ownerResource.properties.propertyField field.
func fuzzySetResources(ownerResource genruntime.ARMResourceSpec, resources []genruntime.ARMResourceSpec, propertyField string) (err error) {
	if len(resources) == 0 {
		// Nothing to do
		return nil
	}

	defer func() {
		if x := recover(); x != nil {
			err = errors.Errorf("caught panic: %s", x)
		}
	}()

	// Here be dragons
	ownerValue := reflect.ValueOf(ownerResource)
	ownerValue = reflect.Indirect(ownerValue)
	if (ownerValue == reflect.Value{}) {
		return errors.Errorf("cannot assign to nil %s", strings.ToLower(ownerResource.GetType()))
	}

	propertiesField := ownerValue.FieldByName("Properties")
	if (propertiesField == reflect.Value{}) {
		return errors.Errorf("couldn't find properties field on %s", strings.ToLower(ownerResource.GetType()))
	}

	propertiesValue := reflect.Indirect(propertiesField)
	if (propertiesValue == reflect.Value{}) {
		// If the properties field is nil, we must construct an entirely new properties and assign it here
		temp := reflect.New(propertiesField.Type().Elem())
		propertiesField.Set(temp)
		propertiesValue = reflect.Indirect(temp)
	}

	resourcePropertyField := propertiesValue.FieldByName(propertyField)
	if (resourcePropertyField == reflect.Value{}) {
		return errors.Errorf("couldn't find %s field on %s", propertyField, strings.ToLower(ownerResource.GetType()))
	}

	if resourcePropertyField.Type().Kind() != reflect.Slice {
		return errors.Errorf("%s field on %s was not of kind Slice", propertyField, strings.ToLower(ownerResource.GetType()))
	}

	elemType := resourcePropertyField.Type().Elem()
	resourceSlice := reflect.MakeSlice(resourcePropertyField.Type(), 0, 0)

	for _, resource := range resources {
		embeddedResource := reflect.New(elemType)
		err := fuzzySetResource(resource, embeddedResource)
		if err != nil {
			return err
		}

		resourceSlice = reflect.Append(resourceSlice, reflect.Indirect(embeddedResource))
	}

	// Now do the assignment. We do it differently here, as we need to make sure to retain current/updated/deleted resource present on the ownerResource.
	resourcePropertyField.Set(reflect.AppendSlice(resourcePropertyField, resourceSlice))

	return nil
}

func fuzzySetResource(resource genruntime.ARMResourceSpec, embeddedResource reflect.Value) error {
	resourceJSON, err := json.Marshal(resource)
	if err != nil {
		return errors.Wrapf(err, "failed to marshal %s json", strings.ToLower(resource.GetType()))
	}

	err = json.Unmarshal(resourceJSON, embeddedResource.Interface())
	if err != nil {
		return errors.Wrapf(err, "failed to unmarshal %s JSON", strings.ToLower(resource.GetType()))
	}

	// Safety check that these are actually the same:
	// We can't use reflect.DeepEqual because the types are not the same.
	embeddedResourceJSON, err := json.Marshal(embeddedResource.Interface())
	if err != nil {
		return errors.Wrapf(err, "unable to check that embedded resource is the same as %s", strings.ToLower(resource.GetType()))
	}

	err = fuzzyEqualityComparison(embeddedResourceJSON, resourceJSON)
	if err != nil {
		return errors.Wrapf(err, "failed during comparison for embedded %sJSON and %sJSON", strings.ToLower(resource.GetType()), strings.ToLower(resource.GetType()))
	}

	return nil
}

func fuzzyEqualityComparison(embeddedResourceJSON, resourceJSON []byte) error {
	var embeddedResourceJSONMap map[string]interface{}
	err := json.Unmarshal(embeddedResourceJSON, &embeddedResourceJSONMap)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("unable to unmarshal (%s)", embeddedResourceJSONMap))
	}

	var resourceJSONMap map[string]interface{}
	err = json.Unmarshal(resourceJSON, &resourceJSONMap)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("unable to unmarshal (%s)", resourceJSONMap))
	}

	if !reflect.DeepEqual(embeddedResourceJSONMap, resourceJSONMap) {
		return errors.Errorf(" (%s) != (%s)", string(embeddedResourceJSON), string(resourceJSON))
	}

	return nil
}
