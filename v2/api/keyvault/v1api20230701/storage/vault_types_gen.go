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

// +kubebuilder:rbac:groups=keyvault.azure.com,resources=vaults,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=keyvault.azure.com,resources={vaults/status,vaults/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20230701.Vault
// Generator information:
// - Generated from: /keyvault/resource-manager/Microsoft.KeyVault/stable/2023-07-01/keyvault.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}
type Vault struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Vault_Spec   `json:"spec,omitempty"`
	Status            Vault_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &Vault{}

// GetConditions returns the conditions of the resource
func (vault *Vault) GetConditions() conditions.Conditions {
	return vault.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (vault *Vault) SetConditions(conditions conditions.Conditions) {
	vault.Status.Conditions = conditions
}

var _ configmaps.Exporter = &Vault{}

// ConfigMapDestinationExpressions returns the Spec.OperatorSpec.ConfigMapExpressions property
func (vault *Vault) ConfigMapDestinationExpressions() []*core.DestinationExpression {
	if vault.Spec.OperatorSpec == nil {
		return nil
	}
	return vault.Spec.OperatorSpec.ConfigMapExpressions
}

var _ secrets.Exporter = &Vault{}

// SecretDestinationExpressions returns the Spec.OperatorSpec.SecretExpressions property
func (vault *Vault) SecretDestinationExpressions() []*core.DestinationExpression {
	if vault.Spec.OperatorSpec == nil {
		return nil
	}
	return vault.Spec.OperatorSpec.SecretExpressions
}

var _ genruntime.KubernetesResource = &Vault{}

// AzureName returns the Azure name of the resource
func (vault *Vault) AzureName() string {
	return vault.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2023-07-01"
func (vault Vault) GetAPIVersion() string {
	return "2023-07-01"
}

// GetResourceScope returns the scope of the resource
func (vault *Vault) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (vault *Vault) GetSpec() genruntime.ConvertibleSpec {
	return &vault.Spec
}

// GetStatus returns the status of this resource
func (vault *Vault) GetStatus() genruntime.ConvertibleStatus {
	return &vault.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (vault *Vault) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.KeyVault/vaults"
func (vault *Vault) GetType() string {
	return "Microsoft.KeyVault/vaults"
}

// NewEmptyStatus returns a new empty (blank) status
func (vault *Vault) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &Vault_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (vault *Vault) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(vault.Spec)
	return vault.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (vault *Vault) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*Vault_STATUS); ok {
		vault.Status = *st
		return nil
	}

	// Convert status to required version
	var st Vault_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return eris.Wrap(err, "failed to convert status")
	}

	vault.Status = st
	return nil
}

// Hub marks that this Vault is the hub type for conversion
func (vault *Vault) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (vault *Vault) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: vault.Spec.OriginalVersion,
		Kind:    "Vault",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20230701.Vault
// Generator information:
// - Generated from: /keyvault/resource-manager/Microsoft.KeyVault/stable/2023-07-01/keyvault.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}
type VaultList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Vault `json:"items"`
}

// Storage version of v1api20230701.APIVersion
// +kubebuilder:validation:Enum={"2023-07-01"}
type APIVersion string

const APIVersion_Value = APIVersion("2023-07-01")

// Storage version of v1api20230701.Vault_Spec
type Vault_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName       string             `json:"azureName,omitempty"`
	Location        *string            `json:"location,omitempty"`
	OperatorSpec    *VaultOperatorSpec `json:"operatorSpec,omitempty"`
	OriginalVersion string             `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a resources.azure.com/ResourceGroup resource
	Owner       *genruntime.KnownResourceReference `group:"resources.azure.com" json:"owner,omitempty" kind:"ResourceGroup"`
	Properties  *VaultProperties                   `json:"properties,omitempty"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	Tags        map[string]string                  `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &Vault_Spec{}

// ConvertSpecFrom populates our Vault_Spec from the provided source
func (vault *Vault_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == vault {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(vault)
}

// ConvertSpecTo populates the provided destination from our Vault_Spec
func (vault *Vault_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == vault {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(vault)
}

// Storage version of v1api20230701.Vault_STATUS
// Resource information with extended details.
type Vault_STATUS struct {
	Conditions  []conditions.Condition  `json:"conditions,omitempty"`
	Id          *string                 `json:"id,omitempty"`
	Location    *string                 `json:"location,omitempty"`
	Name        *string                 `json:"name,omitempty"`
	Properties  *VaultProperties_STATUS `json:"properties,omitempty"`
	PropertyBag genruntime.PropertyBag  `json:"$propertyBag,omitempty"`
	SystemData  *SystemData_STATUS      `json:"systemData,omitempty"`
	Tags        map[string]string       `json:"tags,omitempty"`
	Type        *string                 `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &Vault_STATUS{}

