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

// +kubebuilder:rbac:groups=signalrservice.azure.com,resources=signalrs,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=signalrservice.azure.com,resources={signalrs/status,signalrs/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20211001.SignalR
// Generator information:
// - Generated from: /signalr/resource-manager/Microsoft.SignalRService/stable/2021-10-01/signalr.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SignalRService/signalR/{resourceName}
type SignalR struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SignalR_Spec   `json:"spec,omitempty"`
	Status            SignalR_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &SignalR{}

// GetConditions returns the conditions of the resource
func (signalR *SignalR) GetConditions() conditions.Conditions {
	return signalR.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (signalR *SignalR) SetConditions(conditions conditions.Conditions) {
	signalR.Status.Conditions = conditions
}

var _ configmaps.Exporter = &SignalR{}

// ConfigMapDestinationExpressions returns the Spec.OperatorSpec.ConfigMapExpressions property
func (signalR *SignalR) ConfigMapDestinationExpressions() []*core.DestinationExpression {
	if signalR.Spec.OperatorSpec == nil {
		return nil
	}
	return signalR.Spec.OperatorSpec.ConfigMapExpressions
}

var _ secrets.Exporter = &SignalR{}

// SecretDestinationExpressions returns the Spec.OperatorSpec.SecretExpressions property
func (signalR *SignalR) SecretDestinationExpressions() []*core.DestinationExpression {
	if signalR.Spec.OperatorSpec == nil {
		return nil
	}
	return signalR.Spec.OperatorSpec.SecretExpressions
}

var _ genruntime.KubernetesResource = &SignalR{}

// AzureName returns the Azure name of the resource
func (signalR *SignalR) AzureName() string {
	return signalR.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-10-01"
func (signalR SignalR) GetAPIVersion() string {
	return "2021-10-01"
}

// GetResourceScope returns the scope of the resource
func (signalR *SignalR) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (signalR *SignalR) GetSpec() genruntime.ConvertibleSpec {
	return &signalR.Spec
}

// GetStatus returns the status of this resource
func (signalR *SignalR) GetStatus() genruntime.ConvertibleStatus {
	return &signalR.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (signalR *SignalR) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.SignalRService/signalR"
func (signalR *SignalR) GetType() string {
	return "Microsoft.SignalRService/signalR"
}

// NewEmptyStatus returns a new empty (blank) status
func (signalR *SignalR) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &SignalR_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (signalR *SignalR) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(signalR.Spec)
	return signalR.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (signalR *SignalR) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*SignalR_STATUS); ok {
		signalR.Status = *st
		return nil
	}

	// Convert status to required version
	var st SignalR_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return eris.Wrap(err, "failed to convert status")
	}

	signalR.Status = st
	return nil
}

// Hub marks that this SignalR is the hub type for conversion
func (signalR *SignalR) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (signalR *SignalR) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: signalR.Spec.OriginalVersion,
		Kind:    "SignalR",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20211001.SignalR
// Generator information:
// - Generated from: /signalr/resource-manager/Microsoft.SignalRService/stable/2021-10-01/signalr.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SignalRService/signalR/{resourceName}
type SignalRList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SignalR `json:"items"`
}

// Storage version of v1api20211001.APIVersion
// +kubebuilder:validation:Enum={"2021-10-01"}
type APIVersion string

const APIVersion_Value = APIVersion("2021-10-01")

