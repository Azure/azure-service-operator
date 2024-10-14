// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package storage

import (
	"context"
	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/configmaps"
	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// +kubebuilder:rbac:groups=machinelearningservices.azure.com,resources=registries,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=machinelearningservices.azure.com,resources={registries/status,registries/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20240401.Registry
// Generator information:
// - Generated from: /machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/registries.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/registries/{registryName}
type Registry struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Registry_Spec                  `json:"spec,omitempty"`
	Status            RegistryTrackedResource_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &Registry{}

// GetConditions returns the conditions of the resource
func (registry *Registry) GetConditions() conditions.Conditions {
	return registry.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (registry *Registry) SetConditions(conditions conditions.Conditions) {
	registry.Status.Conditions = conditions
}

var _ genruntime.KubernetesExporter = &Registry{}

// ExportKubernetesResources defines a resource which can create other resources in Kubernetes.
func (registry *Registry) ExportKubernetesResources(_ context.Context, _ genruntime.MetaObject, _ *genericarmclient.GenericClient, _ logr.Logger) ([]client.Object, error) {
	collector := configmaps.NewCollector(registry.Namespace)
	if registry.Spec.OperatorSpec != nil && registry.Spec.OperatorSpec.ConfigMaps != nil {
		if registry.Status.DiscoveryUrl != nil {
			collector.AddValue(registry.Spec.OperatorSpec.ConfigMaps.DiscoveryUrl, *registry.Status.DiscoveryUrl)
		}
	}
	if registry.Spec.OperatorSpec != nil && registry.Spec.OperatorSpec.ConfigMaps != nil {
		if registry.Status.MlFlowRegistryUri != nil {
			collector.AddValue(registry.Spec.OperatorSpec.ConfigMaps.MlFlowRegistryUri, *registry.Status.MlFlowRegistryUri)
		}
	}
	result, err := collector.Values()
	if err != nil {
		return nil, err
	}
	return configmaps.SliceToClientObjectSlice(result), nil
}

var _ genruntime.KubernetesResource = &Registry{}

// AzureName returns the Azure name of the resource
func (registry *Registry) AzureName() string {
	return registry.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2024-04-01"
func (registry Registry) GetAPIVersion() string {
	return "2024-04-01"
}

// GetResourceScope returns the scope of the resource
func (registry *Registry) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (registry *Registry) GetSpec() genruntime.ConvertibleSpec {
	return &registry.Spec
}

// GetStatus returns the status of this resource
func (registry *Registry) GetStatus() genruntime.ConvertibleStatus {
	return &registry.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (registry *Registry) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.MachineLearningServices/registries"
func (registry *Registry) GetType() string {
	return "Microsoft.MachineLearningServices/registries"
}

// NewEmptyStatus returns a new empty (blank) status
func (registry *Registry) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &RegistryTrackedResource_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (registry *Registry) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(registry.Spec)
	return registry.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (registry *Registry) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*RegistryTrackedResource_STATUS); ok {
		registry.Status = *st
		return nil
	}

	// Convert status to required version
	var st RegistryTrackedResource_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	registry.Status = st
	return nil
}

// Hub marks that this Registry is the hub type for conversion
func (registry *Registry) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (registry *Registry) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: registry.Spec.OriginalVersion,
		Kind:    "Registry",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20240401.Registry
// Generator information:
// - Generated from: /machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/registries.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/registries/{registryName}
type RegistryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Registry `json:"items"`
}

// Storage version of v1api20240401.APIVersion
// +kubebuilder:validation:Enum={"2024-04-01"}
type APIVersion string

const APIVersion_Value = APIVersion("2024-04-01")

// Storage version of v1api20240401.Registry_Spec
type Registry_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName                     string                  `json:"azureName,omitempty"`
	DiscoveryUrl                  *string                 `json:"discoveryUrl,omitempty"`
	Identity                      *ManagedServiceIdentity `json:"identity,omitempty"`
	IntellectualPropertyPublisher *string                 `json:"intellectualPropertyPublisher,omitempty"`
	Kind                          *string                 `json:"kind,omitempty"`
	Location                      *string                 `json:"location,omitempty"`
	ManagedResourceGroup          *ArmResourceId          `json:"managedResourceGroup,omitempty"`
	MlFlowRegistryUri             *string                 `json:"mlFlowRegistryUri,omitempty"`
	OperatorSpec                  *RegistryOperatorSpec   `json:"operatorSpec,omitempty"`
	OriginalVersion               string                  `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a resources.azure.com/ResourceGroup resource
	Owner                              *genruntime.KnownResourceReference  `group:"resources.azure.com" json:"owner,omitempty" kind:"ResourceGroup"`
	PropertyBag                        genruntime.PropertyBag              `json:"$propertyBag,omitempty"`
	PublicNetworkAccess                *string                             `json:"publicNetworkAccess,omitempty"`
	RegionDetails                      []RegistryRegionArmDetails          `json:"regionDetails,omitempty"`
	RegistryPrivateEndpointConnections []RegistryPrivateEndpointConnection `json:"registryPrivateEndpointConnections,omitempty"`
	Sku                                *Sku                                `json:"sku,omitempty"`
	Tags                               map[string]string                   `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &Registry_Spec{}

// ConvertSpecFrom populates our Registry_Spec from the provided source
func (registry *Registry_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == registry {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(registry)
}

// ConvertSpecTo populates the provided destination from our Registry_Spec
func (registry *Registry_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == registry {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(registry)
}

// Storage version of v1api20240401.RegistryTrackedResource_STATUS
type RegistryTrackedResource_STATUS struct {
	Conditions                         []conditions.Condition                     `json:"conditions,omitempty"`
	DiscoveryUrl                       *string                                    `json:"discoveryUrl,omitempty"`
	Id                                 *string                                    `json:"id,omitempty"`
	Identity                           *ManagedServiceIdentity_STATUS             `json:"identity,omitempty"`
	IntellectualPropertyPublisher      *string                                    `json:"intellectualPropertyPublisher,omitempty"`
	Kind                               *string                                    `json:"kind,omitempty"`
	Location                           *string                                    `json:"location,omitempty"`
	ManagedResourceGroup               *ArmResourceId_STATUS                      `json:"managedResourceGroup,omitempty"`
	MlFlowRegistryUri                  *string                                    `json:"mlFlowRegistryUri,omitempty"`
	Name                               *string                                    `json:"name,omitempty"`
	PropertyBag                        genruntime.PropertyBag                     `json:"$propertyBag,omitempty"`
	PublicNetworkAccess                *string                                    `json:"publicNetworkAccess,omitempty"`
	RegionDetails                      []RegistryRegionArmDetails_STATUS          `json:"regionDetails,omitempty"`
	RegistryPrivateEndpointConnections []RegistryPrivateEndpointConnection_STATUS `json:"registryPrivateEndpointConnections,omitempty"`
	Sku                                *Sku_STATUS                                `json:"sku,omitempty"`
	SystemData                         *SystemData_STATUS                         `json:"systemData,omitempty"`
	Tags                               map[string]string                          `json:"tags,omitempty"`
	Type                               *string                                    `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &RegistryTrackedResource_STATUS{}

// ConvertStatusFrom populates our RegistryTrackedResource_STATUS from the provided source
func (resource *RegistryTrackedResource_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == resource {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(resource)
}

// ConvertStatusTo populates the provided destination from our RegistryTrackedResource_STATUS
func (resource *RegistryTrackedResource_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == resource {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(resource)
}

// Storage version of v1api20240401.ArmResourceId
// ARM ResourceId of a resource
type ArmResourceId struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`

	// ResourceReference: Arm ResourceId is in the format
	// "/subscriptions/{SubscriptionId}/resourceGroups/{ResourceGroupName}/providers/Microsoft.Storage/storageAccounts/{StorageAccountName}"
	// or
	// "/subscriptions/{SubscriptionId}/resourceGroups/{ResourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{AcrName}"
	ResourceReference *genruntime.ResourceReference `armReference:"ResourceId" json:"resourceReference,omitempty"`
}

// Storage version of v1api20240401.ArmResourceId_STATUS
// ARM ResourceId of a resource
type ArmResourceId_STATUS struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ResourceId  *string                `json:"resourceId,omitempty"`
}

// Storage version of v1api20240401.ManagedServiceIdentity
// Managed service identity (system assigned and/or user assigned identities)
type ManagedServiceIdentity struct {
	PropertyBag            genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	Type                   *string                       `json:"type,omitempty"`
	UserAssignedIdentities []UserAssignedIdentityDetails `json:"userAssignedIdentities,omitempty"`
}

// Storage version of v1api20240401.ManagedServiceIdentity_STATUS
// Managed service identity (system assigned and/or user assigned identities)
type ManagedServiceIdentity_STATUS struct {
	PrincipalId            *string                                `json:"principalId,omitempty"`
	PropertyBag            genruntime.PropertyBag                 `json:"$propertyBag,omitempty"`
	TenantId               *string                                `json:"tenantId,omitempty"`
	Type                   *string                                `json:"type,omitempty"`
	UserAssignedIdentities map[string]UserAssignedIdentity_STATUS `json:"userAssignedIdentities,omitempty"`
}

// Storage version of v1api20240401.RegistryOperatorSpec
// Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure
type RegistryOperatorSpec struct {
	ConfigMaps  *RegistryOperatorConfigMaps `json:"configMaps,omitempty"`
	PropertyBag genruntime.PropertyBag      `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20240401.RegistryPrivateEndpointConnection
// Private endpoint connection definition.
type RegistryPrivateEndpointConnection struct {
	GroupIds          []string                 `json:"groupIds,omitempty"`
	Location          *string                  `json:"location,omitempty"`
	PrivateEndpoint   *PrivateEndpointResource `json:"privateEndpoint,omitempty"`
	PropertyBag       genruntime.PropertyBag   `json:"$propertyBag,omitempty"`
	ProvisioningState *string                  `json:"provisioningState,omitempty"`

	// Reference: This is the private endpoint connection name created on SRP
	// Full resource id:
	// /subscriptions/{subId}/resourceGroups/{rgName}/providers/Microsoft.MachineLearningServices/{resourceType}/{resourceName}/registryPrivateEndpointConnections/{peConnectionName}
	Reference                                 *genruntime.ResourceReference              `armReference:"Id" json:"reference,omitempty"`
	RegistryPrivateLinkServiceConnectionState *RegistryPrivateLinkServiceConnectionState `json:"registryPrivateLinkServiceConnectionState,omitempty"`
}

// Storage version of v1api20240401.RegistryPrivateEndpointConnection_STATUS
// Private endpoint connection definition.
type RegistryPrivateEndpointConnection_STATUS struct {
	GroupIds                                  []string                                          `json:"groupIds,omitempty"`
	Id                                        *string                                           `json:"id,omitempty"`
	Location                                  *string                                           `json:"location,omitempty"`
	PrivateEndpoint                           *PrivateEndpointResource_STATUS                   `json:"privateEndpoint,omitempty"`
	PropertyBag                               genruntime.PropertyBag                            `json:"$propertyBag,omitempty"`
	ProvisioningState                         *string                                           `json:"provisioningState,omitempty"`
	RegistryPrivateLinkServiceConnectionState *RegistryPrivateLinkServiceConnectionState_STATUS `json:"registryPrivateLinkServiceConnectionState,omitempty"`
}

// Storage version of v1api20240401.RegistryRegionArmDetails
// Details for each region the registry is in
type RegistryRegionArmDetails struct {
	AcrDetails            []AcrDetails            `json:"acrDetails,omitempty"`
	Location              *string                 `json:"location,omitempty"`
	PropertyBag           genruntime.PropertyBag  `json:"$propertyBag,omitempty"`
	StorageAccountDetails []StorageAccountDetails `json:"storageAccountDetails,omitempty"`
}

// Storage version of v1api20240401.RegistryRegionArmDetails_STATUS
// Details for each region the registry is in
type RegistryRegionArmDetails_STATUS struct {
	AcrDetails            []AcrDetails_STATUS            `json:"acrDetails,omitempty"`
	Location              *string                        `json:"location,omitempty"`
	PropertyBag           genruntime.PropertyBag         `json:"$propertyBag,omitempty"`
	StorageAccountDetails []StorageAccountDetails_STATUS `json:"storageAccountDetails,omitempty"`
}

// Storage version of v1api20240401.Sku
// The resource model definition representing SKU
type Sku struct {
	Capacity    *int                   `json:"capacity,omitempty"`
	Family      *string                `json:"family,omitempty"`
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Size        *string                `json:"size,omitempty"`
	Tier        *string                `json:"tier,omitempty"`
}

// Storage version of v1api20240401.Sku_STATUS
// The resource model definition representing SKU
type Sku_STATUS struct {
	Capacity    *int                   `json:"capacity,omitempty"`
	Family      *string                `json:"family,omitempty"`
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Size        *string                `json:"size,omitempty"`
	Tier        *string                `json:"tier,omitempty"`
}

// Storage version of v1api20240401.SystemData_STATUS
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

// Storage version of v1api20240401.AcrDetails
// Details of ACR account to be used for the Registry
type AcrDetails struct {
	PropertyBag             genruntime.PropertyBag   `json:"$propertyBag,omitempty"`
	SystemCreatedAcrAccount *SystemCreatedAcrAccount `json:"systemCreatedAcrAccount,omitempty"`
	UserCreatedAcrAccount   *UserCreatedAcrAccount   `json:"userCreatedAcrAccount,omitempty"`
}

// Storage version of v1api20240401.AcrDetails_STATUS
// Details of ACR account to be used for the Registry
type AcrDetails_STATUS struct {
	PropertyBag             genruntime.PropertyBag          `json:"$propertyBag,omitempty"`
	SystemCreatedAcrAccount *SystemCreatedAcrAccount_STATUS `json:"systemCreatedAcrAccount,omitempty"`
	UserCreatedAcrAccount   *UserCreatedAcrAccount_STATUS   `json:"userCreatedAcrAccount,omitempty"`
}

// Storage version of v1api20240401.PrivateEndpointResource
// The PE network resource that is linked to this PE connection.
type PrivateEndpointResource struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`

	// SubnetArmReference: The subnetId that the private endpoint is connected to.
	SubnetArmReference *genruntime.ResourceReference `armReference:"SubnetArmId" json:"subnetArmReference,omitempty"`
}

// Storage version of v1api20240401.PrivateEndpointResource_STATUS
// The PE network resource that is linked to this PE connection.
type PrivateEndpointResource_STATUS struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	SubnetArmId *string                `json:"subnetArmId,omitempty"`
}

// Storage version of v1api20240401.RegistryOperatorConfigMaps
type RegistryOperatorConfigMaps struct {
	DiscoveryUrl      *genruntime.ConfigMapDestination `json:"discoveryUrl,omitempty"`
	MlFlowRegistryUri *genruntime.ConfigMapDestination `json:"mlFlowRegistryUri,omitempty"`
	PropertyBag       genruntime.PropertyBag           `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20240401.RegistryPrivateLinkServiceConnectionState
// The connection state.
type RegistryPrivateLinkServiceConnectionState struct {
	ActionsRequired *string                `json:"actionsRequired,omitempty"`
	Description     *string                `json:"description,omitempty"`
	PropertyBag     genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Status          *string                `json:"status,omitempty"`
}

// Storage version of v1api20240401.RegistryPrivateLinkServiceConnectionState_STATUS
// The connection state.
type RegistryPrivateLinkServiceConnectionState_STATUS struct {
	ActionsRequired *string                `json:"actionsRequired,omitempty"`
	Description     *string                `json:"description,omitempty"`
	PropertyBag     genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Status          *string                `json:"status,omitempty"`
}

// Storage version of v1api20240401.StorageAccountDetails
// Details of storage account to be used for the Registry
type StorageAccountDetails struct {
	PropertyBag                 genruntime.PropertyBag       `json:"$propertyBag,omitempty"`
	SystemCreatedStorageAccount *SystemCreatedStorageAccount `json:"systemCreatedStorageAccount,omitempty"`
	UserCreatedStorageAccount   *UserCreatedStorageAccount   `json:"userCreatedStorageAccount,omitempty"`
}

// Storage version of v1api20240401.StorageAccountDetails_STATUS
// Details of storage account to be used for the Registry
type StorageAccountDetails_STATUS struct {
	PropertyBag                 genruntime.PropertyBag              `json:"$propertyBag,omitempty"`
	SystemCreatedStorageAccount *SystemCreatedStorageAccount_STATUS `json:"systemCreatedStorageAccount,omitempty"`
	UserCreatedStorageAccount   *UserCreatedStorageAccount_STATUS   `json:"userCreatedStorageAccount,omitempty"`
}

// Storage version of v1api20240401.UserAssignedIdentity_STATUS
// User assigned identity properties
type UserAssignedIdentity_STATUS struct {
	ClientId    *string                `json:"clientId,omitempty"`
	PrincipalId *string                `json:"principalId,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20240401.UserAssignedIdentityDetails
// Information about the user assigned identity for the resource
type UserAssignedIdentityDetails struct {
	PropertyBag genruntime.PropertyBag       `json:"$propertyBag,omitempty"`
	Reference   genruntime.ResourceReference `armReference:"Reference" json:"reference,omitempty"`
}

// Storage version of v1api20240401.SystemCreatedAcrAccount
type SystemCreatedAcrAccount struct {
	AcrAccountName *string                `json:"acrAccountName,omitempty"`
	AcrAccountSku  *string                `json:"acrAccountSku,omitempty"`
	PropertyBag    genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20240401.SystemCreatedAcrAccount_STATUS
type SystemCreatedAcrAccount_STATUS struct {
	AcrAccountName *string                `json:"acrAccountName,omitempty"`
	AcrAccountSku  *string                `json:"acrAccountSku,omitempty"`
	ArmResourceId  *ArmResourceId_STATUS  `json:"armResourceId,omitempty"`
	PropertyBag    genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20240401.SystemCreatedStorageAccount
type SystemCreatedStorageAccount struct {
	AllowBlobPublicAccess    *bool                  `json:"allowBlobPublicAccess,omitempty"`
	PropertyBag              genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	StorageAccountHnsEnabled *bool                  `json:"storageAccountHnsEnabled,omitempty"`
	StorageAccountName       *string                `json:"storageAccountName,omitempty"`
	StorageAccountType       *string                `json:"storageAccountType,omitempty"`
}

// Storage version of v1api20240401.SystemCreatedStorageAccount_STATUS
type SystemCreatedStorageAccount_STATUS struct {
	AllowBlobPublicAccess    *bool                  `json:"allowBlobPublicAccess,omitempty"`
	ArmResourceId            *ArmResourceId_STATUS  `json:"armResourceId,omitempty"`
	PropertyBag              genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	StorageAccountHnsEnabled *bool                  `json:"storageAccountHnsEnabled,omitempty"`
	StorageAccountName       *string                `json:"storageAccountName,omitempty"`
	StorageAccountType       *string                `json:"storageAccountType,omitempty"`
}

// Storage version of v1api20240401.UserCreatedAcrAccount
type UserCreatedAcrAccount struct {
	ArmResourceId *ArmResourceId         `json:"armResourceId,omitempty"`
	PropertyBag   genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20240401.UserCreatedAcrAccount_STATUS
type UserCreatedAcrAccount_STATUS struct {
	ArmResourceId *ArmResourceId_STATUS  `json:"armResourceId,omitempty"`
	PropertyBag   genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20240401.UserCreatedStorageAccount
type UserCreatedStorageAccount struct {
	ArmResourceId *ArmResourceId         `json:"armResourceId,omitempty"`
	PropertyBag   genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20240401.UserCreatedStorageAccount_STATUS
type UserCreatedStorageAccount_STATUS struct {
	ArmResourceId *ArmResourceId_STATUS  `json:"armResourceId,omitempty"`
	PropertyBag   genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

func init() {
	SchemeBuilder.Register(&Registry{}, &RegistryList{})
}
