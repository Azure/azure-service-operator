// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20231101

import (
	"encoding/json"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type BackupVaults_BackupInstance_Spec_ARM struct {
	Name string `json:"name,omitempty"`

	// Properties: BackupInstanceResource properties
	Properties *BackupInstance_ARM `json:"properties,omitempty"`

	// Tags: Proxy Resource tags.
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &BackupVaults_BackupInstance_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2023-11-01"
func (instance BackupVaults_BackupInstance_Spec_ARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (instance *BackupVaults_BackupInstance_Spec_ARM) GetName() string {
	return instance.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DataProtection/backupVaults/backupInstances"
func (instance *BackupVaults_BackupInstance_Spec_ARM) GetType() string {
	return "Microsoft.DataProtection/backupVaults/backupInstances"
}

// Backup Instance
type BackupInstance_ARM struct {
	// DataSourceInfo: Gets or sets the data source information.
	DataSourceInfo *Datasource_ARM `json:"dataSourceInfo,omitempty"`

	// DataSourceSetInfo: Gets or sets the data source set information.
	DataSourceSetInfo *DatasourceSet_ARM `json:"dataSourceSetInfo,omitempty"`

	// DatasourceAuthCredentials: Credentials to use to authenticate with data source provider.
	DatasourceAuthCredentials *AuthCredentials_ARM `json:"datasourceAuthCredentials,omitempty"`

	// FriendlyName: Gets or sets the Backup Instance friendly name.
	FriendlyName *string `json:"friendlyName,omitempty"`

	// IdentityDetails: Contains information of the Identity Details for the BI.
	// If it is null, default will be considered as System Assigned.
	IdentityDetails *IdentityDetails_ARM `json:"identityDetails,omitempty"`
	ObjectType      *string              `json:"objectType,omitempty"`

	// PolicyInfo: Gets or sets the policy information.
	PolicyInfo *PolicyInfo_ARM `json:"policyInfo,omitempty"`

	// ValidationType: Specifies the type of validation. In case of DeepValidation, all validations from /validateForBackup API
	// will run again.
	ValidationType *BackupInstance_ValidationType_ARM `json:"validationType,omitempty"`
}

type AuthCredentials_ARM struct {
	// SecretStoreBasedAuthCredentials: Mutually exclusive with all other properties
	SecretStoreBasedAuthCredentials *SecretStoreBasedAuthCredentials_ARM `json:"secretStoreBasedAuthCredentials,omitempty"`
}

// MarshalJSON defers JSON marshaling to the first non-nil property, because AuthCredentials_ARM represents a discriminated union (JSON OneOf)
func (credentials AuthCredentials_ARM) MarshalJSON() ([]byte, error) {
	if credentials.SecretStoreBasedAuthCredentials != nil {
		return json.Marshal(credentials.SecretStoreBasedAuthCredentials)
	}
	return nil, nil
}

// UnmarshalJSON unmarshals the AuthCredentials_ARM
func (credentials *AuthCredentials_ARM) UnmarshalJSON(data []byte) error {
	var rawJson map[string]interface{}
	err := json.Unmarshal(data, &rawJson)
	if err != nil {
		return err
	}
	discriminator := rawJson["objectType"]
	if discriminator == "SecretStoreBasedAuthCredentials" {
		credentials.SecretStoreBasedAuthCredentials = &SecretStoreBasedAuthCredentials_ARM{}
		return json.Unmarshal(data, credentials.SecretStoreBasedAuthCredentials)
	}

	// No error
	return nil
}

// +kubebuilder:validation:Enum={"DeepValidation","ShallowValidation"}
type BackupInstance_ValidationType_ARM string

const (
	BackupInstance_ValidationType_ARM_DeepValidation    = BackupInstance_ValidationType_ARM("DeepValidation")
	BackupInstance_ValidationType_ARM_ShallowValidation = BackupInstance_ValidationType_ARM("ShallowValidation")
)

// Mapping from string to BackupInstance_ValidationType_ARM
var backupInstance_ValidationType_ARM_Values = map[string]BackupInstance_ValidationType_ARM{
	"deepvalidation":    BackupInstance_ValidationType_ARM_DeepValidation,
	"shallowvalidation": BackupInstance_ValidationType_ARM_ShallowValidation,
}

// Datasource to be backed up
type Datasource_ARM struct {
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
	ResourceProperties *BaseResourceProperties_ARM `json:"resourceProperties,omitempty"`

	// ResourceType: Resource Type of Datasource.
	ResourceType *string `json:"resourceType,omitempty"`

	// ResourceUri: Uri of the resource.
	ResourceUri *string `json:"resourceUri,omitempty"`
}

// DatasourceSet details of datasource to be backed up
type DatasourceSet_ARM struct {
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
	ResourceProperties *BaseResourceProperties_ARM `json:"resourceProperties,omitempty"`

	// ResourceType: Resource Type of Datasource.
	ResourceType *string `json:"resourceType,omitempty"`

	// ResourceUri: Uri of the resource.
	ResourceUri *string `json:"resourceUri,omitempty"`
}

type IdentityDetails_ARM struct {
	// UseSystemAssignedIdentity: Specifies if the BI is protected by System Identity.
	UseSystemAssignedIdentity *bool `json:"useSystemAssignedIdentity,omitempty"`

	// UserAssignedIdentityArmUrl: ARM URL for User Assigned Identity.
	UserAssignedIdentityArmUrl *string `json:"userAssignedIdentityArmUrl,omitempty"`
}

// Policy Info in backupInstance
type PolicyInfo_ARM struct {
	PolicyId *string `json:"policyId,omitempty"`

	// PolicyParameters: Policy parameters for the backup instance
	PolicyParameters *PolicyParameters_ARM `json:"policyParameters,omitempty"`
}

type BaseResourceProperties_ARM struct {
	// DefaultResourceProperties: Mutually exclusive with all other properties
	DefaultResourceProperties *DefaultResourceProperties_ARM `json:"defaultResourceProperties,omitempty"`
}

// MarshalJSON defers JSON marshaling to the first non-nil property, because BaseResourceProperties_ARM represents a discriminated union (JSON OneOf)
func (properties BaseResourceProperties_ARM) MarshalJSON() ([]byte, error) {
	if properties.DefaultResourceProperties != nil {
		return json.Marshal(properties.DefaultResourceProperties)
	}
	return nil, nil
}

// UnmarshalJSON unmarshals the BaseResourceProperties_ARM
func (properties *BaseResourceProperties_ARM) UnmarshalJSON(data []byte) error {
	var rawJson map[string]interface{}
	err := json.Unmarshal(data, &rawJson)
	if err != nil {
		return err
	}
	discriminator := rawJson["objectType"]
	if discriminator == "DefaultResourceProperties" {
		properties.DefaultResourceProperties = &DefaultResourceProperties_ARM{}
		return json.Unmarshal(data, properties.DefaultResourceProperties)
	}

	// No error
	return nil
}

// Parameters in Policy
type PolicyParameters_ARM struct {
	// BackupDatasourceParametersList: Gets or sets the Backup Data Source Parameters
	BackupDatasourceParametersList []BackupDatasourceParameters_ARM `json:"backupDatasourceParametersList,omitempty"`

	// DataStoreParametersList: Gets or sets the DataStore Parameters
	DataStoreParametersList []DataStoreParameters_ARM `json:"dataStoreParametersList,omitempty"`
}

type SecretStoreBasedAuthCredentials_ARM struct {
	// ObjectType: Type of the specific object - used for deserializing
	ObjectType SecretStoreBasedAuthCredentials_ObjectType_ARM `json:"objectType,omitempty"`

	// SecretStoreResource: Secret store resource
	SecretStoreResource *SecretStoreResource_ARM `json:"secretStoreResource,omitempty"`
}

type BackupDatasourceParameters_ARM struct {
	// Blob: Mutually exclusive with all other properties
	Blob *BlobBackupDatasourceParameters_ARM `json:"blobBackupDatasourceParameters,omitempty"`

	// KubernetesCluster: Mutually exclusive with all other properties
	KubernetesCluster *KubernetesClusterBackupDatasourceParameters_ARM `json:"kubernetesClusterBackupDatasourceParameters,omitempty"`
}

// MarshalJSON defers JSON marshaling to the first non-nil property, because BackupDatasourceParameters_ARM represents a discriminated union (JSON OneOf)
func (parameters BackupDatasourceParameters_ARM) MarshalJSON() ([]byte, error) {
	if parameters.Blob != nil {
		return json.Marshal(parameters.Blob)
	}
	if parameters.KubernetesCluster != nil {
		return json.Marshal(parameters.KubernetesCluster)
	}
	return nil, nil
}

// UnmarshalJSON unmarshals the BackupDatasourceParameters_ARM
func (parameters *BackupDatasourceParameters_ARM) UnmarshalJSON(data []byte) error {
	var rawJson map[string]interface{}
	err := json.Unmarshal(data, &rawJson)
	if err != nil {
		return err
	}
	discriminator := rawJson["objectType"]
	if discriminator == "BlobBackupDatasourceParameters" {
		parameters.Blob = &BlobBackupDatasourceParameters_ARM{}
		return json.Unmarshal(data, parameters.Blob)
	}
	if discriminator == "KubernetesClusterBackupDatasourceParameters" {
		parameters.KubernetesCluster = &KubernetesClusterBackupDatasourceParameters_ARM{}
		return json.Unmarshal(data, parameters.KubernetesCluster)
	}

	// No error
	return nil
}

type DataStoreParameters_ARM struct {
	// AzureOperationalStoreParameters: Mutually exclusive with all other properties
	AzureOperationalStoreParameters *AzureOperationalStoreParameters_ARM `json:"azureOperationalStoreParameters,omitempty"`
}

// MarshalJSON defers JSON marshaling to the first non-nil property, because DataStoreParameters_ARM represents a discriminated union (JSON OneOf)
func (parameters DataStoreParameters_ARM) MarshalJSON() ([]byte, error) {
	if parameters.AzureOperationalStoreParameters != nil {
		return json.Marshal(parameters.AzureOperationalStoreParameters)
	}
	return nil, nil
}

// UnmarshalJSON unmarshals the DataStoreParameters_ARM
func (parameters *DataStoreParameters_ARM) UnmarshalJSON(data []byte) error {
	var rawJson map[string]interface{}
	err := json.Unmarshal(data, &rawJson)
	if err != nil {
		return err
	}
	discriminator := rawJson["objectType"]
	if discriminator == "AzureOperationalStoreParameters" {
		parameters.AzureOperationalStoreParameters = &AzureOperationalStoreParameters_ARM{}
		return json.Unmarshal(data, parameters.AzureOperationalStoreParameters)
	}

	// No error
	return nil
}

type DefaultResourceProperties_ARM struct {
	// ObjectType: Type of the specific object - used for deserializing
	ObjectType DefaultResourceProperties_ObjectType_ARM `json:"objectType,omitempty"`
}

// +kubebuilder:validation:Enum={"SecretStoreBasedAuthCredentials"}
type SecretStoreBasedAuthCredentials_ObjectType_ARM string

const SecretStoreBasedAuthCredentials_ObjectType_ARM_SecretStoreBasedAuthCredentials = SecretStoreBasedAuthCredentials_ObjectType_ARM("SecretStoreBasedAuthCredentials")

// Mapping from string to SecretStoreBasedAuthCredentials_ObjectType_ARM
var secretStoreBasedAuthCredentials_ObjectType_ARM_Values = map[string]SecretStoreBasedAuthCredentials_ObjectType_ARM{
	"secretstorebasedauthcredentials": SecretStoreBasedAuthCredentials_ObjectType_ARM_SecretStoreBasedAuthCredentials,
}

// Class representing a secret store resource.
type SecretStoreResource_ARM struct {
	// SecretStoreType: Gets or sets the type of secret store
	SecretStoreType *SecretStoreResource_SecretStoreType_ARM `json:"secretStoreType,omitempty"`

	// Uri: Uri to get to the resource
	Uri *string `json:"uri,omitempty"`

	// Value: Gets or sets value stored in secret store resource
	Value *string `json:"value,omitempty"`
}

type AzureOperationalStoreParameters_ARM struct {
	// DataStoreType: type of datastore; Operational/Vault/Archive
	DataStoreType *AzureOperationalStoreParameters_DataStoreType_ARM `json:"dataStoreType,omitempty"`

	// ObjectType: Type of the specific object - used for deserializing
	ObjectType      AzureOperationalStoreParameters_ObjectType_ARM `json:"objectType,omitempty"`
	ResourceGroupId *string                                        `json:"resourceGroupId,omitempty"`
}

type BlobBackupDatasourceParameters_ARM struct {
	// ContainersList: List of containers to be backed up during configuration of backup of blobs
	ContainersList []string `json:"containersList,omitempty"`

	// ObjectType: Type of the specific object - used for deserializing
	ObjectType BlobBackupDatasourceParameters_ObjectType_ARM `json:"objectType,omitempty"`
}

// +kubebuilder:validation:Enum={"DefaultResourceProperties"}
type DefaultResourceProperties_ObjectType_ARM string

const DefaultResourceProperties_ObjectType_ARM_DefaultResourceProperties = DefaultResourceProperties_ObjectType_ARM("DefaultResourceProperties")

// Mapping from string to DefaultResourceProperties_ObjectType_ARM
var defaultResourceProperties_ObjectType_ARM_Values = map[string]DefaultResourceProperties_ObjectType_ARM{
	"defaultresourceproperties": DefaultResourceProperties_ObjectType_ARM_DefaultResourceProperties,
}

type KubernetesClusterBackupDatasourceParameters_ARM struct {
	// BackupHookReferences: Gets or sets the backup hook references. This property sets the hook reference to be executed
	// during backup.
	BackupHookReferences []NamespacedNameResource_ARM `json:"backupHookReferences,omitempty"`

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
	ObjectType KubernetesClusterBackupDatasourceParameters_ObjectType_ARM `json:"objectType,omitempty"`

	// SnapshotVolumes: Gets or sets the volume snapshot property. This property if enabled will take volume snapshots during
	// backup.
	SnapshotVolumes *bool `json:"snapshotVolumes,omitempty"`
}

// +kubebuilder:validation:Enum={"AzureKeyVault","Invalid"}
type SecretStoreResource_SecretStoreType_ARM string

const (
	SecretStoreResource_SecretStoreType_ARM_AzureKeyVault = SecretStoreResource_SecretStoreType_ARM("AzureKeyVault")
	SecretStoreResource_SecretStoreType_ARM_Invalid       = SecretStoreResource_SecretStoreType_ARM("Invalid")
)

// Mapping from string to SecretStoreResource_SecretStoreType_ARM
var secretStoreResource_SecretStoreType_ARM_Values = map[string]SecretStoreResource_SecretStoreType_ARM{
	"azurekeyvault": SecretStoreResource_SecretStoreType_ARM_AzureKeyVault,
	"invalid":       SecretStoreResource_SecretStoreType_ARM_Invalid,
}

// +kubebuilder:validation:Enum={"ArchiveStore","OperationalStore","VaultStore"}
type AzureOperationalStoreParameters_DataStoreType_ARM string

const (
	AzureOperationalStoreParameters_DataStoreType_ARM_ArchiveStore     = AzureOperationalStoreParameters_DataStoreType_ARM("ArchiveStore")
	AzureOperationalStoreParameters_DataStoreType_ARM_OperationalStore = AzureOperationalStoreParameters_DataStoreType_ARM("OperationalStore")
	AzureOperationalStoreParameters_DataStoreType_ARM_VaultStore       = AzureOperationalStoreParameters_DataStoreType_ARM("VaultStore")
)

// Mapping from string to AzureOperationalStoreParameters_DataStoreType_ARM
var azureOperationalStoreParameters_DataStoreType_ARM_Values = map[string]AzureOperationalStoreParameters_DataStoreType_ARM{
	"archivestore":     AzureOperationalStoreParameters_DataStoreType_ARM_ArchiveStore,
	"operationalstore": AzureOperationalStoreParameters_DataStoreType_ARM_OperationalStore,
	"vaultstore":       AzureOperationalStoreParameters_DataStoreType_ARM_VaultStore,
}

// +kubebuilder:validation:Enum={"AzureOperationalStoreParameters"}
type AzureOperationalStoreParameters_ObjectType_ARM string

const AzureOperationalStoreParameters_ObjectType_ARM_AzureOperationalStoreParameters = AzureOperationalStoreParameters_ObjectType_ARM("AzureOperationalStoreParameters")

// Mapping from string to AzureOperationalStoreParameters_ObjectType_ARM
var azureOperationalStoreParameters_ObjectType_ARM_Values = map[string]AzureOperationalStoreParameters_ObjectType_ARM{
	"azureoperationalstoreparameters": AzureOperationalStoreParameters_ObjectType_ARM_AzureOperationalStoreParameters,
}

// +kubebuilder:validation:Enum={"BlobBackupDatasourceParameters"}
type BlobBackupDatasourceParameters_ObjectType_ARM string

const BlobBackupDatasourceParameters_ObjectType_ARM_BlobBackupDatasourceParameters = BlobBackupDatasourceParameters_ObjectType_ARM("BlobBackupDatasourceParameters")

// Mapping from string to BlobBackupDatasourceParameters_ObjectType_ARM
var blobBackupDatasourceParameters_ObjectType_ARM_Values = map[string]BlobBackupDatasourceParameters_ObjectType_ARM{
	"blobbackupdatasourceparameters": BlobBackupDatasourceParameters_ObjectType_ARM_BlobBackupDatasourceParameters,
}

// +kubebuilder:validation:Enum={"KubernetesClusterBackupDatasourceParameters"}
type KubernetesClusterBackupDatasourceParameters_ObjectType_ARM string

const KubernetesClusterBackupDatasourceParameters_ObjectType_ARM_KubernetesClusterBackupDatasourceParameters = KubernetesClusterBackupDatasourceParameters_ObjectType_ARM("KubernetesClusterBackupDatasourceParameters")

// Mapping from string to KubernetesClusterBackupDatasourceParameters_ObjectType_ARM
var kubernetesClusterBackupDatasourceParameters_ObjectType_ARM_Values = map[string]KubernetesClusterBackupDatasourceParameters_ObjectType_ARM{
	"kubernetesclusterbackupdatasourceparameters": KubernetesClusterBackupDatasourceParameters_ObjectType_ARM_KubernetesClusterBackupDatasourceParameters,
}

// Class to refer resources which contains namespace and name
type NamespacedNameResource_ARM struct {
	// Name: Name of the resource
	Name *string `json:"name,omitempty"`

	// Namespace: Namespace in which the resource exists
	Namespace *string `json:"namespace,omitempty"`
}
