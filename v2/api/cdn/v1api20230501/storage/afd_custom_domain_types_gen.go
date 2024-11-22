// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/configmaps"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/core"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/secrets"
	"github.com/rotisserie/eris"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=cdn.azure.com,resources=afdcustomdomains,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=cdn.azure.com,resources={afdcustomdomains/status,afdcustomdomains/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20230501.AfdCustomDomain
// Generator information:
// - Generated from: /cdn/resource-manager/Microsoft.Cdn/stable/2023-05-01/afdx.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/customDomains/{customDomainName}
type AfdCustomDomain struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AfdCustomDomain_Spec   `json:"spec,omitempty"`
	Status            AfdCustomDomain_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &AfdCustomDomain{}

// GetConditions returns the conditions of the resource
func (domain *AfdCustomDomain) GetConditions() conditions.Conditions {
	return domain.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (domain *AfdCustomDomain) SetConditions(conditions conditions.Conditions) {
	domain.Status.Conditions = conditions
}

var _ configmaps.Exporter = &AfdCustomDomain{}

// ConfigMapDestinationExpressions returns the Spec.OperatorSpec.ConfigMapExpressions property
func (domain *AfdCustomDomain) ConfigMapDestinationExpressions() []*core.DestinationExpression {
	if domain.Spec.OperatorSpec == nil {
		return nil
	}
	return domain.Spec.OperatorSpec.ConfigMapExpressions
}

var _ secrets.Exporter = &AfdCustomDomain{}

// SecretDestinationExpressions returns the Spec.OperatorSpec.SecretExpressions property
func (domain *AfdCustomDomain) SecretDestinationExpressions() []*core.DestinationExpression {
	if domain.Spec.OperatorSpec == nil {
		return nil
	}
	return domain.Spec.OperatorSpec.SecretExpressions
}

var _ genruntime.KubernetesResource = &AfdCustomDomain{}

// AzureName returns the Azure name of the resource
func (domain *AfdCustomDomain) AzureName() string {
	return domain.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2023-05-01"
func (domain AfdCustomDomain) GetAPIVersion() string {
	return "2023-05-01"
}

// GetResourceScope returns the scope of the resource
func (domain *AfdCustomDomain) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (domain *AfdCustomDomain) GetSpec() genruntime.ConvertibleSpec {
	return &domain.Spec
}

// GetStatus returns the status of this resource
func (domain *AfdCustomDomain) GetStatus() genruntime.ConvertibleStatus {
	return &domain.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (domain *AfdCustomDomain) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Cdn/profiles/customDomains"
func (domain *AfdCustomDomain) GetType() string {
	return "Microsoft.Cdn/profiles/customDomains"
}

// NewEmptyStatus returns a new empty (blank) status
func (domain *AfdCustomDomain) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &AfdCustomDomain_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (domain *AfdCustomDomain) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(domain.Spec)
	return domain.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (domain *AfdCustomDomain) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*AfdCustomDomain_STATUS); ok {
		domain.Status = *st
		return nil
	}

	// Convert status to required version
	var st AfdCustomDomain_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return eris.Wrap(err, "failed to convert status")
	}

	domain.Status = st
	return nil
}

// Hub marks that this AfdCustomDomain is the hub type for conversion
func (domain *AfdCustomDomain) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (domain *AfdCustomDomain) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: domain.Spec.OriginalVersion,
		Kind:    "AfdCustomDomain",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20230501.AfdCustomDomain
// Generator information:
// - Generated from: /cdn/resource-manager/Microsoft.Cdn/stable/2023-05-01/afdx.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/customDomains/{customDomainName}
type AfdCustomDomainList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AfdCustomDomain `json:"items"`
}

