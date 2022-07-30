// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201201

import "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"

// Deprecated version of VirtualMachine_Status. Use v1beta20201201.VirtualMachine_Status instead
type VirtualMachine_StatusARM struct {
	ExtendedLocation *ExtendedLocation_StatusARM         `json:"extendedLocation,omitempty"`
	Id               *string                             `json:"id,omitempty"`
	Identity         *VirtualMachineIdentity_StatusARM   `json:"identity,omitempty"`
	Location         *string                             `json:"location,omitempty"`
	Name             *string                             `json:"name,omitempty"`
	Plan             *Plan_StatusARM                     `json:"plan,omitempty"`
	Properties       *VirtualMachineProperties_StatusARM `json:"properties,omitempty"`
	Resources        []VirtualMachineExtension_StatusARM `json:"resources,omitempty"`
	Tags             map[string]string                   `json:"tags,omitempty"`
	Type             *string                             `json:"type,omitempty"`
	Zones            []string                            `json:"zones,omitempty"`
}

// Deprecated version of VirtualMachineExtension_Status. Use v1beta20201201.VirtualMachineExtension_Status instead
type VirtualMachineExtension_StatusARM struct {
	Id         *string                                      `json:"id,omitempty"`
	Location   *string                                      `json:"location,omitempty"`
	Name       *string                                      `json:"name,omitempty"`
	Properties *VirtualMachineExtensionProperties_StatusARM `json:"properties,omitempty"`
	Tags       map[string]string                            `json:"tags,omitempty"`
	Type       *string                                      `json:"type,omitempty"`
}

// Deprecated version of VirtualMachineIdentity_Status. Use v1beta20201201.VirtualMachineIdentity_Status instead
type VirtualMachineIdentity_StatusARM struct {
	PrincipalId *string                           `json:"principalId,omitempty"`
	TenantId    *string                           `json:"tenantId,omitempty"`
	Type        *VirtualMachineIdentityStatusType `json:"type,omitempty"`
}

// Deprecated version of VirtualMachineProperties_Status. Use v1beta20201201.VirtualMachineProperties_Status instead
type VirtualMachineProperties_StatusARM struct {
	AdditionalCapabilities  *AdditionalCapabilities_StatusARM     `json:"additionalCapabilities,omitempty"`
	AvailabilitySet         *SubResource_StatusARM                `json:"availabilitySet,omitempty"`
	BillingProfile          *BillingProfile_StatusARM             `json:"billingProfile,omitempty"`
	DiagnosticsProfile      *DiagnosticsProfile_StatusARM         `json:"diagnosticsProfile,omitempty"`
	EvictionPolicy          *EvictionPolicy_Status                `json:"evictionPolicy,omitempty"`
	ExtensionsTimeBudget    *string                               `json:"extensionsTimeBudget,omitempty"`
	HardwareProfile         *HardwareProfile_StatusARM            `json:"hardwareProfile,omitempty"`
	Host                    *SubResource_StatusARM                `json:"host,omitempty"`
	HostGroup               *SubResource_StatusARM                `json:"hostGroup,omitempty"`
	InstanceView            *VirtualMachineInstanceView_StatusARM `json:"instanceView,omitempty"`
	LicenseType             *string                               `json:"licenseType,omitempty"`
	NetworkProfile          *NetworkProfile_StatusARM             `json:"networkProfile,omitempty"`
	OsProfile               *OSProfile_StatusARM                  `json:"osProfile,omitempty"`
	PlatformFaultDomain     *int                                  `json:"platformFaultDomain,omitempty"`
	Priority                *Priority_Status                      `json:"priority,omitempty"`
	ProvisioningState       *string                               `json:"provisioningState,omitempty"`
	ProximityPlacementGroup *SubResource_StatusARM                `json:"proximityPlacementGroup,omitempty"`
	SecurityProfile         *SecurityProfile_StatusARM            `json:"securityProfile,omitempty"`
	StorageProfile          *StorageProfile_StatusARM             `json:"storageProfile,omitempty"`
	VirtualMachineScaleSet  *SubResource_StatusARM                `json:"virtualMachineScaleSet,omitempty"`
	VmId                    *string                               `json:"vmId,omitempty"`
}

