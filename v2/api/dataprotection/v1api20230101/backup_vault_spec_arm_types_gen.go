// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20230101

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type BackupVault_Spec_ARM struct {
	// Identity: Input Managed Identity Details
	Identity *DppIdentityDetails_ARM `json:"identity,omitempty"`

	// Location: Resource location.
	Location *string `json:"location,omitempty"`
	Name     string  `json:"name,omitempty"`

	// Properties: BackupVaultResource properties
	Properties *BackupVaultSpec_ARM `json:"properties,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &BackupVault_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2023-01-01"
func (vault BackupVault_Spec_ARM) GetAPIVersion() string {
	return "2023-01-01"
}

// GetName returns the Name of the resource
func (vault *BackupVault_Spec_ARM) GetName() string {
	return vault.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DataProtection/backupVaults"
func (vault *BackupVault_Spec_ARM) GetType() string {
	return "Microsoft.DataProtection/backupVaults"
}

// Backup Vault
type BackupVaultSpec_ARM struct {
	// FeatureSettings: Feature Settings
	FeatureSettings *FeatureSettings_ARM `json:"featureSettings,omitempty"`

	// MonitoringSettings: Monitoring Settings
	MonitoringSettings *MonitoringSettings_ARM `json:"monitoringSettings,omitempty"`

	// SecuritySettings: Security Settings
	SecuritySettings *SecuritySettings_ARM `json:"securitySettings,omitempty"`

	// StorageSettings: Storage Settings
	StorageSettings []StorageSetting_ARM `json:"storageSettings,omitempty"`
}

// Identity details
type DppIdentityDetails_ARM struct {
	// Type: The identityType which can be either SystemAssigned or None
	Type *string `json:"type,omitempty"`
}

// Class containing feature settings of vault
type FeatureSettings_ARM struct {
	// CrossSubscriptionRestoreSettings: CrossSubscriptionRestore Settings
	CrossSubscriptionRestoreSettings *CrossSubscriptionRestoreSettings_ARM `json:"crossSubscriptionRestoreSettings,omitempty"`
}

// Monitoring Settings
type MonitoringSettings_ARM struct {
	// AzureMonitorAlertSettings: Settings for Azure Monitor based alerts
	AzureMonitorAlertSettings *AzureMonitorAlertSettings_ARM `json:"azureMonitorAlertSettings,omitempty"`
}

// Class containing security settings of vault
type SecuritySettings_ARM struct {
	// ImmutabilitySettings: Immutability Settings at vault level
	ImmutabilitySettings *ImmutabilitySettings_ARM `json:"immutabilitySettings,omitempty"`

	// SoftDeleteSettings: Soft delete related settings
	SoftDeleteSettings *SoftDeleteSettings_ARM `json:"softDeleteSettings,omitempty"`
}

// Storage setting
type StorageSetting_ARM struct {
	// DatastoreType: Gets or sets the type of the datastore.
	DatastoreType *StorageSetting_DatastoreType_ARM `json:"datastoreType,omitempty"`

	// Type: Gets or sets the type.
	Type *StorageSetting_Type_ARM `json:"type,omitempty"`
}

// Settings for Azure Monitor based alerts
type AzureMonitorAlertSettings_ARM struct {
	AlertsForAllJobFailures *AzureMonitorAlertSettings_AlertsForAllJobFailures_ARM `json:"alertsForAllJobFailures,omitempty"`
}

// CrossSubscriptionRestore Settings
type CrossSubscriptionRestoreSettings_ARM struct {
	// State: CrossSubscriptionRestore state
	State *CrossSubscriptionRestoreSettings_State_ARM `json:"state,omitempty"`
}

// Immutability Settings at vault level
type ImmutabilitySettings_ARM struct {
	// State: Immutability state
	State *ImmutabilitySettings_State_ARM `json:"state,omitempty"`
}

// Soft delete related settings
type SoftDeleteSettings_ARM struct {
	// RetentionDurationInDays: Soft delete retention duration
	RetentionDurationInDays *float64 `json:"retentionDurationInDays,omitempty"`

	// State: State of soft delete
	State *SoftDeleteSettings_State_ARM `json:"state,omitempty"`
}

// +kubebuilder:validation:Enum={"ArchiveStore","OperationalStore","VaultStore"}
type StorageSetting_DatastoreType_ARM string

const (
	StorageSetting_DatastoreType_ARM_ArchiveStore     = StorageSetting_DatastoreType_ARM("ArchiveStore")
	StorageSetting_DatastoreType_ARM_OperationalStore = StorageSetting_DatastoreType_ARM("OperationalStore")
	StorageSetting_DatastoreType_ARM_VaultStore       = StorageSetting_DatastoreType_ARM("VaultStore")
)

// Mapping from string to StorageSetting_DatastoreType_ARM
var storageSetting_DatastoreType_ARM_Values = map[string]StorageSetting_DatastoreType_ARM{
	"archivestore":     StorageSetting_DatastoreType_ARM_ArchiveStore,
	"operationalstore": StorageSetting_DatastoreType_ARM_OperationalStore,
	"vaultstore":       StorageSetting_DatastoreType_ARM_VaultStore,
}

// +kubebuilder:validation:Enum={"GeoRedundant","LocallyRedundant","ZoneRedundant"}
type StorageSetting_Type_ARM string

const (
	StorageSetting_Type_ARM_GeoRedundant     = StorageSetting_Type_ARM("GeoRedundant")
	StorageSetting_Type_ARM_LocallyRedundant = StorageSetting_Type_ARM("LocallyRedundant")
	StorageSetting_Type_ARM_ZoneRedundant    = StorageSetting_Type_ARM("ZoneRedundant")
)

// Mapping from string to StorageSetting_Type_ARM
var storageSetting_Type_ARM_Values = map[string]StorageSetting_Type_ARM{
	"georedundant":     StorageSetting_Type_ARM_GeoRedundant,
	"locallyredundant": StorageSetting_Type_ARM_LocallyRedundant,
	"zoneredundant":    StorageSetting_Type_ARM_ZoneRedundant,
}

// +kubebuilder:validation:Enum={"Disabled","Enabled"}
type AzureMonitorAlertSettings_AlertsForAllJobFailures_ARM string

const (
	AzureMonitorAlertSettings_AlertsForAllJobFailures_ARM_Disabled = AzureMonitorAlertSettings_AlertsForAllJobFailures_ARM("Disabled")
	AzureMonitorAlertSettings_AlertsForAllJobFailures_ARM_Enabled  = AzureMonitorAlertSettings_AlertsForAllJobFailures_ARM("Enabled")
)

// Mapping from string to AzureMonitorAlertSettings_AlertsForAllJobFailures_ARM
var azureMonitorAlertSettings_AlertsForAllJobFailures_ARM_Values = map[string]AzureMonitorAlertSettings_AlertsForAllJobFailures_ARM{
	"disabled": AzureMonitorAlertSettings_AlertsForAllJobFailures_ARM_Disabled,
	"enabled":  AzureMonitorAlertSettings_AlertsForAllJobFailures_ARM_Enabled,
}

// +kubebuilder:validation:Enum={"Disabled","Enabled","PermanentlyDisabled"}
type CrossSubscriptionRestoreSettings_State_ARM string

const (
	CrossSubscriptionRestoreSettings_State_ARM_Disabled            = CrossSubscriptionRestoreSettings_State_ARM("Disabled")
	CrossSubscriptionRestoreSettings_State_ARM_Enabled             = CrossSubscriptionRestoreSettings_State_ARM("Enabled")
	CrossSubscriptionRestoreSettings_State_ARM_PermanentlyDisabled = CrossSubscriptionRestoreSettings_State_ARM("PermanentlyDisabled")
)

// Mapping from string to CrossSubscriptionRestoreSettings_State_ARM
var crossSubscriptionRestoreSettings_State_ARM_Values = map[string]CrossSubscriptionRestoreSettings_State_ARM{
	"disabled":            CrossSubscriptionRestoreSettings_State_ARM_Disabled,
	"enabled":             CrossSubscriptionRestoreSettings_State_ARM_Enabled,
	"permanentlydisabled": CrossSubscriptionRestoreSettings_State_ARM_PermanentlyDisabled,
}

// +kubebuilder:validation:Enum={"Disabled","Locked","Unlocked"}
type ImmutabilitySettings_State_ARM string

const (
	ImmutabilitySettings_State_ARM_Disabled = ImmutabilitySettings_State_ARM("Disabled")
	ImmutabilitySettings_State_ARM_Locked   = ImmutabilitySettings_State_ARM("Locked")
	ImmutabilitySettings_State_ARM_Unlocked = ImmutabilitySettings_State_ARM("Unlocked")
)

// Mapping from string to ImmutabilitySettings_State_ARM
var immutabilitySettings_State_ARM_Values = map[string]ImmutabilitySettings_State_ARM{
	"disabled": ImmutabilitySettings_State_ARM_Disabled,
	"locked":   ImmutabilitySettings_State_ARM_Locked,
	"unlocked": ImmutabilitySettings_State_ARM_Unlocked,
}

// +kubebuilder:validation:Enum={"AlwaysOn","Off","On"}
type SoftDeleteSettings_State_ARM string

const (
	SoftDeleteSettings_State_ARM_AlwaysOn = SoftDeleteSettings_State_ARM("AlwaysOn")
	SoftDeleteSettings_State_ARM_Off      = SoftDeleteSettings_State_ARM("Off")
	SoftDeleteSettings_State_ARM_On       = SoftDeleteSettings_State_ARM("On")
)

// Mapping from string to SoftDeleteSettings_State_ARM
var softDeleteSettings_State_ARM_Values = map[string]SoftDeleteSettings_State_ARM{
	"alwayson": SoftDeleteSettings_State_ARM_AlwaysOn,
	"off":      SoftDeleteSettings_State_ARM_Off,
	"on":       SoftDeleteSettings_State_ARM_On,
}
