// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20220901storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=search.azure.com,resources=searchservices,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=search.azure.com,resources={searchservices/status,searchservices/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20220901.SearchService
// Generator information:
// - Generated from: /search/resource-manager/Microsoft.Search/stable/2022-09-01/search.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Search/searchServices/{searchServiceName}
type SearchService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SearchService_Spec   `json:"spec,omitempty"`
	Status            SearchService_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &SearchService{}

// GetConditions returns the conditions of the resource
func (service *SearchService) GetConditions() conditions.Conditions {
	return service.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (service *SearchService) SetConditions(conditions conditions.Conditions) {
	service.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &SearchService{}

// AzureName returns the Azure name of the resource
func (service *SearchService) AzureName() string {
	return service.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2022-09-01"
func (service SearchService) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (service *SearchService) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (service *SearchService) GetSpec() genruntime.ConvertibleSpec {
	return &service.Spec
}

// GetStatus returns the status of this resource
func (service *SearchService) GetStatus() genruntime.ConvertibleStatus {
	return &service.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Search/searchServices"
func (service *SearchService) GetType() string {
	return "Microsoft.Search/searchServices"
}

// NewEmptyStatus returns a new empty (blank) status
func (service *SearchService) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &SearchService_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (service *SearchService) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(service.Spec)
	return service.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (service *SearchService) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*SearchService_STATUS); ok {
		service.Status = *st
		return nil
	}

	// Convert status to required version
	var st SearchService_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	service.Status = st
	return nil
}

// Hub marks that this SearchService is the hub type for conversion
func (service *SearchService) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (service *SearchService) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: service.Spec.OriginalVersion,
		Kind:    "SearchService",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20220901.SearchService
// Generator information:
// - Generated from: /search/resource-manager/Microsoft.Search/stable/2022-09-01/search.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Search/searchServices/{searchServiceName}
type SearchServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SearchService `json:"items"`
}

// Storage version of v1api20220901.APIVersion
// +kubebuilder:validation:Enum={"2022-09-01"}
type APIVersion string

const APIVersion_Value = APIVersion("2022-09-01")

// Storage version of v1api20220901.SearchService_Spec
type SearchService_Spec struct {
	AuthOptions *DataPlaneAuthOptions `json:"authOptions,omitempty"`

	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName         string                     `json:"azureName,omitempty"`
	DisableLocalAuth  *bool                      `json:"disableLocalAuth,omitempty"`
	EncryptionWithCmk *EncryptionWithCmk         `json:"encryptionWithCmk,omitempty"`
	HostingMode       *string                    `json:"hostingMode,omitempty"`
	Identity          *Identity                  `json:"identity,omitempty"`
	Location          *string                    `json:"location,omitempty"`
	NetworkRuleSet    *NetworkRuleSet            `json:"networkRuleSet,omitempty"`
	OperatorSpec      *SearchServiceOperatorSpec `json:"operatorSpec,omitempty"`
	OriginalVersion   string                     `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a resources.azure.com/ResourceGroup resource
	Owner               *genruntime.KnownResourceReference `group:"resources.azure.com" json:"owner,omitempty" kind:"ResourceGroup"`
	PartitionCount      *int                               `json:"partitionCount,omitempty"`
	PropertyBag         genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	PublicNetworkAccess *string                            `json:"publicNetworkAccess,omitempty"`
	ReplicaCount        *int                               `json:"replicaCount,omitempty"`
	Sku                 *Sku                               `json:"sku,omitempty"`
	Tags                map[string]string                  `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &SearchService_Spec{}

// ConvertSpecFrom populates our SearchService_Spec from the provided source
func (service *SearchService_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == service {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(service)
}

// ConvertSpecTo populates the provided destination from our SearchService_Spec
func (service *SearchService_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == service {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(service)
}

// Storage version of v1api20220901.SearchService_STATUS
// Describes an Azure Cognitive Search service and its current state.
type SearchService_STATUS struct {
	AuthOptions                *DataPlaneAuthOptions_STATUS       `json:"authOptions,omitempty"`
	Conditions                 []conditions.Condition             `json:"conditions,omitempty"`
	DisableLocalAuth           *bool                              `json:"disableLocalAuth,omitempty"`
	EncryptionWithCmk          *EncryptionWithCmk_STATUS          `json:"encryptionWithCmk,omitempty"`
	HostingMode                *string                            `json:"hostingMode,omitempty"`
	Id                         *string                            `json:"id,omitempty"`
	Identity                   *Identity_STATUS                   `json:"identity,omitempty"`
	Location                   *string                            `json:"location,omitempty"`
	Name                       *string                            `json:"name,omitempty"`
	NetworkRuleSet             *NetworkRuleSet_STATUS             `json:"networkRuleSet,omitempty"`
	PartitionCount             *int                               `json:"partitionCount,omitempty"`
	PrivateEndpointConnections []PrivateEndpointConnection_STATUS `json:"privateEndpointConnections,omitempty"`
	PropertyBag                genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	ProvisioningState          *string                            `json:"provisioningState,omitempty"`
	PublicNetworkAccess        *string                            `json:"publicNetworkAccess,omitempty"`
	ReplicaCount               *int                               `json:"replicaCount,omitempty"`
	SharedPrivateLinkResources []SharedPrivateLinkResource_STATUS `json:"sharedPrivateLinkResources,omitempty"`
	Sku                        *Sku_STATUS                        `json:"sku,omitempty"`
	Status                     *string                            `json:"status,omitempty"`
	StatusDetails              *string                            `json:"statusDetails,omitempty"`
	Tags                       map[string]string                  `json:"tags,omitempty"`
	Type                       *string                            `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &SearchService_STATUS{}

// ConvertStatusFrom populates our SearchService_STATUS from the provided source
func (service *SearchService_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == service {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(service)
}

// ConvertStatusTo populates the provided destination from our SearchService_STATUS
func (service *SearchService_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == service {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(service)
}

// Storage version of v1api20220901.DataPlaneAuthOptions
// Defines the options for how the data plane API of a Search service authenticates requests. This cannot be set if
// 'disableLocalAuth' is set to true.
type DataPlaneAuthOptions struct {
	AadOrApiKey *DataPlaneAadOrApiKeyAuthOption `json:"aadOrApiKey,omitempty"`
	PropertyBag genruntime.PropertyBag          `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20220901.DataPlaneAuthOptions_STATUS
// Defines the options for how the data plane API of a Search service authenticates requests. This cannot be set if
// 'disableLocalAuth' is set to true.
type DataPlaneAuthOptions_STATUS struct {
	AadOrApiKey *DataPlaneAadOrApiKeyAuthOption_STATUS `json:"aadOrApiKey,omitempty"`
	ApiKeyOnly  map[string]v1.JSON                     `json:"apiKeyOnly,omitempty"`
	PropertyBag genruntime.PropertyBag                 `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20220901.EncryptionWithCmk
// Describes a policy that determines how resources within the search service are to be encrypted with Customer Managed
// Keys.
type EncryptionWithCmk struct {
	Enforcement *string                `json:"enforcement,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20220901.EncryptionWithCmk_STATUS
// Describes a policy that determines how resources within the search service are to be encrypted with Customer Managed
// Keys.
type EncryptionWithCmk_STATUS struct {
	EncryptionComplianceStatus *string                `json:"encryptionComplianceStatus,omitempty"`
	Enforcement                *string                `json:"enforcement,omitempty"`
	PropertyBag                genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20220901.Identity
// Identity for the resource.
type Identity struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Type        *string                `json:"type,omitempty"`
}

// Storage version of v1api20220901.Identity_STATUS
// Identity for the resource.
type Identity_STATUS struct {
	PrincipalId *string                `json:"principalId,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	TenantId    *string                `json:"tenantId,omitempty"`
	Type        *string                `json:"type,omitempty"`
}

// Storage version of v1api20220901.NetworkRuleSet
// Network specific rules that determine how the Azure Cognitive Search service may be reached.
type NetworkRuleSet struct {
	IpRules     []IpRule               `json:"ipRules,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20220901.NetworkRuleSet_STATUS
// Network specific rules that determine how the Azure Cognitive Search service may be reached.
type NetworkRuleSet_STATUS struct {
	IpRules     []IpRule_STATUS        `json:"ipRules,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20220901.PrivateEndpointConnection_STATUS
// Describes an existing Private Endpoint connection to the Azure Cognitive Search service.
type PrivateEndpointConnection_STATUS struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20220901.SearchServiceOperatorSpec
// Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure
type SearchServiceOperatorSpec struct {
	PropertyBag genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	Secrets     *SearchServiceOperatorSecrets `json:"secrets,omitempty"`
}

// Storage version of v1api20220901.SharedPrivateLinkResource_STATUS
// Describes a Shared Private Link Resource managed by the Azure Cognitive Search service.
type SharedPrivateLinkResource_STATUS struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20220901.Sku
// Defines the SKU of an Azure Cognitive Search Service, which determines price tier and capacity limits.
type Sku struct {
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20220901.Sku_STATUS
// Defines the SKU of an Azure Cognitive Search Service, which determines price tier and capacity limits.
type Sku_STATUS struct {
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20220901.DataPlaneAadOrApiKeyAuthOption
// Indicates that either the API key or an access token from Azure Active Directory can be used for authentication.
type DataPlaneAadOrApiKeyAuthOption struct {
	AadAuthFailureMode *string                `json:"aadAuthFailureMode,omitempty"`
	PropertyBag        genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20220901.DataPlaneAadOrApiKeyAuthOption_STATUS
// Indicates that either the API key or an access token from Azure Active Directory can be used for authentication.
type DataPlaneAadOrApiKeyAuthOption_STATUS struct {
	AadAuthFailureMode *string                `json:"aadAuthFailureMode,omitempty"`
	PropertyBag        genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20220901.IpRule
// The IP restriction rule of the Azure Cognitive Search service.
type IpRule struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Value       *string                `json:"value,omitempty"`
}

// Storage version of v1api20220901.IpRule_STATUS
// The IP restriction rule of the Azure Cognitive Search service.
type IpRule_STATUS struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Value       *string                `json:"value,omitempty"`
}

// Storage version of v1api20220901.SearchServiceOperatorSecrets
type SearchServiceOperatorSecrets struct {
	AdminPrimaryKey   *genruntime.SecretDestination `json:"adminPrimaryKey,omitempty"`
	AdminSecondaryKey *genruntime.SecretDestination `json:"adminSecondaryKey,omitempty"`
	PropertyBag       genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	QueryKey          *genruntime.SecretDestination `json:"queryKey,omitempty"`
}

func init() {
	SchemeBuilder.Register(&SearchService{}, &SearchServiceList{})
}
