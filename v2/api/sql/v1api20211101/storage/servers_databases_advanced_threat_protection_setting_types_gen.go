// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=sql.azure.com,resources=serversdatabasesadvancedthreatprotectionsettings,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=sql.azure.com,resources={serversdatabasesadvancedthreatprotectionsettings/status,serversdatabasesadvancedthreatprotectionsettings/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20211101.ServersDatabasesAdvancedThreatProtectionSetting
// Generator information:
// - Generated from: /sql/resource-manager/Microsoft.Sql/stable/2021-11-01/DatabaseAdvancedThreatProtectionSettings.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/advancedThreatProtectionSettings/Default
type ServersDatabasesAdvancedThreatProtectionSetting struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServersDatabasesAdvancedThreatProtectionSetting_Spec   `json:"spec,omitempty"`
	Status            ServersDatabasesAdvancedThreatProtectionSetting_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &ServersDatabasesAdvancedThreatProtectionSetting{}

// GetConditions returns the conditions of the resource
func (setting *ServersDatabasesAdvancedThreatProtectionSetting) GetConditions() conditions.Conditions {
	return setting.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (setting *ServersDatabasesAdvancedThreatProtectionSetting) SetConditions(conditions conditions.Conditions) {
	setting.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &ServersDatabasesAdvancedThreatProtectionSetting{}

// AzureName returns the Azure name of the resource (always "Default")
func (setting *ServersDatabasesAdvancedThreatProtectionSetting) AzureName() string {
	return "Default"
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-11-01"
func (setting ServersDatabasesAdvancedThreatProtectionSetting) GetAPIVersion() string {
	return "2021-11-01"
}

// GetResourceScope returns the scope of the resource
func (setting *ServersDatabasesAdvancedThreatProtectionSetting) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (setting *ServersDatabasesAdvancedThreatProtectionSetting) GetSpec() genruntime.ConvertibleSpec {
	return &setting.Spec
}

// GetStatus returns the status of this resource
func (setting *ServersDatabasesAdvancedThreatProtectionSetting) GetStatus() genruntime.ConvertibleStatus {
	return &setting.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (setting *ServersDatabasesAdvancedThreatProtectionSetting) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Sql/servers/databases/advancedThreatProtectionSettings"
func (setting *ServersDatabasesAdvancedThreatProtectionSetting) GetType() string {
	return "Microsoft.Sql/servers/databases/advancedThreatProtectionSettings"
}

// NewEmptyStatus returns a new empty (blank) status
func (setting *ServersDatabasesAdvancedThreatProtectionSetting) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &ServersDatabasesAdvancedThreatProtectionSetting_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (setting *ServersDatabasesAdvancedThreatProtectionSetting) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(setting.Spec)
	return setting.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (setting *ServersDatabasesAdvancedThreatProtectionSetting) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*ServersDatabasesAdvancedThreatProtectionSetting_STATUS); ok {
		setting.Status = *st
		return nil
	}

	// Convert status to required version
	var st ServersDatabasesAdvancedThreatProtectionSetting_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	setting.Status = st
	return nil
}

// Hub marks that this ServersDatabasesAdvancedThreatProtectionSetting is the hub type for conversion
func (setting *ServersDatabasesAdvancedThreatProtectionSetting) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (setting *ServersDatabasesAdvancedThreatProtectionSetting) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: setting.Spec.OriginalVersion,
		Kind:    "ServersDatabasesAdvancedThreatProtectionSetting",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20211101.ServersDatabasesAdvancedThreatProtectionSetting
// Generator information:
// - Generated from: /sql/resource-manager/Microsoft.Sql/stable/2021-11-01/DatabaseAdvancedThreatProtectionSettings.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/advancedThreatProtectionSettings/Default
type ServersDatabasesAdvancedThreatProtectionSettingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServersDatabasesAdvancedThreatProtectionSetting `json:"items"`
}

// Storage version of v1api20211101.ServersDatabasesAdvancedThreatProtectionSetting_Spec
type ServersDatabasesAdvancedThreatProtectionSetting_Spec struct {
	OriginalVersion string `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a sql.azure.com/ServersDatabase resource
	Owner       *genruntime.KnownResourceReference `group:"sql.azure.com" json:"owner,omitempty" kind:"ServersDatabase"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	State       *string                            `json:"state,omitempty"`
}

var _ genruntime.ConvertibleSpec = &ServersDatabasesAdvancedThreatProtectionSetting_Spec{}

// ConvertSpecFrom populates our ServersDatabasesAdvancedThreatProtectionSetting_Spec from the provided source
func (setting *ServersDatabasesAdvancedThreatProtectionSetting_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == setting {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(setting)
}

// ConvertSpecTo populates the provided destination from our ServersDatabasesAdvancedThreatProtectionSetting_Spec
func (setting *ServersDatabasesAdvancedThreatProtectionSetting_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == setting {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(setting)
}

// Storage version of v1api20211101.ServersDatabasesAdvancedThreatProtectionSetting_STATUS
type ServersDatabasesAdvancedThreatProtectionSetting_STATUS struct {
	Conditions   []conditions.Condition `json:"conditions,omitempty"`
	CreationTime *string                `json:"creationTime,omitempty"`
	Id           *string                `json:"id,omitempty"`
	Name         *string                `json:"name,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	State        *string                `json:"state,omitempty"`
	SystemData   *SystemData_STATUS     `json:"systemData,omitempty"`
	Type         *string                `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &ServersDatabasesAdvancedThreatProtectionSetting_STATUS{}

// ConvertStatusFrom populates our ServersDatabasesAdvancedThreatProtectionSetting_STATUS from the provided source
func (setting *ServersDatabasesAdvancedThreatProtectionSetting_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == setting {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(setting)
}

// ConvertStatusTo populates the provided destination from our ServersDatabasesAdvancedThreatProtectionSetting_STATUS
func (setting *ServersDatabasesAdvancedThreatProtectionSetting_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == setting {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(setting)
}

func init() {
	SchemeBuilder.Register(&ServersDatabasesAdvancedThreatProtectionSetting{}, &ServersDatabasesAdvancedThreatProtectionSettingList{})
}