// Deprecated version of BillingProfile_Status. Use v1beta20201201.BillingProfile_Status instead
type BillingProfile_StatusARM struct {
	MaxPrice *float64 `json:"maxPrice,omitempty"`
}

// Deprecated version of DiagnosticsProfile_Status. Use v1beta20201201.DiagnosticsProfile_Status instead
type DiagnosticsProfile_StatusARM struct {
	BootDiagnostics *BootDiagnostics_StatusARM `json:"bootDiagnostics,omitempty"`
}

// Deprecated version of HardwareProfile_Status. Use v1beta20201201.HardwareProfile_Status instead
type HardwareProfile_StatusARM struct {
	VmSize *HardwareProfileStatusVmSize `json:"vmSize,omitempty"`
}

// Deprecated version of NetworkProfile_Status. Use v1beta20201201.NetworkProfile_Status instead
type NetworkProfile_StatusARM struct {
	NetworkInterfaces []NetworkInterfaceReference_StatusARM `json:"networkInterfaces,omitempty"`
}

// Deprecated version of OSProfile_Status. Use v1beta20201201.OSProfile_Status instead
type OSProfile_StatusARM struct {
	AdminUsername               *string                         `json:"adminUsername,omitempty"`
	AllowExtensionOperations    *bool                           `json:"allowExtensionOperations,omitempty"`
	ComputerName                *string                         `json:"computerName,omitempty"`
	CustomData                  *string                         `json:"customData,omitempty"`
	LinuxConfiguration          *LinuxConfiguration_StatusARM   `json:"linuxConfiguration,omitempty"`
	RequireGuestProvisionSignal *bool                           `json:"requireGuestProvisionSignal,omitempty"`
	Secrets                     []VaultSecretGroup_StatusARM    `json:"secrets,omitempty"`
	WindowsConfiguration        *WindowsConfiguration_StatusARM `json:"windowsConfiguration,omitempty"`
}

// Deprecated version of SecurityProfile_Status. Use v1beta20201201.SecurityProfile_Status instead
type SecurityProfile_StatusARM struct {
	EncryptionAtHost *bool                              `json:"encryptionAtHost,omitempty"`
	SecurityType     *SecurityProfileStatusSecurityType `json:"securityType,omitempty"`
	UefiSettings     *UefiSettings_StatusARM            `json:"uefiSettings,omitempty"`
}

// Deprecated version of StorageProfile_Status. Use v1beta20201201.StorageProfile_Status instead
type StorageProfile_StatusARM struct {
	DataDisks      []DataDisk_StatusARM      `json:"dataDisks,omitempty"`
	ImageReference *ImageReference_StatusARM `json:"imageReference,omitempty"`
	OsDisk         *OSDisk_StatusARM         `json:"osDisk,omitempty"`
}

// Deprecated version of VirtualMachineExtensionProperties_Status. Use v1beta20201201.VirtualMachineExtensionProperties_Status instead
type VirtualMachineExtensionProperties_StatusARM struct {
	AutoUpgradeMinorVersion *bool                                          `json:"autoUpgradeMinorVersion,omitempty"`
	EnableAutomaticUpgrade  *bool                                          `json:"enableAutomaticUpgrade,omitempty"`
	ForceUpdateTag          *string                                        `json:"forceUpdateTag,omitempty"`
	InstanceView            *VirtualMachineExtensionInstanceView_StatusARM `json:"instanceView,omitempty"`
	ProtectedSettings       map[string]v1.JSON                             `json:"protectedSettings,omitempty"`
	ProvisioningState       *string                                        `json:"provisioningState,omitempty"`
	Publisher               *string                                        `json:"publisher,omitempty"`
	Settings                map[string]v1.JSON                             `json:"settings,omitempty"`
	Type                    *string                                        `json:"type,omitempty"`
	TypeHandlerVersion      *string                                        `json:"typeHandlerVersion,omitempty"`
}

// Deprecated version of VirtualMachineIdentityStatusType. Use v1beta20201201.VirtualMachineIdentityStatusType instead
type VirtualMachineIdentityStatusType string