// Storage version of v1api20230501.AfdCustomDomain_Spec
type AfdCustomDomain_Spec struct {
	AzureDnsZone *ResourceReference `json:"azureDnsZone,omitempty"`

	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName          string                       `json:"azureName,omitempty"`
	ExtendedProperties map[string]string            `json:"extendedProperties,omitempty"`
	HostName           *string                      `json:"hostName,omitempty"`
	OperatorSpec       *AfdCustomDomainOperatorSpec `json:"operatorSpec,omitempty"`
	OriginalVersion    string                       `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a cdn.azure.com/Profile resource
	Owner                              *genruntime.KnownResourceReference `group:"cdn.azure.com" json:"owner,omitempty" kind:"Profile"`
	PreValidatedCustomDomainResourceId *ResourceReference                 `json:"preValidatedCustomDomainResourceId,omitempty"`
	PropertyBag                        genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	TlsSettings                        *AFDDomainHttpsParameters          `json:"tlsSettings,omitempty"`
}

var _ genruntime.ConvertibleSpec = &AfdCustomDomain_Spec{}

// ConvertSpecFrom populates our AfdCustomDomain_Spec from the provided source
func (domain *AfdCustomDomain_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == domain {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(domain)
}

// ConvertSpecTo populates the provided destination from our AfdCustomDomain_Spec
func (domain *AfdCustomDomain_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == domain {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(domain)
}

// Storage version of v1api20230501.AfdCustomDomain_STATUS
type AfdCustomDomain_STATUS struct {
	AzureDnsZone                       *ResourceReference_STATUS          `json:"azureDnsZone,omitempty"`
	Conditions                         []conditions.Condition             `json:"conditions,omitempty"`
	DeploymentStatus                   *string                            `json:"deploymentStatus,omitempty"`
	DomainValidationState              *string                            `json:"domainValidationState,omitempty"`
	ExtendedProperties                 map[string]string                  `json:"extendedProperties,omitempty"`
	HostName                           *string                            `json:"hostName,omitempty"`
	Id                                 *string                            `json:"id,omitempty"`
	Name                               *string                            `json:"name,omitempty"`
	PreValidatedCustomDomainResourceId *ResourceReference_STATUS          `json:"preValidatedCustomDomainResourceId,omitempty"`
	ProfileName                        *string                            `json:"profileName,omitempty"`
	PropertyBag                        genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	ProvisioningState                  *string                            `json:"provisioningState,omitempty"`
	SystemData                         *SystemData_STATUS                 `json:"systemData,omitempty"`
	TlsSettings                        *AFDDomainHttpsParameters_STATUS   `json:"tlsSettings,omitempty"`
	Type                               *string                            `json:"type,omitempty"`
	ValidationProperties               *DomainValidationProperties_STATUS `json:"validationProperties,omitempty"`
}

var _ genruntime.ConvertibleStatus = &AfdCustomDomain_STATUS{}

// ConvertStatusFrom populates our AfdCustomDomain_STATUS from the provided source
func (domain *AfdCustomDomain_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == domain {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(domain)
}

// ConvertStatusTo populates the provided destination from our AfdCustomDomain_STATUS
func (domain *AfdCustomDomain_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == domain {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(domain)
}

// Storage version of v1api20230501.APIVersion
// +kubebuilder:validation:Enum={"2023-05-01"}
type APIVersion string

const APIVersion_Value = APIVersion("2023-05-01")

// Storage version of v1api20230501.AfdCustomDomainOperatorSpec
// Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure
type AfdCustomDomainOperatorSpec struct {
	ConfigMapExpressions []*core.DestinationExpression `json:"configMapExpressions,omitempty"`
	PropertyBag          genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	SecretExpressions    []*core.DestinationExpression `json:"secretExpressions,omitempty"`
}

// Storage version of v1api20230501.AFDDomainHttpsParameters
// The JSON object that contains the properties to secure a domain.
type AFDDomainHttpsParameters struct {
	CertificateType   *string                `json:"certificateType,omitempty"`
	MinimumTlsVersion *string                `json:"minimumTlsVersion,omitempty"`
	PropertyBag       genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Secret            *ResourceReference     `json:"secret,omitempty"`
}

// Storage version of v1api20230501.AFDDomainHttpsParameters_STATUS
// The JSON object that contains the properties to secure a domain.
type AFDDomainHttpsParameters_STATUS struct {
	CertificateType   *string                   `json:"certificateType,omitempty"`
	MinimumTlsVersion *string                   `json:"minimumTlsVersion,omitempty"`
	PropertyBag       genruntime.PropertyBag    `json:"$propertyBag,omitempty"`
	Secret            *ResourceReference_STATUS `json:"secret,omitempty"`
}

// Storage version of v1api20230501.DomainValidationProperties_STATUS
// The JSON object that contains the properties to validate a domain.
type DomainValidationProperties_STATUS struct {
	ExpirationDate  *string                `json:"expirationDate,omitempty"`
	PropertyBag     genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ValidationToken *string                `json:"validationToken,omitempty"`
}

// Storage version of v1api20230501.ResourceReference
// Reference to another resource.
type ResourceReference struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`

	// Reference: Resource ID.
	Reference *genruntime.ResourceReference `armReference:"Id" json:"reference,omitempty"`
}

// Storage version of v1api20230501.ResourceReference_STATUS
// Reference to another resource.
type ResourceReference_STATUS struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20230501.SystemData_STATUS
// Read only system data
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
	SchemeBuilder.Register(&AfdCustomDomain{}, &AfdCustomDomainList{})
}