// Storage version of v1api20211001.SignalR_Spec
type SignalR_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName        string               `json:"azureName,omitempty"`
	Cors             *SignalRCorsSettings `json:"cors,omitempty"`
	DisableAadAuth   *bool                `json:"disableAadAuth,omitempty"`
	DisableLocalAuth *bool                `json:"disableLocalAuth,omitempty"`
	Features         []SignalRFeature     `json:"features,omitempty"`
	Identity         *ManagedIdentity     `json:"identity,omitempty"`
	Kind             *string              `json:"kind,omitempty"`
	Location         *string              `json:"location,omitempty"`
	NetworkACLs      *SignalRNetworkACLs  `json:"networkACLs,omitempty"`
	OperatorSpec     *SignalROperatorSpec `json:"operatorSpec,omitempty"`
	OriginalVersion  string               `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a resources.azure.com/ResourceGroup resource
	Owner                    *genruntime.KnownResourceReference `group:"resources.azure.com" json:"owner,omitempty" kind:"ResourceGroup"`
	PropertyBag              genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	PublicNetworkAccess      *string                            `json:"publicNetworkAccess,omitempty"`
	ResourceLogConfiguration *ResourceLogConfiguration          `json:"resourceLogConfiguration,omitempty"`
	Sku                      *ResourceSku                       `json:"sku,omitempty"`
	Tags                     map[string]string                  `json:"tags,omitempty"`
	Tls                      *SignalRTlsSettings                `json:"tls,omitempty"`
	Upstream                 *ServerlessUpstreamSettings        `json:"upstream,omitempty"`
}

var _ genruntime.ConvertibleSpec = &SignalR_Spec{}

// ConvertSpecFrom populates our SignalR_Spec from the provided source
func (signalR *SignalR_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == signalR {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(signalR)
}

// ConvertSpecTo populates the provided destination from our SignalR_Spec
func (signalR *SignalR_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == signalR {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(signalR)
}

// Storage version of v1api20211001.SignalR_STATUS
type SignalR_STATUS struct {
	Conditions                 []conditions.Condition                                         `json:"conditions,omitempty"`
	Cors                       *SignalRCorsSettings_STATUS                                    `json:"cors,omitempty"`
	DisableAadAuth             *bool                                                          `json:"disableAadAuth,omitempty"`
	DisableLocalAuth           *bool                                                          `json:"disableLocalAuth,omitempty"`
	ExternalIP                 *string                                                        `json:"externalIP,omitempty"`
	Features                   []SignalRFeature_STATUS                                        `json:"features,omitempty"`
	HostName                   *string                                                        `json:"hostName,omitempty"`
	HostNamePrefix             *string                                                        `json:"hostNamePrefix,omitempty"`
	Id                         *string                                                        `json:"id,omitempty"`
	Identity                   *ManagedIdentity_STATUS                                        `json:"identity,omitempty"`
	Kind                       *string                                                        `json:"kind,omitempty"`
	Location                   *string                                                        `json:"location,omitempty"`
	Name                       *string                                                        `json:"name,omitempty"`
	NetworkACLs                *SignalRNetworkACLs_STATUS                                     `json:"networkACLs,omitempty"`
	PrivateEndpointConnections []PrivateEndpointConnection_STATUS_SignalR_SubResourceEmbedded `json:"privateEndpointConnections,omitempty"`
	PropertyBag                genruntime.PropertyBag                                         `json:"$propertyBag,omitempty"`
	ProvisioningState          *string                                                        `json:"provisioningState,omitempty"`
	PublicNetworkAccess        *string                                                        `json:"publicNetworkAccess,omitempty"`
	PublicPort                 *int                                                           `json:"publicPort,omitempty"`
	ResourceLogConfiguration   *ResourceLogConfiguration_STATUS                               `json:"resourceLogConfiguration,omitempty"`
	ServerPort                 *int                                                           `json:"serverPort,omitempty"`
	SharedPrivateLinkResources []SharedPrivateLinkResource_STATUS_SignalR_SubResourceEmbedded `json:"sharedPrivateLinkResources,omitempty"`
	Sku                        *ResourceSku_STATUS                                            `json:"sku,omitempty"`
	SystemData                 *SystemData_STATUS                                             `json:"systemData,omitempty"`
	Tags                       map[string]string                                              `json:"tags,omitempty"`
	Tls                        *SignalRTlsSettings_STATUS                                     `json:"tls,omitempty"`
	Type                       *string                                                        `json:"type,omitempty"`
	Upstream                   *ServerlessUpstreamSettings_STATUS                             `json:"upstream,omitempty"`
	Version                    *string                                                        `json:"version,omitempty"`
}

var _ genruntime.ConvertibleStatus = &SignalR_STATUS{}

// ConvertStatusFrom populates our SignalR_STATUS from the provided source
func (signalR *SignalR_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == signalR {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(signalR)
}

// ConvertStatusTo populates the provided destination from our SignalR_STATUS
func (signalR *SignalR_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == signalR {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(signalR)
}

// Storage version of v1api20211001.ManagedIdentity
// A class represent managed identities used for request and response
type ManagedIdentity struct {
	PropertyBag            genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	Type                   *string                       `json:"type,omitempty"`
	UserAssignedIdentities []UserAssignedIdentityDetails `json:"userAssignedIdentities,omitempty"`
}

// Storage version of v1api20211001.ManagedIdentity_STATUS
// A class represent managed identities used for request and response
type ManagedIdentity_STATUS struct {
	PrincipalId            *string                                        `json:"principalId,omitempty"`
	PropertyBag            genruntime.PropertyBag                         `json:"$propertyBag,omitempty"`
	TenantId               *string                                        `json:"tenantId,omitempty"`
	Type                   *string                                        `json:"type,omitempty"`
	UserAssignedIdentities map[string]UserAssignedIdentityProperty_STATUS `json:"userAssignedIdentities,omitempty"`
}

// Storage version of v1api20211001.PrivateEndpointConnection_STATUS_SignalR_SubResourceEmbedded
// A private endpoint connection to an azure resource
type PrivateEndpointConnection_STATUS_SignalR_SubResourceEmbedded struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20211001.ResourceLogConfiguration
// Resource log configuration of a Microsoft.SignalRService resource.
type ResourceLogConfiguration struct {
	Categories  []ResourceLogCategory  `json:"categories,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20211001.ResourceLogConfiguration_STATUS