const (
	VirtualMachineIdentityStatusTypeNone                       = VirtualMachineIdentityStatusType("None")
	VirtualMachineIdentityStatusTypeSystemAssigned             = VirtualMachineIdentityStatusType("SystemAssigned")
	VirtualMachineIdentityStatusTypeSystemAssignedUserAssigned = VirtualMachineIdentityStatusType("SystemAssigned, UserAssigned")
	VirtualMachineIdentityStatusTypeUserAssigned               = VirtualMachineIdentityStatusType("UserAssigned")
)

// Deprecated version of VirtualMachineInstanceView_Status. Use v1beta20201201.VirtualMachineInstanceView_Status instead
type VirtualMachineInstanceView_StatusARM struct {
	AssignedHost              *string                                           `json:"assignedHost,omitempty"`
	BootDiagnostics           *BootDiagnosticsInstanceView_StatusARM            `json:"bootDiagnostics,omitempty"`
	ComputerName              *string                                           `json:"computerName,omitempty"`
	Disks                     []DiskInstanceView_StatusARM                      `json:"disks,omitempty"`
	Extensions                []VirtualMachineExtensionInstanceView_StatusARM   `json:"extensions,omitempty"`
	HyperVGeneration          *VirtualMachineInstanceViewStatusHyperVGeneration `json:"hyperVGeneration,omitempty"`
	MaintenanceRedeployStatus *MaintenanceRedeployStatus_StatusARM              `json:"maintenanceRedeployStatus,omitempty"`
	OsName                    *string                                           `json:"osName,omitempty"`
	OsVersion                 *string                                           `json:"osVersion,omitempty"`
	PatchStatus               *VirtualMachinePatchStatus_StatusARM              `json:"patchStatus,omitempty"`
	PlatformFaultDomain       *int                                              `json:"platformFaultDomain,omitempty"`
	PlatformUpdateDomain      *int                                              `json:"platformUpdateDomain,omitempty"`
	RdpThumbPrint             *string                                           `json:"rdpThumbPrint,omitempty"`
	Statuses                  []InstanceViewStatus_StatusARM                    `json:"statuses,omitempty"`
	VmAgent                   *VirtualMachineAgentInstanceView_StatusARM        `json:"vmAgent,omitempty"`
	VmHealth                  *VirtualMachineHealthStatus_StatusARM             `json:"vmHealth,omitempty"`
}

// Deprecated version of BootDiagnosticsInstanceView_Status. Use v1beta20201201.BootDiagnosticsInstanceView_Status instead
type BootDiagnosticsInstanceView_StatusARM struct {
	ConsoleScreenshotBlobUri *string                       `json:"consoleScreenshotBlobUri,omitempty"`
	SerialConsoleLogBlobUri  *string                       `json:"serialConsoleLogBlobUri,omitempty"`
	Status                   *InstanceViewStatus_StatusARM `json:"status,omitempty"`
}

// Deprecated version of BootDiagnostics_Status. Use v1beta20201201.BootDiagnostics_Status instead
type BootDiagnostics_StatusARM struct {
	Enabled    *bool   `json:"enabled,omitempty"`
	StorageUri *string `json:"storageUri,omitempty"`
}

// Deprecated version of DataDisk_Status. Use v1beta20201201.DataDisk_Status instead
type DataDisk_StatusARM struct {
	Caching                 *Caching_Status                  `json:"caching,omitempty"`
	CreateOption            *CreateOption_Status             `json:"createOption,omitempty"`
	DetachOption            *DetachOption_Status             `json:"detachOption,omitempty"`
	DiskIOPSReadWrite       *int                             `json:"diskIOPSReadWrite,omitempty"`
	DiskMBpsReadWrite       *int                             `json:"diskMBpsReadWrite,omitempty"`
	DiskSizeGB              *int                             `json:"diskSizeGB,omitempty"`
	Image                   *VirtualHardDisk_StatusARM       `json:"image,omitempty"`
	Lun                     *int                             `json:"lun,omitempty"`
	ManagedDisk             *ManagedDiskParameters_StatusARM `json:"managedDisk,omitempty"`
	Name                    *string                          `json:"name,omitempty"`
	ToBeDetached            *bool                            `json:"toBeDetached,omitempty"`
	Vhd                     *VirtualHardDisk_StatusARM       `json:"vhd,omitempty"`
	WriteAcceleratorEnabled *bool                            `json:"writeAcceleratorEnabled,omitempty"`
}

