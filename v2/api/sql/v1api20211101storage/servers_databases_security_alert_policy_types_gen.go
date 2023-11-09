// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20211101storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=sql.azure.com,resources=serversdatabasessecurityalertpolicies,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=sql.azure.com,resources={serversdatabasessecurityalertpolicies/status,serversdatabasessecurityalertpolicies/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20211101.ServersDatabasesSecurityAlertPolicy
// Generator information:
// - Generated from: /sql/resource-manager/Microsoft.Sql/stable/2021-11-01/DatabaseSecurityAlertPolicies.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/securityAlertPolicies/default
type ServersDatabasesSecurityAlertPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Servers_Databases_SecurityAlertPolicy_Spec   `json:"spec,omitempty"`
	Status            Servers_Databases_SecurityAlertPolicy_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &ServersDatabasesSecurityAlertPolicy{}

// GetConditions returns the conditions of the resource
func (policy *ServersDatabasesSecurityAlertPolicy) GetConditions() conditions.Conditions {
	return policy.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (policy *ServersDatabasesSecurityAlertPolicy) SetConditions(conditions conditions.Conditions) {
	policy.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &ServersDatabasesSecurityAlertPolicy{}

// AzureName returns the Azure name of the resource (always "default")
func (policy *ServersDatabasesSecurityAlertPolicy) AzureName() string {
	return "default"
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-11-01"
func (policy ServersDatabasesSecurityAlertPolicy) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (policy *ServersDatabasesSecurityAlertPolicy) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (policy *ServersDatabasesSecurityAlertPolicy) GetSpec() genruntime.ConvertibleSpec {
	return &policy.Spec
}

// GetStatus returns the status of this resource
func (policy *ServersDatabasesSecurityAlertPolicy) GetStatus() genruntime.ConvertibleStatus {
	return &policy.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (policy *ServersDatabasesSecurityAlertPolicy) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Sql/servers/databases/securityAlertPolicies"
func (policy *ServersDatabasesSecurityAlertPolicy) GetType() string {
	return "Microsoft.Sql/servers/databases/securityAlertPolicies"
}

// NewEmptyStatus returns a new empty (blank) status
func (policy *ServersDatabasesSecurityAlertPolicy) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &Servers_Databases_SecurityAlertPolicy_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (policy *ServersDatabasesSecurityAlertPolicy) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(policy.Spec)
	return policy.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (policy *ServersDatabasesSecurityAlertPolicy) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*Servers_Databases_SecurityAlertPolicy_STATUS); ok {
		policy.Status = *st
		return nil
	}

	// Convert status to required version
	var st Servers_Databases_SecurityAlertPolicy_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	policy.Status = st
	return nil
}

// Hub marks that this ServersDatabasesSecurityAlertPolicy is the hub type for conversion
func (policy *ServersDatabasesSecurityAlertPolicy) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (policy *ServersDatabasesSecurityAlertPolicy) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: policy.Spec.OriginalVersion,
		Kind:    "ServersDatabasesSecurityAlertPolicy",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20211101.ServersDatabasesSecurityAlertPolicy
// Generator information:
// - Generated from: /sql/resource-manager/Microsoft.Sql/stable/2021-11-01/DatabaseSecurityAlertPolicies.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/securityAlertPolicies/default
type ServersDatabasesSecurityAlertPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServersDatabasesSecurityAlertPolicy `json:"items"`
}

// Storage version of v1api20211101.Servers_Databases_SecurityAlertPolicy_Spec
type Servers_Databases_SecurityAlertPolicy_Spec struct {
	DisabledAlerts     []string `json:"disabledAlerts,omitempty"`
	EmailAccountAdmins *bool    `json:"emailAccountAdmins,omitempty"`
	EmailAddresses     []string `json:"emailAddresses,omitempty"`
	OriginalVersion    string   `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a sql.azure.com/ServersDatabase resource
	Owner                   *genruntime.KnownResourceReference `group:"sql.azure.com" json:"owner,omitempty" kind:"ServersDatabase"`
	PropertyBag             genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	RetentionDays           *int                               `json:"retentionDays,omitempty"`
	State                   *string                            `json:"state,omitempty"`
	StorageAccountAccessKey *genruntime.SecretReference        `json:"storageAccountAccessKey,omitempty"`
	StorageEndpoint         *string                            `json:"storageEndpoint,omitempty"`
}

var _ genruntime.ConvertibleSpec = &Servers_Databases_SecurityAlertPolicy_Spec{}

// ConvertSpecFrom populates our Servers_Databases_SecurityAlertPolicy_Spec from the provided source
func (policy *Servers_Databases_SecurityAlertPolicy_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == policy {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(policy)
}

// ConvertSpecTo populates the provided destination from our Servers_Databases_SecurityAlertPolicy_Spec
func (policy *Servers_Databases_SecurityAlertPolicy_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == policy {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(policy)
}

// Storage version of v1api20211101.Servers_Databases_SecurityAlertPolicy_STATUS
type Servers_Databases_SecurityAlertPolicy_STATUS struct {
	Conditions         []conditions.Condition `json:"conditions,omitempty"`
	CreationTime       *string                `json:"creationTime,omitempty"`
	DisabledAlerts     []string               `json:"disabledAlerts,omitempty"`
	EmailAccountAdmins *bool                  `json:"emailAccountAdmins,omitempty"`
	EmailAddresses     []string               `json:"emailAddresses,omitempty"`
	Id                 *string                `json:"id,omitempty"`
	Name               *string                `json:"name,omitempty"`
	PropertyBag        genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	RetentionDays      *int                   `json:"retentionDays,omitempty"`
	State              *string                `json:"state,omitempty"`
	StorageEndpoint    *string                `json:"storageEndpoint,omitempty"`
	SystemData         *SystemData_STATUS     `json:"systemData,omitempty"`
	Type               *string                `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &Servers_Databases_SecurityAlertPolicy_STATUS{}

// ConvertStatusFrom populates our Servers_Databases_SecurityAlertPolicy_STATUS from the provided source
func (policy *Servers_Databases_SecurityAlertPolicy_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == policy {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(policy)
}

// ConvertStatusTo populates the provided destination from our Servers_Databases_SecurityAlertPolicy_STATUS
func (policy *Servers_Databases_SecurityAlertPolicy_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == policy {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(policy)
}

func init() {
	SchemeBuilder.Register(&ServersDatabasesSecurityAlertPolicy{}, &ServersDatabasesSecurityAlertPolicyList{})
}
