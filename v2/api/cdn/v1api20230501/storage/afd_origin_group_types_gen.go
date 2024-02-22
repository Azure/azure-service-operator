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

// +kubebuilder:rbac:groups=cdn.azure.com,resources=afdorigingroups,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=cdn.azure.com,resources={afdorigingroups/status,afdorigingroups/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20230501.AfdOriginGroup
// Generator information:
// - Generated from: /cdn/resource-manager/Microsoft.Cdn/stable/2023-05-01/afdx.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/originGroups/{originGroupName}
type AfdOriginGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Profiles_OriginGroup_Spec   `json:"spec,omitempty"`
	Status            Profiles_OriginGroup_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &AfdOriginGroup{}

// GetConditions returns the conditions of the resource
func (group *AfdOriginGroup) GetConditions() conditions.Conditions {
	return group.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (group *AfdOriginGroup) SetConditions(conditions conditions.Conditions) {
	group.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &AfdOriginGroup{}

// AzureName returns the Azure name of the resource
func (group *AfdOriginGroup) AzureName() string {
	return group.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2023-05-01"
func (group AfdOriginGroup) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (group *AfdOriginGroup) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (group *AfdOriginGroup) GetSpec() genruntime.ConvertibleSpec {
	return &group.Spec
}

// GetStatus returns the status of this resource
func (group *AfdOriginGroup) GetStatus() genruntime.ConvertibleStatus {
	return &group.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (group *AfdOriginGroup) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Cdn/profiles/originGroups"
func (group *AfdOriginGroup) GetType() string {
	return "Microsoft.Cdn/profiles/originGroups"
}

// NewEmptyStatus returns a new empty (blank) status
func (group *AfdOriginGroup) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &Profiles_OriginGroup_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (group *AfdOriginGroup) Owner() *genruntime.ResourceReference {
	ownerGroup, ownerKind := genruntime.LookupOwnerGroupKind(group.Spec)
	return group.Spec.Owner.AsResourceReference(ownerGroup, ownerKind)
}

// SetStatus sets the status of this resource
func (group *AfdOriginGroup) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*Profiles_OriginGroup_STATUS); ok {
		group.Status = *st
		return nil
	}

	// Convert status to required version
	var st Profiles_OriginGroup_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	group.Status = st
	return nil
}

// Hub marks that this AfdOriginGroup is the hub type for conversion
func (group *AfdOriginGroup) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (group *AfdOriginGroup) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: group.Spec.OriginalVersion,
		Kind:    "AfdOriginGroup",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20230501.AfdOriginGroup
// Generator information:
// - Generated from: /cdn/resource-manager/Microsoft.Cdn/stable/2023-05-01/afdx.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/originGroups/{originGroupName}
type AfdOriginGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AfdOriginGroup `json:"items"`
}

// Storage version of v1api20230501.Profiles_OriginGroup_Spec
type Profiles_OriginGroup_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName             string                           `json:"azureName,omitempty"`
	HealthProbeSettings   *HealthProbeParameters           `json:"healthProbeSettings,omitempty"`
	LoadBalancingSettings *LoadBalancingSettingsParameters `json:"loadBalancingSettings,omitempty"`
	OriginalVersion       string                           `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a cdn.azure.com/Profile resource
	Owner                                                 *genruntime.KnownResourceReference `group:"cdn.azure.com" json:"owner,omitempty" kind:"Profile"`
	PropertyBag                                           genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	SessionAffinityState                                  *string                            `json:"sessionAffinityState,omitempty"`
	TrafficRestorationTimeToHealedOrNewEndpointsInMinutes *int                               `json:"trafficRestorationTimeToHealedOrNewEndpointsInMinutes,omitempty"`
}

var _ genruntime.ConvertibleSpec = &Profiles_OriginGroup_Spec{}

// ConvertSpecFrom populates our Profiles_OriginGroup_Spec from the provided source
func (group *Profiles_OriginGroup_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == group {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(group)
}

// ConvertSpecTo populates the provided destination from our Profiles_OriginGroup_Spec
func (group *Profiles_OriginGroup_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == group {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(group)
}

// Storage version of v1api20230501.Profiles_OriginGroup_STATUS
type Profiles_OriginGroup_STATUS struct {
	Conditions                                            []conditions.Condition                  `json:"conditions,omitempty"`
	DeploymentStatus                                      *string                                 `json:"deploymentStatus,omitempty"`
	HealthProbeSettings                                   *HealthProbeParameters_STATUS           `json:"healthProbeSettings,omitempty"`
	Id                                                    *string                                 `json:"id,omitempty"`
	LoadBalancingSettings                                 *LoadBalancingSettingsParameters_STATUS `json:"loadBalancingSettings,omitempty"`
	Name                                                  *string                                 `json:"name,omitempty"`
	ProfileName                                           *string                                 `json:"profileName,omitempty"`
	PropertyBag                                           genruntime.PropertyBag                  `json:"$propertyBag,omitempty"`
	ProvisioningState                                     *string                                 `json:"provisioningState,omitempty"`
	SessionAffinityState                                  *string                                 `json:"sessionAffinityState,omitempty"`
	SystemData                                            *SystemData_STATUS                      `json:"systemData,omitempty"`
	TrafficRestorationTimeToHealedOrNewEndpointsInMinutes *int                                    `json:"trafficRestorationTimeToHealedOrNewEndpointsInMinutes,omitempty"`
	Type                                                  *string                                 `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &Profiles_OriginGroup_STATUS{}

// ConvertStatusFrom populates our Profiles_OriginGroup_STATUS from the provided source
func (group *Profiles_OriginGroup_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == group {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(group)
}

// ConvertStatusTo populates the provided destination from our Profiles_OriginGroup_STATUS
func (group *Profiles_OriginGroup_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == group {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(group)
}

// Storage version of v1api20230501.HealthProbeParameters
// The JSON object that contains the properties to send health probes to origin.
type HealthProbeParameters struct {
	ProbeIntervalInSeconds *int                   `json:"probeIntervalInSeconds,omitempty"`
	ProbePath              *string                `json:"probePath,omitempty"`
	ProbeProtocol          *string                `json:"probeProtocol,omitempty"`
	ProbeRequestType       *string                `json:"probeRequestType,omitempty"`
	PropertyBag            genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20230501.HealthProbeParameters_STATUS
// The JSON object that contains the properties to send health probes to origin.
type HealthProbeParameters_STATUS struct {
	ProbeIntervalInSeconds *int                   `json:"probeIntervalInSeconds,omitempty"`
	ProbePath              *string                `json:"probePath,omitempty"`
	ProbeProtocol          *string                `json:"probeProtocol,omitempty"`
	ProbeRequestType       *string                `json:"probeRequestType,omitempty"`
	PropertyBag            genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20230501.LoadBalancingSettingsParameters
// Round-Robin load balancing settings for a backend pool
type LoadBalancingSettingsParameters struct {
	AdditionalLatencyInMilliseconds *int                   `json:"additionalLatencyInMilliseconds,omitempty"`
	PropertyBag                     genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	SampleSize                      *int                   `json:"sampleSize,omitempty"`
	SuccessfulSamplesRequired       *int                   `json:"successfulSamplesRequired,omitempty"`
}

// Storage version of v1api20230501.LoadBalancingSettingsParameters_STATUS
// Round-Robin load balancing settings for a backend pool
type LoadBalancingSettingsParameters_STATUS struct {
	AdditionalLatencyInMilliseconds *int                   `json:"additionalLatencyInMilliseconds,omitempty"`
	PropertyBag                     genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	SampleSize                      *int                   `json:"sampleSize,omitempty"`
	SuccessfulSamplesRequired       *int                   `json:"successfulSamplesRequired,omitempty"`
}

func init() {
	SchemeBuilder.Register(&AfdOriginGroup{}, &AfdOriginGroupList{})
}