// Deprecated version of DiskInstanceView_Status. Use v1beta20201201.DiskInstanceView_Status instead
type DiskInstanceView_StatusARM struct {
	EncryptionSettings []DiskEncryptionSettings_StatusARM `json:"encryptionSettings,omitempty"`
	Name               *string                            `json:"name,omitempty"`
	Statuses           []InstanceViewStatus_StatusARM     `json:"statuses,omitempty"`
}

// Deprecated version of ImageReference_Status. Use v1beta20201201.ImageReference_Status instead
type ImageReference_StatusARM struct {
	ExactVersion *string `json:"exactVersion,omitempty"`
	Id           *string `json:"id,omitempty"`
	Offer        *string `json:"offer,omitempty"`
	Publisher    *string `json:"publisher,omitempty"`
	Sku          *string `json:"sku,omitempty"`
	Version      *string `json:"version,omitempty"`
}

// Deprecated version of InstanceViewStatus_Status. Use v1beta20201201.InstanceViewStatus_Status instead
type InstanceViewStatus_StatusARM struct {
	Code          *string                        `json:"code,omitempty"`
	DisplayStatus *string                        `json:"displayStatus,omitempty"`
	Level         *InstanceViewStatusStatusLevel `json:"level,omitempty"`
	Message       *string                        `json:"message,omitempty"`
	Time          *string                        `json:"time,omitempty"`
}

// Deprecated version of LinuxConfiguration_Status. Use v1beta20201201.LinuxConfiguration_Status instead
type LinuxConfiguration_StatusARM struct {
	DisablePasswordAuthentication *bool                         `json:"disablePasswordAuthentication,omitempty"`
	PatchSettings                 *LinuxPatchSettings_StatusARM `json:"patchSettings,omitempty"`
	ProvisionVMAgent              *bool                         `json:"provisionVMAgent,omitempty"`
	Ssh                           *SshConfiguration_StatusARM   `json:"ssh,omitempty"`
}

// Deprecated version of MaintenanceRedeployStatus_Status. Use v1beta20201201.MaintenanceRedeployStatus_Status instead
type MaintenanceRedeployStatus_StatusARM struct {
	IsCustomerInitiatedMaintenanceAllowed *bool                                                   `json:"isCustomerInitiatedMaintenanceAllowed,omitempty"`
	LastOperationMessage                  *string                                                 `json:"lastOperationMessage,omitempty"`
	LastOperationResultCode               *MaintenanceRedeployStatusStatusLastOperationResultCode `json:"lastOperationResultCode,omitempty"`
	MaintenanceWindowEndTime              *string                                                 `json:"maintenanceWindowEndTime,omitempty"`
	MaintenanceWindowStartTime            *string                                                 `json:"maintenanceWindowStartTime,omitempty"`
	PreMaintenanceWindowEndTime           *string                                                 `json:"preMaintenanceWindowEndTime,omitempty"`
	PreMaintenanceWindowStartTime         *string                                                 `json:"preMaintenanceWindowStartTime,omitempty"`
}

// Deprecated version of NetworkInterfaceReference_Status. Use v1beta20201201.NetworkInterfaceReference_Status instead
type NetworkInterfaceReference_StatusARM struct {
	Id         *string                                        `json:"id,omitempty"`
	Properties *NetworkInterfaceReferenceProperties_StatusARM `json:"properties,omitempty"`
}

// Deprecated version of OSDisk_Status. Use v1beta20201201.OSDisk_Status instead
type OSDisk_StatusARM struct {
	Caching                 *Caching_Status                   `json:"caching,omitempty"`
	CreateOption            *CreateOption_Status              `json:"createOption,omitempty"`
	DiffDiskSettings        *DiffDiskSettings_StatusARM       `json:"diffDiskSettings,omitempty"`
	DiskSizeGB              *int                              `json:"diskSizeGB,omitempty"`
	EncryptionSettings      *DiskEncryptionSettings_StatusARM `json:"encryptionSettings,omitempty"`
	Image                   *VirtualHardDisk_StatusARM        `json:"image,omitempty"`
	ManagedDisk             *ManagedDiskParameters_StatusARM  `json:"managedDisk,omitempty"`
	Name                    *string                           `json:"name,omitempty"`
	OsType                  *OSDiskStatusOsType               `json:"osType,omitempty"`
	Vhd                     *VirtualHardDisk_StatusARM        `json:"vhd,omitempty"`
	WriteAcceleratorEnabled *bool                             `json:"writeAcceleratorEnabled,omitempty"`
}

