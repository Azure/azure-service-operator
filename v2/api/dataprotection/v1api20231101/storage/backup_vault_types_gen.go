// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package storage

import (
	"context"
	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/configmaps"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/core"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/secrets"
	"github.com/go-logr/logr"
	"github.com/rotisserie/eris"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// +kubebuilder:rbac:groups=dataprotection.azure.com,resources=backupvaults,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=dataprotection.azure.com,resources={backupvaults/status,backupvaults/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20231101.BackupVault
// Generator information:
// - Generated from: /dataprotection/resource-manager/Microsoft.DataProtection/stable/2023-11-01/dataprotection.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataProtection/backupVaults/{vaultName}
type BackupVault struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BackupVault_Spec           `json:"spec,omitempty"`
	Status            BackupVaultResource_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &BackupVault{}

// GetConditions returns the conditions of the resource
func (vault *BackupVault) GetConditions() conditions.Conditions {
	return vault.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (vault *BackupVault) SetConditions(conditions conditions.Conditions) {
	vault.Status.Conditions = conditions
}

var _ configmaps.Exporter = &BackupVault{}

// ConfigMapDestinationExpressions returns the Spec.OperatorSpec.ConfigMapExpressions property
func (vault *BackupVault) ConfigMapDestinationExpressions() []*core.DestinationExpression {
	if vault.Spec.OperatorSpec == nil {
		return nil
	}
	return vault.Spec.OperatorSpec.ConfigMapExpressions
}

var _ secrets.Exporter = &BackupVault{}

// SecretDestinationExpressions returns the Spec.OperatorSpec.SecretExpressions property
func (vault *BackupVault) SecretDestinationExpressions() []*core.DestinationExpression {
	if vault.Spec.OperatorSpec == nil {
		return nil
	}
	return vault.Spec.OperatorSpec.SecretExpressions
}

var _ genruntime.KubernetesConfigExporter = &BackupVault{}

// ExportKubernetesConfigMaps defines a resource which can create ConfigMaps in Kubernetes.
func (vault *BackupVault) ExportKubernetesConfigMaps(_ context.Context, _ genruntime.MetaObject, _ *genericarmclient.GenericClient, _ logr.Logger) ([]client.Object, error) {
	collector := configmaps.NewCollector(vault.Namespace)
	if vault.Spec.OperatorSpec != nil && vault.Spec.OperatorSpec.ConfigMaps != nil {
		if vault.Status.Identity != nil {
			if vault.Status.Identity.PrincipalId != nil {
				collector.AddValue(vault.Spec.OperatorSpec.ConfigMaps.PrincipalId, *vault.Status.Identity.PrincipalId)
			}
		}
	}
	result, err := collector.Values()
	if err != nil {
		return nil, err
	}
	return configmaps.SliceToClientObjectSlice(result), nil
}

var _ genruntime.KubernetesResource = &BackupVault{}

// AzureName returns the Azure name of the resource
func (vault *BackupVault) AzureName() string {
	return vault.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2023-11-01"
func (vault BackupVault) GetAPIVersion() string {
	return "2023-11-01"
}

// GetResourceScope returns the scope of the resource
func (vault *BackupVault) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (vault *BackupVault) GetSpec() genruntime.ConvertibleSpec {
	return &vault.Spec
}

// GetStatus returns the status of this resource
func (vault *BackupVault) GetStatus() genruntime.ConvertibleStatus {
	return &vault.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (vault *BackupVault) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DataProtection/backupVaults"
func (vault *BackupVault) GetType() string {
	return "Microsoft.DataProtection/backupVaults"
}

// NewEmptyStatus returns a new empty (blank) status
func (vault *BackupVault) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &BackupVaultResource_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (vault *BackupVault) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(vault.Spec)
	return vault.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (vault *BackupVault) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*BackupVaultResource_STATUS); ok {
		vault.Status = *st
		return nil
	}

	// Convert status to required version
	var st BackupVaultResource_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return eris.Wrap(err, "failed to convert status")
	}

	vault.Status = st
	return nil
}

// Hub marks that this BackupVault is the hub type for conversion
func (vault *BackupVault) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (vault *BackupVault) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: vault.Spec.OriginalVersion,
		Kind:    "BackupVault",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20231101.BackupVault
// Generator information:
// - Generated from: /dataprotection/resource-manager/Microsoft.DataProtection/stable/2023-11-01/dataprotection.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataProtection/backupVaults/{vaultName}
type BackupVaultList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BackupVault `json:"items"`
}

