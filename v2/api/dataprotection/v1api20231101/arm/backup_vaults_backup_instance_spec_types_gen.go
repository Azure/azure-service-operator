// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

import (
	"encoding/json"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type BackupVaultsBackupInstance_Spec struct {
	Name string `json:"name,omitempty"`

	// Properties: BackupInstanceResource properties
	Properties *BackupInstance `json:"properties,omitempty"`

	// Tags: Proxy Resource tags.
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &BackupVaultsBackupInstance_Spec{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2023-11-01"
func (instance BackupVaultsBackupInstance_Spec) GetAPIVersion() string {
	return "2023-11-01"
}

// GetName returns the Name of the resource
func (instance *BackupVaultsBackupInstance_Spec) GetName() string {
	return instance.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DataProtection/backupVaults/backupInstances"
func (instance *BackupVaultsBackupInstance_Spec) GetType() string {
	return "Microsoft.DataProtection/backupVaults/backupInstances"
}

// Backup Instance
type BackupInstance struct {
	// DataSourceInfo: Gets or sets the data source information.
	DataSourceInfo *Datasource `json:"dataSourceInfo,omitempty"`

	// DataSourceSetInfo: Gets or sets the data source set information.
	DataSourceSetInfo *DatasourceSet `json:"dataSourceSetInfo,omitempty"`

	// DatasourceAuthCredentials: Credentials to use to authenticate with data source provider.
	DatasourceAuthCredentials *AuthCredentials `json:"datasourceAuthCredentials,omitempty"`

	// FriendlyName: Gets or sets the Backup Instance friendly name.
	FriendlyName *string `json:"friendlyName,omitempty"`

	// IdentityDetails: Contains information of the Identity Details for the BI.
	// If it is null, default will be considered as System Assigned.
	IdentityDetails *IdentityDetails `json:"identityDetails,omitempty"`
	ObjectType      *string          `json:"objectType,omitempty"`

	// PolicyInfo: Gets or sets the policy information.
	PolicyInfo *PolicyInfo `json:"policyInfo,omitempty"`

	// ValidationType: Specifies the type of validation. In case of DeepValidation, all validations from /validateForBackup API
	// will run again.
	ValidationType *BackupInstance_ValidationType `json:"validationType,omitempty"`
}

type AuthCredentials struct {
	// SecretStoreBasedAuthCredentials: Mutually exclusive with all other properties
	SecretStoreBasedAuthCredentials *SecretStoreBasedAuthCredentials `json:"secretStoreBasedAuthCredentials,omitempty"`
}

// MarshalJSON defers JSON marshaling to the first non-nil property, because AuthCredentials represents a discriminated union (JSON OneOf)
func (credentials AuthCredentials) MarshalJSON() ([]byte, error) {
	if credentials.SecretStoreBasedAuthCredentials != nil {
		return json.Marshal(credentials.SecretStoreBasedAuthCredentials)
	}
	return nil, nil
}

// UnmarshalJSON unmarshals the AuthCredentials
func (credentials *AuthCredentials) UnmarshalJSON(data []byte) error {
	var rawJson map[string]interface{}
	err := json.Unmarshal(data, &rawJson)
	if err != nil {
		return err
	}
	discriminator := rawJson["objectType"]
	if discriminator == "SecretStoreBasedAuthCredentials" {
		credentials.SecretStoreBasedAuthCredentials = &SecretStoreBasedAuthCredentials{}
		return json.Unmarshal(data, credentials.SecretStoreBasedAuthCredentials)
	}

	// No error
	return nil
}

// +kubebuilder:validation:Enum={"DeepValidation","ShallowValidation"}
type BackupInstance_ValidationType string

const (
	BackupInstance_ValidationType_DeepValidation    = BackupInstance_ValidationType("DeepValidation")
	BackupInstance_ValidationType_ShallowValidation = BackupInstance_ValidationType("ShallowValidation")
)

// Mapping from string to BackupInstance_ValidationType
var backupInstance_ValidationType_Values = map[string]BackupInstance_ValidationType{
	"deepvalidation":    BackupInstance_ValidationType_DeepValidation,
	"shallowvalidation": BackupInstance_ValidationType_ShallowValidation,
}

// Datasource to be backed up
type Datasource struct {
	// DatasourceType: DatasourceType of the resource.
	DatasourceType *string `json:"datasourceType,omitempty"`

	// ObjectType: Type of Datasource object, used to initialize the right inherited type
	ObjectType *string `json:"objectType,omitempty"`
	ResourceID *string `json:"resourceID,omitempty"`

	// ResourceLocation: Location of datasource.
	ResourceLocation *string `json:"resourceLocation,omitempty"`

	// ResourceName: Unique identifier of the resource in the context of parent.
	ResourceName *string `json:"resourceName,omitempty"`

	// ResourceProperties: Properties specific to data source
	ResourceProperties *BaseResourceProperties `json:"resourceProperties,omitempty"`

	// ResourceType: Resource Type of Datasource.
	ResourceType *string `json:"resourceType,omitempty"`

	// ResourceUri: Uri of the resource.
	ResourceUri *string `json:"resourceUri,omitempty"`
}

// DatasourceSet details of datasource to be backed up
type DatasourceSet struct {
	// DatasourceType: DatasourceType of the resource.
	DatasourceType *string `json:"datasourceType,omitempty"`

	// ObjectType: Type of Datasource object, used to initialize the right inherited type
	ObjectType *string `json:"objectType,omitempty"`
	ResourceID *string `json:"resourceID,omitempty"`

	// ResourceLocation: Location of datasource.
	ResourceLocation *string `json:"resourceLocation,omitempty"`

	// ResourceName: Unique identifier of the resource in the context of parent.
	ResourceName *string `json:"resourceName,omitempty"`

	// ResourceProperties: Properties specific to data source set
	ResourceProperties *BaseResourceProperties `json:"resourceProperties,omitempty"`

	// ResourceType: Resource Type of Datasource.
	ResourceType *string `json:"resourceType,omitempty"`

	// ResourceUri: Uri of the resource.
	ResourceUri *string `json:"resourceUri,omitempty"`
}

type IdentityDetails struct {
	// UseSystemAssignedIdentity: Specifies if the BI is protected by System Identity.
	UseSystemAssignedIdentity *bool `json:"useSystemAssignedIdentity,omitempty"`

	// UserAssignedIdentityArmUrl: ARM URL for User Assigned Identity.
	UserAssignedIdentityArmUrl *string `json:"userAssignedIdentityArmUrl,omitempty"`
}

// Policy Info in backupInstance
type PolicyInfo struct {
	PolicyId *string `json:"policyId,omitempty"`

	// PolicyParameters: Policy parameters for the backup instance
	PolicyParameters *PolicyParameters `json:"policyParameters,omitempty"`
}

type BaseResourceProperties struct {
	// DefaultResourceProperties: Mutually exclusive with all other properties
	DefaultResourceProperties *DefaultResourceProperties `json:"defaultResourceProperties,omitempty"`
}

// MarshalJSON defers JSON marshaling to the first non-nil property, because BaseResourceProperties represents a discriminated union (JSON OneOf)
func (properties BaseResourceProperties) MarshalJSON() ([]byte, error) {
	if properties.DefaultResourceProperties != nil {
		return json.Marshal(properties.DefaultResourceProperties)
	}
	return nil, nil
}

// UnmarshalJSON unmarshals the BaseResourceProperties
func (properties *BaseResourceProperties) UnmarshalJSON(data []byte) error {
	var rawJson map[string]interface{}
	err := json.Unmarshal(data, &rawJson)
	if err != nil {
		return err
	}
	discriminator := rawJson["objectType"]
	if discriminator == "DefaultResourceProperties" {
		properties.DefaultResourceProperties = &DefaultResourceProperties{}
		return json.Unmarshal(data, properties.DefaultResourceProperties)
	}

	// No error
	return nil
}

// Parameters in Policy
type PolicyParameters struct {
	// BackupDatasourceParametersList: Gets or sets the Backup Data Source Parameters
	BackupDatasourceParametersList []BackupDatasourceParameters `json:"backupDatasourceParametersList,omitempty"`

	// DataStoreParametersList: Gets or sets the DataStore Parameters
	DataStoreParametersList []DataStoreParameters `json:"dataStoreParametersList,omitempty"`
}

type SecretStoreBasedAuthCredentials struct {
	// ObjectType: Type of the specific object - used for deserializing
	ObjectType SecretStoreBasedAuthCredentials_ObjectType `json:"objectType,omitempty"`

	// SecretStoreResource: Secret store resource
	SecretStoreResource *SecretStoreResource `json:"secretStoreResource,omitempty"`
}

type BackupDatasourceParameters struct {
	// Blob: Mutually exclusive with all other properties
	Blob *BlobBackupDatasourceParameters `json:"blobBackupDatasourceParameters,omitempty"`

	// KubernetesCluster: Mutually exclusive with all other properties
	KubernetesCluster *KubernetesClusterBackupDatasourceParameters `json:"kubernetesClusterBackupDatasourceParameters,omitempty"`
}

// MarshalJSON defers JSON marshaling to the first non-nil property, because BackupDatasourceParameters represents a discriminated union (JSON OneOf)
func (parameters BackupDatasourceParameters) MarshalJSON() ([]byte, error) {
	if parameters.Blob != nil {
		return json.Marshal(parameters.Blob)
	}
	if parameters.KubernetesCluster != nil {
		return json.Marshal(parameters.KubernetesCluster)
	}
	return nil, nil
}

// UnmarshalJSON unmarshals the BackupDatasourceParameters
func (parameters *BackupDatasourceParameters) UnmarshalJSON(data []byte) error {
	var rawJson map[string]interface{}
	err := json.Unmarshal(data, &rawJson)
	if err != nil {
		return err
	}
	discriminator := rawJson["objectType"]
	if discriminator == "BlobBackupDatasourceParameters" {
		parameters.Blob = &BlobBackupDatasourceParameters{}
		return json.Unmarshal(data, parameters.Blob)
	}
	if discriminator == "KubernetesClusterBackupDatasourceParameters" {
		parameters.KubernetesCluster = &KubernetesClusterBackupDatasourceParameters{}
		return json.Unmarshal(data, parameters.KubernetesCluster)
	}

	// No error
	return nil
}

type DataStoreParameters struct {
	// AzureOperationalStoreParameters: Mutually exclusive with all other properties
	AzureOperationalStoreParameters *AzureOperationalStoreParameters `json:"azureOperationalStoreParameters,omitempty"`
}

// MarshalJSON defers JSON marshaling to the first non-nil property, because DataStoreParameters represents a discriminated union (JSON OneOf)
func (parameters DataStoreParameters) MarshalJSON() ([]byte, error) {
	if parameters.AzureOperationalStoreParameters != nil {
		return json.Marshal(parameters.AzureOperationalStoreParameters)
	}
	return nil, nil
}

// UnmarshalJSON unmarshals the DataStoreParameters
func (parameters *DataStoreParameters) UnmarshalJSON(data []byte) error {
	var rawJson map[string]interface{}
	err := json.Unmarshal(data, &rawJson)
	if err != nil {
		return err
	}
	discriminator := rawJson["objectType"]
	if discriminator == "AzureOperationalStoreParameters" {
		parameters.AzureOperationalStoreParameters = &AzureOperationalStoreParameters{}
		return json.Unmarshal(data, parameters.AzureOperationalStoreParameters)
	}

	// No error
	return nil
}

type DefaultResourceProperties struct {
	// ObjectType: Type of the specific object - used for deserializing
	ObjectType DefaultResourceProperties_ObjectType `json:"objectType,omitempty"`
}

// +kubebuilder:validation:Enum={"SecretStoreBasedAuthCredentials"}
type SecretStoreBasedAuthCredentials_ObjectType string

const SecretStoreBasedAuthCredentials_ObjectType_SecretStoreBasedAuthCredentials = SecretStoreBasedAuthCredentials_ObjectType("SecretStoreBasedAuthCredentials")

// Mapping from string to SecretStoreBasedAuthCredentials_ObjectType
var secretStoreBasedAuthCredentials_ObjectType_Values = map[string]SecretStoreBasedAuthCredentials_ObjectType{
	"secretstorebasedauthcredentials": SecretStoreBasedAuthCredentials_ObjectType_SecretStoreBasedAuthCredentials,
}

// Class representing a secret store resource.
type SecretStoreResource struct {
	// SecretStoreType: Gets or sets the type of secret store
	SecretStoreType *SecretStoreResource_SecretStoreType `json:"secretStoreType,omitempty"`

	// Uri: Uri to get to the resource
	Uri *string `json:"uri,omitempty"`

	// Value: Gets or sets value stored in secret store resource
	Value *string `json:"value,omitempty"`
}

type AzureOperationalStoreParameters struct {
	// DataStoreType: type of datastore; Operational/Vault/Archive
	DataStoreType *AzureOperationalStoreParameters_DataStoreType `json:"dataStoreType,omitempty"`

	// ObjectType: Type of the specific object - used for deserializing
	ObjectType      AzureOperationalStoreParameters_ObjectType `json:"objectType,omitempty"`
	ResourceGroupId *string                                    `json:"resourceGroupId,omitempty"`
}

type BlobBackupDatasourceParameters struct {
	// ContainersList: List of containers to be backed up during configuration of backup of blobs
	ContainersList []string `json:"containersList,omitempty"`

	// ObjectType: Type of the specific object - used for deserializing
	ObjectType BlobBackupDatasourceParameters_ObjectType `json:"objectType,omitempty"`
}

// +kubebuilder:validation:Enum={"DefaultResourceProperties"}
type DefaultResourceProperties_ObjectType string

const DefaultResourceProperties_ObjectType_DefaultResourceProperties = DefaultResourceProperties_ObjectType("DefaultResourceProperties")

// Mapping from string to DefaultResourceProperties_ObjectType
var defaultResourceProperties_ObjectType_Values = map[string]DefaultResourceProperties_ObjectType{
	"defaultresourceproperties": DefaultResourceProperties_ObjectType_DefaultResourceProperties,
}

type KubernetesClusterBackupDatasourceParameters struct {
	// BackupHookReferences: Gets or sets the backup hook references. This property sets the hook reference to be executed
	// during backup.
	BackupHookReferences []NamespacedNameResource `json:"backupHookReferences,omitempty"`

	// ExcludedNamespaces: Gets or sets the exclude namespaces property. This property sets the namespaces to be excluded
	// during backup.
	ExcludedNamespaces []string `json:"excludedNamespaces,omitempty"`

	// ExcludedResourceTypes: Gets or sets the exclude resource types property. This property sets the resource types to be
	// excluded during backup.
	ExcludedResourceTypes []string `json:"excludedResourceTypes,omitempty"`

	// IncludeClusterScopeResources: Gets or sets the include cluster resources property. This property if enabled will include
	// cluster scope resources during backup.
	IncludeClusterScopeResources *bool `json:"includeClusterScopeResources,omitempty"`

	// IncludedNamespaces: Gets or sets the include namespaces property. This property sets the namespaces to be included
	// during backup.
	IncludedNamespaces []string `json:"includedNamespaces,omitempty"`

	// IncludedResourceTypes: Gets or sets the include resource types property. This property sets the resource types to be
	// included during backup.
	IncludedResourceTypes []string `json:"includedResourceTypes,omitempty"`

	// LabelSelectors: Gets or sets the LabelSelectors property. This property sets the resource with such label selectors to
	// be included during backup.
	LabelSelectors []string `json:"labelSelectors,omitempty"`

	// ObjectType: Type of the specific object - used for deserializing
	ObjectType KubernetesClusterBackupDatasourceParameters_ObjectType `json:"objectType,omitempty"`

	// SnapshotVolumes: Gets or sets the volume snapshot property. This property if enabled will take volume snapshots during
	// backup.
	SnapshotVolumes *bool `json:"snapshotVolumes,omitempty"`
}

// +kubebuilder:validation:Enum={"AzureKeyVault","Invalid"}
type SecretStoreResource_SecretStoreType string

const (
	SecretStoreResource_SecretStoreType_AzureKeyVault = SecretStoreResource_SecretStoreType("AzureKeyVault")
	SecretStoreResource_SecretStoreType_Invalid       = SecretStoreResource_SecretStoreType("Invalid")
)

// Mapping from string to SecretStoreResource_SecretStoreType
var secretStoreResource_SecretStoreType_Values = map[string]SecretStoreResource_SecretStoreType{
	"azurekeyvault": SecretStoreResource_SecretStoreType_AzureKeyVault,
	"invalid":       SecretStoreResource_SecretStoreType_Invalid,
}

// +kubebuilder:validation:Enum={"ArchiveStore","OperationalStore","VaultStore"}
type AzureOperationalStoreParameters_DataStoreType string

const (
	AzureOperationalStoreParameters_DataStoreType_ArchiveStore     = AzureOperationalStoreParameters_DataStoreType("ArchiveStore")
	AzureOperationalStoreParameters_DataStoreType_OperationalStore = AzureOperationalStoreParameters_DataStoreType("OperationalStore")
	AzureOperationalStoreParameters_DataStoreType_VaultStore       = AzureOperationalStoreParameters_DataStoreType("VaultStore")
)

// Mapping from string to AzureOperationalStoreParameters_DataStoreType
var azureOperationalStoreParameters_DataStoreType_Values = map[string]AzureOperationalStoreParameters_DataStoreType{
	"archivestore":     AzureOperationalStoreParameters_DataStoreType_ArchiveStore,
	"operationalstore": AzureOperationalStoreParameters_DataStoreType_OperationalStore,
	"vaultstore":       AzureOperationalStoreParameters_DataStoreType_VaultStore,
}

// +kubebuilder:validation:Enum={"AzureOperationalStoreParameters"}
type AzureOperationalStoreParameters_ObjectType string

const AzureOperationalStoreParameters_ObjectType_AzureOperationalStoreParameters = AzureOperationalStoreParameters_ObjectType("AzureOperationalStoreParameters")

// Mapping from string to AzureOperationalStoreParameters_ObjectType
var azureOperationalStoreParameters_ObjectType_Values = map[string]AzureOperationalStoreParameters_ObjectType{
	"azureoperationalstoreparameters": AzureOperationalStoreParameters_ObjectType_AzureOperationalStoreParameters,
}

// +kubebuilder:validation:Enum={"BlobBackupDatasourceParameters"}
type BlobBackupDatasourceParameters_ObjectType string

const BlobBackupDatasourceParameters_ObjectType_BlobBackupDatasourceParameters = BlobBackupDatasourceParameters_ObjectType("BlobBackupDatasourceParameters")

// Mapping from string to BlobBackupDatasourceParameters_ObjectType
var blobBackupDatasourceParameters_ObjectType_Values = map[string]BlobBackupDatasourceParameters_ObjectType{
	"blobbackupdatasourceparameters": BlobBackupDatasourceParameters_ObjectType_BlobBackupDatasourceParameters,
}

// +kubebuilder:validation:Enum={"KubernetesClusterBackupDatasourceParameters"}
type KubernetesClusterBackupDatasourceParameters_ObjectType string

const KubernetesClusterBackupDatasourceParameters_ObjectType_KubernetesClusterBackupDatasourceParameters = KubernetesClusterBackupDatasourceParameters_ObjectType("KubernetesClusterBackupDatasourceParameters")

// Mapping from string to KubernetesClusterBackupDatasourceParameters_ObjectType
var kubernetesClusterBackupDatasourceParameters_ObjectType_Values = map[string]KubernetesClusterBackupDatasourceParameters_ObjectType{
	"kubernetesclusterbackupdatasourceparameters": KubernetesClusterBackupDatasourceParameters_ObjectType_KubernetesClusterBackupDatasourceParameters,
}

// Class to refer resources which contains namespace and name
type NamespacedNameResource struct {
	// Name: Name of the resource
	Name *string `json:"name,omitempty"`

	// Namespace: Namespace in which the resource exists
	Namespace *string `json:"namespace,omitempty"`
}