// Deprecated version of UefiSettings_Status. Use v1beta20201201.UefiSettings_Status instead
type UefiSettings_StatusARM struct {
	SecureBootEnabled *bool `json:"secureBootEnabled,omitempty"`
	VTpmEnabled       *bool `json:"vTpmEnabled,omitempty"`
}

// Deprecated version of VaultSecretGroup_Status. Use v1beta20201201.VaultSecretGroup_Status instead
type VaultSecretGroup_StatusARM struct {
	SourceVault       *SubResource_StatusARM       `json:"sourceVault,omitempty"`
	VaultCertificates []VaultCertificate_StatusARM `json:"vaultCertificates,omitempty"`
}

// Deprecated version of VirtualMachineAgentInstanceView_Status. Use v1beta20201201.VirtualMachineAgentInstanceView_Status instead
type VirtualMachineAgentInstanceView_StatusARM struct {
	ExtensionHandlers []VirtualMachineExtensionHandlerInstanceView_StatusARM `json:"extensionHandlers,omitempty"`
	Statuses          []InstanceViewStatus_StatusARM                         `json:"statuses,omitempty"`
	VmAgentVersion    *string                                                `json:"vmAgentVersion,omitempty"`
}

// Deprecated version of VirtualMachineExtensionInstanceView_Status. Use v1beta20201201.VirtualMachineExtensionInstanceView_Status instead
type VirtualMachineExtensionInstanceView_StatusARM struct {
	Name               *string                        `json:"name,omitempty"`
	Statuses           []InstanceViewStatus_StatusARM `json:"statuses,omitempty"`
	Substatuses        []InstanceViewStatus_StatusARM `json:"substatuses,omitempty"`
	Type               *string                        `json:"type,omitempty"`
	TypeHandlerVersion *string                        `json:"typeHandlerVersion,omitempty"`
}

// Deprecated version of VirtualMachineHealthStatus_Status. Use v1beta20201201.VirtualMachineHealthStatus_Status instead
type VirtualMachineHealthStatus_StatusARM struct {
	Status *InstanceViewStatus_StatusARM `json:"status,omitempty"`
}

// Deprecated version of VirtualMachinePatchStatus_Status. Use v1beta20201201.VirtualMachinePatchStatus_Status instead
type VirtualMachinePatchStatus_StatusARM struct {
	AvailablePatchSummary        *AvailablePatchSummary_StatusARM        `json:"availablePatchSummary,omitempty"`
	ConfigurationStatuses        []InstanceViewStatus_StatusARM          `json:"configurationStatuses,omitempty"`
	LastPatchInstallationSummary *LastPatchInstallationSummary_StatusARM `json:"lastPatchInstallationSummary,omitempty"`
}

// Deprecated version of WindowsConfiguration_Status. Use v1beta20201201.WindowsConfiguration_Status instead
type WindowsConfiguration_StatusARM struct {
	AdditionalUnattendContent []AdditionalUnattendContent_StatusARM `json:"additionalUnattendContent,omitempty"`
	EnableAutomaticUpdates    *bool                                 `json:"enableAutomaticUpdates,omitempty"`
	PatchSettings             *PatchSettings_StatusARM              `json:"patchSettings,omitempty"`
	ProvisionVMAgent          *bool                                 `json:"provisionVMAgent,omitempty"`
	TimeZone                  *string                               `json:"timeZone,omitempty"`
	WinRM                     *WinRMConfiguration_StatusARM         `json:"winRM,omitempty"`
}

