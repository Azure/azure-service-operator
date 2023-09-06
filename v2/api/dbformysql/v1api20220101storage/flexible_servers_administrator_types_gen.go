// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20220101storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=dbformysql.azure.com,resources=flexibleserversadministrators,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=dbformysql.azure.com,resources={flexibleserversadministrators/status,flexibleserversadministrators/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20220101.FlexibleServersAdministrator
// Generator information:
// - Generated from: /mysql/resource-manager/Microsoft.DBforMySQL/AAD/stable/2022-01-01/AzureADAdministrator.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/flexibleServers/{serverName}/administrators/{administratorName}
type FlexibleServersAdministrator struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FlexibleServers_Administrator_Spec   `json:"spec,omitempty"`
	Status            FlexibleServers_Administrator_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &FlexibleServersAdministrator{}

// GetConditions returns the conditions of the resource
func (administrator *FlexibleServersAdministrator) GetConditions() conditions.Conditions {
	return administrator.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (administrator *FlexibleServersAdministrator) SetConditions(conditions conditions.Conditions) {
	administrator.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &FlexibleServersAdministrator{}

// AzureName returns the Azure name of the resource (always "ActiveDirectory")
func (administrator *FlexibleServersAdministrator) AzureName() string {
	return "ActiveDirectory"
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2022-01-01"
func (administrator FlexibleServersAdministrator) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (administrator *FlexibleServersAdministrator) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (administrator *FlexibleServersAdministrator) GetSpec() genruntime.ConvertibleSpec {
	return &administrator.Spec
}

// GetStatus returns the status of this resource
func (administrator *FlexibleServersAdministrator) GetStatus() genruntime.ConvertibleStatus {
	return &administrator.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DBforMySQL/flexibleServers/administrators"
func (administrator *FlexibleServersAdministrator) GetType() string {
	return "Microsoft.DBforMySQL/flexibleServers/administrators"
}

// NewEmptyStatus returns a new empty (blank) status
func (administrator *FlexibleServersAdministrator) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &FlexibleServers_Administrator_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (administrator *FlexibleServersAdministrator) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(administrator.Spec)
	return administrator.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (administrator *FlexibleServersAdministrator) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*FlexibleServers_Administrator_STATUS); ok {
		administrator.Status = *st
		return nil
	}

	// Convert status to required version
	var st FlexibleServers_Administrator_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	administrator.Status = st
	return nil
}

// Hub marks that this FlexibleServersAdministrator is the hub type for conversion
func (administrator *FlexibleServersAdministrator) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (administrator *FlexibleServersAdministrator) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: administrator.Spec.OriginalVersion,
		Kind:    "FlexibleServersAdministrator",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20220101.FlexibleServersAdministrator
// Generator information:
// - Generated from: /mysql/resource-manager/Microsoft.DBforMySQL/AAD/stable/2022-01-01/AzureADAdministrator.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/flexibleServers/{serverName}/administrators/{administratorName}
type FlexibleServersAdministratorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FlexibleServersAdministrator `json:"items"`
}

// Storage version of v1api20220101.APIVersion
// +kubebuilder:validation:Enum={"2022-01-01"}
type APIVersion string

const APIVersion_Value = APIVersion("2022-01-01")

// Storage version of v1api20220101.FlexibleServers_Administrator_Spec
type FlexibleServers_Administrator_Spec struct {
	AdministratorType *string `json:"administratorType,omitempty"`

	// IdentityResourceReference: The resource id of the identity used for AAD Authentication.
	IdentityResourceReference *genruntime.ResourceReference `armReference:"IdentityResourceId" json:"identityResourceReference,omitempty"`
	Login                     *string                       `json:"login,omitempty"`
	OriginalVersion           string                        `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a dbformysql.azure.com/FlexibleServer resource
	Owner              *genruntime.KnownResourceReference `group:"dbformysql.azure.com" json:"owner,omitempty" kind:"FlexibleServer"`
	PropertyBag        genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	Sid                *string                            `json:"sid,omitempty" optionalConfigMapPair:"Sid"`
	SidFromConfig      *genruntime.ConfigMapReference     `json:"sidFromConfig,omitempty" optionalConfigMapPair:"Sid"`
	TenantId           *string                            `json:"tenantId,omitempty" optionalConfigMapPair:"TenantId"`
	TenantIdFromConfig *genruntime.ConfigMapReference     `json:"tenantIdFromConfig,omitempty" optionalConfigMapPair:"TenantId"`
}

var _ genruntime.ConvertibleSpec = &FlexibleServers_Administrator_Spec{}

// ConvertSpecFrom populates our FlexibleServers_Administrator_Spec from the provided source
func (administrator *FlexibleServers_Administrator_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == administrator {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(administrator)
}

// ConvertSpecTo populates the provided destination from our FlexibleServers_Administrator_Spec
func (administrator *FlexibleServers_Administrator_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == administrator {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(administrator)
}

// Storage version of v1api20220101.FlexibleServers_Administrator_STATUS
type FlexibleServers_Administrator_STATUS struct {
	AdministratorType  *string                `json:"administratorType,omitempty"`
	Conditions         []conditions.Condition `json:"conditions,omitempty"`
	Id                 *string                `json:"id,omitempty"`
	IdentityResourceId *string                `json:"identityResourceId,omitempty"`
	Login              *string                `json:"login,omitempty"`
	Name               *string                `json:"name,omitempty"`
	PropertyBag        genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Sid                *string                `json:"sid,omitempty"`
	SystemData         *SystemData_STATUS     `json:"systemData,omitempty"`
	TenantId           *string                `json:"tenantId,omitempty"`
	Type               *string                `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &FlexibleServers_Administrator_STATUS{}

// ConvertStatusFrom populates our FlexibleServers_Administrator_STATUS from the provided source
func (administrator *FlexibleServers_Administrator_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == administrator {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(administrator)
}

// ConvertStatusTo populates the provided destination from our FlexibleServers_Administrator_STATUS
func (administrator *FlexibleServers_Administrator_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == administrator {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(administrator)
}

// Storage version of v1api20220101.SystemData_STATUS
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

func init() {
	SchemeBuilder.Register(&FlexibleServersAdministrator{}, &FlexibleServersAdministratorList{})
}
