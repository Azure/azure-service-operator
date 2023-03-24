// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20211101storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=sql.azure.com,resources=serversazureadonlyauthentications,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=sql.azure.com,resources={serversazureadonlyauthentications/status,serversazureadonlyauthentications/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1beta20211101.ServersAzureADOnlyAuthentication
// Generator information:
// - Generated from: /sql/resource-manager/Microsoft.Sql/stable/2021-11-01/ServerAzureADOnlyAuthentications.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/azureADOnlyAuthentications/Default
type ServersAzureADOnlyAuthentication struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Servers_AzureADOnlyAuthentication_Spec   `json:"spec,omitempty"`
	Status            Servers_AzureADOnlyAuthentication_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &ServersAzureADOnlyAuthentication{}

// GetConditions returns the conditions of the resource
func (authentication *ServersAzureADOnlyAuthentication) GetConditions() conditions.Conditions {
	return authentication.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (authentication *ServersAzureADOnlyAuthentication) SetConditions(conditions conditions.Conditions) {
	authentication.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &ServersAzureADOnlyAuthentication{}

// AzureName returns the Azure name of the resource (always "Default")
func (authentication *ServersAzureADOnlyAuthentication) AzureName() string {
	return "Default"
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-11-01"
func (authentication ServersAzureADOnlyAuthentication) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (authentication *ServersAzureADOnlyAuthentication) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (authentication *ServersAzureADOnlyAuthentication) GetSpec() genruntime.ConvertibleSpec {
	return &authentication.Spec
}

// GetStatus returns the status of this resource
func (authentication *ServersAzureADOnlyAuthentication) GetStatus() genruntime.ConvertibleStatus {
	return &authentication.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Sql/servers/azureADOnlyAuthentications"
func (authentication *ServersAzureADOnlyAuthentication) GetType() string {
	return "Microsoft.Sql/servers/azureADOnlyAuthentications"
}

// NewEmptyStatus returns a new empty (blank) status
func (authentication *ServersAzureADOnlyAuthentication) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &Servers_AzureADOnlyAuthentication_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (authentication *ServersAzureADOnlyAuthentication) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(authentication.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  authentication.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (authentication *ServersAzureADOnlyAuthentication) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*Servers_AzureADOnlyAuthentication_STATUS); ok {
		authentication.Status = *st
		return nil
	}

	// Convert status to required version
	var st Servers_AzureADOnlyAuthentication_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	authentication.Status = st
	return nil
}

// Hub marks that this ServersAzureADOnlyAuthentication is the hub type for conversion
func (authentication *ServersAzureADOnlyAuthentication) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (authentication *ServersAzureADOnlyAuthentication) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: authentication.Spec.OriginalVersion,
		Kind:    "ServersAzureADOnlyAuthentication",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1beta20211101.ServersAzureADOnlyAuthentication
// Generator information:
// - Generated from: /sql/resource-manager/Microsoft.Sql/stable/2021-11-01/ServerAzureADOnlyAuthentications.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/azureADOnlyAuthentications/Default
type ServersAzureADOnlyAuthenticationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServersAzureADOnlyAuthentication `json:"items"`
}

// Storage version of v1beta20211101.Servers_AzureADOnlyAuthentication_Spec
type Servers_AzureADOnlyAuthentication_Spec struct {
	AzureADOnlyAuthentication *bool  `json:"azureADOnlyAuthentication,omitempty"`
	OriginalVersion           string `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a sql.azure.com/Server resource
	Owner       *genruntime.KnownResourceReference `group:"sql.azure.com" json:"owner,omitempty" kind:"Server"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
}

var _ genruntime.ConvertibleSpec = &Servers_AzureADOnlyAuthentication_Spec{}

// ConvertSpecFrom populates our Servers_AzureADOnlyAuthentication_Spec from the provided source
func (authentication *Servers_AzureADOnlyAuthentication_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == authentication {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(authentication)
}

// ConvertSpecTo populates the provided destination from our Servers_AzureADOnlyAuthentication_Spec
func (authentication *Servers_AzureADOnlyAuthentication_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == authentication {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(authentication)
}

// Storage version of v1beta20211101.Servers_AzureADOnlyAuthentication_STATUS
type Servers_AzureADOnlyAuthentication_STATUS struct {
	AzureADOnlyAuthentication *bool                  `json:"azureADOnlyAuthentication,omitempty"`
	Conditions                []conditions.Condition `json:"conditions,omitempty"`
	Id                        *string                `json:"id,omitempty"`
	Name                      *string                `json:"name,omitempty"`
	PropertyBag               genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Type                      *string                `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &Servers_AzureADOnlyAuthentication_STATUS{}

// ConvertStatusFrom populates our Servers_AzureADOnlyAuthentication_STATUS from the provided source
func (authentication *Servers_AzureADOnlyAuthentication_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == authentication {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(authentication)
}

// ConvertStatusTo populates the provided destination from our Servers_AzureADOnlyAuthentication_STATUS
func (authentication *Servers_AzureADOnlyAuthentication_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == authentication {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(authentication)
}

func init() {
	SchemeBuilder.Register(&ServersAzureADOnlyAuthentication{}, &ServersAzureADOnlyAuthenticationList{})
}
