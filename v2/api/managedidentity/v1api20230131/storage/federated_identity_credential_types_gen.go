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

// +kubebuilder:rbac:groups=managedidentity.azure.com,resources=federatedidentitycredentials,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=managedidentity.azure.com,resources={federatedidentitycredentials/status,federatedidentitycredentials/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20230131.FederatedIdentityCredential
// Generator information:
// - Generated from: /msi/resource-manager/Microsoft.ManagedIdentity/stable/2023-01-31/ManagedIdentity.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{resourceName}/federatedIdentityCredentials/{federatedIdentityCredentialResourceName}
type FederatedIdentityCredential struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FederatedIdentityCredential_Spec   `json:"spec,omitempty"`
	Status            FederatedIdentityCredential_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &FederatedIdentityCredential{}

// GetConditions returns the conditions of the resource
func (credential *FederatedIdentityCredential) GetConditions() conditions.Conditions {
	return credential.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (credential *FederatedIdentityCredential) SetConditions(conditions conditions.Conditions) {
	credential.Status.Conditions = conditions
}

var _ configmaps.Exporter = &FederatedIdentityCredential{}

// ConfigMapDestinationExpressions returns the Spec.OperatorSpec.ConfigMapExpressions property
func (credential *FederatedIdentityCredential) ConfigMapDestinationExpressions() []*core.DestinationExpression {
	if credential.Spec.OperatorSpec == nil {
		return nil
	}
	return credential.Spec.OperatorSpec.ConfigMapExpressions
}

var _ secrets.Exporter = &FederatedIdentityCredential{}

// SecretDestinationExpressions returns the Spec.OperatorSpec.SecretExpressions property
func (credential *FederatedIdentityCredential) SecretDestinationExpressions() []*core.DestinationExpression {
	if credential.Spec.OperatorSpec == nil {
		return nil
	}
	return credential.Spec.OperatorSpec.SecretExpressions
}

var _ genruntime.KubernetesResource = &FederatedIdentityCredential{}

// AzureName returns the Azure name of the resource
func (credential *FederatedIdentityCredential) AzureName() string {
	return credential.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2023-01-31"
func (credential FederatedIdentityCredential) GetAPIVersion() string {
	return "2023-01-31"
}

// GetResourceScope returns the scope of the resource
func (credential *FederatedIdentityCredential) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (credential *FederatedIdentityCredential) GetSpec() genruntime.ConvertibleSpec {
	return &credential.Spec
}

// GetStatus returns the status of this resource
func (credential *FederatedIdentityCredential) GetStatus() genruntime.ConvertibleStatus {
	return &credential.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (credential *FederatedIdentityCredential) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.ManagedIdentity/userAssignedIdentities/federatedIdentityCredentials"
func (credential *FederatedIdentityCredential) GetType() string {
	return "Microsoft.ManagedIdentity/userAssignedIdentities/federatedIdentityCredentials"
}

// NewEmptyStatus returns a new empty (blank) status
func (credential *FederatedIdentityCredential) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &FederatedIdentityCredential_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (credential *FederatedIdentityCredential) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(credential.Spec)
	return credential.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (credential *FederatedIdentityCredential) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*FederatedIdentityCredential_STATUS); ok {
		credential.Status = *st
		return nil
	}

	// Convert status to required version
	var st FederatedIdentityCredential_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return eris.Wrap(err, "failed to convert status")
	}

	credential.Status = st
	return nil
}

// Hub marks that this FederatedIdentityCredential is the hub type for conversion
func (credential *FederatedIdentityCredential) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (credential *FederatedIdentityCredential) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: credential.Spec.OriginalVersion,
		Kind:    "FederatedIdentityCredential",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20230131.FederatedIdentityCredential
// Generator information:
// - Generated from: /msi/resource-manager/Microsoft.ManagedIdentity/stable/2023-01-31/ManagedIdentity.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{resourceName}/federatedIdentityCredentials/{federatedIdentityCredentialResourceName}
type FederatedIdentityCredentialList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FederatedIdentityCredential `json:"items"`
}

// Storage version of v1api20230131.APIVersion
// +kubebuilder:validation:Enum={"2023-01-31"}
type APIVersion string

const APIVersion_Value = APIVersion("2023-01-31")

// Storage version of v1api20230131.FederatedIdentityCredential_Spec
type FederatedIdentityCredential_Spec struct {
	Audiences []string `json:"audiences,omitempty"`

	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName        string                                   `json:"azureName,omitempty"`
	Issuer           *string                                  `json:"issuer,omitempty" optionalConfigMapPair:"Issuer"`
	IssuerFromConfig *genruntime.ConfigMapReference           `json:"issuerFromConfig,omitempty" optionalConfigMapPair:"Issuer"`
	OperatorSpec     *FederatedIdentityCredentialOperatorSpec `json:"operatorSpec,omitempty"`
	OriginalVersion  string                                   `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a managedidentity.azure.com/UserAssignedIdentity resource
	Owner             *genruntime.KnownResourceReference `group:"managedidentity.azure.com" json:"owner,omitempty" kind:"UserAssignedIdentity"`
	PropertyBag       genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	Subject           *string                            `json:"subject,omitempty" optionalConfigMapPair:"Subject"`
	SubjectFromConfig *genruntime.ConfigMapReference     `json:"subjectFromConfig,omitempty" optionalConfigMapPair:"Subject"`
}

var _ genruntime.ConvertibleSpec = &FederatedIdentityCredential_Spec{}

// ConvertSpecFrom populates our FederatedIdentityCredential_Spec from the provided source
func (credential *FederatedIdentityCredential_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == credential {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(credential)
}

// ConvertSpecTo populates the provided destination from our FederatedIdentityCredential_Spec
func (credential *FederatedIdentityCredential_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == credential {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(credential)
}

// Storage version of v1api20230131.FederatedIdentityCredential_STATUS
type FederatedIdentityCredential_STATUS struct {
	Audiences   []string               `json:"audiences,omitempty"`
	Conditions  []conditions.Condition `json:"conditions,omitempty"`
	Id          *string                `json:"id,omitempty"`
	Issuer      *string                `json:"issuer,omitempty"`
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Subject     *string                `json:"subject,omitempty"`
	SystemData  *SystemData_STATUS     `json:"systemData,omitempty"`
	Type        *string                `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &FederatedIdentityCredential_STATUS{}

// ConvertStatusFrom populates our FederatedIdentityCredential_STATUS from the provided source
func (credential *FederatedIdentityCredential_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == credential {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(credential)
}

// ConvertStatusTo populates the provided destination from our FederatedIdentityCredential_STATUS
func (credential *FederatedIdentityCredential_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == credential {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(credential)
}

// Storage version of v1api20230131.FederatedIdentityCredentialOperatorSpec
// Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure
type FederatedIdentityCredentialOperatorSpec struct {
	ConfigMapExpressions []*core.DestinationExpression `json:"configMapExpressions,omitempty"`
	PropertyBag          genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	SecretExpressions    []*core.DestinationExpression `json:"secretExpressions,omitempty"`
}

// Storage version of v1api20230131.SystemData_STATUS
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
	SchemeBuilder.Register(&FederatedIdentityCredential{}, &FederatedIdentityCredentialList{})
}
