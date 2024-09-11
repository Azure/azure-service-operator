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

// +kubebuilder:rbac:groups=cdn.azure.com,resources=afdendpoints,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=cdn.azure.com,resources={afdendpoints/status,afdendpoints/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20230501.AfdEndpoint
// Generator information:
// - Generated from: /cdn/resource-manager/Microsoft.Cdn/stable/2023-05-01/afdx.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/afdEndpoints/{endpointName}
type AfdEndpoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Profiles_AfdEndpoint_Spec   `json:"spec,omitempty"`
	Status            Profiles_AfdEndpoint_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &AfdEndpoint{}

// GetConditions returns the conditions of the resource
func (endpoint *AfdEndpoint) GetConditions() conditions.Conditions {
	return endpoint.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (endpoint *AfdEndpoint) SetConditions(conditions conditions.Conditions) {
	endpoint.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &AfdEndpoint{}

// AzureName returns the Azure name of the resource
func (endpoint *AfdEndpoint) AzureName() string {
	return endpoint.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2023-05-01"
func (endpoint AfdEndpoint) GetAPIVersion() string {
	return "2023-05-01"
}

// GetResourceScope returns the scope of the resource
func (endpoint *AfdEndpoint) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (endpoint *AfdEndpoint) GetSpec() genruntime.ConvertibleSpec {
	return &endpoint.Spec
}

// GetStatus returns the status of this resource
func (endpoint *AfdEndpoint) GetStatus() genruntime.ConvertibleStatus {
	return &endpoint.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (endpoint *AfdEndpoint) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Cdn/profiles/afdEndpoints"
func (endpoint *AfdEndpoint) GetType() string {
	return "Microsoft.Cdn/profiles/afdEndpoints"
}

// NewEmptyStatus returns a new empty (blank) status
func (endpoint *AfdEndpoint) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &Profiles_AfdEndpoint_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (endpoint *AfdEndpoint) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(endpoint.Spec)
	return endpoint.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (endpoint *AfdEndpoint) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*Profiles_AfdEndpoint_STATUS); ok {
		endpoint.Status = *st
		return nil
	}

	// Convert status to required version
	var st Profiles_AfdEndpoint_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	endpoint.Status = st
	return nil
}

// Hub marks that this AfdEndpoint is the hub type for conversion
func (endpoint *AfdEndpoint) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (endpoint *AfdEndpoint) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: endpoint.Spec.OriginalVersion,
		Kind:    "AfdEndpoint",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20230501.AfdEndpoint
// Generator information:
// - Generated from: /cdn/resource-manager/Microsoft.Cdn/stable/2023-05-01/afdx.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/afdEndpoints/{endpointName}
type AfdEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AfdEndpoint `json:"items"`
}

// Storage version of v1api20230501.Profiles_AfdEndpoint_Spec
type Profiles_AfdEndpoint_Spec struct {
	AutoGeneratedDomainNameLabelScope *string `json:"autoGeneratedDomainNameLabelScope,omitempty"`

	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName       string  `json:"azureName,omitempty"`
	EnabledState    *string `json:"enabledState,omitempty"`
	Location        *string `json:"location,omitempty"`
	OriginalVersion string  `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a cdn.azure.com/Profile resource
	Owner       *genruntime.KnownResourceReference `group:"cdn.azure.com" json:"owner,omitempty" kind:"Profile"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	Tags        map[string]string                  `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &Profiles_AfdEndpoint_Spec{}

// ConvertSpecFrom populates our Profiles_AfdEndpoint_Spec from the provided source
func (endpoint *Profiles_AfdEndpoint_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == endpoint {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(endpoint)
}

// ConvertSpecTo populates the provided destination from our Profiles_AfdEndpoint_Spec
func (endpoint *Profiles_AfdEndpoint_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == endpoint {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(endpoint)
}

// Storage version of v1api20230501.Profiles_AfdEndpoint_STATUS
type Profiles_AfdEndpoint_STATUS struct {
	AutoGeneratedDomainNameLabelScope *string                `json:"autoGeneratedDomainNameLabelScope,omitempty"`
	Conditions                        []conditions.Condition `json:"conditions,omitempty"`
	DeploymentStatus                  *string                `json:"deploymentStatus,omitempty"`
	EnabledState                      *string                `json:"enabledState,omitempty"`
	HostName                          *string                `json:"hostName,omitempty"`
	Id                                *string                `json:"id,omitempty"`
	Location                          *string                `json:"location,omitempty"`
	Name                              *string                `json:"name,omitempty"`
	ProfileName                       *string                `json:"profileName,omitempty"`
	PropertyBag                       genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ProvisioningState                 *string                `json:"provisioningState,omitempty"`
	SystemData                        *SystemData_STATUS     `json:"systemData,omitempty"`
	Tags                              map[string]string      `json:"tags,omitempty"`
	Type                              *string                `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &Profiles_AfdEndpoint_STATUS{}

// ConvertStatusFrom populates our Profiles_AfdEndpoint_STATUS from the provided source
func (endpoint *Profiles_AfdEndpoint_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == endpoint {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(endpoint)
}

// ConvertStatusTo populates the provided destination from our Profiles_AfdEndpoint_STATUS
func (endpoint *Profiles_AfdEndpoint_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == endpoint {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(endpoint)
}

func init() {
	SchemeBuilder.Register(&AfdEndpoint{}, &AfdEndpointList{})
}
