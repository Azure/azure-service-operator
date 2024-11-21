// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package storage

import (
	storage "github.com/Azure/azure-service-operator/v2/api/insights/v1api20220615/storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/configmaps"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/core"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/secrets"
	"github.com/rotisserie/eris"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=insights.azure.com,resources=diagnosticsettings,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=insights.azure.com,resources={diagnosticsettings/status,diagnosticsettings/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20210501preview.DiagnosticSetting
// Generator information:
// - Generated from: /monitor/resource-manager/Microsoft.Insights/preview/2021-05-01-preview/diagnosticsSettings_API.json
// - ARM URI: /{resourceUri}/providers/Microsoft.Insights/diagnosticSettings/{name}
type DiagnosticSetting struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DiagnosticSetting_Spec   `json:"spec,omitempty"`
	Status            DiagnosticSetting_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &DiagnosticSetting{}

// GetConditions returns the conditions of the resource
func (setting *DiagnosticSetting) GetConditions() conditions.Conditions {
	return setting.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (setting *DiagnosticSetting) SetConditions(conditions conditions.Conditions) {
	setting.Status.Conditions = conditions
}

var _ configmaps.Exporter = &DiagnosticSetting{}

// ConfigMapDestinationExpressions returns the Spec.OperatorSpec.ConfigMapExpressions property
func (setting *DiagnosticSetting) ConfigMapDestinationExpressions() []*core.DestinationExpression {
	if setting.Spec.OperatorSpec == nil {
		return nil
	}
	return setting.Spec.OperatorSpec.ConfigMapExpressions
}

var _ secrets.Exporter = &DiagnosticSetting{}

// SecretDestinationExpressions returns the Spec.OperatorSpec.SecretExpressions property
func (setting *DiagnosticSetting) SecretDestinationExpressions() []*core.DestinationExpression {
	if setting.Spec.OperatorSpec == nil {
		return nil
	}
	return setting.Spec.OperatorSpec.SecretExpressions
}

var _ genruntime.KubernetesResource = &DiagnosticSetting{}

// AzureName returns the Azure name of the resource
func (setting *DiagnosticSetting) AzureName() string {
	return setting.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-05-01-preview"
func (setting DiagnosticSetting) GetAPIVersion() string {
	return "2021-05-01-preview"
}

// GetResourceScope returns the scope of the resource
func (setting *DiagnosticSetting) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeExtension
}

// GetSpec returns the specification of this resource
func (setting *DiagnosticSetting) GetSpec() genruntime.ConvertibleSpec {
	return &setting.Spec
}

// GetStatus returns the status of this resource
func (setting *DiagnosticSetting) GetStatus() genruntime.ConvertibleStatus {
	return &setting.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (setting *DiagnosticSetting) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Insights/diagnosticSettings"
func (setting *DiagnosticSetting) GetType() string {
	return "Microsoft.Insights/diagnosticSettings"
}

// NewEmptyStatus returns a new empty (blank) status
func (setting *DiagnosticSetting) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &DiagnosticSetting_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (setting *DiagnosticSetting) Owner() *genruntime.ResourceReference {
	if setting.Spec.Owner == nil {
		return nil
	}

	return setting.Spec.Owner.AsResourceReference()
}

// SetStatus sets the status of this resource
func (setting *DiagnosticSetting) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*DiagnosticSetting_STATUS); ok {
		setting.Status = *st
		return nil
	}

	// Convert status to required version
	var st DiagnosticSetting_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return eris.Wrap(err, "failed to convert status")
	}

	setting.Status = st
	return nil
}

// Hub marks that this DiagnosticSetting is the hub type for conversion
func (setting *DiagnosticSetting) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (setting *DiagnosticSetting) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: setting.Spec.OriginalVersion,
		Kind:    "DiagnosticSetting",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20210501preview.DiagnosticSetting
// Generator information:
// - Generated from: /monitor/resource-manager/Microsoft.Insights/preview/2021-05-01-preview/diagnosticsSettings_API.json
// - ARM URI: /{resourceUri}/providers/Microsoft.Insights/diagnosticSettings/{name}
type DiagnosticSettingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DiagnosticSetting `json:"items"`
}

// Storage version of v1api20210501preview.APIVersion
// +kubebuilder:validation:Enum={"2021-05-01-preview"}
type APIVersion string

const APIVersion_Value = APIVersion("2021-05-01-preview")

