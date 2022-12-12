// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20220301

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type Serverfarm_Spec_ARM struct {
	ExtendedLocation *ExtendedLocation_ARM `json:"extendedLocation,omitempty"`

	// Kind: Kind of resource.
	Kind *string `json:"kind,omitempty"`

	// Location: Resource Location.
	Location *string `json:"location,omitempty"`
	Name     string  `json:"name,omitempty"`

	// Properties: AppServicePlan resource specific properties
	Properties *Serverfarm_Properties_Spec_ARM `json:"properties,omitempty"`
	Sku        *SkuDescription_ARM             `json:"sku,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &Serverfarm_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2022-03-01"
func (serverfarm Serverfarm_Spec_ARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (serverfarm *Serverfarm_Spec_ARM) GetName() string {
	return serverfarm.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Web/serverfarms"
func (serverfarm *Serverfarm_Spec_ARM) GetType() string {
	return "Microsoft.Web/serverfarms"
}

type ExtendedLocation_ARM struct {
	// Name: Name of extended location.
	Name *string `json:"name,omitempty"`
}

type Serverfarm_Properties_Spec_ARM struct {
	// ElasticScaleEnabled: ServerFarm supports ElasticScale. Apps in this plan will scale as if the ServerFarm was
	// ElasticPremium sku
	ElasticScaleEnabled *bool `json:"elasticScaleEnabled,omitempty"`

	// FreeOfferExpirationTime: The time when the server farm free offer expires.
	FreeOfferExpirationTime *string `json:"freeOfferExpirationTime,omitempty"`

	// HostingEnvironmentProfile: Specification for the App Service Environment to use for the App Service plan.
	HostingEnvironmentProfile *HostingEnvironmentProfile_ARM `json:"hostingEnvironmentProfile,omitempty"`

	// HyperV: If Hyper-V container app service plan <code>true</code>, <code>false</code> otherwise.
	HyperV *bool `json:"hyperV,omitempty"`

	// IsSpot: If <code>true</code>, this App Service Plan owns spot instances.
	IsSpot *bool `json:"isSpot,omitempty"`

	// IsXenon: Obsolete: If Hyper-V container app service plan <code>true</code>, <code>false</code> otherwise.
	IsXenon *bool `json:"isXenon,omitempty"`

	// KubeEnvironmentProfile: Specification for the Kubernetes Environment to use for the App Service plan.
	KubeEnvironmentProfile *KubeEnvironmentProfile_ARM `json:"kubeEnvironmentProfile,omitempty"`

	// MaximumElasticWorkerCount: Maximum number of total workers allowed for this ElasticScaleEnabled App Service Plan
	MaximumElasticWorkerCount *int `json:"maximumElasticWorkerCount,omitempty"`

	// PerSiteScaling: If <code>true</code>, apps assigned to this App Service plan can be scaled independently.
	// If <code>false</code>, apps assigned to this App Service plan will scale to all instances of the plan.
	PerSiteScaling *bool `json:"perSiteScaling,omitempty"`

	// Reserved: If Linux app service plan <code>true</code>, <code>false</code> otherwise.
	Reserved *bool `json:"reserved,omitempty"`

	// SpotExpirationTime: The time when the server farm expires. Valid only if it is a spot server farm.
	SpotExpirationTime *string `json:"spotExpirationTime,omitempty"`

	// TargetWorkerCount: Scaling worker count.
	TargetWorkerCount *int `json:"targetWorkerCount,omitempty"`

	// TargetWorkerSizeId: Scaling worker size ID.
	TargetWorkerSizeId *int `json:"targetWorkerSizeId,omitempty"`

	// WorkerTierName: Target worker tier assigned to the App Service plan.
	WorkerTierName *string `json:"workerTierName,omitempty"`

	// ZoneRedundant: If <code>true</code>, this App Service Plan will perform availability zone balancing.
	// If <code>false</code>, this App Service Plan will not perform availability zone balancing.
	ZoneRedundant *bool `json:"zoneRedundant,omitempty"`
}

type SkuDescription_ARM struct {
	// Capabilities: Capabilities of the SKU, e.g., is traffic manager enabled?
	Capabilities []Capability_ARM `json:"capabilities,omitempty"`

	// Capacity: Current number of instances assigned to the resource.
	Capacity *int `json:"capacity,omitempty"`

	// Family: Family code of the resource SKU.
	Family *string `json:"family,omitempty"`

	// Locations: Locations of the SKU.
	Locations []string `json:"locations,omitempty"`

	// Name: Name of the resource SKU.
	Name *string `json:"name,omitempty"`

	// Size: Size specifier of the resource SKU.
	Size *string `json:"size,omitempty"`

	// SkuCapacity: Min, max, and default scale values of the SKU.
	SkuCapacity *SkuCapacity_ARM `json:"skuCapacity,omitempty"`

	// Tier: Service tier of the resource SKU.
	Tier *string `json:"tier,omitempty"`
}

type Capability_ARM struct {
	// Name: Name of the SKU capability.
	Name *string `json:"name,omitempty"`

	// Reason: Reason of the SKU capability.
	Reason *string `json:"reason,omitempty"`

	// Value: Value of the SKU capability.
	Value *string `json:"value,omitempty"`
}

type HostingEnvironmentProfile_ARM struct {
	Id *string `json:"id,omitempty"`
}

type KubeEnvironmentProfile_ARM struct {
	Id *string `json:"id,omitempty"`
}

type SkuCapacity_ARM struct {
	// Default: Default number of workers for this App Service plan SKU.
	Default *int `json:"default,omitempty"`

	// ElasticMaximum: Maximum number of Elastic workers for this App Service plan SKU.
	ElasticMaximum *int `json:"elasticMaximum,omitempty"`

	// Maximum: Maximum number of workers for this App Service plan SKU.
	Maximum *int `json:"maximum,omitempty"`

	// Minimum: Minimum number of workers for this App Service plan SKU.
	Minimum *int `json:"minimum,omitempty"`

	// ScaleType: Available scale configurations for an App Service plan.
	ScaleType *string `json:"scaleType,omitempty"`
}
