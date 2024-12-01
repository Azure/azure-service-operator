// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type BackupVault_Spec struct {
	// Identity: Input Managed Identity Details
	Identity *DppIdentityDetails `json:"identity,omitempty"`

	// Location: Resource location.
	Location *string `json:"location,omitempty"`
	Name     string  `json:"name,omitempty"`

	// Properties: BackupVaultResource properties
	Properties *BackupVaultSpec `json:"properties,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &BackupVault_Spec{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2023-01-01"
func (vault BackupVault_Spec) GetAPIVersion() string {
	return "2023-01-01"
}

// GetName returns the Name of the resource
func (vault *BackupVault_Spec) GetName() string {
	return vault.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DataProtection/backupVaults"
func (vault *BackupVault_Spec) GetType() string {
	return "Microsoft.DataProtection/backupVaults"
}

// Backup Vault
type BackupVaultSpec struct {
	// FeatureSettings: Feature Settings
	FeatureSettings *FeatureSettings `json:"featureSettings,omitempty"`

	// MonitoringSettings: Monitoring Settings
	MonitoringSettings *MonitoringSettings `json:"monitoringSettings,omitempty"`

	// SecuritySettings: Security Settings
	SecuritySettings *SecuritySettings `json:"securitySettings,omitempty"`

	// StorageSettings: Storage Settings
	StorageSettings []StorageSetting `json:"storageSettings,omitempty"`
}

// Identity details
type DppIdentityDetails struct {
	// Type: The identityType which can be either SystemAssigned or None
	Type *string `json:"type,omitempty"`
}

// Class containing feature settings of vault
type FeatureSettings struct {
	// CrossSubscriptionRestoreSettings: CrossSubscriptionRestore Settings
	CrossSubscriptionRestoreSettings *CrossSubscriptionRestoreSettings `json:"crossSubscriptionRestoreSettings,omitempty"`
}

// Monitoring Settings
type MonitoringSettings struct {
	// AzureMonitorAlertSettings: Settings for Azure Monitor based alerts
	AzureMonitorAlertSettings *AzureMonitorAlertSettings `json:"azureMonitorAlertSettings,omitempty"`
}

// Class containing security settings of vault
type SecuritySettings struct {
	// ImmutabilitySettings: Immutability Settings at vault level
	ImmutabilitySettings *ImmutabilitySettings `json:"immutabilitySettings,omitempty"`

	// SoftDeleteSettings: Soft delete related settings
	SoftDeleteSettings *SoftDeleteSettings `json:"softDeleteSettings,omitempty"`
}

// Storage setting
type StorageSetting struct {
	// DatastoreType: Gets or sets the type of the datastore.
	DatastoreType *StorageSetting_DatastoreType `json:"datastoreType,omitempty"`

	// Type: Gets or sets the type.
	Type *StorageSetting_Type `json:"type,omitempty"`
}

// Settings for Azure Monitor based alerts
type AzureMonitorAlertSettings struct {
	AlertsForAllJobFailures *AzureMonitorAlertSettings_AlertsForAllJobFailures `json:"alertsForAllJobFailures,omitempty"`
}

// CrossSubscriptionRestore Settings
type CrossSubscriptionRestoreSettings struct {
	// State: CrossSubscriptionRestore state
	State *CrossSubscriptionRestoreSettings_State `json:"state,omitempty"`
}

// Immutability Settings at vault level
type ImmutabilitySettings struct {
	// State: Immutability state
	State *ImmutabilitySettings_State `json:"state,omitempty"`
}

// Soft delete related settings
type SoftDeleteSettings struct {
	// RetentionDurationInDays: Soft delete retention duration
	RetentionDurationInDays *float64 `json:"retentionDurationInDays,omitempty"`

	// State: State of soft delete
	State *SoftDeleteSettings_State `json:"state,omitempty"`
}

// +kubebuilder:validation:Enum={"ArchiveStore","OperationalStore","VaultStore"}
type StorageSetting_DatastoreType string

const (
	StorageSetting_DatastoreType_ArchiveStore     = StorageSetting_DatastoreType("ArchiveStore")
	StorageSetting_DatastoreType_OperationalStore = StorageSetting_DatastoreType("OperationalStore")
	StorageSetting_DatastoreType_VaultStore       = StorageSetting_DatastoreType("VaultStore")
)

// Mapping from string to StorageSetting_DatastoreType
var storageSetting_DatastoreType_Values = map[string]StorageSetting_DatastoreType{
	"archivestore":     StorageSetting_DatastoreType_ArchiveStore,
	"operationalstore": StorageSetting_DatastoreType_OperationalStore,
	"vaultstore":       StorageSetting_DatastoreType_VaultStore,
}

// +kubebuilder:validation:Enum={"GeoRedundant","LocallyRedundant","ZoneRedundant"}
type StorageSetting_Type string

const (
	StorageSetting_Type_GeoRedundant     = StorageSetting_Type("GeoRedundant")
	StorageSetting_Type_LocallyRedundant = StorageSetting_Type("LocallyRedundant")
	StorageSetting_Type_ZoneRedundant    = StorageSetting_Type("ZoneRedundant")
)

// Mapping from string to StorageSetting_Type
var storageSetting_Type_Values = map[string]StorageSetting_Type{
	"georedundant":     StorageSetting_Type_GeoRedundant,
	"locallyredundant": StorageSetting_Type_LocallyRedundant,
	"zoneredundant":    StorageSetting_Type_ZoneRedundant,
}

// +kubebuilder:validation:Enum={"Disabled","Enabled"}
type AzureMonitorAlertSettings_AlertsForAllJobFailures string

const (
	AzureMonitorAlertSettings_AlertsForAllJobFailures_Disabled = AzureMonitorAlertSettings_AlertsForAllJobFailures("Disabled")
	AzureMonitorAlertSettings_AlertsForAllJobFailures_Enabled  = AzureMonitorAlertSettings_AlertsForAllJobFailures("Enabled")
)

// Mapping from string to AzureMonitorAlertSettings_AlertsForAllJobFailures
var azureMonitorAlertSettings_AlertsForAllJobFailures_Values = map[string]AzureMonitorAlertSettings_AlertsForAllJobFailures{
	"disabled": AzureMonitorAlertSettings_AlertsForAllJobFailures_Disabled,
	"enabled":  AzureMonitorAlertSettings_AlertsForAllJobFailures_Enabled,
}

// +kubebuilder:validation:Enum={"Disabled","Enabled","PermanentlyDisabled"}
type CrossSubscriptionRestoreSettings_State string

const (
	CrossSubscriptionRestoreSettings_State_Disabled            = CrossSubscriptionRestoreSettings_State("Disabled")
	CrossSubscriptionRestoreSettings_State_Enabled             = CrossSubscriptionRestoreSettings_State("Enabled")
	CrossSubscriptionRestoreSettings_State_PermanentlyDisabled = CrossSubscriptionRestoreSettings_State("PermanentlyDisabled")
)

// Mapping from string to CrossSubscriptionRestoreSettings_State
var crossSubscriptionRestoreSettings_State_Values = map[string]CrossSubscriptionRestoreSettings_State{
	"disabled":            CrossSubscriptionRestoreSettings_State_Disabled,
	"enabled":             CrossSubscriptionRestoreSettings_State_Enabled,
	"permanentlydisabled": CrossSubscriptionRestoreSettings_State_PermanentlyDisabled,
}

// +kubebuilder:validation:Enum={"Disabled","Locked","Unlocked"}
type ImmutabilitySettings_State string

const (
	ImmutabilitySettings_State_Disabled = ImmutabilitySettings_State("Disabled")
	ImmutabilitySettings_State_Locked   = ImmutabilitySettings_State("Locked")
	ImmutabilitySettings_State_Unlocked = ImmutabilitySettings_State("Unlocked")
)

// Mapping from string to ImmutabilitySettings_State
var immutabilitySettings_State_Values = map[string]ImmutabilitySettings_State{
	"disabled": ImmutabilitySettings_State_Disabled,
	"locked":   ImmutabilitySettings_State_Locked,
	"unlocked": ImmutabilitySettings_State_Unlocked,
}

// +kubebuilder:validation:Enum={"AlwaysOn","Off","On"}
type SoftDeleteSettings_State string

const (
	SoftDeleteSettings_State_AlwaysOn = SoftDeleteSettings_State("AlwaysOn")
	SoftDeleteSettings_State_Off      = SoftDeleteSettings_State("Off")
	SoftDeleteSettings_State_On       = SoftDeleteSettings_State("On")
)

// Mapping from string to SoftDeleteSettings_State
var softDeleteSettings_State_Values = map[string]SoftDeleteSettings_State{
	"alwayson": SoftDeleteSettings_State_AlwaysOn,
	"off":      SoftDeleteSettings_State_Off,
	"on":       SoftDeleteSettings_State_On,
}
