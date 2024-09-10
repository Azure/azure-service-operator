// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=synapse.azure.com,resources=workspaces,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=synapse.azure.com,resources={workspaces/status,workspaces/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20210601.Workspace
// Generator information:
// - Generated from: /synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/workspace.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}
type Workspace struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Workspace_Spec   `json:"spec,omitempty"`
	Status            Workspace_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &Workspace{}

// GetConditions returns the conditions of the resource
func (workspace *Workspace) GetConditions() conditions.Conditions {
	return workspace.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (workspace *Workspace) SetConditions(conditions conditions.Conditions) {
	workspace.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &Workspace{}

// AzureName returns the Azure name of the resource
func (workspace *Workspace) AzureName() string {
	return workspace.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-06-01"
func (workspace Workspace) GetAPIVersion() string {
	return "2021-06-01"
}

// GetResourceScope returns the scope of the resource
func (workspace *Workspace) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (workspace *Workspace) GetSpec() genruntime.ConvertibleSpec {
	return &workspace.Spec
}

// GetStatus returns the status of this resource
func (workspace *Workspace) GetStatus() genruntime.ConvertibleStatus {
	return &workspace.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (workspace *Workspace) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Synapse/workspaces"
func (workspace *Workspace) GetType() string {
	return "Microsoft.Synapse/workspaces"
}

// NewEmptyStatus returns a new empty (blank) status
func (workspace *Workspace) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &Workspace_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (workspace *Workspace) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(workspace.Spec)
	return workspace.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (workspace *Workspace) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*Workspace_STATUS); ok {
		workspace.Status = *st
		return nil
	}

	// Convert status to required version
	var st Workspace_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	workspace.Status = st
	return nil
}

// Hub marks that this Workspace is the hub type for conversion
func (workspace *Workspace) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (workspace *Workspace) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: workspace.Spec.OriginalVersion,
		Kind:    "Workspace",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20210601.Workspace
// Generator information:
// - Generated from: /synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/workspace.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}
type WorkspaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Workspace `json:"items"`
}

// Storage version of v1api20210601.APIVersion
// +kubebuilder:validation:Enum={"2021-06-01"}
type APIVersion string

const APIVersion_Value = APIVersion("2021-06-01")

// Storage version of v1api20210601.Workspace_Spec
type Workspace_Spec struct {
	AzureADOnlyAuthentication *bool `json:"azureADOnlyAuthentication,omitempty"`

	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName                     string                         `json:"azureName,omitempty"`
	CspWorkspaceAdminProperties   *CspWorkspaceAdminProperties   `json:"cspWorkspaceAdminProperties,omitempty"`
	DefaultDataLakeStorage        *DataLakeStorageAccountDetails `json:"defaultDataLakeStorage,omitempty"`
	Encryption                    *EncryptionDetails             `json:"encryption,omitempty"`
	Identity                      *ManagedIdentity               `json:"identity,omitempty"`
	Location                      *string                        `json:"location,omitempty"`
	ManagedResourceGroupName      *string                        `json:"managedResourceGroupName,omitempty"`
	ManagedVirtualNetwork         *string                        `json:"managedVirtualNetwork,omitempty"`
	ManagedVirtualNetworkSettings *ManagedVirtualNetworkSettings `json:"managedVirtualNetworkSettings,omitempty"`
	OriginalVersion               string                         `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a resources.azure.com/ResourceGroup resource
	Owner                            *genruntime.KnownResourceReference `group:"resources.azure.com" json:"owner,omitempty" kind:"ResourceGroup"`
	PropertyBag                      genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	PublicNetworkAccess              *string                            `json:"publicNetworkAccess,omitempty"`
	PurviewConfiguration             *PurviewConfiguration              `json:"purviewConfiguration,omitempty"`
	SqlAdministratorLogin            *string                            `json:"sqlAdministratorLogin,omitempty"`
	SqlAdministratorLoginPassword    *genruntime.SecretReference        `json:"sqlAdministratorLoginPassword,omitempty"`
	Tags                             map[string]string                  `json:"tags,omitempty"`
	TrustedServiceBypassEnabled      *bool                              `json:"trustedServiceBypassEnabled,omitempty"`
	VirtualNetworkProfile            *VirtualNetworkProfile             `json:"virtualNetworkProfile,omitempty"`
	WorkspaceRepositoryConfiguration *WorkspaceRepositoryConfiguration  `json:"workspaceRepositoryConfiguration,omitempty"`
}

var _ genruntime.ConvertibleSpec = &Workspace_Spec{}

// ConvertSpecFrom populates our Workspace_Spec from the provided source
func (workspace *Workspace_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == workspace {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(workspace)
}

// ConvertSpecTo populates the provided destination from our Workspace_Spec
func (workspace *Workspace_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == workspace {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(workspace)
}

// Storage version of v1api20210601.Workspace_STATUS
// A workspace
type Workspace_STATUS struct {
	AdlaResourceId                   *string                                  `json:"adlaResourceId,omitempty"`
	AzureADOnlyAuthentication        *bool                                    `json:"azureADOnlyAuthentication,omitempty"`
	Conditions                       []conditions.Condition                   `json:"conditions,omitempty"`
	ConnectivityEndpoints            map[string]string                        `json:"connectivityEndpoints,omitempty"`
	CspWorkspaceAdminProperties      *CspWorkspaceAdminProperties_STATUS      `json:"cspWorkspaceAdminProperties,omitempty"`
	DefaultDataLakeStorage           *DataLakeStorageAccountDetails_STATUS    `json:"defaultDataLakeStorage,omitempty"`
	Encryption                       *EncryptionDetails_STATUS                `json:"encryption,omitempty"`
	ExtraProperties                  map[string]v1.JSON                       `json:"extraProperties,omitempty"`
	Id                               *string                                  `json:"id,omitempty"`
	Identity                         *ManagedIdentity_STATUS                  `json:"identity,omitempty"`
	Location                         *string                                  `json:"location,omitempty"`
	ManagedResourceGroupName         *string                                  `json:"managedResourceGroupName,omitempty"`
	ManagedVirtualNetwork            *string                                  `json:"managedVirtualNetwork,omitempty"`
	ManagedVirtualNetworkSettings    *ManagedVirtualNetworkSettings_STATUS    `json:"managedVirtualNetworkSettings,omitempty"`
	Name                             *string                                  `json:"name,omitempty"`
	PrivateEndpointConnections       []PrivateEndpointConnection_STATUS       `json:"privateEndpointConnections,omitempty"`
	PropertyBag                      genruntime.PropertyBag                   `json:"$propertyBag,omitempty"`
	ProvisioningState                *string                                  `json:"provisioningState,omitempty"`
	PublicNetworkAccess              *string                                  `json:"publicNetworkAccess,omitempty"`
	PurviewConfiguration             *PurviewConfiguration_STATUS             `json:"purviewConfiguration,omitempty"`
	Settings                         map[string]v1.JSON                       `json:"settings,omitempty"`
	SqlAdministratorLogin            *string                                  `json:"sqlAdministratorLogin,omitempty"`
	Tags                             map[string]string                        `json:"tags,omitempty"`
	TrustedServiceBypassEnabled      *bool                                    `json:"trustedServiceBypassEnabled,omitempty"`
	Type                             *string                                  `json:"type,omitempty"`
	VirtualNetworkProfile            *VirtualNetworkProfile_STATUS            `json:"virtualNetworkProfile,omitempty"`
	WorkspaceRepositoryConfiguration *WorkspaceRepositoryConfiguration_STATUS `json:"workspaceRepositoryConfiguration,omitempty"`
	WorkspaceUID                     *string                                  `json:"workspaceUID,omitempty"`
}

var _ genruntime.ConvertibleStatus = &Workspace_STATUS{}

// ConvertStatusFrom populates our Workspace_STATUS from the provided source
func (workspace *Workspace_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == workspace {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(workspace)
}

// ConvertStatusTo populates the provided destination from our Workspace_STATUS
func (workspace *Workspace_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == workspace {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(workspace)
}

// Storage version of v1api20210601.CspWorkspaceAdminProperties
// Initial workspace AAD admin properties for a CSP subscription
type CspWorkspaceAdminProperties struct {
	InitialWorkspaceAdminObjectId *string                `json:"initialWorkspaceAdminObjectId,omitempty"`
	PropertyBag                   genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20210601.CspWorkspaceAdminProperties_STATUS
// Initial workspace AAD admin properties for a CSP subscription
type CspWorkspaceAdminProperties_STATUS struct {
	InitialWorkspaceAdminObjectId *string                `json:"initialWorkspaceAdminObjectId,omitempty"`
	PropertyBag                   genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20210601.DataLakeStorageAccountDetails
// Details of the data lake storage account associated with the workspace
type DataLakeStorageAccountDetails struct {
	AccountUrl                   *string                        `json:"accountUrl,omitempty" optionalConfigMapPair:"AccountUrl"`
	AccountUrlFromConfig         *genruntime.ConfigMapReference `json:"accountUrlFromConfig,omitempty" optionalConfigMapPair:"AccountUrl"`
	CreateManagedPrivateEndpoint *bool                          `json:"createManagedPrivateEndpoint,omitempty"`
	Filesystem                   *string                        `json:"filesystem,omitempty"`
	PropertyBag                  genruntime.PropertyBag         `json:"$propertyBag,omitempty"`

	// ResourceReference: ARM resource Id of this storage account
	ResourceReference *genruntime.ResourceReference `armReference:"ResourceId" json:"resourceReference,omitempty"`
}

// Storage version of v1api20210601.DataLakeStorageAccountDetails_STATUS
// Details of the data lake storage account associated with the workspace
type DataLakeStorageAccountDetails_STATUS struct {
	AccountUrl                   *string                `json:"accountUrl,omitempty"`
	CreateManagedPrivateEndpoint *bool                  `json:"createManagedPrivateEndpoint,omitempty"`
	Filesystem                   *string                `json:"filesystem,omitempty"`
	PropertyBag                  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ResourceId                   *string                `json:"resourceId,omitempty"`
}

// Storage version of v1api20210601.EncryptionDetails
// Details of the encryption associated with the workspace
type EncryptionDetails struct {
	Cmk         *CustomerManagedKeyDetails `json:"cmk,omitempty"`
	PropertyBag genruntime.PropertyBag     `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20210601.EncryptionDetails_STATUS
// Details of the encryption associated with the workspace
type EncryptionDetails_STATUS struct {
	Cmk                     *CustomerManagedKeyDetails_STATUS `json:"cmk,omitempty"`
	DoubleEncryptionEnabled *bool                             `json:"doubleEncryptionEnabled,omitempty"`
	PropertyBag             genruntime.PropertyBag            `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20210601.ManagedIdentity
// The workspace managed identity
type ManagedIdentity struct {
	PropertyBag            genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	Type                   *string                       `json:"type,omitempty"`
	UserAssignedIdentities []UserAssignedIdentityDetails `json:"userAssignedIdentities,omitempty"`
}

// Storage version of v1api20210601.ManagedIdentity_STATUS
// The workspace managed identity
type ManagedIdentity_STATUS struct {
	PrincipalId            *string                                       `json:"principalId,omitempty"`
	PropertyBag            genruntime.PropertyBag                        `json:"$propertyBag,omitempty"`
	TenantId               *string                                       `json:"tenantId,omitempty"`
	Type                   *string                                       `json:"type,omitempty"`
	UserAssignedIdentities map[string]UserAssignedManagedIdentity_STATUS `json:"userAssignedIdentities,omitempty"`
}

// Storage version of v1api20210601.ManagedVirtualNetworkSettings
// Managed Virtual Network Settings
type ManagedVirtualNetworkSettings struct {
	AllowedAadTenantIdsForLinking     []string               `json:"allowedAadTenantIdsForLinking,omitempty"`
	LinkedAccessCheckOnTargetResource *bool                  `json:"linkedAccessCheckOnTargetResource,omitempty"`
	PreventDataExfiltration           *bool                  `json:"preventDataExfiltration,omitempty"`
	PropertyBag                       genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20210601.ManagedVirtualNetworkSettings_STATUS
// Managed Virtual Network Settings
type ManagedVirtualNetworkSettings_STATUS struct {
	AllowedAadTenantIdsForLinking     []string               `json:"allowedAadTenantIdsForLinking,omitempty"`
	LinkedAccessCheckOnTargetResource *bool                  `json:"linkedAccessCheckOnTargetResource,omitempty"`
	PreventDataExfiltration           *bool                  `json:"preventDataExfiltration,omitempty"`
	PropertyBag                       genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20210601.PrivateEndpointConnection_STATUS
// A private endpoint connection
type PrivateEndpointConnection_STATUS struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20210601.PurviewConfiguration
// Purview Configuration
type PurviewConfiguration struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`

	// PurviewResourceReference: Purview Resource ID
	PurviewResourceReference *genruntime.ResourceReference `armReference:"PurviewResourceId" json:"purviewResourceReference,omitempty"`
}

// Storage version of v1api20210601.PurviewConfiguration_STATUS
// Purview Configuration
type PurviewConfiguration_STATUS struct {
	PropertyBag       genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	PurviewResourceId *string                `json:"purviewResourceId,omitempty"`
}

// Storage version of v1api20210601.VirtualNetworkProfile
// Virtual Network Profile
type VirtualNetworkProfile struct {
	ComputeSubnetId *string                `json:"computeSubnetId,omitempty"`
	PropertyBag     genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20210601.VirtualNetworkProfile_STATUS
// Virtual Network Profile
type VirtualNetworkProfile_STATUS struct {
	ComputeSubnetId *string                `json:"computeSubnetId,omitempty"`
	PropertyBag     genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20210601.WorkspaceRepositoryConfiguration
// Git integration settings
type WorkspaceRepositoryConfiguration struct {
	AccountName         *string                `json:"accountName,omitempty"`
	CollaborationBranch *string                `json:"collaborationBranch,omitempty"`
	HostName            *string                `json:"hostName,omitempty"`
	LastCommitId        *string                `json:"lastCommitId,omitempty"`
	ProjectName         *string                `json:"projectName,omitempty"`
	PropertyBag         genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	RepositoryName      *string                `json:"repositoryName,omitempty"`
	RootFolder          *string                `json:"rootFolder,omitempty"`
	TenantId            *string                `json:"tenantId,omitempty"`
	Type                *string                `json:"type,omitempty"`
}

// Storage version of v1api20210601.WorkspaceRepositoryConfiguration_STATUS
// Git integration settings
type WorkspaceRepositoryConfiguration_STATUS struct {
	AccountName         *string                `json:"accountName,omitempty"`
	CollaborationBranch *string                `json:"collaborationBranch,omitempty"`
	HostName            *string                `json:"hostName,omitempty"`
	LastCommitId        *string                `json:"lastCommitId,omitempty"`
	ProjectName         *string                `json:"projectName,omitempty"`
	PropertyBag         genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	RepositoryName      *string                `json:"repositoryName,omitempty"`
	RootFolder          *string                `json:"rootFolder,omitempty"`
	TenantId            *string                `json:"tenantId,omitempty"`
	Type                *string                `json:"type,omitempty"`
}

// Storage version of v1api20210601.CustomerManagedKeyDetails
// Details of the customer managed key associated with the workspace
type CustomerManagedKeyDetails struct {
	KekIdentity *KekIdentityProperties `json:"kekIdentity,omitempty"`
	Key         *WorkspaceKeyDetails   `json:"key,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20210601.CustomerManagedKeyDetails_STATUS
// Details of the customer managed key associated with the workspace
type CustomerManagedKeyDetails_STATUS struct {
	KekIdentity *KekIdentityProperties_STATUS `json:"kekIdentity,omitempty"`
	Key         *WorkspaceKeyDetails_STATUS   `json:"key,omitempty"`
	PropertyBag genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	Status      *string                       `json:"status,omitempty"`
}

// Storage version of v1api20210601.UserAssignedIdentityDetails
// Information about the user assigned identity for the resource
type UserAssignedIdentityDetails struct {
	PropertyBag genruntime.PropertyBag       `json:"$propertyBag,omitempty"`
	Reference   genruntime.ResourceReference `armReference:"Reference" json:"reference,omitempty"`
}

// Storage version of v1api20210601.UserAssignedManagedIdentity_STATUS
// User Assigned Managed Identity
type UserAssignedManagedIdentity_STATUS struct {
	ClientId    *string                `json:"clientId,omitempty"`
	PrincipalId *string                `json:"principalId,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20210601.KekIdentityProperties
// Key encryption key properties
type KekIdentityProperties struct {
	PropertyBag               genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	UseSystemAssignedIdentity *v1.JSON               `json:"useSystemAssignedIdentity,omitempty"`

	// UserAssignedIdentityReference: User assigned identity resource Id
	UserAssignedIdentityReference *genruntime.ResourceReference `armReference:"UserAssignedIdentity" json:"userAssignedIdentityReference,omitempty"`
}

// Storage version of v1api20210601.KekIdentityProperties_STATUS
// Key encryption key properties
type KekIdentityProperties_STATUS struct {
	PropertyBag               genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	UseSystemAssignedIdentity *v1.JSON               `json:"useSystemAssignedIdentity,omitempty"`
	UserAssignedIdentity      *string                `json:"userAssignedIdentity,omitempty"`
}

// Storage version of v1api20210601.WorkspaceKeyDetails
// Details of the customer managed key associated with the workspace
type WorkspaceKeyDetails struct {
	KeyVaultUrl *string                `json:"keyVaultUrl,omitempty"`
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20210601.WorkspaceKeyDetails_STATUS
// Details of the customer managed key associated with the workspace
type WorkspaceKeyDetails_STATUS struct {
	KeyVaultUrl *string                `json:"keyVaultUrl,omitempty"`
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

func init() {
	SchemeBuilder.Register(&Workspace{}, &WorkspaceList{})
}