// Storage version of v1api20231101.APIVersion
// +kubebuilder:validation:Enum={"2023-11-01"}
type APIVersion string

const APIVersion_Value = APIVersion("2023-11-01")

// Storage version of v1api20231101.BackupVault_Spec
type BackupVault_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName       string                   `json:"azureName,omitempty"`
	Identity        *DppIdentityDetails      `json:"identity,omitempty"`
	Location        *string                  `json:"location,omitempty"`
	OperatorSpec    *BackupVaultOperatorSpec `json:"operatorSpec,omitempty"`
	OriginalVersion string                   `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a resources.azure.com/ResourceGroup resource
	Owner       *genruntime.KnownResourceReference `group:"resources.azure.com" json:"owner,omitempty" kind:"ResourceGroup"`
	Properties  *BackupVaultSpec                   `json:"properties,omitempty"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	Tags        map[string]string                  `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &BackupVault_Spec{}

// ConvertSpecFrom populates our BackupVault_Spec from the provided source
func (vault *BackupVault_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == vault {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(vault)
}

// ConvertSpecTo populates the provided destination from our BackupVault_Spec
func (vault *BackupVault_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == vault {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(vault)
}

// Storage version of v1api20231101.BackupVaultResource_STATUS
// Backup Vault Resource
type BackupVaultResource_STATUS struct {
	Conditions  []conditions.Condition     `json:"conditions,omitempty"`
	ETag        *string                    `json:"eTag,omitempty"`
	Id          *string                    `json:"id,omitempty"`
	Identity    *DppIdentityDetails_STATUS `json:"identity,omitempty"`
	Location    *string                    `json:"location,omitempty"`
	Name        *string                    `json:"name,omitempty"`
	Properties  *BackupVault_STATUS        `json:"properties,omitempty"`
	PropertyBag genruntime.PropertyBag     `json:"$propertyBag,omitempty"`
	SystemData  *SystemData_STATUS         `json:"systemData,omitempty"`
	Tags        map[string]string          `json:"tags,omitempty"`
	Type        *string                    `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &BackupVaultResource_STATUS{}

