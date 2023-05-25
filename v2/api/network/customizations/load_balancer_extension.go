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

// Attention: A lot of code in this file is very similar to the logic in route_table_extensions.go and virtual_network_extensions.go.
// The two should be kept in sync as much as possible.

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

	inboundNatRules := &network.LoadBalancersInboundNatRuleList{}
	matchingFields := client.MatchingFields{".metadata.ownerReferences[0]": string(obj.GetUID())}
	err := kubeClient.List(ctx, inboundNatRules, matchingFields)
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

	err = fuzzySetInboundNatRules(armObj.Spec(), armInboundNatRules)
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

// fuzzySetSubnets assigns a collection of subnets to the subnets property of the loadBalancer. Since there are
// many possible ARM API versions and we don't know which one we're using, we cannot do this statically.
// To make matters even more horrible, the type used in the loadBalancer.properties.inboundNatRule property is not the same
// type as used for inboundNatRule.properties (although structurally they are basically the same). To overcome this
// we JSON serialize the subnet and deserialize it into the loadBalancer.properties.inboundNatRule field.
func fuzzySetInboundNatRules(lb genruntime.ARMResourceSpec, inboundNatRules []genruntime.ARMResourceSpec) (err error) {
	if len(inboundNatRules) == 0 {
		// Nothing to do
		return nil
	}

	defer func() {
		if x := recover(); x != nil {
			err = errors.Errorf("caught panic: %s", x)
		}
	}()

	// Here be dragons
	lbValue := reflect.ValueOf(lb)
	lbValue = reflect.Indirect(lbValue)
	if (lbValue == reflect.Value{}) {
		return errors.Errorf("cannot assign to nil loadbalancer")
	}

	propertiesField := lbValue.FieldByName("Properties")
	if (propertiesField == reflect.Value{}) {
		return errors.Errorf("couldn't find properties field on loadbalancer")
	}

	propertiesValue := reflect.Indirect(propertiesField)
	if (propertiesValue == reflect.Value{}) {
		// If the properties field is nil, we must construct an entirely new properties and assign it here
		temp := reflect.New(propertiesField.Type().Elem())
		propertiesField.Set(temp)
		propertiesValue = reflect.Indirect(temp)
	}

	inboundNatRulesField := propertiesValue.FieldByName("InboundNatRules")
	if (inboundNatRulesField == reflect.Value{}) {
		return errors.Errorf("couldn't find InboundNatRules field on LoadBalancer")
	}

	if inboundNatRulesField.Type().Kind() != reflect.Slice {
		return errors.Errorf("inboundNatRules field on LoadBalancer was not of kind Slice")
	}

	elemType := inboundNatRulesField.Type().Elem()
	inboundNatRuleSlice := reflect.MakeSlice(inboundNatRulesField.Type(), 0, 0)

	for _, inboundNatRule := range inboundNatRules {
		embeddedInboundNatRule := reflect.New(elemType)
		err := fuzzySetInboundNatRule(inboundNatRule, embeddedInboundNatRule)
		if err != nil {
			return err
		}

		inboundNatRuleSlice = reflect.Append(inboundNatRuleSlice, reflect.Indirect(embeddedInboundNatRule))
	}

	// Now do the assignment. We do it differently here, as we need to make sure to retain current/updated/deleted InboundNatRules present on the LoadBalancer.
	inboundNatRulesField.Set(reflect.AppendSlice(inboundNatRulesField, inboundNatRuleSlice))

	return nil
}

func fuzzySetInboundNatRule(inboundNatRule genruntime.ARMResourceSpec, embeddedInboundNatRule reflect.Value) error {
	inboundNatRuleJSON, err := json.Marshal(inboundNatRule)
	if err != nil {
		return errors.Wrapf(err, "failed to marshal inboundNatRule json")
	}

	err = json.Unmarshal(inboundNatRuleJSON, embeddedInboundNatRule.Interface())
	if err != nil {
		return errors.Wrapf(err, "failed to unmarshal inboundNatRule JSON")
	}

	// Safety check that these are actually the same:
	// We can't use reflect.DeepEqual because the types are not the same.
	embeddedInboundNatRuleJSON, err := json.Marshal(embeddedInboundNatRule.Interface())
	if err != nil {
		return errors.Wrap(err, "unable to check that embedded inboundNatRule is the same as inboundNatRule")
	}

	err = fuzzyEqualityComparison(embeddedInboundNatRuleJSON, inboundNatRuleJSON)
	if err != nil {
		return errors.Wrap(err, "failed during comparison for embeddedInboundNatRuleJSON and inboundNatRuleJSON")
	}

	return nil
}