// Resource log configuration of a Microsoft.SignalRService resource.
type ResourceLogConfiguration_STATUS struct {
	Categories  []ResourceLogCategory_STATUS `json:"categories,omitempty"`
	PropertyBag genruntime.PropertyBag       `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20211001.ResourceSku
// The billing information of the resource.
type ResourceSku struct {
	Capacity    *int                   `json:"capacity,omitempty"`
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Tier        *string                `json:"tier,omitempty"`
}

// Storage version of v1api20211001.ResourceSku_STATUS
// The billing information of the resource.
type ResourceSku_STATUS struct {
	Capacity    *int                   `json:"capacity,omitempty"`
	Family      *string                `json:"family,omitempty"`
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Size        *string                `json:"size,omitempty"`
	Tier        *string                `json:"tier,omitempty"`
}

// Storage version of v1api20211001.ServerlessUpstreamSettings
// The settings for the Upstream when the service is in server-less mode.
type ServerlessUpstreamSettings struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Templates   []UpstreamTemplate     `json:"templates,omitempty"`
}

// Storage version of v1api20211001.ServerlessUpstreamSettings_STATUS
// The settings for the Upstream when the service is in server-less mode.
type ServerlessUpstreamSettings_STATUS struct {
	PropertyBag genruntime.PropertyBag    `json:"$propertyBag,omitempty"`
	Templates   []UpstreamTemplate_STATUS `json:"templates,omitempty"`
}

// Storage version of v1api20211001.SharedPrivateLinkResource_STATUS_SignalR_SubResourceEmbedded
// Describes a Shared Private Link Resource
type SharedPrivateLinkResource_STATUS_SignalR_SubResourceEmbedded struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20211001.SignalRCorsSettings
// Cross-Origin Resource Sharing (CORS) settings.
type SignalRCorsSettings struct {
	AllowedOrigins []string               `json:"allowedOrigins,omitempty"`
	PropertyBag    genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20211001.SignalRCorsSettings_STATUS
// Cross-Origin Resource Sharing (CORS) settings.
type SignalRCorsSettings_STATUS struct {
	AllowedOrigins []string               `json:"allowedOrigins,omitempty"`
	PropertyBag    genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20211001.SignalRFeature
// Feature of a resource, which controls the runtime behavior.
type SignalRFeature struct {
	Flag        *string                `json:"flag,omitempty"`
	Properties  map[string]string      `json:"properties,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Value       *string                `json:"value,omitempty"`
}

