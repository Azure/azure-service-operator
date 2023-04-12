// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20201201

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

// Deprecated version of VirtualMachineScaleSet_Spec. Use v1api20201201.VirtualMachineScaleSet_Spec instead
type VirtualMachineScaleSet_Spec_ARM struct {
	ExtendedLocation *ExtendedLocation_ARM                 `json:"extendedLocation,omitempty"`
	Identity         *VirtualMachineScaleSetIdentity_ARM   `json:"identity,omitempty"`
	Location         *string                               `json:"location,omitempty"`
	Name             string                                `json:"name,omitempty"`
	Plan             *Plan_ARM                             `json:"plan,omitempty"`
	Properties       *VirtualMachineScaleSetProperties_ARM `json:"properties,omitempty"`
	Sku              *Sku_ARM                              `json:"sku,omitempty"`
	Tags             map[string]string                     `json:"tags,omitempty"`
	Zones            []string                              `json:"zones,omitempty"`
}

var _ genruntime.ARMResourceSpec = &VirtualMachineScaleSet_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-12-01"
func (scaleSet VirtualMachineScaleSet_Spec_ARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (scaleSet *VirtualMachineScaleSet_Spec_ARM) GetName() string {
	return scaleSet.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Compute/virtualMachineScaleSets"
func (scaleSet *VirtualMachineScaleSet_Spec_ARM) GetType() string {
	return "Microsoft.Compute/virtualMachineScaleSets"
}

// Deprecated version of ExtendedLocation. Use v1api20201201.ExtendedLocation instead
type ExtendedLocation_ARM struct {
	Name *string               `json:"name,omitempty"`
	Type *ExtendedLocationType `json:"type,omitempty"`
}

// Deprecated version of Plan. Use v1api20201201.Plan instead
type Plan_ARM struct {
	Name          *string `json:"name,omitempty"`
	Product       *string `json:"product,omitempty"`
	PromotionCode *string `json:"promotionCode,omitempty"`
	Publisher     *string `json:"publisher,omitempty"`
}

// Deprecated version of Sku. Use v1api20201201.Sku instead
type Sku_ARM struct {
	Capacity *int    `json:"capacity,omitempty"`
	Name     *string `json:"name,omitempty"`
	Tier     *string `json:"tier,omitempty"`
}

// Deprecated version of VirtualMachineScaleSetIdentity. Use v1api20201201.VirtualMachineScaleSetIdentity instead
type VirtualMachineScaleSetIdentity_ARM struct {
	Type                   *VirtualMachineScaleSetIdentity_Type       `json:"type,omitempty"`
	UserAssignedIdentities map[string]UserAssignedIdentityDetails_ARM `json:"userAssignedIdentities,omitempty"`
}

// Deprecated version of VirtualMachineScaleSetProperties. Use v1api20201201.VirtualMachineScaleSetProperties instead
type VirtualMachineScaleSetProperties_ARM struct {
	AdditionalCapabilities                 *AdditionalCapabilities_ARM          `json:"additionalCapabilities,omitempty"`
	AutomaticRepairsPolicy                 *AutomaticRepairsPolicy_ARM          `json:"automaticRepairsPolicy,omitempty"`
	DoNotRunExtensionsOnOverprovisionedVMs *bool                                `json:"doNotRunExtensionsOnOverprovisionedVMs,omitempty"`
	HostGroup                              *SubResource_ARM                     `json:"hostGroup,omitempty"`
	OrchestrationMode                      *OrchestrationMode                   `json:"orchestrationMode,omitempty"`
	Overprovision                          *bool                                `json:"overprovision,omitempty"`
	PlatformFaultDomainCount               *int                                 `json:"platformFaultDomainCount,omitempty"`
	ProximityPlacementGroup                *SubResource_ARM                     `json:"proximityPlacementGroup,omitempty"`
	ScaleInPolicy                          *ScaleInPolicy_ARM                   `json:"scaleInPolicy,omitempty"`
	SinglePlacementGroup                   *bool                                `json:"singlePlacementGroup,omitempty"`
	UpgradePolicy                          *UpgradePolicy_ARM                   `json:"upgradePolicy,omitempty"`
	VirtualMachineProfile                  *VirtualMachineScaleSetVMProfile_ARM `json:"virtualMachineProfile,omitempty"`
	ZoneBalance                            *bool                                `json:"zoneBalance,omitempty"`
}

// Deprecated version of AdditionalCapabilities. Use v1api20201201.AdditionalCapabilities instead
type AdditionalCapabilities_ARM struct {
	UltraSSDEnabled *bool `json:"ultraSSDEnabled,omitempty"`
}

// Deprecated version of AutomaticRepairsPolicy. Use v1api20201201.AutomaticRepairsPolicy instead
type AutomaticRepairsPolicy_ARM struct {
	Enabled     *bool   `json:"enabled,omitempty"`
	GracePeriod *string `json:"gracePeriod,omitempty"`
}

// Deprecated version of ExtendedLocationType. Use v1api20201201.ExtendedLocationType instead
// +kubebuilder:validation:Enum={"EdgeZone"}
type ExtendedLocationType string

const ExtendedLocationType_EdgeZone = ExtendedLocationType("EdgeZone")

// Deprecated version of ScaleInPolicy. Use v1api20201201.ScaleInPolicy instead
type ScaleInPolicy_ARM struct {
	Rules []ScaleInPolicy_Rules `json:"rules,omitempty"`
}

// Deprecated version of SubResource. Use v1api20201201.SubResource instead
type SubResource_ARM struct {
	Id *string `json:"id,omitempty"`
}

// Deprecated version of UpgradePolicy. Use v1api20201201.UpgradePolicy instead
type UpgradePolicy_ARM struct {
	AutomaticOSUpgradePolicy *AutomaticOSUpgradePolicy_ARM `json:"automaticOSUpgradePolicy,omitempty"`
	Mode                     *UpgradePolicy_Mode           `json:"mode,omitempty"`
	RollingUpgradePolicy     *RollingUpgradePolicy_ARM     `json:"rollingUpgradePolicy,omitempty"`
}

// Information about the user assigned identity for the resource
type UserAssignedIdentityDetails_ARM struct {
}

// Deprecated version of VirtualMachineScaleSetIdentity_Type. Use v1api20201201.VirtualMachineScaleSetIdentity_Type instead
// +kubebuilder:validation:Enum={"None","SystemAssigned","SystemAssigned, UserAssigned","UserAssigned"}
type VirtualMachineScaleSetIdentity_Type string

const (
	VirtualMachineScaleSetIdentity_Type_None                       = VirtualMachineScaleSetIdentity_Type("None")
	VirtualMachineScaleSetIdentity_Type_SystemAssigned             = VirtualMachineScaleSetIdentity_Type("SystemAssigned")
	VirtualMachineScaleSetIdentity_Type_SystemAssignedUserAssigned = VirtualMachineScaleSetIdentity_Type("SystemAssigned, UserAssigned")
	VirtualMachineScaleSetIdentity_Type_UserAssigned               = VirtualMachineScaleSetIdentity_Type("UserAssigned")
)

// Deprecated version of VirtualMachineScaleSetVMProfile. Use v1api20201201.VirtualMachineScaleSetVMProfile instead
type VirtualMachineScaleSetVMProfile_ARM struct {
	BillingProfile         *BillingProfile_ARM                         `json:"billingProfile,omitempty"`
	DiagnosticsProfile     *DiagnosticsProfile_ARM                     `json:"diagnosticsProfile,omitempty"`
	EvictionPolicy         *EvictionPolicy                             `json:"evictionPolicy,omitempty"`
	ExtensionProfile       *VirtualMachineScaleSetExtensionProfile_ARM `json:"extensionProfile,omitempty"`
	LicenseType            *string                                     `json:"licenseType,omitempty"`
	NetworkProfile         *VirtualMachineScaleSetNetworkProfile_ARM   `json:"networkProfile,omitempty"`
	OsProfile              *VirtualMachineScaleSetOSProfile_ARM        `json:"osProfile,omitempty"`
	Priority               *Priority                                   `json:"priority,omitempty"`
	ScheduledEventsProfile *ScheduledEventsProfile_ARM                 `json:"scheduledEventsProfile,omitempty"`
	SecurityProfile        *SecurityProfile_ARM                        `json:"securityProfile,omitempty"`
	StorageProfile         *VirtualMachineScaleSetStorageProfile_ARM   `json:"storageProfile,omitempty"`
}

// Deprecated version of AutomaticOSUpgradePolicy. Use v1api20201201.AutomaticOSUpgradePolicy instead
type AutomaticOSUpgradePolicy_ARM struct {
	DisableAutomaticRollback *bool `json:"disableAutomaticRollback,omitempty"`
	EnableAutomaticOSUpgrade *bool `json:"enableAutomaticOSUpgrade,omitempty"`
}

// Deprecated version of RollingUpgradePolicy. Use v1api20201201.RollingUpgradePolicy instead
type RollingUpgradePolicy_ARM struct {
	EnableCrossZoneUpgrade              *bool   `json:"enableCrossZoneUpgrade,omitempty"`
	MaxBatchInstancePercent             *int    `json:"maxBatchInstancePercent,omitempty"`
	MaxUnhealthyInstancePercent         *int    `json:"maxUnhealthyInstancePercent,omitempty"`
	MaxUnhealthyUpgradedInstancePercent *int    `json:"maxUnhealthyUpgradedInstancePercent,omitempty"`
	PauseTimeBetweenBatches             *string `json:"pauseTimeBetweenBatches,omitempty"`
	PrioritizeUnhealthyInstances        *bool   `json:"prioritizeUnhealthyInstances,omitempty"`
}

// Deprecated version of ScheduledEventsProfile. Use v1api20201201.ScheduledEventsProfile instead
type ScheduledEventsProfile_ARM struct {
	TerminateNotificationProfile *TerminateNotificationProfile_ARM `json:"terminateNotificationProfile,omitempty"`
}

// Deprecated version of VirtualMachineScaleSetExtensionProfile. Use v1api20201201.VirtualMachineScaleSetExtensionProfile instead
type VirtualMachineScaleSetExtensionProfile_ARM struct {
	Extensions           []VirtualMachineScaleSetExtension_ARM `json:"extensions,omitempty"`
	ExtensionsTimeBudget *string                               `json:"extensionsTimeBudget,omitempty"`
}

// Deprecated version of VirtualMachineScaleSetNetworkProfile. Use v1api20201201.VirtualMachineScaleSetNetworkProfile instead
type VirtualMachineScaleSetNetworkProfile_ARM struct {
	HealthProbe                    *ApiEntityReference_ARM                          `json:"healthProbe,omitempty"`
	NetworkInterfaceConfigurations []VirtualMachineScaleSetNetworkConfiguration_ARM `json:"networkInterfaceConfigurations,omitempty"`
}

// Deprecated version of VirtualMachineScaleSetOSProfile. Use v1api20201201.VirtualMachineScaleSetOSProfile instead
type VirtualMachineScaleSetOSProfile_ARM struct {
	AdminPassword        *string                   `json:"adminPassword,omitempty"`
	AdminUsername        *string                   `json:"adminUsername,omitempty"`
	ComputerNamePrefix   *string                   `json:"computerNamePrefix,omitempty"`
	CustomData           *string                   `json:"customData,omitempty"`
	LinuxConfiguration   *LinuxConfiguration_ARM   `json:"linuxConfiguration,omitempty"`
	Secrets              []VaultSecretGroup_ARM    `json:"secrets,omitempty"`
	WindowsConfiguration *WindowsConfiguration_ARM `json:"windowsConfiguration,omitempty"`
}

// Deprecated version of VirtualMachineScaleSetStorageProfile. Use v1api20201201.VirtualMachineScaleSetStorageProfile instead
type VirtualMachineScaleSetStorageProfile_ARM struct {
	DataDisks      []VirtualMachineScaleSetDataDisk_ARM `json:"dataDisks,omitempty"`
	ImageReference *ImageReference_ARM                  `json:"imageReference,omitempty"`
	OsDisk         *VirtualMachineScaleSetOSDisk_ARM    `json:"osDisk,omitempty"`
}

// Deprecated version of ApiEntityReference. Use v1api20201201.ApiEntityReference instead
type ApiEntityReference_ARM struct {
	Id *string `json:"id,omitempty"`
}

// Deprecated version of TerminateNotificationProfile. Use v1api20201201.TerminateNotificationProfile instead
type TerminateNotificationProfile_ARM struct {
	Enable           *bool   `json:"enable,omitempty"`
	NotBeforeTimeout *string `json:"notBeforeTimeout,omitempty"`
}

// Deprecated version of VirtualMachineScaleSetDataDisk. Use v1api20201201.VirtualMachineScaleSetDataDisk instead
type VirtualMachineScaleSetDataDisk_ARM struct {
	Caching                 *Caching                                         `json:"caching,omitempty"`
	CreateOption            *CreateOption                                    `json:"createOption,omitempty"`
	DiskIOPSReadWrite       *int                                             `json:"diskIOPSReadWrite,omitempty"`
	DiskMBpsReadWrite       *int                                             `json:"diskMBpsReadWrite,omitempty"`
	DiskSizeGB              *int                                             `json:"diskSizeGB,omitempty"`
	Lun                     *int                                             `json:"lun,omitempty"`
	ManagedDisk             *VirtualMachineScaleSetManagedDiskParameters_ARM `json:"managedDisk,omitempty"`
	Name                    *string                                          `json:"name,omitempty"`
	WriteAcceleratorEnabled *bool                                            `json:"writeAcceleratorEnabled,omitempty"`
}

// Deprecated version of VirtualMachineScaleSetExtension. Use v1api20201201.VirtualMachineScaleSetExtension instead
type VirtualMachineScaleSetExtension_ARM struct {
	Name       *string                                        `json:"name,omitempty"`
	Properties *VirtualMachineScaleSetExtensionProperties_ARM `json:"properties,omitempty"`
}

// Deprecated version of VirtualMachineScaleSetNetworkConfiguration. Use v1api20201201.VirtualMachineScaleSetNetworkConfiguration instead
type VirtualMachineScaleSetNetworkConfiguration_ARM struct {
	Id         *string                                                   `json:"id,omitempty"`
	Name       *string                                                   `json:"name,omitempty"`
	Properties *VirtualMachineScaleSetNetworkConfigurationProperties_ARM `json:"properties,omitempty"`
}

// Deprecated version of VirtualMachineScaleSetOSDisk. Use v1api20201201.VirtualMachineScaleSetOSDisk instead
type VirtualMachineScaleSetOSDisk_ARM struct {
	Caching                 *Caching                                         `json:"caching,omitempty"`
	CreateOption            *CreateOption                                    `json:"createOption,omitempty"`
	DiffDiskSettings        *DiffDiskSettings_ARM                            `json:"diffDiskSettings,omitempty"`
	DiskSizeGB              *int                                             `json:"diskSizeGB,omitempty"`
	Image                   *VirtualHardDisk_ARM                             `json:"image,omitempty"`
	ManagedDisk             *VirtualMachineScaleSetManagedDiskParameters_ARM `json:"managedDisk,omitempty"`
	Name                    *string                                          `json:"name,omitempty"`
	OsType                  *VirtualMachineScaleSetOSDisk_OsType             `json:"osType,omitempty"`
	VhdContainers           []string                                         `json:"vhdContainers,omitempty"`
	WriteAcceleratorEnabled *bool                                            `json:"writeAcceleratorEnabled,omitempty"`
}

// Deprecated version of VirtualMachineScaleSetExtensionProperties. Use v1api20201201.VirtualMachineScaleSetExtensionProperties instead
type VirtualMachineScaleSetExtensionProperties_ARM struct {
	AutoUpgradeMinorVersion  *bool              `json:"autoUpgradeMinorVersion,omitempty"`
	EnableAutomaticUpgrade   *bool              `json:"enableAutomaticUpgrade,omitempty"`
	ForceUpdateTag           *string            `json:"forceUpdateTag,omitempty"`
	ProtectedSettings        map[string]v1.JSON `json:"protectedSettings,omitempty"`
	ProvisionAfterExtensions []string           `json:"provisionAfterExtensions,omitempty"`
	Publisher                *string            `json:"publisher,omitempty"`
	Settings                 map[string]v1.JSON `json:"settings,omitempty"`
	Type                     *string            `json:"type,omitempty"`
	TypeHandlerVersion       *string            `json:"typeHandlerVersion,omitempty"`
}

// Deprecated version of VirtualMachineScaleSetManagedDiskParameters. Use v1api20201201.VirtualMachineScaleSetManagedDiskParameters instead
type VirtualMachineScaleSetManagedDiskParameters_ARM struct {
	DiskEncryptionSet  *SubResource_ARM    `json:"diskEncryptionSet,omitempty"`
	StorageAccountType *StorageAccountType `json:"storageAccountType,omitempty"`
}

// Deprecated version of VirtualMachineScaleSetNetworkConfigurationProperties. Use v1api20201201.VirtualMachineScaleSetNetworkConfigurationProperties instead
type VirtualMachineScaleSetNetworkConfigurationProperties_ARM struct {
	DnsSettings                 *VirtualMachineScaleSetNetworkConfigurationDnsSettings_ARM `json:"dnsSettings,omitempty"`
	EnableAcceleratedNetworking *bool                                                      `json:"enableAcceleratedNetworking,omitempty"`
	EnableFpga                  *bool                                                      `json:"enableFpga,omitempty"`
	EnableIPForwarding          *bool                                                      `json:"enableIPForwarding,omitempty"`
	IpConfigurations            []VirtualMachineScaleSetIPConfiguration_ARM                `json:"ipConfigurations,omitempty"`
	NetworkSecurityGroup        *SubResource_ARM                                           `json:"networkSecurityGroup,omitempty"`
	Primary                     *bool                                                      `json:"primary,omitempty"`
}

// Deprecated version of VirtualMachineScaleSetIPConfiguration. Use v1api20201201.VirtualMachineScaleSetIPConfiguration instead
type VirtualMachineScaleSetIPConfiguration_ARM struct {
	Id         *string                                              `json:"id,omitempty"`
	Name       *string                                              `json:"name,omitempty"`
	Properties *VirtualMachineScaleSetIPConfigurationProperties_ARM `json:"properties,omitempty"`
}

// Deprecated version of VirtualMachineScaleSetNetworkConfigurationDnsSettings. Use v1api20201201.VirtualMachineScaleSetNetworkConfigurationDnsSettings instead
type VirtualMachineScaleSetNetworkConfigurationDnsSettings_ARM struct {
	DnsServers []string `json:"dnsServers,omitempty"`
}

// Deprecated version of VirtualMachineScaleSetIPConfigurationProperties. Use v1api20201201.VirtualMachineScaleSetIPConfigurationProperties instead
type VirtualMachineScaleSetIPConfigurationProperties_ARM struct {
	ApplicationGatewayBackendAddressPools []SubResource_ARM                                                        `json:"applicationGatewayBackendAddressPools,omitempty"`
	ApplicationSecurityGroups             []SubResource_ARM                                                        `json:"applicationSecurityGroups,omitempty"`
	LoadBalancerBackendAddressPools       []SubResource_ARM                                                        `json:"loadBalancerBackendAddressPools,omitempty"`
	LoadBalancerInboundNatPools           []SubResource_ARM                                                        `json:"loadBalancerInboundNatPools,omitempty"`
	Primary                               *bool                                                                    `json:"primary,omitempty"`
	PrivateIPAddressVersion               *VirtualMachineScaleSetIPConfigurationProperties_PrivateIPAddressVersion `json:"privateIPAddressVersion,omitempty"`
	PublicIPAddressConfiguration          *VirtualMachineScaleSetPublicIPAddressConfiguration_ARM                  `json:"publicIPAddressConfiguration,omitempty"`
	Subnet                                *ApiEntityReference_ARM                                                  `json:"subnet,omitempty"`
}

// Deprecated version of VirtualMachineScaleSetPublicIPAddressConfiguration. Use v1api20201201.VirtualMachineScaleSetPublicIPAddressConfiguration instead
type VirtualMachineScaleSetPublicIPAddressConfiguration_ARM struct {
	Name       *string                                                           `json:"name,omitempty"`
	Properties *VirtualMachineScaleSetPublicIPAddressConfigurationProperties_ARM `json:"properties,omitempty"`
}

// Deprecated version of VirtualMachineScaleSetPublicIPAddressConfigurationProperties. Use v1api20201201.VirtualMachineScaleSetPublicIPAddressConfigurationProperties instead
type VirtualMachineScaleSetPublicIPAddressConfigurationProperties_ARM struct {
	DnsSettings            *VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettings_ARM                   `json:"dnsSettings,omitempty"`
	IdleTimeoutInMinutes   *int                                                                                 `json:"idleTimeoutInMinutes,omitempty"`
	IpTags                 []VirtualMachineScaleSetIpTag_ARM                                                    `json:"ipTags,omitempty"`
	PublicIPAddressVersion *VirtualMachineScaleSetPublicIPAddressConfigurationProperties_PublicIPAddressVersion `json:"publicIPAddressVersion,omitempty"`
	PublicIPPrefix         *SubResource_ARM                                                                     `json:"publicIPPrefix,omitempty"`
}

// Deprecated version of VirtualMachineScaleSetIpTag. Use v1api20201201.VirtualMachineScaleSetIpTag instead
type VirtualMachineScaleSetIpTag_ARM struct {
	IpTagType *string `json:"ipTagType,omitempty"`
	Tag       *string `json:"tag,omitempty"`
}

// Deprecated version of VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettings. Use v1api20201201.VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettings instead
type VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettings_ARM struct {
	DomainNameLabel *string `json:"domainNameLabel,omitempty"`
}
