// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package customizations

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	v1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/conversion"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"

	network "github.com/Azure/azure-service-operator/v2/api/network/v1api20220701/storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/configmaps"

	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	. "github.com/Azure/azure-service-operator/v2/internal/logging"
	"github.com/Azure/azure-service-operator/v2/internal/resolver"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/extensions"
)

var _ extensions.PostReconciliationChecker = &PrivateEndpointExtension{}

func (extension *PrivateEndpointExtension) PostReconcileCheck(
	_ context.Context,
	obj genruntime.MetaObject,
	_ genruntime.MetaObject,
	_ *resolver.Resolver,
	_ *genericarmclient.GenericClient,
	_ logr.Logger,
	_ extensions.PostReconcileCheckFunc,
) (extensions.PostReconcileCheckResult, error) {
	endpoint, ok := obj.(*network.PrivateEndpoint)
	if !ok {
		return extensions.PostReconcileCheckResult{},
			errors.Errorf("cannot run on unknown resource type %T, expected *network.PrivateEndpoint", obj)
	}

	// Type assert that we are the hub type. This will fail to compile if
	// the hub type has been changed but this extension has not
	var _ conversion.Hub = endpoint

	var reqApprovals []string
	// We want to check `ManualPrivateLinkServiceConnections` as these are the ones which are not auto-approved.
	if connections := endpoint.Status.ManualPrivateLinkServiceConnections; connections != nil {
		for _, connection := range connections {
			if *connection.PrivateLinkServiceConnectionState.Status != "Approved" {
				reqApprovals = append(reqApprovals, *connection.Id)
			}
		}
	}

	if len(reqApprovals) > 0 {
		// Returns 'conditions.NewReadyConditionImpactingError' error
		return extensions.PostReconcileCheckResultFailure(
			fmt.Sprintf(
				"Private connection(s) '%q' to the PrivateEndpoint requires approval",
				reqApprovals)), nil
	}

	return extensions.PostReconcileCheckResultSuccess(), nil
}

func (extension *PrivateEndpointExtension) ExportKubernetesResources(
	ctx context.Context,
	obj genruntime.MetaObject,
	armClient *genericarmclient.GenericClient,
	log logr.Logger,
) ([]client.Object, error) {
	// This has to be the current hub storage version. It will need to be updated
	// if the hub storage version changes.
	endpoint, ok := obj.(*network.PrivateEndpoint)
	if !ok {
		return nil, errors.Errorf("cannot run on unknown resource type %T, expected *network.PrivateEndpoint", obj)
	}

	// Type assert that we are the hub type. This will fail to compile if
	// the hub type has been changed but this extension has not
	var _ conversion.Hub = endpoint

	hasIpConfiguration := configMapSpecified(endpoint)
	if !hasIpConfiguration {
		log.V(Debug).Info("no configmap retrieval to perform as operatorSpec is empty")
		return nil, nil
	}

	if endpoint.Status.NetworkInterfaces == nil || len(endpoint.Status.NetworkInterfaces) == 0 {
		log.V(Debug).Info("no configmap retrieval to perform as there is no NetworkInterfaces attached")
		return nil, nil
	}

	nicId, err := arm.ParseResourceID(*endpoint.Status.NetworkInterfaces[0].Id)
	if err != nil {
		return nil, err
	}

	// The default primary ip configuration for PrivateEndpoint is on NetworkInterfaceController. Hence, we fetch it from there.
	var interfacesClient *armnetwork.InterfacesClient
	interfacesClient, err = armnetwork.NewInterfacesClient(nicId.SubscriptionID, armClient.Creds(), armClient.ClientOptions())
	if err != nil {
		return nil, errors.Wrapf(err, "failed to create new PrivateEndpointsClient")
	}

	var resp armnetwork.InterfacesClientGetResponse
	resp, err = interfacesClient.Get(ctx, nicId.ResourceGroupName, nicId.Name, nil)
	if err != nil {
		return nil, errors.Wrapf(err, "failed getting PrivateEndpoint")
	}

	configsByName := configByName(resp)
	configs, err := configMapToWrite(endpoint, configsByName)
	if err != nil {
		return nil, err
	}

	return configmaps.SliceToClientObjectSlice(configs), nil
}

func configByName(resp armnetwork.InterfacesClientGetResponse) map[string]string {
	result := make(map[string]string)

	if resp.Properties != nil && resp.Properties.IPConfigurations != nil && len(resp.Properties.IPConfigurations) > 0 {
		result["PrimaryNicPrivateIPAddress"] = *resp.Properties.IPConfigurations[0].Properties.PrivateIPAddress
	}

	return result
}

func configMapToWrite(obj *network.PrivateEndpoint, configs map[string]string) ([]*v1.ConfigMap, error) {
	operatorSpecConfigs := obj.Spec.OperatorSpec.ConfigMaps
	if operatorSpecConfigs == nil {
		return nil, errors.Errorf("unexpected nil operatorspec")
	}

	collector := configmaps.NewCollector(obj.Namespace)
	
	primaryNicPrivateIPAddress, ok := configs["PrimaryNicPrivateIPAddress"]
	if ok {
		collector.AddValue(operatorSpecConfigs.PrimaryNicPrivateIPAddress, primaryNicPrivateIPAddress)
	}

	return collector.Values()
}

func configMapSpecified(endpoint *network.PrivateEndpoint) bool {
	hasIpConfiguration := false
	configMaps := endpoint.Spec.OperatorSpec.ConfigMaps

	if configMaps != nil && configMaps.PrimaryNicPrivateIPAddress != nil {
		hasIpConfiguration = true
	}

	return hasIpConfiguration
}