// Storage version of v1api20210501preview.DiagnosticSetting_Spec
type DiagnosticSetting_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName string `json:"azureName,omitempty"`

	// EventHubAuthorizationRuleReference: The resource Id for the event hub authorization rule.
	EventHubAuthorizationRuleReference *genruntime.ResourceReference `armReference:"EventHubAuthorizationRuleId" json:"eventHubAuthorizationRuleReference,omitempty"`
	EventHubName                       *string                       `json:"eventHubName,omitempty"`
	LogAnalyticsDestinationType        *string                       `json:"logAnalyticsDestinationType,omitempty"`
	Logs                               []LogSettings                 `json:"logs,omitempty"`

	// MarketplacePartnerReference: The full ARM resource ID of the Marketplace resource to which you would like to send
	// Diagnostic Logs.
	MarketplacePartnerReference *genruntime.ResourceReference  `armReference:"MarketplacePartnerId" json:"marketplacePartnerReference,omitempty"`
	Metrics                     []MetricSettings               `json:"metrics,omitempty"`
	OperatorSpec                *DiagnosticSettingOperatorSpec `json:"operatorSpec,omitempty"`
	OriginalVersion             string                         `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. This resource is an
	// extension resource, which means that any other Azure resource can be its owner.
	Owner            *genruntime.ArbitraryOwnerReference `json:"owner,omitempty"`
	PropertyBag      genruntime.PropertyBag              `json:"$propertyBag,omitempty"`
	ServiceBusRuleId *string                             `json:"serviceBusRuleId,omitempty"`

	// StorageAccountReference: The resource ID of the storage account to which you would like to send Diagnostic Logs.
	StorageAccountReference *genruntime.ResourceReference `armReference:"StorageAccountId" json:"storageAccountReference,omitempty"`

	// WorkspaceReference: The full ARM resource ID of the Log Analytics workspace to which you would like to send Diagnostic
	// Logs. Example:
	// /subscriptions/4b9e8510-67ab-4e9a-95a9-e2f1e570ea9c/resourceGroups/insights-integration/providers/Microsoft.OperationalInsights/workspaces/viruela2
	WorkspaceReference *genruntime.ResourceReference `armReference:"WorkspaceId" json:"workspaceReference,omitempty"`
}

var _ genruntime.ConvertibleSpec = &DiagnosticSetting_Spec{}

// ConvertSpecFrom populates our DiagnosticSetting_Spec from the provided source
func (setting *DiagnosticSetting_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == setting {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(setting)
}

// ConvertSpecTo populates the provided destination from our DiagnosticSetting_Spec
func (setting *DiagnosticSetting_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == setting {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(setting)
}

// Storage version of v1api20210501preview.DiagnosticSetting_STATUS
type DiagnosticSetting_STATUS struct {
	Conditions                  []conditions.Condition  `json:"conditions,omitempty"`
	EventHubAuthorizationRuleId *string                 `json:"eventHubAuthorizationRuleId,omitempty"`
	EventHubName                *string                 `json:"eventHubName,omitempty"`
	Id                          *string                 `json:"id,omitempty"`
	LogAnalyticsDestinationType *string                 `json:"logAnalyticsDestinationType,omitempty"`
	Logs                        []LogSettings_STATUS    `json:"logs,omitempty"`
	MarketplacePartnerId        *string                 `json:"marketplacePartnerId,omitempty"`
	Metrics                     []MetricSettings_STATUS `json:"metrics,omitempty"`
	Name                        *string                 `json:"name,omitempty"`
	PropertyBag                 genruntime.PropertyBag  `json:"$propertyBag,omitempty"`
	ServiceBusRuleId            *string                 `json:"serviceBusRuleId,omitempty"`
	StorageAccountId            *string                 `json:"storageAccountId,omitempty"`
	SystemData                  *SystemData_STATUS      `json:"systemData,omitempty"`
	Type                        *string                 `json:"type,omitempty"`
	WorkspaceId                 *string                 `json:"workspaceId,omitempty"`
}

var _ genruntime.ConvertibleStatus = &DiagnosticSetting_STATUS{}

// ConvertStatusFrom populates our DiagnosticSetting_STATUS from the provided source
func (setting *DiagnosticSetting_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == setting {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(setting)
}

// ConvertStatusTo populates the provided destination from our DiagnosticSetting_STATUS
func (setting *DiagnosticSetting_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == setting {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(setting)
}

// Storage version of v1api20210501preview.DiagnosticSettingOperatorSpec
// Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure
type DiagnosticSettingOperatorSpec struct {
	ConfigMapExpressions []*core.DestinationExpression `json:"configMapExpressions,omitempty"`
	PropertyBag          genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	SecretExpressions    []*core.DestinationExpression `json:"secretExpressions,omitempty"`
}

// Storage version of v1api20210501preview.LogSettings
// Part of MultiTenantDiagnosticSettings. Specifies the settings for a particular log.
type LogSettings struct {
	Category        *string                `json:"category,omitempty"`
	CategoryGroup   *string                `json:"categoryGroup,omitempty"`
	Enabled         *bool                  `json:"enabled,omitempty"`
	PropertyBag     genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	RetentionPolicy *RetentionPolicy       `json:"retentionPolicy,omitempty"`
}

// Storage version of v1api20210501preview.LogSettings_STATUS
// Part of MultiTenantDiagnosticSettings. Specifies the settings for a particular log.
type LogSettings_STATUS struct {
	Category        *string                 `json:"category,omitempty"`
	CategoryGroup   *string                 `json:"categoryGroup,omitempty"`
	Enabled         *bool                   `json:"enabled,omitempty"`
	PropertyBag     genruntime.PropertyBag  `json:"$propertyBag,omitempty"`
	RetentionPolicy *RetentionPolicy_STATUS `json:"retentionPolicy,omitempty"`
}

// Storage version of v1api20210501preview.MetricSettings
// Part of MultiTenantDiagnosticSettings. Specifies the settings for a particular metric.
type MetricSettings struct {
	Category        *string                `json:"category,omitempty"`
	Enabled         *bool                  `json:"enabled,omitempty"`
	PropertyBag     genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	RetentionPolicy *RetentionPolicy       `json:"retentionPolicy,omitempty"`
	TimeGrain       *string                `json:"timeGrain,omitempty"`
}

// Storage version of v1api20210501preview.MetricSettings_STATUS
// Part of MultiTenantDiagnosticSettings. Specifies the settings for a particular metric.
type MetricSettings_STATUS struct {
	Category        *string                 `json:"category,omitempty"`
	Enabled         *bool                   `json:"enabled,omitempty"`
	PropertyBag     genruntime.PropertyBag  `json:"$propertyBag,omitempty"`
	RetentionPolicy *RetentionPolicy_STATUS `json:"retentionPolicy,omitempty"`
	TimeGrain       *string                 `json:"timeGrain,omitempty"`
}

// Storage version of v1api20210501preview.SystemData_STATUS
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

// AssignProperties_From_SystemData_STATUS populates our SystemData_STATUS from the provided source SystemData_STATUS
func (data *SystemData_STATUS) AssignProperties_From_SystemData_STATUS(source *storage.SystemData_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// CreatedAt
	data.CreatedAt = genruntime.ClonePointerToString(source.CreatedAt)

	// CreatedBy
	data.CreatedBy = genruntime.ClonePointerToString(source.CreatedBy)

	// CreatedByType
	data.CreatedByType = genruntime.ClonePointerToString(source.CreatedByType)

	// LastModifiedAt
	data.LastModifiedAt = genruntime.ClonePointerToString(source.LastModifiedAt)

	// LastModifiedBy
	data.LastModifiedBy = genruntime.ClonePointerToString(source.LastModifiedBy)

	// LastModifiedByType
	data.LastModifiedByType = genruntime.ClonePointerToString(source.LastModifiedByType)

	// Update the property bag
	if len(propertyBag) > 0 {
		data.PropertyBag = propertyBag
	} else {
		data.PropertyBag = nil
	}

	// Invoke the augmentConversionForSystemData_STATUS interface (if implemented) to customize the conversion
	var dataAsAny any = data
	if augmentedData, ok := dataAsAny.(augmentConversionForSystemData_STATUS); ok {
		err := augmentedData.AssignPropertiesFrom(source)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_SystemData_STATUS populates the provided destination SystemData_STATUS from our SystemData_STATUS
func (data *SystemData_STATUS) AssignProperties_To_SystemData_STATUS(destination *storage.SystemData_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(data.PropertyBag)

	// CreatedAt
	destination.CreatedAt = genruntime.ClonePointerToString(data.CreatedAt)

	// CreatedBy
	destination.CreatedBy = genruntime.ClonePointerToString(data.CreatedBy)

	// CreatedByType
	destination.CreatedByType = genruntime.ClonePointerToString(data.CreatedByType)

	// LastModifiedAt
	destination.LastModifiedAt = genruntime.ClonePointerToString(data.LastModifiedAt)

	// LastModifiedBy
	destination.LastModifiedBy = genruntime.ClonePointerToString(data.LastModifiedBy)

	// LastModifiedByType
	destination.LastModifiedByType = genruntime.ClonePointerToString(data.LastModifiedByType)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForSystemData_STATUS interface (if implemented) to customize the conversion
	var dataAsAny any = data
	if augmentedData, ok := dataAsAny.(augmentConversionForSystemData_STATUS); ok {
		err := augmentedData.AssignPropertiesTo(destination)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

type augmentConversionForSystemData_STATUS interface {
	AssignPropertiesFrom(src *storage.SystemData_STATUS) error
	AssignPropertiesTo(dst *storage.SystemData_STATUS) error
}

// Storage version of v1api20210501preview.RetentionPolicy
// Specifies the retention policy for the log.
type RetentionPolicy struct {
	Days        *int                   `json:"days,omitempty"`
	Enabled     *bool                  `json:"enabled,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20210501preview.RetentionPolicy_STATUS
// Specifies the retention policy for the log.
type RetentionPolicy_STATUS struct {
	Days        *int                   `json:"days,omitempty"`
	Enabled     *bool                  `json:"enabled,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

func init() {
	SchemeBuilder.Register(&DiagnosticSetting{}, &DiagnosticSettingList{})
}
