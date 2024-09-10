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

// +kubebuilder:rbac:groups=machinelearningservices.azure.com,resources=workspaces,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=machinelearningservices.azure.com,resources={workspaces/status,workspaces/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20210701.Workspace
// Generator information:
// - Generated from: /machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2021-07-01/machineLearningServices.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}
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

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-07-01"
func (workspace Workspace) GetAPIVersion() string {
	return "2021-07-01"
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

// GetType returns the ARM Type of the resource. This is always "Microsoft.MachineLearningServices/workspaces"
func (workspace *Workspace) GetType() string {
	return "Microsoft.MachineLearningServices/workspaces"
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
// Storage version of v1api20210701.Workspace
// Generator information:
// - Generated from: /machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2021-07-01/machineLearningServices.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}
type WorkspaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Workspace `json:"items"`
}

// Storage version of v1api20210701.APIVersion
// +kubebuilder:validation:Enum={"2021-07-01"}
type APIVersion string

const APIVersion_Value = APIVersion("2021-07-01")

// Storage version of v1api20210701.Workspace_Spec
type Workspace_Spec struct {
	AllowPublicAccessWhenBehindVnet *bool `json:"allowPublicAccessWhenBehindVnet,omitempty"`

	// ApplicationInsightsReference: ARM id of the application insights associated with this workspace. This cannot be changed
	// once the workspace has been created
	ApplicationInsightsReference *genruntime.ResourceReference `armReference:"ApplicationInsights" json:"applicationInsightsReference,omitempty"`

	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName string `json:"azureName,omitempty"`

	// ContainerRegistryReference: ARM id of the container registry associated with this workspace. This cannot be changed once
	// the workspace has been created
	ContainerRegistryReference *genruntime.ResourceReference `armReference:"ContainerRegistry" json:"containerRegistryReference,omitempty"`
	Description                *string                       `json:"description,omitempty"`
	DiscoveryUrl               *string                       `json:"discoveryUrl,omitempty"`
	Encryption                 *EncryptionProperty           `json:"encryption,omitempty"`
	FriendlyName               *string                       `json:"friendlyName,omitempty"`
	HbiWorkspace               *bool                         `json:"hbiWorkspace,omitempty"`
	Identity                   *Identity                     `json:"identity,omitempty"`
	ImageBuildCompute          *string                       `json:"imageBuildCompute,omitempty"`

	// KeyVaultReference: ARM id of the key vault associated with this workspace. This cannot be changed once the workspace has
	// been created
	KeyVaultReference *genruntime.ResourceReference `armReference:"KeyVault" json:"keyVaultReference,omitempty"`
	Location          *string                       `json:"location,omitempty"`
	OperatorSpec      *WorkspaceOperatorSpec        `json:"operatorSpec,omitempty"`
	OriginalVersion   string                        `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a resources.azure.com/ResourceGroup resource
	Owner *genruntime.KnownResourceReference `group:"resources.azure.com" json:"owner,omitempty" kind:"ResourceGroup"`

	// PrimaryUserAssignedIdentityReference: The user assigned identity resource id that represents the workspace identity.
	PrimaryUserAssignedIdentityReference *genruntime.ResourceReference    `armReference:"PrimaryUserAssignedIdentity" json:"primaryUserAssignedIdentityReference,omitempty"`
	PropertyBag                          genruntime.PropertyBag           `json:"$propertyBag,omitempty"`
	PublicNetworkAccess                  *string                          `json:"publicNetworkAccess,omitempty"`
	ServiceManagedResourcesSettings      *ServiceManagedResourcesSettings `json:"serviceManagedResourcesSettings,omitempty"`
	SharedPrivateLinkResources           []SharedPrivateLinkResource      `json:"sharedPrivateLinkResources,omitempty"`
	Sku                                  *Sku                             `json:"sku,omitempty"`

	// StorageAccountReference: ARM id of the storage account associated with this workspace. This cannot be changed once the
	// workspace has been created
	StorageAccountReference *genruntime.ResourceReference `armReference:"StorageAccount" json:"storageAccountReference,omitempty"`
	SystemData              *SystemData                   `json:"systemData,omitempty"`
	Tags                    map[string]string             `json:"tags,omitempty"`
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

// Storage version of v1api20210701.Workspace_STATUS
// An object that represents a machine learning workspace.
type Workspace_STATUS struct {
	AllowPublicAccessWhenBehindVnet *bool                                   `json:"allowPublicAccessWhenBehindVnet,omitempty"`
	ApplicationInsights             *string                                 `json:"applicationInsights,omitempty"`
	Conditions                      []conditions.Condition                  `json:"conditions,omitempty"`
	ContainerRegistry               *string                                 `json:"containerRegistry,omitempty"`
	Description                     *string                                 `json:"description,omitempty"`
	DiscoveryUrl                    *string                                 `json:"discoveryUrl,omitempty"`
	Encryption                      *EncryptionProperty_STATUS              `json:"encryption,omitempty"`
	FriendlyName                    *string                                 `json:"friendlyName,omitempty"`
	HbiWorkspace                    *bool                                   `json:"hbiWorkspace,omitempty"`
	Id                              *string                                 `json:"id,omitempty"`
	Identity                        *Identity_STATUS                        `json:"identity,omitempty"`
	ImageBuildCompute               *string                                 `json:"imageBuildCompute,omitempty"`
	KeyVault                        *string                                 `json:"keyVault,omitempty"`
	Location                        *string                                 `json:"location,omitempty"`
	MlFlowTrackingUri               *string                                 `json:"mlFlowTrackingUri,omitempty"`
	Name                            *string                                 `json:"name,omitempty"`
	NotebookInfo                    *NotebookResourceInfo_STATUS            `json:"notebookInfo,omitempty"`
	PrimaryUserAssignedIdentity     *string                                 `json:"primaryUserAssignedIdentity,omitempty"`
	PrivateEndpointConnections      []PrivateEndpointConnection_STATUS      `json:"privateEndpointConnections,omitempty"`
	PrivateLinkCount                *int                                    `json:"privateLinkCount,omitempty"`
	PropertyBag                     genruntime.PropertyBag                  `json:"$propertyBag,omitempty"`
	ProvisioningState               *string                                 `json:"provisioningState,omitempty"`
	PublicNetworkAccess             *string                                 `json:"publicNetworkAccess,omitempty"`
	ServiceManagedResourcesSettings *ServiceManagedResourcesSettings_STATUS `json:"serviceManagedResourcesSettings,omitempty"`
	ServiceProvisionedResourceGroup *string                                 `json:"serviceProvisionedResourceGroup,omitempty"`
	SharedPrivateLinkResources      []SharedPrivateLinkResource_STATUS      `json:"sharedPrivateLinkResources,omitempty"`
	Sku                             *Sku_STATUS                             `json:"sku,omitempty"`
	StorageAccount                  *string                                 `json:"storageAccount,omitempty"`
	StorageHnsEnabled               *bool                                   `json:"storageHnsEnabled,omitempty"`
	SystemData                      *SystemData_STATUS                      `json:"systemData,omitempty"`
	Tags                            map[string]string                       `json:"tags,omitempty"`
	TenantId                        *string                                 `json:"tenantId,omitempty"`
	Type                            *string                                 `json:"type,omitempty"`
	WorkspaceId                     *string                                 `json:"workspaceId,omitempty"`
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

// Storage version of v1api20210701.EncryptionProperty
type EncryptionProperty struct {
	Identity           *IdentityForCmk        `json:"identity,omitempty"`
	KeyVaultProperties *KeyVaultProperties    `json:"keyVaultProperties,omitempty"`
	PropertyBag        genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Status             *string                `json:"status,omitempty"`
}

// Storage version of v1api20210701.EncryptionProperty_STATUS
type EncryptionProperty_STATUS struct {
	Identity           *IdentityForCmk_STATUS     `json:"identity,omitempty"`
	KeyVaultProperties *KeyVaultProperties_STATUS `json:"keyVaultProperties,omitempty"`
	PropertyBag        genruntime.PropertyBag     `json:"$propertyBag,omitempty"`
	Status             *string                    `json:"status,omitempty"`
}

// Storage version of v1api20210701.Identity
// Identity for the resource.
type Identity struct {
	PropertyBag            genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	Type                   *string                       `json:"type,omitempty"`
	UserAssignedIdentities []UserAssignedIdentityDetails `json:"userAssignedIdentities,omitempty"`
}

// Storage version of v1api20210701.Identity_STATUS
// Identity for the resource.
type Identity_STATUS struct {
	PrincipalId            *string                                `json:"principalId,omitempty"`
	PropertyBag            genruntime.PropertyBag                 `json:"$propertyBag,omitempty"`
	TenantId               *string                                `json:"tenantId,omitempty"`
	Type                   *string                                `json:"type,omitempty"`
	UserAssignedIdentities map[string]UserAssignedIdentity_STATUS `json:"userAssignedIdentities,omitempty"`
}

// Storage version of v1api20210701.NotebookResourceInfo_STATUS
type NotebookResourceInfo_STATUS struct {
	Fqdn                     *string                          `json:"fqdn,omitempty"`
	NotebookPreparationError *NotebookPreparationError_STATUS `json:"notebookPreparationError,omitempty"`
	PropertyBag              genruntime.PropertyBag           `json:"$propertyBag,omitempty"`
	ResourceId               *string                          `json:"resourceId,omitempty"`
}

// Storage version of v1api20210701.PrivateEndpointConnection_STATUS
// The Private Endpoint Connection resource.
type PrivateEndpointConnection_STATUS struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20210701.ServiceManagedResourcesSettings
type ServiceManagedResourcesSettings struct {
	CosmosDb    *CosmosDbSettings      `json:"cosmosDb,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20210701.ServiceManagedResourcesSettings_STATUS
type ServiceManagedResourcesSettings_STATUS struct {
	CosmosDb    *CosmosDbSettings_STATUS `json:"cosmosDb,omitempty"`
	PropertyBag genruntime.PropertyBag   `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20210701.SharedPrivateLinkResource
type SharedPrivateLinkResource struct {
	GroupId *string `json:"groupId,omitempty"`
	Name    *string `json:"name,omitempty"`

	// PrivateLinkResourceReference: The resource id that private link links to.
	PrivateLinkResourceReference *genruntime.ResourceReference `armReference:"PrivateLinkResourceId" json:"privateLinkResourceReference,omitempty"`
	PropertyBag                  genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	RequestMessage               *string                       `json:"requestMessage,omitempty"`
	Status                       *string                       `json:"status,omitempty"`
}

// Storage version of v1api20210701.SharedPrivateLinkResource_STATUS
type SharedPrivateLinkResource_STATUS struct {
	GroupId               *string                `json:"groupId,omitempty"`
	Name                  *string                `json:"name,omitempty"`
	PrivateLinkResourceId *string                `json:"privateLinkResourceId,omitempty"`
	PropertyBag           genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	RequestMessage        *string                `json:"requestMessage,omitempty"`
	Status                *string                `json:"status,omitempty"`
}

// Storage version of v1api20210701.Sku
// Sku of the resource
type Sku struct {
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Tier        *string                `json:"tier,omitempty"`
}

// Storage version of v1api20210701.Sku_STATUS
// Sku of the resource
type Sku_STATUS struct {
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Tier        *string                `json:"tier,omitempty"`
}

// Storage version of v1api20210701.SystemData
// Metadata pertaining to creation and last modification of the resource.
type SystemData struct {
	CreatedAt          *string                `json:"createdAt,omitempty"`
	CreatedBy          *string                `json:"createdBy,omitempty"`
	CreatedByType      *string                `json:"createdByType,omitempty"`
	LastModifiedAt     *string                `json:"lastModifiedAt,omitempty"`
	LastModifiedBy     *string                `json:"lastModifiedBy,omitempty"`
	LastModifiedByType *string                `json:"lastModifiedByType,omitempty"`
	PropertyBag        genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20210701.SystemData_STATUS
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

// Storage version of v1api20210701.WorkspaceOperatorSpec
// Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure
type WorkspaceOperatorSpec struct {
	PropertyBag genruntime.PropertyBag    `json:"$propertyBag,omitempty"`
	Secrets     *WorkspaceOperatorSecrets `json:"secrets,omitempty"`
}

// Storage version of v1api20210701.CosmosDbSettings
type CosmosDbSettings struct {
	CollectionsThroughput *int                   `json:"collectionsThroughput,omitempty"`
	PropertyBag           genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20210701.CosmosDbSettings_STATUS
type CosmosDbSettings_STATUS struct {
	CollectionsThroughput *int                   `json:"collectionsThroughput,omitempty"`
	PropertyBag           genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20210701.IdentityForCmk
// Identity that will be used to access key vault for encryption at rest
type IdentityForCmk struct {
	PropertyBag          genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	UserAssignedIdentity *string                `json:"userAssignedIdentity,omitempty"`
}

// Storage version of v1api20210701.IdentityForCmk_STATUS
// Identity that will be used to access key vault for encryption at rest
type IdentityForCmk_STATUS struct {
	PropertyBag          genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	UserAssignedIdentity *string                `json:"userAssignedIdentity,omitempty"`
}

// Storage version of v1api20210701.KeyVaultProperties
type KeyVaultProperties struct {
	IdentityClientId *string                `json:"identityClientId,omitempty"`
	KeyIdentifier    *string                `json:"keyIdentifier,omitempty"`
	KeyVaultArmId    *string                `json:"keyVaultArmId,omitempty"`
	PropertyBag      genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20210701.KeyVaultProperties_STATUS
type KeyVaultProperties_STATUS struct {
	IdentityClientId *string                `json:"identityClientId,omitempty"`
	KeyIdentifier    *string                `json:"keyIdentifier,omitempty"`
	KeyVaultArmId    *string                `json:"keyVaultArmId,omitempty"`
	PropertyBag      genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20210701.NotebookPreparationError_STATUS
type NotebookPreparationError_STATUS struct {
	ErrorMessage *string                `json:"errorMessage,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	StatusCode   *int                   `json:"statusCode,omitempty"`
}

// Storage version of v1api20210701.UserAssignedIdentity_STATUS
// User Assigned Identity
type UserAssignedIdentity_STATUS struct {
	ClientId    *string                `json:"clientId,omitempty"`
	PrincipalId *string                `json:"principalId,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	TenantId    *string                `json:"tenantId,omitempty"`
}

// Storage version of v1api20210701.UserAssignedIdentityDetails
// Information about the user assigned identity for the resource
type UserAssignedIdentityDetails struct {
	PropertyBag genruntime.PropertyBag       `json:"$propertyBag,omitempty"`
	Reference   genruntime.ResourceReference `armReference:"Reference" json:"reference,omitempty"`
}

// Storage version of v1api20210701.WorkspaceOperatorSecrets
type WorkspaceOperatorSecrets struct {
	AppInsightsInstrumentationKey *genruntime.SecretDestination `json:"appInsightsInstrumentationKey,omitempty"`
	ContainerRegistryPassword     *genruntime.SecretDestination `json:"containerRegistryPassword,omitempty"`
	ContainerRegistryPassword2    *genruntime.SecretDestination `json:"containerRegistryPassword2,omitempty"`
	ContainerRegistryUserName     *genruntime.SecretDestination `json:"containerRegistryUserName,omitempty"`
	PrimaryNotebookAccessKey      *genruntime.SecretDestination `json:"primaryNotebookAccessKey,omitempty"`
	PropertyBag                   genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	SecondaryNotebookAccessKey    *genruntime.SecretDestination `json:"secondaryNotebookAccessKey,omitempty"`
	UserStorageKey                *genruntime.SecretDestination `json:"userStorageKey,omitempty"`
}

func init() {
	SchemeBuilder.Register(&Workspace{}, &WorkspaceList{})
}
