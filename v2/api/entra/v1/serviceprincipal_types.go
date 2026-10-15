// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package v1

import (
	"github.com/microsoftgraph/msgraph-beta-sdk-go/models"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/conversion"

	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
)

// +kubebuilder:rbac:groups=entra.azure.com,resources=serviceprincipals,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=entra.azure.com,resources={serviceprincipals/status,serviceprincipals/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:resource:categories={azure,entra}
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// +kubebuilder:storageversion
// ServicePrincipal is an Entra service principal for an application.
type ServicePrincipal struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServicePrincipalSpec   `json:"spec,omitempty"`
	Status            ServicePrincipalStatus `json:"status,omitempty"`
}

var (
	_ conditions.Conditioner = &ServicePrincipal{}
	_ conversion.Hub         = &ServicePrincipal{}
)

func (sp *ServicePrincipal) GetConditions() conditions.Conditions {
	return sp.Status.Conditions
}

func (sp *ServicePrincipal) SetConditions(value conditions.Conditions) {
	sp.Status.Conditions = value
}

func (sp *ServicePrincipal) Hub() {}

// +kubebuilder:object:root=true
type ServicePrincipalList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServicePrincipal `json:"items"`
}

// +kubebuilder:validation:XValidation:rule="has(self.appId) || has(self.displayName)",message="appId or displayName must be specified"
type ServicePrincipalSpec struct {
	// AppId is the global application (client) GUID, not the tenant-specific Entra object ID.
	// Required to create a new service principal; existing principals can be adopted by displayName.
	// +kubebuilder:validation:Pattern="^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$"
	AppId *string `json:"appId,omitempty"`

	// DisplayName identifies an existing principal for adoption, following Application's
	// behavior: multiple matches are an error. An appId is also required if creation is needed.
	// +kubebuilder:validation:MinLength=1
	DisplayName *string `json:"displayName,omitempty"`

	// OperatorSpec configures adoption and destinations for the resolved IDs.
	OperatorSpec *ServicePrincipalOperatorSpec `json:"operatorSpec,omitempty"`
}

func (spec *ServicePrincipalSpec) OriginalVersion() string {
	return GroupVersion.Version
}

func (spec *ServicePrincipalSpec) AssignToServicePrincipal(model models.ServicePrincipalable) {
	model.SetAppId(spec.AppId)
	if spec.DisplayName != nil {
		model.SetDisplayName(spec.DisplayName)
	}
}

type ServicePrincipalStatus struct {
	// EntraID is the tenant-specific object ID, suitable for role assignments.
	EntraID *string `json:"entraID,omitempty"`
	// AppId is the global application (client) GUID.
	AppId       *string                `json:"appId,omitempty"`
	DisplayName *string                `json:"displayName,omitempty"`
	Conditions  []conditions.Condition `json:"conditions,omitempty"`
}

func (status *ServicePrincipalStatus) AssignFromServicePrincipal(model models.ServicePrincipalable) {
	if model == nil {
		return
	}
	status.EntraID = model.GetId()
	status.AppId = model.GetAppId()
	status.DisplayName = model.GetDisplayName()
}

type ServicePrincipalOperatorSpec struct {
	// CreationMode defaults to AdoptOrCreate. Set AdoptOnly to resolve an existing
	// service principal without creating one if it is absent in this tenant.
	CreationMode *CreationMode                       `json:"creationMode,omitempty"`
	ConfigMaps   *ServicePrincipalOperatorConfigMaps `json:"configmaps,omitempty"`
}

func (spec *ServicePrincipalOperatorSpec) CreationAllowed() bool {
	return spec.CreationMode == nil || spec.CreationMode.AllowsCreation()
}

func (spec *ServicePrincipalOperatorSpec) AdoptionAllowed() bool {
	return spec.CreationMode == nil || spec.CreationMode.AllowsAdoption()
}

type ServicePrincipalOperatorConfigMaps struct {
	EntraID *genruntime.ConfigMapDestination `json:"entraID,omitempty"`
	AppId   *genruntime.ConfigMapDestination `json:"appId,omitempty"`
}

func init() {
	SchemeBuilder.Register(&ServicePrincipal{}, &ServicePrincipalList{})
}