// Deprecated version of AdditionalUnattendContent_Status. Use v1beta20201201.AdditionalUnattendContent_Status instead
type AdditionalUnattendContent_StatusARM struct {
	ComponentName *AdditionalUnattendContentStatusComponentName `json:"componentName,omitempty"`
	Content       *string                                       `json:"content,omitempty"`
	PassName      *AdditionalUnattendContentStatusPassName      `json:"passName,omitempty"`
	SettingName   *AdditionalUnattendContentStatusSettingName   `json:"settingName,omitempty"`
}

// Deprecated version of AvailablePatchSummary_Status. Use v1beta20201201.AvailablePatchSummary_Status instead
type AvailablePatchSummary_StatusARM struct {
	AssessmentActivityId          *string                            `json:"assessmentActivityId,omitempty"`
	CriticalAndSecurityPatchCount *int                               `json:"criticalAndSecurityPatchCount,omitempty"`
	Error                         *ApiError_StatusARM                `json:"error,omitempty"`
	LastModifiedTime              *string                            `json:"lastModifiedTime,omitempty"`
	OtherPatchCount               *int                               `json:"otherPatchCount,omitempty"`
	RebootPending                 *bool                              `json:"rebootPending,omitempty"`
	StartTime                     *string                            `json:"startTime,omitempty"`
	Status                        *AvailablePatchSummaryStatusStatus `json:"status,omitempty"`
}

// Deprecated version of DiffDiskSettings_Status. Use v1beta20201201.DiffDiskSettings_Status instead
type DiffDiskSettings_StatusARM struct {
	Option    *DiffDiskOption_Status    `json:"option,omitempty"`
	Placement *DiffDiskPlacement_Status `json:"placement,omitempty"`
}

// Deprecated version of DiskEncryptionSettings_Status. Use v1beta20201201.DiskEncryptionSettings_Status instead
type DiskEncryptionSettings_StatusARM struct {
	DiskEncryptionKey *KeyVaultSecretReference_StatusARM `json:"diskEncryptionKey,omitempty"`
	Enabled           *bool                              `json:"enabled,omitempty"`
	KeyEncryptionKey  *KeyVaultKeyReference_StatusARM    `json:"keyEncryptionKey,omitempty"`
}

// Deprecated version of LastPatchInstallationSummary_Status. Use v1beta20201201.LastPatchInstallationSummary_Status instead
type LastPatchInstallationSummary_StatusARM struct {
	Error                     *ApiError_StatusARM                       `json:"error,omitempty"`
	ExcludedPatchCount        *int                                      `json:"excludedPatchCount,omitempty"`
	FailedPatchCount          *int                                      `json:"failedPatchCount,omitempty"`
	InstallationActivityId    *string                                   `json:"installationActivityId,omitempty"`
	InstalledPatchCount       *int                                      `json:"installedPatchCount,omitempty"`
	LastModifiedTime          *string                                   `json:"lastModifiedTime,omitempty"`
	MaintenanceWindowExceeded *bool                                     `json:"maintenanceWindowExceeded,omitempty"`
	NotSelectedPatchCount     *int                                      `json:"notSelectedPatchCount,omitempty"`
	PendingPatchCount         *int                                      `json:"pendingPatchCount,omitempty"`
	StartTime                 *string                                   `json:"startTime,omitempty"`
	Status                    *LastPatchInstallationSummaryStatusStatus `json:"status,omitempty"`
}

// Deprecated version of LinuxPatchSettings_Status. Use v1beta20201201.LinuxPatchSettings_Status instead
type LinuxPatchSettings_StatusARM struct {
	PatchMode *LinuxPatchSettingsStatusPatchMode `json:"patchMode,omitempty"`
}

// Deprecated version of ManagedDiskParameters_Status. Use v1beta20201201.ManagedDiskParameters_Status instead
type ManagedDiskParameters_StatusARM struct {
	DiskEncryptionSet  *SubResource_StatusARM     `json:"diskEncryptionSet,omitempty"`
	Id                 *string                    `json:"id,omitempty"`
	StorageAccountType *StorageAccountType_Status `json:"storageAccountType,omitempty"`
}

// Deprecated version of NetworkInterfaceReferenceProperties_Status. Use v1beta20201201.NetworkInterfaceReferenceProperties_Status instead
type NetworkInterfaceReferenceProperties_StatusARM struct {
	Primary *bool `json:"primary,omitempty"`
}

