// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package storage

import (
	"github.com/Azure/azure-service-operator/v2/api/machinelearningservices/v1api20240401/storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

var _ augmentConversionForAKS_Properties = &AKS_Properties{}

func (aks *AKS_Properties) AssignPropertiesFrom(src *storage.AKS_Properties) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(src.PropertyBag)

	if src.LoadBalancerSubnetReference != nil {
		isNotKubeRef := !src.LoadBalancerSubnetReference.IsKubernetesReference()
		// Note that using isNotKubeRef is a bit awkward because in reality it shouldn't be possible to have a genruntime.ResourceReference with no
		// kube ref AND no ARM ref, but if that does happen we pass the empty-string along to maintain the round-trip invariant.
		if len(src.LoadBalancerSubnetReference.ARMID) > 0 || isNotKubeRef {
			aks.LoadBalancerSubnet = &src.LoadBalancerSubnetReference.ARMID
			propertyBag.Remove("LoadBalancerSubnetReference") // Remove it from property bag added by code generated code
		}
		// No need to handle the other case, as we would just put it into the property bag, which was already done by the generated code
	}

	return nil
}

func (aks *AKS_Properties) AssignPropertiesTo(dst *storage.AKS_Properties) error {
	// Clone the existing property bag
	dstPropertyBag := genruntime.NewPropertyBag(dst.PropertyBag)

	// ProximityPlacementGroupID
	if aks.LoadBalancerSubnet != nil {
		dst.LoadBalancerSubnetReference = &genruntime.ResourceReference{
			ARMID: *aks.LoadBalancerSubnet,
		}
	}
	// Ensure that this field is not set aks the destination property bag (it shouldn't ever be there)
	dstPropertyBag.Remove("LoadBalancerSubnet")

	return nil
}