// ConvertStatusFrom populates our BackupVaultResource_STATUS from the provided source
func (resource *BackupVaultResource_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == resource {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(resource)
}

// ConvertStatusTo populates the provided destination from our BackupVaultResource_STATUS
func (resource *BackupVaultResource_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == resource {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(resource)
}

// Storage version of v1api20231101.BackupVault_STATUS
// Backup Vault
type BackupVault_STATUS struct {
	FeatureSettings                 *FeatureSettings_STATUS     `json:"featureSettings,omitempty"`
	IsVaultProtectedByResourceGuard *bool                       `json:"isVaultProtectedByResourceGuard,omitempty"`
	MonitoringSettings              *MonitoringSettings_STATUS  `json:"monitoringSettings,omitempty"`
	PropertyBag                     genruntime.PropertyBag      `json:"$propertyBag,omitempty"`
	ProvisioningState               *string                     `json:"provisioningState,omitempty"`
	ReplicatedRegions               []string                    `json:"replicatedRegions,omitempty"`
	ResourceMoveDetails             *ResourceMoveDetails_STATUS `json:"resourceMoveDetails,omitempty"`
	ResourceMoveState               *string                     `json:"resourceMoveState,omitempty"`
	SecureScore                     *string                     `json:"secureScore,omitempty"`
	SecuritySettings                *SecuritySettings_STATUS    `json:"securitySettings,omitempty"`
	StorageSettings                 []StorageSetting_STATUS     `json:"storageSettings,omitempty"`
}

// Storage version of v1api20231101.BackupVaultOperatorSpec
// Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure
type BackupVaultOperatorSpec struct {
	ConfigMapExpressions []*core.DestinationExpression  `json:"configMapExpressions,omitempty"`
	ConfigMaps           *BackupVaultOperatorConfigMaps `json:"configMaps,omitempty"`
	PropertyBag          genruntime.PropertyBag         `json:"$propertyBag,omitempty"`
	SecretExpressions    []*core.DestinationExpression  `json:"secretExpressions,omitempty"`
}

// Storage version of v1api20231101.BackupVaultSpec
// Backup Vault
type BackupVaultSpec struct {
	FeatureSettings    *FeatureSettings       `json:"featureSettings,omitempty"`
	MonitoringSettings *MonitoringSettings    `json:"monitoringSettings,omitempty"`
	PropertyBag        genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ReplicatedRegions  []string               `json:"replicatedRegions,omitempty"`
	SecuritySettings   *SecuritySettings      `json:"securitySettings,omitempty"`
	StorageSettings    []StorageSetting       `json:"storageSettings,omitempty"`
}

// Storage version of v1api20231101.DppIdentityDetails
// Identity details
type DppIdentityDetails struct {
	PropertyBag            genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	Type                   *string                       `json:"type,omitempty"`
	UserAssignedIdentities []UserAssignedIdentityDetails `json:"userAssignedIdentities,omitempty"`
}

// Storage version of v1api20231101.DppIdentityDetails_STATUS
// Identity details
type DppIdentityDetails_STATUS struct {
	PrincipalId            *string                                `json:"principalId,omitempty"`
	PropertyBag            genruntime.PropertyBag                 `json:"$propertyBag,omitempty"`
	TenantId               *string                                `json:"tenantId,omitempty"`
	Type                   *string                                `json:"type,omitempty"`
	UserAssignedIdentities map[string]UserAssignedIdentity_STATUS `json:"userAssignedIdentities,omitempty"`
}

// Storage version of v1api20231101.SystemData_STATUS
// Metadata pertaining to creation and last modification of the resource.
type SystemData_STATUS struct {
	CreatedAt          *string                `json:"createdAt,omitempty"`
	CreatedBy          *string                `json:"createdBy,omitempty"`
	CreatedByType      *string                `json:"createdByType,omitempty"`
	LastModifiedAt     *string                `json:"lastModifiedAt,omitempty"`
	LastModifiedBy     *string                `json:"lastModifiedBy,omitempty"`
	LastModifiedByType *string                `json:"lastModifiedByType,omitempty"`
	PropertyBag        genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20231101.BackupVaultOperatorConfigMaps
type BackupVaultOperatorConfigMaps struct {
	PrincipalId *genruntime.ConfigMapDestination `json:"principalId,omitempty"`
	PropertyBag genruntime.PropertyBag           `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20231101.FeatureSettings
// Class containing feature settings of vault
type FeatureSettings struct {
	CrossRegionRestoreSettings       *CrossRegionRestoreSettings       `json:"crossRegionRestoreSettings,omitempty"`
	CrossSubscriptionRestoreSettings *CrossSubscriptionRestoreSettings `json:"crossSubscriptionRestoreSettings,omitempty"`
	PropertyBag                      genruntime.PropertyBag            `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20231101.FeatureSettings_STATUS
// Class containing feature settings of vault
type FeatureSettings_STATUS struct {
	CrossRegionRestoreSettings       *CrossRegionRestoreSettings_STATUS       `json:"crossRegionRestoreSettings,omitempty"`
	CrossSubscriptionRestoreSettings *CrossSubscriptionRestoreSettings_STATUS `json:"crossSubscriptionRestoreSettings,omitempty"`
	PropertyBag                      genruntime.PropertyBag                   `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20231101.MonitoringSettings
// Monitoring Settings
type MonitoringSettings struct {
	AzureMonitorAlertSettings *AzureMonitorAlertSettings `json:"azureMonitorAlertSettings,omitempty"`
	PropertyBag               genruntime.PropertyBag     `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20231101.MonitoringSettings_STATUS
// Monitoring Settings
type MonitoringSettings_STATUS struct {
	AzureMonitorAlertSettings *AzureMonitorAlertSettings_STATUS `json:"azureMonitorAlertSettings,omitempty"`
	PropertyBag               genruntime.PropertyBag            `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20231101.ResourceMoveDetails_STATUS
// ResourceMoveDetails will be returned in response to GetResource call from ARM
type ResourceMoveDetails_STATUS struct {
	CompletionTimeUtc  *string                `json:"completionTimeUtc,omitempty"`
	OperationId        *string                `json:"operationId,omitempty"`
	PropertyBag        genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	SourceResourcePath *string                `json:"sourceResourcePath,omitempty"`
	StartTimeUtc       *string                `json:"startTimeUtc,omitempty"`
	TargetResourcePath *string                `json:"targetResourcePath,omitempty"`
}

// Storage version of v1api20231101.SecuritySettings
// Class containing security settings of vault
type SecuritySettings struct {
	ImmutabilitySettings *ImmutabilitySettings  `json:"immutabilitySettings,omitempty"`
	PropertyBag          genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	SoftDeleteSettings   *SoftDeleteSettings    `json:"softDeleteSettings,omitempty"`
}

// Storage version of v1api20231101.SecuritySettings_STATUS
// Class containing security settings of vault
type SecuritySettings_STATUS struct {
	ImmutabilitySettings *ImmutabilitySettings_STATUS `json:"immutabilitySettings,omitempty"`
	PropertyBag          genruntime.PropertyBag       `json:"$propertyBag,omitempty"`
	SoftDeleteSettings   *SoftDeleteSettings_STATUS   `json:"softDeleteSettings,omitempty"`
}

// Storage version of v1api20231101.StorageSetting
// Storage setting
type StorageSetting struct {
	DatastoreType *string                `json:"datastoreType,omitempty"`
	PropertyBag   genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Type          *string                `json:"type,omitempty"`
}

// Storage version of v1api20231101.StorageSetting_STATUS
// Storage setting
type StorageSetting_STATUS struct {
	DatastoreType *string                `json:"datastoreType,omitempty"`
	PropertyBag   genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Type          *string                `json:"type,omitempty"`
}

// Storage version of v1api20231101.UserAssignedIdentity_STATUS
// User assigned identity properties
type UserAssignedIdentity_STATUS struct {
	ClientId    *string                `json:"clientId,omitempty"`
	PrincipalId *string                `json:"principalId,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20231101.UserAssignedIdentityDetails
// Information about the user assigned identity for the resource
type UserAssignedIdentityDetails struct {
	PropertyBag genruntime.PropertyBag       `json:"$propertyBag,omitempty"`
	Reference   genruntime.ResourceReference `armReference:"Reference" json:"reference,omitempty"`
}

// Storage version of v1api20231101.AzureMonitorAlertSettings
// Settings for Azure Monitor based alerts
type AzureMonitorAlertSettings struct {
	AlertsForAllJobFailures *string                `json:"alertsForAllJobFailures,omitempty"`
	PropertyBag             genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20231101.AzureMonitorAlertSettings_STATUS
// Settings for Azure Monitor based alerts
type AzureMonitorAlertSettings_STATUS struct {
	AlertsForAllJobFailures *string                `json:"alertsForAllJobFailures,omitempty"`
	PropertyBag             genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20231101.CrossRegionRestoreSettings
type CrossRegionRestoreSettings struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	State       *string                `json:"state,omitempty"`
}

// Storage version of v1api20231101.CrossRegionRestoreSettings_STATUS
type CrossRegionRestoreSettings_STATUS struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	State       *string                `json:"state,omitempty"`
}

// Storage version of v1api20231101.CrossSubscriptionRestoreSettings
// CrossSubscriptionRestore Settings
type CrossSubscriptionRestoreSettings struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	State       *string                `json:"state,omitempty"`
}

// Storage version of v1api20231101.CrossSubscriptionRestoreSettings_STATUS
// CrossSubscriptionRestore Settings
type CrossSubscriptionRestoreSettings_STATUS struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	State       *string                `json:"state,omitempty"`
}

// Storage version of v1api20231101.ImmutabilitySettings
// Immutability Settings at vault level
type ImmutabilitySettings struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	State       *string                `json:"state,omitempty"`
}

// Storage version of v1api20231101.ImmutabilitySettings_STATUS
// Immutability Settings at vault level
type ImmutabilitySettings_STATUS struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	State       *string                `json:"state,omitempty"`
}

// Storage version of v1api20231101.SoftDeleteSettings
// Soft delete related settings
type SoftDeleteSettings struct {
	PropertyBag             genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	RetentionDurationInDays *float64               `json:"retentionDurationInDays,omitempty"`
	State                   *string                `json:"state,omitempty"`
}

// Storage version of v1api20231101.SoftDeleteSettings_STATUS
// Soft delete related settings
type SoftDeleteSettings_STATUS struct {
	PropertyBag             genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	RetentionDurationInDays *float64               `json:"retentionDurationInDays,omitempty"`
	State                   *string                `json:"state,omitempty"`
}

func init() {
	SchemeBuilder.Register(&BackupVault{}, &BackupVaultList{})
}