// Deprecated version of PatchSettings_Status. Use v1beta20201201.PatchSettings_Status instead
type PatchSettings_StatusARM struct {
	EnableHotpatching *bool                         `json:"enableHotpatching,omitempty"`
	PatchMode         *PatchSettingsStatusPatchMode `json:"patchMode,omitempty"`
}

// Deprecated version of SshConfiguration_Status. Use v1beta20201201.SshConfiguration_Status instead
type SshConfiguration_StatusARM struct {
	PublicKeys []SshPublicKey_StatusARM `json:"publicKeys,omitempty"`
}

// Deprecated version of VaultCertificate_Status. Use v1beta20201201.VaultCertificate_Status instead
type VaultCertificate_StatusARM struct {
	CertificateStore *string `json:"certificateStore,omitempty"`
	CertificateUrl   *string `json:"certificateUrl,omitempty"`
}

// Deprecated version of VirtualHardDisk_Status. Use v1beta20201201.VirtualHardDisk_Status instead
type VirtualHardDisk_StatusARM struct {
	Uri *string `json:"uri,omitempty"`
}

// Deprecated version of VirtualMachineExtensionHandlerInstanceView_Status. Use v1beta20201201.VirtualMachineExtensionHandlerInstanceView_Status instead
type VirtualMachineExtensionHandlerInstanceView_StatusARM struct {
	Status             *InstanceViewStatus_StatusARM `json:"status,omitempty"`
	Type               *string                       `json:"type,omitempty"`
	TypeHandlerVersion *string                       `json:"typeHandlerVersion,omitempty"`
}

// Deprecated version of WinRMConfiguration_Status. Use v1beta20201201.WinRMConfiguration_Status instead
type WinRMConfiguration_StatusARM struct {
	Listeners []WinRMListener_StatusARM `json:"listeners,omitempty"`
}

// Deprecated version of ApiError_Status. Use v1beta20201201.ApiError_Status instead
type ApiError_StatusARM struct {
	Code       *string                  `json:"code,omitempty"`
	Details    []ApiErrorBase_StatusARM `json:"details,omitempty"`
	Innererror *InnerError_StatusARM    `json:"innererror,omitempty"`
	Message    *string                  `json:"message,omitempty"`
	Target     *string                  `json:"target,omitempty"`
}

// Deprecated version of KeyVaultKeyReference_Status. Use v1beta20201201.KeyVaultKeyReference_Status instead
type KeyVaultKeyReference_StatusARM struct {
	KeyUrl      *string                `json:"keyUrl,omitempty"`
	SourceVault *SubResource_StatusARM `json:"sourceVault,omitempty"`
}

// Deprecated version of KeyVaultSecretReference_Status. Use v1beta20201201.KeyVaultSecretReference_Status instead
type KeyVaultSecretReference_StatusARM struct {
	SecretUrl   *string                `json:"secretUrl,omitempty"`
	SourceVault *SubResource_StatusARM `json:"sourceVault,omitempty"`
}

// Deprecated version of SshPublicKey_Status. Use v1beta20201201.SshPublicKey_Status instead
type SshPublicKey_StatusARM struct {
	KeyData *string `json:"keyData,omitempty"`
	Path    *string `json:"path,omitempty"`
}

// Deprecated version of WinRMListener_Status. Use v1beta20201201.WinRMListener_Status instead
type WinRMListener_StatusARM struct {
	CertificateUrl *string                      `json:"certificateUrl,omitempty"`
	Protocol       *WinRMListenerStatusProtocol `json:"protocol,omitempty"`
}

// Deprecated version of ApiErrorBase_Status. Use v1beta20201201.ApiErrorBase_Status instead
type ApiErrorBase_StatusARM struct {
	Code    *string `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
	Target  *string `json:"target,omitempty"`
}

// Deprecated version of InnerError_Status. Use v1beta20201201.InnerError_Status instead
type InnerError_StatusARM struct {
	Errordetail   *string `json:"errordetail,omitempty"`
	Exceptiontype *string `json:"exceptiontype,omitempty"`
}