// ConvertStatusFrom populates our Vault_STATUS from the provided source
func (vault *Vault_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == vault {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(vault)
}

// ConvertStatusTo populates the provided destination from our Vault_STATUS
func (vault *Vault_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == vault {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(vault)
}

// Storage version of v1api20230701.SystemData_STATUS
// Metadata pertaining to creation and last modification of the key vault resource.
type SystemData_STATUS struct {
	CreatedAt          *string                `json:"createdAt,omitempty"`
	CreatedBy          *string                `json:"createdBy,omitempty"`
	CreatedByType      *string                `json:"createdByType,omitempty"`
	LastModifiedAt     *string                `json:"lastModifiedAt,omitempty"`
	LastModifiedBy     *string                `json:"lastModifiedBy,omitempty"`
	LastModifiedByType *string                `json:"lastModifiedByType,omitempty"`
	PropertyBag        genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20230701.VaultOperatorSpec
// Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure
type VaultOperatorSpec struct {
	ConfigMapExpressions []*core.DestinationExpression `json:"configMapExpressions,omitempty"`
	PropertyBag          genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	SecretExpressions    []*core.DestinationExpression `json:"secretExpressions,omitempty"`
}

// Storage version of v1api20230701.VaultProperties
// Properties of the vault
type VaultProperties struct {
	AccessPolicies               []AccessPolicyEntry            `json:"accessPolicies,omitempty"`
	CreateMode                   *string                        `json:"createMode,omitempty"`
	EnablePurgeProtection        *bool                          `json:"enablePurgeProtection,omitempty"`
	EnableRbacAuthorization      *bool                          `json:"enableRbacAuthorization,omitempty"`
	EnableSoftDelete             *bool                          `json:"enableSoftDelete,omitempty"`
	EnabledForDeployment         *bool                          `json:"enabledForDeployment,omitempty"`
	EnabledForDiskEncryption     *bool                          `json:"enabledForDiskEncryption,omitempty"`
	EnabledForTemplateDeployment *bool                          `json:"enabledForTemplateDeployment,omitempty"`
	NetworkAcls                  *NetworkRuleSet                `json:"networkAcls,omitempty"`
	PropertyBag                  genruntime.PropertyBag         `json:"$propertyBag,omitempty"`
	ProvisioningState            *string                        `json:"provisioningState,omitempty"`
	PublicNetworkAccess          *string                        `json:"publicNetworkAccess,omitempty"`
	Sku                          *Sku                           `json:"sku,omitempty"`
	SoftDeleteRetentionInDays    *int                           `json:"softDeleteRetentionInDays,omitempty"`
	TenantId                     *string                        `json:"tenantId,omitempty" optionalConfigMapPair:"TenantId"`
	TenantIdFromConfig           *genruntime.ConfigMapReference `json:"tenantIdFromConfig,omitempty" optionalConfigMapPair:"TenantId"`
	VaultUri                     *string                        `json:"vaultUri,omitempty"`
}

// Storage version of v1api20230701.VaultProperties_STATUS
// Properties of the vault
type VaultProperties_STATUS struct {
	AccessPolicies               []AccessPolicyEntry_STATUS             `json:"accessPolicies,omitempty"`
	CreateMode                   *string                                `json:"createMode,omitempty"`
	EnablePurgeProtection        *bool                                  `json:"enablePurgeProtection,omitempty"`
	EnableRbacAuthorization      *bool                                  `json:"enableRbacAuthorization,omitempty"`
	EnableSoftDelete             *bool                                  `json:"enableSoftDelete,omitempty"`
	EnabledForDeployment         *bool                                  `json:"enabledForDeployment,omitempty"`
	EnabledForDiskEncryption     *bool                                  `json:"enabledForDiskEncryption,omitempty"`
	EnabledForTemplateDeployment *bool                                  `json:"enabledForTemplateDeployment,omitempty"`
	HsmPoolResourceId            *string                                `json:"hsmPoolResourceId,omitempty"`
	NetworkAcls                  *NetworkRuleSet_STATUS                 `json:"networkAcls,omitempty"`
	PrivateEndpointConnections   []PrivateEndpointConnectionItem_STATUS `json:"privateEndpointConnections,omitempty"`
	PropertyBag                  genruntime.PropertyBag                 `json:"$propertyBag,omitempty"`
	ProvisioningState            *string                                `json:"provisioningState,omitempty"`
	PublicNetworkAccess          *string                                `json:"publicNetworkAccess,omitempty"`
	Sku                          *Sku_STATUS                            `json:"sku,omitempty"`
	SoftDeleteRetentionInDays    *int                                   `json:"softDeleteRetentionInDays,omitempty"`
	TenantId                     *string                                `json:"tenantId,omitempty"`
	VaultUri                     *string                                `json:"vaultUri,omitempty"`
}

// Storage version of v1api20230701.AccessPolicyEntry
// An identity that have access to the key vault. All identities in the array must use the same tenant ID as the key
// vault's tenant ID.
type AccessPolicyEntry struct {
	ApplicationId           *string                        `json:"applicationId,omitempty" optionalConfigMapPair:"ApplicationId"`
	ApplicationIdFromConfig *genruntime.ConfigMapReference `json:"applicationIdFromConfig,omitempty" optionalConfigMapPair:"ApplicationId"`
	ObjectId                *string                        `json:"objectId,omitempty" optionalConfigMapPair:"ObjectId"`
	ObjectIdFromConfig      *genruntime.ConfigMapReference `json:"objectIdFromConfig,omitempty" optionalConfigMapPair:"ObjectId"`
	Permissions             *Permissions                   `json:"permissions,omitempty"`
	PropertyBag             genruntime.PropertyBag         `json:"$propertyBag,omitempty"`
	TenantId                *string                        `json:"tenantId,omitempty" optionalConfigMapPair:"TenantId"`
	TenantIdFromConfig      *genruntime.ConfigMapReference `json:"tenantIdFromConfig,omitempty" optionalConfigMapPair:"TenantId"`
}

// Storage version of v1api20230701.AccessPolicyEntry_STATUS
// An identity that have access to the key vault. All identities in the array must use the same tenant ID as the key
// vault's tenant ID.
type AccessPolicyEntry_STATUS struct {
	ApplicationId *string                `json:"applicationId,omitempty"`
	ObjectId      *string                `json:"objectId,omitempty"`
	Permissions   *Permissions_STATUS    `json:"permissions,omitempty"`
	PropertyBag   genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	TenantId      *string                `json:"tenantId,omitempty"`
}

// Storage version of v1api20230701.NetworkRuleSet
// A set of rules governing the network accessibility of a vault.
type NetworkRuleSet struct {
	Bypass              *string                `json:"bypass,omitempty"`
	DefaultAction       *string                `json:"defaultAction,omitempty"`
	IpRules             []IPRule               `json:"ipRules,omitempty"`
	PropertyBag         genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	VirtualNetworkRules []VirtualNetworkRule   `json:"virtualNetworkRules,omitempty"`
}

// Storage version of v1api20230701.NetworkRuleSet_STATUS
// A set of rules governing the network accessibility of a vault.
type NetworkRuleSet_STATUS struct {
	Bypass              *string                     `json:"bypass,omitempty"`
	DefaultAction       *string                     `json:"defaultAction,omitempty"`
	IpRules             []IPRule_STATUS             `json:"ipRules,omitempty"`
	PropertyBag         genruntime.PropertyBag      `json:"$propertyBag,omitempty"`
	VirtualNetworkRules []VirtualNetworkRule_STATUS `json:"virtualNetworkRules,omitempty"`
}

// Storage version of v1api20230701.PrivateEndpointConnectionItem_STATUS
// Private endpoint connection item.
type PrivateEndpointConnectionItem_STATUS struct {
	Etag                              *string                                   `json:"etag,omitempty"`
	Id                                *string                                   `json:"id,omitempty"`
	PrivateEndpoint                   *PrivateEndpoint_STATUS                   `json:"privateEndpoint,omitempty"`
	PrivateLinkServiceConnectionState *PrivateLinkServiceConnectionState_STATUS `json:"privateLinkServiceConnectionState,omitempty"`
	PropertyBag                       genruntime.PropertyBag                    `json:"$propertyBag,omitempty"`
	ProvisioningState                 *string                                   `json:"provisioningState,omitempty"`
}

// Storage version of v1api20230701.Sku
// SKU details
type Sku struct {
	Family      *string                `json:"family,omitempty"`
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20230701.Sku_STATUS
// SKU details
type Sku_STATUS struct {
	Family      *string                `json:"family,omitempty"`
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20230701.IPRule
// A rule governing the accessibility of a vault from a specific ip address or ip range.
type IPRule struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Value       *string                `json:"value,omitempty"`
}

// Storage version of v1api20230701.IPRule_STATUS
// A rule governing the accessibility of a vault from a specific ip address or ip range.
type IPRule_STATUS struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Value       *string                `json:"value,omitempty"`
}

// Storage version of v1api20230701.Permissions
// Permissions the identity has for keys, secrets, certificates and storage.
type Permissions struct {
	Certificates []string               `json:"certificates,omitempty"`
	Keys         []string               `json:"keys,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Secrets      []string               `json:"secrets,omitempty"`
	Storage      []string               `json:"storage,omitempty"`
}

// Storage version of v1api20230701.Permissions_STATUS
// Permissions the identity has for keys, secrets, certificates and storage.
type Permissions_STATUS struct {
	Certificates []string               `json:"certificates,omitempty"`
	Keys         []string               `json:"keys,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Secrets      []string               `json:"secrets,omitempty"`
	Storage      []string               `json:"storage,omitempty"`
}

// Storage version of v1api20230701.PrivateEndpoint_STATUS
// Private endpoint object properties.
type PrivateEndpoint_STATUS struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20230701.PrivateLinkServiceConnectionState_STATUS
// An object that represents the approval state of the private link connection.
type PrivateLinkServiceConnectionState_STATUS struct {
	ActionsRequired *string                `json:"actionsRequired,omitempty"`
	Description     *string                `json:"description,omitempty"`
	PropertyBag     genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Status          *string                `json:"status,omitempty"`
}

// Storage version of v1api20230701.VirtualNetworkRule
// A rule governing the accessibility of a vault from a specific virtual network.
type VirtualNetworkRule struct {
	IgnoreMissingVnetServiceEndpoint *bool                  `json:"ignoreMissingVnetServiceEndpoint,omitempty"`
	PropertyBag                      genruntime.PropertyBag `json:"$propertyBag,omitempty"`

	// +kubebuilder:validation:Required
	// Reference: Full resource id of a vnet subnet, such as
	// '/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/subnet1'.
	Reference *genruntime.ResourceReference `armReference:"Id" json:"reference,omitempty"`
}

// Storage version of v1api20230701.VirtualNetworkRule_STATUS
// A rule governing the accessibility of a vault from a specific virtual network.
type VirtualNetworkRule_STATUS struct {
	Id                               *string                `json:"id,omitempty"`
	IgnoreMissingVnetServiceEndpoint *bool                  `json:"ignoreMissingVnetServiceEndpoint,omitempty"`
	PropertyBag                      genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

func init() {
	SchemeBuilder.Register(&Vault{}, &VaultList{})
}