// Storage version of v1api20211001.SignalRFeature_STATUS
// Feature of a resource, which controls the runtime behavior.
type SignalRFeature_STATUS struct {
	Flag        *string                `json:"flag,omitempty"`
	Properties  map[string]string      `json:"properties,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Value       *string                `json:"value,omitempty"`
}

// Storage version of v1api20211001.SignalRNetworkACLs
// Network ACLs for the resource
type SignalRNetworkACLs struct {
	DefaultAction    *string                `json:"defaultAction,omitempty"`
	PrivateEndpoints []PrivateEndpointACL   `json:"privateEndpoints,omitempty"`
	PropertyBag      genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	PublicNetwork    *NetworkACL            `json:"publicNetwork,omitempty"`
}

// Storage version of v1api20211001.SignalRNetworkACLs_STATUS
// Network ACLs for the resource
type SignalRNetworkACLs_STATUS struct {
	DefaultAction    *string                     `json:"defaultAction,omitempty"`
	PrivateEndpoints []PrivateEndpointACL_STATUS `json:"privateEndpoints,omitempty"`
	PropertyBag      genruntime.PropertyBag      `json:"$propertyBag,omitempty"`
	PublicNetwork    *NetworkACL_STATUS          `json:"publicNetwork,omitempty"`
}

// Storage version of v1api20211001.SignalROperatorSpec
// Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure
type SignalROperatorSpec struct {
	ConfigMapExpressions []*core.DestinationExpression `json:"configMapExpressions,omitempty"`
	PropertyBag          genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	SecretExpressions    []*core.DestinationExpression `json:"secretExpressions,omitempty"`
	Secrets              *SignalROperatorSecrets       `json:"secrets,omitempty"`
}

// Storage version of v1api20211001.SignalRTlsSettings
// TLS settings for the resource
type SignalRTlsSettings struct {
	ClientCertEnabled *bool                  `json:"clientCertEnabled,omitempty"`
	PropertyBag       genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20211001.SignalRTlsSettings_STATUS
// TLS settings for the resource
type SignalRTlsSettings_STATUS struct {
	ClientCertEnabled *bool                  `json:"clientCertEnabled,omitempty"`
	PropertyBag       genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20211001.SystemData_STATUS
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

// Storage version of v1api20211001.NetworkACL
// Network ACL
type NetworkACL struct {
	Allow       []string               `json:"allow,omitempty"`
	Deny        []string               `json:"deny,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20211001.NetworkACL_STATUS
// Network ACL
type NetworkACL_STATUS struct {
	Allow       []string               `json:"allow,omitempty"`
	Deny        []string               `json:"deny,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20211001.PrivateEndpointACL
// ACL for a private endpoint
type PrivateEndpointACL struct {
	Allow       []string               `json:"allow,omitempty"`
	Deny        []string               `json:"deny,omitempty"`
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20211001.PrivateEndpointACL_STATUS
// ACL for a private endpoint
type PrivateEndpointACL_STATUS struct {
	Allow       []string               `json:"allow,omitempty"`
	Deny        []string               `json:"deny,omitempty"`
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20211001.ResourceLogCategory
// Resource log category configuration of a Microsoft.SignalRService resource.
type ResourceLogCategory struct {
	Enabled     *string                `json:"enabled,omitempty"`
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20211001.ResourceLogCategory_STATUS
// Resource log category configuration of a Microsoft.SignalRService resource.
type ResourceLogCategory_STATUS struct {
	Enabled     *string                `json:"enabled,omitempty"`
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20211001.SignalROperatorSecrets
type SignalROperatorSecrets struct {
	PrimaryConnectionString   *genruntime.SecretDestination `json:"primaryConnectionString,omitempty"`
	PrimaryKey                *genruntime.SecretDestination `json:"primaryKey,omitempty"`
	PropertyBag               genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	SecondaryConnectionString *genruntime.SecretDestination `json:"secondaryConnectionString,omitempty"`
	SecondaryKey              *genruntime.SecretDestination `json:"secondaryKey,omitempty"`
}

// Storage version of v1api20211001.UpstreamTemplate
// Upstream template item settings. It defines the Upstream URL of the incoming requests.
// The template defines the pattern
// of the event, the hub or the category of the incoming request that matches current URL template.
type UpstreamTemplate struct {
	Auth            *UpstreamAuthSettings  `json:"auth,omitempty"`
	CategoryPattern *string                `json:"categoryPattern,omitempty"`
	EventPattern    *string                `json:"eventPattern,omitempty"`
	HubPattern      *string                `json:"hubPattern,omitempty"`
	PropertyBag     genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	UrlTemplate     *string                `json:"urlTemplate,omitempty"`
}

// Storage version of v1api20211001.UpstreamTemplate_STATUS
// Upstream template item settings. It defines the Upstream URL of the incoming requests.
// The template defines the pattern
// of the event, the hub or the category of the incoming request that matches current URL template.
type UpstreamTemplate_STATUS struct {
	Auth            *UpstreamAuthSettings_STATUS `json:"auth,omitempty"`
	CategoryPattern *string                      `json:"categoryPattern,omitempty"`
	EventPattern    *string                      `json:"eventPattern,omitempty"`
	HubPattern      *string                      `json:"hubPattern,omitempty"`
	PropertyBag     genruntime.PropertyBag       `json:"$propertyBag,omitempty"`
	UrlTemplate     *string                      `json:"urlTemplate,omitempty"`
}

// Storage version of v1api20211001.UserAssignedIdentityDetails
// Information about the user assigned identity for the resource
type UserAssignedIdentityDetails struct {
	PropertyBag genruntime.PropertyBag       `json:"$propertyBag,omitempty"`
	Reference   genruntime.ResourceReference `armReference:"Reference" json:"reference,omitempty"`
}

// Storage version of v1api20211001.UserAssignedIdentityProperty_STATUS
// Properties of user assigned identity.
type UserAssignedIdentityProperty_STATUS struct {
	ClientId    *string                `json:"clientId,omitempty"`
	PrincipalId *string                `json:"principalId,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20211001.UpstreamAuthSettings
// Upstream auth settings. If not set, no auth is used for upstream messages.
type UpstreamAuthSettings struct {
	ManagedIdentity *ManagedIdentitySettings `json:"managedIdentity,omitempty"`
	PropertyBag     genruntime.PropertyBag   `json:"$propertyBag,omitempty"`
	Type            *string                  `json:"type,omitempty"`
}

// Storage version of v1api20211001.UpstreamAuthSettings_STATUS
// Upstream auth settings. If not set, no auth is used for upstream messages.
type UpstreamAuthSettings_STATUS struct {
	ManagedIdentity *ManagedIdentitySettings_STATUS `json:"managedIdentity,omitempty"`
	PropertyBag     genruntime.PropertyBag          `json:"$propertyBag,omitempty"`
	Type            *string                         `json:"type,omitempty"`
}

// Storage version of v1api20211001.ManagedIdentitySettings
// Managed identity settings for upstream.
type ManagedIdentitySettings struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Resource    *string                `json:"resource,omitempty"`
}

// Storage version of v1api20211001.ManagedIdentitySettings_STATUS
// Managed identity settings for upstream.
type ManagedIdentitySettings_STATUS struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Resource    *string                `json:"resource,omitempty"`
}

func init() {
	SchemeBuilder.Register(&SignalR{}, &SignalRList{})
}
