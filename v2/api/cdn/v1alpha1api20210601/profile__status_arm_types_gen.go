// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210601

// Deprecated version of Profile_Status. Use v1beta20210601.Profile_Status instead
type Profile_StatusARM struct {
	Id         *string                      `json:"id,omitempty"`
	Kind       *string                      `json:"kind,omitempty"`
	Location   *string                      `json:"location,omitempty"`
	Name       *string                      `json:"name,omitempty"`
	Properties *ProfileProperties_StatusARM `json:"properties,omitempty"`
	Sku        *Sku_StatusARM               `json:"sku,omitempty"`
	SystemData *SystemData_StatusARM        `json:"systemData,omitempty"`
	Tags       map[string]string            `json:"tags,omitempty"`
	Type       *string                      `json:"type,omitempty"`
}

// Deprecated version of ProfileProperties_Status. Use v1beta20210601.ProfileProperties_Status instead
type ProfileProperties_StatusARM struct {
	FrontDoorId                  *string                                   `json:"frontDoorId,omitempty"`
	OriginResponseTimeoutSeconds *int                                      `json:"originResponseTimeoutSeconds,omitempty"`
	ProvisioningState            *ProfilePropertiesStatusProvisioningState `json:"provisioningState,omitempty"`
	ResourceState                *ProfilePropertiesStatusResourceState     `json:"resourceState,omitempty"`
}

// Deprecated version of Sku_Status. Use v1beta20210601.Sku_Status instead
type Sku_StatusARM struct {
	Name *SkuStatusName `json:"name,omitempty"`
}

// Deprecated version of SkuStatusName. Use v1beta20210601.SkuStatusName instead
type SkuStatusName string

const (
	SkuStatusNameCustomVerizon                    = SkuStatusName("Custom_Verizon")
	SkuStatusNamePremiumAzureFrontDoor            = SkuStatusName("Premium_AzureFrontDoor")
	SkuStatusNamePremiumVerizon                   = SkuStatusName("Premium_Verizon")
	SkuStatusNameStandard955BandWidthChinaCdn     = SkuStatusName("Standard_955BandWidth_ChinaCdn")
	SkuStatusNameStandardAkamai                   = SkuStatusName("Standard_Akamai")
	SkuStatusNameStandardAvgBandWidthChinaCdn     = SkuStatusName("Standard_AvgBandWidth_ChinaCdn")
	SkuStatusNameStandardAzureFrontDoor           = SkuStatusName("Standard_AzureFrontDoor")
	SkuStatusNameStandardChinaCdn                 = SkuStatusName("Standard_ChinaCdn")
	SkuStatusNameStandardMicrosoft                = SkuStatusName("Standard_Microsoft")
	SkuStatusNameStandardPlus955BandWidthChinaCdn = SkuStatusName("StandardPlus_955BandWidth_ChinaCdn")
	SkuStatusNameStandardPlusAvgBandWidthChinaCdn = SkuStatusName("StandardPlus_AvgBandWidth_ChinaCdn")
	SkuStatusNameStandardPlusChinaCdn             = SkuStatusName("StandardPlus_ChinaCdn")
	SkuStatusNameStandardVerizon                  = SkuStatusName("Standard_Verizon")
)
